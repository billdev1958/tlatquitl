package usecases

import (
	"context"
	"user/internal/domain/account"
	"user/internal/domain/user"
	apperror "user/pkg/error"
	"user/pkg/hash"
	"user/pkg/uuid"
	"user/pkg/validations"
)

func validateAccount(a *account.Account) (*account.Account, error) {
	if validations.IsValidEmailSyntax(a.Email) == false {
		return &account.Account{}, apperror.ErrInvalidEmail
	}

	err := validations.IsValidPassword(a.Password)
	if err != nil {
		return &account.Account{}, err
	}

	return a, nil

}

func (uc *useCases) Create(ctx context.Context, u *user.User, a *account.Account) (user.User, error) {
	userID := uuid.NewUUID()
	accountID := uuid.NewUUID()

	if u.Name == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	if u.Lastname1 == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	if u.Lastname2 == "" {
		return user.User{}, apperror.ErrEmptyInput
	}

	u.ID = userID

	_, err := validateAccount(a)
	if err != nil {
		return user.User{}, err
	}

	a.ID = accountID
	a.UserID = userID

	hashedPassword, err := hash.HashPassword(a.Password)
	if err != nil {
		return user.User{}, err
	}
	a.Password = hashedPassword

	createdUser, err := uc.userRepo.Create(ctx, u, a)
	if err != nil {
		return user.User{}, err
	}

	return createdUser, nil

}
