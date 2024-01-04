package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func main() {
	for {
		percent, err := cpu.Percent(time.Second, false)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("CPU Usage: %.2f%%\n", percent[0])
		if percent[0] > 50 {
			fmt.Println("CPU Usage is above 50%")
			// send email
		} else if percent[0] > 50 {
			fmt.Println("CPU Usage is less then 50%")
			fmt.Println("you are good to go for deploy")
		}
		time.Sleep(2 * time.Second)
	}
}
