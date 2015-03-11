package reset

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/ikeikeikeike/gopkg/convert"

	"bitbucket.org/ikeikeikeike/antenna/lib/mailer"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

type ResetForm struct {
	Password   string `form:"Password" valid:"Required;MinSize(6)"`
	Repassword string `form:"Repassword" valid:"Required"`
}

func (u *ResetForm) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "2つのパスワードが一致していません")
	}
}

func ResetPassword(email string) (*models.Tmpuser, error) {

	if !models.Users().Filter("email", email).Exist() {
		msg := "入力された情報ではメールアドレスが見つかりませんでした"
		return nil, errors.New(msg)
	}

	tu := &models.Tmpuser{}
	tu.Using = "reset"
	tu.Email = email
	tu.Token = convert.GetRandomString(32)
	tu.ExpiredAt = time.Now().AddDate(0, 0, 1) // next 1 days

	return tu, tu.Insert()
}

/*
 */
func RegisterPassword(tu *models.Tmpuser) (*models.User, error) {
	var (
		err error
		msg string
	)

	u := &models.User{Email: tu.Email}
	err = u.Read("Email")
	if err != nil {
		return u, err
	}

	u.Password = convert.StrTo(tu.Password).Md5()
	err = u.Update("Password")
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
func SendResetMail(tu *models.Tmpuser, confirm string) error {
	sn := beego.AppConfig.String("SiteName")

	m := mailer.NewMailer()
	err := m.SendMail(
		tu.Email,
		fmt.Sprintf("「%s」アカウント、パスワードリセットのご確認", sn),
		fmt.Sprintf(`
		%s 様

		<br/><br />
		本メールはパスワードリセットの確認メールです。<br />
		下記リンクをクリックすることによりパスワードのリセット操作を続けられます。<br/>
		<a href="%s">%s</a><br /><br />
		
		「%s」のパスワードリセット操作をしていないのにも関わらず、<br />
		このメッセージを受け取った場合はこちらより<a href="mailto:%s">ご連絡ください</a><br /><br />

		---<br />
		Copyright (c)2015 %s All rights reserved.<br />
		本メール掲載内容の無断複製・転載を固く禁じます。<br />
		`, tu.Email, confirm, confirm, sn, beego.AppConfig.String("Email"), sn),
	)

	if err != nil {
		beego.Error("[SendResetMail] Error:", err, "(TmpuserID=", tu.Id, ")")
		return errors.New(fmt.Sprintf("メールの送信に失敗しました: %s", err.Error()))
	}
	return nil
}

/*
	return Tmpuser
*/
func ReceiveResetMail(token string) (*models.Tmpuser, error) {
	if token == "" {
		return nil, errors.New("token is blank")
	}

	qs := models.Tmpusers().Filter("expired_at__gte", time.Now())
	qs = qs.Filter("token", token).Filter("using", "reset")
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
