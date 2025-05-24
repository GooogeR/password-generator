package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	lettersLower = "qwertyuiopasdfghjklzxcvbnm"
	lettersUpper = "QWERTYUIOPASDFGHJKLZXCVBNM"
	digits       = "0123456789"
	symbols      = "!@#$%^&*()_+-=[]{}<>?"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите длину пароля: ")
	scanner.Scan()
	lengthStr := scanner.Text()
	length, err := strconv.Atoi(lengthStr)
	if err != nil || length < 0 {
		fmt.Println("Ошибка, введено не верное число")
		return
	}

	fmt.Println("Добавлять цифры? (Да/Нет): ")
	scanner.Scan()
	useDigits := strings.ToLower(scanner.Text()) == "да"

	fmt.Println("Добавлять Символы? (Да/Нет): ")
	scanner.Scan()
	useSymbols := strings.ToLower(scanner.Text()) == "да"

	fmt.Println("Использовать заглавные буквы? (Да/Нет): ")
	scanner.Scan()
	useUpper := strings.ToLower(scanner.Text()) == "да"

	fmt.Println("Использовать строчные буквы? (Да/Нет): ")
	scanner.Scan()
	useLower := strings.ToLower(scanner.Text()) == "да"

	chars := ""
	if useLower {
		chars += lettersLower
	}
	if useUpper {
		chars += lettersUpper
	}
	if useDigits {
		chars += digits
	}
	if useSymbols {
		chars += symbols
	}

	if len(chars) == 0 {
		fmt.Println("Ошибка: не выбрано ни одного набора символов")
		return
	}

	password := generatePassword(length, chars)
	fmt.Println("Пароль сгенерирован:", password)
}

func generatePassword(length int, chars string) string {
	password := make([]byte, length)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		index := rand.Intn(len(chars))
		password[i] = chars[index]
	}

	return string(password)
}
