package auth

import (
	inter "../interaction"
    // "io/ioutil"
    "os/user"
	"net/http"
    "encoding/json"
    // "fmt"
)

func Register(url string, client *http.Client) string {
	user, err := user.Current()
    if err != nil {
        panic(err)
    }

    name := user.Name
    username := user.Username
    result := inter.PutOrPostBecauseArturDoesNotWantToAddOneClassHttp(url+"/api/client/", client, name, username, "PUT")

    token := inter.Token{}

    json.Unmarshal([]byte(result), &token)

    return token.Token
}