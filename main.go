package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
	playerSpeed  = 5
	bulletSpeed  = 8
	enemySpeed   = 2
)

type Player struct {
	x, y   float64
	width  float64
	height float64
}

type Bullet struct {
	x, y  float64
	w, h  float64
	alive bool
}

type Enemy struct {
	x, y  float64
	w, h  float64
	alive bool
}

type Game struct {
	player    Player
	bullets   []*Bullet
	enemies   []*Enemy
	score     int
	gameOver  bool
	spawnCD   int
}

func (g *Game) Layout(outW, outH int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) || inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.reset()
		}
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.x -= playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.x += playerSpeed
	}
	g.player.x = clamp(g.player.x, 0, screenWidth-g.player.width)

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.shoot()
	}

	for _, b := range g.bullets {
		if b.alive {
			b.y -= bulletSpeed
			if b.y+b.h < 0 {
				b.alive = false
			}
		}
	}

	for _, e := range g.enemies {
		if e.alive {
			e.y += enemySpeed
			if e.y > screenHeight {
				e.alive = false
				g.gameOver = true
			}
		}
	}

	for _, b := range g.bullets {
		if !b.alive {
			continue
		}
		for _, e := range g.enemies {
			if !e.alive {
				continue
			}
			if rectOverlap(b.x, b.y, b.w, b.h, e.x, e.y, e.w, e.h) {
				b.alive = false
				e.alive = false
				g.score++
			}
		}
	}

	for _, e := range g.enemies {
		if !e.alive {
			continue
		}
		if rectOverlap(g.player.x, g.player.y, g.player.width, g.player.height, e.x, e.y, e.w, e.h) {
			g.gameOver = true
		}
	}

	g.spawnCD--
	if g.spawnCD <= 0 {
		g.spawnEnemy()
		g.spawnCD = 30 - int(math.Min(float64(g.score)/10, 20))
		if g.spawnCD < 8 {
			g.spawnCD = 8
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{10, 10, 30, 255})

	ebitenutil.DrawRect(screen, g.player.x, g.player.y, g.player.width, g.player.height, color.RGBA{0, 200, 255, 255})
	ebitenutil.DrawRect(screen, g.player.x+4, g.player.y-6, g.player.width-8, 6, color.RGBA{100, 200, 255, 255})

	for _, b := range g.bullets {
		if b.alive {
			ebitenutil.DrawRect(screen, b.x, b.y, b.w, b.h, color.RGBA{255, 255, 100, 255})
		}
	}

	for _, e := range g.enemies {
		if e.alive {
			ebitenutil.DrawRect(screen, e.x, e.y, e.w, e.h, color.RGBA{255, 60, 60, 255})
			ebitenutil.DrawRect(screen, e.x+3, e.y-4, e.w-6, 4, color.RGBA{200, 80, 80, 255})
		}
	}

	ebitenutil.DebugPrint(screen, "Score: "+itoa(g.score))

	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "GAME OVER", screenWidth/2-40, screenHeight/2-20)
		ebitenutil.DebugPrintAt(screen, "Press SPACE or ENTER to restart", screenWidth/2-80, screenHeight/2+10)
	}
}

func (g *Game) shoot() {
	b := &Bullet{
		x:     g.player.x + g.player.width/2 - 2,
		y:     g.player.y - 10,
		w:     4,
		h:     10,
		alive: true,
	}
	g.bullets = append(g.bullets, b)
}

func (g *Game) spawnEnemy() {
	w := 20.0 + rand.Float64()*20
	h := 15.0 + rand.Float64()*15
	e := &Enemy{
		x:     rand.Float64() * (screenWidth - w),
		y:     -h,
		w:     w,
		h:     h,
		alive: true,
	}
	g.enemies = append(g.enemies, e)
}

func (g *Game) reset() {
	g.player = Player{x: screenWidth/2 - 15, y: screenHeight - 50, width: 30, height: 20}
	g.bullets = nil
	g.enemies = nil
	g.score = 0
	g.gameOver = false
	g.spawnCD = 30
}

func main() {
	g := &Game{}
	g.reset()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Space Shooter")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func clamp(v, min, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func rectOverlap(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	return x1 < x2+w2 && x1+w1 > x2 && y1 < y2+h2 && y1+h1 > y2
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var buf [12]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
