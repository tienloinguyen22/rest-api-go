package utils

type signupprovider struct {
	EMAIL string
	FACEBOOK string
	GOOGLE string
	APPLE string
}

var SignupProvider = &signupprovider{
	EMAIL: "EMAIL",
	FACEBOOK: "FACEBOOK",
	GOOGLE: "GOOGLE",
	APPLE: "APPLE",
}