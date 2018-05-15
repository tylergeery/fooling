package main

import (
    "fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(p []byte) (bytes int, err error) {
    bytes, err = r.r.Read(p)

    if err == nil {
        for i := 0; i < bytes; i++ {
            if p[i] <= 'Z' && p[i] >= 'A' {
                p[i] += 13

                if p[i] > 'Z' {
                    p[i] -= 26
                }
            } else if (p[i] <= 'z' && p[i] >= 'a') {
                p[i] += 13

                if p[i] > 'z' {
                    p[i] -= 26
                }
            }
        }
    }

    return
}

func main() {
    if (len(os.Args) != 2) {
        help()
        return
    }

	s := strings.NewReader(os.Args[1])
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

    return
}

func help() {
    fmt.Printf("Usage: %s <integer>\n", os.Args[0])
}
