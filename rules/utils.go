package rules

import (
	"errors"
	"strings"
)

var validVMAdminUserNames = map[string]bool{
	"azureuser":        false,
	"1":                false,
	"123":              false,
	"a":                false,
	"actuser":          false,
	"adm":              false,
	"admin":            false,
	"admin1":           false,
	"admin2":           false,
	"administrator":    false,
	"aspnet":           false,
	"backup":           false,
	"console":          false,
	"david":            false,
	"guest":            false,
	"john":             false,
	"owner":            false,
	"root":             false,
	"server":           false,
	"sql":              false,
	"support_388945a0": false,
	"support":          false,
	"sys":              false,
	"test":             false,
	"test1":            false,
	"test2":            false,
	"test3":            false,
	"user":             false,
	"user1":            false,
	"user2":            false,
	"user3":            false,
	"user4":            false,
	"user5":            false,
	"video":            false,
}

func isVlidVMAdminUserNames(userName string) (bool, error) {
	if userName == "" {
		return false, errors.New("userName is empty")
	}
	return validVMAdminUserNames[strings.ToLower(userName)], nil
}
