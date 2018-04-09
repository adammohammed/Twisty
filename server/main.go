package main

import (
	"fmt"
	"github.com/adammohammed/groupmebot"
	"log"
	"net"
	"net/http"
	"time"
)

var queue = make(chan byte)

/*
 Test hook functions
 Each hook should match a certain string, and if it matches
 it should return a string of text
 Hooks will be traversed until match occurs
*/
func direction(msg groupmebot.InboundMessage) string {
	text := msg.Text[0]
	log.Printf("Inserted %c into queue\n", text)
	queue <- text
	return ""
}

func handleConnection(c net.Conn) {
	for {
		select {
			case out, ok := <-queue:
				      if ok {
						x := fmt.Sprintf("%c", out)
						bs := []byte(x)
						c.Write(bs)
				      } else {
					      log.Println("Closed")
				      }
			default:
				      time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	bot, err := groupmebot.NewBotFromJson("bot_cfg.json")
	if err != nil {
		log.Fatal("Could not create bot structure")
	}

	// Make a list of functions
	bot.AddHook("d$", direction)
	bot.AddHook("a$", direction)
	bot.AddHook("w$", direction)

	// Create Server to listen for incoming POST from GroupMe
	log.Printf("Listening on %v...\n", bot.Server)

	botMux := http.NewServeMux()
	botMux.HandleFunc("/", bot.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(bot.Server, botMux))
	}()

	ln, err := net.Listen("tcp", ":9090")
	if err != nil {
		// handle error
	}
	conn, err := ln.Accept()
	if err != nil {
		// handle error
	}
	handleConnection(conn)
}
