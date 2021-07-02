package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func init() {
	flag.Parse()
}

func main() {
	err := run()
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}
}

func run() error {
	var filename string

	if args := flag.Args(); len(args) > 0 {
		filename = args[0]
	}

	var r io.Reader
	switch filename {
	case "":
		r = os.Stdin
	case "-":
		if args := flag.Args(); len(args) > 1 {
			r = strings.NewReader(args[1])
		} else {
			r = os.Stdin
		}
	default:
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		r = f
	}
	t, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	rep := regexp.MustCompile("[1-9]+[mps]|[TNSPWGR]")
	raw_string := rep.FindAllStringSubmatch(string(t), -1)
	// fmt.Println(string(t))
	var all_tiles string
	for _, str := range raw_string {
		tiles := convertToTiles(str[0])
		all_tiles += tiles
	}
	fmt.Println(all_tiles)

	return nil
}

func convertToTiles(str string) string {
	var ret string
	lastChar := str[len(str)-1]
	switch lastChar {
	case 'm':
		for _, c := range str[:len(str)-1] {
			rune := ('🀇' + c - 49)
			ret += string(rune)
		}
		return ret
	case 'p':
		for _, c := range str[:len(str)-1] {
			rune := ('🀙' + c - 49)
			ret += string(rune)
		}
		return ret
	case 's':
		for _, c := range str[:len(str)-1] {
			rune := ('🀐' + c - 49)
			ret += string(rune)
		}
		return ret
	case 'T':
		return "🀀"
	case 'N':
		return "🀁"
	case 'S':
		return "🀂"
	case 'P':
		return "🀃"
	case 'W':
		return "🀆"
	case 'G':
		return "🀅"
	case 'R':
		return "🀄"
	}
	return "a"
}
