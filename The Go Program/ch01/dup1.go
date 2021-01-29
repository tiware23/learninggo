// Dup1 exibe o texto de toda a liinha que aparece mais
// de uma vez no stdin, precedida por sua contagem

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// Nota: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
