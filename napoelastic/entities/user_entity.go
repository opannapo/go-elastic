package entities

type UserEntity struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Gender    string   `json:"gender"`
	Interests []string `json:"interests"`
}
