package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	saveFolderName  string = "save"
	watchFolderName string = "sharedCsvs"
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
}

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
				log.Println("event error")
				continue
			}
			log.Println("event:", event)
			if event.Has(fsnotify.Create) {
				log.Println("modified file:", event.Name)

				outFileName := fmt.Sprintf("output_%s.txt", time.Now().Format("2006_01_02_150405_.000000000"))
				outFilePath := filepath.Join(saveFolderName, outFileName)
				fmt.Printf("save path: %s\n", outFilePath)
				if err := moveFile(event.Name, outFilePath); err != nil {
					log.Println("error: ", err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
			log.Println("error:", err)
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
