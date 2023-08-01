package helper

import "github.com/jaevor/go-nanoid"

const defaultCharsId = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func RandomChar(length int) string {
	gen, _ := nanoid.CustomUnicode(defaultCharsId, length)
	return gen()
}

