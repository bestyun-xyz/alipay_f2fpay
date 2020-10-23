package system

import (
	"bufio"
	"github.com/smartwalle/alipay/v3"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)



var ListenPort string


//支付宝公钥地址
var alipay_public_key = ""

//商户私钥地址
var private_key = ""

//应用id
var appid  = ""



var Client *alipay.Client


var StartTime time.Time

//读取key=value类型的配置文件
func InitConfig() map[string]string {
	config := make(map[string]string)
	ostype := runtime.GOOS

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return config
	}

	if ostype == "windows"{
		//strings.Replace(dir, "\\", "/", -1)
		dir = dir + "\\conf\\server.conf"
	}else {
		//strings.Replace(dir, "\\", "/", -1)
		dir = dir + "/conf/server.conf"
	}

	f, err := os.Open(dir)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}


func InitPara() {
	config := InitConfig()

	ListenPort = ":"+config["listen_port"]




	alipay_public_key = config["alipay_public_key"]
	private_key = config["private_key"]
	appid = config["appid"]



	Client, _ = alipay.New(appid, private_key, true)
	Client.LoadAliPayPublicKey(alipay_public_key)


	timely,_  := time.LoadLocation("Asia/Shanghai")
	time.Local = timely

	StartTime = time.Now().Local()
}