package model

type Message struct {
	MID      int64  `json:"mid"`
	SendName string `json:"UID"`
	RecUName string `json:"recUID"`
	Detail   string `json:"detail"`
}