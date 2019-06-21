package main

import (
	"errors"
	"fmt"
	"github.com/lanrion/gopkg-examples/file"
	"net/textproto"
	"strconv"
	"strings"
)

func main() {
	// sync.Log(os.Stdout, "path", "/search?q=flowers")
	// sync.TestPoolList()
	// container.TestList()
	// container.TestHeap()
	// container.TestRing()

	// errors.TestError()
	// sync.TestWaitGroup()
	// sync.TestOnce()
	// sync.TestMutex()
	// sync.TestCond()
	//data.TestArray()

	//file.TestMime()

	//err := TestDefer()
	//fmt.Println(err.Error())

	//var str = "hello 你好"
	//fmt.Println("len(str):", len(str))
	//strings.ToLower("ss")
	//data.CovertToInt("12")
	//qrcode.TestScanner()
	//file.TestSeek()

	//fmt.Println("")

	//ranges, err := parseRange("", 178953)
	//if err == nil {
	//	fmt.Println(ranges)
	//	fmt.Println(len(ranges))
	//	fmt.Println(sumRangesSize(ranges))
	//	//fmt.Println(ranges[0])
	//	//ra := ranges[0]
	//
	//	//fmt.Println(ranges[1])
	//
	//} else {
	//	fmt.Println("parse err: ", err)
	//}

	//fmt.Println(len(""))
	//var i string = ""

	file.TestMime()

}

func TestDefer() (err error) {
	defer func() {
		fmt.Print("1231232132")
	}()
	return errors.New("ssdf")
}

var errNoOverlap = errors.New("invalid range: failed to overlap")

// httpRange specifies the byte range to be sent to the client.
type httpRange struct {
	start, length int64
}

func (r httpRange) contentRange(size int64) string {
	return fmt.Sprintf("bytes %d-%d/%d", r.start, r.start+r.length-1, size)
}

func (r httpRange) mimeHeader(contentType string, size int64) textproto.MIMEHeader {
	return textproto.MIMEHeader{
		"Content-Range": {r.contentRange(size)},
		"Content-Type":  {contentType},
	}
}

func sumRangesSize(ranges []httpRange) (size int64) {
	for _, ra := range ranges {
		size += ra.length
	}
	return
}

func parseRange(s string, size int64) ([]httpRange, error) {
	if s == "" {
		return nil, nil // header not present
	}
	const b = "bytes="
	if !strings.HasPrefix(s, b) {
		return nil, errors.New("invalid range")
	}
	var ranges []httpRange
	noOverlap := false
	for _, ra := range strings.Split(s[len(b):], ",") {
		ra = strings.TrimSpace(ra)
		if ra == "" {
			continue
		}
		i := strings.Index(ra, "-")
		if i < 0 {
			return nil, errors.New("invalid range")
		}
		start, end := strings.TrimSpace(ra[:i]), strings.TrimSpace(ra[i+1:])
		var r httpRange
		if start == "" {
			// If no start is specified, end specifies the
			// range start relative to the end of the file.
			i, err := strconv.ParseInt(end, 10, 64)
			if err != nil {
				return nil, errors.New("invalid range")
			}
			if i > size {
				i = size
			}
			r.start = size - i
			r.length = size - r.start
		} else {
			i, err := strconv.ParseInt(start, 10, 64)
			if err != nil || i < 0 {
				return nil, errors.New("invalid range")
			}
			if i >= size {
				// If the range begins after the size of the content,
				// then it does not overlap.
				noOverlap = true
				continue
			}
			r.start = i
			if end == "" {
				// If no end is specified, range extends to end of the file.
				r.length = size - r.start
			} else {
				i, err := strconv.ParseInt(end, 10, 64)
				if err != nil || r.start > i {
					return nil, errors.New("invalid range")
				}
				if i >= size {
					i = size - 1
				}
				r.length = i - r.start + 1
			}
		}
		ranges = append(ranges, r)
	}
	if noOverlap && len(ranges) == 0 {
		// The specified ranges did not overlap with the content.
		return nil, errNoOverlap
	}
	return ranges, nil
}
