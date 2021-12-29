package main

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	latihan_module "github.com/ferriantitus/latihan-module"
)

const (
	serverKey = "Your-Key"
	topic = "/topics/news"
	title = "Test Tugas GoModules-Notif"
	tanggal = "29/12/2021"
)

func main(){
	data := map[string]string{
		"title": title,
		"tanggal": tanggal,
	}

	notifPayload := fcm.NotificationPayload{
		Title: "Ada Berita Baru",
		Body:  title,
	}

	c := fcm.NewFcmClient(serverKey)
	c.SetNotificationPayload(&notifPayload)
	c.NewFcmMsgTo(topic, data)

	status, err := c.Send()


	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
	fmt.Println(latihan_module.PrintKalimat())
}