package today

import "time"

// MockNow mocks the now time function, it makes return a different
// date.
func MockNow(today string) {
	now = func() time.Time {
		t, _ := time.Parse("2006-01-02", today)
		return t
	}
}

// ResetMocks reset the functions defined to the initial purpose.
// Making mocking really simple.
func ResetMocks() {
	now = time.Now
}
