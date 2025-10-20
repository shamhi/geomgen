package render

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/shamhi/geomgen"
)

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

	tmpDir := os.TempDir()
	tmpMD := filepath.Join(tmpDir, "geomgen_"+strings.ReplaceAll(result.Config.Title, " ", "_")+".md")
	tmpDOCX := filepath.Join(tmpDir, "geomgen_output.docx")

	if err := os.WriteFile(tmpMD, []byte(md), 0o644); err != nil {
		return nil, fmt.Errorf("write temp markdown: %w", err)
	}
	defer os.Remove(tmpMD)
	defer os.Remove(tmpDOCX)

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

	cmd := exec.Command("pandoc", args...)
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
