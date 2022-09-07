package util

import "fmt"

func LoginKey(uuid string) string {
	return fmt.Sprintf("login_key:%s", uuid)
}
