package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Jordan Braz", "jb@gmail.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Jordan Braz", user.Name)
	assert.Equal(t, "jb@gmail.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Jordan Braz", "jb@gmail.com", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("123456"))
	assert.NotEqual(t, "123456", user.Password)
}
