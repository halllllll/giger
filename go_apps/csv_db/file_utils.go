package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/halllllll/golog"
)

func readActionLogCsv(fileName string) ([][]interface{}, error) {
	if ext := filepath.Ext(fileName); ext != ".csv" {
		return nil, fmt.Errorf("%s is not csv file", fileName)
	}
	f, err := os.Open(filepath.FromSlash(fileName))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("not exist file: %w", err)
		}
		return nil, fmt.Errorf("open csv file error: %w", err)
	}
	defer f.Close()

	c := make(chan ActionLog)

	done := make(chan bool)
	go func() {
		if err := gocsv.UnmarshalToChan(f, c); err != nil {
			golog.ErrLog.Println("unmarshal csv file error: %w", err)
		}
		done <- true
	}()
	var csvRows [][]interface{}
	fmt.Println("let's goooo")
	for {
		select {
		case v := <-c:
			parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", v.CreatedAt)
			vv := []interface{}{parsedCreatedAt, v.LgateAction, v.UserName, v.FamilyName, v.GivenName, v.SchoolClassName, v.SchoolName, v.RemoteAddress, v.ContentName}
			csvRows = append(csvRows, vv)
		case <-done:
			return csvRows, nil
		}
	}
}

func readUsersCsv(fileName string) ([][]interface{}, error) {
	if ext := filepath.Ext(fileName); ext != ".csv" {
		return nil, fmt.Errorf("%s is not csv file", fileName)
	}
	f, err := os.Open(filepath.FromSlash(fileName))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("not exist file: %w", err)
		}
		return nil, fmt.Errorf("open csv file error: %w", err)
	}
	defer f.Close()

	c := make(chan Users)

	done := make(chan bool)
	go func() {
		if err := gocsv.UnmarshalToChan(f, c); err != nil {
			golog.ErrLog.Println("unmarshal csv file error: %w", err)
		}
		done <- true
	}()
	var csvRows [][]interface{}
	for {
		select {
		case v := <-c:
			vv := []interface{}{v.UserName, v.Password, v.EnabledUser, v.IsLocal, v.SchoolCode, v.SchoolName, v.FamilyName, v.GivenName, v.FamilyKanaName, v.GivenKanaName, v.RenewName, v.RenewPassword, v.RenewClass, v.TermName1, v.ClassName1, v.ClassRole1, v.ClassNumber1, v.TermName2, v.ClassName2, v.ClassRole2, v.ClassNumber2, v.TermName3, v.ClassName3, v.ClassRole3, v.ClassNumber3, v.TermName4, v.ClassName4, v.ClassRole4, v.ClassNumber4, v.TermName5, v.ClassName5, v.ClassRole5, v.ClassNumber5, v.TermName6, v.ClassName6, v.ClassRole6, v.ClassNumber6, v.TermName7, v.ClassName7, v.ClassRole7, v.ClassNumber7, v.TermName8, v.ClassName8, v.ClassRole8, v.ClassNumber8, v.TermName9, v.ClassName9, v.ClassRole9, v.ClassNumber9, v.TermName10, v.ClassName10, v.ClassRole10, v.ClassNumber10, v.TermName11, v.ClassName11, v.ClassRole11, v.ClassNumber11, v.TermName12, v.ClassName12, v.ClassRole12, v.ClassNumber12, v.TermName13, v.ClassName13, v.ClassRole13, v.ClassNumber13, v.TermName14, v.ClassName14, v.ClassRole14, v.ClassNumber14, v.TermName15, v.ClassName15, v.ClassRole15, v.ClassNumber15, v.TermName16, v.ClassName16, v.ClassRole16, v.ClassNumber16, v.TermName17, v.ClassName17, v.ClassRole17, v.ClassNumber17, v.TermName18, v.ClassName18, v.ClassRole18, v.ClassNumber18, v.TermName19, v.ClassName19, v.ClassRole19, v.ClassNumber19, v.TermName20, v.ClassName20, v.ClassRole20, v.ClassNumber20, v.TermName21, v.ClassName21, v.ClassRole21, v.ClassNumber21, v.TermName22, v.ClassName22, v.ClassRole22, v.ClassNumber22}
			csvRows = append(csvRows, vv)
		case <-done:
			return csvRows, nil
		}
	}
}

func moveFile(in string, out string) error {
	src, err := os.OpenFile(in, os.O_RDONLY, 0644)
	if err != nil {
		err = fmt.Errorf("file '%s' open error: %w", in, err)
	}
	dst, err := os.OpenFile(out, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		err = fmt.Errorf("file '%s' open error: %w", in, err)
	}
	_, err = io.Copy(dst, src)
	if err != nil {
		err = fmt.Errorf("file copy error: %w", err)
		return err
	}
	return os.Remove(in)
}
