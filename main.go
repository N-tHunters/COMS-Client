package main

import (
	inter "../Coms/interaction"
	auth "../Coms/auth"
	"net/http"
	"fmt"
	"net/url"
	"runtime"
	"os"
    "io/ioutil"
    "strings"
)

func main() {
	myUrl := "http://127.0.0.1:8000"
	proxy := false
	OS := runtime.GOOS
	client := &http.Client{}
	home := os.Getenv("HOME") + "/"

	err := os.Mkdir(home + ".coms", 8567)

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
    if err != nil {
		token = []byte(auth.Register(myUrl, client))
		inter.PutHttpToken(myUrl + "/api/client/connect", client, token)
		f, _ := os.Create(home + ".coms/token")
		defer f.Close()
		f.Write(token)
	} else {
		fmt.Println(inter.PutHttpToken(myUrl + "/api/connect", client, token))
	}
	fmt.Println(string(token))
}