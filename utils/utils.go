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
	snapshot_len int32 = 1024
	promiscuous  bool  = false
	err          error
	handle *pcap.Handle
)

type DomainCount struct {
	Domain string
	Count  int
}

// NormalizeDomain removes common prefixes like "www." from domain names
func NormalizeDomain(domain string) string {
	return strings.TrimPrefix(domain, "www.")
}

// LoadMalwareList loads a list of malware domains from a file and stores them in a map for quick lookup
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

// SortMap sorts a map of domains and their respective counts in descending order
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

// CaptureDNSPcts captures DNS packets from a network interface, updates domain counts, and tracks malware domains based on a preloaded list
func CaptureDNSPcts(device string, domainMap map[string]int, malwareMap map[string]int, lock *sync.Mutex, malwareList map[string]bool) {
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Capture only DNS packets (port 53)
	if err := handle.SetBPFFilter("port 53"); err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println(packetSource)

	for packet := range packetSource.Packets() {
		dnsLayer := packet.Layer(layers.LayerTypeDNS)
		if dnsLayer != nil {
			dns := dnsLayer.(*layers.DNS)
			for _, query := range dns.Questions {
				domainName := NormalizeDomain(string(query.Name))
				fmt.Println(domainName)
				lock.Lock()
				domainMap[domainName]++
				if malwareList[domainName] {
					malwareMap[domainName]++
				}
				lock.Unlock()
			}
		}
	}
}
