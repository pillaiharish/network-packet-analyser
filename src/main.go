package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"network-packet-analyser/utils"

	"github.com/gin-gonic/gin"
)

var (
	device string
	domainMap  map[string]int
	lock       sync.Mutex
	malwareMap map[string]int
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <interface>")
	}
	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	device = os.Args[1]
	fmt.Println(device)
	domainMap = make(map[string]int)
	malwareMap = make(map[string]int)
	malwareList := utils.LoadMalwareList("monitor_list.txt")

	go utils.CaptureDNSPcts(device, domainMap, malwareMap, &lock, malwareList)

	r.GET("/data", func(c *gin.Context) {
		lock.Lock()
		sortedDomains := utils.SortMap(domainMap)
		lock.Unlock()
		c.JSON(200, sortedDomains)
	})
	r.GET("/monitor", func(c *gin.Context) {
		lock.Lock()
		out := utils.SortMap(malwareMap)
		defer lock.Unlock()
		c.JSON(200, out)
	})
	r.Run(":8880")
}
