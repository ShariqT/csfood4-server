package core

func AuthenticateUser(username string, password string, db DB) (*User, error) {
	user, err := db.AuthUser(username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
