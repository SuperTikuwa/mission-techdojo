package model

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// MySQL Model

type User struct {
	ID    int    `gorm:"id,omitempty"`
	Name  string `gorm:"name"`
	Token string `gorm:"token"`
}

func (u *User) GenerateToken() {
	nonce := time.Now().Format("2006-01-02 15:04:05")
	token := sha256.Sum256([]byte(u.Name + nonce))
	u.Token = hex.EncodeToString(token[:])
}

type Character struct {
	ID     int    `gorm:"id"`
	Name   string `gorm:"name"`
	Weight int    `gorm:"weight"`
}

type UserOwnedCharacter struct {
	UserID          int    `gorm:"user_id"`
	CharacterID     int    `gorm:"character_id"`
	UserCharacterID string `gorm:"user_character_id"`
}

type Gacha struct {
	ID   int    `gorm:"id"`
	Name string `gorm:"name"`
}

type GachaEmissions struct {
	GachaID     int `gorm:"gacha_id"`
	CharacterID int `gorm:"character_id"`
	Weight      int `gorm:"weight"`
}

// HTTP Model

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	Name string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
}

type GachaDrawRequest struct {
	Times   int `json:"times"`
	GachaID int `json:"gacha_id"`
}

type GachaResult struct {
	CharacterID string `json:"characterID"`
	Name        string `json:"name"`
}

type GachaDrawResponse struct {
	Results []GachaResult `json:"results"`
}

type UserCharacter struct {
	UserCharacterID string `json:"userCharacterID"`
	CharacterID     string `json:"characterID"`
	Name            string `json:"name"`
}

type CharacterListResponse struct {
	Characters []UserCharacter `json:"characters"`
}
