# Space Shooter

A 2D space shooter built with [Go](https://go.dev/) and [Ebitengine](https://ebitengine.org/).

## How to Play

| Key | Action |
|-----|--------|
| ← → / A D | Move |
| Space | Shoot |
| Space / Enter | Restart on game over |

Shoot enemies before they reach the bottom. Speed increases as your score grows.

## Build & Run

```bash
go run .
```

Or compile a binary:

```bash
go build -o space-shooter .
```

## CI Builds

Pushing to `master` triggers a [GitHub Actions](https://github.com/lostsys311-arch/space-shooter/actions) workflow that builds for:

| Platform | Format |
|----------|--------|
| **Linux** | `space-shooter-linux.tar.gz` |
| **Windows** | `space-shooter-windows.zip` (contains `space-shooter.exe`) |
| **macOS** | `space-shooter-macos.dmg` |
| **Android** | `space-shooter.apk` |

Download the artifact for your platform from the Actions tab.
