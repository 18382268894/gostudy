package main

import (
	"os"
	"compress/gzip"
	"archive/tar"
	"time"
	"io"
	"fmt"
)

func main(){
	var src = "./src/"
	var dst = "test.tar.gz"
	err := compress(src,dst)
	if err != nil{
		fmt.Println(err)
	}
	err = decompress("test.tar.gz","./")
	if err != nil{
		fmt.Println(err)
	}
}

func compress(src string,dst string)error{
	dstFile,err := os.Create(dst)  //创建归档压缩目的文件
	if err != nil{
		fmt.Println(0)
		return err
	}
	defer dstFile.Close()
	gz := gzip.NewWriter(dstFile)
	if err != nil{
		fmt.Println(1)
		return err
	}
	defer gz.Close()
	tg := tar.NewWriter(gz)
	defer tg.Close()

	srcFile,err := os.Open(src)  //读取目标文件夹或文件
	if err != nil {
		fmt.Println(2)
		return err
	}
	defer srcFile.Close()
	srcStat,err := srcFile.Stat()
	if err != nil {
		fmt.Println(3)
		return err
	}
	if srcStat.IsDir() {
		fileInfos,err := srcFile.Readdir(0)
		if err != nil {
			fmt.Println(4)
			return err
		}
		for _,fileInfo := range fileInfos{
			if !fileInfo.IsDir(){
				header := new(tar.Header)
				header.Name = fileInfo.Name()
				header.Size = fileInfo.Size()  //必须设置不然会出现write too long错误
				header.ModTime = fileInfo.ModTime()
				header.AccessTime = time.Now()
				header.Mode = int64(fileInfo.Mode())
				err := tg.WriteHeader(header)
				if err != nil {
					fmt.Println(5)
					return err
				}
				fr,err := os.Open(srcFile.Name()+fileInfo.Name())
				fmt.Println(fr.Name())
				if err != nil {
					fmt.Println(6)
					return err
				}
				_,err = io.Copy(tg,fr)
				if err != nil {
					fmt.Println(7)
					return err
				}

			}
		}
	}else{
		header := new(tar.Header)
		header.Name = srcStat.Name()
		header.Size = srcStat.Size()
		header.ModTime = srcStat.ModTime()
		header.AccessTime = time.Now()
		header.Mode = int64(srcStat.Mode())
		err := tg.WriteHeader(header)
		if err != nil {
			return err
		}
		_,err = io.Copy(tg,srcFile)
		if err != nil {
			return nil
		}
		fmt.Println(srcStat.Name())
	}
	fmt.Println("tar.gz is ok!")
	return nil
}


func decompress(src string,dst string)error{
	srcFile,err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gz,err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gz.Close()
	tg := tar.NewReader(gz)
	for{
		h,err := tg.Next()  //便利访问要读取的文件的tar.Header
		if err == io.EOF{
			break
		}
		if err != nil {
			return err
		}
		dstFile,err := os.OpenFile(dst+"/"+h.Name,os.O_CREATE | os.O_WRONLY,os.FileMode(h.Mode))  //os.FileMode(h.Mode)会涉及到位截断
		if err != nil{
			return err
		}
		_,err = io.Copy(dstFile,tg)
		if err != nil {
			return err
		}
		fmt.Println(dstFile.Name())

	}
	fmt.Println("Decompress secussed")
	return nil
}