package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func setCPUOnline(cpu int) (string, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo 1 | sudo tee /sys/devices/system/cpu/cpu%d/online", cpu))
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return "", nil
}
func main() {
	// For example, to adjust the CPU frequency using cpupower (not actually overclocking)
	// THIS IS JUST AN EXAMPLE. PLEASE BE CAREFUL.
	err := ExecuteCommand("sudo", "apt", "update")
	if err != nil {
		fmt.Println("Error:", err)
	}
	err1 := ExecuteCommand("sudo", "apt", "install", "linux-tools-common", "linux-tools-generic", "-y")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}
	for i := 12; i <= 15; i++ {
		errOutput, err := setCPUOnline(i)
		if err != nil {
			fmt.Printf("Failed to set CPU%d online: %s\nError Output: %s\n", i, err, errOutput)
		} else {
			fmt.Printf("CPU%d is now online\n", i)
		}
	}
	err2 := ExecuteCommand("sudo", "cpupower", "frequency-set", "--max", "4.00GHz")
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
}
