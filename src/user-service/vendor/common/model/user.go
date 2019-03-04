package model

type User struct {
	ID       uint32 `json:"id"`       //
	Username string `json:"username"` //
	Nickname string `json:"nickname"` //
	Password string `json:"password"` //
	ModelTime
}

type UserLoginHistory struct {
	ID       uint32 `json:"id"`       //
	UserId   uint32 `json:"userId"`   //
	Token    string `json:"token"`    //
	Ip       string `json:"ip"`       //
	Logout   int    `json:"logout"`   //
	Duration int    `json:"duration"` //
	ModelTime
}

type UserSession struct {
	ID       uint32 `json:"id"`       //
	Username string `json:"username"` //
	Nickname string `json:"nickname"` //
	ModelTime
}
