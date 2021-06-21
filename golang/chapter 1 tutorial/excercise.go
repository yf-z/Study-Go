// This file includes all exercises problems in chapter 1
// written by Yifan Zhang
// 06/09/2021

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"bufio"
	"strconv"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"net/http"
	"io"
	//"io/ioutil"
)

// later
// func solve_problems(var prob_num)

func prob_1() {
	fmt.Println(strings.Join(os.Args, " "))
}

func prob_2() {
	for idx, arg := range os.Args {
		fmt.Println("idx: ", idx, "arg: ", arg)
	}
}

func prob_3_ineff() {
	var s, sep string  // implicitly initialized as empty strings ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func prob_3_eff() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func prob_3() {
	start := time.Now().UnixNano()
	prob_3_ineff()
	end := time.Now().UnixNano()
	duration_ineff := float64((end-start))

	start = time.Now().UnixNano()
	prob_3_eff()
	end = time.Now().UnixNano()
	duration_eff := float64((end-start))

	fmt.Println("ineff: ", duration_ineff, "eff: ", duration_eff)
}

func prob_4() {
	cur_line := 0
	counts := make(map[string][]string)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { // does not return false when the input is empty
		cur_line++
		if (input.Text() == "") {
			break
		}
		counts[input.Text()] = append(counts[input.Text()], strconv.Itoa(cur_line))
	}

	// Note: ignoring potential errors from input.Err()
	for line, n := range counts {
		if len(n) > 1 {
			fmt.Println(strings.Join(n, " "), "\n", line)
		}
	}
}

// func lissajous(f* os.File)
func prob_5(out *os.File) {
	var palette = []color.Color{color.Black, color.RGBA{123, 239, 178, 1}}

	const (
		blackIndex = 0 // first color in palette
		greenIndex = 1 // next color in pallete
	)

	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)

	freq := rand.Float64()*4.0 // relative frequency if y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: igoring encoding errors
	out.Close()
}

func prob_6(out *os.File, color_nums int) {
	seed := rand.NewSource(time.Now().UnixNano())
	ran := rand.New(seed)
	var palette = []color.Color{}
	for i := 0; i < color_nums; i++ {
		// r, g, b, y := rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256)
		// fmt.Println(r, g, b, y)
		palette = append(palette, color.RGBA{uint8(ran.Intn(256)), uint8(ran.Intn(256)), uint8(ran.Intn(256)), uint8(ran.Intn(256))})
	}

	blackIndex := uint8(ran.Intn(len(palette))) // first color in palette
	greenIndex := uint8(ran.Intn(len(palette))) // next color in pallete
	for greenIndex == blackIndex {
		greenIndex = uint8(ran.Intn(len(palette)))
	}

	const (
		cycles = 5 // number of complete x oscillator revolutions
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [-size..+size]
		nframes = 64 // number of animation frames
		delay = 8 // delay between frames in 10ms units
	)

	freq := rand.Float64()*4.0 // relative frequency if y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size + int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: igoring encoding errors
	out.Close()
}

func prob_7(links [2]string) {
	for _, url := range links {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func prob_8(links [2]string) {
	for _, url := range links {
		if !(strings.HasPrefix(url, "http://")) {
			url = "http://"+url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func prob_9(links [2]string) {
	for _, url := range links {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("Status: %v\n", resp.Status)
	}
}

func main() {
	// prob_1()
	// prob_2()
	// prob_3()
	// prob_4()
	// f, _ := os.Create("prob_5.gif")
	// prob_5(f)
	// f, _ := os.Create("prob_6.gif")
	// prob_6(f, 20)
	links := [2]string{"http://gopl.io", "http://bad.gopl.io"}
	// prob_7(links)
	// links_8 := [2]string{"gopl.io", "bad.gopl.io"}
	// prob_8(links_8)
	prob_9(links)
}