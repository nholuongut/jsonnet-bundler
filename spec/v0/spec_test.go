// Copyright 2018 Nho Luong DevOps
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"encoding/json"
	"testing"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const jsonJF = `{
  "dependencies": [
    {
      "name": "grafana-builder",
      "source": {
        "git": {
          "remote": "https://github.com/grafana/jsonnet-libs",
          "subdir": "grafana-builder"
        }
      },
      "version": "54865853ebc1f901964e25a2e7a0e4d2cb6b9648",
      "sum": "ELsYwK+kGdzX1mee2Yy+/b2mdO4Y503BOCDkFzwmGbE="
    },
    {
      "name": "prometheus-mixin",
      "source": {
        "git": {
          "remote": "https://github.com/prometheus/prometheus",
          "subdir": "documentation/prometheus-mixin"
        }
      },
      "version": "7c039a6b3b4b2a9d7c613ac8bd3fc16e8ca79684",
      "sum": "bVGOsq3hLOw2irNPAS91a5dZJqQlBUNWy3pVwM4+kIY="
    }
  ]
}`

func testData() JsonnetFile {
	f := JsonnetFile{Dependencies: orderedmap.NewOrderedMap[string, Dependency]()}

	f.Dependencies.Set("grafana-builder", Dependency{
		Name: "grafana-builder",
		Source: Source{
			GitSource: &GitSource{
				Remote: "https://github.com/grafana/jsonnet-libs",
				Subdir: "grafana-builder",
			},
		},
		Version: "54865853ebc1f901964e25a2e7a0e4d2cb6b9648",
		Sum:     "ELsYwK+kGdzX1mee2Yy+/b2mdO4Y503BOCDkFzwmGbE=",
	})
	f.Dependencies.Set("prometheus-mixin", Dependency{
		Name: "prometheus-mixin",
		Source: Source{
			GitSource: &GitSource{
				Remote: "https://github.com/prometheus/prometheus",
				Subdir: "documentation/prometheus-mixin",
			},
		},
		Version: "7c039a6b3b4b2a9d7c613ac8bd3fc16e8ca79684",
		Sum:     "bVGOsq3hLOw2irNPAS91a5dZJqQlBUNWy3pVwM4+kIY=",
	})

	return f
}

// TestUnmarshal checks that unmarshalling works
func TestUnmarshal(t *testing.T) {
	var dst JsonnetFile
	err := json.Unmarshal([]byte(jsonJF), &dst)
	require.NoError(t, err)
	assert.Equal(t, testData(), dst)
}

// TestMarshal checks that marshalling works
func TestMarshal(t *testing.T) {
	data, err := json.Marshal(testData())
	require.NoError(t, err)
	assert.JSONEq(t, jsonJF, string(data))
}

// TestRemarshal checks that unmarshalling a previously marshalled object yields
// the same object
func TestRemarshal(t *testing.T) {
	jf := testData()

	data, err := json.Marshal(jf)
	require.NoError(t, err)

	var dst JsonnetFile
	err = json.Unmarshal(data, &dst)
	require.NoError(t, err)

	assert.Equal(t, jf, dst)
}
