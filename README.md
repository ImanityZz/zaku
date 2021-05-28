# Zaku
本项目致力于打造轻量的跨平台口令爆破工具，得益于Go优秀的并发能力和跨平台能力，帮助红队在内网渗透中快速定位弱口令应用。

## 功能

1、内置弱口令爆破字典，提高爆破效率

2、多种IP格式解析支持

3、同时支持命令行和文件中字典的自定义功能

4、可自定义爆破应用的端口号

5、可关闭Ping主机存活探测

## Usage Of zaku.exe
 -Np  
        no ping detect host alive  
        
  -fp string
        custom crack password file or commandline data
         if specify file use prefix 'file:' like '-fp file:user.txt', if specify commandline like '-fp password1,password2'
  -fu string
        custom crack user name file or commandline data
         if specify file use prefix 'file:' like '-fu file:user.txt', if specify commandline like '-fu admin,admin1,admin2'
  -ip string
        ip range, support like commandline x.x.x.x/n or x.x.x.x-x.x.x.x or load from file must specify 'file:filename' , such as '-ip file:1.txt'
  -p string
        custom crack service port, need mapping service index 'servie:port'
         For example, modify service ssh and mysql like '-p 0:1022,1:306,....'
  -s string
        service define A:All | 0:SSH | 1:MYSQL | 2:MSSQL | 3:REDIS | 4:ORACLE | 5:SMB | 6:FTP | 7: MONGODB
         example scan service MSSQL,REDIS,ORACLE, then set '-s 2,3,4'
  -t int
        run thread number (default 100)
  -timeout int
        connect timeout (default 500)

![zaku](https://github.com/L4ml3da/zaku/blob/main/img/zaku.jpg)

## 示例

例如：爆破SSH、MYSQL、MSSQL弱口令，且指定其用户名密码文件，同时修改SSH和MSSQL的默认端口为1122和2323，IP范围为192.168.1.1-192.168.1.5,192.168.1.0/24，最后禁止使用ping探测（-Np)

```
zaku.exe -s 0,1,2 -fu file:./1.txt -fp file:./1.txt -p 0:1122,2:2323 -ip 192.168.1.1-192.168.1.5,192.168.1.0/24 -Np

```


## TODO

1、其他已知的流行应用口令爆破能力

2、常见的web弱口令探测

3、其他常用应用端口探测

4、快速生成弱口令字典



## 免责申明

本项目仅供学习交流使用，请勿用于违法犯罪行为。

本软件不得用于从事违反中国人民共和国相关法律所禁止的活动，由此导致的任何法律问题与本项目和开发人员无关。

