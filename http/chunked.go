package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/chunk/text", texHandler)
	http.HandleFunc("/video", videoHandler)
	http.HandleFunc("/image", imageHandler)
	http.HandleFunc("/hls/playlist/", hlsHandler)

	http.ListenAndServe(":3000", nil)
}

// 使用以下命令生成 m3u8以及ts文件
//ffmpeg -re -i text_chunk.mp4 -c copy -f hls -hls_time 9 -hls_list_size 0 -bsf:v h264_mp4toannexb output.m3u8
func hlsHandler(w http.ResponseWriter, r *http.Request) {

	switch path.Ext(r.URL.Path) {
	case ".m3u8":
		m3u8File, _ := os.Open("/Users/lanrion/Documents/test_hls/output.m3u8")
		buf := bytes.NewBuffer(nil)
		io.Copy(buf, m3u8File);
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Type", "application/x-mpegURL")
		w.Header().Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-cache")

		w.Write(buf.Bytes())
	case ".ts":
		fmt.Println("r.URL.Path: ", r.URL.Path)
		pathstr := r.URL.Path
		// 解析出实际的ts文件名称
		pathstr = strings.TrimLeft(pathstr, "/")
		paths := strings.SplitN(pathstr, "/", 3)

		tsPath := fmt.Sprintf("/Users/lanrion/Documents/test_hls/%s", paths[2])
		tsFile, _ := os.Open(tsPath)

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, tsFile);

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "video/mp2ts")
		w.Header().Set("Content-Length", strconv.Itoa(len(buf.Bytes())))
		w.Write(buf.Bytes())
	}
}

func imageHandler(w http.ResponseWriter, req *http.Request) {
	file, _ := os.Open("./assets/bigimg.jpg")
	defer file.Close()
	buffer := make([]byte, 200)
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", mime.TypeByExtension(file.Name()))
	w.WriteHeader(http.StatusOK)
	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		w.Write(buffer)
		w.(http.Flusher).Flush()

		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
}

func videoHandler(w http.ResponseWriter, req *http.Request)  {
	file, _ := os.Open("./assets/hugou.mp4")
	defer file.Close()

	buffer := make([]byte, 100)

	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Set("Content-Type", mime.TypeByExtension(file.Name()))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Disposition", "attachment;filename=FileName.mp4")

	for {
		bytesread, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		w.Write(buffer)
		w.(http.Flusher).Flush()

		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
}

func texHandler(w http.ResponseWriter, req *http.Request) {
	for i := 0 ; i < 3; i++ {
		w.Write([]byte("x"))
		log.Printf("echo \"x\" to client.")
		//time.Sleep(time.Duration(i+1 * 5) * time.Second)
		w.(http.Flusher).Flush()
	}
}