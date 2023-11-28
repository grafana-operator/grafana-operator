package grafana

import (
	"fmt"
	"testing"

	config2 "github.com/grafana-operator/grafana-operator/v5/controllers/config"

	"github.com/stretchr/testify/assert"
)

func Test_setGrafanaImage(t *testing.T) {
	expectedDeploymentImage := fmt.Sprintf("%s:%s", config2.GrafanaImage, config2.GrafanaVersion)

	assert.Equal(t, expectedDeploymentImage, setGrafanaImage())
}

func Test_setGrafanaImage_withEnvironmentOverride(t *testing.T) {
	expectedDeploymentImage := "I want this grafana image"
	t.Setenv("RELATED_IMAGE_GRAFANA", expectedDeploymentImage)

	assert.Equal(t, expectedDeploymentImage, setGrafanaImage())
}
