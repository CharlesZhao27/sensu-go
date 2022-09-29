// Code generated by interfacer; DO NOT EDIT

package mockapi

import (
	"context"
	"github.com/sensu/sensu-go/api/core/v2"
)

// CheckClient is an interface generated for "github.com/sensu/sensu-go/backend/api.CheckClient".
type CheckClient interface {
	CreateCheck(context.Context, *v2.CheckConfig) error
	DeleteCheck(context.Context, string) error
	ExecuteCheck(context.Context, string, *v2.AdhocRequest) error
	FetchCheck(context.Context, string) (*v2.CheckConfig, error)
	ListChecks(context.Context) ([]*v2.CheckConfig, error)
	UpdateCheck(context.Context, *v2.CheckConfig) error
}