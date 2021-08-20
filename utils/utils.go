package utils

import (
	"fmt"
	"io"
	"strings"
)

func Type(w io.Writer, content string) {
	chars := strings.Split(content, "")

	for _, c := range chars {
		// fmt.Println(c)
		fmt.Fprint(w, c)
	}
}

func ClearTerm(w io.Writer) {
	w.Write([]byte("\033c"))
}
