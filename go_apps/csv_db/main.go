package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/halllllll/golog"
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
					if err := flowActionCsvToDB(rows); err != nil {
						golog.ErrLog.Println(err)
					}
					fmt.Println("うまくいっとるやないか")

				case USERS:
					fmt.Printf("get user csv, DBに保存\n")
					fmt.Println("テーブルごと作り直す")

					rows, err := readUsersCsv(event.Name)
					if err != nil {
						golog.ErrLog.Println(err)
					}
					fmt.Printf("一行の長さが知りたいな %d\n", len(rows[0]))
					fmt.Println(rows[0])
					if err := flowUsersCsvToDB(rows); err != nil {
						golog.ErrLog.Println(err)
					}

				default:
					golog.InfoLog.Printf("UNKNOWN KIND OF FILE NAME: %s\n", kindOfFile)
				}
				_, fileName = filepath.Split(event.Name)

				outFilePath := filepath.Join(saveFolderName, fileName)
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
