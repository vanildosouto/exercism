package lasagna

const (
	// OvenTime constant with how many minutes the lasagna should be in the oven
	OvenTime = 40
)

// RemainingOvenTime take the actual minutes the lasagna has been in the oven
// and returns how many minutes the lasagna still has to remain
func RemainingOvenTime(time int) int {
	return OvenTime - time
}

// PreparationTime function that takes the number of layers you added to the lasagna as a parameter and returns how many minutes you spent preparing the lasagna, assuming each layer takes you 2 minutes to prepare.
func PreparationTime(layers int) int {
	return layers * 2
}

// ElapsedTime The function should return how many minutes in total you've worked on cooking the lasagna, which is the sum of the preparation time in minutes, and the time in minutes the lasagna has spent in the oven at the moment
func ElapsedTime(layers int, minutes int) int {
	return PreparationTime(layers) + minutes
}
