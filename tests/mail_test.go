package main

import (
	"testing"

	"github.com/astaxie/beego"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/exp/ses"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"
)

var (
	Auth   *aws.Auth
	Mailer *ses.SES
)

func TestSES(t *testing.T) {
	email := ses.NewEmail()
	email.AddTo("jpt.ne.co.jp@gmail.com")
	email.SetSource("jp.ne.co.jp@gmail.com")
	email.SetSubject("test2")
	email.SetBodyHtml("test3")

	err := Mailer.SendEmail(email)
	if err != nil {
		t.Error(err)
	}
}

func init() {
	auth := aws.Auth{AccessKey: beego.AppConfig.String("AwsAccessKey"), SecretKey: beego.AppConfig.String("AwsSecretKey")}
	Mailer = ses.NewSES(auth, aws.Region{Name: "us-east-1", SESEndpoint: "https://email.us-east-1.amazonaws.com"})
}
