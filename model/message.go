package model

import "time"

type Message struct {
	MID      int64     `json:"mid"`
	UID      int64     `json:"uid"`
	SendName string    `json:"UID"`
	RecUName string    `json:"recUID"`
	Detail   string    `json:"detail"`
	Time     time.Time `json:"time"`
	IsDelete bool      `json:"is_delete"`
}
