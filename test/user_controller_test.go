package test

import (
	"backend-golang/model/domain"
	"backend-golang/repository"
	"context"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func truncateUser(db *sql.DB) {
	db.Exec(`TRUNCATE TABLE "user" RESTART IDENTITY`)
}

func TestCreateUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Darul Ikhsan", "email":"darulikhssan@gmail.com", "password":"123456789"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/user", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Darul Ikhsan", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateUserFailed(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"", "email":"darulikhssan@gmail.com", "password":"123456789"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/user", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()

	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name:     "Darul",
		Email:    "darulikhssan@gmail.com",
		Password: "123456789",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"Darul Ikhsan", "email":"darulikhssan@gmail.com", "password":"123456789"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/user/"+strconv.Itoa(user.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Darul Ikhsan", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateUserFailed(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name:     "Darul",
		Email:    "darulikhssa@gmail.com",
		Password: "12345",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name":"", "email":"darulikhssan@gmail.com", "password":"12345"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/user/"+strconv.Itoa(user.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name:     "Darul",
		Email:    "darulikhssan@gmail.com",
		Password: "12345",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/user/"+strconv.Itoa(user.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, user.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Darul", responseBody["data"].(map[string]interface{})["name"])
}

func TestGetUserFailed(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/user/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name:     "Darul",
		Email:    "darulikhssan@gmail.com",
		Password: "12345",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/user/"+strconv.Itoa(user.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteUserFailed(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/user/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListUserSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user1 := userRepository.Save(context.Background(), tx, domain.User{
		Name:     "Darul",
		Email:    "darulikhssan@gmail.com",
		Password: "12345",
	})
	user2 := userRepository.Save(context.Background(), tx, domain.User{
		Name:     "Ikhsan",
		Email:    "darulikhssan@gmail.com",
		Password: "12345",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/user", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var users = responseBody["data"].([]interface{})

	userResponse1 := users[0].(map[string]interface{})
	userResponse2 := users[1].(map[string]interface{})

	assert.Equal(t, user1.Id, int(userResponse1["id"].(float64)))
	assert.Equal(t, user1.Name, userResponse1["name"])

	assert.Equal(t, user2.Id, int(userResponse2["id"].(float64)))
	assert.Equal(t, user2.Name, userResponse2["name"])
}

func TestListUserFailed(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/user", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(bytes, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
