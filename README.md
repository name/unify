# Unify Installer Wrapper

A simple Go utility to standardize silent installations for `.msi` and `.exe` installers. Automatically adds appropriate silent and logging parameters and prints the install log to the console.

## Download

If you do not want to build from source, download the latest release from GitHub:

- [Download unify.exe v1.0.0](https://github.com/name/unify/releases/download/v1.0.0/unify.exe)

Place `unify.exe` in the same folder as your installer or add it to your system PATH.

## Build

1. Install [Go](https://golang.org/dl/).
2. Clone this repository.
3. Build the executable:

   ```sh
   go build -o unify.exe
   ```

## Usage

```sh
unify.exe "<installer-path>" [params...]
```

- Do **not** quote multiple parameters together.
- For `.msi` files, `/qn /l*v install.log` are added by default unless overridden.
- For `.exe` files, `/S` is added by default unless overridden.

### Examples

Install an MSI silently with no restart and custom directory:

```sh
unify.exe "C:\path\to\installer.msi" /norestart INSTALLDIR="C:\custom\dir"
```

Install an EXE silently with no restart and custom directory:

```sh
unify.exe "C:\path\to\installer.exe" /norestart /D=C:\custom\dir
```

## Using in Intune with PowerShell

You can call `unify.exe` from a PowerShell script in your Intune Win32 package:

```powershell
Start-Process -FilePath ".\unify.exe" -ArgumentList "C:\Path\To\Installer.msi", '/norestart', 'INSTALLDIR="C:\Program Files\MyApp"' -Wait -NoNewWindow
```

Or for an EXE installer:

```powershell
Start-Process -FilePath ".\unify.exe" -ArgumentList "C:\Path\To\Installer.exe", '/norestart', '/D=C:\Program Files\MyApp' -Wait -NoNewWindow
```

**Tip:** Place `unify.exe` and your installer in the same folder, or provide the full path to each.

---
