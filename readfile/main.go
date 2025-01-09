package main

import (
	"fmt"
	"os"
)

func main() {
	// readfile
	chLine := make(chan struct {
		linno int
		line  string
	})
	go func() {
		f, err := os.OpenFile("test.txt", os.O_RDONLY, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		buf := make([]byte, 1)
		line := make([]byte, 0)
		lineno := 1
		for {
			_, err := f.Read(buf)
			if err != nil {
				//fmt.Println(string(line))
				//fmt.Println(err)
				chLine <- struct {
					linno int
					line  string
				}{lineno, string(line)}
				close(chLine)
				break
			}
			if buf[0] == '\n' {
				fmt.Println(string(line))
				chLine <- struct {
					linno int
					line  string
				}{lineno, string(line)}
				line = make([]byte, 0)
				lineno++
			} else {
				line = append(line, buf[0])
			}
			//fmt.Println(string(buf[:n]))
		}
	}()

	for line := range chLine {
		fmt.Println(line.linno, "-->", line.line)
	}

}
