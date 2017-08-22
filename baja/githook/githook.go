// https://github.com/rjz/githubhook/blob/master/githubhook.go
package githubhook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

// Hook is an inbound github webhook
type Hook struct {

	// Id specifies the Id of a github webhook request.
	//
	// Id is extracted from the inbound request's `X-Github-Delivery` header.
	Id string

	// Event specifies the event name of a github webhook request.
	//
	// Event is extracted from the inbound request's `X-GitHub-Event` header.
	// See: https://developer.github.com/webhooks/#events
	Event string

	// Signature specifies the signature of a github webhook request.
	//
	// Signature is extracted from the inbound request's `X-Hub-Signature` header.
	Signature string

	// Payload contains the raw contents of the webhook request.
	//
	// Payload is extracted from the JSON-formatted body of the inbound request.
	Payload []byte
}

const signaturePrefix = "sha1="
const signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

// SignedBy checks that the provided secret matches the hook Signature
//
// Implements validation described in github's documentation:
// https://developer.github.com/webhooks/securing/
func (h *Hook) SignedBy(secret []byte) bool {
	if len(h.Signature) != signatureLength || !strings.HasPrefix(h.Signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(h.Signature[5:]))

	return hmac.Equal(signBody(secret, h.Payload), actual)
}

// Extract unmarshals Payload into a destination interface.
func (h *Hook) Extract(dst interface{}) error {
	return json.Unmarshal(h.Payload, dst)
}

// New reads a Hook from an incoming HTTP Request.
func New(req *http.Request) (hook *Hook, err error) {
	hook = new(Hook)
	if !strings.EqualFold(req.Method, "POST") {
		return nil, errors.New("Unknown method!")
	}

	if hook.Signature = req.Header.Get("x-hub-signature"); len(hook.Signature) == 0 {
		return nil, errors.New("No signature!")
	}

	if hook.Event = req.Header.Get("x-github-event"); len(hook.Event) == 0 {
		return nil, errors.New("No event!")
	}

	if hook.Id = req.Header.Get("x-github-delivery"); len(hook.Id) == 0 {
		return nil, errors.New("No event Id!")
	}

	hook.Payload, err = ioutil.ReadAll(req.Body)
	return
}

// Parse reads and verifies the hook in an inbound request.
func Parse(secret []byte, req *http.Request) (hook *Hook, err error) {
	hook, err = New(req)
	if err == nil && !hook.SignedBy(secret) {
		err = errors.New("Invalid signature")
	}
	return
}
