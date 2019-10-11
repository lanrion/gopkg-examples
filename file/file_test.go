package file

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMime(t *testing.T) {
	GetMime()
}

func TestMultipleFile(t *testing.T) {
	MultipleFile()
}

func TestStat(t *testing.T) {
	file, err := os.Open("/Users/lanrion/Projects/go/gopath/src/github.com/lanrion/gopkg-examples/assets/weixin_pay.JPG")
	defer file.Close()

	if err != nil {
		fmt.Println("TestSeek error: ", err)
	} else {
		fileInfo, _ := file.Stat()
		fmt.Println(fileInfo.ModTime().Unix())
	}

}
