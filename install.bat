@echo off
setlocal enabledelayedexpansion

echo ========================================
echo  Task CLI Go - Installer
echo ========================================
echo.

cd /d "%~dp0"

:: Try to build from source if Go is available
where go >nul 2>&1
if %ERRORLEVEL% equ 0 (
    echo [1/2] Building task-cli.exe from source...
    go build -o task-cli.exe .
    if %ERRORLEVEL% neq 0 (
        echo ERROR: Build failed.
        pause
        exit /b 1
    )
    echo  OK - Built task-cli.exe
) else (
    echo [1/2] Go not found, using pre-built task-cli.exe...
    if not exist "task-cli.exe" (
        echo ERROR: task-cli.exe not found and Go is not installed.
        pause
        exit /b 1
    )
    echo  OK - Using existing task-cli.exe
)
echo.

:: If this is the project folder (has source), just build locally
if exist "go.mod" (
    echo [2/2] Project detected. Binary is ready in the current folder.
    echo.
    echo ========================================
    echo  Installation complete!
    echo.
    echo  Use .\task-cli add "my task" in this folder.
    echo  Data will be saved to task.json in this folder.
    echo ========================================
    pause
    exit /b
)

:: Standalone install - copy to task-cli-bin and add to PATH
set "INSTALL_DIR=%USERPROFILE%\task-cli-bin"
echo [2/2] Standalone install. Copying to %INSTALL_DIR%...
if not exist "%INSTALL_DIR%" mkdir "%INSTALL_DIR%"
copy /y "task-cli.exe" "%INSTALL_DIR%\task-cli.exe" >nul

echo "%PATH%" | find /i "%INSTALL_DIR%" >nul
if %ERRORLEVEL% neq 0 (
    setx PATH "%INSTALL_DIR%;%PATH%"
    echo  OK - Added to PATH. Restart your terminal.
) else (
    echo  OK - Already in PATH.
)
echo.

echo ========================================
echo  Installation complete!
echo.
echo  Usage: task-cli add "my task"
echo  Data will be saved to %INSTALL_DIR%\task.json
echo ========================================
pause
