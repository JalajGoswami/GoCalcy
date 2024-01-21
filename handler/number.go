package handler

import (
	"example/store"
	"strings"
)

func HandleNumberPressed(s string) func() {
	return func() {
		currentInput := store.GlobalStore.CurrentInput
		prev, err := currentInput.Get()
		if err != nil {
			store.ErrorChannel <- err
			return
		}
		if err != nil {
			store.ErrorChannel <- err
			return
		}
		currentInput.Set(prev + s)
	}
}

func ValidateNumberInput(r rune) bool {
	if !(r >= '0' && r <= '9') && r != '.' {
		return false
	}
	currentInput := store.GlobalStore.CurrentInput
	currValue, err := currentInput.Get()
	if err != nil {
		store.ErrorChannel <- err
		return false
	}
	if r == '.' {
		exist := strings.Contains(currValue, ".")
		return !exist
	}
	return true
}
