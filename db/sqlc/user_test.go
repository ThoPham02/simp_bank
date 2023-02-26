package db

import (
	"context"
	"testing"
	"time"

	"github.com/ThoPham02/simp_bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	hasPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username: util.RandomOwner(8),
		HashPass: hasPassword,
		FullName: util.RandomOwner(16),
		Email:    util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, user.Username)
	require.NotEmpty(t, user.CreateAt)

	require.Equal(t, user.Username, arg.Username)
	require.Equal(t, user.HashPass, arg.HashPass)
	require.Equal(t, user.FullName, arg.FullName)
	require.Equal(t, user.Email, arg.Email)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	newUser := CreateRandomUser(t)

	user, err := testQueries.GetUser(context.Background(), newUser.Username)

	require.NoError(t, err)
	require.Equal(t, newUser.Username, user.Username)
	require.Equal(t, newUser.HashPass, user.HashPass)
	require.Equal(t, newUser.FullName, user.FullName)
	require.Equal(t, newUser.Email, user.Email)
	require.WithinDuration(t, newUser.CreateAt, user.CreateAt, time.Second)
}
