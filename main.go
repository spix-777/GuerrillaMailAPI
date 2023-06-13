package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	bannerLogger *log.Logger
	infoLogger   *log.Logger
	warnLogger   *log.Logger
)

func init() {
	bannerLogger = log.New(os.Stdout, "", 0)
	infoLogger = log.New(os.Stdout, " [ + ] ", 0)
	warnLogger = log.New(os.Stdout, " [ ! ] ", 0)
}

func banner() {
	banner := `
     ___                      _ _ _                    _ _ 
    / _ \_   _  ___ _ __ _ __(_) | | __ _  /\/\   __ _(_) |
   / /_\/ | | |/ _ \ '__| '__| | | |/ _  |/    \ / _  | | |
  / /_\\| |_| |  __/ |  | |  | | | | (_| / /\/\ \ (_| | | |
  \____/ \__,_|\___|_|  |_|  |_|_|_|\__,_\/    \/\__,_|_|_|
`
	bannerLogger.Println(banner)
}

func main() {
	var version bool
	flag.BoolVar(&version, "v", false, "Print version")
	flag.Parse()
	if version {
		bannerLogger.Println("Guerrilla Mail JSON API")
		bannerLogger.Println("Version:          1.0.0")
		bannerLogger.Println("Made by:       SpiX-777")
		os.Exit(0)
	}
	banner()
	jsonDataGet := getEmailAddress()
	checkEmail(jsonDataGet)
}

func getEmailAddress() map[string]interface{} {
	resp, err := http.Get("http://api.guerrillamail.com/ajax.php?f=get_email_address&ip=127.0.0.1&agent=Mozilla_foo_bar")
	if err != nil {
		warnLogger.Println(err)
		warnLogger.Panicln("http.Get error in getEmailAddress")
	}
	defer resp.Body.Close()

	var jsonDataGet map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonDataGet)
	if err != nil {
		warnLogger.Println(err)
		warnLogger.Panicln("json.NewDecoder error in getEmailAddress")
	}

	infoLogger.Println("eMail:              ", jsonDataGet["email_addr"])
	infoLogger.Println("alias:              ", jsonDataGet["alias"])
	infoLogger.Println("sid_token           ", jsonDataGet["sid_token"])
	bannerLogger.Println()

	return jsonDataGet
}

func checkEmail(jsonDataGet map[string]interface{}) {
	for {
		resp, err := http.Get("http://api.guerrillamail.com/ajax.php?f=check_email&sid_token=" + fmt.Sprintf("%v", jsonDataGet["sid_token"]) + "&seq=0")
		if err != nil {
			warnLogger.Println(err)
			warnLogger.Panicln("http.Get error in checkEmail")
		}
		defer resp.Body.Close()

		var jsonData map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&jsonData)
		if err != nil {
			warnLogger.Println(err)
			warnLogger.Panicln("json.NewDecoder error in checkEmail")
		}

		emails, ok := jsonData["list"].([]interface{})
		if !ok {
			warnLogger.Println("JSON error in checkEmail")
		}

		if len(emails) > 0 {
			firstEmail := emails[0].(map[string]interface{})
			if firstEmail["mail_subject"] == "Welcome to Guerrilla Mail" {
				continue
			}

			bannerLogger.Println("from: ", firstEmail["mail_from"])
			bannerLogger.Println("subject: ", firstEmail["mail_subject"])
			bannerLogger.Println("")
			bannerLogger.Println(firstEmail["mail_excerpt"])
			bannerLogger.Println("")
			bannerLogger.Println("--------------------------------------------------")
			bannerLogger.Println("")
		}

		time.Sleep(10 * time.Second)
	}
}
