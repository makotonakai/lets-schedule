package models_test

import (
	"os"
	"fmt"
	"log"
	"time"
	"context"
	"testing"
	"path/filepath"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
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
	endpoint := "unix:///Users/makotonakai/.docker/run/docker.sock"
	pool, err := dockertest.NewPool(endpoint)
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	pwd, err := os.Getwd()
	if err != nil {
    panic(err)
	}
	parent := filepath.Dir(pwd)

	mariaDBExists := MariaDBExists()

	// A new MariaDB container will be created if one is not found
	if mariaDBExists == true {
		continue 
	} else {
		runOptions := &dockertest.RunOptions{
			Repository: "mariadb",
			Tag: "10.2",
			// latest だと本番とマッチしなくなる場合があるのでバージョン指定
			// ポート番号は固定せずに 0 で listen する
			Env: []string{
				"MYSQL_ROOT_HOST=%",
				"MYSQL_DATABASE=test",
				"MYSQL_ROOT_PASSWORD=secret",
			},
			ExposedPorts: []string {
				"3306",
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"3306/tcp": {{HostIP: "localhost", HostPort: "3306/tcp"}},
			},
			// ここでデータベースの初期化ファイルを渡す
			Mounts: []string{
				parent + "/db/init.sql:/docker-entrypoint-initdb.d/init.sql",
			},
		}

		resource, err := pool.RunWithOptions(runOptions,
			func(config *docker.HostConfig) {
				// 処理が終了したらインスタンスを削除する
				config.AutoRemove = true
				config.RestartPolicy = docker.RestartPolicy{
					Name: "no",
				}
			})
		if err != nil {
			log.Fatalf("Could not start resource: %s", err)
		}

	if err := pool.Retry(func() error {
		var err error
		dsn := fmt.Sprintf("root:secret@tcp(localhost:%s)/test?charset=utf8mb4&parseTime=True&loc=Local", resource.GetPort("3306/tcp"))
		mockDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
		db, _ := mockDB.DB()
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)

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


