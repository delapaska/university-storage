package todo

type ProjectList struct {
	Id        int    `json:"id"`
	Title     string `json:"title" binding:"required"`
	Directory string `json:"directory"`
}
