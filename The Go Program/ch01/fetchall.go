// Fetchall busca urls em paralelo e informe os tempos gastos e os tamanhos.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // inicia uma gorotine
	}
	// for range os.Args[1:] {
	// 	fmt.Println(<-ch) // recebe do canal ch
	// }

	for range os.Args[1:] {
		respOutput(<-ch) // recebe do canal ch
	}

	fmt.Printf("%2.fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

func respOutput(output string) {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error to create %s", err)
	}
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println(err)
		file.Close()
	}
	file.Close()
}
