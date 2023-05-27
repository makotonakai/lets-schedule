package models_test

import (

	"fmt"
	"time"
	"bytes"
	"context"
	"testing"

	"gorm.io/gorm"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/container"
	"github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/stretchr/testify/assert"
	"github.com/MakotoNakai/lets-schedule/config"
	"github.com/MakotoNakai/lets-schedule/models"
)

const (
	host = "localhost"
	port = "3306"

	image = "mariadb:10.2"
	containerName = "testdb"
	architecture = "amd4"
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



func MariaDBDoesNotExist() bool {
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
			return false
		}
	}
	return true
}

func pullNewImage(ctx context.Context, cli *client.Client) (error) {

	read, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(read)
	imageRespBytes := buf.String()
	imageRespString := string(imageRespBytes)
	fmt.Printf("%s\n", imageRespString)

	return nil
}

func createDBContainer(ctx context.Context, cli *client.Client) container.CreateResponse {

	containerConfig := &container.Config{
		Image:        image,
		ExposedPorts: nat.PortSet{port: struct{}{}},
		Env: []string{
			"MYSQL_ROOT_HOST=%",
			"MYSQL_ROOT_USER=root",
			"MYSQL_ROOT_PASSWORD=root",
			"MYSQL_USER=user",
      "MYSQL_PASSWORD=password",
			"MYSQL_DATABASE=test",
		},
	}

	portBindings := map[nat.Port][]nat.PortBinding{nat.Port(port): {{HostIP: host, HostPort: port}}}
	hostConfig := &container.HostConfig{
		PortBindings: portBindings,
		Mounts: []mount.Mount{
			{
				Type: mount.TypeBind,
				Source: "/Users/makotonakai/golang-docker/database/data",
				Target: "/var/lib/mariadb",
			},
			{
				Type: mount.TypeBind,
				Source: "/Users/makotonakai/golang-docker/database/my.cnf",
				Target: "/etc/mysql/conf.d/my.cnf",
			},
			{
				Type: mount.TypeBind,
				Source: "/Users/makotonakai/golang-docker/database/initdb.d",
				Target: "/docker-entrypoint-initdb.d",
			},
		},
	}

	platform := &v1.Platform{
		Architecture: architecture,
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, platform, containerName)
	if err != nil {
			panic(err)
	}

	return resp
}




func TestMain(m *testing.M) {

	ctx := context.Background()
	mariaDBDoesNotExist := MariaDBDoesNotExist()
	if (mariaDBDoesNotExist) {
		cli, err := client.NewClientWithOpts(client.FromEnv)
		if err != nil {
			panic(err)
		}

		err = pullNewImage(ctx, cli)
		if err != nil {
			panic(err)
		}

		resp := createDBContainer(ctx, cli)
		err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
		if err != nil {
				panic(err)
		}
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


