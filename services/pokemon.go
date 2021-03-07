package services

import "github.com/awize/pokeapi/models"

// GetPokemon handler
func GetPokemon(pokemonID int) (models.Pokemon, error) {
	return models.GetPokemon(pokemonID)
}

// GetPokemons handler
func GetPokemons() (map[int]models.Pokemon, error) {
	return models.GetPokemons()
}
