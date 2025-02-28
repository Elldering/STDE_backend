package models

type VerifyCode struct {
	Code int    `json:"code"`
	Type string `json:"type"`
}
