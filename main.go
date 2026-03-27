package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
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

func rollD() uint8 {
	return uint8(rand.Intn(6) + 1)
}

func roll2D() uint8 {
	return rollD() + rollD()
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

type planetTemperature struct {
	Type        string
	AverageTemp string
	Description string
}

func determineTemperature(atmosphere uint8) planetTemperature {
	diceMod := map[uint8]int8{
		0:   0,
		1:   0,
		2:   -2,
		3:   -2,
		4:   -1,
		5:   -1,
		0xE: -1,
		6:   0,
		7:   0,
		8:   1,
		9:   1,
		0xA: 2,
		0xD: 2,
		0xF: 2,
		0xB: 6,
		0xC: 6,
	}

	switch result := int8(roll2D()) + diceMod[atmosphere]; {
	case result <= 2:
		return planetTemperature{"Frozen", "-51 degrees or less", "Frozen world. No liquid water, very dry atmosphere."}
	case result < 5:
		return planetTemperature{"Cold", "-51 degrees to 0 degrees", "Icy world. Little liquid water, extensive ice caps, few clouds."}
	case result < 10:
		return planetTemperature{"Temperate", "0 degrees to 30 degrees", "Temperate world. Earth-like. Liquid & vaporized water are common, moderate ice caps."}
	case result < 12:
		return planetTemperature{"Hot", "31 degrees to 80 degrees", "Hot world. Small or no ice caps, little liquid water. Most water in the form of clouds."}
	default:
		return planetTemperature{"Boiling", "81 degrees or more", "Boiling world. No ice caps, little liquid water."}
	}
}

type population struct {
	Id          uint8
	Inhabitants string
	Range       string
	Description string
}

func determinePopulation() population {
	populations := map[uint8]population{
		0:   population{0, "None", "0", ""},
		1:   population{1, "Few", "1+", "A tiny farmstead or a single family"},
		2:   population{2, "Hundreds", "100+", "A village"},
		3:   population{3, "Thousands", "1,000+", ""},
		4:   population{4, "Tens of thousands", "10,000+", "Small town"},
		5:   population{5, "Hundreds of thousands", "100,000+", "Average city"},
		6:   population{6, "Millions", "1,000,000+", ""},
		7:   population{7, "Tens of millions", "10,000,000+", "Large city"},
		8:   population{8, "Hundreds of millions", "100,000,000+", ""},
		9:   population{9, "Billions", "1,000,000,000", "Present day Earth"},
		0xA: population{0xA, "Tens of billions", "10,000,000,000+", ""},
		0xB: population{0xB, "Hundreds of billions", "100,000,000,000", "Incredibly crowded world"},
		0xC: population{0xC, "Trillions", "1,000,000,000,000+", "World-city"},
	}

	return populations[roll2D()]
}

type government struct {
	Id                uint8
	governmentType    string
	description       string
	examples          string
	exampleContraband []string
}

func determineGovernment(population uint8) government {
	governments := map[uint8]government{
		0:   government{0, "None", "No government structure. In many cases, family bonds predominate.", "Family, clan, anarchy", []string{}},
		1:   government{1, "Company/Corporation", "Ruling functions are assumed by a company managerial elite and most citizenry are company employees or dependants.", "Corporate outpost, asteroid mine, feudal domain", []string{"Weapons", "Drugs", "Travellers"}},
		2:   government{2, "Participating Democracy", "Ruling functions are reached by the advice and consent of the citizenry directly.", "Collective, tribal council, comm-linked consensus", []string{"Drugs"}},
		3:   government{3, "Self-perpetuating Oligarchy", "Ruling functions are performed by a restricted minority, with little or no input from the mass of citizenry.", "Plutocracy, hereditary ruling caste", []string{"Technology", "Weapons", "Travllers"}},
		4:   government{4, "Representative Democracy", "Ruling functions are performed by elected representatives.", "Republic, democracy", []string{"Drugs", "Weapons", "Psionics"}},
		5:   government{5, "Feudal Technocracy", "Ruling functions are performed by specific individuals for persons who agree to be ruled by them. Relationships are based on the performance of technical activities that a mutually beneficial.", "Those with access to advanced technology tend to have higher social status", []string{"Technology", "Weapons", "Computers"}},
		6:   government{6, "Captive Government", "Ruling functions are performed by an imposed leadership answerable to an outside group.", "A colony or conquered area", []string{"Weapons", "Technology", "Travellers"}},
		7:   government{7, "Balkanisation", "No central authority exists; rival governments compete for control. Law level refers to the government nearest the starport.", "Multiple governments, civil war", []string{"Varies"}},
		8:   government{8, "Civil Service Bureaucracy", "Ruling functions are performed by government agencies employing individuals selected to their expertise.", "Technocracy, Communism", []string{"Drugs", "Weapons"}},
		9:   government{9, "Impersonal Bureaucracy", "Ruling functions are performed by agencies that have become insulated from the governed citizens.", "Entrenched castes of bureaucrats, decaying empire", []string{"Drugs", "Weapons"}},
		0xA: government{0xA, "Charismatic Dictator", "Ruling functions are performed by agencies directed by a single leader who enjoys the overwhelming confidence of the citizens.", "Revolutionary leader, messiah, emperor", []string{}},
		0xB: government{0xB, "Non-Charismatic Leader", "A previous charasmatic dictator has been replaced by a leader through normal channels.", "Military dictatorship, hereditary kingship", []string{"Weapons", "Technology", "Computers"}},
		0xC: government{0xC, "Charismatic Oligarchy", "Ruling funtcions are performed by a select group of members of an organization or class that enjoys the overwhelming confidence of the citizenry.", "Junta, revolutionary council", []string{"Weapons"}},
		0xD: government{0xD, "Religious Dictatorship", "Ruling functions are performed by a religious organization without regard to the specific individual needs of the citizenry.", "Cult, transcendent philosophy, psionic group mind", []string{"Varies"}},
		0xE: government{0xE, "Religious Autocracy", "Government by a single religious leader having absolute power over the citizenry.", "Messiah", []string{"Varies"}},
		0xF: government{0xF, "Totalitarian Oligarchy", "Government by an all-powerful minority which maintains absolute control through widespread coercion and oppression.", "World church, ruthless corporation", []string{"Varies"}},
	}

	return governments[roll2D()-7+population]
}

type faction struct {
	GovernmentType   government
	RelativeStrength string
}

func determineFactions(government uint8, population uint8) []faction {
	var modifier int8
	switch {
	case government == 0 || government == 7:
		modifier = 1
	case government >= 10:
		modifier = -1
	default:
		modifier = 0
	}

	numberOfFactions := int8(roll2D()) + int8(rollD()) + modifier

	factions := []faction{}

	for i := 0; i <= int(numberOfFactions); i++ {
		factionGovernment := determineGovernment(population)

		var relativeStrength string
		switch result := roll2D(); {
		case result < 3:
			relativeStrength = "Obscure group - few have heard of them, no popular support"
		case result < 6:
			relativeStrength = "Fringe group - few supporters"
		case result < 8:
			relativeStrength = "Minor group - some supporters"
		case result < 10:
			relativeStrength = "Notable group - significant support, well known"
		case result < 12:
			relativeStrength = "Significant - nearly as powerful as government"
		default:
			relativeStrength = "Overwhelming popular support - more powerful than government"
		}

		factions = append(factions, faction{factionGovernment, relativeStrength})
	}

	return factions
}

func determineCulturalDifferences() string {
	result := strconv.Itoa(int(rollD())) + strconv.Itoa(int(rollD()))

	toInt, err := strconv.ParseUint(result, 10, 8)
	if err != nil {
		fmt.Printf("error:", err)
	}

	culturalDifferences := map[uint8]string{
		11: "Sexist - one gender is considered subservient or inferior to the other.",
		12: "Religious - culture is heavily influenced by a religion or belief systems, possibly one unique to this world.",
		13: "Artistic - art and culture are highly prized. Aesthetic design is important in all artefacts produced on world.",
		14: "Ritualised – social interaction and trade is highly formalised. Politeness and adherence to traditional forms is considered very important.",
		15: "Conservative – the culture resists change and outside influences.",
		16: "Xenophobic – the culture distrusts outsiders and alien influences. Offworlders will face considerable prejudice.",
		21: "Taboo – a particular topic is forbidden and cannot be discussed. Travellers who unwittingly mention this topic will be ostracised.",
		22: "Deceptive – trickery and equivocation are considered acceptable. Honesty is a sign of weakness.",
		23: "Liberal – the culture welcomes change and offworld influence. Travellers who bring new and strange ideas will be welcomed.",
		24: "Honourable – one’s word is one’s bond in the culture. Lying is both rare and despised.",
		25: "Influenced – the culture is heavily influenced by another, neighbouring world. Roll again for a cultural quirk that has been inherited from the culture.",
		26: "Fusion – the culture is a merger of two distinct cultures. Roll again twice to determine the quirks inherited from these cultures. If the quirks are incompatible, then the culture is likely divided.",
		31: "Barbaric – physical strength and combat prowess are highly valued in the culture. Travellers may be challenged to a fight or dismissed if they seem incapable of defending themselves. Sports tend towards the bloody and violent.",
		32: "Remnant – the culture is a surviving remnant of a once-great and vibrant civilisation, clinging to its former glory. The world is filled with crumbling ruins and every story revolves around the good old days.",
		33: "Degenerate – the culture is falling apart and is on the brink of war or economic collapse. Violent protests are common and the social order is decaying.",
		34: "Progressive – the culture is expanding and vibrant. Fortunes are being made in trade; science is forging bravely ahead.",
		35: "Recovering – a recent trauma, such as a plague, war, disaster or despotic regime has left scars on the culture.",
		36: "Nexus – members of many different cultures and species visit here.",
		41: "Tourist Attraction – some aspect of the culture or the planet draws visitors from all over civilised space.",
		42: "Violent – physical conflict is common, taking the form of duels, brawls or other contests. Trial by combat is a part of their judicial system.",
		43: "Peaceful – physical conflict is almost unheard- of. The culture produces few soldiers and diplomacy reigns supreme. Forceful Travellers will be ostracised.",
		44: "Obsessed – everyone is obsessed with or addicted to a substance, personality, act or item. This monomania pervades every aspect of the culture.",
		45: "Fashion – fine clothing and decoration are considered vitally important in the culture. Underdressed Travellers have no standing here.",
		46: "At war – the culture is at war, either with another planet or polity, or is troubled by terrorists or rebels.",
		51: "Unusual Custom: Offworlders – space travellers hold a unique position in the culture’s mythology or beliefs and travellers will be expected to live up to these myths.",
		52: "Unusual Custom: Starport – the planet’s starport is more than a commercial centre; it might be a religious temple or be seen as highly controversial and surrounded by protestors.",
		53: "Unusual Custom: Media – news agencies and telecommunications channels are especially strange here. Getting accurate information may be difficult.",
		54: "Unusual Customs: Technology – the culture interacts with technology in an unusual way. Telecommunications might be banned, robots might have civil rights or cyborgs might be property.",
		55: "Unusual Customs: Lifecycle – there might be a mandatory age of termination or anagathics might be widely used. Family units might be different, with children being raised by the state or banned in favour of cloning.",
		56: "Unusual Customs: Social Standings – the culture has a distinct caste system. Travellers of a low social standing who do not behave appropriately will face punishment.",
		61: "Unusual Customs: Trade – the culture has an odd attitude towards some aspect of commerce, which may interfere with trade at the spaceport. For example, merchants might expect a gift as part of a deal or some goods may only be handled by certain families.",
		62: "Unusual Customs: Nobility – those of high social standing have a strange custom associated with them; perhaps nobles are blinded, must live in gilded cages or only serve for a single year before being exiled.",
		63: "Unusual Customs: Sex – the culture has an unusual attitude towards intercourse and reproduction. Perhaps cloning is used instead or sex is used to seal commercial deals.",
		64: "Unusual Customs: Eating – food and drink occupies an unusual place in the culture. Perhaps eating is a private affair or banquets and formal dinners are seen as the highest form of politeness.",
		65: "Unusual Customs: Travel – Travellers may be distrusted or feted, or perhaps the culture frowns on those who leave their homes.",
		66: "Unusual Custom: Conspiracy – something strange is going on. The government is being subverted by another group or agency.",
	}

	return culturalDifferences[toInt]
}

type lawLevel struct {
	Level         uint8
	WeaponsBanned string
	Armour        string
}

func determineLawLevel(government uint8) lawLevel {
	result := roll2D() - 7 + government

	lawLevels := map[uint8]lawLevel{
		0: lawLevel{0, "No restrictions - heavy armour and a handy weapon recommended...", ""},
		1: lawLevel{1, "Poison gas, explosives, undetectable weapons, WMD", "Battle dress"},
		2: lawLevel{2, "Portable energy and laser weapons", "Combat armour"},
		3: lawLevel{3, "Military weapons", "Flak"},
		4: lawLevel{4, "Light assault weapons and submachine guns", "Cloth"},
		5: lawLevel{5, "Personal concealable weapons", "Mesh"},
		6: lawLevel{6, "All firearms except shotguns & stunners", ""},
		7: lawLevel{7, "Shotguns", ""},
		8: lawLevel{8, "All bladed weapons, stunners", "All visible armour"},
		9: lawLevel{9, "All weapons", "All armour"},
	}
}

type starport struct {
	Class        string
	Quality      string
	BerthingCost string
	Fuel         string
	Facilities   string
	Bases        string
}

func determineStarport(population uint8) starport {
	var modifier int8
	switch {
	case population == 8 || population == 9:
		modifier = 1
	case population >= 10:
		modifier = 2
	case population == 3 || population == 4:
		modifier = -1
	case population <= 2:
		modifier = -2
	default:
		modifier = 0
	}

	result := roll2D() + uint8(modifier)

	var starportClass string
	switch {
	case result <= 2:
		starportClass = "X"
	case result < 5:
		starportClass = "E"
	case result < 7:
		starportClass = "D"
	case result < 9:
		starportClass = "C"
	case result < 11:
		starportClass = "B"
	default:
		starportClass = "A"
	}

	starports := map[string]starport{
		"A": starport{"A", "Excellent", "1DxCr1000", "Refined", "Shipyard (all), Repair, Highport 6+", "Military 8+, Naval 8+, Scout 10+"},
		"B": starport{"B", "Good", "1DxCr500", "Refined", "Shipyard (spacecraft), Repair, Highport 8+", "Military 8+, Naval 8+, Scout 9+"},
		"C": starport{"C", "Routine", "1DxCr100", "Unrefined", "Shipyard (small craft), Repair, Highport 10+", "Military 10+, Scout 9+"},
		"D": starport{"D", "Poor", "1DxCr10", "Unrefined", "Limited Repair, Highport 12+", "Scout 8+, Corsair 12+"},
		"E": starport{"E", "Frontier", "0", "None", "None", "Corsair 10+"},
		"X": starport{"X", "No Starport", "0", "None", "None", "Corsair 10+"},
	}

	return starports[starportClass]
}

func determineTechLevel(starport string, size uint8, atmosphere uint8, hydrographics uint8, population uint8, government uint8) uint8 {
	var starportMod int8
	switch {
	case starport == "A":
		starportMod = 6
	case starport == "B":
		starportMod = 4
	case starport == "C":
		starportMod = 2
	case starport == "X":
		starportMod = -4
	default:
		starportMod = 0
	}

	var sizeMod int8
	switch {
	case size < 2:
		sizeMod = 2
	case size < 5:
		sizeMod = 1
	default:
		sizemod = 0
	}

	var atmosphereMod int8
	if atmosphere < 4 || (atmosphere <= 0xF && atmosphere >= 0xA) {
		atmosphereMod = 1
	} else {
		atmosphereMod = 0
	}

	var hydrographicsMod int8
	switch {
	case hydrographics == 0 || hydrographics == 9:
		hydrographicsMod = 1
	case hydrographics == 0xA:
		hydrographicsMod = 2
	default:
		hydrographicsMod = 0
	}

	var populationMod int8
	switch {
	case (population > 0 && population < 6) || population == 8:
		populationMod = 1
	case population == 9:
		populationMod = 2
	case population == 0xA:
		populationMod = 4
	default:
		populationMod = 0
	}

	var governmentMod int8
	switch {
	case government == 0 || government == 5:
		governmentMod = 1
	case government == 7:
		governmentMod = 2
	case government == 13 || government == 14:
		governmentMod = -2
	default:
		governmentMod = 0
	}

	return rollD() + uint8(starportMod) + uint8(sizeMod) + uint8(atmosphereMod) + uint8(hydrographicsMod) + uint8(populationMod) + uint8(governmentMod)
}

func main() {
	cont := true

	for cont {
		planetSize := determinePlanetSize()
		atmosphere := determineAtmosphere(planetSize.Id)
		hydrographics := determineHydrographics(atmosphere.Id)
		temperature := determineTemperature(atmosphere.Id)
		population := determinePopulation()
		government := determineGovernment(population.Id)
		factions := determineFactions(government.Id, population.Id)
		culturalDifferences := determineCulturalDifferences()
		starport := determineStarport(population.Id)
		techLevel := determineTechLevel(starport.Class, planetSize.Id, atmosphere.Id, hydrographics.Id, population.Id, government.Id)

		println(planetSize)
		println(atmosphere)
		println(hydrographics)
		println(temperature)
		println(population)
		println(government)
		println(factions)
		println(culturalDifferences)
		println(starport)
		println(techLevel)

		fmt.Println("Would you like to continue? yes/no?")

		scanner := bufio.NewScanner(os.Stdin)

		cont = continueOrStop(scanner)
	}
}
