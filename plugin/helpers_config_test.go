package plugin_test

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
	"github.com/mattermost/mattermost-server/v5/plugin/plugintest"
	"github.com/stretchr/testify/assert"
)

func TestCheckRequiredServerConfiguration(t *testing.T) {
	for name, test := range map[string]struct {
		SetupAPI     func(*plugintest.API) *plugintest.API
		Input        *model.Config
		ShouldReturn bool
		ShouldError  bool
	}{
		"no required config therefore it should be compatible": {
			SetupAPI: func(api *plugintest.API) *plugintest.API {
				return api
			},
			Input:        nil,
			ShouldReturn: true,
			ShouldError:  false,
		},
		"same configurations": {
			SetupAPI: func(api *plugintest.API) *plugintest.API {
				api.On("GetConfig").Return(&model.Config{
					ServiceSettings: model.ServiceSettings{
						EnableCommands: model.NewBool(true),
					},
				})

				return api
			},
			Input: &model.Config{
				ServiceSettings: model.ServiceSettings{
					EnableCommands: model.NewBool(true),
				},
			},
			ShouldReturn: true,
			ShouldError:  false,
		},
		"different configurations": {
			SetupAPI: func(api *plugintest.API) *plugintest.API {
				api.On("GetConfig").Return(&model.Config{})

				return api
			},
			Input: &model.Config{
				ServiceSettings: model.ServiceSettings{
					EnableCommands: model.NewBool(true),
				},
			},
			ShouldReturn: false,
			ShouldError:  false,
		},
	} {
		t.Run(name, func(t *testing.T) {
			api := test.SetupAPI(&plugintest.API{})
			defer api.AssertExpectations(t)

			p := &plugin.HelpersImpl{}
			p.API = api

			ok, err := p.CheckRequiredServerConfiguration(test.Input)

			if !ok {
				assert.False(t, ok)
			}
			if test.ShouldError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
