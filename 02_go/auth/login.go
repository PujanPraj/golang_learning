package auth

import "fmt"

func Login(name string) string {
	res := fmt.Sprintf("Successful login in by %s", name)
	return res
} 