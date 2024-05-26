package container

import (
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/config"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/middlewares"
	"github.com/go-chi/jwtauth/v5"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

type Container struct {
	Middlewares
	Services
	Controllers
}

type Middlewares struct {
	AuthMw func(http.Handler) http.Handler
}

type Services struct {
	app.AuthService
	app.UserService
	app.OrganizationService
	app.RoomService
	app.DeviceService
}

type Controllers struct {
	AuthController         controllers.AuthController
	UserController         controllers.UserController
	OrganizationController controllers.OrganizationController
	RoomController         controllers.RoomController
	DeviceController       controllers.DeviceController
}

func New(conf config.Configuration) Container {
	tknAuth := jwtauth.New("HS256", []byte(conf.JwtSecret), nil)
	sess := getDbSess(conf)

	sessionRepository := database.NewSessRepository(sess)
	userRepository := database.NewUserRepository(sess)
	organizationRepository := database.NewOrganizationRepository(sess)
	roomRepository := database.NewRoomRepository(sess)
	deviceRepository := database.NewDeviceRepository(sess)

	userService := app.NewUserService(userRepository)
	authService := app.NewAuthService(sessionRepository, userRepository, tknAuth, conf.JwtTTL)
	organizationService := app.NewOrganizationService(organizationRepository, roomRepository)
	roomServise := app.NewRoomService(roomRepository, organizationRepository)
	deviceSevise := app.NewDeviceService(deviceRepository)

	authController := controllers.NewAuthController(authService, userService)
	userController := controllers.NewUserController(userService, authService)
	organizationController := controllers.NewOrganizationController(organizationService)
	roomController := controllers.NewRoomController(roomServise)
	deviceController := controllers.NewDeviceController(deviceSevise)

	authMiddleware := middlewares.AuthMiddleware(tknAuth, authService, userService)

	return Container{
		Middlewares: Middlewares{
			AuthMw: authMiddleware,
		},
		Services: Services{
			authService,
			userService,
			organizationService,
			roomServise,
			deviceSevise,
		},
		Controllers: Controllers{
			authController,
			userController,
			organizationController,
			roomController,
			deviceController,
		},
	}
}

func getDbSess(conf config.Configuration) db.Session {
	sess, err := postgresql.Open(
		postgresql.ConnectionURL{
			User:     conf.DatabaseUser,
			Host:     conf.DatabaseHost,
			Password: conf.DatabasePassword,
			Database: conf.DatabaseName,
		})
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}
	return sess
}
