package main

import (
	"fmt"
	"math/rand"
	"github.com/jmcvetta/randutil"
)

// Variables, can be customized

min_num_of_traits := 3
max_num_of_traits := 6
chance_of_genetic := 80 # in percent
max_num_of_genetic := 1
min_age := 16
max_age := 30
min_stat := 9 // For cost calculations, we calculate base stats between 0 and 20
max_stat := 12
var education_level_weights = []int{15,30,30,10,5}
var genetics_level_weights = []int{1,2,4,4,2,1}
var sex_orientation_weights = []int{15,3,3,3}


// Game Data
var stats_cost = [int]{0,2,4,6,8,12,16,20,24,31,38,45,52,63,74,85,96,113,130,147,164}
var age_cost = [int]{0,2,4,6,8,10,12,14,16,18,22,24,27,29,31,33,40,42,48,51,58,60,66,66,67,67,67,67,67,66,66,65,64,62,61,59,57,55,53,50,48,45,42,38,35,31,27,23,19,14,10,10,10,10,10,11,11,11,11,11,6,6,6,6,6,6,6,6,6,6,0}
var education_cost = [int]{0,20,40,80,150}

educations := [][]string{
	[]string{"Naive Appeaser", "Adequate Bargainer", "Charismatic Negotiator", "Grey Eminence","Virtuoso Arbitrator"},
	[]string{"Misguided Warrior", "Tough Soldier", "Skilled Tactician", "Brilliant Strategist","Exalted Warlord"},
	[]string{"Indulgent Wastrel", "Thrifty Clerk", "Fortune Builder", "Midas Touched","Golden Sovereign"},
	[]string{"Amateurish Plotter", "Flamboyant Trickster", "Intricate Webweaver", "Elusive Shadow","Conniving Puppetmaster"},
	[]string{"Conscientious Scribe", "Insightful Thinker", "Astute Intellectual", "Mastermind Philosopher","Erudite Oracle"},
} 

traits :=[][]map[string]int{

	{"Chaste":20, "Lustful":25}
	{"Temperate":40, "Gluttonous":20}
	{"Generous":20, "Greedy":30}
	{"Diligent":40, "Lazy":-10}
	{"Patient":30, "Impatient":25}
	{"Humble":20, "Arrogant":20}
	{"Ambitious":40, "Content":20}
	{"Just":40, "Arbitrary":30}
	{"Brave":40, "Craven":-10}
	{"Calm":25, "Wrathful":30}
	{"Forgiving":25, "Vengeful":30}
	{"Honest":20, "Deceitful":30}
	{"Zealous":30, "Cynical":30}
	{"Gregarious":30, "Shy":-10}
	{"Trusting":10, "Paranoid":-10}
	{"Compassionate":10, "Callous":40, "Sadistic":40}
	{"Fickle":25, "Stubborn":30, "Eccentric":15}
}

genetics := [][]map[string]int{

	{"Hideous":-30, "Ugly":-20, "Homely":-10, "Comely":40, "Pretty / Handsome":80, "Beautiful":120}
	{"Imbecile":-45, "Stupid":-30, "Slow":-15, "Quick":80, "Intelligent":160, "Genius":240}
	{"Feeble":-45, "Frail":-30, "Delicate":-15, "Hale":60, "Robust":120, "Amazonian / Herculean":180}

}

stats := []string{"Diplomacy", "Martial", "Stewardship", "Intrigue", "Learning","Prowess"}

//Character Creation
func Creation(educ int) {
    cost := 0

	// Age and sex

	sex := []string{"Male", "Female"}[rand.Intn(2)]

	sex_oris := []randutil.Choice{
		randutil.Choice{5, "Hetero"},
		randutil.Choice{1, "Homo"},
		randutil.Choice{1, "Bi"},
		randutil.Choice{1, "Ace"},
	}

	sex_ori_ret, err := randutil.WeightedChoice(sex_oris)
		if err != nil {
			panic(err)
		}
	sex_orient := sex_ori_ret.Item
	age := rand.Intn(71)
	if age > 70 { 
        cost = append(cost, 0)
    } else {
		cost = append(cost, age_cost[age])
	}

	// Education

	// Traits

	// Genetics

	// Stats

	// Printing and Formating


}