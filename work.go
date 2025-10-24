package geomgen

// Core types for orchestrating works (KR, Exam, Typical) and rendering

type WorkType string

const (
	WorkTypeKR      WorkType = "kr"
	WorkTypeExam    WorkType = "exam"
	WorkTypeTypical WorkType = "typical"
)

type OutputFormat string

const (
	OutputHTML OutputFormat = "html"
	OutputDOCX OutputFormat = "docx"
)

// Options control generation ranges, difficulty and presentation preferences
type Options struct {
	Difficulty  int  `json:"difficulty"` // 1-easy, 2-medium, 3-hard
	NiceAnswers bool `json:"nice_answers"`
	Use2D       bool `json:"use_2d"`
	MaxAttempts int  `json:"max_attempts"`
}

// WorkItem selects a generator by key and specifies repetitions and per-item options
type WorkItem struct {
	Key     string  `json:"key"`
	Count   int     `json:"count"`
	Options Options `json:"options"`
}

// WorkConfig is the input describing the work to assemble
type WorkConfig struct {
	Type    WorkType     `json:"type"`
	Format  OutputFormat `json:"format"`
	Seed    string       `json:"seed"`
	Title   string       `json:"title"`
	Items   []WorkItem   `json:"items"`
	Options Options      `json:"options"`
}

// Problem holds a single statement and its solution
type Problem struct {
	Category  string `json:"category"`
	Title     string `json:"title"`
	Statement string `json:"statement"`
	Solution  string `json:"solution"`
}

// WorkResult carries generated problems along with original config
type WorkResult struct {
	Config   WorkConfig `json:"config"`
	Problems []Problem  `json:"problems"`
}
