package today

import (
	"time"

	"github.com/ddspog/hashlab-hiring-backapply/service1/model/user"
)

// IsUserBDay checks if today is a spcified user birthday.
func IsUserBDay(id string) (ans bool, err error) {
	var usr *user.User
	if usr, err = user.FetchByID(id); err == nil {
		ans = usr.BDay.Format("01-02") == Date()
	}

	return
}

// IsBlackFriday checks if today is black friday.
func IsBlackFriday() (ans bool) {
	ans = Date() == "11-25"
	return
}

// Date returns current time parsed as a date informing day and month.
func Date() (d string) {
	d = now().Format("01-02")
	return
}

var (
	// now returns current time.
	now = time.Now
)
