package interaction

import (
	"net/http"
    "encoding/json"
    "bytes"
    "fmt"
)

func GetHttp(url string, client *http.Client) string {
	resp, _ := client.Get(url)

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

func PutOrPostBecauseArturDoesNotWantToAddOneClassHttp(url string, client *http.Client, name string, username string, method string) string {
    data := Data {
        Computer: name,
        Username: username,
    }
    json_data, err := json.Marshal(data)

    req, _ := http.NewRequest(method, url, bytes.NewBuffer(json_data))
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, err := client.Do(req)

    if err != nil {
        return "[ERROR CANNOT CONNECT]"
    }

    defer resp.Body.Close()

    result := ""

    for true {
             
        bs := make([]byte, 4096)
        n, err := resp.Body.Read(bs)
        result += string(bs[:n])
         
        if n == 0 || err != nil{
            break
        }
    }

    return result
}

func PutOrPostBecauseArturDoesNotWantToAddOneClassHttpToken(url string, client *http.Client, token []byte, method string) string {
    data := Token {
        Token: string(token),
    }

    json, _ := json.Marshal(data)
    req, _ := http.NewRequest(method, url, bytes.NewBuffer(json))
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
    resp, err := client.Do(req)

    for err != nil {
        fmt.Println(err)
        req, _ = http.NewRequest(method, url, bytes.NewBuffer(json))
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

func PutHttp(url string, client *http.Client, name string, username string) string {
    return PutOrPostBecauseArturDoesNotWantToAddOneClassHttp(url, client, name, username, "PUT")
}

func PutHttpToken(url string, client *http.Client, token []byte) string {
    return PutOrPostBecauseArturDoesNotWantToAddOneClassHttpToken(url, client, token, "PUT")
}

func PostHttp(url string, client *http.Client, name string, username string) string {
    return PutOrPostBecauseArturDoesNotWantToAddOneClassHttp(url, client, name, username, "POST")
}

func PostHttpToken(url string, client *http.Client, token []byte) string {
    return PutOrPostBecauseArturDoesNotWantToAddOneClassHttpToken(url, client, token, "POST")
}