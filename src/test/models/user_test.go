package models_test

import (
	"io"
	"os"
	"fmt"
	// "log"
	"time"
	"context"
	"testing"
	"path/filepath"

	"gorm.io/gorm"
	// "gorm.io/driver/mysql"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"

	// "github.com/ory/dockertest/v3"
	// "github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

var mockDB *gorm.DB

var emptyUser = models.User{	
	Id: 0,
	UserName: "",
	EmailAddress: "",
	Password: "",
	IsAdmin: false,
	CanLogin: true,
	CreatedAt: time.Now(),	
	UpdatedAt: time.Now(),
}

var user = models.User{	
	Id: 1,
	UserName: "user",
	EmailAddress: "user@email.com",
	Password: "password",
	IsAdmin: true,
	CanLogin: true,
	CreatedAt: time.Now(),	
	UpdatedAt: time.Now(),
}

func MariaDBExists() bool {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if container.Image == "mariadb:10.2" {
			return true
		}
	}
	return false
}

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	// endpoint := "unix:///Users/makotonakai/.docker/run/docker.sock"
	// pool, err := dockertest.NewPool(endpoint)
	// if err != nil {
	// 	log.Fatalf("Could not construct pool: %s", err)
	// }

	// // uses pool to try to connect to Docker
	// err = pool.Client.Ping()
	// if err != nil {
	// 	log.Fatalf("Could not connect to Docker: %s", err)
	// }

	mariaDBExists := MariaDBExists()

	// A new MariaDB container will be created if one is not found
	if mariaDBExists == true {
	} else {
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}
		cli.NegotiateAPIVersion(ctx)

		imageName := "mariadb:10.2"
		out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
		if err != nil {
			panic(err)
		}
		io.Copy(os.Stdout, out)

		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		parent := filepath.Dir(pwd)

		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image: "mariadb:10.2",
			ExposedPorts: []string{
				"3306",
			},
			Env: []string{
				"MYSQL_ROOT_HOST=%",
				"MYSQL_DATABASE=test",
				"MYSQL_ROOT_PASSWORD=secret",
			},
			Volumes: map[string]{
				parent + "/db/init.sql": "/docker-entrypoint-initdb.d/init.sql"
			}


		}, nil, nil, nil, "")
		if err != nil {
			panic(err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			panic(err)
		}

			fmt.Println(resp.ID)
		}
		
}

func TestIsEmailAddressEmptySuccess(t *testing.T) {
	result := models.IsEmailAddressEmptyOrNull(emptyUser)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsEmailAddressEmptyFail(t *testing.T) {
	result := models.IsEmailAddressEmptyOrNull(user)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsUserNameEmptySuccess(t *testing.T) {
	result := models.IsUserNameEmptyOrNull(emptyUser)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsUserNameEmptyFail(t *testing.T) {
	result := models.IsUserNameEmptyOrNull(user)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestIsPasswordSuccess(t *testing.T) {
	result := models.IsPasswordEmptyOrNull(emptyUser)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestIsPasswordEmptyFail(t *testing.T) {
	result := models.IsPasswordEmptyOrNull(user)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestErrorExistsSuccess(t *testing.T) {
	errorMessageList := []string{config.EmailAddressIsEmpty}
	result := models.ErrorsExist(errorMessageList)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestErrorExistsFail(t *testing.T) {
	errorMessageList := []string{}
	result := models.ErrorsExist(errorMessageList)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestAlreadyExistsSuccess(t *testing.T) {
	result := models.AlreadyExists(mockDB, user)
	expectedResult := true
	assert.Equal(t, result, expectedResult)
}

func TestAlreadyExistsFail(t *testing.T) {
	result := models.AlreadyExists(mockDB, emptyUser)
	expectedResult := false
	assert.Equal(t, result, expectedResult)
}

func TestGetUserNameFromUserIdSuccess(t *testing.T) {
	userId := 1
	userName := models.GetUserNameFromUserId(mockDB, userId)
	expectedUserName := "user"
	assert.Equal(t, userName, expectedUserName)
}

func TestGetUserNameFromUserIdFail(t *testing.T) {
	userId := 100
	userName := models.GetUserNameFromUserId(mockDB, userId)
	expectedUserName := "user"
	assert.NotEqual(t, userName, expectedUserName)
}

func TestGetUserIdFromEmailAddressSuccess(t *testing.T) {
	emailAddress := "user@email.com"
	userId := models.GetUserIdFromEmailAddress(mockDB, emailAddress)
	expectedUserId := 1
	assert.Equal(t, userId, expectedUserId)
}

func TestGetUserIdFromEmailAddressFail(t *testing.T) {
	emailAddress := "user!@email.com"
	userId := models.GetUserIdFromEmailAddress(mockDB, emailAddress)
	expectedUserId := 1
	assert.NotEqual(t, userId, expectedUserId)
}

func TestResetPassword(t *testing.T) {
	id := 1
	newPassword := "hoge"
	err := models.ResetPassword(mockDB, id, newPassword)
	assert.Nil(t, err)
}


