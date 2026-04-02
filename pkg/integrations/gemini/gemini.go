package gemini

import (
	"github.com/superplanehq/superplane/pkg/configuration"
	"github.com/superplanehq/superplane/pkg/core"
	"github.com/superplanehq/superplane/pkg/registry"
)

func init() {
	registry.RegisterIntegration("gemini", &Gemini{})
}

type Configuration struct {
	APIKey string `json:"apiKey"`
}

type Gemini struct{}

func (g Gemini) Name() string {
	return "gemini"
}

func (g Gemini) Label() string {
	return "Gemini"
}

func (g Gemini) Icon() string {
	return "sparkles"
}

func (g Gemini) Description() string {
	return "Interact with Gemini models"
}

func (g Gemini) Instructions() string {
	return ""
}

func (g Gemini) Configuration() []configuration.Field {
	return []configuration.Field{
		{
			Name:        "apiKey",
			Label:       "API Key",
			Type:        configuration.FieldTypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Gemini API key",
		},
	}
}

func (g Gemini) Components() []core.Component {
	//TODO implement me
	panic("implement me")
}

func (g Gemini) Triggers() []core.Trigger {
	return []core.Trigger{}
}

func (g Gemini) Sync(ctx core.SyncContext) error {
	//TODO implement me
	panic("implement me")
}

func (g Gemini) Cleanup(ctx core.IntegrationCleanupContext) error {
	return nil
}

func (g Gemini) Actions() []core.Action {
	return []core.Action{}
}

func (g Gemini) HandleAction(ctx core.IntegrationActionContext) error {
	return nil
}

func (g Gemini) ListResources(resourceType string, ctx core.ListResourcesContext) ([]core.IntegrationResource, error) {
	//TODO implement me
	panic("implement me")
}

func (g Gemini) HandleRequest(ctx core.HTTPRequestContext) {
}
