package requestmetrics

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics - сущность метрики
type metrics struct {
	prefix          string
	requestsTotal   *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

// todo: consider keeping these vars in middleware, as we do not support multiple instances of this middleware.
var once sync.Once
var metr *metrics

func New(prefix string) gin.HandlerFunc {
	once.Do(func() {
		metr = &metrics{
			prefix: prefix,
		}
		metr.initMetrics()
	})

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := time.Since(start)
		metr.requestsTotal.With(prometheus.Labels{
			"method":  c.Request.Method,
			"handler": c.FullPath(),
			"code":    status,
		}).Inc()

		metr.requestDuration.With(prometheus.Labels{
			"method":  c.Request.Method,
			"handler": c.FullPath(),
			"code":    status,
		}).Observe(elapsed.Seconds())
	}
}

func (m *metrics) initMetrics() {
	// Для чего теги
	// Например идет запрос и мы детализируем
	// теги method=get handler=fetchUser code=200
	// с помощью тегов мы можем понимать сколько таких запросов было с кодом 200
	// или с кодом 401 - (возможно отклеился сервис авторизации)
	m.requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: fmt.Sprintf("%s_requests_total", m.prefix),
			Help: "Кол-во запросов к сервису",
		},
		[]string{"method", "handler", "code"},
	)

	m.requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    fmt.Sprintf("%s_requests_duration", m.prefix),
			Help:    "Время обработки запроса",
			Buckets: prometheus.LinearBuckets(0.020, 0.020, 5),
		},
		[]string{"method", "handler", "code"},
	)

	prometheus.MustRegister(m.requestsTotal, m.requestDuration)
	m.setDefaultValue()
}

func (m *metrics) setDefaultValue() {
	m.requestsTotal.With(prometheus.Labels{
		"method":  "",
		"handler": "",
		"code":    "",
	}).Add(0)
}

type ContextCaller interface {
	HandlerName() string
}
