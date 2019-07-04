package main

import (
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {
	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	err = watch.Watch("/tmp")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case ev := <-watch.Event:
				{
					if ev.IsCreate() {
						log.Println("创建文件 : ", ev.Name)
					}
					if ev.IsDelete() {
						log.Println("删除文件 : ", ev.Name)
					}
					if ev.IsRename() {
						log.Println("重命名文件 : ", ev.Name)
					}
					if ev.IsModify() {
						log.Println("文件属性更改 : ", ev.Name)
					}
				}
			case err := <-watch.Error:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()
	select {}
}

