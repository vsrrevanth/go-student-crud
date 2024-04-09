package student

import (
	"errors"
)

var students []Student

func init() {
	students = []Student{
		{ID: "1", Name: "Revanth Varanasi", Department: "Computer Science", Age: 26},
		{ID: "2", Name: "Karthik Smith", Department: "Electrical Engineering", Age: 24},
	}
}

func GetAllStudents() []Student {
	return students
}

func GetStudentById(id string) (*Student, error) {
	for _, s := range students {
		if s.ID == id {
			return &s, nil
		}
	}
	return nil, errors.New("Student with ID not found")
}

func CreateStudent(s Student) error {
	for _, current := range students {
		if s.ID == current.ID {
			return errors.New("Student with ID already exists ")
		}
	}
	students = append(students, s)
	return nil
}

func UpdateStudent(id string, stu Student) error {
	for i, s := range students {
		if s.ID == id {
			stu.ID = id
			students[i] = stu
			return nil
		}
	}
	return errors.New("cannot update the student")
}

func DeleteStudent(id string) error {
	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}
