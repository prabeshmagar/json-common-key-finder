package main

import (
	"flag"
	"fmt"
)

func main() {
	var outputFileName string
	flag.StringVar(&outputFileName, "output", "CommonKeys.xlsx", "Specifies the name of the output Excel file where the common keys will be saved. Default is 'CommonKeys.xlsx'.")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("No input files provided. Usage: common-key-finder --output=name_of_file.xlsx file1.json file2.json ...")
		return
	}

	keysInFiles, err := processFiles(files)
	if err != nil {
		fmt.Println("Error processing files:", err)
		return
	}

	if err := printResultsToExcel(keysInFiles, outputFileName); err != nil {
		fmt.Println("Error saving results:", err)
	}
}
