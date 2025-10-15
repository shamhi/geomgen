package geomgen

import "time"

type Task struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Statement string    `json:"statement"`
	Solution  string    `json:"solution"`
	Seed      string    `json:"seed"`
	CreatedAt time.Time `json:"created_at"`
}

type Category string

const (
	Vectors  Category = "vectors"
	Lines    Category = "lines"
	Curves   Category = "curves"
	Surfaces Category = "surfaces"
)
