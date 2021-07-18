package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
)

func GenerateExcel(rows [][]string, c *gin.Context) {
	excel := excelize.NewFile()
	sheet1 := "Sheet1"

	for i, row := range rows {
		for j, s := range row {
			xy := string(byte(j+65)) + strconv.Itoa(i+1)
			_ = excel.SetCellStr(sheet1, xy, s)
		}
	}
	filename := time.Now().Format("20060102150405") + ".xlsx"
	err := excel.SaveAs("/root/excel/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"data": filename,
	})
}
