package model

import "time"

type CustomerNotify struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customerID"`
	NotifyType string    `json:"-"`
	Event      string    `json:"-"`
	ObjID      string    `json:"-"`
	Phone      string    `json:"-"`
	Title      string    `json:"title"`
	ShortText  string    `json:"shortText"`
	Text       string    `json:"Text"`
	Img        string    `json:"-"`
	ImgURL     string    `json:"imgURL"`
	Video      string    `json:"-"`
	VideoURL   string    `json:"videoURL"`
	NotifyTime time.Time `json:"notifyTime"`
	Readed     bool      `json:"readed"`
}
