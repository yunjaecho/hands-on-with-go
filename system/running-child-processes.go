package system

import (
	"fmt"
	"os/exec"
)

func init() {
	fmt.Println("===== Running Child Processes =====")

	//lsCommand := exec.Command("ls")
	lsCommand := exec.Command("ls", "-al")
	output, _ := lsCommand.Output()
	fmt.Println(string(output))
	lsCommand.Start()
	lsCommand.Run()
	fmt.Println(lsCommand.Process.Pid)
}
