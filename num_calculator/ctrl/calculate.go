package ctrl

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"num_calculator/service"

	"github.com/gin-gonic/gin"
)

type Num struct {
	detail []*Detail
}

type Detail struct {
	numStr string
	nums   *[]*NumInfo
}

type NumInfo struct {
	num     int    // 数字
	isPrime bool   // 是否是质数
	isEven  bool   // 是否是偶数
	isEx    bool   // 是都是补全的数
	rank    string // 位数
}

func Add0(str []string) Num {
	max := 0
	for _, s := range str {
		if len(s) > max {
			max = len(s)
		}
	}
	numInfo := make([]*Detail, 0)
	for _, s := range str {
		var niSlice = make([]*NumInfo, 0)
		if len(s) < max {
			num0 := max - len(s)
			for i := 0; i < num0; i++ {
				ni := &NumInfo{
					num:     0,
					isPrime: false,
					isEven:  false,
					isEx:    true,
				}
				niSlice = append(niSlice, ni)
			}
		}
		for _, i := range s {
			ni := &NumInfo{
				num:     int(i - 48),
				isPrime: service.IsPrime(int(i - 48)),
				isEven:  service.IsEven(int(i - 48)),
				isEx:    false,
			}
			niSlice = append(niSlice, ni)
		}
		d := &Detail{
			numStr: s,
			nums:   &niSlice,
		}
		numInfo = append(numInfo, d)
	}
	return Num{
		detail: numInfo,
	}
}

func (num *Num) CalRank() {
	for _, n := range num.detail {
		for i, info := range *n.nums {
			index := len(*n.nums) - i - 1
			if index >= len(service.Ranks) {
				info.rank = ""
				continue
			}
			info.rank = service.Ranks[index]
		}
	}
}

func (num *Num) BuildExcelRows() [][]string {
	nums := num.detail
	max := len(*nums[0].nums)
	rows := [][]string{}
	firstRow := []string{""}
	for i := max - 1; i >= 0; i-- {
		firstRow = append(firstRow, service.Ranks[i])
	}
	rows = append(rows, firstRow)
	for _, detail := range nums {
		r := []string{detail.numStr}
		for _, info := range *detail.nums {
			if info.isEx {
				r = append(r, "")
			} else {
				var s1 string
				if info.num == 1 || info.num == 0 {
					s1 += "      "
				} else {
					if info.isPrime {
						s1 += "质 "
					} else {
						s1 += "合 "
					}
				}
				if info.isEven {
					s1 += "偶"
				} else {
					s1 += "奇"
				}
				r = append(r, s1)
			}
		}
		rows = append(rows, r)
	}
	return rows
}

func Calculate(c *gin.Context) {
	con, _ := ioutil.ReadAll(c.Request.Body)
	//con := []byte("[\n    \"4\",\n    \"51\",\n    \"1\",\n    \"2\",\n    \"2231\",\n    \"38957\",\n    \"25567\"\n]")
	var nums = make([]string, 0)
	err := json.Unmarshal(con, &nums)
	if err != nil {
		return
	}

	num := Add0(nums)
	num.CalRank()
	rows := num.BuildExcelRows()
	service.GenerateExcel(rows, c)
}

func Download(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(404, "404")
		return
	}
	f, err := os.Open("/root/excel/" + filename)
	defer f.Close()
	if err != nil {
		c.JSON(404, "404")
		return
	}

	// 将文件读取出来
	data, err := ioutil.ReadAll(f)
	if err != nil {
		c.JSON(404, "404")
		return
	}
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+filename)
	c.Writer.Write(data)
}
