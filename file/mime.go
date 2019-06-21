package file

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
)

func TestMime(){

	file, _ := os.Open("./qrcode/weixin_pay.JPG")
	defer file.Close()

	//
	mime_type := mime.TypeByExtension(path.Ext(file.Name()))
	fmt.Println(mime_type)

	var buf [5]byte
	n, _ := io.ReadFull(file, buf[:])
	ctype := http.DetectContentType(buf[:n])

	fmt.Println("buf len: ", len(buf))

	fmt.Println("ctype: ", ctype)
}
