package main

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"strconv"
	"time"
	"algorithm"
	"io"
)

var infile *string = flag.String("i", "infile", "File contains value for sorting.")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values.")
var algo *string = flag.String("a", "qsort", "algorithm to sort.")


func readValues(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("failed to open ", filePath)
		return nil, nil
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values := make([]int, 0)
	for {
		context, isPrefix, err1 := br.ReadLine()
		
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		
		if isPrefix {
			fmt.Println("the line readed is too long.")
			break
		}

		str := string(context)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			break
		}

		values = append(values, value)

	}
	return values, err
}

func writeValues(filePath string, values []int) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("failed to create file: ", filePath)
		return err
	}
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	flag.Parse()
	if infile == nil {
		fmt.Println("解析命令行参数失败！")
		return
	}
	fmt.Println("infile:", *infile, "outfile:", *outfile, "algorithm:", *algo)
	t1 := time.Now()
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println("读取数组失败！", err)
		return
	}
	fmt.Println("old array:", values)
	switch *algo {
		case "qsort":
			algorithm.QuickSort(values)
		case "bubblesort":
			algorithm.BubbleSort(values)
		default:
			fmt.Println("Sorting algorthm", *algo, "is either unkonw or unsupported.")
	}
	fmt.Println("sorted array:", values)
	t2 := time.Now()
	fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
	writeValues(*outfile, values)
}