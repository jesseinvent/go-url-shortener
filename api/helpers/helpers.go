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

func RemoveSERVER_HOSTError(url string) bool {
	
	if url == os.Getenv("SERVER_HOST") {
		return false;
	}

	// Remove "http://", "https://" "www." from supplied url
	// And check it matches server's SERVER_HOST
	newUrl := strings.Replace(url, "http://", "", 1);
	newUrl = strings.Replace(newUrl, "https://", "", 1);
	newUrl = strings.Replace(newUrl, "www.", "", 1);
	newUrl = strings.Split(newUrl, "/")[0];

	if newUrl == os.Getenv("SERVER_HOST") {
		return false;
	}

	return true;
}
