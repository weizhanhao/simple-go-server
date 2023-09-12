package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CoverageData struct {
	StartRow    int64
	StartColumn int64
	EndRow      int64
	EndColumn   int64
	Module      int64
	Count       int64
}

type CoverageRow struct {
	CoverageRowCount int64
	AllRowCount      int64
}

func main333() {
	filePath := "coveragenew1.out" // 覆盖率文件的路径

	var coverageDataList []CoverageData

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	coverRow := int64(0)
	allRow := int64(0)
	for scanner.Scan() {
		thisLine := scanner.Text()
		separator := "go:"
		result := strings.Split(thisLine, separator)
		if len(result) < 2 {
			continue
		}
		line := result[1]
		line = strings.Replace(line, ".", " ", -1)
		line = strings.Replace(line, ",", " ", -1)
		if strings.TrimSpace(line) != "" {
			fields := strings.Split(line, " ")

			if len(fields) >= 6 {
				startRow, err := strconv.ParseInt(fields[0], 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse startRow: %s", err)
				}

				startColumn, err := strconv.ParseInt(fields[1], 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse startColumn: %s", err)
				}

				endRow, err := strconv.ParseInt(fields[2], 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse endRow: %s", err)
				}

				endColumn, err := strconv.ParseInt(fields[3], 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse endColumn: %s", err)
				}

				module, err := strconv.ParseInt(fields[4], 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse module: %s", err)
				}

				count, err := strconv.ParseInt(fields[5], 10, 64)
				if err != nil {
					log.Fatalf("Failed to parse count: %s", err)
				}
				thisLineHaveTheRow := endRow - startRow + 1
				allRow += thisLineHaveTheRow
				if count > 0 {
					coverRow += thisLineHaveTheRow
				}

				coverageData := CoverageData{
					StartRow:    startRow,
					StartColumn: startColumn,
					EndRow:      endRow,
					EndColumn:   endColumn,
					Module:      module,
					Count:       count,
				}

				coverageDataList = append(coverageDataList, coverageData)
			}
		}
	}

	coverageRow := CoverageRow{
		CoverageRowCount: coverRow,
		AllRowCount:      allRow,
	}

	fmt.Println(coverageRow)
	//if err := scanner.Err(); err != nil {
	//	log.Fatalf("Failed to read file: %s", err)
	//}
	//
	//for _, data := range coverageDataList {
	//	fmt.Printf("StartRow: %d, StartColumn: %d, EndRow: %d, EndColumn: %d, Module: %s, Count: %d\n",
	//		data.StartRow, data.StartColumn, data.EndRow, data.EndColumn, data.Module, data.Count)
	//}
	//fmt.Println(1111)
}
