package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	fmt.Println("vim-go")
	w, _ := fsnotify.NewWatcher()
	w.Add(".")
	time.Sleep(time.Second)
}
