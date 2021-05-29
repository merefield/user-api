package model

import (
	"time"

	"github.com/form3tech-oss/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// EmailTokenModel is an abstract model which can be used for objects from which
// we derive redirect emails (email confirmation, password reset and such)
type EmailTokenModel struct {
	IDRecord
	Reference   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	EmailSent   bool      `bun:"index;not null"`
	EmailSentAt *time.Time
	ExpiresAt   time.Time `bun:"index;not null"`
}

type EmailTokenClaimsModel struct {
	Username  string `json:"username"`
	Reference string `json:"reference"`
	jwt.StandardClaims
}

type MailgunEmailModel struct {
	Recipient string
	Subject   string
	Template  string
}

// TableName specifies table name
func (ac *EmailTokenModel) TableName() string {
	return "oauth_email_tokens"
}

// NewEmailToken creates new OauthEmailToken instance
func NewOauthEmailToken(expiresIn *time.Duration) *EmailTokenModel {
	return &EmailTokenModel{
		EmailSent: false,
		ExpiresAt: time.Now().UTC().Add(*expiresIn),
	}
}

// NewOauthEmailTokenClaims creates new NewOauthEmailTokenClaims instance
func NewOauthEmailTokenClaims(email string, emailToken *EmailTokenModel) *EmailTokenClaimsModel {
	return &EmailTokenClaimsModel{
		Username:  email,
		Reference: emailToken.Reference.String(),
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: emailToken.ExpiresAt.Unix(),
		},
	}
}

// NewOauthEmail
func NewOauthEmail(email, subject, template string) *MailgunEmailModel {
	return &MailgunEmailModel{
		Recipient: email,
		Subject:   subject,
		Template:  template,
	}
}
