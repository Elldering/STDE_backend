package RegisterService

import (
	"STDE_proj/internal/repositories/RegisterRepository"
	"STDE_proj/utils/hash"
	"log"
)

func Register(login string, password string) error {

	hashedPassword, err := hash.HashPassword(password)
	if err != nil {
		log.Printf("Ошибка при хешировании пароля: %v", err)
		return err
	}
	password = hashedPassword
	return RegisterRepository.Register(login, password)
}
