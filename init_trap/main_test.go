package init

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tigren/go-tour/init_trap/api"
	"github.com/tigren/go-tour/init_trap/config"
)

func TestUseApi(t *testing.T) {
	// init config
	config.Init()

	assert.True(t, api.IsEnabled())
}
