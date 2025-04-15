package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func HashPassword(password string) (string, error) {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	return string(bcryptHash), err
}

func ComparePassword(hashedPassword string, plainPassword string) error {
	byteHash := []byte(hashedPassword)
	return bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword))
}

func HashMd5(data ...any) string {
	str := fmt.Sprintf("%v", data)
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func ToBson(doc interface{}) bson.M {
	var res bson.M
	b, _ := bson.Marshal(doc)
	_ = bson.Unmarshal(b, &res)
	return res
}

func TimeString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(time.RFC3339Nano)
}

func NewUuidV7String() string {
	id, err := uuid.NewV7()
	if err != nil {
		return uuid.NewString()
	}
	return id.String()
}
