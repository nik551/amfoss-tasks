package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	var User string
	fmt.Println("Enter the name of the user?")
	fmt.Scanln(&User)
	tn := flag.String("twitter", User, "it contains the name of twitter")
	flag.Parse()
	config := oauth1.NewConfig("EMHQ8tNJ3MM29cXD2C6MCwYCZ", "NFJqZxIeQ8NIjAj9MbiDkdbX5uZCFDNNivwVMz8hUF7VYE4W8o")
	token := oauth1.NewToken("1171096328459444224-geNoko5XEhB6gwn6BBaBaZUsWgbG0b", "wwMDcjEY3QDhs3NGdsSDpqTHzYD6Tnnwsu40RPera56Px")
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	g, err := os.Create("Followers.txt")
	params := &twitter.FollowerListParams{ScreenName: *tn}

	followers, resp, err := client.Followers.List(params)
	var numb int = 0
	fmt.Println(resp, err)
	g.WriteString("Followers - " + *tn)
	for _, follower := range followers.Users {
		numb++
		g.WriteString("\n" + follower.ScreenName)
	}
	g.WriteString("\n")
	g.WriteString(fmt.Sprintf("\n", numb))
	g.Close()

}
