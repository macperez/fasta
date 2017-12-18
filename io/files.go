package io

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	str "strings"
	uni "unicode"
)

const Reference string = "referen"

/*
	WriteLine writes one line in a specific file
*/
func WriteLine(file *os.File, line string) {
	w := bufio.NewWriter(file)
	n4, err := w.WriteString(line)
	Check(err)
	fmt.Printf("Vrote %d bytes\n", n4)
	w.Flush()
}

/*
OpenReferenceFile lookfor the reference file and all the fasta files
*/
func OpenReferenceFile(folderPath string) (string, []string) {
	var fastas []string
	var reference string

	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if str.Contains(info.Name(), Reference) {
				reference = path
			} else if uni.IsDigit(rune(info.Name()[0])) {
				fastas = append(fastas, path)
			}

		}

		return nil
	})

	return reference, fastas
}

/*
ReadFirstLine is for reading the first line of a given file
*/
func ReadFirstLine(filePath string) (string, error) {
	file, err := os.Open(filePath)

	defer file.Close()
	if err != nil {
		return "", err
	}

	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}
	return line, scanner.Err()

}
