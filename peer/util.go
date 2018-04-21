package peer

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetMyExternalIP returns public ip by local.
func GetMyExternalIP() (string, error) {
	resp, err := http.Get("https://ipv4.myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	buf, err := ioutil.ReadAll(io.LimitReader(resp.Body, 64))
	if err != nil {
		return "", err
	}
	if len(buf) == 0 {
		return "", errors.New("myexternalip.com returned a 0 length IP address")
	}
	return strings.TrimSpace(string(buf)), nil
}
