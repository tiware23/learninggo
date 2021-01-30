// Dup2 exibe a contagem e o texto das linhas que aparecem
// mais de uma vez na entrada. Ele le de stdin ou de uma lista
// de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for k, v := range counts {
		for line, n := range v {
			if n > 1 {
				fmt.Printf("%d\t%s\t%s\n", n, line, k)
			}
		}
	}

}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	counts[f.Name()] = map[string]int{}
	for input.Scan() {
		_, ok := counts[f.Name()]
		if ok {
			counts[f.Name()][input.Text()]++
		}
	}
}
