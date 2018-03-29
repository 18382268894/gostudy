package main

import (
	"fmt"
	"strings"
)

type studentInfo struct {
	name  string
	age   uint8
	score uint8
}

func main() {

	//map的初始化
	var m1 = map[string]int{
		"std01": 1,
		"std02": 2,
	}
	fmt.Println(m1)
	var m2 = make(map[string]int, 16)
	m2["std01"] = 1
	m2["std02"] = 2
	fmt.Println(m2)

	//查询map中是否存在该key-value
	val, ok := m2["std03"]
	if !ok {
		m2["std03"] = 5
		val = m2["std03"]
	}
	fmt.Println(val)
	fmt.Println(m2)

	//对map进行排序(依据key)
	/*var m3 = make(map[string]int, 100)
	for i := 0; i < 100; i++ { //填充map
		key := fmt.Sprintf("std%02d", i)
		m3[key] = i
	}
	var slice = make([]string, 0, 100)
	for key, _ := range m3 {
		slice = append(slice, key)
	}
	sort.Strings(slice)
	delete(m3, "std99") //删除不存在的key不会报错
	for _, v := range slice {
		fmt.Printf("%s:%d\n", v, m3[v])
	}*/

	str := "What can i do for you? You shuold not hidden the trueth! I can not belive that!"
	count(str)
	students := holdInfo()
	var studentId int64 = 201402040217
	student := searchInfo(studentId, students)
	fmt.Printf("name:%s\tage:%d\tscore:%d\n", student.name, student.age, student.score)

}

//统计句子里出现的单词次数
func count(str string) {
	var m = make(map[string]int, 1024)
	strSlice := strings.Split(str, " ")
	for _, val := range strSlice {
		m[val]++
	}
	fmt.Println(m)

}

//储存学生信息，key为学号
func holdInfo() map[int64]studentInfo {
	var student = make(map[int64]studentInfo, 3)
	student[201402040216] = studentInfo{
		name:  "cxz",
		age:   23,
		score: 59,
	}
	student[201402040215] = studentInfo{
		name:  "bhj",
		age:   23,
		score: 58,
	}
	student[201402040214] = studentInfo{
		name:  "wqy",
		age:   23,
		score: 57,
	}
	return student

}

//通过学号，在学生map里查询该学生信息
func searchInfo(id int64, students map[int64]studentInfo) studentInfo {
	_, ok := students[id]
	if !ok {
		fmt.Println("没有该学生信息")
		return studentInfo{}
	}
	return students[id]
}
