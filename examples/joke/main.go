package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/enrico5b1b4/tbwrap"
)

func main() {
	telegramBotToken := mustGetEnv("TELEGRAM_BOT_TOKEN")
	allowedUsers := parseUsersAndGroups(os.Getenv("ALLOWED_USERS"))
	allowedGroups := parseUsersAndGroups(os.Getenv("ALLOWED_GROUPS"))

	botConfig := tbwrap.Config{
		Token:         telegramBotToken,
		AllowedUsers:  allowedUsers,
		AllowedGroups: allowedGroups,
	}
	telegramBot, err := tbwrap.NewBot(botConfig)
	if err != nil {
		log.Fatal(err)
	}

	telegramBot.Add("/joke", HandleJoke())
	telegramBot.Start()
}

func mustGetEnv(name string) string {
	value := os.Getenv(name)
	if value == "" {
		log.Fatalln(fmt.Sprintf("%s must be set", name))
	}

	return value
}

func parseUsersAndGroups(list string) []int {
	if len(list) == 0 {
		return nil
	}

	sepList := strings.Split(list, ",")
	intList := make([]int, len(sepList))
	var err error

	for i := range sepList {
		intList[i], err = strconv.Atoi(strings.TrimSpace(sepList[i]))
		if err != nil {
			log.Fatalln(err)
		}
	}

	return intList
}
