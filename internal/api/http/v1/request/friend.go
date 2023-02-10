package request

type AddFriendRequest struct {
	FriendID int `uri:"user_id" binding:"required"`
}

type DeleteFriendRequest struct {
	FriendID int `uri:"user_id" binding:"required"`
}
