package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("usage: mlsort $path\n"))
	}
	name:=os.Args[1]
	ss := readLinesFromFile(name)
	sort.Strings(ss)
	printStrings(ss)
}
func printStrings(ss []string) {
	for i := range ss {
        fmt.Println(ss[i])
	}
}

func readLinesFromFile(name string) []string {

	var ss []string
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	buf := bufio.NewReader(f)
	for {
		s, err := buf.ReadString('\n')
		strings.TrimRight(s,"\n")
		s  = strings.TrimSpace(s)
		if len(s) != 0 {
			ss = append(ss, s)
		}
		if err != nil {
			if err  == io.EOF {
				break
			}
			panic(err)
		}
	}
	return ss
}
