package file

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
)

func GetMime() {

	file, err := os.Open("/Users/lanrion/Projects/go/gopath/src/github.com/lanrion/gopkg-examples/assets/weixin_pay.JPG")
	if err != nil {
		fmt.Println("Open file error: ", err.Error())
	}
	defer file.Close()

	//
	mime_type := mime.TypeByExtension(path.Ext(file.Name()))
	fmt.Println(mime_type)

	var buf [5]byte
	n, _ := io.ReadFull(file, buf[:])

	fmt.Printf("buf %b", buf[:n])

	sign := []byte("\x89\x50\x4E\x47\x0D\x0A\x1A\x0A")

	res := bytes.HasPrefix(buf[:n], sign)

	fmt.Println("res: ", res)

	ctype := http.DetectContentType(buf[:n])

	fmt.Println("buf len: ", len(buf[:n]))

	fmt.Println("ctype: ", ctype)
}
