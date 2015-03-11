package authc

import (
	"errors"
	"time"

	"github.com/ikeikeikeike/gopkg/convert"

	"bitbucket.org/ikeikeikeike/antenna/models"
)

/*
 Get authenticated user and update logintime
*/
func Authenticate(email string, password string) (user *models.User, err error) {
	msg := "メールアドレスまたはパスワードが間違っています。"
	user = &models.User{Email: email}

	if err := user.Read("Email"); err != nil {
		if err.Error() == "<QuerySeter> no row found" {
			err = errors.New(msg)
		}
		return user, err
	} else if user.Id < 1 {
		// No user
		return user, errors.New(msg)
	} else if user.Password != convert.StrTo(password).Md5() {
		// No matched password
		return user, errors.New(msg)
	} else {
		user.Lastlogintime = time.Now()
		user.Update("Lastlogintime")
		return user, nil
	}
}
