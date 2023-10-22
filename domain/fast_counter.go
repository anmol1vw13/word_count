package domain

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

const fileBufferSize = 10

func (c Counter) Start() {

	fileBufferChannel := make(chan int, fileBufferSize)
	wg := &sync.WaitGroup{}

	if len(c.Files) == 0 {
		c.Files = []string{"-"}
	}

	for _, f := range c.Files {
		fmt.Println("File is", f)
		fileBufferChannel <- 1
		wg.Add(1)
		go c.processFile(f, fileBufferChannel, wg)
	}
	wg.Wait()
}

func (c Counter) processFile(fileName string, fileBufferChannel chan int, wg *sync.WaitGroup) {

	lineChan := make(chan string)
	errChan := make(chan error)

	defer func() { <-fileBufferChannel }()
	defer wg.Done()

	var scanner *bufio.Scanner
	if fileName != "-" {
		file, err := os.Open(fileName)
		if err != nil {
			errChan <- err
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)

	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	scanner.Buffer(make([]byte, chunkSize), chunkSize)

	go func() {
		defer close(lineChan)
		defer close(errChan)
		for scanner.Scan() {
			line := scanner.Text()
			lineChan <- line
		}

		if err := scanner.Err(); err != nil {
			errChan <- err
		}
	}()

	countInfo, err := count(lineChan, errChan)

	if c.Flag.Line {
		fmt.Printf("%d\t", countInfo.Line)
	}
	if c.Flag.Word {
		fmt.Printf("%d\t", countInfo.Word)
	}
	if c.Flag.Char {
		fmt.Printf("%d\t", countInfo.Char)
	}
	fmt.Printf("%s\n", fileName)

	if err != nil {
		fmt.Println(err)
	}

}

func count(lineChan chan string, errorChan chan error) (CountInfo, error) {

	countInfo := CountInfo{}

	for {

		select {
		case line, ok := <-lineChan:
			if !ok {
				return countInfo, nil
			}
			countInfo.Line++
			countInfo.Word = countInfo.Word + int64(len(strings.Fields(line)))
			countInfo.Char = countInfo.Char + int64(len(line))

		case err := <-errorChan:
			if err != nil {
				return countInfo, err
			}
		}

	}
}
