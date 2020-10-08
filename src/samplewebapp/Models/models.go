package models

//Subscription struct
type Subscription struct {
	Product string `json:"product"`
	Type    string `json:"type"`
}

var subs []Subscription
