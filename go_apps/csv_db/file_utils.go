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
	for {
		select {
		case v := <-c:
			parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", v.CreatedAt)
			vv := []interface{}{parsedCreatedAt, v.Action, v.UserName, v.FamilyName, v.GivenName, v.SchoolClassName, v.SchoolName, v.RemoteAddress, v.ContentName}
			csvRows = append(csvRows, vv)
		case <-done:
			return csvRows, nil
		}
	}
}

// func readUsersCsv(fileName string) ([][]interface{}, error) {
// 	if ext := filepath.Ext(fileName); ext != ".csv" {
// 		return csvRows, fmt.Errorf("%s is not csv file", fileName)
// 	}
// 	f, err := os.Open(filepath.FromSlash(fileName))
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return csvRows, fmt.Errorf("not exist file: %w", err)
// 		}
// 		return csvRows, fmt.Errorf("open csv file error: %w", err)
// 	}
// 	defer f.Close()

// 	c := make(chan Users)

// 	done := make(chan bool)
// 	go func() {
// 		if err := gocsv.UnmarshalToChan(f, c); err != nil {
// 			log.Fatalf("unmarshal csv file error: %w", err)
// 		}
// 		done <- true
// 	}()
// 	for {
// 		select {
// 		case v := <-c:
// 			parsedCreatedAt, _ := time.Parse("2006-01-02 15:04:05", v.CreatedAt)
// 			vv := []interface{}{parsedCreatedAt, v.Action, v.UserName, v.FamilyName, v.GivenName, v.SchoolClassName, v.SchoolName, v.RemoteAddress, v.ContentName}
// 			csvRows = append(csvRows, vv)
// 		case <-done:
// 			return csvRows, nil
// 		}
// 	}
// }

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
