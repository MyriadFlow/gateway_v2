package server

import (
	"os"

	"app.myriadflow.com/controllers"
	"app.myriadflow.com/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	if len(os.Getenv("GIN_MODE")) == 0 {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// adding middleware server
	router.Use(middleware.CORSMiddleware())

	// health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Routes(router)
	router.Run(":9090") // listen and serve on 0.0.0.0:808
}

func Routes(r *gin.Engine) {
	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users/all", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Brand routes
	r.POST("/brands", controllers.CreateBrand)
	r.GET("/brands/:id", controllers.GetBrand)
	r.GET("/brands/all", controllers.GetAllBrands)
	r.PUT("/brands/:id", controllers.UpdateBrand)
	r.DELETE("/brands/:id", controllers.DeleteBrand)
	r.GET("/brands/manager/:manager_id", controllers.GetBrandsByManager)


	// Collection routes
	r.POST("/collections", controllers.CreateCollection)
	r.GET("/collections/:id", controllers.GetCollection)
	r.GET("/collections/brand-id/:brandId", controllers.GetCollectionByBrandId)
	r.GET("/collections/all", controllers.GetAllCollections)
	r.PUT("/collections/:id", controllers.UpdateCollection)
	r.PUT("/collections/brand-id/:brandId", controllers.UpdateCollectionByBrandId)
	r.DELETE("/collections/:id", controllers.DeleteCollection)

	// Phygital routes
	r.POST("/phygitals", controllers.CreatePhygital)
	r.GET("/phygitals/:id", controllers.GetPhygital)
	r.GET("/phygitals/all", controllers.GetAllPhygital)
	r.PUT("/phygitals/:id", controllers.UpdatePhygital)
	r.DELETE("/phygitals/:id", controllers.DeletePhygital)

	// WebXR routes
	r.POST("/webxr", controllers.CreateWebXR)
	r.GET("/webxr/:id", controllers.GetWebXR)
	r.GET("/webxr/all", controllers.GetAllWebXR)
	r.PUT("/webxr/:id", controllers.UpdateWebXR)
	r.DELETE("/webxr/:id", controllers.DeleteWebXR)
	r.GET("/webxr/phygital/:phygital_id", controllers.GetWebXRByPhygitalID)

	// Avatar routes
	r.POST("/avatars", controllers.CreateAvatar)
	r.GET("/avatars/:id", controllers.GetAvatar)
	r.GET("/avatars/all", controllers.GetAllAvatars)
	r.PUT("/avatars/:id", controllers.UpdateAvatar)
	r.DELETE("/avatars/:id", controllers.DeleteAvatar)
	r.GET("/avatars/phygital/:phygital_id", controllers.GetAvatarByPhygitalID)

	// Variant routes
	r.POST("/variants", controllers.CreateVariant)
	r.GET("/variants/:id", controllers.GetVariant)
	r.GET("/variants/all", controllers.GetAllVariant)
	r.PUT("/variants/:id", controllers.UpdateVariant)
	r.DELETE("/variants/:id", controllers.DeleteVariant)
}
