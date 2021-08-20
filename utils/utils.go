package utils

import (
	// "fmt"
	"io"
	"strings"
  "time"
)

func Type(w io.Writer, content string ) {
	chars := strings.Split(content, "")

	for _, c := range chars {
    time.Sleep(50 * time.Millisecond)

    w.Write([]byte(c))
	}
}

func AddText(w io.Writer, content string ) {
	w.Write([]byte(content))
}

func ClearTerm(w io.Writer) {
	w.Write([]byte("\033c"))
}
