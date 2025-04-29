package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: unify \"<installer-path>\" [params...]")
		fmt.Println("Description: Unifies the installation process for .msi and .exe installers.")
		fmt.Println("             Parameters are passed directly to the installer.")
		fmt.Println("             The installer type is determined by the file extension.")
		fmt.Println("             A log file is automatically generated for .msi installers.")
		fmt.Println("             Do not include a log file name in the params, it will be generated automatically.")
		fmt.Println("Note: Pass each parameter as a separate argument. Do NOT quote multiple parameters together.")
		fmt.Println("\nDefaults:")
		fmt.Println("  .msi: /qn /l*v install.log")
		fmt.Println("  .exe: /S")
		fmt.Println("\nExamples:")
		fmt.Println("  Install .msi silently with no restart and custom directory:")
		fmt.Println("    .\\unify.exe \"C:\\path\\to\\installer.msi\" /norestart INSTALLDIR=\"C:\\custom\\dir\"")
		fmt.Println("  Install .exe silently with no restart and custom directory:")
		fmt.Println("    .\\unify.exe \"C:\\path\\to\\installer.exe\" /norestart /D=C:\\custom\\dir")
		os.Exit(1)
	}
	installerPath := os.Args[1]
	params := os.Args[2:]

	output, exitCode, err := InstallApp(installerPath, params)
	fmt.Printf("Exit code: %d\n", exitCode)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Output:\n", output)
}
