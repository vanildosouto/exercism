package lasagna

// PreparationTime return estimate total preparation
func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}

	return len(layers) * preparationTime
}

// Quantities return quantities of noodls and sauce
func Quantities(layers []string) (int, float64) {
	saucePerLayer := 0.2  // in liters
	noodlesPerLayer := 50 // in grams
	var sauceQuantities, noodlesQuantities int

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesQuantities++
		}

		if layer == "sauce" {
			sauceQuantities++
		}
	}

	return noodlesPerLayer * noodlesQuantities, saucePerLayer * float64(sauceQuantities)
}

// AddSecretIngredient return my list with a secret ingredient
func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// ScaleRecipe return amounts needed for the desired number of portions
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaleQuantities []float64
	onePortion := float64(portions) / 2.0

	for _, quantity := range quantities {
		scaleQuantities = append(scaleQuantities, quantity*onePortion)
	}

	return scaleQuantities
}
