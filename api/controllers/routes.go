package controllers

import "penyuluhan2/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	// s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// // Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Admins routes
	s.Router.HandleFunc("/admin", middlewares.SetMiddlewareJSON(s.CreateAdmin)).Methods("POST")
	// s.Router.HandleFunc("/admin", middlewares.SetMiddlewareJSON(s.GetAdmins)).Methods("GET")
	s.Router.HandleFunc("/admin", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetAdmins))).Methods("GET")
	s.Router.HandleFunc("/admin/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/admin/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateAdmin))).Methods("PUT")
	s.Router.HandleFunc("/admin/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Penyuluh routes
	s.Router.HandleFunc("/penyuluh", middlewares.SetMiddlewareJSON(s.CreatePenyuluh)).Methods("POST")
	s.Router.HandleFunc("/penyuluh", middlewares.SetMiddlewareJSON(s.GetPenyuluhs)).Methods("GET")
	// s.Router.HandleFunc("/penyuluh", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPenyuluhs))).Methods("GET")
	s.Router.HandleFunc("/penyuluh/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/penyuluh/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/penyuluh/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	s.Router.HandleFunc("/changepassword/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.ChangePasswordUser))).Methods("PUT")

	//Majelis routes
	s.Router.HandleFunc("/majelis", middlewares.SetMiddlewareJSON(s.CreateMajelis)).Methods("POST")
	s.Router.HandleFunc("/majelis", middlewares.SetMiddlewareJSON(s.GetMajeliss)).Methods("GET")
	// s.Router.HandleFunc("/majelis", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetPenyuluhs))).Methods("GET")
	s.Router.HandleFunc("/majelis/{id}", middlewares.SetMiddlewareJSON(s.GetMajelis)).Methods("GET")
	s.Router.HandleFunc("/majelis/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/majelis/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Jadwal routes
	s.Router.HandleFunc("/jadwal", middlewares.SetMiddlewareJSON(s.CreateJadwal)).Methods("POST")
	s.Router.HandleFunc("/jadwal", middlewares.SetMiddlewareJSON(s.GetJadwals)).Methods("GET")
	// s.Router.HandleFunc("/jadwal", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetJadwals))).Methods("GET")
	s.Router.HandleFunc("/jadwal/{id}", middlewares.SetMiddlewareJSON(s.GetJadwal)).Methods("GET")
	s.Router.HandleFunc("/jadwal/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/jadwal/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//laporan routes
	s.Router.HandleFunc("/laporan", middlewares.SetMiddlewareJSON(s.CreateLaporan)).Methods("POST")
	// s.Router.HandleFunc("/laporan", middlewares.SetMiddlewareJSON(s.GetLaporans)).Methods("GET")
	// s.Router.HandleFunc("/laporan", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetLaporans))).Methods("GET")
	s.Router.HandleFunc("/laporan/{id}", middlewares.SetMiddlewareJSON(s.GetLaporan)).Methods("GET")
	s.Router.HandleFunc("/laporan/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/laporan/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// //Users routes
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	// s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetUsers))).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	// //Posts routes
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	// s.Router.HandleFunc("/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	// s.Router.HandleFunc("/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
