package utils

import (
	"strconv"
	"strings"
)

// IsHex проверяет, является ли строка hex-строкой
func IsHex(s string) bool {
	// Проверяем, что строка состоит только из hex-символов
	for _, c := range s {
		if !strings.ContainsRune("0123456789abcdefABCDEF", c) {
			return false
		}
	}
	return true
}

// StringToUint64 конвертирует строку в uint64
func StringToUint64(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}

// Uint64ToString конвертирует uint64 в строку
func Uint64ToString(u uint64) string {
	return strconv.FormatUint(u, 10)
}

// StringToFloat64 конвертирует строку в float64
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Float64ToString конвертирует float64 в строку
func Float64ToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
