package main

import "regexp"

func GetSubMatchMap(re *regexp.Regexp, str string) (map[string]string, error) {
	match := re.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	if match != nil {
		for i, name := range re.SubexpNames() {
			if i != 0 {
				subMatchMap[name] = match[i]
			}
		}
	}
	return subMatchMap, nil
}
