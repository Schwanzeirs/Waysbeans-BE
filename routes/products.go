package routes

import (
	"waysbeans/handlers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	ProductRepository := repositories.RepositoryProduct(mysql.DB)
	h := handlers.HandlerProduct(ProductRepository)

	r.HandleFunc("/products", h.FindProducts).Methods("GET")
	r.HandleFunc("/product/{id}", h.GetProduct).Methods("GET")
	r.HandleFunc("/product", middleware.UploadFile(h.CreateProduct)).Methods("POST")
	r.HandleFunc("/product/{id}", middleware.UploadFile(h.UpdateProduct)).Methods("PATCH")
	r.HandleFunc("/product/{id}", h.DeleteProduct).Methods("DELETE")
}
