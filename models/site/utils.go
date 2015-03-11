package site

import (
	"regexp"

	"github.com/robpike/filter"
)

/*
	Fuzzy select to domains by name
*/
func SelectDomains(name string) []string {
	return filter.Choose(DOMAINS, func(d string) bool {
		return regexp.MustCompile(d).MatchString(name)
	}).([]string)
}
