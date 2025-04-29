package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func InstallApp(installerPath string, params []string) (string, int, error) {
	ext := strings.ToLower(filepath.Ext(installerPath))
	var cmd *exec.Cmd

	switch ext {
	case ".msi":
		quiet := false
		logParamIdx := -1
		logFile := "install.log"
		for i, p := range params {
			if strings.EqualFold(p, "/qn") || strings.EqualFold(p, "/quiet") {
				quiet = true
			}
			if strings.HasPrefix(strings.ToLower(p), "/l*v") {
				logParamIdx = i
				parts := strings.Fields(p)
				if len(parts) == 2 {
					logFile = parts[1]
				} else if len(params) > i+1 {
					logFile = params[i+1]
				}
			}
		}

		args := []string{"/i", installerPath}
		if !quiet {
			args = append(args, "/qn")
		}
		if logParamIdx == -1 {
			args = append(args, "/l*v", logFile)
		}
		args = append(args, params...)

		cmd = exec.Command("msiexec", args...)
		output, err := cmd.CombinedOutput()
		exitCode := -1
		if cmd.ProcessState != nil {
			exitCode = cmd.ProcessState.ExitCode()
		}

		logContent, logErr := os.ReadFile(logFile)
		if logErr == nil {
			fmt.Printf("%s content:\n", logFile)
			fmt.Println(string(logContent))
		} else {
			fmt.Printf("Error reading %s: %v\n", logFile, logErr)
		}

		return string(output), exitCode, err

	case ".exe":
		silent := false
		for _, p := range params {
			if strings.EqualFold(p, "/S") {
				silent = true
				break
			}
		}
		args := []string{}
		if !silent {
			args = append(args, "/S")
		}
		args = append(args, params...)
		cmd = exec.Command(installerPath, args...)

	default:
		return "", -1, fmt.Errorf("unsupported installer type: %s", ext)
	}

	output, err := cmd.CombinedOutput()
	exitCode := -1
	if cmd.ProcessState != nil {
		exitCode = cmd.ProcessState.ExitCode()
	}
	return string(output), exitCode, err
}
