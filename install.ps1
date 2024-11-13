# Проверка, запущен ли скрипт с правами администратора
if (-not ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Host "Пожалуйста, запустите скрипт от имени администратора."
    exit
}

# Определение путей
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$SourceBinary = Join-Path $ScriptDir "recall.exe"
$SourceScript = Join-Path $ScriptDir "recall.ps1"
$destinationFolder = "C:\Program Files\recall"
$destination = "$destinationFolder\recall.exe"

# Определение пути к папке PowerShell Scripts
$PowerShellScriptsFolder = Join-Path ([Environment]::GetFolderPath("MyDocuments")) "PowerShell\Scripts"
$ScriptDestination = Join-Path $PowerShellScriptsFolder "recall.ps1"

# Создание папки для установки recall.exe, если она не существует
if (-Not (Test-Path -Path $destinationFolder)) {
    New-Item -Path $destinationFolder -ItemType Directory -Force
    Write-Host "Создана папка: $destinationFolder"
}

# Создание папки PowerShell Scripts, если она не существует
if (-Not (Test-Path -Path $PowerShellScriptsFolder)) {
    New-Item -Path $PowerShellScriptsFolder -ItemType Directory -Force
    Write-Host "Создана папка: $PowerShellScriptsFolder"
}

# Копирование файлов
Copy-Item -Path $SourceBinary -Destination $destination -Force
Write-Host "Скопирован файл: recall.exe"

Copy-Item -Path $SourceScript -Destination $ScriptDestination -Force
Write-Host "Скопирован файл: recall.ps1"

# Обновленный способ добавления в PATH
$currentPath = [Environment]::GetEnvironmentVariable('Path', [EnvironmentVariableTarget]::Machine)
$pathArray = $currentPath -split ';' | Where-Object { $_ -ne '' }

if ($pathArray -notcontains $destinationFolder) {
    $newPath = ($pathArray + $destinationFolder) -join ';'
    try {
        [Environment]::SetEnvironmentVariable('Path', $newPath, [EnvironmentVariableTarget]::Machine)
        Write-Host "Путь успешно добавлен в системную переменную PATH."
    } catch {
        Write-Host "Ошибка при добавлении пути в PATH: $_"
    }
} else {
    Write-Host "Путь уже существует в системной переменной PATH."
}

# Проверка успешности установки
$installationSuccess = $true
if (-not (Test-Path $destination)) {
    Write-Host "Ошибка: recall.exe не был установлен корректно." -ForegroundColor Red
    $installationSuccess = $false
}
if (-not (Test-Path $ScriptDestination)) {
    Write-Host "Ошибка: recall.ps1 не был установлен корректно." -ForegroundColor Red
    $installationSuccess = $false
}

if ($installationSuccess) {
    Write-Host "`nУстановка успешно завершена!" -ForegroundColor Green
    Write-Host "recall.exe установлен в: $destinationFolder"
    Write-Host "recall.ps1 установлен в: $PowerShellScriptsFolder"
    Write-Host "Вы можете использовать 'recall' из командной строки после перезапуска PowerShell."
} else {
    Write-Host "`nУстановка завершена с ошибками. Пожалуйста, проверьте сообщения выше." -ForegroundColor Red
}