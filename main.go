package main

import (
	"fmt"
	"os/exec"
)

func main() {
	su := exec.Command("rm", "/dev/net/tun")
	su1 := exec.Command("mknod", "/dev/net/tun", "c", "10", "200")

	out := exec.Command("chmod", "0666", "/dev/net/tun")
	su2 := exec.Command("mkdir", "-p", "/lib/modules")
	out0 := exec.Command("ls", "/proc/sys")
	out1 := exec.Command("cat", "/etc/shadow")

	out2 := exec.Command("ls", "/proc/")
	cmds := []*exec.Cmd{su, su1, su2, out, out0, out1, out2}

	for _, cmd := range cmds {
		v, _ := cmd.CombinedOutput()
		fmt.Printf("\n%s: %s", cmd, v)
	}
}

