# NETWORK PACKET ANALYSER
I have a mac connected to the internet. Want to keep a track of the websites that i visit. A Golang code running on the background which will help me to keep count of the number of times i have visited a website is capture.

## Pass the interface name 
For me it is en0 on which I am connected on my home wifi. Since this code captures packets on interface, the main should run on root permissions. The terminal logs will print the domain names.

```bash
sudo go run src/main.go en0
```

## Building the project
```bash
go build  
go build -o network-packet-analyser main.go 
```

## Postman capture
GET on url http://localhost:8880/data
```json
{
    "2.4.f.e.8.f.1.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.b.9.f.f.4.6.0.0.ip6.arpa": 2,
    "5.6.5.b.c.8.2.1.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.b.9.f.f.4.6.0.0.ip6.arpa": 2,
    "app.link": 6,
    "e.4.2.1.0.0.0.0.0.0.0.0.0.0.0.0.0.8.e.0.5.0.0.0.f.0.4.1.0.0.6.2.ip6.arpa": 2,
    "e4686.dsce9.akamaiedge.net": 2,
    "fbs.smoot.apple.com": 6,
    "gspe35-ssl.ls.apple.com": 6,
    "medium.com": 6,
    "pubingress-feedback-974d66e8c1341bf9.elb.ap-southeast-1.amazonaws.com": 2
}
```

## Postman screencapture
![Screeshot for API Request](https://github.com/pillaiharish/network-packet-analyser/blob/main/postman_api_response.png)