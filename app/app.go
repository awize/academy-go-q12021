package app

import (
	"log"
	"net/http"

	"github.com/awize/pokeapi/controllers"
	"github.com/julienschmidt/httprouter"
)

// StartApp give us the main router
func StartApp() {
	router := httprouter.New()
	router.GET("/", controllers.GetPokemons)
	router.GET("/:id", controllers.GetPokemon)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
