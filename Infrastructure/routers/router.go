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
	rolController           controllers.RolController           = controllers.NewRolController()
	detachmentController    controllers.DetachmentController    = controllers.NewDetachmentController()
	churchController        controllers.ChurchController        = controllers.NewChurchController()
	moduleController        controllers.ModuleController        = controllers.NewModuleController()
	subDetachmentController controllers.SubdetachmentController = controllers.NewSubdetachmentController()
	patrolController        controllers.PatrolController        = controllers.NewPatrolController()
	kanbanController        controllers.KanbanController        = controllers.NewKanbanController()
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
	authRoutes := r.Group("api/login")
	{
		authRoutes.POST("/", authController.Login)
		//authRoutes.POST("/register", authController.Register)
	}
	//userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
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
	kanbanRoutes := r.Group("api/kanban", middleware.AuthorizeJWT(jwtService))
	{
		kanbanRoutes.GET("/kanban", kanbanController.GetKanban)
		kanbanRoutes.GET("/count_kanban", kanbanController.CountKanban)
	}
	rolRoutes := r.Group("api/rol", middleware.AuthorizeJWT(jwtService))
	{
		rolRoutes.GET("/", rolController.All)
		rolRoutes.POST("/create", rolController.Create)
		rolRoutes.PUT("/", rolController.Update)
		rolRoutes.DELETE("/:id", rolController.Delete)
		rolRoutes.GET("/:id", rolController.FindRol)
	}
	// middleware.AuthorizeJWT(jwtService)
	detachmentRoutes := r.Group("api/detachment", middleware.AuthorizeJWT(jwtService))
	{
		detachmentRoutes.GET("/", detachmentController.All)
		detachmentRoutes.GET("/:id", detachmentController.FindById)
		detachmentRoutes.POST("/", detachmentController.Create)
		detachmentRoutes.PUT("/", detachmentController.Update)
		detachmentRoutes.DELETE("/:id", detachmentController.Delete)

	}
	churchRoutes := r.Group("api/church", middleware.AuthorizeJWT(jwtService))
	{
		churchRoutes.GET("/", churchController.All)
		churchRoutes.GET("/:id", churchController.FindById)
		churchRoutes.POST("/", churchController.Create)
		churchRoutes.PUT("/", churchController.Update)
		churchRoutes.DELETE("/:id", churchController.Delete)

	}
	moduleRoutes := r.Group("api/module", middleware.AuthorizeJWT(jwtService))
	{
		moduleRoutes.GET("/", moduleController.All)
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
