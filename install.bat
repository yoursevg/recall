@echo off
chcp 65001 > nul
powershell -ExecutionPolicy Bypass -Command "Start-Process powershell -ArgumentList '-ExecutionPolicy Bypass -File ""%~dp0install.ps1""' -Verb RunAs"
pause