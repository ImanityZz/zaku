package main

import (
	"flag"
	"fmt"
	"github.com/L4ml3da/zaku/conf"
	"github.com/L4ml3da/zaku/crackdict"
	"github.com/L4ml3da/zaku/scanner"
	"github.com/L4ml3da/zaku/util"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main(){
	var ports,ipParam,serviceUse string
	var noPing bool
	var timeout,thread int
	var serviceList[]int
	var customerUsr, customerPass string
	fmt.Println("\n               ,---.       ,--.-.,-.                \n  ,--,----.  .--.'  \\     /==/- |\\  \\  .--.-. .-.-. \n /==/` - ./  \\==\\-/\\ \\    |==|_ `/_ / /==/ -|/=/  | \n `--`=/. /   /==/-|_\\ |   |==| ,   /  |==| ,||=| -| \n  /==/- /    \\==\\,   - \\  |==|-  .|   |==|- | =/  | \n /==/- /-.   /==/ -   ,|  |==| _ , \\  |==|,  \\/ - | \n/==/, `--`\\ /==/-  /\\ - \\ /==/  '\\  | |==|-   ,   / \n\\==\\-  -, | \\==\\ _.\\=\\.-' \\==\\ /\\=\\.' /==/ , _  .'  \n `--`.-.--`  `--`          `--`       `--`..---'    \n" )
	fmt.Println("                                      Authored by L4ml3da version@1.0")
	fmt.Println("=====================================================================")
	flag.StringVar(&ports, "p", "", "custom crack service port, need mapping service index 'servie:port'\n For example, modify service ssh and mysql like '-p 0:1022,1:306,....'")
	flag.BoolVar(&noPing, "Np", false, "no ping detect host alive")
	flag.StringVar(&ipParam, "ip", "","ip range, support like commandline x.x.x.x/n or x.x.x.x-x.x.x.x or load from file must specify 'file:filename' , such as '-ip file:1.txt'")
	flag.StringVar(&serviceUse, "s", "","service define A:All | 0:SSH | 1:MYSQL | 2:MSSQL | 3:REDIS | 4:ORACLE | 5:SMB | 6:FTP | 7: MONGODB \n example scan service MSSQL,REDIS,ORACLE, then set '-s 2,3,4' ")
	flag.IntVar(&timeout, "timeout", 500,"connect timeout")
	flag.IntVar(&thread, "t", 100,"run thread number")
	flag.StringVar(&customerUsr, "fu", "","custom crack user name file or commandline data \n if specify file use prefix 'file:' like '-fu file:user.txt', if specify commandline like '-fu admin,admin1,admin2'")
	flag.StringVar(&customerPass, "fp", "","custom crack password file or commandline data \n if specify file use prefix 'file:' like '-fp file:user.txt', if specify commandline like '-fp password1,password2'")
	flag.Parse()

	var counter int = 0
	var wg sync.WaitGroup
	var chIP = make(chan string, 100000)
	var chAlive = make(chan string, 10000)
	var chCrack = make(chan scanner.CrackTarget, 10000)
	var chanPortExit   = make(chan int, thread)
	var aliveHosts[]string

	if len(serviceUse) == 0 {
		fmt.Println("[-] Must specify service ")
		os.Exit(-1)
	}

	scanner.PortServiceInit()

	if len(customerUsr) != 0 || len(customerPass) != 0 {
		userList := crackdict.DictRead(customerUsr)
		if userList != nil {
			scanner.SetDict(crackdict.USER_DICT,userList)
		}
		passList := crackdict.DictRead(customerPass)
		if passList != nil {
			scanner.SetDict(crackdict.PASS_DICT, passList)
		}
	}

	if len(ports) != 0 {
		portMap := make(map[int]string)
		tmpP := strings.Split(ports, ",")
		for _,tmp := range tmpP {
			s,_ := strconv.Atoi(strings.Split(tmp, ":")[0])
			p := strings.Split(tmp, ":")[1]
			portMap[s] = p
		}
		scanner.SetPort(portMap)
	}

	tmpList := strings.Split(serviceUse, ",")
	fmt.Print("[*] Load service ")
	for i:=0; i < len(tmpList); i++ {
		if strings.Compare(tmpList[i], "A") == 0 {
			for j:=0; j < conf.DEFEND; j++ {
				serviceList = append(serviceList, j)
				fmt.Print(" [" +scanner.GetPortServiceName(j) + "]")
			}
			break
		}
		idx,_ := strconv.Atoi(tmpList[i])
		serviceList = append(serviceList, idx)
		fmt.Print(" [" + scanner.GetPortServiceName(idx) +"]")
	}
	fmt.Println("")
	//fmt.Println(serviceList)

	iplist := util.ListIPAddress(ipParam)
	//fmt.Println(iplist)
	if noPing {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, host := range iplist {
				chIP <- host
			}
			close(chIP)
		}()
		fmt.Println("[*] Disable ping host")
	} else {
		fmt.Println("[*] Enable ping host")
		//fmt.Println(iplist)

		for i:=0; i < len(iplist); i++ {
			//fmt.Println("create counter" + strconv.Itoa(counter))
			counter++
			wg.Add(1)
			go scanner.PingAlive(iplist[i], chAlive, &wg)
			//time.Sleep(50*time.Millisecond)
		}

		var cnt int = 0
		for h:= range chAlive{
			if len(h) != 0{
				aliveHosts = append(aliveHosts, h)
			}
			//fmt.Println("got counter" + strconv.Itoa(cnt))
			cnt++
			if cnt == counter {
				close(chAlive)
			}
		}

		//fmt.Println("get all alive host")
		//fmt.Println(aliveHosts)
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, host := range aliveHosts {
				chIP <- host
			}
			close(chIP)
		}()
	}

	for i:= 0; i< thread; i++ {
		go scanner.PortScan(i, chIP, chCrack, chanPortExit, serviceList, timeout, &wg)
		wg.Add(1)
	}

	for i:= 0; i< thread; i++ {
		info := <- chCrack
		if info.Service != conf.DEFEND {
			go scanner.CrackWeak(info.Service, info.Ip, info.Port, &wg)
			wg.Add(1)
		}
	}

	wg.Wait()
	close(chCrack)
	fmt.Println("[+] zaku finish")
	//fmt.Println("cpu info")
	//fmt.Println(runtime.NumCPU())
}