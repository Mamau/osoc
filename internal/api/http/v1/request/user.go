package request

type UserRequest struct {
	UserID int `uri:"id" binding:"required"`
}
type UserSearch struct {
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
}
