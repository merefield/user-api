package jwt

import (
	"fmt"
	"time"

	uuid "github.com/google/uuid"

	"github.com/resonatecoop/user-api/model"

	jwt "github.com/form3tech-oss/jwt-go"
)

// New instantiates new JWT service
func New(key string, duration int, algo string) *JWT {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid signing method")
	}
	return &JWT{
		key:      []byte(key),
		duration: time.Duration(duration) * time.Minute,
		algo:     signingMethod,
	}
}

// JWT contains data necessery for jwt auth
type JWT struct {
	// Secret key used for signing.
	key []byte

	// Duration for which the jwt token is valid.
	duration time.Duration

	// JWT signing algorithm
	algo jwt.SigningMethod
}

// GenerateToken generates new jwt token
func (j *JWT) GenerateToken(u *model.AuthUser) (string, error) {
	t := jwt.NewWithClaims(j.algo, jwt.MapClaims{
		"id":  u.ID.String(),
		"t":   u.TenantID,
		"u":   u.Username,
		"r":   u.Role,
		"exp": time.Now().Add(j.duration).Unix(),
	})

	return t.SignedString(j.key)
}

// ParseToken parses the bearer token
func (j *JWT) ParseToken(token string) (*model.AuthUser, error) {
	claims, err := j.verifyToken(token)
	if err != nil {
		return nil, err
	}

	id, ok := claims["id"]
	if !ok {
		return nil, fmt.Errorf("unauthorized: no id claim present")
	}

	tenantId, ok := claims["t"]
	if !ok {
		return nil, fmt.Errorf("unauthorized: no tenant_id claim present")
	}

	username, ok := claims["u"]
	if !ok {
		return nil, fmt.Errorf("unauthorized: no username claim present")
	}

	email, ok := claims["e"]
	if !ok {
		return nil, fmt.Errorf("unauthorized: no email claim present")
	}

	role, ok := claims["r"]
	if !ok {
		return nil, fmt.Errorf("unauthorized: no role claim present")
	}

	uid, err := uuid.Parse(id.(string))
	if err != nil {
		return nil, fmt.Errorf("id claim is not a valid uuid")
	}

	return &model.AuthUser{
		ID:       uid,
		TenantID: int32(tenantId.(float64)),
		Username: username.(string),
		Email:    email.(string),
		Role:     model.AccessRole(role.(float64)),
	}, nil

}

func (j *JWT) verifyToken(token string) (map[string]interface{}, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != j.algo {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.key, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("could not parse provided token")
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("jwt token could not be verified")
}
