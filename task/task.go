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

type Report struct {
	Result string `json:"content"`
	Id int `json:"task"`
	Token string `json:"id"`
}

func GetTask(url string, client *http.Client, token []byte) Tasks {
	tasks := Tasks{}

	data := inter.Token {
		Token: string(token),
	}

	dataInJson, _ := json.Marshal(data)

	result := inter.PostHttp(url + "/api/task/", client, dataInJson)

	json.Unmarshal([]byte(result), &tasks)

	return tasks
}

func SendReport(url string, client *http.Client, token []byte, result string, task_id int) {
	report := Report {
		Result: result,
		Id: task_id,
		Token: string(token),
	}

	reportInJson, _ := json.Marshal(report)

	inter.PutHttp(url + "/api/report/", client, reportInJson)
}