# NETWORK PACKET ANALYSER
I have a Mac connected to the Internet. Want to keep a track of the websites that I visit. A Golang code running on the background which will help me to keep count of the number of times I have visited a website is captured. Also I have a monitor list which will monitor the domains I want to observe and flag them separately when visited along with its frequency.

## Pass the interface name 
Find out the active interface which is serving network traffic. It can be Wi-Fi or Ethernet interface. For me it is en0 on which I am connected on my home wifi. Since this code captures packets on interface, the main.go should run on root permissions. 
```bash
sudo go run src/main.go en0
```
The terminal logs will print the domain names. The UI will load the visited domains traffic and the malware domains. The malware domains are the domains which are visited domains and also in the [monitor list](https://github.com/pillaiharish/network-packet-analyser/blob/main/monitor_list.txt)



## Building the project
```bash
go build  
go build -o network-packet-analyser main.go 
```

## Postman capture for Visited domains 
/data route has all the dns traffic visited. GET visited websites data from url http://localhost:8880/data
```json
[
    {
        "Domain": "google.com",
        "Count": 8
    },
    {
        "Domain": "app.link",
        "Count": 4
    },
    {
        "Domain": "content-autofill.googleapis.com",
        "Count": 4
    },
    {
        "Domain": "googletagmanager.com",
        "Count": 4
    },
    {
        "Domain": "49-courier.push.apple.com",
        "Count": 4
    },
]
```

## Postman capture for Monitor domains 
/monitor route has the domains that are visited and in the [monitor_list.txt](https://github.com/pillaiharish/network-packet-analyser/blob/main/monitor_list.txt). GET monitor data from url http://localhost:8880/monitor
```json
[
    {
        "Domain": "newpointglobal.com",
        "Count": 10
    },
    {
        "Domain": "semaglutidereview.us",
        "Count": 4
    }
]
```

## Postman screencapture
![Monitor and Visited Domains](https://github.com/pillaiharish/network-packet-analyser/blob/main/snapshots/monitored_visited_sites.png)
![Visited Domains JSON](https://github.com/pillaiharish/network-packet-analyser/blob/main/snapshots/get_for_visited_sites.png)
![Monitor Domains JSON](https://github.com/pillaiharish/network-packet-analyser/blob/main/snapshots/get_for_monitor_sites.png)