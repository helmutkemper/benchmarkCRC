package crc_md5

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
)

// fixme: key must be set

func GetHash(a string) (hash.Hash, error) {
	var h hash.Hash
	switch a {
	case "adler32":
		h = adler32.New()
	case "crc32", "crc32ieee":
		h = crc32.New(crc32.MakeTable(crc32.IEEE))
	case "crc32castagnoli":
		h = crc32.New(crc32.MakeTable(crc32.Castagnoli))
	case "crc32koopman":
		h = crc32.New(crc32.MakeTable(crc32.Koopman))
	case "crc64", "crc64iso":
		h = crc64.New(crc64.MakeTable(crc64.ISO))
	case "crc64ecma":
		h = crc64.New(crc64.MakeTable(crc64.ECMA))
	case "fnv", "fnv32":
		h = fnv.New32()
	case "fnv32a":
		h = fnv.New32a()
	case "fnv64":
		h = fnv.New64()
	case "fnv64a":
		h = fnv.New64a()
	case "hmac", "hmacsha256":
		key := "SHA-256"
		h = hmac.New(sha256.New, []byte(key))
	case "hmacmd5":
		key := "SHA-256"
		h = hmac.New(md5.New, []byte(key))
	case "hmacsha1":
		key := "SHA-256"
		h = hmac.New(sha1.New, []byte(key))
	case "hmacsha512":
		key := "SHA-256"
		h = hmac.New(sha512.New, []byte(key))
	case "md5":
		h = md5.New()
	case "sha1":
		h = sha1.New()
	case "sha224":
		h = sha256.New224()
	case "sha256":
		h = sha256.New()
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		return nil, errors.New("Invalid algorithm")
	}
	return h, nil
}

func HashString(h hash.Hash, s string) string {
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
