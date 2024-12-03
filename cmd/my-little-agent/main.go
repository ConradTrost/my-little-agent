package main

import (
	"log"
	"net/http"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	"github.com/gin-gonic/gin"
)

func getLoadAvg() *types.LoadAverageInfo {
	process, err := sysinfo.Self()
	if err != nil {
		log.Fatal(err)
	}

	loadAvg := &types.LoadAverageInfo{}
	if hello, ok := process.(types.LoadAverage); ok {
		loadAvg, err := hello.LoadAverage()
		if err != nil {
			log.Fatal(err)
		}
		return loadAvg
	}

	return loadAvg
}

func getOpenHandles() int {
	process, err := sysinfo.Self()
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	if handleCounter, ok := process.(types.OpenHandleCounter); ok {
		count, err := handleCounter.OpenHandleCount()
		if err != nil {
			log.Fatal(err)
		}
		return count
	}

	return count
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/info", func(c *gin.Context) {
		host, _ := sysinfo.Host()

		cpu, _ := host.CPUTime()
		memory, _ := host.Memory()
		c.JSON(http.StatusOK, gin.H{
			"count":  getOpenHandles(),
			"load":   *getLoadAvg(),
			"host":   host.Info(),
			"cpu":    cpu,
			"memory": *memory,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
