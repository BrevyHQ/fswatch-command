package command

import (
  "bytes"
  "fmt"
  "os/exec"
)

func Cmd(cmd string, shell bool) {

  if shell {
    cmd := exec.Command("/bin/sh", "-c", cmd)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
      fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
      return
    }
  } else {
    cmd := exec.Command(cmd)
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
      fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
      return
    }
  }
}
