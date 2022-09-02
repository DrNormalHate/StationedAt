package user

import (
	"StationedAtAuth/data"
	"database/sql"
	"regexp"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
}

type LoginRequest struct {
	Password string `json:"password" sql:"password" validate:"password"`
	Match    bool   `json:"match" sql:"match"`
}

type NewUser struct {
	DisplayName string `json:"display_name" sql:"display_name" validate:"-"`
	Email       string `json:"email" sql:"email" validate:"email"`
	Password    string `json:"password" sql:"password" validate:"gt=9,password"`
	Hash        string `json:"hash" sql:"hash" validate:"-"`
}

type LoggedInUser struct {
	DisplayName string          `json:"display_name" sql:"display_name" validate:"required"`
	Email       EmailComponents `json:"email" sql:"email" validate:"structonly"`
	Branch      string          `json:"branch" sql:"branch" validate:"-"`
	Component   string          `json:"component" sql:"component" validate:"-"`
	DeersStatus string          `json:"deers_status" sql:"deers_status" validate:"-"`
	DutyStation string          `json:"duty_station" sql:"duty_station" validate:"-"`
}

type EmailComponents struct {
	Address string `json:"address" sql:"address" validate:"required"`
	Domain  string `json:"domain" sql:"domain" validate:"required"`
}

//function to validate password

func (u *Client) LoginUser() error {
	//login function

	return nil
}

func (u *LoggedInUser) CreateLoggedInUser() error {
	// Remove NewUser and create LoggedInUser

	return nil
}

func (u *Client) CreateClient() error {
	// Determine LoggedIn status

	return nil
}

func (r *LoginRequest) MatchHashes() bool {
	return false
}

func (u *NewUser) CreateNewUser() error {
	//Register a new user

	validate := validator.New()
	// Password must be 10 or more characters, contain 1 upper case letter, 1 lower case letter, 1 digit, and one special character
	//[!\"#$%&'()*+,-./:;<=>?@\\^_`~]
	validate.RegisterValidation("password", validatePassword)

	// validate newuser struct
	if err := validate.Struct(u); err != nil {
		panic(err.(validator.ValidationErrors))
	}

	// load newuser substruct of newpassword
	if err := u.CreatePassword(); err != nil {
		panic(err)
	}
	//check if email already exists

	//check if phone already exists (can this be done in the query itself)

	//compose sql and send

	if data, err := data.CreateDBEngine(); err == nil {
		defer data.Close()
		if tx, err := data.Begin(); err == nil {
			defer tx.Rollback()
			if stmt, err := tx.Prepare("INSERT INTO `Users` (`display_name`, `email`, `password`) VALUES (?, ?, ?)"); err == nil {
				defer stmt.Close()
				if _, err := stmt.Exec(u.DisplayName, u.Email, u.Hash); err == nil {
					if err := tx.Commit(); err == nil {
						return nil
					} else {
						return err
					}
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}
}

func (u *LoggedInUser) Sql2User(res *sql.Rows) error {
	if err := res.Scan(&u.DisplayName); err != nil {
		return err
	}
	return nil
}

func (u *NewUser) CreatePassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	u.Hash = string(hash)

	return nil
}

// Validation Functions

func validatePassword(fl validator.FieldLevel) bool {
	if match, err := regexp.Match("[^A-Za-z0-9!\"#$%&'()*+,-./:;<=>?@\\^_`~]", []byte(fl.Field().String())); match {
		if err != nil {
			panic("unaccepted character")
		}
		return false
	}
	if match, err := regexp.Match("[A-Z]+", []byte(fl.Field().String())); match && err == nil {
		if match, err := regexp.Match("[a-z]+", []byte(fl.Field().String())); match && err == nil {
			if match, err := regexp.Match("[0-9]+", []byte(fl.Field().String())); match && err == nil {
				if match, err := regexp.Match("[!\"#$%&'()*+,-./:;<=>?@\\^_`~]+", []byte(fl.Field().String())); match && err == nil {
					return true
				} else {
					if err != nil {
						panic("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
					}
				}
			} else {
				if err != nil {
					panic("0-9")
				}
			}
		} else {
			if err != nil {
				panic("a-z")
			}
		}
	} else {
		if err != nil {
			panic("A-z")
		}
	}
	return false
}
