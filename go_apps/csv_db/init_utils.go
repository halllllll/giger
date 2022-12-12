package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/halllllll/golog"
	"github.com/jackc/pgx/v4"
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
var ctx *context.Context
var txn *pgx.Tx

func initDatabaseInfo() {
	envData, err := Env.ReadFile("secret.json")
	if err != nil {
		err = fmt.Errorf("env file read error: %w", err)
		panic(err)
	}
	if err = json.Unmarshal(envData, &dbEnv); err != nil {
		panic(err)
	}
	_ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbEnv.DbUser, dbEnv.DbPw, dbEnv.Host, dbEnv.DbPort, dbEnv.DbName))
	if err != nil {
		panic(err)
	}

	_txn, err := conn.Begin(_ctx)
	if err != nil {
		panic(err)
	}
	txn = &_txn
	ctx = &_ctx
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
