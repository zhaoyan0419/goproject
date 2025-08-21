package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// 定义一个 Counter 指标
	counter := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_app_requests_total",
			Help: "Total number of requests",
		},
	)

	// 注册指标到默认的 Register
	prometheus.MustRegister(counter)

	// 模拟业务逻辑：增加计数器
	counter.Inc()

	// 暴露指标端点
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
