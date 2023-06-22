package functions

import (
	"os"
	"testing"

	"github.com/rancher/rancher/tests/framework/pkg/config"
	set "github.com/josh-diamond/rancher-terratest-framework/functions/set"
	"github.com/josh-diamond/rancher-terratest-framework/terratest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BuildModule(t *testing.T) (bool, error) {
	clusterConfig := new(terratest.TerratestConfig)
	config.LoadConfig("terratest", clusterConfig)

	keyPath := set.SetKeyPath()

	result, err := set.SetConfigTF(t, clusterConfig.KubernetesVersion, clusterConfig.Nodepools)
	require.NoError(t, err)
	assert.Equal(t, true, result)

	// Log module
	module, err := os.ReadFile(keyPath + "/main.tf")

	if err != nil {
		t.Logf("Failed to read/grab main.tf file contents. Error: %v", err)
		return false, err
	}

	t.Log(string(module))

	return true, nil
}
