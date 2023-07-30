package userclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	userclientpb "github.com/requiemofthesouls/user-client/pb"
)

const (
	contextUserClient = "userClient"
	contextRequestID  = "requestID"
)

var ErrClientIsNotSetToContext = errors.New("client is not set to context")

func ExtractContextToClient(ctx context.Context) *userclientpb.Client {
	if v, ok := ctx.Value(contextUserClient).(*userclientpb.Client); ok {
		return v
	}

	return nil
}

func ExtractContextToClientWithError(ctx context.Context) (*userclientpb.Client, error) {
	var v any
	if v = ctx.Value(contextUserClient); v == nil {
		return nil, ErrClientIsNotSetToContext
	}

	if c, ok := v.(*userclientpb.Client); ok {
		return c, nil
	}

	return nil, fmt.Errorf("can't cast client %v to *client.Client", v)
}

func ClientToContext(ctx context.Context, c *userclientpb.Client) context.Context {
	return context.WithValue(ctx, contextUserClient, c)
}

func ExtractContextToRequestID(ctx context.Context) string {
	if v, ok := ctx.Value(contextRequestID).(string); ok {
		return v
	}

	return ""
}

func RequestIDToContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, contextRequestID, requestID)
}

func GenerateRequestID() string {
	return uuid.Must(uuid.NewV4()).String()
}

func MarshallClient(c *userclientpb.Client) []byte {
	bytes, _ := json.Marshal(c)
	return bytes
}

func UnmarshallClient(data []byte) *userclientpb.Client {
	var client userclientpb.Client
	_ = json.Unmarshal(data, &client)
	return &client
}
