// +build windows

package autorunInstall

import (
	"golang.org/x/sys/windows/registry"
	"fmt"
)

func Install(path string) {
	fmt.Println("Windows")
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.QUERY_VALUE|registry.SET_VALUE)

	if err != nil {
		fmt.Println(err)
	}

	defer k.Close()

	k.SetStringValue("COMS", path)
}