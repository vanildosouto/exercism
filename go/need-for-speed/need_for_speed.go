package speed

type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{100, batteryDrain, speed, 0}
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	batteryLeft := car.battery - car.batteryDrain
	if batteryLeft < 0 {
		return car
	}

	car.battery = batteryLeft
	car.distance += car.speed
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	lapsBatteryRemains := car.battery / car.batteryDrain
	lapsRemains := track.distance / car.speed

	return lapsBatteryRemains > lapsRemains
}
