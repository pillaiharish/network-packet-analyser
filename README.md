# NETWORL PACKET ANALYSER
I have a mac connected to the internet. Want to keep a track of the websites that i visit. A Golang code running on the background which will help me to keep count of the number of times i have visited a website is capture.

## Pass the interface name 
For me it is en0 on which I am connected on my home wifi. Since this code captures packets on interface, the main should run on root permissions. The terminal logs will print the domain names.

```bash
sudo go run main.go en0
```

## Building the project
```bash
harish $ go build  
harish $ go build -o network-packet-analyser main.go 
```