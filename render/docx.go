package render

import (
	"bytes"
	"fmt"

	"github.com/shamhi/geomgen"
)

// DOCXRenderer stub returning raw bytes placeholder. TODO: soon
func DOCXRenderer(result geomgen.WorkResult) []byte {
	var buf bytes.Buffer
	buf.WriteString("DOCX generation is not yet implemented.\n")
	fmt.Fprintf(&buf, "Title: %s\n", result.Config.Title)
	fmt.Fprintf(&buf, "Problems: %d\n", len(result.Problems))
	return buf.Bytes()
}
