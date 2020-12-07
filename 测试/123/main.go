package main

import (
	"bytes"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
	"github.com/golang/freetype"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
)

type TMainForm struct {
	*vcl.TForm
	CbbPrinters *vcl.TComboBox
	Btn1        *vcl.TButton
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm)
	vcl.Application.Run()

}

//绘制 顺丰快递 面单信息
func NewParseFont(img *image.Gray) (*image.Gray, error) {

	//绘制条形码
	cs, _ := code128.Encode("SF1042356749791")
	// 设置图片像素大小
	qrCode_t, _ := barcode.Scale(cs, 500, 100)
	// 将code128的条形码编码为png图片

	for x := 0; x < qrCode_t.Bounds().Dx(); x++ {
		for y := 0; y < qrCode_t.Bounds().Dy(); y++ {
			img.Set(x, y+90, qrCode_t.At(x, y))
		}
	}

	//绘制二维码
	er, _ := qr.Encode("MMM={'k1':'377VA','k2':'377JA','k3':'037','k4':'T4','k5':'SF1042356749791','k6':'','k7':'4d68ca6b'}", qr.M, qr.Auto)
	// 设置图片像素大小
	qrCode_e, _ := barcode.Scale(er, 168, 168)
	for x := 0; x < qrCode_e.Bounds().Dx(); x++ {
		for y := 0; y < qrCode_e.Bounds().Dy(); y++ {
			img.Set(x+266, y+476, qrCode_e.At(x, y))
		}
	}

	//开始绘制快递面单 收件人和寄件人 文字信息
	//读取字体数据
	fontBytes, err := ioutil.ReadFile("./msyh.ttc")
	if err != nil {
		return img, err
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return img, err
	}
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色(红色)
	f.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))

	//代收货款
	str := "代收货款"
	f.SetFontSize(20)
	pt := freetype.Pt(588, 340+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString(str, pt)
	if err != nil {
		return img, err
	}
	str = "20000"
	f.SetFontSize(20)
	pt = freetype.Pt(588, 370+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString(str, pt)
	if err != nil {
		return img, err
	}
	//代收回款图片绘制  代收货款条件
	//if a==""{}
	file_COD, _ := os.Open("./image/COD.png")
	img_COD, _ := png.Decode(file_COD)
	for x := 1; x < img_COD.Bounds().Dx(); x++ {
		for y := 1; y < img_COD.Bounds().Dy(); y++ {
			img.Set(x+588, y+266, img_COD.At(x, y))
		}
	}

	//设置  大头笔  位置
	f.SetFontSize(60)
	pt = freetype.Pt(28, 294+int(f.PointToFixed(70))>>8) //设置大头笔位置
	f.DrawString("377JA-037", pt)
	if err != nil {
		return img, err
	}

	//设置 收件人 姓名 位置
	str = "柴雪新"
	name_s := []rune(str) //
	f.SetFontSize(20)
	pt = freetype.Pt(112, 364+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString(string(name_s[:len(name_s)-len(name_s)/2-1])+"**", pt)
	if err != nil {
		return img, err
	}

	//设置 收件人 电话 位置
	f.SetFontSize(20)
	pt = freetype.Pt(280, 364+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString("178****8284", pt)
	if err != nil {
		return img, err
	}

	//设置 收件人 地址 位置
	str = "陕西省临汾市浮山县北韩乡北韩村8号"
	address_j := []rune(str)
	f.SetFontSize(20)
	pt = freetype.Pt(112, 420+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString(string(address_j[:len(address_j)-len(address_j)/3*2])+"******", pt)
	if err != nil {
		return img, err
	}

	//设置 付款方式位置 位置
	f.SetFontSize(20)
	pt = freetype.Pt(98, 504+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString("寄付月结", pt)
	if err != nil {
		return img, err
	}

	//设置 已验视 位置
	f.SetFontSize(30)
	pt = freetype.Pt(465, 504+int(f.PointToFixed(30))>>8)
	_, err = f.DrawString("已", pt)
	if err != nil {
		return img, err
	}
	pt = freetype.Pt(465, 560+int(f.PointToFixed(32))>>8)
	_, err = f.DrawString("验", pt)
	if err != nil {
		return img, err
	}
	pt = freetype.Pt(465, 609+int(f.PointToFixed(32))>>8)
	_, err = f.DrawString("视", pt)
	if err != nil {
		return img, err
	}

	//设置 寄件人姓名 位置
	f.SetFontSize(20)
	pt = freetype.Pt(112, 675+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString("柴雪新", pt)
	if err != nil {
		return img, err
	}

	//设置 寄件人电话 位置
	f.SetFontSize(20)
	pt = freetype.Pt(280, 675+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString("178****8284", pt)
	if err != nil {
		return img, err
	}

	//设置 寄件人地址 位置
	str = "陕西省临汾市浮山县北韩乡北韩村8号"
	address_s := []rune(str)
	f.SetFontSize(20)
	pt = freetype.Pt(112, 714+int(f.PointToFixed(20))>>8)
	_, err = f.DrawString(string(address_s[:len(address_s)-len(address_s)/3*2])+"******", pt)
	if err != nil {
		return img, err
	}

	//设置 订单号 位置
	f.SetFontSize(24)
	pt = freetype.Pt(28, 770+int(f.PointToFixed(24))>>8)
	_, err = f.DrawString(fmt.Sprintf("订单号：%v", 0), pt)
	if err != nil {
		return img, err
	}

	//设置 托寄物 位置
	f.SetFontSize(24)
	pt = freetype.Pt(28, 812+int(f.PointToFixed(24))>>8)
	_, err = f.DrawString(fmt.Sprintf("托寄物：%v", 0), pt)
	if err != nil {
		return img, err
	}

	return img, err
}

//绘制顺丰快递电子面单图片
func SF_image_init() *image.Gray {

	ima := image.NewGray(image.Rect(0, 0, 700, 1050))
	for x := 0; x <= 700; x++ {
		for y := 0; y <= 1050; y++ {
			ima.Set(x, y, color.White) //绘制边框 横线
		}
	}

	//边框
	//////////////////////////////////////////
	for x := 0; x <= 700; x++ {
		ima.Set(x, 252, color.RGBA{0, 0, 0, 0}) //绘制边框 横线
	}
	for x := 0; x <= 700; x++ {
		ima.Set(x, 1036, color.RGBA{0, 0, 0, 0}) //绘制边框 横线
	}

	//for y := 180; y <= 740; y++ {
	//	ima.Set(10, y, color.RGBA{0, 0, 0, 0}) //绘制边框 竖线
	//}
	//
	//for y := 180; y <= 740; y++ {
	//	ima.Set(490, y, color.RGBA{0, 0, 0, 0}) //绘制边框 竖线
	//}
	////////////////////////////////

	/////////////////////内线横线
	for x := 0; x <= 700; x++ {
		ima.Set(x, 469, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	for x := 0; x <= 700; x++ {
		ima.Set(x, 651, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	for x := 0; x <= 700; x++ {
		ima.Set(x, 742, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	for x := 0; x <= 700; x++ {
		ima.Set(x, 971, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	/////////////////////内线横线短
	for x := 0; x <= 259; x++ {
		ima.Set(x, 532, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	for x := 0; x <= 259; x++ {
		ima.Set(x, 588, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	for x := 518; x <= 700; x++ {
		ima.Set(x, 560, color.RGBA{0, 0, 0, 0}) //内线横线
	}

	/////////////////////内线竖线
	for y := 469; y <= 651; y++ {
		ima.Set(259, y, color.RGBA{0, 0, 0, 0}) //绘制边框 竖线
	}

	for y := 469; y <= 651; y++ {
		ima.Set(441, y, color.RGBA{0, 0, 0, 0}) //绘制边框 竖线
	}

	for y := 469; y <= 651; y++ {
		ima.Set(518, y, color.RGBA{0, 0, 0, 0}) //绘制边框 竖线
	}

	//寄件人 图片绘制
	file, _ := os.Open("./image/寄icon.png")
	img_ji, _ := png.Decode(file)
	for x := 1; x < img_ji.Bounds().Dx(); x++ {
		for y := 1; y < img_ji.Bounds().Dy(); y++ {
			ima.Set(x+20, y+660, img_ji.At(x, y))
		}
	}

	//收件人图片绘制
	file, _ = os.Open("./image/收icon.png")
	img_shou, _ := png.Decode(file)
	for x := 1; x < img_shou.Bounds().Dx(); x++ {
		for y := 1; y < img_shou.Bounds().Dy(); y++ {
			ima.Set(x+20, y+364, img_shou.At(x, y))
		}
	}

	//顺丰客服电话图片绘制
	file, _ = os.Open("./image/qiao.png")
	img_qiao, _ := png.Decode(file)
	for x := 1; x < img_qiao.Bounds().Dx(); x++ {
		for y := 1; y < img_qiao.Bounds().Dy(); y++ {
			ima.Set(x+720, y+10, img_qiao.At(x, y))
		}
	}

	//sflogo图片绘制
	file, _ = os.Open("./image/sflogo.png")
	img_sflogo, _ := png.Decode(file)
	for x := 1; x < img_sflogo.Bounds().Dx(); x++ {
		for y := 1; y < img_sflogo.Bounds().Dy(); y++ {
			ima.Set(x+10, y+10, img_sflogo.At(x, y))
		}
	}

	return ima
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.Btn1 = vcl.NewButton(f)
	f.Btn1.SetParent(f)
	f.Btn1.SetBounds(10, 10, 88, 28)
	f.Btn1.SetCaption("打印")
	f.Btn1.SetOnClick(f.OnButtonClick)

	f.CbbPrinters = vcl.NewComboBox(f)
	f.CbbPrinters.SetParent(f)
	f.CbbPrinters.SetBounds(f.Btn1.Left(), f.Btn1.Top()+f.Btn1.Height()+20, 200, f.CbbPrinters.Height())
	f.CbbPrinters.Items().Assign(vcl.Printer.Printers())
	f.CbbPrinters.SetOnChange(f.OnCbbChange)

	// Printer.Orientation 方向

}

func (f *TMainForm) OnButtonClick(sender vcl.IObject) {

	if f.CbbPrinters.ItemIndex() == -1 {
		vcl.ShowMessage("先设置一个打印机.")
		return
	}

	// 注意：由于打印机的Canvas与系统的DPI有关，所以很多都需要通过换算
	// vcl.Screen.PixelsPerInch()

	rx := float64(vcl.Screen.PixelsPerInch()) / 96.0
	fmt.Println("vcl.Screen.PixelsPerInch(): ", vcl.Screen.PixelsPerInch(), ", rx: ", rx)
	//vcl.Printer.SetTitle("标题啊。。。。 ")

	vcl.Printer.BeginDoc()
	defer vcl.Printer.EndDoc()

	canvas := vcl.Printer.Canvas()
	//canvas.Brush().SetColor(colors.ClWhite)
	canvas.Brush().SetStyle(types.BsClear)
	// 这里画个图片
	img := SF_image_init()
	img, _ = NewParseFont(img)
	//ima := resize.Thumbnail(1000, 1500, img, resize.Lanczos2)
	buf := new(bytes.Buffer)
	//file, _ := os.Create("./12.jpg")
	jpeg.Encode(buf, img, nil)
	jpgImg := vcl.NewJPEGImage()
	stream := vcl.NewMemoryStreamFromBytes(buf.Bytes())
	stream.SetPosition(0)
	jpgImg.LoadFromStream(stream)
	//canvas.Draw(0, 0, jpgImg)
	canvas.StretchDraw(types.TRect{0, 0, int32(float64(jpgImg.Width()) * rx), int32(float64(jpgImg.Height()) * rx)}, jpgImg)
	stream.Free()
	jpgImg.Free()

}

func (f *TMainForm) OnCbbChange(sender vcl.IObject) {
	if f.CbbPrinters.ItemIndex() != -1 {
		vcl.Printer.SetPrinterIndex(f.CbbPrinters.ItemIndex())
		fmt.Println(f.CbbPrinters.ItemIndex())
	}
}
