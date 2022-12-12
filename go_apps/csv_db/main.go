package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/halllllll/golog"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	saveFolderName  string = "save"
	watchFolderName string = "sharedCsvs"
)

type csvType string

const (
	ACTIONLOG csvType = "useractionlog"
	USERS     csvType = "users"
)

func watch(targetPath string) error {
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

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	err = watcher.Add(targetPath)
	if err != nil {
		return err
	}
	// Start listening for events.
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				golog.ErrLog.Println("event error")
				continue
			}
			log.Println("event:", event)
			if event.Has(fsnotify.Create) {
				golog.InfoLog.Println("modified file:", event.Name)
				// ファイル名で分ける
				_, fileName := filepath.Split(event.Name)
				kindOfFile := strings.Split(fileName, "_")[0]
				switch csvType(kindOfFile) {
				case ACTIONLOG:
					fmt.Printf("get action log csv, ここでDBに保存\n")

					rows, err := readActionLogCsv(event.Name)
					if err != nil {
						golog.ErrLog.Println(err)
					}
					if err := pourActionLog(rows, *txn, *ctx, string(LGATE_ACTIONLOG_TABLE)); err != nil {
						golog.ErrLog.Println(err)
					}
					// なぜか上記が失敗するのでここで直接やってみる
					// golog.InfoLog.Printf("here is ActionLog Method!!!!!! table name: %s\n", string(LGATE_ACTIONLOG_TABLE))
					// _, err = txn.CopyFrom(ctx, pgx.Identifier{string(LGATE_ACTIONLOG_TABLE)}, []string{"created_at", "lgate_action", "user_name", "family_name", "given_name", "school_class_name", "school_name", "remote_address", "content_name"}, pgx.CopyFromRows(rows))

					if err != nil {
						err := fmt.Errorf("actin log csv pouring db error: %w", err)
						return err
					}
					if err := _txn.Commit(*ctx); err != nil {
						err := fmt.Errorf("transaction error: %w", err)
						return err
					}
					return nil

				case USERS:
					fmt.Printf("get user csv, DBに保存\n")
				default:
					golog.InfoLog.Printf("UNKNOWN KIND OF FILE NAME: %s\n", kindOfFile)
				}

				outFilePath := filepath.Join(saveFolderName, event.Name)
				fmt.Printf("save path: %s\n", outFilePath)
				if err := moveFile(event.Name, outFilePath); err != nil {
					log.Println("error: ", err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				golog.ErrLog.Println("error:", err)
				return err
			}
		}
	}

	return nil
}

func main() {
	fmt.Println("監視開始")

	if err := watch(watchFolderName); err != nil {
		panic(err)
	}
}
