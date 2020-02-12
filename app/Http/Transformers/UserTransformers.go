package Transformers

// PostLogin POST "/login" response object
type PostLogin struct {
	Username string `json:"username"`
	ID       int64  `json:"id"`
	Token    string `json:"token"`
}
