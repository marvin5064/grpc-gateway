package main

import (
	"github.com/marvin5064/grpc-gateway/protobuf/university"
)

var student1 = &university.Student{
	Id:     "1",
	Age:    20,
	Gender: university.Gender_MALE,
	RegistedCourses: []*university.Course{
		course1,
	},
	Name: "Knight",
}

var student2 = &university.Student{
	Id:     "2",
	Age:    22,
	Gender: university.Gender_MALE,
	RegistedCourses: []*university.Course{
		course2,
	},
	Name: "Farmer",
}

var student3 = &university.Student{
	Id:     "3",
	Age:    19,
	Gender: university.Gender_FEMALE,
	RegistedCourses: []*university.Course{
		course1,
		course2,
	},
	Name: "Princess",
}

var course1 = &university.Course{
	Id:         "1",
	Name:       "COMP",
	Professors: []*university.Professor{professor1},
}

var course2 = &university.Course{
	Id:         "2",
	Name:       "MATH",
	Professors: []*university.Professor{professor1, professor2},
}

var professor1 = &university.Professor{
	Id:     "1",
	Gender: university.Gender_MALE,
	Name:   "Master",
	Age:    39,
}

var professor2 = &university.Professor{
	Id:     "2",
	Gender: university.Gender_MALE,
	Name:   "God",
	Age:    43,
}
