package pkg

func HashPassword(password string) string {
	return password
}

func CheckPasswordHash(password, hash string) bool {
	return password == hash
}
