package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/chunk/text", texHandler)
	http.HandleFunc("/video", videoHandler)
	http.ListenAndServe(":3000", nil)
}

func imageHandler(w http.ResponseWriter, req *http.Request) {
	file, _ := os.Open("./assets/weixin_pay.JPG")
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