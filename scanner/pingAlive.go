package scanner

import (
	"fmt"
	"github.com/tatsushid/go-fastping"
	"net"
	"sync"
	"time"
)

func ICMPPing(host string) bool {
	var alive bool = false
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", host)
	if err != nil {
		fmt.Println(err)
		return false
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		alive = true
		//fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		//fmt.Println("finish")
	}
	err = p.Run()

	return alive
}

func PingAlive (host string, chAlive chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	if ICMPPing(host) {
		//fmt.Println("send alive host" + host)
		chAlive <- host
	} else{
		chAlive <- ""
	}
}