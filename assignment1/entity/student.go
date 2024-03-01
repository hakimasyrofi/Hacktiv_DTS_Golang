package entity

import "fmt"

type Student struct {
	ID      int
	Name    string
	Address string
	Job     string
	Reason  string
}

var students = []Student{
	{ID: 1, Name: "John Doe", Address: "123 Main St", Job: "Software Engineer", Reason: "Interested in learning new technologies"},
	{ID: 2, Name: "Jane Smith", Address: "456 Elm St", Job: "Data Scientist", Reason: "Passionate about data analysis"},
	{ID: 3, Name: "Michael Johnson", Address: "789 Oak St", Job: "UX Designer", Reason: "Enjoys creating user-friendly interfaces"},
}

func (s *Student) GetStudent(ID int) (Student, error) {
	for _, student := range students {
		if student.ID == ID {
			return student, nil
		}
	}
	return Student{}, fmt.Errorf("Student with ID %d not found", ID)
}
