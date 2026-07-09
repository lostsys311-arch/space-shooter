$src = "C:\Users\sylwe\OneDrive\Documents"
$dst = "D:\"
$projects = @("space-shooter","tibia-android-game","tibia-apk","PixelDungeonRemastered","amnesic-wipe","archlive")

foreach ($p in $projects) {
    $sp = Join-Path $src $p
    $dp = Join-Path $dst $p
    if (Test-Path $sp) {
        Copy-Item -Path $sp -Destination $dst -Recurse -Force
        Write-Output "Copied $p"
    }
}

$files = @("bomb.py")
foreach ($f in $files) {
    $sf = Join-Path $src $f
    if (Test-Path $sf) {
        Copy-Item -Path $sf -Destination $dst -Force
        Write-Output "Copied $f"
    }
}

Write-Output "Backup complete!"
