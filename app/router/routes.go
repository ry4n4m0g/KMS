package router

// appendRoutes registers routes in the router
func (r *Router) appendRoutes() {
	// API
	api := r.e.Group("/api")
	api.Use(r.mwBasicAuth())

	// API, Version 1
	v1 := api.Group("/v1")

	// Endpoints
	v1.GET("/key/generate_key", r.c.GenerateKey)
	v1.GET("/key/get_key", r.c.GetKey)
}
