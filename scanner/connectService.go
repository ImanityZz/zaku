package scanner

import (
	"database/sql"
	"fmt"
	"context"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jlaffaye/ftp"
	"github.com/stacktitan/smb/smb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
	"time"
)

type connectDaemon func(string, string, string, string) bool

func sshConnect(ip string, port string, user string, passwd string) bool {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		ClientVersion:   "",
		Timeout:         5000 * time.Millisecond,
	}
	_, err := ssh.Dial("tcp", ip+ ":" +port, sshConfig)
	if err != nil {
		//fmt.Println("ssh connect fail")
		return false
	}
	//fmt.Println("ssh connect success")
	return true
}
func mysqlConnect(ip string, port string, user string, passwd string) bool {
	var DB *sql.DB
	path := strings.Join([]string{user, ":", passwd, "@tcp(",ip, ":", port, ")/", "mysql", "?charset=utf8"}, "")
	DB, _ = sql.Open("mysql", path)
	if err := DB.Ping(); err != nil {
		//fmt.Println("open database fail")
		return false
	}
	return true
}

func mssqlConnect(ip string, port string, user string, passwd string) bool {
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", ip, port, "master", user, passwd)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return false
	}
	//fmt.Println("connect success")
	conn.Close()
	return true
}

func redisConnect(ip string, port string, user string, passwd string) bool {
	client := redis.NewClient(&redis.Options{
		Addr: ip+":"+port,
		Password: passwd,
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		if strings.Contains(err.Error(), "NOAUTH"){
			return true
		}
		return false
	}
	//fmt.Println("redis connect success")
	return true
}

func ftpConnect(ip string, port string, user string, passwd string) bool {
	//fmt.Println("try user " + user + " passwd " + passwd)
	c, err := ftp.Dial(ip+":"+port, ftp.DialWithTimeout(5000*time.Millisecond))
	if err != nil {
		return false
	}

	err = c.Login(user, passwd)
	if err != nil {
		//fmt.Println("ftp login fail")
		return false
	}
	//fmt.Println("ftp login success")
	return true
}


func smbConnect(ip string, port string, user string, passwd string) bool {
	port_i,_ := strconv.Atoi(port)
	options := smb.Options{
		Host:        ip,
		Port:        port_i,
		User:        user,
		Domain:      "",
		Workstation: "",
		Password:    passwd,
	}
	debug := false
	session, err := smb.NewSession(options, debug)
	if err != nil {
		//fmt.Println("[+] Login fail")
		return false
	}
	if session.IsAuthenticated {
		session.Close()
		//fmt.Println("[+] Login successful")
		return true
	}
	//fmt.Println("[+] Login fail")
	return false
}

func oracleConnect(ip string, port string, user string, passwd string) bool{
	db, err := sql.Open("godror", "user="+user+" password="+passwd+" connectString="+ip+":"+port+"/helowin")
	if err != nil {
		return false
	}
	err = db.Ping()
	if err != nil {
		return false
	}
	db.Close()
	return true
}

func mongoDBConnect(ip string, port string, user string, passwd string) bool {
	var mgoCli *mongo.Client
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://" + ip +":" + port)

	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		//fmt.Println("Mongodb connect error")
		return false
	}
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		//fmt.Println("Mongodb connect fail")
		return false
	}

	//fmt.Println("ip "+ ip +" Mongodb connect success")
	return true
}
