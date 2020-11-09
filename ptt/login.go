package ptt

func Login(userID string, passwd string, ip string, isHashed bool) (bool, error) {
	if userID == STR_GUEST {
		return true, nil
	}

	if !isValidUserID(userID) {
		return false, nil
	}

	user, err := initUser(userID)
	if err != nil {
		return false, err
	}
	if user == nil {
		logAttempt(userID, ip)
		return false, nil
	}

	isValid, err := checkPasswd(user.PasswdBig5, []byte(passwd), isHashed)
	if err != nil {
		return false, err
	}
	if !isValid {
		logAttempt(userID, ip)
		return false, nil
	}

	return false, nil
}
