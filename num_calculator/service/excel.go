package service

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func GenerateExcel(rows [][]string, filename string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic at ")
		}
	}()
	excel := excelize.NewFile()
	sheet1 := "Sheet1"

	for i, row := range rows {
		for j, s := range row {
			xy := string(byte(j+65)) + strconv.Itoa(i+1)
			_ = excel.SetCellStr(sheet1, xy, s)
		}
	}

	err := excel.SaveAs("/root/excel/" + filename)
	if err != nil {
		fmt.Println(err)
	}
}
