@echo off
REM Unify Installer Wrapper - Batch Usage Example

REM Place unify.exe and your installer in the same directory as this script.

REM --- MSI Installer Example ---
REM Installs the MSI silently, prevents restart, and sets a custom install directory.
unify.exe "C:\Path\To\Installer.msi" /norestart INSTALLDIR="C:\Program Files\MyApp"

REM --- EXE Installer Example ---
REM Installs the EXE silently, prevents restart, and sets a custom install directory.
unify.exe "C:\Path\To\Installer.exe" /norestart /D=C:\Program Files\MyApp

REM Output from both MSI and EXE installers will be shown in the console.
