package switch_conditions

func switch_complex_conditions(city string) (string, string) {
	var region string
	var continent string

	switch city {

	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"

	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"

	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"

	default:
		region, continent = "Unknown", "Unknown"

	}

	return region, continent
}
