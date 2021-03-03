package interaction

import (
	"net/http"
    // "encoding/json"
    "bytes"
    "fmt"
    "time"
)

func GetHttp(url string, client *http.Client) string {
	resp, err := client.Get(url)

    for err != nil {
        resp, err = client.Get(url)
        time.Sleep(5 * time.Second)
    }

	defer resp.Body.Close()

	result := ""

    for true {
             
        bs := make([]byte, 1014)
        n, err := resp.Body.Read(bs)
        result += string(bs[:n])
         
        if n == 0 || err != nil{
            break
        }
    }

    return result
}

type Data struct {
    Computer string `json:"computer"`
    Username string `json:"username"`
}

type Token struct {
    Token string `json:"token"`
}

type Response struct {
    Result string `json:"result"`
    Error  string `json:"error"`
}

func SendHttp(url string, client *http.Client, method string, data []byte) string {
    req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, err := client.Do(req)

    for err != nil {
        fmt.Println(err)
        req, _ = http.NewRequest(method, url, bytes.NewBuffer(data))
        req.Header.Set("Content-Type", "application/json; charset=utf-8")
        resp, err = client.Do(req)
    }

    defer resp.Body.Close()

    result := ""

    for true {
             
        bs := make([]byte, 1014)
        n, err := resp.Body.Read(bs)
        result += string(bs[:n])
         
        if n == 0 || err != nil{
            break
        }
    }

    return result
}

func PutHttp(url string, client *http.Client, data []byte) string {
    return SendHttp(url, client, "PUT", data)
}

func PostHttp(url string, client *http.Client, data []byte) string {
    return SendHttp(url, client, "POST", data)
}