package middlewares

import (
	"../../api/user/model"
	"../../db"
	"../../helpers/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/url"
	"strings"
	"time"
)


type Claims struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	jwt.StandardClaims
}

//generate token
func GenerateJWT(username string, id string) (string, error) {
	env := config.GetEnvValue()
	expirationTime := time.Now().UTC().Add(time.Duration(env.Server.ExpireToken) * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		Id:       id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string

	var jwtKey = []byte(env.Secret.JwtSecretKey)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

var (
	ErrEmptyAuthHeader   = errors.New("auth header is empty")
	ErrInvalidAuthHeader = errors.New("auth header is invalid")
)

//get token from cookie or header
func JwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)
	if authHeader == "" {
		authHeader, _ = JwtFromCookie(c, "Token")
		return authHeader, nil
	} else {
		if authHeader == "" {
			return "", ErrEmptyAuthHeader
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Token") {
			return "", ErrInvalidAuthHeader
		}

		return parts[1], nil
	}

}

//get token from cookie
func JwtFromCookie(c *gin.Context, key string) (string, error) {
	cookie, _ := c.Request.Cookie(key)
	if cookie != nil {
		authHeader, _ := url.QueryUnescape(cookie.Value)
		if authHeader == "" {
			return "", ErrEmptyAuthHeader
		}
		return authHeader, nil
	}
	return "", nil
}

func VerifyToken(c *gin.Context) (model.User, error) {
	var jwtKey = []byte(config.GetSecret())
	var token, _ = JwtFromHeader(c, "Authorization")

	var user model.User
	claims := &Claims{}
	_, errTK := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if errTK != nil {
		return user, errTK
	}

	objectIdUser, _ := primitive.ObjectIDFromHex(claims.Id)
	errUser := db.Collection(model.CollectionUser).FindOne(db.GetContext(), bson.M{"_id": objectIdUser}).Decode(&user)

	if errUser != nil {
		return user, errors.New("token invalid")
	}
	return user, errTK
}

//
//func ExtractToken(r *http.Request) string {
//	bearToken := r.Header.Get("Authorization")
//	//normally Authorization the_token_xxx
//	strArr := strings.Split(bearToken, " ")
//	if len(strArr) == 2 {
//		return strArr[1]
//	}
//	return ""
//}
//
//func CheckToken(r *http.Request) (*jwt.Token, error) {
//	tokenString := ExtractToken(r)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		//Make sure that the token method conform to "SigningMethodHMAC"
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("ACCESS_SECRET")), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return token, nil
//}
//
//type AccessDetails struct {
//	AccessUuid string
//	UserId     uint64
//}
//
////Logout
//func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
//	token, err := CheckToken(r)
//	if err != nil {
//		return nil, err
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		accessUuid, ok := claims["access_uuid"].(string)
//		if !ok {
//			return nil, err
//		}
//		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
//		if err != nil {
//			return nil, err
//		}
//		return &AccessDetails{
//			AccessUuid: accessUuid,
//			UserId:     userId,
//		}, nil
//	}
//	return nil, err
//}
//

//
//func DeleteAuth(c *gin.Context, givenUuid string) (int64, error) {
//	deleted, err := client.Del(c, givenUuid).Result()
//	if err != nil {
//		return 0, err
//	}
//	return deleted, nil
//}
//
