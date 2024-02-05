package switch_conditions

func switch_complex_conditions_2(city string) (string, string) {
	var region string
	var continent string

	switch city {

	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"

		fallthrough
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
		fallthrough

	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"

	default:
		region, continent = "Unknown", "Unknown"

	}

	return region, continent
}
