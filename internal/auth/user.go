package auth

// import (
// 	"campsite/internal/config"
// 	"net/http"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// JWT Secret Key
// var jwtSecret = []byte("campsite")

// // User struct
// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // JWT Claims
// type Claims struct {
// 	UserID int `json:"userId"`
// 	jwt.StandardClaims
// }

// // Login handler
// func login(c echo.Context) error {
// 	// Handle login logic
// 	// ...

// 	// Create JWT token
// 	claims := &Claims{
// 		UserID: user.ID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"token": tokenString,
// 	})
// }

// // JWT middleware
// func isAuthorized(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		token := c.Request().Header.Get("Authorization")
// 		if token == "" {
// 			return echo.ErrUnauthorized
// 		}

// 		claims := &Claims{}
// 		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtSecret, nil
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		// Set user ID to context
// 		c.Set("userId", claims.UserID)

// 		return next(c)
// 	}
// }

// func Start(config *config.Config, echo *echo.Echo) {

// 	// Middleware
// 	echo.Use(middleware.Logger())
// 	echo.Use(middleware.Recover())

// 	// Routes
// 	echo.POST("/login", login)

// 	// Secure group with JWT middleware
// 	secured := e.Group("")
// 	secured.Use(isAuthorized)
// 	{
// 		// Campsite routes
// 		secured.POST("/campsites", addCampsite)
// 		secured.GET("/campsites", getAllCampsites)
// 		secured.GET("/campsites/:id", getCampsite)
// 		secured.PUT("/campsites/:id", updateCampsite)
// 		secured.DELETE("/campsites/:id", deleteCampsite)

// 		// Booking routes
// 		secured.POST("/bookings", createBooking)
// 		secured.GET("/bookings", getAllBookings)
// 		secured.GET("/bookings/:id", getBooking)
// 		secured.PUT("/bookings/:id", updateBooking)
// 		secured.DELETE("/bookings/:id", cancelBooking)

// 		// Review routes
// 		secured.POST("/reviews", addReview)
// 		secured.GET("/reviews", getAllReviews)
// 		secured.GET("/reviews/:id", getReview)
// 		secured.PUT("/reviews/:id", updateReview)
// 		secured.DELETE("/reviews/:id", deleteReview)

// 		// User routes
// 		secured.GET("/users/:id", getUser)
// 		secured.PUT("/users/:id", updateUser)

// 		// Admin routes
// 		secured.GET("/admin/campsites", getAllCampsitesAdmin)
// 		secured.GET("/admin/bookings", getAllBookingsAdmin)
// 		secured.GET("/admin/reviews", getAllReviewsAdmin)
// 		secured.GET("/admin/users", getAllUsers)
// 	}

// 	e.Logger.Fatal(e.Start(":8000"))
// }
