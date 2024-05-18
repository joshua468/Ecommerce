package auth

import (
	"time"

	"strconv"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string,error) {
	expiration:= time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)
token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims {
	"userID" : strconv.Iota(userID),
	"expiredAt" : time.Now().Add(expiration).Unix(),
})
tokenString,err := token.SignedString(secret)
if err!= nil {
	return "",err
}
return tokenString,nil
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandleFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)
		token,err:= validateToken(tokenString)

		if err!= nil {
			log.Printf("failed to validate token: %v",err)
			permissionDenied(w)
			return
		}
		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		userID,_ := strconv(str)

		u,err := store.GetUserByID(userID)
		if err!= nil {
			log.Printf("failed to get  user by id: %v", err)
			permissionDenied(w)
			return
		}
		ctx := r.Context()
		ctx := context.WithValue(ctx,"userID",u.ID)
		r = r.WithContext(ctx)
		handlerfunc()
	}
}

func getTokenFromRequest(r *http.Request) string {
	tokenAuth :=  r.Header.Get("Authorization")
	if tokenAuth != "" {
		return tokenAuth
	}
	return  ""
}

func validateToken(t string)  (*jwt.Token,error) {
	return  jwt.Parse(t,func(t *jwt.Token) (interface{},error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil,fmt.Erorrf("unexpected signing method : %v",t.Header["alg"])
		}
		return  []byte(config.Envs.JWTSecret),nil
	})

}

func  permissionDenied( ) {
	utils.WriteError(w,http.StatusForbidden,fmt.Errorf("Permission denied"))
}