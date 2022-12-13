package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/halllllll/golog"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func initFolder(folder string, mode fs.FileMode) error {
	_, err := os.Stat(folder)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(folder, mode); err != nil {
				return err
			}
		} else {
			if err := os.Chmod(folder, mode); err != nil {
				return err
			}
		}
	}
	return nil
}

var dbEnv *EnvJson

func initDatabaseInfo() {
	envData, err := Env.ReadFile("secret.json")
	if err != nil {
		err = fmt.Errorf("env file read error: %w", err)
		panic(err)
	}
	if err = json.Unmarshal(envData, &dbEnv); err != nil {
		panic(err)
	}
}

func init() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = loc
	if err := initFolder(saveFolderName, 0777); err != nil {
		panic(err)
	}
	if err = initFolder(watchFolderName, 0777); err != nil {
		panic(err)
	}
	golog.LoggingSetting("lget.log")
	initDatabaseInfo()
}
