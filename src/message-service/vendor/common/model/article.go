package model

type Article struct {
	ID      uint32 `json:"id"`      //
	UserId  uint32 `json:"userId"`  //
	Title   string `json:"title"`   //
	Content string `json:"content"` //
	ModelTime
}
