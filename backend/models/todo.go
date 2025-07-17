package models

type Todo struct {
    ID         int    `json:"id"`
    Task       string `json:"task"`
    IsComplete bool   `json:"is_complete"`
    CreatedAt  string `json:"created_at"`
}