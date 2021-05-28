package crackdict

import (
	"github.com/L4ml3da/zaku/conf"
)

var normalWeakPass = []string {
	"root","admin","password","administrator","123","1234","12345","123456","1234567",
	"12345678","123456789","1234567890","rootroot","!QAZ2wsx","!QAZxsw2","1qaz!QAZ","1qaz@WSX",
	"a123456","123123","888888","88888888","123456qwerty","Passw0rd","admin123","admin1234",
	"admin999","admin888","admin@888","admin123!@#","admin123","admin@123","admin#123","password",
	"Password","P@ssw0rd","root","root123","root@123","root1234","root@1234","Root@123",
	"Root@1234","root888","administrator","000000","00000000","111111","11111111","qwerty",
	"test","test123","test1234","1q2w3e4r","1qaz2wsx","qazwsx","123qwe","123qaz",
	"0000","password123","1q2w3e","abc123","pass","pass123","pass1234","pwd",
	"pwd123","pwd1234","pwd888",
}

var oracleWeakPass = []string {
	"oracle","system","test","12345", "123456","12345678","123456789",
	"P@ssw0rd","admin123","1qaz@WSX",
}

var mssqlWeakPass = []string {
	"sa","root","admin","password","administrator","123","1234","12345","123456","1234567",
	"12345678","123456789","1234567890","rootroot","!QAZ2wsx","!QAZxsw2","1qaz!QAZ","1qaz@WSX",
	"a123456","123123","888888","88888888","123456qwerty","Passw0rd","admin123","admin1234",
	"admin999","admin888","admin@888","admin123!@#","admin123","admin@123","admin#123","password",
	"Password","P@ssw0rd","root123","root@123","root1234","root@1234","Root@123",
	"Root@1234","root888","administrator","000000","00000000","111111","11111111","qwerty",
	"test","test123","test1234","1q2w3e4r","1qaz2wsx","qazwsx","123qwe","123qaz",
	"0000","password123","1q2w3e","abc123","pass","pass123","pass1234","pwd",
	"pwd123","pwd1234","pwd888",
}

var redisWeakPass = []string {
	"",
}
var ftpWeakPass = []string {
	"anonymous", "ftp", "test", "root","admin","password","123","1234","12345","123456","1234567",
	"12345678","123456789","1234567890","rootroot","!QAZ2wsx","!QAZxsw2","1qaz!QAZ","1qaz@WSX",
}

var mongoDBWeakPass = []string {
	"",
}

var mssqlWeakUser = []string {
	"sa", "test", "root","admin",
}
var oracleWeakUser = []string {
	"system" ,"oracle", "sys", "test","web", "root", "admin",
}

var mysqlWeakUser = []string {
	"root", "admin",
}
var sshWeakUser = []string {
	"root",
}
var redisWeakUser = []string {
	"",
}

var mongoDBWeakUser = []string {
	"",
}

var ftpWeakUser = []string {
	"anonymous", "ftp", "administrator", "test","admin",
}

var smbWeakUser = []string {
	"administrator","admin","test", "manager","webadmin","guest",
}
const (
	USER_DICT = iota
	PASS_DICT
)

func GetWeakDict(service int, dicttype int) []string{
	switch service {
	case conf.MYSQL:
		if dicttype == USER_DICT{
			return mysqlWeakUser
		} else{
			return normalWeakPass
		}
	case conf.SSH:
		if dicttype == USER_DICT{
			return sshWeakUser
		} else{
			return normalWeakPass
		}
	case conf.MSSQL:
		if dicttype == USER_DICT{
			return mssqlWeakUser
		} else{
			return mssqlWeakPass
		}
	case conf.REDIS:
		if dicttype == USER_DICT{
			return redisWeakUser
		} else{
			return redisWeakPass
		}
	case conf.SMB:
		if dicttype == USER_DICT{
			return smbWeakUser
		} else{
			return normalWeakPass
		}
	case conf.ORACLE:
		if dicttype == USER_DICT{
			return oracleWeakUser
		} else{
			return oracleWeakPass
		}
	case conf.FTP:
		if dicttype == USER_DICT{
			return ftpWeakUser
		} else{
			return ftpWeakPass
		}
	case conf.MONGODB:
		if dicttype == USER_DICT{
			return mongoDBWeakUser
		} else{
			return mongoDBWeakPass
		}
	default:
		return nil;
	}
	return nil;
}