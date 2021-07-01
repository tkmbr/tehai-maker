package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var dict = map[int]string{
	int('T'): "🀀",
	int('N'): "🀁",
	int('S'): "🀂",
	int('P'): "🀃",
}

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
	case "", "-":
		r = os.Stdin
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
	// fmt.Print(string(t))

	for _, char := range string(t) {
		fmt.Print(dict[int(char)])
	}

	return nil
}
