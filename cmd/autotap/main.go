package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	quaff "github.com/chrissexton/quaff"
)

var (
	token  = flag.String("token", "", "Untappd token")
	users  = flag.String("users", "", "Comma separated users to tap")
	whoami = flag.String("whoami", "", "Username of toaster")
	limit  = flag.Int("limit", 25, "Number of checkins to examine")
)

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
	if *whoami == "" {
		fmt.Println("Error: The -whoami flag is required.")
		os.Exit(2)
	}

	userMap := map[string]bool{}
	for _, u := range strings.Split(*users, ",") {
		userMap[strings.ToLower(strings.TrimSpace(u))] = true
	}

	u := quaff.New(*token)
	u.Limit = *limit

	for name, _ := range userMap {
		checkins, err := u.PullUserCheckins(name)
		if err != nil {
			log.Fatal(err)
		}

		for _, checkin := range checkins {
			log.Printf("Checkin by %s", checkin.User.UserName)
			log.Println(checkin.CheckinID)

			found := false
			for _, toast := range checkin.Toasts.Items {
				found = found || toast.User.UserName == *whoami
			}

			if !found {
				_, err := u.Toast(checkin.CheckinID)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Toasted %d", checkin.CheckinID)
			}
		}
	}
}
