package main

import (
	"fmt"
    "regexp"
    "strings"
)

type url struct {
    protocol string
    host string
    path string
    params map[string][]string
}

func parseURL(urlString string) url {
    u := url{}
    re := regexp.MustCompile(`^(?P<protocol>[a-z,0-9]+)?(?:://)?(?P<host>[\.,a-z,0-9]+)?(?P<path>[^\?]+)`)
    groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(urlString, -1) {
    		for groupIdx, value := range match {
        		switch groupNames[groupIdx] {
                case "protocol":
                    u.protocol = value
                case "host":
                    u.host = value
                case "path":
                    u.path = value
                }
    		}
	}

    u.params = make(map[string][]string)
    queries, urlParts := []string{}, strings.Split(urlString, "?")
    if len(urlParts) > 1 {
        queries = strings.Split(urlParts[1], "&")
    }

    for _, query := range queries {
        parts, queryValue := strings.Split(query, "="), ""
        if len(parts) > 1 {
            queryValue = parts[1]
        }

        if _, ok := u.params[parts[0]]; ok {
            u.params[parts[0]] = append(u.params[parts[0]], queryValue)
        } else {
            u.params[parts[0]] = []string{queryValue}
        }
    }

    return u
}

func main() {
	fmt.Println(parseURL("http://apple.com/search?term=ipad"))
	fmt.Println(parseURL("https://facebook.com/create/user?firstName=john&lastName=doe&city=venice"))
}
