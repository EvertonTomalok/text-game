package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func runCmd(name string, arg ...string) {
	command := exec.Command(name, arg...)
	command.Stdout = os.Stdout
	_ = command.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
