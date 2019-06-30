package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func countStringAppearInFile(path, str string) (result string, err error) {
	var count int

	file, err := os.Open(path)
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if strings.Contains(scan.Text(), str) {
			count++
		}
	}
	if scan.Err() == io.EOF {
		err = nil
	}

	result = fmt.Sprintf("String \"%s\" appears %d times in file.\n", str, count)
	return
}

func replaceStringsInFile(path, str1, str2 string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't read this file:\n %s", err))
	}

	newContent := strings.Replace(string(read), str1, str2, -1)

	err = ioutil.WriteFile(path, []byte(newContent), 0)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't write to this file:\n %s", err))
	}
	return nil
}

func main() {
	if len(os.Args) == 3 {
		result, err := countStringAppearInFile(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)

	} else if len(os.Args) == 4 {
		if err := replaceStringsInFile(os.Args[1], os.Args[2], os.Args[3]); err != nil {
			fmt.Println("Can't make a replacement in a reason: \n", err)
		}
		fmt.Printf("Finished replacing %s by %s.\n", os.Args[2], os.Args[3])
	} else {
		fmt.Println("Command-Line arguments required:\n" +
			"1. <file-path> <\"string\"> for counting string appearances in text file\n" +
			"2. <file-path> <\"string1\"> <\"string2\"> for replacing string1 with string2 in file.")
	}
}
