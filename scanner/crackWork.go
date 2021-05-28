package scanner

import (
	"fmt"
	"sync"
)

type CrackTarget struct {
	Ip      string
	Port    string
	Service int
}

func CrackWeak(service int,ip string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	var cracked bool = false
	fmt.Println("[!] Start crack"+" service: " + GetPortServiceName(service) + " ip: " + ip + " port: " + port)
	for _,u:= range portServiceMapping[service].crackUserDict {
		for _, p := range portServiceMapping[service].crackPasswdDict {
			if portServiceMapping[service].callbackFunc(ip, port, u, p) {
				fmt.Println("[+] Success " + GetPortServiceName(service) + " ip: " + ip + " port: " + port + " user: " + u + " password: " + p)
				cracked = true
				break
			} else {
				//fmt.Println("target ip: " + ip + " port: " + port + " user " + u + "password " + p + " fail")
			}
		}
		if cracked{
			break
		}
	}
}
