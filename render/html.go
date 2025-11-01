package render

import (
	"fmt"
	"html"
	"strings"

	"github.com/shamhi/geomgen/v2"
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

	b.WriteString("<link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/katex@0.16.9/dist/katex.min.css\">\n")
	b.WriteString("<script defer src=\"https://cdn.jsdelivr.net/npm/katex@0.16.9/dist/katex.min.js\"></script>\n")
	b.WriteString("<script defer src=\"https://cdn.jsdelivr.net/npm/katex@0.16.9/dist/contrib/auto-render.min.js\" onload=\"renderMathInElement(document.body,{delimiters:[{left:'$$',right:'$$',display:true},{left:'$',right:'$',display:false},{left:'\\\\(',right:'\\\\)',display:false},{left:'\\\\[',right:'\\\\]',display:true}],throwOnError:false});\"></script>\n")
	b.WriteString("<style>")
	b.WriteString("body{font-family:system-ui,-apple-system,'Segoe UI',Roboto,Arial,sans-serif;font-size:18px;line-height:1.6;margin:24px;}")
	b.WriteString("h1{margin:0 0 0.75em} h2{margin:1.25em 0 0.5em}")
	b.WriteString("ol{padding-left:28px;} li{margin:0.5em 0;}")
	b.WriteString("p{margin:0.25em 0;}")
	b.WriteString(".katex{font-size:1.08em;}")
	b.WriteString(".katex-display{margin:0.75em 0;}")
	b.WriteString("</style>\n")
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
