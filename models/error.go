package models

type Error struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}
