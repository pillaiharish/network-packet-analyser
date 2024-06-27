package utils

import (
	"fmt"
	"log"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	// device       string
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	// timeout      pcap.BlockForever
	handle *pcap.Handle
	// domainMap map[string]int
	lock sync.Mutex
)

func CaptureDNSPcts(device string, domainMap map[string]int) {
	// Capturing live packets using pcap
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	//capture only dsn packets
	if err := handle.SetBPFFilter("port 53"); err != nil {
		log.Fatal(err)
	}

	// https://pkg.go.dev/github.com/google/gopacket#hdr-Reading_Packets_From_A_Source
	// handle.LinkType hanles the interface type like ethernet /wifi
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println(packetSource)

	for packet := range packetSource.Packets() {
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		if dnsLayer != nil {
			dns := dnsLayer.(*layers.DNS)
			for _, query := range dns.Questions {
				domainName := string(query.Name)
				// recordType := int(query.Type)
				fmt.Println(domainName)
				lock.Lock()
				// domainName = domainName + "," + string(recordType)
				domainMap[domainName]++
				lock.Unlock()
			}
		}
	}
}
