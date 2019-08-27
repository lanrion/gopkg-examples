package file

import (
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"strings"
)

var errSeeker = errors.New("seeker can't seek")

func sizeFunc(content io.ReadSeeker) (int64, error) {
	size, err := content.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, errSeeker
	}
	_, err = content.Seek(0, io.SeekStart)
	if err != nil {
		return 0, errSeeker
	}
	return size, nil
}

func TestSeek() {
	file, err := os.Open("./assets/weixin_pay.JPG")
	defer file.Close()
	if err != nil {
		fmt.Println("TestSeek error: ", err)
	} else {
		size, err := sizeFunc(file)
		if err == nil {
			fmt.Println("Size: ", size)
		}
		file.Seek(0, 10)

		//fmt.Println("TestSeek: ", len())
	}

	reader := strings.NewReader("Go语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}

func MultipleFile() {
	path := "/Users/lanrion/Projects/go/gopath/src/github.com/lanrion/gopkg-examples/assets/bigimg.jpg"
	fi, err := os.Open(path)
	defer fi.Close()
	if err != nil {
		panic(err)
	}
	_, name, _ := image.Decode(fi)
	fmt.Println("name: ", name)

	_, err = fi.Seek(0, io.SeekStart)
	_, name, err = image.Decode(fi)
	if err != nil {
		panic(err)
	}
	fmt.Println("name2: ", name)
	


}

