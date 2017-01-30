package dockerfile

import (
	"bytes"
	"io"
)

// IndentedStream will indent each line printed to a stream
type IndentedStream struct {
	par    io.Writer
	prefix []byte
	next   bool
}

// NewIndentedStream creates a new indented stream.
func NewIndentedStream(par io.Writer) *IndentedStream {
	return &IndentedStream{
		par:    par,
		prefix: []byte("    "),
		next:   true,
	}
}

// Write per io.Writer. Count of bytes written does not include prefixes added.
func (is *IndentedStream) Write(p []byte) (n int, err error) {
	written := 0
	for len(p) != 0 {
		if is.next {
			_, err = is.par.Write(is.prefix)
			if err != nil {
				return written, err
			}
			is.next = false
		}

		i := bytes.IndexByte(p, '\n')
		if i == -1 {
			i = len(p) - 1
		} else {
			is.next = true
		}

		w, err := is.par.Write(p[:i+1])
		written += w
		if err != nil {
			return written, err
		}

		p = p[i+1:]
	}
	return written, nil
}
