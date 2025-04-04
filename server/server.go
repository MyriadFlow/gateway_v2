package server

import (
	"net/http"
	"os"

	"app.myriadflow.com/controllers"
	phygital_controllers "app.myriadflow.com/controllers/phygital"
	"app.myriadflow.com/db"
	"app.myriadflow.com/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	if os.Getenv("GIN_MODE") == "DEBUG" || os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// adding middleware server
	router.Use(middleware.CORSMiddleware())

	// health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/ping_database", func(c *gin.Context) {
		DB, err := db.Connect()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to connect to database: " + err.Error(),
			})
			return
		}

		// Retrieve the underlying SQL database object from GORM
		sqlDB, err := DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to retrieve database object: " + err.Error(),
			})
			return
		}

		// Ping the database to check if the connection is active
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database connection error: " + err.Error(),
			})
			return
		}

		// If the ping is successful
		c.JSON(http.StatusOK, gin.H{
			"message": "Database connection successful!",
		})
	})

	Routes(router)
	router.Run(":" + os.Getenv("APP_PORT")) // listen and serve on 0.0.0.0:808
}

func Routes(r *gin.Engine) {
	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users/all", controllers.GetAllUsers)
	r.GET("/users/all/:chaintype_id", controllers.GetAllUsersByChainType)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Brand routes
	r.POST("/brands", controllers.CreateBrand)
	r.GET("/brands/:id", controllers.GetBrand)
	r.GET("/brand/name/:name", controllers.GetBrandByName)
	r.GET("/brands/all/:chaintype_id", controllers.GetAllBrandsByChainType)
	r.GET("/brands/all", controllers.GetAllBrands)
	r.PUT("/brands/:id", controllers.UpdateBrand)
	r.DELETE("/brands/:id", controllers.DeleteBrand)
	r.GET("/brands/manager/:manager_id", controllers.GetBrandsByManager)
	r.GET("/brands/region", controllers.GetAllBrandsByRegion)

	// Collection routes
	r.POST("/collections", controllers.CreateCollection)
	r.GET("/collections/:id", controllers.GetCollection)
	r.GET("/collections/brand-id/:brandId", controllers.GetCollectionByBrandId)
	r.GET("/collections/all/:chaintype_id", controllers.GetAllCollectionsByChainType)
	r.GET("/collections/all", controllers.GetAllCollections)
	r.PUT("/collections/:id", controllers.UpdateCollection)
	r.PUT("/collections/brand-id/:brandId", controllers.UpdateCollectionByBrandId)
	r.DELETE("/collections/:id", controllers.DeleteCollection)
	r.GET("/collections/region", controllers.GetAllCollectionByRegion)

	// Phygital routes
	r.POST("/phygitals", phygital_controllers.CreatePhygital)
	r.GET("/phygitals/:id", phygital_controllers.GetPhygital)
	r.GET("/phygitals/deployer_address/:deployer_address", phygital_controllers.GetPhygitalByWalletAddress)
	r.GET("/phygitals/all/:chaintype_id", phygital_controllers.GetAllPhygitalByChainType)
	// r.GET("/phygitals/all", phygital_controllers.GetAllPhygital)
	r.PUT("/phygitals/:id", phygital_controllers.UpdatePhygital)
	r.DELETE("/phygitals/:id", phygital_controllers.DeletePhygital)
	r.GET("/phygitals/region", phygital_controllers.GetAllPhygitalByRegion)

	// WebXR routes
	r.POST("/webxr", controllers.CreateWebXR)
	r.GET("/webxr/:id", controllers.GetWebXR)
	r.GET("/webxr/all/:chaintype_id", controllers.GetAllWebXRByChainType)
	r.GET("/webxr/all", controllers.GetAllWebXR)
	r.PUT("/webxr/:id", controllers.UpdateWebXR)
	r.DELETE("/webxr/:id", controllers.DeleteWebXR)
	r.GET("/webxr/phygital/:phygital_id", controllers.GetWebXRByPhygitalID)
	r.GET("/webxr/region", controllers.GetAllWebxrByRegion)

	// Avatar routes
	r.POST("/avatars", controllers.CreateAvatar)
	r.GET("/avatars/:id", controllers.GetAvatar)
	r.GET("/avatars/all/:chaintype_id", controllers.GetAllAvatarsByChainType)
	r.GET("/avatars/all", controllers.GetAllAvatars)
	r.PUT("/avatars/:id", controllers.UpdateAvatar)
	r.DELETE("/avatars/:id", controllers.DeleteAvatar)
	r.GET("/avatars/phygital/:phygital_id", controllers.GetAvatarByPhygitalID)
	r.GET("/avatars/region", controllers.GetAllAvatarByRegion)

	// Variant routes
	r.POST("/variants", controllers.CreateVariant)
	r.GET("/variants/:id", controllers.GetVariant)
	r.GET("/variants/all/:chaintype_id", controllers.GetAllVariantByChainType)
	r.GET("/variants/all", controllers.GetAllVariant)
	r.PUT("/variants/:id", controllers.UpdateVariant)
	r.DELETE("/variants/:id", controllers.DeleteVariant)
	r.GET("/variants/region", controllers.GetAllVariantByRegion)

	//FanToken routes
	r.POST("/fantoken", controllers.CreateFanToken)
	r.GET("/fantoken/:id", controllers.GetFanToken)
	r.GET("fantoken/all/:chaintype_id", controllers.GetAllFanTokenByChainType)
	r.GET("fantoken/all", controllers.GetAllFanToken)
	r.PUT("fantoken/:id", controllers.UpdateFanToken)
	r.DELETE("fantoken/:id", controllers.DeleteFanToken)

	//ChainType routes
	r.POST("/chains", controllers.CreateChain)
	r.GET("/chains/:id", controllers.GetChain)
	r.PUT("/chains/:id", controllers.UpdateChain)
	r.DELETE("/chains/:id", controllers.DeleteChain)
	r.GET("/chains", controllers.GetChains)

	// NftEntries routes
	r.POST("/nftentries", controllers.CreateNftEntries)
	r.GET("/nftentries/:id", controllers.GetNftEntriesById)
	r.GET("/nftentries/phygital/:phygital_id", controllers.GetNftEntriesByPhygitalID)
	r.GET("/nftentries/all/:chaintype_id", controllers.GetAllNftEntriesByChainType)
	r.GET("/nftentries/owner/:phygital_id/:copy_number", controllers.GetOwnerByPhygitalAndCopyNumber)
	r.PUT("/nftentries/:id", controllers.UpdateNftEntries)
	r.DELETE("/nftentries/:id", controllers.DeleteNftEntries)

	//Profile routes
	r.POST("/profiles", controllers.CreateProfile)
	r.POST("/profiles/addresses/:profile_id", controllers.SaveAddresses)
	r.GET("/profiles/addresses/:profile_id", controllers.GetAddresses)
	r.GET("/profiles/:id", controllers.GetProfile)
	r.GET("/profiles/all", controllers.GetAllProfiles)
	r.GET("/profiles/all/:chaintype_id", controllers.GetAllProfilesByChainType)
	// r.GET("/profiles/email/:walletAddress", controllers.GetEmailByWalletAddress)
	r.GET("/profiles/wallet/:walletAddress", controllers.GetProfileByWalletAddress)
	r.GET("/profiles/username/:username", controllers.GetProfileByUsername)
	r.PUT("/profiles/:walletAddress", controllers.UpdateProfile)
	r.PUT("/profiles/addresses/:profile_id", controllers.UpdateAddress)
	r.DELETE("/profiles/:id", controllers.DeleteProfile)
	r.DELETE("/profiles/walletandemail/:walletAddress/:email", controllers.DeleteProfileByWalletAndEmail)
	r.DELETE("/profiles/addresses/:id", controllers.DeleteAddress)

	// Cart routes
	r.POST("/cart", controllers.AddToCart)
	r.DELETE("/cart/:wallet_address/:phygital_id", controllers.RemoveFromCart)
	r.GET("/cart/:wallet_address", controllers.GetCartItems)

	// OTP routes
	r.POST("/send-otp", controllers.SendOTPHandler)
	r.POST("/verify-otp", controllers.VerifyOTPHandler)

	// CreateFanToken routes
	r.POST("/create-fantoken", controllers.CreateMainnetFanTokenRequest)
	r.POST("/delegate-mint-fantoken", controllers.DelegateMintFanToken)
	r.POST("/create-mint-fantoken", controllers.CreateMintFanToken)
	r.GET("/get-mint-fantoken", controllers.GetAllMintFanToken)
	r.GET("/get-mint-fantoken/:creator_wallet", controllers.GetMintFanTokenByWalletAddress)
	r.PUT("/update-mint-fantoken/:id", controllers.UpdateMintFanToken)
	r.DELETE("/delete-mint-fantoken/:id", controllers.DeleteMintFanToken)

	//Elevate routes
	r.POST("/elevate", controllers.CreateElevate)
	r.GET("/elevate/:id", controllers.GetElevate)
	r.GET("/elevate/all", controllers.GetAllElevate)
	r.GET("/elevate/walletaddress/:walletAddress", controllers.GetElevateByWalletAddress)
	r.GET("/elevate/all/:chaintype_id", controllers.GetAllElevateByChainType)
	r.PUT("/elevate/:id", controllers.UpdateElevate)
	r.DELETE("/elevate/:id", controllers.DeleteElevate)

	r.POST("/agents", controllers.CreateAgent)
	r.GET("/agents", controllers.GetAgents)
	r.GET("/agents/:id", controllers.GetAgentByID)
	r.PUT("/agents/:brand_id", controllers.UpdateAgent)
	r.DELETE("/agents/:id", controllers.DeleteAgent)

	RoutesV2(r)

}
func RoutesV2(r *gin.Engine) {
	r.GET("v2/all/phygitals", phygital_controllers.GetAllPhygital)
}
