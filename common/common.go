package common

const (
	WclSolutionsApiUrl = "http://localhost:8881"
)

type Msg struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
