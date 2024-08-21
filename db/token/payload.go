package token

import (
	"errors"
	"time"

	// "github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ClaimStrings []string

type NumericDate struct {
	time.Time
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
	// Audience string `json:""`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil,nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	out := jwt.NumericDate{payload.IssuedAt}
	return &out,nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	out := jwt.NumericDate{payload.IssuedAt}
	return &out,nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return "raman",nil
}

func (payload *Payload) GetSubject() (string, error) {
	return "",nil
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	out := jwt.NumericDate{payload.ExpiredAt}
	return &out,nil
}


//  GetExpirationTime() (*NumericDate, error)
// 	GetIssuedAt() (*NumericDate, error)
// 	GetNotBefore() (*NumericDate, error)
// 	GetIssuer() (string, error)
// 	GetSubject() (string, error)
// 	GetAudience() (ClaimStrings, error)

// func (payload *Payload) GetExpirationTime() (*time.Time, error) {
// 	return &payload.ExpiredAt, nil
// }
// func (payload *Payload) GetIssuedAt() (*time.Time, error) {
// 	return &payload.IssuedAt, nil
// }
// func (payload *Payload) GetNotBefore() (*time.Time, error)
// func (payload *Payload) GetIssuer() (string, error) {
// 	return "", nil
// }
// func (payload *Payload) GetSubject() (string, error) {
// 	return "", nil
// }

// func (payload *Payload) GetAudience() (ClaimStrings, error) {
// 	return ClaimStrings{""}, nil
// }

// func NewPayload(username string, duration time.Duration) ( *jwt.MapClaims, error) {
// 	tokenID,err := uuid.NewRandom()
// 	if err!=nil {
// 		return nil, err
// 	}

// 	// payload := &Payload{
// 	// 	ID: tokenID,
// 	// 	Username: username,
// 	// 	IssuedAt: time.Now(),
// 	// 	ExpiredAt: time.Now().Add(duration),
// 	// }

// 	return &jwt.MapClaims{
// 		"username": username,
// 		"exp":time.Now().Add(duration),
// 		"ID": tokenID,
// 		"IssuedAt": time.Now(),

// 	},err
// }
