package auth

import (
	"strconv"

	API "github.com/epointpayment/key_management_system/app/services/api"
	User "github.com/epointpayment/key_management_system/app/services/user"

	"github.com/labstack/echo"
)

// BasicValidator is a validator used for basic auth middleware
func BasicValidator(username, password string, c echo.Context) (isValid bool, err error) {
	sa := API.New()

	programID, err := strconv.Atoi(c.QueryParam("program_id"))
	if err != nil {
		err = User.ErrInvalidProgramID
		return
	}

	isValid, err = sa.DoAuth(username, password, programID)
	if !isValid || err != nil {
		return
	}

	us := User.New()
	user, err := us.GetByUsername(username, programID)
	if err != nil {
		return
	}

	c.Set("userID", user.ID)

	return
}
