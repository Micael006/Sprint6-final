package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Все допустимые символы, взято из morse.go, так как в тех. задании
// нет определения ошибочного текста как такового
const allowedChars = "абвгдежзийклмнопрстуфхцчшщъыьэюя0123456789.,:?'-/()“=+@ "

// Функция для проверки является ли переданный текст Морзе или нет и последующего преобразования
func ConvertMessage(msg string) (string, error) {
	if len(msg) == 0 {
		return "", fmt.Errorf("provided string is empty")
	}

	isMorse := true
	for i, r := range []rune(strings.ToLower(msg)) {
		// Ошибка в случае обнаружения недопустимых символов
		if !strings.ContainsRune(allowedChars, r) {
			return "", fmt.Errorf("provided symbol '%c' at index %d is not legal", r, i)
		}
		if isMorse {
			isMorse = (r == ' ' || r == '.' || r == '-')
		}
	}
	if isMorse {
		return morse.ToText(msg), nil // Морзе
	}
	return morse.ToMorse(msg), nil // обычный текст
}
