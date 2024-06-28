package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/miekg/dns"
)

func generateTCP(wg *sync.WaitGroup) {
	defer wg.Done()
	// The Dial function connects to a server on network type (tcp)
	conn, err := net.Dial("tcp", "golang.org:80")
	if err != nil {
		fmt.Println("Error connecting tcp:", err)
	}
	defer conn.Close()
	fmt.Println("TCP connection successfull")

}

func generateUDP(wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("udp", "8.8.8.8:53") // Google dns
	if err != nil {
		fmt.Println("Error connecting udp:", err)
	}
	defer conn.Close()
	fmt.Println("UDP connection successfull")
}

func generateHTTP(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("http://facebook.com")
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTP request done")
}

func generateHTTPS(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("HTTPS error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("HTTPS request done")
}

/*
	Below is from github.com/miekg/dns

// Msg contains the layout of a DNS message.

	type Msg struct {
		MsgHdr
		Compress bool       `json:"-"` // If true, the message will be compressed when converted to wire format.
		Question []Question // Holds the RR(s) of the question section.
		Answer   []RR       // Holds the RR(s) of the answer section.
		Ns       []RR       // Holds the RR(s) of the authority section.
		Extra    []RR       // Holds the RR(s) of the additional section.
	}
*/
func generateDNS(wg *sync.WaitGroup) {
	defer wg.Done()
	m := new(dns.Msg) // Msg contains the layout of a DNS message.
	m.SetQuestion(dns.Fqdn("example.com"), dns.TypeA)
	c := new(dns.Client)                     // A Client defines parameters for a DNS client.
	_, _, err := c.Exchange(m, "8.8.8.8:53") // Exchange sends synchronous message to address and wait for reply
	if err != nil {
		fmt.Println("DNS error:", err)
		return
	}
	fmt.Println("DNS query done")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	// go generateTCP(&wg)
	go generateUDP(&wg)
	// go generateHTTP(&wg)
	go generateDNS(&wg)
	// go generateHTTPS(&wg)
	wg.Wait()

}

//https://www.whoisds.com/
