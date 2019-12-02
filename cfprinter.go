package fcolorprinter

import (
	"fmt"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

/*
var (
	B = "\033[1;30m%s\033[0m"
	R = "\033[1;31m%s\033[0m"
	G = "\033[1;32m%s\033[0m"
	Y = "\033[1;33m%s\033[0m"
	P = "\033[1;34m%s\033[0m"
	M = "\033[1;35m%s\033[0m"
	T = "\033[1;36m%s\033[0m"
	W = "\033[1;37m%s\033[0m"
)
*/

var R, G, B, Y, P, M, D string

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

type ToPrint struct {
	Format string
	Color  string
	Inter  interface{}
}

func fmpr(tp []ToPrint) {
	var rfmt string
	sl := make([]interface{}, len(tp))
	for j, i := range tp {
		fmt.Println("mek", i.Color)
		rfmt += fmt.Sprintf(i.Color, i.Format)
		sl[j] = i.Inter
		fmt.Println("hak", rfmt)
	}
	fmt.Printf(rfmt, sl...)
}

func DeclareColors16() {
	B = "\033[1;30m%s\033[0m"
	R = "\033[1;31m%s\033[0m"
	G = "\033[1;32m%s\033[0m"
	Y = "\033[1;33m%s\033[0m"
	P = "\033[1;34m%s\033[0m"
	M = "\033[1;35m%s\033[0m"
	D = "fuck"
}

func fma(a ...interface{}) error {
	// DeclareColors16()
	fmt.Println(a)
	/*
		var cf string // current formating
		var cc string // current color
		var comb string
		for i, j := range a {
			if i == 0 {
				// format
				switch v := j.(type) {
				case string:
					if strings.HasPrefix(v, "% ") {
						cf = v[2:len(v)]
					} else {
						return errors.New("No format given")
					}
				}
			}
			if i == 1 {
				// Color
				switch v := j.(type) {
				case string:
					if strings.HasPrefix(v, "\033[1;") {
						// if strings.HasPrefix(v, "\x1b[38;") {
						fmt.Println("Color set")
						cc = v
					} else {
						cc = "%s"
					}
				}
				comb = fmt.Sprintf(cc, cf)
			}
			if i == 2 {
				fmt.Printf(comb, j)
			}
		}
		fmt.Printf(comb, "hello")
		fmt.Printf("\nmik%q\n", cf)
	*/
	return nil
}
