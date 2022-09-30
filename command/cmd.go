package command

import (
  "os/exec"
)

func Cmd(cmd string, shell bool) []byte {

  if shell {
    out, err := exec.Command("/bin/sh", "-c", cmd).Output()
    if err != nil {
      panic(err)
    }
    return out
  } else {
    out, err := exec.Command(cmd).Output()
    if err != nil {
      panic(err)
    }
    return out
  }
}
