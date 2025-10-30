package render

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/shamhi/geomgen"
)

// toMarkdown builds a Markdown document preserving TeX math
func toMarkdown(result geomgen.WorkResult) string {
	var b strings.Builder

	title := result.Config.Title
	if title == "" {
		title = "Generated Work"
	}
	fmt.Fprintf(&b, "# %s\n\n", title)

	b.WriteString("## Условия\n\n")
	for i, p := range result.Problems {
		fmt.Fprintf(&b, "%d. %s\n\n", i+1, p.Statement)
	}

	if result.Config.Type == geomgen.WorkTypeExam {
		b.WriteString("## Ответы\n\n")
		for i, p := range result.Problems {
			fmt.Fprintf(&b, "%d. %s\n\n", i+1, p.Solution)
		}
	}

	return b.String()
}

// DOCXRenderer renders WorkResult to a docx file content
func DOCXRenderer(result geomgen.WorkResult, extraArgs ...string) ([]byte, error) {
	md := toMarkdown(result)

	// Create secure unique temporary files
	tmpDir := os.TempDir()
	tmpMDFile, err := os.CreateTemp(tmpDir, "geomgen_*.md")
	if err != nil {
		return nil, fmt.Errorf("create temp markdown file: %w", err)
	}
	tmpMD := tmpMDFile.Name()
	defer func() {
		_ = tmpMDFile.Close()
		_ = os.Remove(tmpMD)
	}()

	// write markdown to temp file
	if _, err := tmpMDFile.WriteString(md); err != nil {
		return nil, fmt.Errorf("write temp markdown: %w", err)
	}

	// prepare temp output docx path
	tmpDOCXFile, err := os.CreateTemp(tmpDir, "geomgen_*.docx")
	if err != nil {
		return nil, fmt.Errorf("create temp docx file: %w", err)
	}
	tmpDOCX := tmpDOCXFile.Name()
	defer func() {
		_ = tmpDOCXFile.Close()
		_ = os.Remove(tmpDOCX)
	}()

	// build pandoc args
	args := []string{
		"--from=markdown+tex_math_dollars+tex_math_single_backslash",
		"--to=docx",
		"--standalone",
		"--output", tmpDOCX,
		tmpMD,
	}
	if len(extraArgs) > 0 {
		args = append(args, extraArgs...)
	}

	// execute pandoc with timeout via context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "pandoc", args...)
	stderr, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("pandoc failed: %v\n%s", err, string(stderr))
	}

	data, err := os.ReadFile(tmpDOCX)
	if err != nil {
		return nil, fmt.Errorf("read generated docx: %w", err)
	}
	return data, nil
}
