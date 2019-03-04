package model

import (
	"time"
	"fmt"
)

//type Result struct {
//	Success bool        `json:"success"`
//	Code    int         `json:"code"`
//	Message string      `json:"message"`
//	Result  interface{} `json:"result"`
//	Page    interface{} `json:"page,omitempty"`
//}

type JSONTime time.Time
type JSONTime2 time.Time

const (
	Time1 = "2006-01-02"
	Time2 = "2006-01-02 15:04:05"
	Time3 = "2006年01月02日"
)

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(Time1))
	return []byte(stamp), nil
}

func (t JSONTime2) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(Time2))
	return []byte(stamp), nil
}

type ModelTime struct {
	CreateTime JSONTime `json:"createTime"`
	UpdateTime JSONTime `json:"updateTime"`
}

type ModelTime2 struct {
	CreateTime JSONTime2 `json:"createTime"`
	UpdateTime JSONTime2 `json:"updateTime"`
}
