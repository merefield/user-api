package userserver

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	uuid "github.com/google/uuid"
	uuidpkg "github.com/resonatecoop/user-api/pkg/uuid"

	"github.com/uptrace/bun"

	"github.com/resonatecoop/user-api/model"
	pbUser "github.com/resonatecoop/user-api/proto/user"
)

// Server implements the UserService
type Server struct {
	db  *bun.DB
	sec Securer
}

// New creates an instance of our server
func New(db *bun.DB, sec Securer) *Server {
	return &Server{db: db, sec: sec}
}

// Securer represents password securing service
type Securer interface {
	Hash(string) string
	Password(string, ...string) bool
}

// AddUser gets a user to the in-memory store.
func (s *Server) AddUser(ctx context.Context, user *pbUser.UserAddRequest) (*pbUser.Empty, error) {

	requiredErr := checkRequiredAddAttributes(user)
	if requiredErr != nil {
		return nil, requiredErr
	}

	newUser := &model.User{
		Username:               user.Username,
		FullName:               user.FullName,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Member:                 user.Member,
		NewsletterNotification: user.NewsletterNotification,
	}
	_, err := s.db.NewInsert().Model(newUser).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// GetUser Gets a user from the DB
func (s *Server) GetUser(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserPublicResponse, error) {

	u := new(model.User)

	err := s.db.NewSelect().Model(u).
		Column("user.*").
		Where("id = ?", user.Id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.UserPublicResponse{
		Username:       u.Username,
		FullName:       u.FullName,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Member:         u.Member,
		Country:        u.Country,
		FollowedGroups: uuidpkg.ConvertUUIDToStrArray(u.FollowedGroups),
	}, nil
}

// string username = 1; // required
// string full_name = 2; // required
// string first_name = 3;
// string last_name = 4;
// string country = 5;
// bool member = 6;
// repeated string personas = 7;
// repeated string owned_groups = 8;
// repeated string followed_groups = 9;

// GetUserRestricted intended for privileged roles only supplies more detailed, private info about user.
func (s *Server) GetUserRestricted(ctx context.Context, user *pbUser.UserRequest) (*pbUser.UserPrivateResponse, error) {

	u := new(model.User)

	err := s.db.NewSelect().Model(u).
		Column("user.*").
		Where("id = ?", user.Id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.UserPrivateResponse{
		Id:                     u.ID.String(),
		Username:               u.Username,
		FullName:               u.FullName,
		FirstName:              u.FirstName,
		LastName:               u.LastName,
		Member:                 u.Member,
		RoleId:                 u.RoleID,
		TenantId:               u.TenantID,
		FollowedGroups:         uuidpkg.ConvertUUIDToStrArray(u.FollowedGroups),
		NewsletterNotification: u.NewsletterNotification,
	}, nil
}

// DeleteUser Deletes a user from the DB
func (s *Server) DeleteUser(ctx context.Context, user *pbUser.UserRequest) (*pbUser.Empty, error) {
	u := new(model.User)

	_, err := s.db.NewUpdate().
		Model(u).
		Set("deleted_at = ?", time.Now().UTC()).
		Where("id = ?", user.Id).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// UpdateUser updates a users basic attributes
func (s *Server) UpdateUser(ctx context.Context, UserUpdateRequest *pbUser.UserUpdateRequest) (*pbUser.Empty, error) {

	var updatedUserValues = make(map[string]interface{})

	if UserUpdateRequest.Username != nil {
		updatedUserValues["username"] = *UserUpdateRequest.Username
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*UserUpdateRequest.Username) {
			return nil, errors.New("username must be a valid email")
		}
	}
	if UserUpdateRequest.FirstName != nil {
		updatedUserValues["first_name"] = *UserUpdateRequest.FirstName
	}
	if UserUpdateRequest.LastName != nil {
		updatedUserValues["last_name"] = *UserUpdateRequest.LastName
	}
	if UserUpdateRequest.FullName != nil {
		updatedUserValues["full_name"] = *UserUpdateRequest.FullName
	}
	if UserUpdateRequest.NewsletterNotification != nil {
		updatedUserValues["newsletter_notification"] = *UserUpdateRequest.NewsletterNotification
	}

	updatedUserValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserValues).TableExpr("users").Where("id = ?", UserUpdateRequest.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// UpdateUserRestricted updates a users more restricted attributes
func (s *Server) UpdateUserRestricted(ctx context.Context, UserUpdateRestrictedRequest *pbUser.UserUpdateRestrictedRequest) (*pbUser.Empty, error) {

	var updatedUserValues = make(map[string]interface{})

	if UserUpdateRestrictedRequest.Username != nil {
		updatedUserValues["username"] = *UserUpdateRestrictedRequest.Username
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(*UserUpdateRestrictedRequest.Username) {
			return nil, errors.New("username must be a valid email")
		}
	}
	if UserUpdateRestrictedRequest.FirstName != nil {
		updatedUserValues["first_name"] = *UserUpdateRestrictedRequest.FirstName
	}
	if UserUpdateRestrictedRequest.LastName != nil {
		updatedUserValues["last_name"] = *UserUpdateRestrictedRequest.LastName
	}
	if UserUpdateRestrictedRequest.FullName != nil {
		updatedUserValues["full_name"] = *UserUpdateRestrictedRequest.FullName
	}
	if UserUpdateRestrictedRequest.Member != nil {
		updatedUserValues["member"] = *UserUpdateRestrictedRequest.Member
	}
	if UserUpdateRestrictedRequest.RoleId != nil {
		updatedUserValues["role_id"] = *UserUpdateRestrictedRequest.RoleId
	}
	if UserUpdateRestrictedRequest.TenantId != nil {
		updatedUserValues["tenant_id"] = *UserUpdateRestrictedRequest.TenantId
	}
	if UserUpdateRestrictedRequest.NewsletterNotification != nil {
		updatedUserValues["newsletter_notification"] = *UserUpdateRestrictedRequest.NewsletterNotification
	}

	updatedUserValues["updated_at"] = time.Now().UTC()

	rows, err := s.db.NewUpdate().Model(&updatedUserValues).TableExpr("users").Where("id = ?", UserUpdateRestrictedRequest.Id).Exec(ctx)

	if err != nil {
		return nil, err
	}

	number, _ := rows.RowsAffected()

	if number == 0 {
		return nil, errors.New("warning: no rows were updated")
	}

	return &pbUser.Empty{}, nil
}

// ResetUserPassword reset's a user's password
func (s *Server) ResetUserPassword(ctx context.Context, ResetUserPasswordRequest *pbUser.ResetUserPasswordRequest) (*pbUser.Empty, error) {
	err := checkRequiredResetPasswordAttributes(ResetUserPasswordRequest, s)

	if err != nil {
		return nil, err
	}

	hashedPassword := s.sec.Hash(ResetUserPasswordRequest.Password)

	u := new(model.User)

	_, err = s.db.NewUpdate().
		Model(u).
		Set("updated_at = ?", time.Now().UTC()).
		Set("password = ?", hashedPassword).
		Where("username = ?", ResetUserPasswordRequest.Username).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return &pbUser.Empty{}, nil
}

// ListUsers lists all users in the store.
func (s *Server) ListUsers(ctx context.Context, Empty *pbUser.Empty) (*pbUser.UserListResponse, error) {

	var users []model.User
	var results pbUser.UserListResponse

	err := s.db.NewSelect().
		Model(&users).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	for _, user := range users {
		var result pbUser.UserPrivateResponse
		result.Id = uuid.UUID.String(user.ID)
		result.Username = user.Username
		result.FullName = user.FullName
		result.FirstName = user.FirstName
		result.LastName = user.LastName
		result.Member = user.Member
		result.NewsletterNotification = user.NewsletterNotification
		result.FollowedGroups = uuidpkg.ConvertUUIDToStrArray(user.FollowedGroups)
		results.User = append(results.User, &result)
	}

	return &results, nil
}

func checkRequiredAddAttributes(user *pbUser.UserAddRequest) error {
	if user.Username == "" || user.FullName == "" {
		var argument string
		switch {
		case user.Username == "":
			argument = "username"
		case user.FullName == "":
			argument = "full_name"
		}
		return fmt.Errorf("argument %v is required", argument)
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Username) {
		return errors.New("username must be a valid email")
	}
	return nil
}

// func checkRequiredUpdateAttributes(user *pbUser.UpdateUserRequest) error {
// 	if user.Username == "" || user.FullName == "" {
// 		var argument string
// 		switch {
// 		case user.Username == "":
// 			argument = "username"
// 		case user.FullName == "":
// 			argument = "full_name"
// 		}
// 		return fmt.Errorf("argument %v is required", argument)
// 	}
// 	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	if !re.MatchString(user.Username) {
// 		return errors.New("username must be a valid email")
// 	}
// 	return nil
// }

// func checkRequiredRestrictedUpdateAttributes(user *pbUser.UpdateUserRestrictedRequest) error {
// 	if user.Username == "" || user.FullName == "" {
// 		var argument string
// 		switch {
// 		case user.Username == "":
// 			argument = "username"
// 		case user.FullName == "":
// 			argument = "full_name"
// 		}
// 		return fmt.Errorf("argument %v is required", argument)
// 	}
// 	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
// 	if !re.MatchString(user.Username) {
// 		return errors.New("username must be a valid email")
// 	}
// 	return nil
// }

func checkRequiredResetPasswordAttributes(user *pbUser.ResetUserPasswordRequest, s *Server) error {
	if user.Username == "" || user.Password == "" {
		var argument string
		switch {
		case user.Username == "":
			argument = "email"
		case user.Password == "":
			argument = "Password"
		}
		return fmt.Errorf("argument %v is required", argument)
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(user.Username) {
		return errors.New("username must be a valid email")
	}
	if !s.sec.Password(user.Password) {
		return errors.New("password is not strong enough")
	}

	return nil
}

func getUserGroupResponse(ownerOfGroup []model.UserGroup) []*pbUser.RelatedUserGroup {
	groups := make([]*pbUser.RelatedUserGroup, len(ownerOfGroup))
	for i, group := range ownerOfGroup {
		groups[i] = &pbUser.RelatedUserGroup{Id: group.ID.String(), DisplayName: group.DisplayName}
	}
	return groups
}

func (s *Server) DerefString(str *string) string {
	if str != nil {
		return *str
	}

	return ""
}
