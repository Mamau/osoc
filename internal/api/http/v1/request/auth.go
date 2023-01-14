package request

type Refresh struct {
	Token string `json:"token" binding:"required"`
}

type Authorization struct {
	FirstName string `json:"first_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
type Registration struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Sex       string `json:"sex" binding:"required"`
	Interests string `json:"interests" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Age       int    `json:"age" binding:"required"`
}
