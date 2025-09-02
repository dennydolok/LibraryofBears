package usecase

import (
	"BearLibrary/Models"
	"BearLibrary/internal/domain/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(user *Models.User) error
	Login(email, password string) (*Models.User, error)
	GetByID(id uint) (*Models.User, error)
	GetAll() ([]Models.User, error)
	Update(user *Models.User) error
	Delete(id uint) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) Register(user *Models.User) error {
	// Check if user already exists
	existingUser, _ := u.userRepo.GetByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.userRepo.Create(user)
}

func (u *userUsecase) Login(email, password string) (*Models.User, error) {
	user, err := u.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	// Don't return password
	user.Password = ""
	return user, nil
}

func (u *userUsecase) GetByID(id uint) (*Models.User, error) {
	user, err := u.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	user.Password = ""
	return user, nil
}

func (u *userUsecase) GetAll() ([]Models.User, error) {
	users, err := u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	// Clear passwords
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

func (u *userUsecase) Update(user *Models.User) error {
	// If password is being updated, hash it
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return u.userRepo.Update(user)
}

func (u *userUsecase) Delete(id uint) error {
	return u.userRepo.Delete(id)
}