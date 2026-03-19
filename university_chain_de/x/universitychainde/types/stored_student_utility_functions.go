package types

import (
	"net/mail"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var course_of_studies map[string]int = map[string]int{
	"cs": 1,
}

func (student_info StudentInfo) GetStudentAddress() (address sdk.AccAddress, err error) {
	address, errAddress := sdk.AccAddressFromBech32(student_info.StudentKey)
	return address, sdkerrors.Wrapf(errAddress, ErrInvalidStudentAddress.Error(), student_info.StudentKey)
}

func (student_info StudentInfo) Validate() (err error) {

	_, err = student_info.GetStudentAddress()
	if err != nil {
		return err
	}
	if student_info.CourseType != "master" && student_info.CourseType != "bachelor" {
		return ErrWrongCourseType
	}
	_, found := course_of_studies[student_info.CourseOfStudy]
	if !found {
		return ErrWrongCourseOfStudy
	}
	return err
}

func (student_info PersonalInfo) Validate() (err error) {

	if student_info.Gender != "male" && student_info.Gender != "female" && student_info.Gender != "non binary" {
		return ErrGenderNotFound
	}
	date, errDate := time.Parse("2006-01-02", student_info.DateOfBirth)
	if errDate != nil {
		return errDate
	}
	today := time.Now()
	if date.After(today.AddDate(-18, 0, 0)) || date.Before(today.AddDate(-100, 0, 0)) {
		return ErrWrongDate
	}
	/*
		student_info.PrimaryNationality
		student_info.CountryOfBirth
		student_info.ProvinceOfBirth
		student_info.TownOfBirth
	*/

	if len(student_info.TaxCode) != 16 {
		return ErrWrongTaxCode
	}

	return err
}

func (student_info ContactInfo) Validate() (err error) {

	_, err = mail.ParseAddress(student_info.Email)
	if err != nil {
		return err
	}
	if len(student_info.MobilePhone) > 13 {
		return ErrWrongMobileNumber
	}

	//student_info.ContactAddress

	return err
}

func (student_info ResidenceInfo) Validate() (err error) {

	if len(student_info.HomePhone) > 13 {
		return ErrWrongMobileNumber
	}
	/*
		student_info.Country
		student_info.Province
		student_info.Town
		student_info.PostCode
		student_info.Address
		student_info.HouseNumber
	*/
	return err
}
