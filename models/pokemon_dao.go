package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetPokemons get pokemons hash
func GetPokemons() (map[int]Pokemon, error) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/assets/poke_list.csv")
	if err != nil {
		return nil, fmt.Errorf("failed opening file because: %s", err.Error())
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("failed opening file because: %s", err.Error())
	}

	pokemons := make(map[int]Pokemon)
	for _, record := range records {
		var pokemon Pokemon
		var err error
		pokemon.ID, err = strconv.Atoi(record[0])
		pokemon.Name = record[1]
		pokemon.Type1 = record[2]
		pokemon.Type2 = record[3]
		pokemon.Total, err = strconv.Atoi(record[4])
		pokemon.HP, err = strconv.Atoi(record[5])
		pokemon.Attack, err = strconv.Atoi(record[6])
		pokemon.Defense, err = strconv.Atoi(record[7])
		pokemon.SpAttack, err = strconv.Atoi(record[8])
		pokemon.SpDefense, err = strconv.Atoi(record[9])
		pokemon.Speed, err = strconv.Atoi(record[10])
		pokemon.Generation, err = strconv.Atoi(record[11])
		pokemon.Legendary, err = strconv.ParseBool(strings.ToLower(record[12]))
		if err != nil {
			fmt.Println(err)
		} else {
			pokemons[pokemon.ID] = pokemon
		}

	}
	return pokemons, nil
}

// GetPokemon retrieves one pokemon
func GetPokemon(pokemonID int) (Pokemon, error) {
	pokemons, err := GetPokemons()
	if err != nil {
		return Pokemon{}, err
	}

	if pokemons == nil {
		return Pokemon{}, fmt.Errorf("There are no pokemons")
	}

	if pokemon, ok := pokemons[pokemonID]; ok {
		return pokemon, nil
	}
	return Pokemon{}, fmt.Errorf("pokemon %v", pokemonID)
}
