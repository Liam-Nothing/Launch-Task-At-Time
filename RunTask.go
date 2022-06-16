package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
	"time"
	"unsafe"

	"github.com/martinlindhe/inputbox"
)

type JsonData struct {
	Disable    int `json: "disable"`
	Delta_time int `json: "delta_time"`
}

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
		NULL     = 0
		MB_YESNO = 4
	)
	return MessageBox(NULL, caption, title, MB_YESNO)
}

// func RdmMessage

func main() {
	url := "https://nothingelse.fr/"
	delta_time := 8
	disable := 0
	accepted_value_message := 6

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	config_badgeuse_path := user.HomeDir + "\\AppData\\Roaming\\config_badgeuse.json"

	if _, err := os.Stat(config_badgeuse_path); err == nil {
		file, _ := ioutil.ReadFile(config_badgeuse_path)
		data := JsonData{}
		_ = json.Unmarshal([]byte(file), &data)
		delta_time = data.Delta_time
		disable = data.Disable
	} else {

		var_message := MessageBoxPlain("Efectis LaunchBadgeuse", "Souhaitez-vous configurer la badgeuse ?")
		if var_message == accepted_value_message {
			// disable_message := MessageBoxPlain("Efectis LaunchBadgeuse", "Êtes-vous un cadre non pointant ? (Si oui alors le rappel pour la badgeuse se désactivera)")
			disable_message := 0
			if disable_message == accepted_value_message {
				disable = 1
			} else {
				got, ok := inputbox.InputBox("Efectis LaunchBadgeuse", "Dans combien de temps (heures) voulez-vous configurer le debadgeage ?", "8")
				if ok && len(got) > 0 {
					delta_time, _ = strconv.Atoi(got)
				}
			}
		}

		//Wirte file
		f, err := os.Create(config_badgeuse_path)
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString("{\n\t\"disable\": " + strconv.Itoa(disable) + ",\n\t\"delta_time\": " + strconv.Itoa(delta_time) + "\n}")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "bytes written successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if disable == 1 {
		os.Exit(0)
	}

	// Run first message
	today := time.Now()
	if int(today.Weekday()) != 0 && int(today.Weekday()) != 6 {
		var_message := MessageBoxPlain("Efectis LaunchBadgeuse", "Bonjour,\nVous allez être redirigé vers la page de la badgeuse.\nPassez une bonne journée !")
		if var_message == accepted_value_message {
			err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	for {

		today := time.Now()
		yyyy, mm, dd := today.Date()
		tomorrow := time.Date(yyyy, mm, dd+1, 7, 50, 0, 0, today.Location())
		evening := time.Now().Local().Add(time.Duration(delta_time) * time.Hour)
		fmt.Println(today)
		fmt.Println(tomorrow)
		fmt.Println(evening)

		fmt.Println("Waiting for evening...")
		for time.Now().Before(evening) {
			time.Sleep(10 * time.Second)
		}

		if int(today.Weekday()) != 0 && int(today.Weekday()) != 6 {
			var_message := MessageBoxPlain("Efectis LaunchBadgeuse", "Bonjour,\nVous allez être redirigé vers la page de la badgeuse.\nPassez une bonne journée !")
			if var_message == accepted_value_message {
				err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		fmt.Println("Waiting for tomorrow...")
		for time.Now().Before(tomorrow) {
			time.Sleep(10 * time.Second)
		}

		if int(today.Weekday()) != 0 && int(today.Weekday()) != 6 {
			var_message := MessageBoxPlain("Efectis LaunchBadgeuse", "Bonjour,\nVous allez être redirigé vers la page de la badgeuse.\nPassez une bonne journée !")
			if var_message == accepted_value_message {
				err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
