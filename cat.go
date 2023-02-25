package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(fPath string) error {
	file, err := os.Open(fPath)
	if err != nil {
		return err
	}
	defer file.Close()
	r := bufio.NewScanner(file)
	for r.Scan() {
		fmt.Println(r.Text())
	}
	if err := r.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("ファイル名を指定してください")
	} else {
		for i, v := range os.Args {
			if i >= 1 {
				err := readFile(v)
				fmt.Println(err)
			}
		}
	}
}
