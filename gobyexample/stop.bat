@echo off
set params=EmsMonitorApplication
for /f "usebackq tokens=1-2" %%a in (`jps -l ^| findstr %params%`) do (
    set pid=%%a
)
if not defined pid (
    exit /b 1
)else (
taskkill /f /pid %pid%
)