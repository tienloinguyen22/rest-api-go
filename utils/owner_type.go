package utils

type ownertype struct {
	STUDENT string
	PARENT string
	TEACHER string
	OTHER string
}

var OwnerType = &ownertype{
	STUDENT: "STUDENT",
	PARENT: "PARENT",
	TEACHER: "TEACHER",
	OTHER: "OTHER",
}