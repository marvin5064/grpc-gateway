package main

import (
	"fmt"

	context "golang.org/x/net/context"

	"github.com/marvin5064/grpc-gateway/protobuf/university"
)

func (s *server) GetStudentInfo(ctx context.Context, request *university.GetEntityInfoRequest) (*university.Student, error) {
	fmt.Println("Processing GetStudentInfo for request", request)
	var student *university.Student
	switch request.Id {
	case "1":
		student = student1
	case "2":
		student = student2
	case "3":
		student = student3
	default:
		return &university.Student{}, nil
	}
	return student, nil
}

func (s *server) GetCourseInfo(ctx context.Context, request *university.GetEntityInfoRequest) (*university.Course, error) {
	fmt.Println("Processing GetCourseInfo for request", request)
	var course *university.Course
	switch request.Id {
	case "1":
		course = course1
	case "2":
		course = course2
	default:
		return &university.Course{}, nil
	}
	return course, nil
}

func (s *server) GetProfessorInfo(ctx context.Context, request *university.GetEntityInfoRequest) (*university.Professor, error) {
	fmt.Println("Processing GetProfessorInfo for request", request)
	var professor *university.Professor
	switch request.Id {
	case "1":
		professor = professor1
	case "2":
		professor = professor2
	default:
		return &university.Professor{}, nil
	}
	return professor, nil
}
