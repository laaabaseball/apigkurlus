package handler

import (
	"net/http"
	"time"

	"github.com/laaabaseball/apigkurlus/pkg/client"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	httpClient := http.DefaultClient
	client.RssHandler("https://old.reddit.com", time.Now, httpClient, client.GetArticle, w, r)
}
