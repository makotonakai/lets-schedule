package controllers_test

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
	// "github.com/MakotoNakai/lets-schedule/config"
	// "github.com/MakotoNakai/lets-schedule/database"
	"github.com/MakotoNakai/lets-schedule/controllers"
)

const (
	host = "localhost"
	port = "3306"

	image         = "mariadb:10.2"
	containerName = "controllertestdb"
	architecture  = "linux/arm64/v8"
)

var mockDB *gorm.DB

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

func pullNewImage(ctx context.Context, cli *client.Client) error {

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
			"MYSQL_DATABASE=db",
			"TZ=Asia/Tokyo",
			"BIND_ADDRESS=127.0.0.1",
		},
	}

	portBindings := map[nat.Port][]nat.PortBinding{nat.Port(port): {{HostIP: host, HostPort: port}}}
	hostConfig := &container.HostConfig{
		PortBindings: portBindings,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: "/Users/makotonakai/golang-docker/database/data",
				Target: "/var/lib/mariadb",
			},
			{
				Type:   mount.TypeBind,
				Source: "/Users/makotonakai/golang-docker/database/my.cnf",
				Target: "/etc/mysql/conf.d/my.cnf",
			},
			{
				Type:   mount.TypeBind,
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
	if mariaDBDoesNotExist {
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

func TestCreateUserSuccess(t *testing.T) {

	userJSON := `{"user_name":"newuser","email_address":"newuser@email.com","password":"password","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestCreateUserNoName(t *testing.T) {

	userJSON := `{"user_name":"","email_address":"newuser@email.com","password":"password","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["Username is empty"]`, rec.Body.String())
	}
}

func TestCreateUserNoEmail(t *testing.T) {

	userJSON := `{"user_name":"newuser","email_address":"","password":"password","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["Email address is empty"]`, rec.Body.String())
	}
}

func TestCreateUserNoPassword(t *testing.T) {

	userJSON := `{"user_name":"newuser","email_address":"newuser@email.com","password":"","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["Password is empty"]`, rec.Body.String())
	}
}

func TestCreateUserUserNameOnly(t *testing.T) {

	userJSON := `{"user_name":"newuser","email_address":"","password":"","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["Email address is empty", "Password is empty"]`, rec.Body.String())
	}
}

func TestCreateUserEmailAddressOnly(t *testing.T) {

	userJSON := `{"user_name":"","email_address":"newuser@email.com","password":"","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["Username is empty", "Password is empty"]`, rec.Body.String())
	}
}

func TestCreateUserPassword(t *testing.T) {

	userJSON := `{"user_name":"","email_address":"","password":"password","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["Username is empty", "Email address is empty"]`, rec.Body.String())
	}
}

func TestCreateUserAlreadyExists(t *testing.T) {

	userJSON := `{"user_name":"newuser","email_address":"newuser@email.com","password":"password","is_admin":false,"can_login":false}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/signup", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["The given user already exists"]`, rec.Body.String())
	}
}

func TestResetPasswordSuccess(t *testing.T) {

	newPasswordJSON := `{"new_password":"password!"}`

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/user/1/reset-password", strings.NewReader(newPasswordJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, controllers.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, `["The given user already exists"]`, rec.Body.String())
	}
}
