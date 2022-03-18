package main

import (
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"syscall"
	"time"
	"unsafe"
)

func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))

	return int(ret)
}

func MessageBoxPlain(title, caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK)
}

func main() {
	url := "https://nothingelse.fr/"
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	startup_path := user.HomeDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"

	if exPath == startup_path {
		// fmt.Println("ok")
	} else {
		err := os.Rename(os.Args[0], user.HomeDir+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\"+filepath.Base(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		MessageBoxPlain("Efectis LaunchBadgeuse", "Installation successful !")
	}

	// fmt.Println(os.Args[0])
	// fmt.Println(filepath.Base(os.Args[0]))

	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	for {
		today := time.Now()
		yyyy, mm, dd := today.Date()
		tomorrow := time.Date(yyyy, mm, dd+1, 7, 50, 0, 0, today.Location())
		diff := tomorrow.Sub(today)
		time.Sleep(diff)

		err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
