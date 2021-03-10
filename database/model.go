package database

type Todo struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"msg"`
}
