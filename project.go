package todo

type ProjectList struct {
	Id      int      `json:"id" db:"id"`
	Title   string   `json:"title" db:"title" binding:"required"`
	Folders []string `json:"folder_name" db:"folder_name" binding:"required"`
}
