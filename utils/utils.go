package utils

import (
	// "fmt"
	"io"
	"strings"
  "time"

  markdownP "github.com/MichaelMure/go-term-markdown"
)


func RenderMarkdownTerminal(markdown string) string{
  result := markdownP.Render(markdown, 80, 6)

  return string(result)
}

func Type(w io.Writer, content string ) {
	chars := strings.Split(content, "")

	for _, c := range chars {
    
    time.Sleep(20 * time.Millisecond)

    w.Write([]byte(c))
	}
}

func AddText(w io.Writer, content string ) {
	w.Write([]byte(content))
}

func ClearTerm(w io.Writer) {
	w.Write([]byte("\033c"))
}
