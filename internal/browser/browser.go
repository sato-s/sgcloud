package browser

import (
	"os"
	"os/exec"
	"runtime"
)

// ref: https://gist.github.com/hyg/9c4afcd91fe24316cbf0
func OpenBrowser(url string) error {

	switch runtime.GOOS {
	case "linux":
		if isWSL() {
			return exec.Command("cmd.exe", "/C", "start", url).Start()
		} else {
			return exec.Command("xdg-open", url).Start()
		}
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	}
	return nil
}

// https://superuser.com/a/1749811
func isWSL() bool {
	const wslFile = "/proc/sys/fs/binfmt_misc/WSLInterop"
	_, err := os.Stat(wslFile)
	return err == nil
}
