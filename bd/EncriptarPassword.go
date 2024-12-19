package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPassword(pass sting) (string, err) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]bytes(pass), costo)
	if err != null {
		return err.Error(), err
	}
	return string(bytes), nil
}
