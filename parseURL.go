package main

import (
	"fmt"
    "regexp"
    "strings"
)

type Url struct {
    protocol string
    host string
    path string
    params map[string][]string
}

func (u Url) Output() string {
    return fmt.Sprintf("protocol: %v\nhost: %v\npath: %v\nparams: %v\n", u.protocol, u.host, u.path, u.params)
}

func trimParamKey(paramKey string) string {
    if strings.HasSuffix(paramKey, "[]") {
        paramKey = paramKey[:len(paramKey)-len("[]")]
    }

    return paramKey
}

func ParseURL(urlString string) Url {
    u := Url{}
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
        key := trimParamKey(parts[0])
        if len(parts) > 1 {
            queryValue = parts[1]
        }

        if _, ok := u.params[key]; ok {
            u.params[key] = append(u.params[key], queryValue)
        } else {
            u.params[key] = []string{queryValue}
        }
    }

    return u
}
