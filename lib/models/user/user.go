package user

import (
	"errors"

	"bitbucket.org/ikeikeikeike/antenna/models"

	"github.com/astaxie/beego/orm"
	"github.com/ikeikeikeike/gopkg/convert"
)

func UpdateUser(u *models.User) (int64, error) {
	if err := u.IsValid(); err != nil {
		return 0, err
	}
	user := make(orm.Params)
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Password) > 0 {
		user["Password"] = convert.StrTo(u.Password).Md5()
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	num, err := models.Users().Filter("Id", u.Id).Update(user)
	return num, err
}
