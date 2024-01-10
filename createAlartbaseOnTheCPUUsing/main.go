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
func printResourceUsage(label string) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("%s:\n", label)
	fmt.Println("timeStamp: ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("  Memory Usage: %v MB\n", bToMb(memStats.Alloc))
	fmt.Printf("  Total Allocated Memory: %v MB\n", bToMb(memStats.TotalAlloc))
	fmt.Printf("  CPU Usage: %v\n", getCpuUsage())

	fmt.Println("----------------------")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func getCpuUsage() string {
	cpuStats := runtime.NumCPU()
	return fmt.Sprintf("%d cores", cpuStats)
}
