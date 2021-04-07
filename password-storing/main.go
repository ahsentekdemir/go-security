package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"context"
	"fmt"
)

const saltSize = 32
func main(){
	ct := context.Background()
	email := []byte("user@userdomain.com")
	password := []byte("userpassword")

	randomSalt  := make([]byte, saltSize)
	_, err := rand.Read(randomSalt)
	if err != nil { panic(err) }

	hashData := sha256.New()
	hashData.Write(password)
	hashData.Write(salt)

	stmt, err := db.PrepareContext(ct, "INSERT INTO user SET hash=?, salt=?, email=?")
	if err != nil { panic(err) }
	result, err := stmt.ExecContext(ct, hashData, salt, email)
	if err != nil { panic(err) }

}
