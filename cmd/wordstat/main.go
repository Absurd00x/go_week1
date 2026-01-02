package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := 0
	if _, err := fmt.Fscan(in, &n); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read n:", err)
		os.Exit(1)
	}

	for i := 0; i < n; i++ {
		s := ""
		if _, err := fmt.Fscan(in, &s); err != nil {
			fmt.Fprintln(os.Stderr, "failed to read word:", err)
			os.Exit(1)
		}
		fmt.Fprintln(out, s)
	}
}
