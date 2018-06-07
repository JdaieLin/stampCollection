package controllers

type BaseRequest struct {
	Version string `json:"version"`
	UserId  int64  `json:"user_id"`
	Session string `json:"session"`
	token   string `json:"token"`
}
