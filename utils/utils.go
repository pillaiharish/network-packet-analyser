package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
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
	// lock sync.Mutex
)

type DomainCount struct {
	Domain string
	Count  int
}

// remove prefixes like "www."
func NormalizeDomain(domain string) string {

	return strings.TrimPrefix(domain, "www.")
}

func LoadMalwareList(filename string) map[string]bool {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening malware list file:", err)
	}
	defer file.Close()

	malwareList := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		normalizedDomain := NormalizeDomain(scanner.Text())
		malwareList[normalizedDomain] = true
	}
	return malwareList
}

func SortMap(DomainMap map[string]int) []DomainCount {
	var domains []DomainCount
	for domain, count := range DomainMap {
		domains = append(domains, DomainCount{Domain: domain, Count: count})
	}

	sort.Slice(domains, func(i, j int) bool {
		return domains[i].Count > domains[j].Count
	})
	return domains
}

func CaptureDNSPcts(device string, domainMap map[string]int, malwareMap map[string]int, lock *sync.Mutex, malwareList map[string]bool) {
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
				domainName := NormalizeDomain(string(query.Name))
				// recordType := int(query.Type)
				fmt.Println(domainName)
				lock.Lock()
				// domainName = domainName + "," + string(recordType)
				domainMap[domainName]++
				if malwareList[domainName] {
					malwareMap[domainName]++
				}
				lock.Unlock()
			}
		}
	}
}
