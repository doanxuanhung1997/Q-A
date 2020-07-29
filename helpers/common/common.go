package common

import (
	"../../api/user/model"
	"../../db"
	"../../helpers/config"
	"../../helpers/constant"
	"crypto/sha256"
	"encoding/hex"
	"gopkg.in/mgo.v2/bson"
	"math/rand"
	"regexp"
	"time"
	"unicode/utf8"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UTC().UnixNano()))

/*General Token Register Code Client*/
func GenerateTokenString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[seededRand.Intn(len(letterRunes))]
	}
	return string(b)
}

//hash password from user pass
func HashPassword(pass string) string {
	env := config.GetEnvValue()
	mixPass := pass + env.Secret.Salt
	bytePass := []byte(mixPass)

	hash := sha256.New()
	hash.Write(bytePass)
	passwordHash := hex.EncodeToString(hash.Sum(nil))
	//return  passwordHash
	return env.Secret.Buffer + passwordHash
}

/*CheckEmailExit return true if email existed in user table */
func CheckEmailExit(value string) bool {
	var user model.User
	err := db.Collection(model.CollectionUser).FindOne(db.GetContext(), bson.M{"email": value}).Decode(&user)
	if err != nil {
		return false
	}
	return true
}

/*IsEmpty return True if string is empty*/
func IsEmpty(value string) bool {
	if utf8.RuneCountInString(value) == 0 || value == "" {
		return true
	}
	return false
}

/*CheckValidationEmail return true when string is email and match regex*/
func CheckValidationEmail(email string) bool {
	re := regexp.MustCompile(`(?:[a-z0-9!#$%&'*+/=?^_'{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_'{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)
	if !re.MatchString(email) {
		return false
	}
	return true
}

/*Check string is format date*/
func CheckFormatDate(value string) bool {
	_, err := time.Parse(constant.TimeLayout, value)
	if err != nil {
		return false
	}
	return true
}

/*CheckLength return true if string length input passed min,max input*/
func CheckLength(value string, minLength int, maxLength int) bool {
	if utf8.RuneCountInString(value) < minLength || utf8.RuneCountInString(value) > maxLength {
		return false
	}
	return true
}

/*Check string is format number*/
func CheckIsNumber(value string) bool {
	re := regexp.MustCompile(`^[0-9]*$`)
	if !re.MatchString(value) {
		return false
	}
	return true
}

/*Check contain Special in string*/
func CheckSpecialCharacters(value string) bool {
	re := regexp.MustCompile(`[!@#$%^&*(),._?'"` + `:{+}|<>/-]`)
	if re.MatchString(value) {
		return false
	}
	return true
}
