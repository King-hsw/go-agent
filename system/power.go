package system

import (
	"os/exec"
	"runtime"
	"time"
)

// Shutdown 关闭电脑
func Shutdown() error {

	switch runtime.GOOS {

	case "windows":

		// Windows关机
		cmd := exec.Command(
			"shutdown",
			"/s",
			"/t",
			"0",
		)

		return cmd.Run()

	case "linux", "darwin":
		time.Sleep(60 * time.Second)
		// Linux关机
		cmd := exec.Command(
			"shutdown",
			"-h",
			"now",
		)

		return cmd.Run()

	default:

		return nil
	}
}
