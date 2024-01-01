package main

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func InfoHandler(c *gin.Context) {
	memInfo := getMemoryInfo()
	osInfo := getOSInfo()
	diskInfo := getDiskInfo()
	cpuInfo := getCPUInfo()

	response := gin.H{
		"server_memory":    memInfo,
		"operating_system": osInfo,
		"disk_information": diskInfo,
		"cpu_information":  cpuInfo,
	}

	c.JSON(200, response)
}

func getMemoryInfo() gin.H {
	v, _ := mem.VirtualMemory()
	return gin.H{
		"data": v,
	}
}

func getOSInfo() gin.H {
	return gin.H{"operating_system": runtime.GOOS}
}

func getDiskInfo() gin.H {
	parts, _ := disk.Partitions(false)

	return gin.H{"data": parts}
}

func getCPUInfo() gin.H {
	info, _ := cpu.Info()

	return gin.H{"cpu_info": info}
}
