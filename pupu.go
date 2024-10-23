package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func FileReader(fname string) ([]string, error) {
	sl := make([]string, 0)
	file, err := os.Open(fname)
	defer file.Close()
	if err != nil {
		return sl, fmt.Errorf("Unable to open file: %s", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			sl = append(sl, strings.TrimRight(line, "\r\n"))
		}
	}
	return sl, err
}

func FindUniq(arr []string) []string {
	stlines := len(arr)
	if stlines >= 1000 {
		stlines /= 10
	} else if stlines > 100 {
		stlines = 100
	}
	dct := make(map[string]int, stlines)
	for _, el := range arr {
		dct[el]++
	}
	uniql := 0
	for _, co := range dct {
		if co == 1 {
			uniql++
		}

	}
	newlst := make([]string, uniql)
	for line, co := range dct {
		if co == 1 {
			newlst = append(newlst, strings.ToUpper(line))
		}
	}
	slices.Sort(newlst)
	return newlst
}

func writer(arr []string) error {
	file, err := os.Create("output.txt")
	if err != nil {
		return fmt.Errorf("Ошибка при создании файла %s", err)
	}
	defer file.Close()

	for _, line := range arr {
		if len(line) > 0 {
			ln := fmt.Sprintf("%s - %d байт \n", line, len(line))
			file.WriteString(ln)
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Передайте корректный путь до входного файла")
		return
	} else {
		sl, err := FileReader(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		uni := FindUniq(sl)
		err = writer(uni)
		if err != nil {
			fmt.Println(err)
		}
	}
}
