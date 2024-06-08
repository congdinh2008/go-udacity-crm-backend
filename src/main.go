package main

import (
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"congdinh.com/crm/controllers"
	"congdinh.com/crm/docs" // Updated import path
	"congdinh.com/crm/services"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	router := mux.NewRouter()
	services := services.NewCustomerService()
	customerController := controllers.NewCustomerController(services)
	customerController.RegisterRoutes(router)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Redirect / to /swagger/index.html
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
	})

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "CRM API"
	docs.SwaggerInfo.Description = "This is a sample server CRM server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	log.Println("Server is running on port 8080")

	// Open http://localhost:8080/swagger/index.html in the default browser
	url := "http://localhost:8080/swagger/index.html"
	openBrowser(url)

	log.Fatal(http.ListenAndServe(":8080", router))
}
