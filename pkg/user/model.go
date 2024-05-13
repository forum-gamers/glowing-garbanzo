package user

type User struct {
	Id          string `json:"id"`
	AccountType string `json:"accountType"`
	Username    string `json:"username"`
}
