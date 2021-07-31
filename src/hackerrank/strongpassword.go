package hackerrank

import "strings"

type Password struct {
	rules []func(string) int32
}

func (p *Password) addRule(rule func(string) int32) {
	p.rules = append(p.rules, rule)
}

func (p *Password) StrongPassword(password string) int32 {
	l := length(password)
	count := int32(0)
	for _, rule := range p.rules {
		count += rule(password)
	}
	if l > count {
		return l
	}
	return count
}

func length(password string) int32 {
	l := 6 - int32(len(password))
	if l >= 1 {
		return l
	}
	return 0
}

func numbers(password string) int32 {
	if !strings.ContainsAny(password, "0123456789") {
		return 1
	}
	return 0
}

func lower(password string) int32 {
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return 1
	}
	return 0
}

func upper(password string) int32 {
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return 1
	}
	return 0
}

func special(password string) int32 {
	if !strings.ContainsAny(password, "!@#$%^&*()-+") {
		return 1
	}
	return 0
}
