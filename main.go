package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	str "strings"
	uni "unicode"
)

const Reference string = "referen"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

/*
OpenReferenceFile lookfor the reference file and all the fasta files
*/
func OpenReferenceFile(folderPath string) (string, []string) {
	//files, err := ioutil.ReadDir(folderPath)
	// walk all files in directory

	var fastas []string
	var reference string
	filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if str.Contains(info.Name(), "referen") {
				reference = filepath.Join(path, info.Name())
			} else if uni.IsDigit(rune(info.Name()[0])) {
				fastas = append(fastas, filepath.Join(path, info.Name()))
			}

		}

		return nil
	})

	return reference, fastas
}

// func ReadFirstLine(file os.FileInfo) (string, error) {
// 	file, err := os.Open(file)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()
//
// 	var line string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
//
// 		line = scanner.Text()
// 	}
// 	return line, scanner.Err()
// }

func main() {

	// The folder name that constains the fasta files is stored here
	// The same for output path in which it will be written the pattern
	var folderPath string
	var outputPath string

	args := os.Args
	if len(args) == 3 {
		folderPath = args[1]
		outputPath = args[2]
	} else {
		fmt.Println("Incorrect number of param Usage: fasta /path/of/folder output.txt")
	}

	// Obtaining the reference file and a slice with all the fasta files found
	ref, fastas := OpenReferenceFile(folderPath)
	if ref == "" {
		fmt.Println("Error: There is no reference file to compare with")
		fmt.Println("Nothing to create")
		os.Exit(1)
	}
	// Opening output file
	f, err := os.Create(outputPath)
	check(err)
	//It’s idiomatic to defer a Close immediately after opening a file.
	defer f.Close()

	for _, fas := range fastas {
		// we read the first line (and the only one)
		fmt.Println(">>", fas)
		// line, err := ReadFirstLine(fas)
		// check(err)
		// fmt.Println(line)
		// fmt.Println("-------------------------------------------------------")

	}
	// iterate over all the files
	// for each file comparing its first line and we write one line in the output
	// closing output file
}
