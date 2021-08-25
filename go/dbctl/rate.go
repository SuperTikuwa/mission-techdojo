package dbctl

import "github.com/SuperTikuwa/mission-techdojo/model"

func CalcEmissionRate(gachaID int) (model.EmissionRateResponse, error) {
	var characters []model.Character
	var err error
	if gachaID == 0 {
		characters, err = selectAllCharacters()
	} else {
		characters, err = selectCharactersByGachaID(gachaID)
	}

	if err != nil {
		writeLog(failure, err.Error())
		return model.EmissionRateResponse{}, err
	}

	rate := model.EmissionRateResponse{}
	rate.GachaID = gachaID
	for _, character := range characters {
		rate.Characters = append(rate.Characters, model.EmissionRateCharacter{
			ID:           character.ID,
			Name:         character.Name,
			EmissionRate: float32(character.EmissionWeight) / float32(len(characters)),
		})
	}

	return rate, nil
}
