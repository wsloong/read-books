// 链路追踪器
package tracer

import (
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	// 为jaeger client的配置项
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ // 固定采样、对所有数据进行采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true, //  是否启动 LoggingReporter
			BufferFlushInterval: time.Second,
			LocalAgentHostPort:  agentHostPort, // 上报的agent地址
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置全局的tracer对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
