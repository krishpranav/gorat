package shell

import (
	"os/exec"
	"syscall"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

func GetShell() *exec.Cmd {
	cmd := exec.Command("C:\\Windows\\SysWOW64\\WindowsPowerShell\\v1.0\\powershell.exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
