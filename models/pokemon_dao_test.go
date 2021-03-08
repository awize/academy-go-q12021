package models

import (
	"fmt"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	pokemon, err := GetPokemon(1)
	emptyPokemon := Pokemon{}
	fmt.Println(err)
	if pokemon == emptyPokemon {
		t.Error("expected a pokemon")
	}

	if err != nil {
		t.Error("we did not expect any error")
	}
}
