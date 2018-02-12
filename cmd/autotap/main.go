package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/velour/untappd"
)

var token = flag.String("token", "", "Untappd token")
var users = flag.String("users", "", "Comma separated users to tap")

func main() {
	flag.Parse()
	if *token == "" {
		fmt.Println("Error: The -token flag is required.")
		os.Exit(2)
	}
	if *users == "" {
		fmt.Println("Error: The -users flag is required.")
		os.Exit(2)
	}

	userMap := map[string]bool{}
	for _, u := range strings.Split(*users, ",") {
		userMap[strings.ToLower(strings.TrimSpace(u))] = true
	}

	u := untappd.New(*token)
	feed, err := u.PullFeed()
	if err != nil {
		log.Fatal(err)
	}
	for _, checkin := range feed {
		log.Printf("Checkin by %s", checkin.User.UserName)
		uname := checkin.User.UserName
		if userMap[uname] {
			log.Println(checkin.CheckinID)
			toast, err := u.Toast(checkin.CheckinID)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(toast)
		}
	}
}
