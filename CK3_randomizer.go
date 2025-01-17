package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/mroth/weightedrand/v2"
)

// Variables, can be customized

var min_num_of_traits = 3
var max_num_of_traits = 6
var chance_of_genetic = 80 // in percent
var max_num_of_genetic = 1
var min_age = 16
var max_age = 30
var min_stat = 9 // For cost calculations, we calculate base stats between 0 and 20
var max_stat = 12

// Game Data
var stats_cost = []int{0, 2, 4, 6, 8, 12, 16, 20, 24, 31, 38, 45, 52, 63, 74, 85, 96, 113, 130, 147, 164}
var age_cost = []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 22, 24, 27, 29, 31, 33, 40, 42, 48, 51, 58, 60, 66, 66, 67, 67, 67, 67, 67, 66, 66, 65, 64, 62, 61, 59, 57, 55, 53, 50, 48, 45, 42, 38, 35, 31, 27, 23, 19, 14, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 0}
var education_cost = []int{0, 20, 40, 80, 150}

var educations = [][]string{
	[]string{"Naive Appeaser (Diplomacy 1)", "Adequate Bargainer (Diplomacy 2)", "Charismatic Negotiator (Diplomacy 3)", "Grey Eminence (Diplomacy 4)", "Virtuoso Arbitrator (Diplomacy 5)"},
	[]string{"Misguided Warrior (Martial 1)", "Tough Soldier (Martial 2)", "Skilled Tactician (Martial 3)", "Brilliant Strategist (Martial 4)", "Exalted Warlord (Martial 5)"},
	[]string{"Indulgent Wastrel (Stewardship 1)", "Thrifty Clerk (Stewardship 2)", "Fortune Builder (Stewardship 3)", "Midas Touched (Stewardship 4)", "Golden Sovereign (Stewardship 5)"},
	[]string{"Amateurish Plotter (Intrigue 1)", "Flamboyant Trickster (Intrigue 2)", "Intricate Webweaver (Intrigue 3)", "Elusive Shadow (Intrigue 4)", "Conniving Puppetmaster (Intrigue 5)"},
	[]string{"Conscientious Scribe (Learning 1)", "Insightful Thinker (Learning 2)", "Astute Intellectual (Learning 3)", "Mastermind Philosopher (Learning 4)", "Erudite Oracle (Learning 5)"},
}

type Traits struct {
	trait string
	cout  int
}

var liste_traits = [][]Traits{

	{Traits{"Chaste", 20}, Traits{"Lustful", 25}},
	{Traits{"Temperate", 40}, Traits{"Gluttonous", 20}},
	{Traits{"Generous", 20}, Traits{"Greedy", 30}},
	{Traits{"Diligent", 40}, Traits{"Lazy", -10}},
	{Traits{"Patient", 30}, Traits{"Impatient", 25}},
	{Traits{"Humble", 20}, Traits{"Arrogant", 20}},
	{Traits{"Ambitious", 40}, Traits{"Content", 20}},
	{Traits{"Just", 40}, Traits{"Arbitrary", 30}},
	{Traits{"Brave", 40}, Traits{"Craven", -10}},
	{Traits{"Calm", 25}, Traits{"Wrathful", 30}},
	{Traits{"Forgiving", 25}, Traits{"Vengeful", 30}},
	{Traits{"Honest", 20}, Traits{"Deceitful", 30}},
	{Traits{"Zealous", 30}, Traits{"Cynical", 30}},
	{Traits{"Gregarious", 30}, Traits{"Shy", -10}},
	{Traits{"Trusting", 10}, Traits{"Paranoid", -10}},
	{Traits{"Compassionate", 10}, Traits{"Callous", 40}},
	{Traits{"Fickle", 25}, Traits{"Stubborn", 30}},
}

var genetic_type = []string{"beauty", "intel", "phys"}

var beauty_gens, _ = weightedrand.NewChooser(
	weightedrand.NewChoice(Traits{"Hideous", -30}, 1),
	weightedrand.NewChoice(Traits{"Ugly", -20}, 2),
	weightedrand.NewChoice(Traits{"Homely", -10}, 4),
	weightedrand.NewChoice(Traits{"Comely", 40}, 4),
	weightedrand.NewChoice(Traits{"Pretty / Handsome", 80}, 2),
	weightedrand.NewChoice(Traits{"Beautiful", 120}, 1),
)

var intel_gens, _ = weightedrand.NewChooser(
	weightedrand.NewChoice(Traits{"Imbecile", -45}, 1),
	weightedrand.NewChoice(Traits{"Stupid", -30}, 2),
	weightedrand.NewChoice(Traits{"Slow", -15}, 4),
	weightedrand.NewChoice(Traits{"Quick", 80}, 4),
	weightedrand.NewChoice(Traits{"Intelligent", 160}, 2),
	weightedrand.NewChoice(Traits{"Genius", 240}, 1),
)

var phys_gens, _ = weightedrand.NewChooser(
	weightedrand.NewChoice(Traits{"Feeble", -45}, 1),
	weightedrand.NewChoice(Traits{"Frail", -30}, 2),
	weightedrand.NewChoice(Traits{"Delicate", -15}, 4),
	weightedrand.NewChoice(Traits{"Hale", 60}, 4),
	weightedrand.NewChoice(Traits{"Robust", 120}, 2),
	weightedrand.NewChoice(Traits{"Amazonian / Herculean", 180}, 1),
)

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func slice_remove(s [][]Traits, i int) [][]Traits {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func slice_remove_str(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Character Creation
func Creation(educ int) {
	cost := []int{}

	// Age and sex

	sex := []string{"Male", "Female"}[rand.Intn(2)]

	sex_oris, _ := weightedrand.NewChooser(
		weightedrand.NewChoice("Heterosexual", 5),
		weightedrand.NewChoice("Homosexual", 1),
		weightedrand.NewChoice("Bisexual", 1),
		weightedrand.NewChoice("Asexual", 1),
	)
	sex_orient := sex_oris.Pick()

	age := randRange(min_age, max_age)
	if age < 70 {
		cost = append(cost, age_cost[age])
	}

	// Education

	education_levels, _ := weightedrand.NewChooser(
		// CUSTOMIZABLE: weightedrand.NewChoice{EDUCATION LEVEL, RELATIVE FREQUENCY}
		weightedrand.NewChoice(0, 3),
		weightedrand.NewChoice(1, 6),
		weightedrand.NewChoice(2, 6),
		weightedrand.NewChoice(3, 2),
		weightedrand.NewChoice(4, 1),
	)
	education_trait := randRange(0, 4)
	education_final := educations[education_trait][education_levels.Pick()]

	// Traits

	num_of_traits := randRange(min_num_of_traits, max_num_of_traits)
	traits_char := []string{}

	for i := 0; i < num_of_traits; i++ {
		pick := randRange(0, len(liste_traits)-1)
		duo := liste_traits[pick] // can be a trio, changes nothing
		liste_traits = slice_remove(liste_traits, pick)
		chosen := duo[randRange(0, len(duo)-1)]
		traits_char = append(traits_char, chosen.trait)
		cost = append(cost, chosen.cout)
	}

	// Genetics
	genetics := []string{}

	for i := 0; i < max_num_of_genetic; i++ {
		gen := randRange(1, 100)
		if gen <= chance_of_genetic {
			genpick := randRange(0, len(genetic_type)-1)
			category := genetic_type[genpick]
			switch category {
			case "beauty":
				gentraits := beauty_gens.Pick()
				genetics = append(genetics, gentraits.trait)
				cost = append(cost, gentraits.cout)
			case "intel":
				gentraits := intel_gens.Pick()
				genetics = append(genetics, gentraits.trait)
				cost = append(cost, gentraits.cout)
			case "phys":
				gentraits := phys_gens.Pick()
				genetics = append(genetics, gentraits.trait)
				cost = append(cost, gentraits.cout)
			}
			genetic_type = slice_remove_str(genetic_type, genpick)
		}

	}

	// Stats

	stats_character := []int{}
	for i := 0; i < 7; i++ {
		stat_value := randRange(min_stat, max_stat)
		stats_character = append(stats_character, stat_value)
		cost = append(cost, stats_cost[stat_value])
	}

	// Printing and Formating
	cout := 0
	for _, i := range cost {
		cout += i
	}

	fmt.Println("Age:", age)
	fmt.Println("Sex:", sex)
	fmt.Println("Sexual Orientation:", sex_orient)
	fmt.Println("Education:", education_final)
	fmt.Println("Traits:", strings.Join(traits_char, ", "))
	fmt.Println("Genetic Traits:", strings.Join(genetics, ", "))
	fmt.Printf("Statistics:\n \tDipl: %v\n \tStew: %v\n \tMart: %v\n \tIntr: %v\n \tLear: %v\n \tProw: %v\n", stats_character[0], stats_character[1], stats_character[2], stats_character[3], stats_character[4], stats_character[5])
	fmt.Println("Cost:", cout)
}

func main() {
	Creation(0)
	fmt.Scanf("h")
}
