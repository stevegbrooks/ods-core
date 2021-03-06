package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func RunScriptFromBaseDir(command string, args []string, envVars []string) (string, string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	return RunCommand(fmt.Sprintf("%s/%s", dir, command), args, envVars)
}

func RunCommand(command string, args []string, envVars []string) (string, string, error) {
	cmd := exec.Command(command, args...)
	cmd.Env = append(os.Environ(), envVars...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

func RunCommandWithWorkDir(command string, args []string, workDir string, envVars []string) (string, string, error) {
	cmd := exec.Command(command, args...)
	cmd.Env = append(os.Environ(), envVars...)
	cmd.Dir = workDir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
