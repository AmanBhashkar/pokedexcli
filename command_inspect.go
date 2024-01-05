package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	// pockeapiClient := pokeapi.NewClient()
	if (len(args)) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("%s hasn't bee caught", pokemonName)
	}
	fmt.Printf("Name %s\n", pokemon.Name)
	fmt.Printf("Height%v\n", pokemon.Height)
	fmt.Printf("Weight%v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Type:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}
	return nil
}
