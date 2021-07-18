package ctrl

import (
	"testing"
)

func TestAdd0(t *testing.T) {
	var nums = make([]string, 0)
	nums = append(nums, "1", "123", "241", "12000", "487", "11", "90")
	num := Add0(nums)
	num.CalRank()
	rows := num.BuildExcelRows()
	rows = rows
	//service.GenerateExcel(rows)
	//fmt.Println(num)
}
