package usecases

import (
	"github.com/banerwai/gather/service"
	"github.com/martini-contrib/sessionauth"
)

// MyUserModel can be any struct that represents a user in my system
type UserModel struct {
	Id            string `form:"id"`
	Email         string `form:"email"`
	Password      string `form:"password"`
	authenticated bool   `form:"-"`
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() sessionauth.User {
	return &UserModel{}
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *UserModel) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *UserModel) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *UserModel) IsAuthenticated() bool {
	return u.authenticated
}

func (u *UserModel) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *UserModel) GetById(id interface{}) error {
	_id := id.(string)
	var _user_service service.UserService
	_user, _err := _user_service.GetUserByIDDto(_id)
	if _err != nil {
		return nil
	}
	u.Id = _id
	u.Email = _user.Email
	return nil
}
