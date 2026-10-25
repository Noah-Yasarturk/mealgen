# Run the Go webserver
Write-Host "Running mealgen app..." -ForegroundColor Green -BackgroundColor Black

# Set env vars
$env:PORT = "8080"
$env:SQLITE_FILEPATH = "file:../db/mealgen.db"

# Run the app
$cmdDir = Join-Path $PSScriptRoot 'cmd'
$binDir = Join-Path $PSScriptRoot 'bin'
$exe    = Join-Path $binDir 'mealgen.exe'
New-Item -ItemType Directory -Force -Path $binDir | Out-Null
go -C $cmdDir build -o $exe .
if ($LASTEXITCODE -eq 0) { & $exe }
