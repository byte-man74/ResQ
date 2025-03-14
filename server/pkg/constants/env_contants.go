package constants

import "os"

var JWTSecretKey = []byte(os.Getenv("JWT_SECRET"))
