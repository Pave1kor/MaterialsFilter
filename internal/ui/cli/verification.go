package cli

import (
	errorsx "MaterialsFilter/pkg/errors"
	"fmt"
	"unicode"
)

// Проверить правильность введенного имени
func verificationName(name string) error {
	if name == "" {
		err := fmt.Errorf("Имя не может быть пустым: %w", errorsx.ErrVerification)
		fmt.Println(err)
		return err
	}

	for _, r := range name {
		if unicode.IsSpace(r) {
			err := fmt.Errorf("В имени не должно быть пробелов: %w", errorsx.ErrVerification)
			fmt.Println(err)
			return err
		}

		if !(unicode.Is(unicode.Latin, r) ||
			unicode.IsDigit(r) ||
			r == '_' ||
			r == '-') {
			err := fmt.Errorf("Имя может содержать только латинские буквы, цифры, '_' и '-': %w", errorsx.ErrVerification)
			fmt.Println(err)
			return err
		}
	}
	return nil
}

// Проверка: да/нет
func Verification() bool {
	for {
		fmt.Println("Подтверждение ввода.")
		fmt.Print("Ожидаю (да/нет) или (yes/no): ")

		input := NewLine()

		switch input {
		case "да", "yes", "y":
			return true
		case "нет", "no", "n":
			return false
		default:
			fmt.Println("Введите 'да' или 'нет'.")
		}
	}
}
