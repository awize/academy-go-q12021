package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pokeapi/services"
)

// GetPokemons retrieves all pokemons
func GetPokemons(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// response with pokemons
	pokemons, err := services.GetPokemons()
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}

	pokemonsJSON, _ := json.Marshal(pokemons)
	resp.Write(pokemonsJSON)
}

// GetPokemon retrieves the pokemon by Id
func GetPokemon(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	pokemonID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("pokemon_id must be a number"))
		return
	}
	pokemon, err := services.GetPokemon(pokemonID)

	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		return
	}

	pokemonJSON, _ := json.Marshal(pokemon)
	resp.Write(pokemonJSON)
}
