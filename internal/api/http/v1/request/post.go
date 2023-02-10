package request

type UpdatePost struct {
	ID   int    `json:"id" binding:"required"`
	Text string `json:"text" binding:"required"`
}
type RetrievePost struct {
	ID int `uri:"id" binding:"required"`
}
type Post struct {
	Text string `json:"text" binding:"required"`
}
