package crc_md5

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	md52 "github.com/golang/go/src/pkg/crypto/md5"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"testing"
)

type User struct {
	Name       string
	FamilyName string
	Gener      string
}

type Address struct {
	Street string
	Number int
}

type test struct {
	User    User
	Address Address
}

var user = test{
	User: User{
		Name:       "Dino",
		FamilyName: "Sauro",
		Gener:      "T-Rex",
	},
	Address: Address{
		Street: "Pagea Street",
		Number: 1000,
	},
}

func BenchmarkAdler32(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := adler32.New()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkCRC32IEEE(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := crc32.New(crc32.MakeTable(crc32.IEEE))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkCRC32Castagnoli(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := crc32.New(crc32.MakeTable(crc32.Castagnoli))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkCRC31Koopman(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := crc32.New(crc32.MakeTable(crc32.Koopman))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkCRC64ISO(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := crc64.New(crc64.MakeTable(crc64.ISO))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkCRC64ECMA(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := crc64.New(crc64.MakeTable(crc64.ECMA))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkFNV32(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := fnv.New32()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkFNV32a(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := fnv.New32a()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkFNV64(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := fnv.New64()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkFNV64a(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := fnv.New64a()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkHmacSha256(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	key := "SHA-256"
	hash := hmac.New(sha256.New, []byte(key))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkHmacMd5(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	key := "123"
	hash := hmac.New(md5.New, []byte(key))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkHmacSha1(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	key := "123"
	hash := hmac.New(sha1.New, []byte(key))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkHmacSha512(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	key := "123"
	hash := hmac.New(sha512.New, []byte(key))

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkMd5(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := md5.New()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkMd52(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := md52.New()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkSha1(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := sha1.New()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkSha254(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := sha256.New224()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkSha256(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := sha256.New()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}

func BenchmarkSha512(b *testing.B) {
	var err error
	var data []byte
	data, err = json.Marshal(&user)
	if err != nil {
		b.FailNow()
	}

	hash := sha512.New384()

	for i := 0; i < b.N; i++ {
		_, err = hash.Write(data)
		if err != nil {
			b.FailNow()
		}
		hash.Sum(nil)
	}
}
