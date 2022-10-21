package studentservice

import (
	conn "echo-task/dbconnection"
	vm "echo-task/model"
	"errors"
	"log"
)

func DeleteStudentDetails(roll_no string) (error, *vm.StudentDetails) {
	req := &vm.StudentDetails{}
	db := conn.DB.Table("student_personal_info").Where("roll_no=$1", roll_no).Delete(&req)
	if db.Error != nil {
		log.Println("func Name : RepoDeleteStudentDetails error=", db.Error)
		return db.Error, nil
	} else if db.RowsAffected < 1 {
		return errors.New("not exist"), req
	}
	return nil, req
}

func GetAllStudentDetails() (error, []vm.StudentDetails) {
	var req []vm.StudentDetails
	err := conn.DB.Table("student_personal_info").Find(&req).Error
	if err != nil {
		return err, nil
	}
	return err, req
}

//Get student details
func GetStudentDetails(roll_no, name string) (error, *vm.StudentDetails) {
	req := &vm.StudentDetails{}
	err := conn.DB.Table("student_personal_info").Where("roll_no=? AND name=?", roll_no, name).Find(&req).Error
	if err != nil {
		return errors.New("STUDENT IS NOT AVAILABLE"), nil
	}
	return nil, req
}

//Registraion SignIn
func Registration(newStudentManagement vm.StudentDetails) error {
	if err := conn.DB.Table("student_personal_info").Create(&newStudentManagement).Error; err != nil {
		log.Println(err)
		return errors.New("error in repo method")
	}
	return nil
}
