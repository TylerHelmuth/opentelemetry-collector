module go.opentelemetry.io/collector/processor/memorylimiterprocessor

go 1.23.0

require (
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v1.37.0
	go.opentelemetry.io/collector/component/componenttest v0.131.0
	go.opentelemetry.io/collector/confmap v1.37.0
	go.opentelemetry.io/collector/consumer v1.37.0
	go.opentelemetry.io/collector/consumer/consumererror v0.131.0
	go.opentelemetry.io/collector/consumer/consumertest v0.131.0
	go.opentelemetry.io/collector/consumer/xconsumer v0.131.0
	go.opentelemetry.io/collector/internal/memorylimiter v0.131.0
	go.opentelemetry.io/collector/internal/telemetry v0.131.0
	go.opentelemetry.io/collector/pdata v1.37.0
	go.opentelemetry.io/collector/pdata/pprofile v0.131.0
	go.opentelemetry.io/collector/pipeline v0.131.0
	go.opentelemetry.io/collector/pipeline/xpipeline v0.131.0
	go.opentelemetry.io/collector/processor v1.37.0
	go.opentelemetry.io/collector/processor/processorhelper v0.131.0
	go.opentelemetry.io/collector/processor/processorhelper/xprocessorhelper v0.131.0
	go.opentelemetry.io/collector/processor/processortest v0.131.0
	go.opentelemetry.io/collector/processor/xprocessor v0.131.0
	go.opentelemetry.io/otel v1.37.0
	go.opentelemetry.io/otel/metric v1.37.0
	go.opentelemetry.io/otel/sdk/metric v1.37.0
	go.opentelemetry.io/otel/trace v1.37.0
	go.uber.org/goleak v1.3.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/ebitengine/purego v0.8.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/knadh/koanf/providers/confmap v1.0.0 // indirect
	github.com/knadh/koanf/v2 v2.2.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20211012122336-39d0f177ccd0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/power-devops/perfstat v0.0.0-20210106213030-5aafc221ea8c // indirect
	github.com/shirou/gopsutil/v4 v4.25.6 // indirect
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/collector/component/componentstatus v0.131.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.37.0 // indirect
	go.opentelemetry.io/collector/pdata/testdata v0.131.0 // indirect
	go.opentelemetry.io/contrib/bridges/otelzap v0.12.0 // indirect
	go.opentelemetry.io/otel/log v0.13.0 // indirect
	go.opentelemetry.io/otel/sdk v1.37.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.yaml.in/yaml/v3 v3.0.4 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
	google.golang.org/grpc v1.74.2 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.opentelemetry.io/collector/processor => ../

replace go.opentelemetry.io/collector/processor/processortest => ../processortest

replace go.opentelemetry.io/collector/component => ../../component

replace go.opentelemetry.io/collector/component/componenttest => ../../component/componenttest

replace go.opentelemetry.io/collector/confmap => ../../confmap

replace go.opentelemetry.io/collector/pdata => ../../pdata

replace go.opentelemetry.io/collector/pdata/testdata => ../../pdata/testdata

replace go.opentelemetry.io/collector/consumer => ../../consumer

retract (
	v0.76.0 // Depends on retracted pdata v1.0.0-rc10 module, use v0.76.1
	v0.69.0 // Release failed, use v0.69.1
)

replace go.opentelemetry.io/collector/pdata/pprofile => ../../pdata/pprofile

replace go.opentelemetry.io/collector/consumer/xconsumer => ../../consumer/xconsumer

replace go.opentelemetry.io/collector/consumer/consumertest => ../../consumer/consumertest

replace go.opentelemetry.io/collector/component/componentstatus => ../../component/componentstatus

replace go.opentelemetry.io/collector/processor/xprocessor => ../xprocessor

replace go.opentelemetry.io/collector/pipeline => ../../pipeline

replace go.opentelemetry.io/collector/internal/memorylimiter => ../../internal/memorylimiter

replace go.opentelemetry.io/collector/internal/telemetry => ../../internal/telemetry

replace go.opentelemetry.io/collector/featuregate => ../../featuregate

replace go.opentelemetry.io/collector/consumer/consumererror => ../../consumer/consumererror

replace go.opentelemetry.io/collector/processor/processorhelper/xprocessorhelper => ../processorhelper/xprocessorhelper

replace go.opentelemetry.io/collector/processor/processorhelper => ../processorhelper

replace go.opentelemetry.io/collector/pipeline/xpipeline => ../../pipeline/xpipeline
