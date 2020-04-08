package site

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	. "github.com/Foxcapades/Go-ChainRequest/simple"
	. "github.com/VEuPathDB/script-public-strategy-runner/internal/log"
)

// ResolveUrl takes the given URL string and formats it
// before using it to connect to the target site and
// resolve any redirects returning the final location.
func ResolveUrl(url string) (string, error) {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	return getActualUrl(url)
}

func getActualUrl(url string) (out string, err error) {
	Log().Tracef("Begin site.getActualUrl <- %s", url)
	defer func() { Log().Tracef("End site.getActualUrl -> %s, %s", out, err) }()

	res := GetRequest(url).
		DisableRedirects().
		Submit()

	if res.GetError() != nil {
		return "", res.GetError()
	}

	switch res.MustGetResponseCode() {
	case http.StatusMovedPermanently:
		if loc, ok := res.MustLookupHeader("Location"); !ok {
			return "", errors.New("site redirect missing \"Location\" header")
		} else {
			return getActualUrl(loc)
		}
	case http.StatusFound:
		if loc, ok := res.MustLookupHeader("Location"); !ok {
			return "", errors.New("site redirect missing \"Location\" header")
		} else {
			return loc, nil
		}
	case http.StatusOK:
		return url, nil
	default:
		return "", fmt.Errorf("unexpected site response code: %d",
			res.MustGetResponseCode())
	}
}
