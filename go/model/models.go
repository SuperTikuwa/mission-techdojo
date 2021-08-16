package model

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Token string `json:"token,omitempty"`
}

func (u *User) GenerateToken() {
	nonce := time.Now().Format("2006-01-02 15:04:05")
	token := sha256.Sum256([]byte(u.Name + nonce))
	u.Token = hex.EncodeToString(token[:])
}

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type LinkUserAndCharacter struct {
	UserID      int `gorm:"user_id"`
	CharacterID int `gorm:"character_id"`
}
