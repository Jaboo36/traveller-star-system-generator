package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func parseBool(string string) (bool, error) {
	switch strings.ToLower(string) {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean: %q", string)
	}
}

func continueOrStop(scanner *bufio.Scanner) bool {
	var continueOrStop bool

	continueLoop := true
	for continueLoop {
		scanner.Scan()

		answer := scanner.Text()

		boolValue, err := parseBool(answer)

		if err != nil {
			fmt.Printf("%q is not a valid answer, please try again...\n", answer)
		} else {
			continueOrStop = boolValue
			continueLoop = false

		}
	}
	return continueOrStop
}

func roll2D() uint8 {
	roll1 := uint8(rand.Intn(6) + 1)
	roll2 := uint8(rand.Intn(6) + 1)

	return roll1 + roll2
}

type planetSize struct {
	Id      uint8
	size    string
	example string
	gravity float32
}

func determinePlanetSize() planetSize {
	sizes := map[uint8]planetSize{
		0:  {0, "Less than 1,600km", "Asteroid, orbital complex", 0.0},
		1:  {1, "1,600km", "Triton", 0.05},
		2:  {2, "3,200km", "Luna, Europa", 0.15},
		3:  {3, "4,800km", "Mercury, Ganymede", 0.25},
		4:  {4, "6,400km", "Mars", 0.35},
		5:  {5, "8,000km", "", 0.45},
		6:  {6, "9,600km", "", 0.7},
		7:  {7, "11,200km", "", 0.9},
		8:  {8, "12,800km", "Earth", 1.0},
		9:  {9, "14,400km", "", 1.25},
		10: {10, "16,00km", "", 1.4},
	}

	result := roll2D() - 2

	return sizes[result]
}

type atmosphere struct {
	Id                   uint8
	Composition          string
	Examples             string
	Pressure             string
	SurvivalGearRequired string
}

func determineAtmosphere(planetSize uint8) atmosphere {
	atmosphereTypes := map[uint8]atmosphere{
		0:  {0x0, "None", "Moon", "0.00", "Vacc Suit"},
		1:  {0x1, "Trace", "Mars", "0.001 to 0.09", "Vacc Suit"},
		2:  {0x2, "Very Thin, Tainted", "", "0.1 to 0.42", "Respirator, Filter"},
		3:  {0x3, "Very Thin", "", "0.1 to 0.42", "Respirator"},
		4:  {0x4, "Thin, Tainted", "", "0.43 to 0.7", "Filter"},
		5:  {0x5, "Thin", "", "0.43 to 0.7", ""},
		6:  {0x6, "Standard", "Earth", "0.71 to 1.49", ""},
		7:  {0x7, "Standard, Tainted", "", "0.71 to 1.49", "Filter"},
		8:  {0x8, "Dense", "", "1.5 to 2.49", ""},
		9:  {0x9, "Dense, Tainted", "", "1.5 to 2.49", "Filter"},
		10: {0xA, "Exotic", "", "Varies", "Air Supply"},
		11: {0xB, "Corosive", "Venus", "Varies", "Vacc Suit"},
		12: {0xC, "Insidious", "", "Varies", "Vacc Suit"},
		13: {0xD, "Very Dense", "", "2.5+", ""},
		14: {0xE, "Low", "", "0.5 or less", ""},
		15: {0xF, "Unusual", "", "Varies", "Varies"},
	}

	result := roll2D() - 7 + planetSize

	return atmosphereTypes[result]
}

type hydrographics struct {
	Id          uint8
	Percentage  string
	Description string
}

func determineHydrographics(atmosphere uint8) hydrographics {
	hydrographicsProfiles := map[uint8]hydrographics{
		0:  {0, "0%-5%", "Desert world"},
		1:  {1, "6%-15%", "Dry world"},
		2:  {2, "16%-25%", "A few small seas"},
		3:  {3, "26%-35%", "Small seas and oceans"},
		4:  {4, "36%-45%", "Wet world"},
		5:  {5, "46%-55%", "A large ocean"},
		6:  {6, "56%-65%", "Large oceans"},
		7:  {7, "66%-75%", "Earth-like world"},
		8:  {8, "76%-85%", "Only a few islands and archipelagos"},
		9:  {9, "86%-95%", "Almost entirely water"},
		10: {0xF, "96%-100%", "Waterworld"},
	}

	result := roll2D() - 7 + atmosphere

	return hydrographicsProfiles[result]
}

func main() {
	cont := true

	for cont {
		planetSize := determinePlanetSize()
		atmosphere := determineAtmosphere(planetSize.Id)
		Hydrographics := determineHydrographics(atmosphere.Id)

		fmt.Println("Would you like to continue? yes/no?")

		scanner := bufio.NewScanner(os.Stdin)

		cont = continueOrStop(scanner)
	}
}
