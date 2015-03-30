package changemail

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/convert"

	"bitbucket.org/ikeikeikeike/antenna/lib/mailer"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

func CreateTmpuser(email string) (*models.Tmpuser, error) {
	if models.Users().Filter("email", email).Exist() {
		msg := "入力されたメールアドレスは既に登録されています。"
		return nil, errors.New(msg)
	}

	tu := &models.Tmpuser{}
	tu.Using = "changemail"
	tu.Email = email
	tu.Token = convert.GetRandomString(32)
	tu.ExpiredAt = time.Now().AddDate(0, 0, 1) // next 1 days

	return tu, tu.Insert()
}

/*
 */
func RegisterMail(tu *models.Tmpuser, u *models.User) (*models.User, error) {
	var (
		err error
		msg string
	)

	u.Email = tu.Email
	err = u.Update("Email", "Updated")
	if err != nil {
		return u, err
	}

	// one time tmpuser
	if err = tu.Delete(); err != nil {
		msg = "[rest.RegisterPassword] Tmpuser delete err:"
		beego.Warning(msg, err)
	}

	return u, err
}

/*
	Send reset password mail to gate(tmp) user.
*/
func SendChangeMail(tu *models.Tmpuser, confirm string) error {
	m := mailer.NewMailer()
	err := m.SendMail(
		tu.Email,
		fmt.Sprintf("「%s」アカウント、メールアドレスのご確認", beego.AppConfig.String("SiteName")),
		fmt.Sprintf(`<strong>最後のステップです</strong><br /><br />

		メールアドレスを確認して****アカウントの設定を完了してください。<br />
		以下のリンクをクリックしてください。<br/>
		<a href="%s">%s</a><br /><br />
		
		自分のアカウントではない場合はこちらより<a href="mailto:%s">ご連絡ください</a><br /><br />
		
		---<br />
		Copyright (c)2015 %s All rights reserved.<br />
		本メール掲載内容の無断複製・転載を固く禁じます。<br />
		`, confirm, confirm, beego.AppConfig.String("Email"), beego.AppConfig.String("SiteName")),
	)

	if err != nil {
		beego.Error("[SendChangeMail] Error:", err, "(TmpuserID=", tu.Id, ")")
		return errors.New(fmt.Sprintf("メールの送信に失敗しました: %s", err.Error()))
	}
	return nil
}

/*
	return Tmpuser
*/
func ReceiveChangeMail(token string) (*models.Tmpuser, error) {
	if token == "" {
		return nil, errors.New("token is blank")
	}

	qs := models.Tmpusers().Filter("expired_at__gte", time.Now())
	qs = qs.Filter("token", token).Filter("using", "changemail")
	if !qs.Exist() {
		msg := "トークンが見つからないまたはトークンが古い可能性があります"
		return nil, errors.New(msg)
	}

	tu := &models.Tmpuser{Token: token}
	if err := tu.Read("Token"); err != nil {
		return nil, err
	}

	return tu, nil
}
