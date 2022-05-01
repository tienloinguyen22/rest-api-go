package utils

func IsValidGender(gender string) bool {
	genders := map[string]bool{
		GENDER_MALE: true,
		GENDER_FEMALE: true,
		GENDER_OTHER: true,
	}
	return genders[gender] || false
}