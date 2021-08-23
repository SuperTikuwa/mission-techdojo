package dbctl

import (
	"fmt"
	"strconv"

	"github.com/SuperTikuwa/mission-techdojo/model"
)

func Query() error {
	db := gormConnect()
	defer db.Close()
	user := model.User{}
	user.ID = 1
	characters := []model.Character{}
	ownershipSlice := []model.UserOwnedCharacter{}

	if result := db.Table("user_owned_characters").Select("characters.id,user_owned_characters.user_character_id,characters.name").Joins("join characters on user_owned_characters.character_id = characters.id").Where("user_owned_characters.user_id = ?", user.ID).Scan(&characters).Scan(&ownershipSlice); result.Error != nil {
		writeLog(failure, result.Error)
		return result.Error
	}

	userCharacters := []model.UserCharacter{}
	for i := range characters {
		userCharacters = append(userCharacters, model.UserCharacter{
			Name:            characters[i].Name,
			CharacterID:     strconv.Itoa(characters[i].ID),
			UserCharacterID: ownershipSlice[i].UserCharacterID,
		})
	}

	fmt.Println(userCharacters)
	return nil
}
