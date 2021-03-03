package auth

import (
	inter "../interaction"
    // "io/ioutil"
    "os/user"
	"net/http"
    "encoding/json"
    // "fmt"
    "os"
)

func Register(url string, client *http.Client) string {
	user, err := user.Current()

    if err != nil {
        panic(err)
    }

    name, _ := os.Hostname()
    username := user.Username

    data := inter.Data {
        Computer: name,
        Username: username,
    }

    dataInJson, _ := json.Marshal(data)

    result := inter.PutHttp(url+"/api/client/", client, dataInJson)

    token := inter.Token{}

    json.Unmarshal([]byte(result), &token)

    return token.Token
}