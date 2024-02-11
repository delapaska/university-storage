package todo

type ProjectList struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title" binding:"required"`
	Directory string `json:"directory" db:"directory"`
}
