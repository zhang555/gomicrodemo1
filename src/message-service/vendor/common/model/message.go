package model

type Message struct {
	ID      uint32 `json:"id"`      //
	UserId  uint32 `json:"userId"`  //
	Message string `json:"message"` //
	ModelTime
}
