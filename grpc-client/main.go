package main

import (
	"fmt"
	"time"

	"github.com/marvin5064/grpc-gateway/protobuf/university"
	context "golang.org/x/net/context"
)

type client struct {
	universityMgrCln university.UniversityManagerClient
}

func main() {
	fmt.Println("You are executing Client at", time.Now())
	cln := &client{}
	cln.setupUniversityMgrClient()
	reply1, err := cln.universityMgrCln.GetStudentInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "1"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply student:", reply1)
	}
	reply2, err := cln.universityMgrCln.GetStudentInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "2"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply student:", reply2)
	}
	reply3, err := cln.universityMgrCln.GetStudentInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "3"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply student:", reply3)
	}
	reply4, err := cln.universityMgrCln.GetProfessorInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "1"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply professor:", reply4)
	}
	reply5, err := cln.universityMgrCln.GetProfessorInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "2"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply professor:", reply5)
	}
	reply6, err := cln.universityMgrCln.GetCourseInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "1"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply course:", reply6)
	}
	reply7, err := cln.universityMgrCln.GetCourseInfo(context.TODO(), &university.GetEntityInfoRequest{Id: "2"})
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("Reply course:", reply7)
	}
}
