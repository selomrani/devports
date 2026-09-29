Write-Host "🚀 Installing devports for Windows..." -ForegroundColor Cyan

$VERSION = "v1.0.0"
$REPO = "selomrani/devports"
$FILE = "devports-windows-amd64.zip"
$URL = "https://github.com/$REPO/releases/download/$VERSION/$FILE"

$InstallDir = "$env:LOCALAPPDATA\devports"
$ZipPath = "$InstallDir\$FILE"

# Create directory
if (!(Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Force -Path $InstallDir | Out-Null
}

Write-Host "⬇️ Downloading $VERSION..."
Invoke-WebRequest -Uri $URL -OutFile $ZipPath

Write-Host "📦 Extracting..."
Expand-Archive -Path $ZipPath -DestinationPath $InstallDir -Force

Write-Host "🔧 Configuring..."
if (Test-Path "$InstallDir\devports-windows-amd64.exe") {
    Rename-Item -Path "$InstallDir\devports-windows-amd64.exe" -NewName "devports.exe" -Force
}

Write-Host "🧹 Cleaning up..."
Remove-Item -Path $ZipPath -Force

# Add to User PATH if it doesn't exist
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notmatch [regex]::Escape($InstallDir)) {
    Write-Host "⚙️ Adding devports to your system PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$InstallDir", "User")
    $NeedsRestart = $true
} else {
    $NeedsRestart = $false
}

Write-Host "`n✅ devports installed successfully! 🎉" -ForegroundColor Green
if ($NeedsRestart) {
    Write-Host "⚠️  IMPORTANT: Please restart your terminal (or open a new tab) so the PATH updates take effect." -ForegroundColor Yellow
}
Write-Host "👉 Just type 'devports' in your terminal to launch it.`n" -ForegroundColor Cyan
