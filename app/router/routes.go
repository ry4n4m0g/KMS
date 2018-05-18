package router

func (r *Router) appendRoutes() {
	api := r.e.Group("/api")
	api.Use(r.mwBasicAuth())
}
