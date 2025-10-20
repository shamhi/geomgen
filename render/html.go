package render

import (
	"fmt"
	"html"
	"strings"

	"github.com/shamhi/geomgen"
)

// HTMLRenderer renders WorkResult to a simple HTML page with LaTeX passthrough
func HTMLRenderer(result geomgen.WorkResult) string {
	var b strings.Builder
	b.WriteString("<!doctype html><html><head><meta charset=\"utf-8\">\n")
	b.WriteString("<title>")
	if result.Config.Title != "" {
		b.WriteString(html.EscapeString(result.Config.Title))
	} else {
		b.WriteString("Generated Work")
	}
	b.WriteString("</title>\n")

	// KaTeX CSS/JS + auto-render
	b.WriteString("<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/katex@0.16.9/dist/katex.min.css\">\n")
	b.WriteString("<script defer src=\"https://cdn.jsdelivr.net/npm/katex@0.16.9/dist/katex.min.js\"></script>\n")
	b.WriteString("<script defer src=\"https://cdn.jsdelivr.net/npm/katex@0.16.9/dist/contrib/auto-render.min.js\"></script>\n")
	b.WriteString("</head><body><main>")

	b.WriteString(fmt.Sprintf("<h1>%s</h1>", html.EscapeString(result.Config.Title)))
	b.WriteString("<h2>Условия</h2><ol>")
	for _, p := range result.Problems {
		b.WriteString("<li><div class=\"statement\"><p>")
		b.WriteString(p.Statement)
		b.WriteString("</p></div></li>")
	}
	b.WriteString("</ol>")

	if result.Config.Type == geomgen.WorkTypeExam {
		b.WriteString("<h2>Ответы</h2><ol>")
		for _, p := range result.Problems {
			b.WriteString("<li><div class=\"solution\"><p>")
			b.WriteString(p.Solution)
			b.WriteString("</p></div></li>")
		}
		b.WriteString("</ol>")
	}

	b.WriteString("<script>")
	b.WriteString("document.addEventListener('DOMContentLoaded', function(){")
	b.WriteString("if (typeof renderMathInElement==='function'){")
	b.WriteString("renderMathInElement(document.body,{")
	b.WriteString("delimiters:[")
	b.WriteString("{left:'$$',right:'$$',display:true},")
	b.WriteString("{left:'$',right:'$',display:false},")
	b.WriteString("{left:'\\\\(',right:'\\\\)',display:false},")
	b.WriteString("{left:'\\\\[',right:'\\\\]',display:true}")
	b.WriteString("],")
	b.WriteString("throwOnError:false")
	b.WriteString("});")
	b.WriteString("}")
	b.WriteString("});")
	b.WriteString("</script>")

	b.WriteString("</main></body></html>")
	return b.String()
}
