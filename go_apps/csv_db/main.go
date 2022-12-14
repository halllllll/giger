package main

import (
	"fmt"
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
			golog.InfoLog.Println("event:", event)
			if event.Has(fsnotify.Create) {
				golog.InfoLog.Println("modified file:", event.Name)
				// ファイル名で分ける
				_, fileName := filepath.Split(event.Name)
				kindOfFile := strings.Split(fileName, "_")[0]
				switch csvType(kindOfFile) {
				case ACTIONLOG:
					golog.InfoLog.Println("start process - action log save")
					rows, err := readActionLogCsv(event.Name)
					if err != nil {
						golog.ErrLog.Println(err)
					}
					// 末尾の行が空なので狂うことがある && なぜかレコードが重複することがある対策
					var result [][]interface{}
					recordMap := make(map[string]bool)
					for _, row := range rows {
						// 先頭3つが空文字でないというのはただのワークアラウンドです
						if len(row) >= 2 && row[0] != "" && row[1] != "" && row[2] == "" {
							continue
						}
						// 時間とIDとIPアドレス（全部複合キー）が重複しないように選ぶ
						x := fmt.Sprintf("%s%s%s", row[0], row[2], row[7])
						if _, ok := recordMap[x]; ok {
							continue
						}
						recordMap[x] = true
						result = append(result, row)
					}
					if err := flowActionCsvToDB(result); err != nil {
						golog.ErrLog.Println(err)
					}
					golog.InfoLog.Println("action log save - process done")

				case USERS:
					golog.InfoLog.Println("start process - user log save")
					rows, err := readUsersCsv(event.Name)
					if err != nil {
						golog.ErrLog.Println(err)
					}
					// 末尾の行が空なので狂うことがある
					var result [][]interface{}
					for _, row := range rows {
						if len(row) != 0 && row[0] != "" {
							result = append(result, row)
						}
					}
					if err := flowUsersCsvToDB(result); err != nil {
						golog.ErrLog.Println(err)
					}
					golog.InfoLog.Println("user log save - process done")

				default:
					golog.InfoLog.Printf("UNKNOWN KIND OF FILE NAME: %s\n", kindOfFile)
				}
				_, fileName = filepath.Split(event.Name)

				outFilePath := filepath.Join(saveFolderName, fileName)
				fmt.Printf("save path: %s\n", outFilePath)
				if err := moveFile(event.Name, outFilePath); err != nil {
					golog.ErrLog.Println(err)
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
		golog.ErrLog.Println(err)
		panic(err)
	}
}
