package elasticsearch

import (
	es "github.com/elastic/go-elasticsearch/v9"
	"github.com/tinh-tinh/tinhtinh/v2/core"
)

const ELASTICSEARCH = "ELASTICSEARCH"

func ForRoot(config es.Config) core.Modules {
	return func(module core.Module) core.Module {
		esModule := module.New(core.NewModuleOptions{})

		client, err := es.NewClient(config)
		if err != nil {
			panic(err)
		}

		esModule.NewProvider(core.ProviderOptions{
			Name:  ELASTICSEARCH,
			Value: client,
		})
		esModule.Export(ELASTICSEARCH)

		return esModule
	}
}

type ConfigFactory func(ref core.RefProvider) es.Config

func ForRootFactory(factory ConfigFactory) core.Modules {
	return func(module core.Module) core.Module {
		esModule := module.New(core.NewModuleOptions{})

		config := factory(module)
		client, err := es.NewClient(config)
		if err != nil {
			panic(err)
		}

		esModule.NewProvider(core.ProviderOptions{
			Name:  ELASTICSEARCH,
			Value: client,
		})
		esModule.Export(ELASTICSEARCH)

		return esModule
	}
}

func InjectClient(module core.RefProvider) *es.Client {
	client, ok := module.Ref(ELASTICSEARCH).(*es.Client)
	if !ok {
		return nil
	}
	return client
}
