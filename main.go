package main

import (
	"fmt"
	"github.com/NaySoftware/go-fcm"
	latihan_module "github.com/ferriantitus/latihan-module"
)

const (
	serverKey = "AAAAuZhqvDY:APA91bE66rRGmyDlYZQ6rALPfX2Z1zdnHbrnyA7gy4ER6rFEybRB_20WqwBXxqowij8h6s7r_qwL1sxY19NE-Y742ciQODpiG4GxTe0pcIIxGdiNR9E9VDUJmEJ8lxmStyoyV7_H-suL"
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