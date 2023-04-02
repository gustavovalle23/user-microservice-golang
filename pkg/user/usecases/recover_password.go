package usecases

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

type RecoverPasswordUseCase struct {
	userRepository domain.UserRepository
	mailer         domain.Mailer
}

func NewRecoverPasswordUseCase(userRepository domain.UserRepository, mailer domain.Mailer) *RecoverPasswordUseCase {
	return &RecoverPasswordUseCase{
		userRepository: userRepository,
		mailer:         mailer,
	}
}

func (uc *RecoverPasswordUseCase) Execute(email string) error {
	user, err := uc.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrUserNotFound
	}

	token, err := uc.generateResetToken()
	if err != nil {
		return err
	}

	user.ResetToken = token
	user.ResetTokenExpiresAt = time.Now().Add(time.Hour * 24)

	err = uc.mailer.SendPasswordResetEmail(user.Email)
	if err != nil {
		return err
	}

	err = uc.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (uc *RecoverPasswordUseCase) generateResetToken() (string, error) {
	const tokenLength = 32

	bytes := make([]byte, tokenLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	token := base64.StdEncoding.EncodeToString(bytes)

	return token, nil
}
