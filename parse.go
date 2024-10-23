package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Event represents a webhook event.
type Event struct {
	// ID is a unique id for this event.
	ID string `json:"id"`

	// WebhookID is the id of the webhook that received this event.
	WebhookID string `json:"webhook_id"`

	// SequenceID is a unique, incrementing sequence id for this event,
	// specific to the webhook.
	SequenceID int64 `json:"sequence_id"`

	// Type is the type of the event.
	Type string `json:"type"`

	// Data is the parsed event data.
	// The concrete type depends on what event it is.
	Data any `json:"data"`
}

// ParseEvent parses a webhook event from the given payload.
func ParseEvent(payload []byte, header, secret string) (*Event, error) {
	const tolerance = 300 * time.Second
	if err := validatePayload(payload, header, secret, tolerance); err != nil {
		return nil, err
	}

	var raw rawEvent
	if err := json.Unmarshal(payload, &raw); err != nil {
		return nil, fmt.Errorf("webhooks: parse event: %v", err)
	}

	data, err := parseEventData(raw.Type, raw.Data)
	if err != nil {
		return nil, fmt.Errorf("webhooks: parse event data: %v", err)
	}

	ev := &Event{
		ID:         raw.ID,
		WebhookID:  raw.WebhookID,
		SequenceID: raw.SequenceID,
		Type:       raw.Type,
		Data:       data,
	}
	return ev, nil
}

type rawEvent struct {
	ID         string          `json:"id"`
	WebhookID  string          `json:"webhook_id"`
	SequenceID int64           `json:"sequence_id"`
	Type       string          `json:"type"`
	Data       json.RawMessage `json:"data"`
}

var errUnknownEvent = errors.New("unknown event type")

func parseEventData(typ string, data json.RawMessage) (any, error) {
	var dst any
	var err error

	switch prefix, _, _ := strings.Cut(typ, "."); prefix {
	case "rollout":
		dst, err = parseRolloutEvent(typ)
	default:
		err = errUnknownEvent
	}

	if err != nil {
		if errors.Is(err, errUnknownEvent) {
			return nil, fmt.Errorf("unknown event type %q", typ)
		}
		return nil, err
	}

	if err := json.Unmarshal(data, dst); err != nil {
		return nil, err
	}
	return dst, nil
}

// ComputeSignature computes the MAC signature for the given payload.
func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(strconv.FormatInt(t.Unix(), 10)))
	h.Write([]byte("."))
	h.Write(payload)
	return h.Sum(nil)
}

var (
	ErrNoSignature      = errors.New("missing KhulnaSoft-Signature header")
	ErrInvalidSignature = errors.New("invalid KhulnaSoft-Signature header")
	ErrNoValidSignature = errors.New("no valid KhulnaSoft-Signature signatures")
	ErrTooOld           = errors.New("webhook event is too old")
)

type signature struct {
	ts   time.Time
	macs [][]byte
}

// parseSignatureHeader parses the KhulnaSoft signature header.
func parseSignatureHeader(val string) (*signature, error) {
	if val == "" {
		return nil, ErrNoSignature
	}
	sig := &signature{}
	parts := strings.Split(val, ",")
	for _, p := range parts {
		key, val, ok := strings.Cut(p, "=")
		if !ok {
			return nil, ErrInvalidSignature
		}
		switch key {
		case "t":
			ts, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return nil, ErrInvalidSignature
			}
			sig.ts = time.Unix(ts, 0)

		case "v1":
			data, err := base64.RawURLEncoding.DecodeString(val)
			if err != nil {
				// Ignore invalid signatures; as long as at least one valid signature
				// exists (checked later), we're good.
				continue
			}
			sig.macs = append(sig.macs, data)

		default:
			continue
		}
	}

	if len(sig.macs) == 0 {
		return nil, ErrNoValidSignature
	} else if sig.ts.IsZero() {
		return nil, ErrInvalidSignature
	}

	return sig, nil
}

// validatePayload validates the payload against the given signature header and secret,
// and ensures the event is not too old.
func validatePayload(payload []byte, sigHeader, secret string, tolerance time.Duration) error {
	sig, err := parseSignatureHeader(sigHeader)
	if err != nil {
		return err
	}
	wantSig := ComputeSignature(sig.ts, payload, secret)
	if time.Since(sig.ts) > tolerance {
		return ErrTooOld
	}
	for _, mac := range sig.macs {
		if hmac.Equal(mac, wantSig) {
			return nil
		}
	}
	return ErrNoValidSignature
}
