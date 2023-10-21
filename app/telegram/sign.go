package telegram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SignParams struct {
	Payload string
	Key     string
	Secret  []byte
}

// Sign signs passed payload using specified key. Function removes such
// technical parameters as "hash" and "auth_date".
func Sign(payload map[string]string, key string, authDate time.Time) string {
	pairs := make([]string, 0, len(payload)+1)

	// Extract all key-value pairs and add them to pairs slice.
	for k, v := range payload {
		// Skip technical fields.
		if k == "hash" || k == "auth_date" {
			continue
		}
		// Append new pair.
		pairs = append(pairs, k+"="+v)
	}

	// Append sign date.
	pairs = append(pairs, "auth_date="+strconv.FormatInt(authDate.Unix(), 10))

	// According to docs, we sort all the pairs in alphabetical order.
	sort.Strings(pairs)

	// Perform signing.
	return sign(&SignParams{
		Payload: strings.Join(pairs, "\n"),
		Key:     key,
	})
}

// SignQueryString signs passed query string.
func SignQueryString(qs, key string, authDate time.Time) (string, error) {
	// Parse query string.
	qp, err := url.ParseQuery(qs)
	if err != nil {
		return "", err
	}

	// Convert query params to map[string]string.
	m := make(map[string]string, len(qp))
	for k, v := range qp {
		m[k] = v[0]
	}
	return Sign(m, key, authDate), nil
}

// Performs payload subscription. Payload itself slice of key-value pairs
// joined with "\n".
func sign(r *SignParams) string {
	if r.Secret == nil {
		r.Secret = []byte("WebAppData")
		skHmac := hmac.New(sha256.New, []byte(r.Secret))
		skHmac.Write([]byte(r.Key))

		impHmac := hmac.New(sha256.New, skHmac.Sum(nil))
		impHmac.Write([]byte(r.Payload))

		return hex.EncodeToString(impHmac.Sum(nil))
	} else {
		h := hmac.New(sha256.New, r.Secret)
		h.Write([]byte(r.Payload))

		return hex.EncodeToString(h.Sum(nil))
	}
}
