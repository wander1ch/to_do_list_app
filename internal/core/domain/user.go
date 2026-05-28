package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/wander1ch/to_do_list_app/internal/core/errors"
)

type User struct {
	ID       int
	Version  int
	FullName string
	PhoneNumber *string
}

func NewUser (id int, version int, fullName string, phoneNumber *string) User {
	return User{
		ID: id,	
		Version: version,
		FullName: fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(fullName string, phoneNumber *string) User {
	return NewUser(UninitializedID, UninitializedVersion, fullName, phoneNumber)
}

func (u *User) Validate() error {
	if len(u.FullName) < 2 || len(u.FullName) > 100 {
		return fmt.Errorf("invalid full name len: %d: %w", len(u.FullName), core_errors.ErrInvalidArgument)
	}
	
	if u.PhoneNumber != nil {
		if len(*u.PhoneNumber) < 10 || len(*u.PhoneNumber) > 15 {
			return fmt.Errorf("invalid phone number len: %d: %w", len(*u.PhoneNumber), core_errors.ErrInvalidArgument)
		}
		re := regexp.MustCompile(`^\+[0-9]{10,15}$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid phone number format: %s: %w", *u.PhoneNumber, core_errors.ErrInvalidArgument)
		}
	}
	return nil
}


type UserPatch struct {
	Name Nullable[string]
	Phone Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.Name.Set && p.Name.Value == nil {
		return fmt.Errorf("invalid name: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate patch: %w", err)
	}

	tmp := *u

	if patch.Name.Set {
		tmp.FullName = *patch.Name.Value
	}

	if patch.Phone.Set {
		tmp.PhoneNumber = patch.Phone.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp
	return nil
}