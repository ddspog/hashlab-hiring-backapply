package user

import (
	"time"

	"github.com/ddsgok/gql"
	"github.com/ddsgok/str"
	"github.com/ddspog/hashlab-hiring-backapply/service1/model/req"
)

// User represents the user modeled on this service DB. Contains only
// needed information for the service. BDay is a converted type for
// ease of use.
type User struct {
	ID   string
	BDay time.Time
}

// New creates a new user, using information received.
func New(id string, bday string) (p *User, err error) {
	var t time.Time
	t, err = time.Parse("2006-01-02", bday)
	p = &User{id, t}
	return
}

// FetchByID retrieve a user from DB, using his ID as key.
func FetchByID(id string) (p *User, err error) {
	var res gql.Response
	if res, err = req.UserByID(id).Run(); err == nil {
		if res.Get("User.#").Int() > 0 {
			p, err = New(id, res.Get("User.0.date_of_birth").String())
		} else {
			err = str.New("no user found with id=%v", id).Error()
		}
	}

	return
}
