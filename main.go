package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCommand(s string) [2]string {
	var cmd *exec.Cmd
	cmd = exec.Command("getdata.exe", s)
	// cmd.Stdin = strings.NewReader(s)
	// stdoutStderr, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var out bytes.Buffer
	cmd.Stdout = &out
	var err bytes.Buffer
	cmd.Stderr = &err
	cmd.Run()
	var result [2]string
	result[0] = string(out.String())
	result[1] = string(err.String())

	return result
}

func main() {
	result := RunCommand("021FI_004.pv")
	fmt.Println("QR code:", result[0])
}
