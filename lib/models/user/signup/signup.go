package signup

import (
	"errors"
	"fmt"
	"time"

	"bitbucket.org/ikeikeikeike/antenna/lib/mailer"
	"bitbucket.org/ikeikeikeike/antenna/lib/verify"
	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/str"
)

func CreateTmpuser(tu *models.Tmpuser) (int64, error) {
	var (
		err error
		msg string
	)

	if models.Users().Filter("email", tu.Email).Exist() {
		msg = "入力されたメールアドレスは既に登録されています。"
		return 0, errors.New(msg)
	}
	if models.Blogs().Filter("Rss", tu.Rss).Exist() {
		msg = "入力されたRSSは既に登録されています。"
		return 0, errors.New(msg)
	}

	tu.Using = "signup"
	tu.Password = convert.StrTo(tu.Password).Md5()
	tu.Token = convert.GetRandomString(32)
	tu.ExpiredAt = time.Now().AddDate(0, 0, 1) // next 1 days

	err = tu.Insert()
	if err != nil {
		return 0, err
	}

	return tu.Id, err
}

/*
	1. Add user
	2. Add blog
	3. Summarize blog rss
*/
func RegisterUser(tu *models.Tmpuser) (*models.User, error) {
	var (
		err error
		msg string
	)

	if models.Blogs().Filter("Rss", tu.Rss).Exist() {
		msg = "入力されたRSSは既に登録されています。"
		return nil, errors.New(msg)
	}

	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		beego.Error("[signup.RegisterUser] Transaction err:", err)
		msg = "エラーが発生しました。しばらくしてからご登録ください。"
		return nil, errors.New(msg)
	}

	defer func() {
		if err != nil {
			o.Rollback()
			msg = "[signup.RegisterUser] Transaction Rollback err:"
			beego.Error(msg, tu.Email, err)
			return
		}

		err = o.Commit()
		if err != nil {
			msg = "[signup.RegisterUser] Transaction Commit err:"
			beego.Error(msg, tu.Email, err)
		}
	}()

	u := &models.User{Email: tu.Email, Password: tu.Password}
	err = u.Insert()
	if err != nil {
		return u, err
	}

	b := &models.Blog{Rss: tu.Rss, User: u,
		Mediatype: tu.Mediatype, Adsensetype: tu.Adsensetype}

	feed, ok := verify.GetFeed(str.Clean(b.Rss))
	if ok && len(feed.Channels) > 0 {
		b.Url = str.Clean(feed.Channels[0].Links[0].Href)
	}

	err = b.Insert()
	if err != nil {
		msg = "[signup.RegisterUser] blog insert:"
		beego.Error(msg, err)
		return u, err
	}

	// one time tmpuser
	if err = tu.Delete(); err != nil {
		msg = "[signup.RegisterUser] Tmpuser delete err:"
		beego.Warning(msg, err)
	}

	// summarize blog
	// if err = summarize.RssFeed(); err != nil {
	// msg = "[signup.RegisterUser] summarize.RssFeed:"
	// beego.Warning(msg, err)
	// }

	return u, err
}

/*
	Send activation mail to gate(tmp) user.
*/
func SendSignupMail(tu *models.Tmpuser, confirm string) error {
	sn := beego.AppConfig.String("SiteName")

	m := mailer.NewMailer()
	err := m.SendMail(
		tu.Email,
		fmt.Sprintf("「%s」掲載用アカウントの作成誠にありがとうございます。", sn),
		fmt.Sprintf(`ようこそ「%s」へ<br/><br/>
		
		こちらはアカウント作成確認メールです。<br/>
		下記リンクより登録を完了させてください。<br/>
		<a href="%s">%s</a> <br/><br/>
		
		「%s」のアカウントを作成していないのにも関わらずこのメッセージを受け取った場合は、<br/>
		恐れ入りますがこちらより<a href="mailto:%s">ご連絡ください</a><br /><br />

		---<br />
		Copyright (c)2015 %s All rights reserved.<br />
		本メール掲載内容の無断複製・転載を固く禁じます。<br />
		`, sn, confirm, confirm, sn, beego.AppConfig.String("Email"), sn),
	)

	if err != nil {
		beego.Error("[SendSignupMail] Error:", err, " TmpuserID=", tu.Id)
		return errors.New(fmt.Sprintf("メールの送信に失敗しました: %s", err.Error()))
	}
	return nil
}

/*
	return Tmpuser
*/
func ReceiveSignupMail(token string) (*models.Tmpuser, error) {
	if token == "" {
		return nil, errors.New("token is blank")
	}

	qs := models.Tmpusers().Filter("expired_at__gte", time.Now())
	qs = qs.Filter("token", token).Filter("using", "signup")
	if !qs.Exist() {
		msg := `トークンが見つからないまたはトークンが古い可能性があります。
		恐れ入りますが再度ご登録ください。
		`
		return nil, errors.New(msg)
	}

	tu := &models.Tmpuser{Token: token}
	if err := tu.Read("Token"); err != nil {
		return nil, err
	}

	return tu, nil
}
