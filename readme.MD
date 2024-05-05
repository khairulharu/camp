# Endpoint

     //ALL SUPPORT ROUTES
     e.POST("/login", login)

     secured.GET("/campsites", getAllCampsites)
     secured.GET("/campsites/:id", getCampsite)
     secured.PUT("/campsites/:id", updateCampsite)
     secured.DELETE("/campsites/:id", deleteCampsite)

     // Booking routes
     secured.POST("/bookings", createBooking)
     secured.GET("/bookings", getAllBookings)
     secured.GET("/bookings/:id", getBooking)
     secured.PUT("/bookings/:id", updateBooking)
     secured.DELETE("/bookings/:id", cancelBooking)

     // Review routes
     secured.POST("/reviews", addReview)
     secured.GET("/reviews", getAllReviews)
     secured.GET("/reviews/:id", getReview)
     secured.PUT("/reviews/:id", updateReview)
     secured.DELETE("/reviews/:id", deleteReview)

     // User routes
     secured.GET("/users/:id", getUser)
     secured.PUT("/users/:id", updateUser)

     // Admin routes
     secured.GET("/admin/campsites", getAllCampsitesAdmin)
     secured.GET("/admin/bookings", getAllBookingsAdmin)
     secured.GET("/admin/reviews", getAllReviewsAdmin)
     secured.GET("/admin/users", getAllUsers)