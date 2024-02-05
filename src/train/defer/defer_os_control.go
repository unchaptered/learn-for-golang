package main

import (
	"fmt"
	"io"
	"os"
)

func defer_os_control() {
	newFile, errorFile := os.Create("learnGo.txt")
	fmt.Println("[defer_os_control]", "os.Create")
	if errorFile != nil {
		fmt.Println("[defer_os_control][errorFile]")
	}

	defer fmt.Println("[defer_os_control]", "before close")
	defer newFile.Close()
	defer fmt.Println("[defer_os_control]", "after close")

	fmt.Println("[defer_os_control]", "io.WriteString")
	if _, errorFile = io.WriteString(newFile, "Learning Go!"); errorFile != nil {
		fmt.Println("[defer_os_control][WriteString]")
		return
	}

	fmt.Println("[defer_os_control]", "newFile.Sync")
	newFile.Sync()
}
