package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	ipcmd := exec.Command("ipconfig")
	ipOutput, err := ipcmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ipOutput))

	grepCmd := exec.Command("ipconfig")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
