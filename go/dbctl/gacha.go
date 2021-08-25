package dbctl

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/SuperTikuwa/mission-techdojo/model"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func DrawGacha(user model.User, gachaParams model.GachaDrawRequest) (model.GachaDrawResponse, error) {

	var characters []model.Character
	var err error
	if gachaParams.GachaID == 0 {
		characters, err = selectAllCharacters()
	} else {
		characters, err = selectCharactersByGachaID(gachaParams.GachaID)
	}

	if err != nil {
		writeLog(failure, "DrawGacha", err)
		return model.GachaDrawResponse{}, err
	}

	if len(characters) == 0 {
		return model.GachaDrawResponse{}, errors.New("no characters")
	}

	lookupTable := createLookupTable(characters)

	resultIDs := lotteryGacha(lookupTable, gachaParams.Times)

	results := extractResultsFromIDs(characters, resultIDs, user)
	if err := insertGachaResults(results, user); err != nil {
		return model.GachaDrawResponse{}, err
	}

	drawResponse := model.GachaDrawResponse{
		Results: results,
	}
	return drawResponse, nil
}

func GachaExists(gachaID int) bool {
	db := gormConnect()
	defer db.Close()
	gachas := make([]model.Gacha, 0)
	if result := db.Table("gachas").Where("id = ?", gachaID).Find(&gachas); result.Error != nil {
		writeLog(failure, "GachaExists", result.Error)
		return false
	}

	fmt.Println(gachas)

	return len(gachas) > 0
}

func selectAllCharacters() ([]model.Character, error) {
	db := gormConnect()
	defer db.Close()
	characters := make([]model.Character, 0)
	if result := db.Table("character_emissions").Select("characters.id,characters.name,emission_weight").Joins("JOIN characters ON character_emissions.character_id = characters.id").Scan(&characters); result.Error != nil {
		writeLog(failure, result.Error)
		return nil, result.Error
	}

	return characters, nil
}

func selectCharactersByGachaID(gachaID int) ([]model.Character, error) {
	db := gormConnect()
	defer db.Close()

	characters := make([]model.Character, 0)
	if result := db.Table("character_emissions").Select("characters.id,characters.name,emission_weight").Joins("JOIN characters ON character_emissions.character_id = characters.id").Where("character_emissions.gacha_id = ?", gachaID).Scan(&characters); result.Error != nil {
		writeLog(failure, "selectCharactersByGachaID", result.Error)
		return nil, result.Error
	}
	return characters, nil
}

func lotteryGacha(table []int, times int) []int {
	results := make([]int, 0)
	for i := 0; i < times; i++ {
		results = append(results, table[rand.Intn(len(table))])
	}
	return results
}

func createLookupTable(characters []model.Character) []int {
	lookupTable := make([]int, 0)
	for _, c := range characters {
		for i := 0; c.EmissionWeight > i; i++ {
			lookupTable = append(lookupTable, c.ID)
		}
	}
	return lookupTable
}

func insertGachaResults(results []model.GachaResult, user model.User) error {
	db := gormConnect()
	defer db.Close()

	ownershipInterfaces := make([]interface{}, 0, len(results))

	for _, result := range results {
		idStr := strings.Split(result.CharacterID, "-")[1]
		characterID, err := strconv.Atoi(idStr)
		if err != nil {
			return err
		}

		ownership := model.UserOwnedCharacter{
			UserID:          user.ID,
			CharacterID:     characterID,
			UserCharacterID: result.CharacterID,
		}

		ownershipInterfaces = append(ownershipInterfaces, ownership)
	}

	if err := gormbulk.BulkInsert(db, ownershipInterfaces, 3000); err != nil {
		writeLog(failure, "insertGachaResults", err)
		return err
	}

	return nil
}

func generateCharacterID(userID, characterID int) string {
	return strconv.Itoa(userID) + "-" + strconv.Itoa(characterID) + "-" + time.Now().Format("20060102150405")
}

func extractResultsFromIDs(characters []model.Character, characterIDs []int, user model.User) []model.GachaResult {
	// Returnするスライスの初期化
	results := make([]model.GachaResult, 0)

	// 一つずつ取り出して，IDからCharacterを取得し，スライスに追加
	for characterNumber, id := range characterIDs {
		for _, character := range characters {
			if id == character.ID {
				result := model.GachaResult{
					CharacterID: generateCharacterID(user.ID, character.ID) + strconv.Itoa(characterNumber),
					Name:        character.Name,
				}
				results = append(results, result)
				break
			}
		}
	}

	return results
}
