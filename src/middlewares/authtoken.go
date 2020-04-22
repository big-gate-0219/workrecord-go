package middlewares

import (
	"strconv"
	"fmt"
	"github.com/pkg/errors"
	"time"
	"os"
	"models"
	"github.com/dgrijalva/jwt-go"
)

const (
	idKey = "sub"
	uidKey = "aud"
	emailKey ="email"
	nameKey ="name"
	iatKey ="iat"
	expKey = "exp"
)

func Generate(user models.User) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
	                           jwt.MapClaims{
								idKey: strconv.FormatUint(user.ID, 10),
								uidKey: user.UID,
								emailKey: user.MailAddress,
								nameKey: user.Name,
								iatKey: time.Now().Unix(),
								expKey: time.Now().Add(time.Hour * 24).Unix(),
							   })
	return jwtToken.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}

func Parse(jwtToken string) (*models.User, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.Wrapf(err, "%s is expired", jwtToken)
			} else {
				return nil, errors.Wrapf(err, "%s is invalid", jwtToken)
			}
		} else {
			return nil, errors.Wrapf(err, "%s is invalid", jwtToken)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s", jwtToken)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", jwtToken)
	}

	idString, ok := claims[idKey].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", idKey, jwtToken)
	}
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("don't parse %s from %s", idKey, idString)
	}

	uid, ok := claims[uidKey].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", uidKey, jwtToken)
	}

	email, ok := claims[emailKey].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", emailKey, jwtToken)
	}

	name, ok := claims[nameKey].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", nameKey, jwtToken)
	}

	return &models.User{
		ID: id,
		UID: uid,
		MailAddress: email,
		Name: name,
	}, nil
}