package stringsandrunes

func checkingPassword(password string) bool {

	resultLen := false
	resultChar := false
	resultNum := false

	if password != "" {
		if len(password) >=5 && len(password) <= 12 {
			resultLen = true
		}

		for _,passcode := range password {

			if passcode >= 'A' && passcode <= 'Z' {
				resultChar = true
			}

			if passcode >= '0' && passcode <= '9' {
				resultNum = true
			}
		}
	}
	return resultChar && resultLen && resultNum
}