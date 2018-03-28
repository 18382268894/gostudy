package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

//type sortSlice []int
/*  实现sort.Interface接口，可以用于各种排序。但是切片可以不用这样，
sort.Ints（）内部实现时已经强制转换为这个接口类型
func (s sortSlice) Len() int {
	return len(s)
}
func (s sortSlice) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}
	return false
}
func (s sortSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
*/

//利用切片做密码的模板
var (
	nums  []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	chars []byte = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	mix   []byte = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
)

func main() {
	/*var str = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		str = append(str, fmt.Sprintf("%v", i))
	}
	fmt.Printf("%v\tlen:%d\tcap:%d\n", str, len(str), cap(str))
	*/

	//利用sort包对切片进行排序
	//var mySlice = sortSlice{5, 2, 9, 11, 7, 3, 20, 14, 6}
	//sort.Ints(mySlice) //升序进行排列
	//fmt.Println(mySlice)

	//利用flag包对用户指令进行解析，并生成用户制定类型和长度的密码
	var cmdLength = flag.Int("l", 6, "length of password")                            //密码长度
	var cmdType = flag.String("t", "num", "type for passwod,eg:num,char,mix,advance") //密码类型
	flag.Parse()
	fmt.Printf("密码长度：%d\t密码类型：%v\n", *cmdLength, *cmdType)
	var password []byte = make([]byte, *cmdLength) //初始化指定长度的密码切片
	rand.Seed(time.Now().UnixNano())               //利用当前的unix时间戳来初始化rand种子
	switch *cmdType {
	case "num":
		for i := 0; i < *cmdLength; i++ {
			password[i] = nums[rand.Intn(len(nums))]
		}
	case "mix":
		for i := 0; i < *cmdLength; i++ {
			password[i] = mix[rand.Intn(len(mix))]
		}
	case "char":
		for i := 0; i < *cmdLength; i++ {
			password[i] = chars[rand.Intn(len(chars))]
		}
	}
	passwordStr := string(password)
	fmt.Println("你的密码为：" + passwordStr)

}
