@echo off
setlocal enabledelayedexpansion

set NGINX_HOME=D:\nginx-1.16.1

cd /d %NGINX_HOME%
nginx -s stop && start nginx.exe

timeout /t 3 /nobreak >nul
tasklist | find "nginx.exe" >nul
