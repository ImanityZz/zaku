package scanner

import (
	"fmt"
	"github.com/L4ml3da/zaku/conf"
	"github.com/L4ml3da/zaku/crackdict"
	"net"
	"sync"
	"time"
)


var portServiceMapping map[int] *portService

type portService struct {
	name string
	port string
	crackUserDict[]string
	crackPasswdDict[]string
	callbackFunc connectDaemon
}

func SetPort(portConfig map[int]string) {
	for k,v := range portConfig {
		fmt.Println("[*] Config service ["  +portServiceMapping[k].name + "] port -> " + v )
		portServiceMapping[k].port = v
	}
}
func GetPortServiceName (service int) string {
	return portServiceMapping[service].name
}

func SetDict(dictType int, dict[] string) {
	for k,_ := range portServiceMapping {
		switch dictType {
		case crackdict.USER_DICT:
			if len(dict) != 0 {
				fmt.Print("[*] Config user dict:")
				fmt.Println(dict)
				portServiceMapping[k].crackUserDict = dict
				return
			}
		case crackdict.PASS_DICT:
			if len(dict) != 0 {
				fmt.Print("[*] Config password dict: ")
				fmt.Println(dict)
				portServiceMapping[k].crackPasswdDict = dict
				return
			}
		}
	}
}

func PortServiceInit() {
	portServiceMapping = make(map[int] *portService)
	portServiceMapping[conf.SSH] = &portService{
											name:"SSH",
											port: "22",
											callbackFunc: mysqlConnect,
											crackUserDict: crackdict.GetWeakDict(conf.SSH, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.SSH, crackdict.PASS_DICT)}
	portServiceMapping[conf.MYSQL] = &portService{
											name:"MYSQL",
											port: "3306",
											callbackFunc: sshConnect,
											crackUserDict: crackdict.GetWeakDict(conf.MYSQL, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.MYSQL, crackdict.PASS_DICT)}
	portServiceMapping[conf.MSSQL] = &portService{
											name:"MSSQL",
											port: "1433",
											callbackFunc: mssqlConnect,
											crackUserDict: crackdict.GetWeakDict(conf.MSSQL, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.MSSQL, crackdict.PASS_DICT)}
	portServiceMapping[conf.REDIS] = &portService{
											name:"REDIS",
											port: "6379",
											callbackFunc: redisConnect,
											crackUserDict: crackdict.GetWeakDict(conf.REDIS, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.REDIS, crackdict.PASS_DICT)}
	portServiceMapping[conf.ORACLE] = &portService{
											name:"ORACLE",
											port: "1521",
											callbackFunc: oracleConnect,
											crackUserDict: crackdict.GetWeakDict(conf.ORACLE, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.ORACLE, crackdict.PASS_DICT)}
	portServiceMapping[conf.SMB] = &portService{
											name:"SMB",
											port: "445",
											callbackFunc: smbConnect,
											crackUserDict: crackdict.GetWeakDict(conf.SMB, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.SMB, crackdict.PASS_DICT)}
	portServiceMapping[conf.FTP] = &portService{
		                                    name:"FTP",
		                                    port: "21",
											callbackFunc: ftpConnect,
											crackUserDict: crackdict.GetWeakDict(conf.FTP, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.FTP, crackdict.PASS_DICT)}
	portServiceMapping[conf.MONGODB] = &portService{
											name:"MONGODB",
											port: "27017",
											callbackFunc: mongoDBConnect,
											crackUserDict: crackdict.GetWeakDict(conf.MONGODB, crackdict.USER_DICT),
											crackPasswdDict: crackdict.GetWeakDict(conf.MONGODB, crackdict.PASS_DICT)}
}

func PortScan(routineID int, chIP chan string, chCrack chan CrackTarget, chExit chan int, servicelist[]int, timeout int, wg *sync.WaitGroup) {
	defer wg.Done()
	for ipstr:= range chIP {
		for _,service := range servicelist {
			port := portServiceMapping[service].port
			//fmt.Println("scan ip addres " + ipstr + " port " + port)
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", ipstr, port), time.Millisecond*time.Duration(timeout))
			if err != nil {
				//fmt.Println("ip addres " + ipstr + " port " + port + " closed")
				continue
				//fmt.Println(err.Error())
			}
			_ = conn.Close()
			tmpCrack := CrackTarget{Ip: ipstr, Port: port, Service: service}
			chCrack <- tmpCrack
			//fmt.Println("[+] Found IP:" + ipstr + " port:" + port +" service:" + portServiceMapping[service].name + " open")
		}
	}
	tmpCrack := CrackTarget{Ip: "", Port:"" , Service: conf.DEFEND}
	chCrack <- tmpCrack
}
