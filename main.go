package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	s := http.NewServeMux()
	s.HandleFunc("/", indexHandler)
	s.HandleFunc("/file", fileHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), s)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	imageStr := string(b)
	splitStr := strings.Split(imageStr, ",")
	formattedImgStr := splitStr[1]
	imageFormat := splitStr[0][11:14]
	fmt.Println(imageFormat)
	// fmt.Println(b)
	// reader := bytes.NewReader(b)
	// img, err := png.Decode(reader)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	f, err := os.Create("img/img2.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	decodeResult, err := b64Decode(formattedImgStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = f.Write(decodeResult)
	if err != nil {
		fmt.Println(err)
		return
	}
	// png.Encode(f, img)
}

func b64Encode(bytes []byte) string {

	x := base64.StdEncoding.EncodeToString(bytes)
	return x
}
func b64Decode(inputStr string) ([]byte, error) {
	x, err := base64.StdEncoding.DecodeString(inputStr)
	if err != nil {
		return nil, err
	}
	return x, nil
}
