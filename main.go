package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int){
  fmt.Println("Pilot, you have", fuel, "units of fuel remaining.")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int{
  var fuel int
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    case "Jupiter":
      fuel = 1000000
    case "Saturn":
      fuel = 1200000
    case "Earth":
      fuel = 0 // No fuel needed to stay on Earth
    default:
      fuel = 0
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("Welcome to", planet + "! We hope you enjoy your stay.")
}


// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else if fuelCost > fuelRemaining {
		cantFly()
	}

	return fuelRemaining
}

// Create the function returnToHomePlanet()
func returnToHomePlanet(fuel int) int {
	homePlanet := "Earth"
	fmt.Println("Returning to", homePlanet + "...")
	// Assume you have enough fuel to return to Earth
	fuel = flyToPlanet(homePlanet, fuel)
	fmt.Println("Arrived back at", homePlanet + "!")
	return fuel
}

func main() {
  // Create `planetChoice` and `fuel`
  fuel := 1000000
  planetChoice := "Venus"

  // Check initial fuel status
  fuelGauge(fuel)

  // Check calculateFuel function
  requiredFuel := calculateFuel(planetChoice)
  fmt.Println("Fuel required to reach", planetChoice, "is", requiredFuel)

  // Optional: Greet Mars for fun
  greetPlanet("Mars")
  

  // Fly to the planet and update remaining fuel
  fuel = flyToPlanet(planetChoice, fuel)

  // Check your gas tank after flying
  fuelGauge(fuel)

  // And then liftoff!
  // Now let's add the feature to travel between multiple planets
  // Let's fly to Mars from Venus
  planetChoice = "Mars"
  requiredFuel = calculateFuel(planetChoice)
  fmt.Println("Fuel required to reach", planetChoice, "is", requiredFuel)

  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)

  // Let's return to Earth
  fuel = returnToHomePlanet(fuel)
  fuelGauge(fuel)
}
