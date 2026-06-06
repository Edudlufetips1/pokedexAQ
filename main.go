package main

import (
    "time"
    "github.com/Edudlufetips1/pokedexAQ/internal/pokeapi"
)

func main() {
    client := pokeapi.NewClient(5*time.Second, 5*time.Minute)
    cfg := &config{
        pokeapiClient: client,
    }
	cfg.caughtPokemon = make(map[string]pokeapi.Pokemon)
    startRepl(cfg)
}