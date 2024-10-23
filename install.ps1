# Проверка, запущен ли скрипт с правами администратора
if (-not ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)) {
    Write-Host "Пожалуйста, запустите скрипт от имени администратора."
    exit
}

# Определение пути к исходному бинарному файлу
$ScriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$SourceBinary = Join-Path $ScriptDir "recall.exe"
$destinationFolder = "C:\Program Files\recall"
$destination = "$destinationFolder\recall.exe"

# Создание папки для установки, если она не существует
if (-Not (Test-Path -Path $destinationFolder)) {
    New-Item -Path $destinationFolder -ItemType Directory -Force
}

# Копирование файла в папку назначения
Copy-Item -Path $SourceBinary -Destination $destination -Force

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

Write-Host "Установка завершена! Вы можете использовать 'recall' из командной строки после перезапуска PowerShell."