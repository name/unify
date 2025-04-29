import subprocess

# Unify Installer Wrapper - Python Usage Example
# Place unify.exe and your installer in the same directory as this script.

# --- MSI Installer Example ---
# Installs the MSI silently, prevents restart, and sets a custom install directory.
msi_args = [
    "unify.exe",
    r"C:\Path\To\Installer.msi",
    "/norestart",
    'INSTALLDIR="C:\\Program Files\\MyApp"'
]
subprocess.run(msi_args, check=True)

# --- EXE Installer Example ---
# Installs the EXE silently, prevents restart, and sets a custom install directory.
exe_args = [
    "unify.exe",
    r"C:\Path\To\Installer.exe",
    "/norestart",
    "/D=C:\\Program Files\\MyApp"
]
subprocess.run(exe_args, check=True)

# Output from both MSI and EXE installers will be shown in the console.
