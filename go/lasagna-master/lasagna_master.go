package lasagna

const (
	gramsOfNoodles = 50
	litersOfSauce  = 0.2
)

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}

	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodlesGrams int
	var sauceLiters float64

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesGrams += gramsOfNoodles
		}

		if layer == "sauce" {
			sauceLiters += litersOfSauce
		}
	}

	return noodlesGrams, sauceLiters
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var newQuantities []float64

	for _, quantity := range quantities {
		singlePortion := quantity / 2
		newQuantities = append(newQuantities, singlePortion*float64(portions))
	}

	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
