package error

import (
	"fmt"
	"os"
)

func Assert(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	}
}
