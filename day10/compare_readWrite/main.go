package main

import (
	"os"
	"fmt"
	"io"
	"time"
	"bufio"
	"io/ioutil"
)

func main(){
	fileName := "./xianni.txt"

	readWrite0(fileName)
	readWrite1(fileName)
	readWrite2(fileName)
}




/*func read0(fileName string){
	var buf [4096]byte
	var t1 = time.Now().UnixNano()
	f,err := os.Open(fileName)
	defer f.Close()
	if err != nil{
		fmt.Println("Open file failed,err:",err)
	}
	for{
		_,err := f.Read(buf[:])
		if err == io.EOF{
			fmt.Println("Read file finished.")
			break
		}
		if err != nil {
			fmt.Println("Read file failed,err:",err)
			return
		}

	}
	var t2 = time.Now().UnixNano()
	fmt.Printf("time cost:%dns\n",t2-t1)
}

func read1(fileName string){
	var t1 = time.Now().UnixNano()
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil{
		fmt.Println("Open file failed,err:",err)
	}
	var buf = bufio.NewReaderSize(f,4096)
	for{
		_,err = buf.ReadBytes('\n')
		if err == io.EOF{
			fmt.Println("Read file finished.")
			break
		}
		if err != nil {
			fmt.Println("Read file failed,err:",err)
			return
		}
	}
	var t2 = time.Now().UnixNano()
	fmt.Printf("time cost:%dns\n",t2-t1)
}

func read2(fileName string){
	var t1 = time.Now().UnixNano()
	_,err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println("Open file failed,err:",err)
	}
	var t2 = time.Now().UnixNano()
	fmt.Printf("time cost:%dns\n",t2-t1)
}*/


func readWrite0(fileName string){
	var buf [4096]byte
	var t1 = time.Now().UnixNano()
	f,err := os.Open(fileName)
	defer f.Close()
	if err != nil{
		fmt.Println("Open file failed,err:",err)
	}
	cp,err := os.OpenFile("cp.txt",os.O_RDWR | os.O_APPEND |os.O_CREATE,0666)
	defer cp.Close()
	if err != nil{
		fmt.Println("Create file failed,err:",err)
	}
	for{
		n,err := f.Read(buf[:])
		if err == io.EOF{
			fmt.Println("Read file finished.")
			break
		}
		if err != nil {
			fmt.Println("Read file failed,err:",err)
			return
		}
		_,err = cp.Write(buf[:n])
		if err != nil {
			fmt.Println("Write file failed,err:",err)
			return
		}
	}
	fmt.Println("write file finished.")
	var t2 = time.Now().UnixNano()
	fmt.Printf("readwrite0 cost time:%dns\n",t2-t1)
}


func readWrite1(fileName string){
	var t1 = time.Now().UnixNano()
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil{
		fmt.Println("Open file failed,err:",err)
	}
	cp,err := os.OpenFile("cp1.txt",os.O_RDWR | os.O_APPEND |os.O_CREATE,0666)
	defer cp.Close()
	if err != nil{
		fmt.Println("Create file failed,err:",err)
	}
	var buf = bufio.NewReaderSize(f,4096)
	var bufWriter = bufio.NewWriterSize(cp,4096)
	for{
		line,err := buf.ReadBytes('\n')
		if err == io.EOF{
			fmt.Println("Read file finished.")
			break
		}
		if err != nil {
			fmt.Println("Read file failed,err:",err)
			return
		}
		_,err = bufWriter.Write(line)
		if err != nil {
			fmt.Println("Write file failed,err:",err)
			return
		}

	}
	fmt.Println("write file finished.")
	var t2 = time.Now().UnixNano()
	fmt.Printf("readwrite1 cost time:%dns\n",t2-t1)
}


func readWrite2(fileName string){
	var t1 = time.Now().UnixNano()
	cp,err := os.OpenFile("cp2.txt",os.O_RDWR | os.O_APPEND |os.O_CREATE,0666)
	defer cp.Close()
	if err != nil{
		fmt.Println("Create file failed,err:",err)
	}
	content,err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println("Open file failed,err:",err)
	}
	_,err = cp.Write(content)
	if err != nil{
		fmt.Println("write file failed,err:",err)
		return
	}
	fmt.Println("write file finished.")
	var t2 = time.Now().UnixNano()
	fmt.Printf("readwrite2 cost time:%dns\n",t2-t1)

}