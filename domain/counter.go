package domain

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const chunkSize = 1024 * 1024

func (c Counter) Count() ([]CountInfo, error) {

	if !(c.Flag.Char || c.Flag.Word || c.Flag.Line) {
		c.Flag.Char = true
		c.Flag.Word = true
		c.Flag.Line = true
	}

	countInfoArr := make([]CountInfo, 0)
	if len(c.Files) != 0 {
		for _, f := range c.Files {
			countInfoRes, err := c.countFile(f)
			if err == nil {
				countInfoArr = append(countInfoArr, *countInfoRes)
			} else {
				return countInfoArr, err
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Buffer(make([]byte, chunkSize), chunkSize)

		countInfoRes, err := c.scanAndCount("", scanner)
		if err != nil {
			countInfoArr = append(countInfoArr, *countInfoRes)
		} else {
			return countInfoArr, err
		}
	}
	return countInfoArr, nil
}

func (c Counter) countFile(f string) (*CountInfo, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Buffer(make([]byte, chunkSize), chunkSize)
	return c.scanAndCount(f, fileScanner)
}

func (c Counter) scanAndCount(f string, scanner *bufio.Scanner) (*CountInfo, error) {
	countFile := &CountInfo{Identifier: f}
	for scanner.Scan() {

		text := scanner.Text()
		fmt.Println("You have entered ", text)
		if c.Flag.Line {
			(countFile.Line)++
		}
		if c.Flag.Word {
			re := regexp.MustCompile("\\s+")
			countFile.Word = countFile.Word + int64(len(re.Split(text, -1)))
		}
		if c.Flag.Char {
			countFile.Char = countFile.Char + int64(len(text))
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
	return countFile, nil
}
