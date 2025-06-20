package pokemons

type Pokedex struct {
	pokemons map[string]Pokemon
}

type Pokemon struct {
	Name      string
	Height    int
	Weight    int
	Abilities []struct {
		IsHidden bool
		Slot     int
		Ability  struct {
			Name string
			URL  string
		}
	}
	Moves []struct {
		Move struct {
			Name string
			URL  string
		}
	}
	Species struct {
		Name string
		URL  string
	}
	Stats []struct {
		BaseStat int
		Effort   int
		Stat     struct {
			Name string
			URL  string
		}
	}
	Types []struct {
		Slot int
		Type struct {
			Name string
			URL  string
		}
	}
}
