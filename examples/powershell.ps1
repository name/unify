# Unify Installer Wrapper - PowerShell Usage Example
#
# Place unify.exe and your installer in the same directory as this script.
# For Intune Win32 packages, include unify.exe in the package root.
#
# IMPORTANT:
# - Pass each installer parameter as a separate argument.
# - Do NOT quote multiple parameters together (e.g., do NOT use "/norestart /D=..." as one string).
#
# --- Default Arguments ---
# - For .msi files: /qn /l*v install.log are added by default unless overridden.
# - For .exe files: /S is added by default unless overridden.
#
# --- MSI Installer Example ---
# Installs the MSI silently, prevents restart, and sets a custom install directory.
Start-Process -FilePath ".\unify.exe" `
  -ArgumentList "C:\Path\To\Installer.msi", "/norestart", 'INSTALLDIR="C:\Program Files\MyApp"' `
  -Wait -NoNewWindow

# --- EXE Installer Example ---
# Installs the EXE silently, prevents restart, and sets a custom install directory.
Start-Process -FilePath ".\unify.exe" `
  -ArgumentList "C:\Path\To\Installer.exe", "/norestart", "/D=C:\Program Files\MyApp" `
  -Wait -NoNewWindow

# Output from both MSI and EXE installers will be shown in the console.
# This allows you to capture logs or output for troubleshooting or reporting.
