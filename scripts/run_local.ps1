param(
    [switch]$Refresh
)

if ($Refresh) {
    Write-Host "Composing with fresh images..."
    docker compose -f compose_local.yml build --no-cache
    docker compose -f compose_local.yml up --force-recreate
} else {
    Write-Host "Leveraging Docker cache to compose..."
    docker compose -f compose_local.yml build
    docker compose -f compose_local.yml up
}
