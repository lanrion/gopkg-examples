package qrcode

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"os"
)

//func init() {
//	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
//	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
//	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
//}

func TestScanner() {
	// open and decode image file
	//file, _ := os.Open("/Users/lanrion/Projects/go/gopath/src/github.com/lanrion/gopkg-examples/qrcode/weixin_pay.JPG")
	//file, _ := os.Open("/Users/lanrion/Downloads/tmp/WechatIMG4360.jpeg")
	//file, _ := os.Open("/Users/lanrion/Downloads/tmp/zhifubao.jpeg")
	file1 := "/Users/lanrion/Projects/go/gopath/src/github.com/lanrion/gopkg-examples/qrcode/qq_android_qrcode.gif"
	file, _ := os.Open(file1)
	img, fileName, err2 := image.Decode(file)

	fmt.Println("err2: ", err2, "fileName: ", fileName)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)
}
