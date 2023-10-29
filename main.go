package main

import (
	"fmt"
)

func main() {
	// Create the Government Singleton instance.
	singletonInstance = newGovernmentSingleton()

	// Get the GovernmentInfo instance.
	government := singletonInstance.GetGovernmentInfo()

	// Print information about the government.
	fmt.Printf("Government Name: %s\n", government.Name)
	fmt.Printf("President: %s\n", government.President)
	fmt.Printf("Capital: %s\n", government.Capital)
	fmt.Printf("Established: %d\n", government.Established)
}
