package utils

func IsValidOwnerType(ownerType string) bool {
	ownerTypes := map[string]bool{
		OWNER_TYPE_STUDENT: true,
		OWNER_TYPE_PARENT: true,
		OWNER_TYPE_TEACHER: true,
		OWNER_TYPE_OTHER: true,
	}
	return ownerTypes[ownerType] || false
}