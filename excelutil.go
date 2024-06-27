package main

import (
	"fmt"
	"sort"

	"github.com/tealeg/xlsx"
)

func printResultsToExcel(keysInFiles map[string][]string, fileName string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Common Keys")
	if err != nil {
		return fmt.Errorf("error creating sheet: %w", err)
	}

	addHeaderRow(sheet)

	keyCounts := getKeyCounts(keysInFiles)
	sortKeyCountsDescending(keyCounts)

	for _, kc := range keyCounts {
		if len(keysInFiles[kc.Key]) > 1 {
			addDataRow(sheet, kc.Key, keysInFiles[kc.Key])
		}
	}

	if err := file.Save(fileName); err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}
	fmt.Printf("Results saved to %s\n", fileName)
	return nil
}

func addHeaderRow(sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	row.AddCell().Value = "Key"
	row.AddCell().Value = "Files"
}

func addDataRow(sheet *xlsx.Sheet, key string, files []string) {
	row := sheet.AddRow()
	row.AddCell().Value = key
	filesCell := row.AddCell()
	for _, file := range files {
		filesCell.Value += file + "; "
	}
}

func getKeyCounts(keysInFiles map[string][]string) []FileKeyInfo {
	var keyCounts []FileKeyInfo
	for key, files := range keysInFiles {
		keyCounts = append(keyCounts, FileKeyInfo{Key: key, Files: files})
	}
	return keyCounts
}

func sortKeyCountsDescending(keyCounts []FileKeyInfo) {
	sort.Slice(keyCounts, func(i, j int) bool {
		return len(keyCounts[i].Files) > len(keyCounts[j].Files)
	})
}
