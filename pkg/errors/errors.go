package errors

import "errors"

var (
	ErrVerification     = errors.New("Ошибка верификации.")
	ErrFileExists       = errors.New("Файл уже существует.")
	ErrSymbolExists     = errors.New("Имя может содержать только латинские буквы, цифры, '_' и '-'.")
	ErrElementExists    = errors.New("Этот элемент уже добавлен. Попробуйте снова.")
	ErrElementNotExists = errors.New("Такого элемента нет в фильтре. Попробуйте снова.")
	ErrFileNotExists    = errors.New("Файл не существует.")
	ErrConfigNotExists  = errors.New("Конфигурационный файл отсутствует. Запустите программу снова и создайте его заново.")
	ErrInputNotExists   = errors.New("Файл с исходными данными отсутствует.")
	ErrOutputNotExists  = errors.New("Файл с результатами обработки отсутствует.")
)
