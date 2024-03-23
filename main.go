package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u, err := url.Parse("ws://43.229.135.146:6611/3/5?DownType=5&jsession=&DevIDNO=23115725861&FILELOC=2&FILESVR=7&FILECHN=0&FILEBEG=9&FILEEND=909&PLAYIFRM=0&PLAYFILE=C:/gStorage/RECORD_FILE/23115725861/2024-03-21/camera-11-1_0-240321-000009-001509-20010100.mp4&PLAYBEG=0&PLAYEND=0&PLAYCHN=0&YEAR=24&MON=3&DAY=21&ttxplayer=1&ttxver=1&rate=1")
	if err != nil {
		log.Fatal("parse:", err)
	}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		log.Printf("recv: %s", message)
	}
}
