package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"os"
)


func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	argv := os.Args;
	argc := len(argv);
	progname := argv[0]
	if argc != 2 {
		fmt.Fprintf(os.Stderr, "%s: error: expected one input file.\n", progname)
		os.Exit(1)
	}
	path := argv[1]

	file, _ := os.ReadFile(path)

	md := []byte(file)
	html := mdToHTML(md)

	fmt.Printf("%s\n", html)
}
