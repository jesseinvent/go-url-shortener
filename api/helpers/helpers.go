package helpers

import (
	"fmt"
	"strings"
	"os"
)

func EnforceHTTPS(url string) string {

	// Add HTTPS to URL
	if url[:5] != "https" && url[:4] != "http" {
		return "https://" + url;
	}

	if url[:5] != "https" && url[:4] == "http" {
		// length := len(url);
		fmt.Print(url[:4])
		return "https" + url[4:];
	}

	return url;
}

func RemoveDomainError(url string) bool {
	
	if url == os.Getenv("DOMAIN") {
		return false;
	}

	// Remove "http://", "https://" "www." from supplied url
	// And check it matches server's domain
	newUrl := strings.Replace(url, "http://", "", 1);
	newUrl = strings.Replace(newUrl, "https://", "", 1);
	newUrl = strings.Replace(newUrl, "www.", "", 1);
	newUrl = strings.Split(newUrl, "/")[0];

	if newUrl == os.Getenv("DOMAIN") {
		return false;
	}

	return true;
}
