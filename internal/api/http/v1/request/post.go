package request

type DeletePost struct {
	ID int `uri:"id" binding:"required"`
}
type UpdatePost struct {
	ID   int    `json:"id" binding:"required"`
	Text string `json:"text" binding:"required"`
}
type Feeds struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}
type RetrievePost struct {
	ID int `uri:"id" binding:"required"`
}
type Post struct {
	Text string `json:"text" binding:"required"`
}
