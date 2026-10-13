package main

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	req, err := selectClient(db, clientID)
	require.NoError(t, err)
	assert.Equal(t, clientID, req.ID)
	assert.NotEmpty(t, req.FIO)
	assert.NotEmpty(t, req.Login)
	assert.NotEmpty(t, req.Birthday)
	assert.NotEmpty(t, req.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	req, err := selectClient(db, clientID)
	require.True(t, errors.Is(err, sql.ErrNoRows))
	assert.Empty(t, req.FIO)
	assert.Empty(t, req.Login)
	assert.Empty(t, req.Birthday)
	assert.Empty(t, req.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	cl.ID = id

	req, err := selectClient(db, cl.ID)
	require.NoError(t, err)
	require.Equal(t, req, cl)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	id, err := insertClient(db, cl)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	cl.ID = id

	req, err := selectClient(db, id)
	require.NoError(t, err)
	require.Equal(t, req, cl)

	err = deleteClient(db, id)
	require.NoError(t, err)

	req, err = selectClient(db, id)
	require.True(t, errors.Is(err, sql.ErrNoRows))
}
