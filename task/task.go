package tasking

import (
	inter "../interaction"
    "encoding/json"
    "net/http"
    // "fmt"
)

type Task struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Token string `json:"client"`
	ModuleName string `json:"module"`
	Args string `json:"arguments"`
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func GetTask(url string, client *http.Client, token []byte) Tasks {
	tasks := Tasks{}

	result := inter.PostHttpToken(url + "/api/task/", client, token)

	json.Unmarshal([]byte(result), &tasks)

	return tasks
}