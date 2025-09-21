package main

import (
	"fmt"
	"sync"

	"github.com/xuri/excelize/v2"
)

func generateExcel(wg *sync.WaitGroup, fileName string, rowCount int) {
	defer wg.Done()
	f := excelize.NewFile()
	sheet := "Sheet1"
	index, _ := f.NewSheet(sheet)

	// Header
	f.SetCellValue(sheet, "A1", "ID")
	f.SetCellValue(sheet, "B1", "Name")
	f.SetCellValue(sheet, "C1", "Email")

	for i := 1; i <= rowCount; i++ {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+1), i)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+1), fmt.Sprintf("User_%d", i))
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+1), fmt.Sprintf("user%d@example.com", i))
		if i%10000 == 0 {
			fmt.Printf("Generated %d rows...\n", i)
		}
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs(fileName); err != nil {
		panic(err)
	}
	fmt.Printf("✅ File created: %s (%d rows)\n", fileName, rowCount)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	generateExcel(&wg, "small.xlsx", 50000)
	generateExcel(&wg, "medium.xlsx", 1000000)
	generateExcel(&wg, "large.xlsx", 5000000)
	generateExcel(&wg, "import.xlsx", 10000)

	wg.Wait()
}
