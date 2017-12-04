package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	str "strings"
	uni "unicode"
)

/*Reference is the pattern name file to compare with */
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

func readFirstLine(filePath string) (string, error) {
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

/**
compare:

*/

func compare(lineRef string, line string) string {
	var b, c string

	if len(lineRef) > len(line) {
		b = lineRef
		c = line
	} else {
		b = line
		c = lineRef
	}
	result := make([]byte, len(c), len(b))
	for i, val := range []byte(b) {
		if c[i] == val {
			result[i] = byte('-')
		} else {
			result[i] = val
		}

	}
	return string(result)
}

func writeLine(file *os.File, line string) {
	w := bufio.NewWriter(file)
	n4, err := w.WriteString(line)
	check(err)
	fmt.Printf("Vrote %d bytes\n", n4)
	w.Flush()
}

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
		fmt.Println("Incorrects number of param Usage: fasta /path/of/folder output.txt")
	}

	// Obtaining the reference file and a slice with all the fasta files found
	referenceFile, fastas := OpenReferenceFile(folderPath)
	if referenceFile == "" {
		fmt.Println("Error: There is no reference file to compare with")
		fmt.Println("Nothing to create")
		os.Exit(1)
	}
	// Opening output file
	output, err := os.Create(outputPath)
	fmt.Println(">>", output)
	check(err)
	//It’s idiomatic to defer a Close immediately after opening a file.
	defer output.Close()

	// Opening reference file and take
	// its first line

	referenceLine, err := readFirstLine(referenceFile)
	check(err)

	for _, fas := range fastas {
		// we read the first line (and the only one)
		fastaLine, err := readFirstLine(fas)
		check(err)
		comparinsonLine := compare(referenceLine, fastaLine)
		writeLine(output, comparinsonLine)

	}
	// iterate over all the files
	// for each file comparing its first line and we write one line in the output
	// closing output file
}
