package Requests


// PostLogin POST "/login" request object
type PostLogin struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}