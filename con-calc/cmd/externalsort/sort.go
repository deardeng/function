package main

import (
	"bufio"
	"fmt"
	"function/con-calc/pipeline"
	"os"
	"strconv"
)

func main() {
	p := createNetPipeLine("large.in", 800000000, 4)
	writeToFile(p, "large.out")
	printFile("large.out")
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	p := pipeline.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}

func createPipeLine(filename string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func createNetPipeLine(filename string, fileSize, chunkCount int) <-chan int {
	pipeline.Init()
	chunkSize := fileSize / chunkCount
	sortAddr := []string{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(9000+i)
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	sortResults := []<-chan int{}
	for _, addr := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}
