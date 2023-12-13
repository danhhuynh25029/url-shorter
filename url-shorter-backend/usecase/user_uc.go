package usecase

import "errors"

func Login() error {
	return errors.New("login is failed")
}

func Logout() error {
	return errors.New("")
}

func Resigter() error {
	return nil
}
