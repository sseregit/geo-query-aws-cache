package service

import (
	"strings"
)

func solveImageExtension(e string) bool {
	extension := strings.ToLower(e)

	switch extension {
	case ".jpeg":
		return true
	case ".jpg":
		return true
	case ".webp":
		return true
	default:
		return false
	}
}
