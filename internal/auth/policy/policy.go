package policy

import (
	"context"
	"crypto/x509"
	"fmt"
	"strings"

	"github.com/open-policy-agent/opa/rego"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrPermissionDenied = status.Error(codes.PermissionDenied, "you are not authorized to perform this operation")

// Attributes is a collection of metadata which can be used by policies to decide whether to accept or reject
// a protected operation.
type Attributes struct {
	Service string `json:"service"` // gRPC service name, fully qualified
	Method  string `json:"method"`  // gRPC method name
	// Metadata about streams, for streaming calls. For unary calls, both IsServerStream and IsClientStream
	// will be false.
	Stream StreamAttributes `json:"stream"`
	// gRPC request message for unary and server streaming calls.
	// For client and bidirectional streaming calls, Request is initially nil, but will contain the latest stream
	// message once Stream.Open is true.
	Request any `json:"request"`

	CertificateValid bool              `json:"certificate_valid"` // A cert is present and validated against the CA
	Certificate      *x509.Certificate `json:"certificate"`       // Claims in the validated certificate
	TokenValid       bool              `json:"token_valid"`       // A token is present and signature validated
	TokenClaims      any               `json:"token_claims"`      // Claims in the validated token
}

type StreamAttributes struct {
	IsServerStream bool `json:"is_server_stream"` // true for server streaming calls and bidirectional streaming calls
	IsClientStream bool `json:"is_client_stream"` // true for client streaming calls and bidirectional streaming calls
	// Open is false when the policy is being evaluated to decide whether the streaming call is allowed to open.
	// It is true when the streaming call is already open and the policy needs to check the latest incoming stream
	// message, which will be present in the Request attribute. Always false for server streaming calls and unary calls,
	// as policy is evaluated only once, with the request present.
	Open bool `json:"open"`
}

type Policy interface {
	EvalPolicy(ctx context.Context, query string, input Attributes) (rego.ResultSet, error)
}

// Validate will validate a set of decision attributes against a policy.
// It uses a hierarchical query system based on the service's name and package, querying the policy from most specific
// to least specific until the query returns a result.
// If the policy permits the request, returns nil. If it denies the request, returns ErrPermissionDenied.
// If the policy could not be evaluated successfully, returns an error.
//
// Services/protobuf packages are assumed to correspond to Rego packages of the same name; the gRPC service `foo.bar.Baz`
// corresponds to the Rego package `foo.bar.Baz`. The package should contain a boolean rule named `allow`.
// The policy accepts the request if querying for `data.foo.bar.Baz.allow` returns a result set with only
//`allow = true`. Otherwise, if the result set is empty then the next package up the hierarchy is tried, and if it's
// not empty then returns ErrPermissionDenied.
// If none of those queries return a result, we evaluate the query `data.grpc_default.allow` as a last resort.
// A package can prevent going further up the hierarchy by including a `default allow := false` rule - this
// is recommended to make policies easier to understand.
//
// For example, when evaluating access to the Service `foo.bar.Baz`, the following queries will be performed in order
// until one returns a non-empty result set:
//   - data.foo.bar.Baz.allow
//   - data.foo.bar.allow
//   - data.foo.allow
//   - data.grpc_default.allow
func Validate(ctx context.Context, policy Policy, attr Attributes) error {
	// the component parts of the protobuf package name (+ service name)
	components := strings.Split(attr.Service, ".")

	for len(components) > 0 {
		query := fmt.Sprintf("data.%s.allow", strings.Join(components, "."))
		result, err := policy.EvalPolicy(ctx, query, attr)
		if err != nil {
			return err
		}

		if len(result) > 0 {
			if result.Allowed() {
				return nil
			} else {
				return ErrPermissionDenied
			}
		}

		// if the result set is empty, we can try the next level up the hierarchy for a match
		components = components[:len(components)-1]
	}

	// try evaluating the fallback policy
	result, err := policy.EvalPolicy(ctx, "data.grpc_default.allow", attr)
	if err != nil {
		return err
	}

	if !result.Allowed() {
		return ErrPermissionDenied
	}
	return nil
}
