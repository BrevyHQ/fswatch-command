package command

import (
	"fmt"
	"os/exec"
)

func Cmd(cmd string, shell bool) []byte {

	if shell {
		out, err := exec.Command("/bin/sh", "-c", cmd).Output()
		if err != nil {
			panic("some error found")
		}
		fmt.Println(string(out))
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}
