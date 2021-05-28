package util

import (
	"bytes"
	"net"
	"strconv"
	"strings"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
func removeDuplicated(elist []string) []string {
	result := make([]string, 0, len(elist))
	temp := map[string]struct{}{}
	for _, item := range elist {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func IpNetParser (cidr string)[] string{
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips

	default:
		return ips[1 : len(ips)-1]
	}
}

func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}

func ipRangeParser(ipStr string)[]string {
	var iplist[]string
	if strings.Contains(ipStr, "-") {
		ipBegin := StringIpToInt(strings.Split(ipStr, "-")[0])
		ipEnd := StringIpToInt(strings.Split(ipStr, "-")[1])
		for tmpip := ipBegin; tmpip <= ipEnd; tmpip++ {
			iplist = append(iplist, IpIntToString(tmpip))
		}
		return iplist
	}

	if strings.Contains(ipStr, "/") {
		return IpNetParser(ipStr)
	}
	iplist = append(iplist, ipStr)
	return iplist
}

func ListIPAddress(ipParam string)[] string {

	var allIP[]string

	if len(ipParam) == 0 {
		return nil
	}

	if strings.Contains(ipParam, "file:") {

		idx := strings.Index(ipParam, ":")
		//fmt.Println("get file " + param[idx + 1:])
		filename := ipParam[idx+1:]
		ipRangeList := ReadFile(filename)
		for _,v := range ipRangeList {
			allIP= append(allIP, ipRangeParser(v)...)
		}
	} else {
		tmpList := strings.Split(ipParam, ",")
		for _,v := range tmpList {
			allIP= append(allIP, ipRangeParser(v)...)
		}
	}
	return removeDuplicated(allIP)
}