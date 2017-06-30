package main

import (

	//	"time"

	//	"github.com/patrickmn/go-cache"

	"fmt"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/message/template"
)

// BookUpdateMSG 消息结构
type BookUpdateMSG struct {
	BookName        template.DataItem `json:"book_name"`
	EndChapterTitle template.DataItem `json:"end_chapter_title"`
}

func main() {
	//	fansCache = cache.New(5*time.Minute, 30*time.Second)
	wxAppID := "wx702b93aef72f3549"
	wxAppSecret := "8b69f45fc737a938cbaaffc05b192394"
	ats := core.NewDefaultAccessTokenServer(wxAppID, wxAppSecret, nil)
	clt := core.NewClient(ats, nil)

	// fmt.Println(clt.Token())

	bn := BookUpdateMSG{
		BookName:        template.DataItem{Value: "xxx", Color: ""},
		EndChapterTitle: template.DataItem{Value: "ooo", Color: ""},
	}

	msg := template.TemplateMessage2{
		ToUser:     "o7UTkjr7if4AQgcPmveQ5wJ5alsA",
		TemplateId: "9S_pcl3gklUmE7jOrHa2mTOcGPhj5_wnuWt2F6MH_qQ",
		URL:        "",
		Data:       bn,
	}

	msgID, err := template.Send(clt, msg)
	fmt.Println(msgID)
	fmt.Println(err)
}
