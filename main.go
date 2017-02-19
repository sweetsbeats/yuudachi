package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Variables used for command line parameters
var (
	botID   string
	botName string
)

const userAgent = "Yuudachi/0.1"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	twitterConsumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	twitterConsumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	twitterAccessToken := flags.String("access-token", "", "Twitter Access Token")
	twitterAccessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	discordBotToken := flags.String("token", "", "Discord Bot Token")
	bibleAccessToken := flags.String("bible", "", "Bible search token")
	printVersion := flags.Bool("v", false, "Display current version")
	flags.Parse(os.Args[1:])

	//TODO(Sjon): Give this a proper place.
	if *printVersion {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	if *twitterConsumerKey == "" || *twitterConsumerSecret == "" || *twitterAccessToken == "" || *twitterAccessSecret == "" || *discordBotToken == "" || *bibleAccessToken == "" {
		log.Println(*twitterConsumerKey, *twitterConsumerSecret, *twitterAccessToken, *twitterAccessSecret, *discordBotToken, *bibleAccessToken)
		log.Fatal("Consumer key/secret and Access token/secret required")
	}
	log.Println("Keys gotten")
	bibleToken = *bibleAccessToken
	config := oauth1.NewConfig(*twitterConsumerKey, *twitterConsumerSecret)
	token := oauth1.NewToken(*twitterAccessToken, *twitterAccessSecret)
	log.Println("Twitter tokens done")
	// OAuth1 http.Client will automatically authorize Requests
	rawTwitterClient = config.Client(oauth1.NoContext, token)

	// Twitter client
	twitterClient = twitter.NewClient(rawTwitterClient)
	log.Println("Twitter set up")
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + *discordBotToken)
	if err != nil {
		log.Fatalln("error creating Discord session,", err)
	}
	log.Println("Discord session created")

	// Get the account information.
	u, err := dg.User("@me")
	if err != nil {
		log.Fatalln("error obtaining account details,", err)
	}
	log.Println("Got bot details")

	botID = u.ID
	botName = u.Username

	//here we add the functions
	dg.AddHandler(personality)
	dg.AddHandler(command)
	dg.AddHandler(tatsumaki)
	log.Println("Handlers added")
	err = dg.Open()
	if err != nil {
		log.Fatalln("error opening connection,", err)
	}
	log.Println("Discord opened")
	fmt.Println("Succesfully initialized")
	<-make(chan struct{})
}
