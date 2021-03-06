// Lissajous gera animacoes GIF de figuras de Lissajous aleatorias.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.RGBA{0x00, 0x4c, 0x99, 0xff}}

const (
	whiteIndex = 0 // primeira cor da paleta
	blackIndex = 1 // proxima cor da paleta
)

func main() {

	http.HandleFunc("/", handler) // cada requisição chama handler
	log.Fatal(http.ListenAndServe(":8000", nil))

	rand.Seed(time.Now().UTC().UnixNano())
	// lissajous(os.Stdout)
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // resolucao angular
		size    = 100   // canvas da imagem de cobre de [-size..+size]
		nframes = 64    // numero de quadros de animacao
		delay   = 8     // tempo entre quadros em unidades
	)

	freq := rand.Float64() * 3.0 // frequencia relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferenca de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: Ignorando erros de codificaçao.
}

// handler ecoa o componente Path do URL requisitado
func handler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["cycles"]

	if !ok || len(keys) < 1 {
		fmt.Println("Key is missing")
		return
	}
	key := keys[0]
	k, err := strconv.Atoi(key)

	if err != nil {
		fmt.Println(err)
	}

	lissajous(w, k)
}
