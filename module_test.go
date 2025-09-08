package elasticsearch_test

import (
	"testing"

	es "github.com/elastic/go-elasticsearch/v9"
	"github.com/stretchr/testify/require"
	"github.com/tinh-tinh/elasticsearch"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

func Test_Connect(t *testing.T) {
	appModule := core.NewModule(core.NewModuleOptions{
		Imports: []core.Modules{
			elasticsearch.ForRoot(es.Config{
				Addresses: []string{
					"http://localhost:9200",
				},
			}),
		},
	})

	es := elasticsearch.InjectClient(appModule)
	require.NotNil(t, es)
}
