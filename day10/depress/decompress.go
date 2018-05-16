package main

import (
	"os"
	"compress/gzip"
	"archive/tar"
	"io"
	"fmt"
)



func main() {
	// file read
	fr, err := os.Open("lin_golang_src.tar.gz")
	if err != nil {
		panic(err)
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		panic(err)
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		// 显示文件
		fmt.Println(h.Name)
		// 打开文件
		fw, err := os.OpenFile("" + h.Name, os.O_CREATE | os.O_WRONLY, 0644/*os.FileMode(h.Mode)*/)
		if err != nil {
			panic(err)
		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("un tar.gz ok")
}

