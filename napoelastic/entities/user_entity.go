package entities

type UserEntity struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Gender    string   `json:"gender"`
	Interests []string `json:"interests"`
}
