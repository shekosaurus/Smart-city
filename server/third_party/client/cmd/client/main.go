package main

import (
	ServerClient "Smart-city/third_party/client/internal"
	"strings"
)

func main() {
	URL := ServerClient.DownloadNews()
	for _, url := range URL {
		ServerClient.GetPic(url, strings.Join([]string{"news", url}, "/"))
	}

	URL = ServerClient.GetBroadcast()
	for _, url := range URL {
		ServerClient.GetPic(url, strings.Join([]string{"broadcast", url}, "/"))
	}
}
