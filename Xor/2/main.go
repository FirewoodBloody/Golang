package main

import (
	"fmt"
	"image/png"
	"os"
	"strconv"

	"github.com/Luxurioust/excelize"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var XorKey []byte = []byte{0xA1, 0xB7, 0xAC, 0x57, 0x1C, 0x63, 0x3B, 0x81}

type Xor struct {
}

//加密  异或加密
func (a *Xor) enc(src string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	return result
}

//生成二维码
func Barcode(phone string) {
	xor := Xor{}
	qrCode, _ := qr.Encode(xor.enc(phone+",20200301,111111"), qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 124, 124)

	file, _ := os.Create("./images/" + phone + ".png")
	defer file.Close()

	png.Encode(file, qrCode)
}

func main() {
	excel, err := excelize.OpenFile("111.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows := excel.GetRows("Sheet1")
	for i := 1; i <= len(rows); i++ {
		cell := excel.GetCellValue("Sheet1", fmt.Sprintf("D%v", i))
		Barcode(cell)

		if err := excel.AddPicture("Sheet1", fmt.Sprintf("E%v", i), "./images/"+cell+".png", ""); err != nil {
			println(err.Error())
		}

	}
	// 保存文件
	if err = excel.Save(); err != nil {
		println(err.Error())
	}

}
