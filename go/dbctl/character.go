package dbctl

import (
	"strconv"

	"github.com/SuperTikuwa/mission-techdojo/model"
)

func SelectAllCharacters() []model.Character {
	db := gormConnect()
	defer db.Close()
	characters := make([]model.Character, 0)
	db.Find(&characters)

	return characters
}

func SelectAllUserCharacter(user model.User) ([]model.UserCharacter, error) {
	db := gormConnect()
	defer db.Close()
	userCharacters := make([]model.UserCharacter, 0)
	links := make([]model.UserAndCharacterLink, 0)

	if result := db.Where("user_id = ?", user.ID).Find(&links); result.Error != nil {
		return nil, result.Error
	}

	for _, link := range links {
		character := model.Character{}
		if result := db.Where("id = ?", link.CharacterID).Find(&character); result.Error != nil {
			return nil, result.Error
		}
		characterResponse := model.UserCharacter{}
		characterResponse.UserCharacterID = link.UserCharacterID
		characterResponse.CharacterID = strconv.Itoa(character.ID)
		characterResponse.Name = character.Name

		userCharacters = append(userCharacters, characterResponse)
	}

	return userCharacters, nil
}
