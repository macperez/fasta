package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	str "strings"
	uni "unicode"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func OpenReferenceFile(folderPath string) (os.FileInfo, []os.FileInfo) {
	files, err := ioutil.ReadDir(folderPath)
	check(err)
	var fastas []os.FileInfo
	fmt.Printf("here")
	var reference os.FileInfo
	for _, f := range files {
		if str.Contains(f.Name(), "referen") {
			reference = f
		} else if uni.IsDigit(rune(f.Name()[0])) {
			fastas = append(fastas, f)
		}

	}
	return reference, fastas
}

func main() {
	// capturamos argumentos de la carpeta
	var folderPath string
	var outputPath string

	args := os.Args
	if len(args) == 3 {
		folderPath = args[1]
		outputPath = args[2]
	} else {
		fmt.Println("Incorrect number of param Usage: fasta /path/of/folder output.txt")
	}

	fmt.Println("Folder: ", folderPath)
	fmt.Println("output: ", outputPath)
	// abrimos el fichero de referencia y lo procesamos
	// reference, fastas := OpenReferenceFile(folderPath)
	ref, fastas := OpenReferenceFile(folderPath)
	fmt.Println("reference file is ", ref.Name())
	fmt.Println("------------------------------")
	for _, fas := range fastas {
		fmt.Println(">>", fas.Name())
	}
	// abrimoms fichero salida
	// recorremos todos ficheros
	// por cada fichero se compara y se escribe en la salida.
	// se cierra fichero de salida
}
