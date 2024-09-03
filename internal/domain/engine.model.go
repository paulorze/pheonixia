package domain

type AskRequest struct {
	TableName string `json:"tablename"`
	Query     string `json:"query"`
}
