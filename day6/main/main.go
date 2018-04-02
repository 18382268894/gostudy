package main

import (
	"fmt"

	"day6/students"
)

func main() {
	/*
		name := "cxz"
		var age uint8 = 23
		photoUrl := "www.cxz.com/jpg.index"
		var user = user.NewUser(name, age, photoUrl)
		fmt.Printf("%#v\n", user)
	*/
	var s = students.NewStudentsForm(20)
	menu(&s)
}

func menu(s *students.Students) {
	menustring := `
Please input your choice:
	1.Add student
	2.Modify student info
	3.Change student score
	4.Quit
`

	for {
		var selector int
		fmt.Println(menustring)
		fmt.Scanf("%d", &selector)
		switch selector {
		case 1:
			err := s.AddStudent()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Add the student success")
		case 2:
			s.ModifyStudentInfo()
		}

	}

}
