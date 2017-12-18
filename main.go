package main

import (
	"fmt"
	"os"

	"github.com/macperez/fasta/io"
)

/**
compare:

*/

func compare(lineRef string, line string) string {
	var b, c string

	if len(lineRef) >= len(line) {
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
		os.Exit(1)
	}

	// Obtaining the reference file and a slice with all the fasta files found
	referenceFile, fastas := io.OpenReferenceFile(folderPath)
	if referenceFile == "" {
		fmt.Println("Error: There is no reference file to compare with")
		fmt.Println("Nothing to create")
		os.Exit(1)
	}
	// Opening output file
	output, err := os.Create(outputPath)
	fmt.Println(">>", output)
	io.Check(err)
	//It’s idiomatic to defer a Close immediately after opening a file.
	defer output.Close()

	// Opening reference file and take
	// its first line

	referenceLine, err := io.ReadFirstLine(referenceFile)
	io.Check(err)

	for _, fas := range fastas {
		// we read the first line (and the only one)
		fastaLine, err := io.ReadFirstLine(fas)
		io.Check(err)
		comparinsonLine := compare(referenceLine, fastaLine)
		io.WriteLine(output, comparinsonLine)

	}
	// iterate over all the files
	// for each file comparing its first line and we write one line in the output
	// closing output file
}
