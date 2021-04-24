package roundrobin

import (
	"log"
	"net/url"
	"strconv"
	"strings"
)

func GetUrlsWithNextPorts(originUrl string, count int) []string {
	u, err := url.Parse(originUrl)
	if err != nil {
		log.Panicln("Cannot create RR for proxy server", err.Error())
		return nil
	}

	port, _ := strconv.Atoi(u.Port())
	validUrls := []string{originUrl}
	for i := 1; i < count; i++ {
		copy := u
		splittedHost := strings.Split(copy.Host, ":")
		splittedHost[1] = strconv.Itoa(port + i)
		copy.Host = strings.Join(splittedHost, ":")
		validUrls = append(validUrls, copy.String())
	}

	return validUrls
}
