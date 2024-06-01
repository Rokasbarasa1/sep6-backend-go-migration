package types

type Todo struct {
	Name string `json:"title" binding:"required,min=2,max=100"`
}
