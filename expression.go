package geomgen

import "time"

type Expression[T any] struct {
	Category  string    `json:"category"`
	Data      T         `json:"data"`
	Statement string    `json:"statement"`
	Solution  string    `json:"solution"`
	Valid     bool      `json:"valid"`
	Seed      string    `json:"seed"`
	CreatedAt time.Time `json:"created_at"`
}
