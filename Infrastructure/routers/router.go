package routers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/services"
	"bete/controllers"
	"time"

	cors "github.com/itsjamie/gin-cors"

	"github.com/gin-gonic/gin"
)

//controller
var (
	jwtService services.JWTService = services.NewJWTService()

	authController          controllers.AuthController          = controllers.NewAuthController()
	userController          controllers.UserController          = controllers.NewUserController()
	scoutController         controllers.ScoutController         = controllers.NewScoutController()
	rolController           controllers.RolController           = controllers.NewRolController()
	detachmentController    controllers.DetachmentController    = controllers.NewDetachmentController()
	churchController        controllers.ChurchController        = controllers.NewChurchController()
	moduleController        controllers.ModuleController        = controllers.NewModuleController()
	subDetachmentController controllers.SubdetachmentController = controllers.NewSubdetachmentController()
	patrolController        controllers.PatrolController        = controllers.NewPatrolController()
	kanbanController        controllers.KanbanController        = controllers.NewKanbanController()
	attendanceController    controllers.AttendanceController    = controllers.NewAttendanceController()
)

func NewRouter() {

	r := gin.New()

	// Set up CORS middleware options
	config := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Authorization",
		Methods:         "GET, POST, PUT,DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	r.Use(cors.Middleware(config))

	//
	authRoutes := r.Group("api/v1/login")
	{
		authRoutes.POST("/", authController.Login)
		//authRoutes.POST("/register", authController.Register)
	}
	//userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	userRoutes := r.Group("api/v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", userController.Profile)
		userRoutes.GET("/users", userController.ListUser)
		userRoutes.GET("/:id", userController.FindUser)
		userRoutes.GET("search/:search", userController.FindUserNameLastName)
		//userRoutes.GET("/users/", userController.All)
		//userRoutes.GET("/usersroles", userController.UsersRoles)
		userRoutes.POST("/create", userController.Create)
		userRoutes.PUT("/edit", userController.Update)
		userRoutes.PUT("/update_password", userController.PasswordChange)
		userRoutes.PUT("/", userController.UpdateProfile)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	//userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	scoutRoutes := r.Group("api/v1/scout", middleware.AuthorizeJWT(jwtService))
	{
		scoutRoutes.GET("/:id", scoutController.ListKingsScouts)
		scoutRoutes.POST("/create", scoutController.Create)
		scoutRoutes.PUT("/edit", scoutController.Update)
		//scoutRoutes.DELETE("/:id", scoutController.Delete)
	}

	attendanceRoutes := r.Group("api/v1/attendance", middleware.AuthorizeJWT(jwtService))
	{
		attendanceRoutes.GET("/", attendanceController.All)
		attendanceRoutes.POST("/create", attendanceController.Create)
		attendanceRoutes.PUT("/edit", attendanceController.Update)
		attendanceRoutes.DELETE("/:id", attendanceController.Remove)
	}
	kanbanRoutes := r.Group("api/v1/kanban", middleware.AuthorizeJWT(jwtService))
	{
		kanbanRoutes.GET("/kanban", kanbanController.GetKanban)
		kanbanRoutes.GET("/count_kanban", kanbanController.CountKanban)
	}
	rolRoutes := r.Group("api/v1/rol", middleware.AuthorizeJWT(jwtService))
	{
		rolRoutes.GET("/", rolController.All)
		rolRoutes.GET("/group", rolController.AllGroupRol)
		rolRoutes.GET("/rolemodule", rolController.AllRoleModule)
		rolRoutes.POST("/create", rolController.Create)
		rolRoutes.PUT("/", rolController.Update)
		rolRoutes.DELETE("/:id", rolController.Delete)
		rolRoutes.GET("/:id", rolController.FindRol)
	}

	detachmentRoutes := r.Group("api/v1/detachment", middleware.AuthorizeJWT(jwtService))
	{
		detachmentRoutes.GET("/", detachmentController.All)
		detachmentRoutes.GET("/:id", detachmentController.FindById)
		detachmentRoutes.POST("/", detachmentController.Create)
		detachmentRoutes.PUT("/", detachmentController.Update)
		detachmentRoutes.DELETE("/:id", detachmentController.Delete)

	}
	churchRoutes := r.Group("api/v1/church", middleware.AuthorizeJWT(jwtService))
	{
		churchRoutes.GET("/", churchController.All)
		churchRoutes.GET("/:id", churchController.FindById)
		churchRoutes.POST("/", churchController.Create)
		churchRoutes.PUT("/", churchController.Update)
		churchRoutes.DELETE("/:id", churchController.Delete)

	}
	moduleRoutes := r.Group("api/v1/module", middleware.AuthorizeJWT(jwtService))
	{
		moduleRoutes.GET("/", moduleController.All)
		moduleRoutes.GET("rolemodule/:id", moduleController.ByRoleModule)
		moduleRoutes.GET("menumodule/:id", moduleController.ByRoleModule)

		moduleRoutes.GET("/:id", moduleController.FindModuleById)
		moduleRoutes.POST("/", moduleController.Create)
		moduleRoutes.PUT("/", moduleController.Update)
		moduleRoutes.DELETE("/:id", moduleController.Delete)
		moduleRoutes.POST("/rolemodule", moduleController.AddModuleRole)
		moduleRoutes.DELETE("/rolemodule/:id", moduleController.DeleteModuleRole)

	}
	subDetachmentRoutes := r.Group("api/v1/sud-detachment", middleware.AuthorizeJWT(jwtService))
	{
		subDetachmentRoutes.GET("/", subDetachmentController.All)
		subDetachmentRoutes.GET("/:id", subDetachmentController.FindById)
		subDetachmentRoutes.GET("/detachment/:id", subDetachmentController.FindByIdDetachment)
		subDetachmentRoutes.POST("/", subDetachmentController.Create)

		subDetachmentRoutes.PUT("/", subDetachmentController.Update)
		subDetachmentRoutes.DELETE("/:id", subDetachmentController.Remove)

	}
	patrolRoutes := r.Group("api/v1/patrol", middleware.AuthorizeJWT(jwtService))
	{
		patrolRoutes.GET("/", patrolController.All)
		patrolRoutes.GET("/:id", patrolController.FindById)
		patrolRoutes.POST("/", patrolController.Create)
		patrolRoutes.PUT("/", patrolController.Update)
		patrolRoutes.DELETE("/:id", patrolController.Remove)

	}
	r.Static("assets", "./assets")
	r.Run(":8080")
	return
}
