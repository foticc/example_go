@echo off
set params=6379
for /f "usebackq tokens=5-6" %%a in (`netstat -ano ^| findstr %params%`) do (
    set pid=%%a
)
if not defined pid (
    exit /b 1
)else (
taskkill /f /pid %pid%
)