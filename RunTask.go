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

	MessageBoxPlain("Efectis LaunchBadgeuse", "Bonjour !\nVous allez être rediriger vers la page de la badgeuse.\nPasser une bonne journée !")
	exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()

	// fmt.Println("Start")

	for {

		today := time.Now()
		yyyy, mm, dd := today.Date()
		tomorrow := time.Date(yyyy, mm, dd+1, 7, 50, 0, 0, today.Location())
		// tomorrow := time.Now().Local().Add(time.Minute * 2)
		// fmt.Println("[Run]")
		// fmt.Println("[Tomorrow] ", tomorrow)

		for time.Now().Before(tomorrow) {
			time.Sleep(10 * time.Second)
			// fmt.Println("[If] true")
			// fmt.Println("[Timenow] ", time.Now())
			// fmt.Println("[Tomorrow] ", tomorrow)
		}

		if int(today.Weekday()) != 0 && int(today.Weekday()) != 6 {
			MessageBoxPlain("Efectis LaunchBadgeuse", "Bonjour !\nVous allez être rediriger vers la page de la badgeuse.\nPasser une bonne journée !")
			err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
