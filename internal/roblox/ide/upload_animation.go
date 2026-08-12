package ide

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/kamsislayerson-netizen/Asset-Reuploader/internal/roblox"
)

var UploadAnimationErrors = struct {
	ErrNotLoggedIn       error
	ErrTokenInvalid      error
	ErrInappropriateName error
}{
	ErrNotLoggedIn:       errors.New("not logged in"),
	ErrTokenInvalid:      errors.New("XSRF token validation failed"),
	ErrInappropriateName: errors.New("inappropriate name or description"),
}

func newAnimationURL(groupID int64, name, description string) string {
	url := fmt.Sprintf("https://www.roblox.com/ide/publish/UploadNewAnimation?assetTypeName=Animation&name=%s&description=%s",
		url.QueryEscape(name),
		url.QueryEscape(description),
	)
	if groupID > 0 {
		url += fmt.Sprintf("&groupId=%d", groupID)
	}

	return url
}

func newUploadAnimationRequest(
	groupID int64,
	name,
	description string,
	data *bytes.Buffer,
) (*http.Request, error) {
	url := newAnimationURL(groupID, name, description)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "RobloxStudio/WinInet")

	return req, nil
}

func NewUploadAnimationHandler(
	c *roblox.Client,
	name,
	description string,
	data *bytes.Buffer,
	groupID ...int64,
) (func() (int64, error), error) {
	group := groupID[0]
	// Store the data bytes to allow resetting for retries
	dataBytes := data.Bytes()

	return func() (int64, error) {
		// Create a fresh request with a new reader for each attempt
		url := newAnimationURL(group, name, description)
		req, err := http.NewRequest("POST", url, bytes.NewReader(dataBytes))
		if err != nil {
			return 0, err
		}
		req.Header.Set("User-Agent", "RobloxStudio/WinInet")
		req.Header.Set("Content-Type", "application/octet-stream")
		req.AddCookie(&http.Cookie{
			Name:  ".ROBLOSECURITY",
			Value: c.Cookie,
		})
		req.Header.Set("x-csrf-token", c.GetToken())

		resp, err := c.DoRequest(req)
		if err != nil {
			return 0, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}

		switch resp.StatusCode {
		case http.StatusOK:
			id, err := strconv.ParseInt(string(body), 10, 64)
			if err != nil {
				return 0, err
			}

			return id, nil
		case http.StatusForbidden:
			if strBody := string(body); strBody == "NotLoggedIn" {
				return 0, UploadAnimationErrors.ErrNotLoggedIn
			} else if strBody == "XSRF Token Validation Failed" {
				c.SetToken(resp.Header.Get("x-csrf-token"))
				return 0, UploadAnimationErrors.ErrTokenInvalid
			}

			return 0, errors.New(resp.Status)
		case http.StatusUnprocessableEntity:
			if string(body) == "Inappropriate name or description." {
				req, _ = newUploadAnimationRequest(group, "[Censored]", description, data)
				return 0, UploadAnimationErrors.ErrInappropriateName
			}

			return 0, errors.New(resp.Status)
		default:
			return 0, errors.New(resp.Status)
		}
	}, nil
}
