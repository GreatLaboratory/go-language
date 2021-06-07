package section6

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func ReadXlsx() {
	wb, err := xlsx.OpenFile("./files/sample.xlsx")
	if err != nil {
		panic(err)
	}
	for _, sh := range wb.Sheets {
		fmt.Println("Max row is", sh.MaxRow)
		sh.ForEachRow(rowVisitor)
	}
}
func cellVisitor(c *xlsx.Cell) error {
	value, err := c.FormattedValue()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Cell value:", value)
	}
	return err
}

func rowVisitor(r *xlsx.Row) error {
	return r.ForEachCell(cellVisitor)
}
