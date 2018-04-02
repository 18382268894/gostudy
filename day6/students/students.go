package students

import (
	"errors"
	"fmt"
)

type Sex uint8

const (
	MAN   Sex = 1
	WOMAN     = 0
)

type Score struct {
	English float32
	Math    float32
	Chinese float32
}
type Student struct {
	ID   uint64
	Name string
	Sex
	Score
}

type Students []Student

//创建学生列表
func NewStudentsForm(studentNum uint8) Students {
	var students = make([]Student, 0, studentNum)
	return students
}

//添加新学生信息
func (s *Students) AddStudent() error {
	length := len(*s)
	var student Student
	fmt.Println("Please input student's name:\t")
	fmt.Scanf("%s", &student.Name)
	fmt.Println("Please input student's ID:\t")
	fmt.Scanf("%d", &student.ID)
	fmt.Println("Please input student's Sex:\t")
	fmt.Scanf("%d", &student.Sex)
	*s = append(*s, student)
	if len(*s) <= length {
		err := errors.New("Cannot add the student")
		return err
	}
	return nil
}

//修改学生信息，不包括成绩
func (s *Students) ModifyStudentInfo() {
	var selector int = 1
	var student Student
	fmt.Println("Please select the way to lock the student:1 is ID, 2 is name.")
	fmt.Scanf("%d", &selector)
	if selector == 1 {
		fmt.Println("Please input student's ID:\t")
		fmt.Scanf("%d", &student.ID)
		for _, v := range *s {
			if v.ID == student.ID {
				fmt.Println("Please input student's Sex:\t")
				fmt.Scanf("%s", &v.Sex)
			}
		}
	} else if selector == 2 {
		fmt.Println("Please input student's name:\t")
		fmt.Scanf("%d", &student.Name)
		for _, v := range *s {
			if v.Name == student.Name {
				fmt.Println("Please input student's Sex:\t")
				fmt.Scanf("%s", &v.Sex)
			}
		}
	}
	fmt.Println("Modify the student info success")
}
