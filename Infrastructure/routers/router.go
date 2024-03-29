package routers

import (
	"bete/Infrastructure/middleware"
	"bete/UseCases/InterfacesService"
	"bete/UseCases/services"
	"bete/controllers"

	docs "bete/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// controller
var (
	jwtService InterfacesService.IJWTService = services.NewJWTService()

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
	cityController          controllers.CityController          = controllers.NewCityController()
	parentController        controllers.ParentController        = controllers.NewParentController()
	visitController         controllers.VisitController         = controllers.NewVisitController()
)

func NewRouter() {

	r := gin.New()
	docs.SwaggerInfo_swagger.BasePath = "/api/v1"
	// Set up CORS middleware options
	// config := cors.Config{
	// 	Origins:         "*",
	// 	RequestHeaders:  "Authorization",
	// 	Methods:         "GET, POST, PUT,DELETE",
	// 	Credentials:     true,
	// 	ValidateHeaders: false,
	// 	MaxAge:          1 * time.Minute,
	// }

	// Apply the middleware to the router (works on groups too)
	//r.Use(cors.Middleware(config))

	//
	authRoutes := r.Group("api/v1/login")
	{
		authRoutes.POST("/", authController.Login)
		//authRoutes.POST("/register", authController.Register)
	}
	userRoutes := r.Group("api/v1/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/", userController.ListUser)
		userRoutes.GET("/:id", userController.FindUser)
		userRoutes.GET("search/:search", userController.FindUserNameLastName)
		userRoutes.POST("/", userController.Create)
		userRoutes.PUT("/:id", userController.Update)
		userRoutes.PUT("/update_password", userController.PasswordChange)
		userRoutes.PUT("/", userController.UpdateProfile)
		userRoutes.DELETE("/:id", userController.Delete)
	}

	//userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	scoutRoutes := r.Group("api/v1/scout", middleware.AuthorizeJWT(jwtService))
	{
		scoutRoutes.GET("/", scoutController.ListKingsScouts)
		scoutRoutes.POST("/create", scoutController.Create)
		scoutRoutes.PUT("/edit", scoutController.Update)
		//scoutRoutes.DELETE("/:id", scoutController.Delete)
	}

	attendanceRoutes := r.Group("api/v1/attendance", middleware.AuthorizeJWT(jwtService))
	{
		attendanceRoutes.GET("/", attendanceController.All)
		attendanceRoutes.GET("/:id", attendanceController.AttendancesSubdetachment)
		attendanceRoutes.GET("/weeks", attendanceController.WeeksbySubDetachments)
		attendanceRoutes.POST("/", attendanceController.Create)
		attendanceRoutes.PUT("/:id", attendanceController.Update)
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
		rolRoutes.POST("/", rolController.Create)
		rolRoutes.PUT("/:id", rolController.Update)
		rolRoutes.DELETE("/:id", rolController.Remove)
		rolRoutes.GET("/:id", rolController.FindRol)
	}

	detachmentRoutes := r.Group("api/v1/detachment", middleware.AuthorizeJWT(jwtService))
	{
		detachmentRoutes.GET("/", detachmentController.All)
		detachmentRoutes.GET("/:id", detachmentController.FindById)
		detachmentRoutes.POST("/", detachmentController.Create)
		detachmentRoutes.PUT("/:id", detachmentController.Update)
		detachmentRoutes.DELETE("/:id", detachmentController.Delete)

	}
	churchRoutes := r.Group("api/v1/church", middleware.AuthorizeJWT(jwtService))
	{
		churchRoutes.GET("/", churchController.All)
		churchRoutes.GET("/:id", churchController.FindById)
		churchRoutes.POST("/", churchController.Create)
		churchRoutes.PUT("/:id", churchController.Update)
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
	patrolRoutes := r.Group("api/v1/patrol")
	{
		patrolRoutes.GET("/", patrolController.All)
		patrolRoutes.GET("/:id", patrolController.FindById)
		patrolRoutes.POST("/", patrolController.Create)
		patrolRoutes.PUT("/:id", patrolController.Update)
		patrolRoutes.DELETE("/:id", patrolController.Remove)

	}
	parentRoutes := r.Group("api/v1/parent", middleware.AuthorizeJWT(jwtService))
	{
		parentRoutes.GET("/", parentController.All)
		parentRoutes.GET("/:id", parentController.AllParentScout)
		parentRoutes.POST("/:id", parentController.Create)
		parentRoutes.PUT("/", parentController.Update)
		parentRoutes.DELETE("/:id", parentController.Remove)
	}
	visitRoutes := r.Group("api/v1/visit", middleware.AuthorizeJWT(jwtService))
	{
		visitRoutes.GET("/", visitController.All)
		visitRoutes.GET("/:id", visitController.AllVisitByUserAndSubDatachment)
		visitRoutes.POST("/", visitController.CreateVisit)
		visitRoutes.PUT("/", visitController.UpdateVisit)
		visitRoutes.DELETE("/:id", visitController.RemoveVisit)
	}
	cityRoutes := r.Group("api/v1/city", middleware.AuthorizeJWT(jwtService))
	{
		cityRoutes.GET("/", cityController.All)
	}
	r.Static("assets", "./assets")

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	r.Run(":8080")
	return
}
