package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime_(layer []string, du int) int {
	if du == 0 {
		du = 2
	}
	return len(layer) * du
}

// TODO: define the 'Quantities()' function
func Quantities(layer []string) (nn int, sn float64) {
	/*
		Quantities([]string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"})
		// => 100, 0.4
	*/
	for _, item := range layer {
		switch item {
		case "noodles":
			nn += 50
		case "sauce":
			sn += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fl, ml []string) {
	ml[len(ml)-1] = fl[len(fl)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(fl []float64, c int) []float64 {
	result := []float64{}
	for _, item := range fl {
		result = append(result, item*float64(c)/2)
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
