package util

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(filename string) []string{
	_,err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err){
			fmt.Println("[-] Dict file " + filename + " not found")
			os.Exit(-1)
		}
	}

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(-1)
	}

	defer fi.Close()

	var lines[] string
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}

