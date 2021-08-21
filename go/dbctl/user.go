package dbctl

import (
	"github.com/SuperTikuwa/mission-techdojo/model"
)

func InsertNewUser(newUser model.User) error {
	db := gormConnect()
	defer db.Close()

	if result := db.Create(&newUser); result.Error != nil {
		writeLog(failure, result.Error)
		return result.Error
	}

	return nil
}

func SelectUserByToken(token string) model.User {
	db := gormConnect()
	defer db.Close()
	var user model.User
	if err := db.Where("token = ?", token).First(&user).Error; err != nil {
		writeLog(failure, err)
		return model.User{}
	}
	return user
}

func UpdateUser(user model.User) error {
	db := gormConnect()
	defer db.Close()

	if err := db.Model(&model.User{}).Where("token = ?", user.Token).Update(user).Error; err != nil {
		writeLog(failure, err)
		return err
	}
	return nil
}

func UserExists(user model.User) bool {
	db := gormConnect()
	defer db.Close()

	checkResult := SelectUserByToken(user.Token)

	if checkResult.Name != "" || checkResult.Token != "" {
		return true
	}

	return false
}
