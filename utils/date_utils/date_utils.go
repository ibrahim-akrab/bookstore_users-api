package date_utils

import "time"

func GetNowString() string {
	return time.Now().UTC().Format(time.ANSIC)
}
