package fcolorprinter

import (
	"fmt"
	"testing"
)

func TestFmpr(t *testing.T) {
	tp := []ToPrint{
		{"%s\n", "red %s", "hello"},
		{"%+v\n", InfoColor, []int{1, 2, 3}},
	}
	fmpr(tp)
}

var ansicols []int = []int{
	1, 2, 3, 4, 5, 7,
	30, 31, 32, 33, 34, 35, 36,
	41, 42, 43, 44, 45, 46, 47,
	90, 91, 92, 93, 94, 95, 96, 97,
	100, 101, 102, 103, 104, 105, 106, 107,
}

var ansirows []int = []int{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	91, 92, 93, 94, 95, 96, 97, 98, 99,
	101, 102, 103, 104, 105, 106, 107, 108, 109,
}

func TestAnsiColor256(t *testing.T) {
	fmt.Printf("\n\n%s", "TestAnsiColor256:")
	for _, i := range ansirows {
		fmt.Printf("\n%03d ", i)
		for _, j := range ansicols {
			fmt.Printf("\033[%d;%dm %03d\033[0m", j, i, j)
		}
	}
	fmt.Printf("\n")
}

/*
func TestAnsiColorAll(t *testing.T) {
	var j int
	for i := 0; i < 257; i++ {
		fmt.Printf("\033[%d;%dm%d\033[0m\n", j, i, i)
		for j := 0; j < 257; j++ {
			fmt.Printf("\033[%d;%dm %d\033[0m", j, i, i)
		}
	}
}
*/

// reset \u001b[0m
func TestUnicodeColor256(t *testing.T) {
	fmt.Printf("\n\n%s", "TestUnicodeColor256:")
	for i := 0; i < 16; i++ {
		fmt.Printf("\n")
		for j := 0; j < 16; j++ {
			code := i*16 + j
			fmt.Printf("\u001b[38;5;%dm %03d", code, code)
		}
	}
	fmt.Printf("\n")
}

func TestColor24bit(t *testing.T) {
	res := 5 // color space resolution
	fmt.Printf("\n\n%s\n", "TestColor24bitHSV:")
	// fmt.Printf("\x1b[38;2;250;100;0m%s\x1b[0m", "hello")
	fmt.Printf("%s\n", "x,0,0")
	for i := 0; i < 256; i += res {
		fmt.Printf("\x1b[38;2;%d;0;0m%03d \x1b[0m", i, i)
	}

	fmt.Printf("%s\n", "0,x,0")
	for i := 0; i < 256; i += res {
		fmt.Printf("\x1b[38;2;0;%d;0m%03d \x1b[0m", i, i)
	}
	fmt.Printf("%s\n", "0,0,x")
	for i := 0; i < 256; i += res {
		fmt.Printf("\x1b[38;2;0;0;%dm%03d \x1b[0m", i, i)
	}
	for i := 0; i < 256; i += res {
		RainbowColor(i)
	}

	fmt.Printf("\n")
}

func RainbowColor(i int) {
	h := i / 43
	f := i - 43*h
	t := f * 255 / 43
	q := 255 - t
	if h == 0 {
		// fmt.Printf("%s\n", "h=0")
		fmt.Printf("\x1b[38;2;255;%d;0m%03d \x1b[0m", t, t)
		fmt.Printf("%s\n", "255,t,0")
	}
	if h == 1 {
		// fmt.Printf("%s\n", "h=1")
		fmt.Printf("\x1b[38;2;%d;255;0m%03d \x1b[0m", q, q)
		fmt.Printf("%s\n", "q,255,0")
	}
	if h == 2 {
		// fmt.Printf("%s\n", "h=2")
		fmt.Printf("\x1b[38;2;0;%d;0m%03d \x1b[0m", t, t)
		fmt.Printf("%s\n", "0,t,0")
	}
	if h == 3 {
		// fmt.Printf("%s\n", "h=2")
		fmt.Printf("\x1b[38;2;0;%d;255m%03d \x1b[0m", q, q)
		fmt.Printf("%s\n", "0,t,0")
	}
	if h == 4 {
		// fmt.Printf("%s\n", "h=2")
		fmt.Printf("\x1b[38;2;%d;0;255m%03d \x1b[0m", t, t)
		fmt.Printf("%s\n", "t,0,255")
	}
	if h == 5 {
		fmt.Printf("\x1b[38;2;255;0;%dm%03d \x1b[0m", q, q)
		fmt.Printf("%s\n", "255,0,q")
	}
}

func TestRainbowColor(t *testing.T) {
	res := 3
	for i := 0; i < 256; i += res {
		RainbowColor(i)
	}
}

func TestFma(t *testing.T) {
	DeclareColors16()
	// err := fma("% %s", "\033[1;30m%s\033[0m")
	// err := fma("% %s", "\x1b[38;2;255;0;0m%s\x1b[0m", "10")
	// err := fma("% %s", D, "10")
	err := fma(D)
	fmt.Println(err)
}
