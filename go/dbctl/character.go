package dbctl

import (
	"strconv"

	"github.com/SuperTikuwa/mission-techdojo/model"
)

func selectAllCharacters() []model.Character {
	db := gormConnect()
	defer db.Close()
	characters := make([]model.Character, 0)
	db.Find(&characters)

	return characters
}

func SelectAllUserCharacter(user model.User) ([]model.UserCharacter, error) {
	db := gormConnect()
	defer db.Close()

	characters := []model.Character{}
	ownershipSlice := []model.UserOwnedCharacter{}

	if result := db.Table("user_owned_characters").Select("characters.id,user_owned_characters.user_character_id,characters.name").Joins("join characters on user_owned_characters.character_id = characters.id").Where("user_owned_characters.user_id = ?", user.ID).Scan(&characters).Scan(&ownershipSlice); result.Error != nil {
		writeLog(failure, result.Error)
		return nil, result.Error
	}

	userCharacters := []model.UserCharacter{}
	for i := range characters {
		userCharacters = append(userCharacters, model.UserCharacter{
			Name:            characters[i].Name,
			CharacterID:     strconv.Itoa(characters[i].ID),
			UserCharacterID: ownershipSlice[i].UserCharacterID,
		})
	}

	return userCharacters, nil
}
