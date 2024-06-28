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
	// snapshot_len int32 = 1024
	// promiscuous  bool  = false
	// err          error
	// timeout      pcap.BlockForever
	// handle    *pcap.Handle
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

// func captureDNSPcts(device string) {
// 	// Capturing live packets using pcap
// 	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, pcap.BlockForever)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer handle.Close()

// 	//capture only dsn packets
// 	if err := handle.SetBPFFilter("port 53"); err != nil {
// 		log.Fatal(err)
// 	}

// 	// https://pkg.go.dev/github.com/google/gopacket#hdr-Reading_Packets_From_A_Source
// 	// handle.LinkType hanles the interface type like ethernet /wifi
// 	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
// 	fmt.Println(packetSource)

// 	for packet := range packetSource.Packets() {
// 		dnsLayer := packet.Layer(layers.LayerTypeDNS)
// 		if dnsLayer != nil {
// 			dns := dnsLayer.(*layers.DNS)
// 			for _, query := range dns.Questions {
// 				domainName := string(query.Name)
// 				// recordType := int(query.Type)
// 				fmt.Println(domainName)
// 				lock.Lock()
// 				// domainName = domainName + "," + string(recordType)
// 				domainMap[domainName]++
// 				lock.Unlock()
// 			}
// 		}
// 	}
// }
