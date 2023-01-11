package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

// Metrics - сущность метрики
type Metrics struct {
	MyCustomMetric *prometheus.CounterVec
}

var once sync.Once
var metr *Metrics

// Instance - (одиночка) получить экземпляр метрик
func Instance() *Metrics {
	once.Do(func() {
		metr = &Metrics{}
		metr.initMetrics()
	})

	return metr
}

func (m *Metrics) initMetrics() {
	// Для чего теги
	// Например идет запрос и мы детализируем
	// теги method=get handler=fetchUser code=200
	// с помощью тегов мы можем понимать сколько таких запросов было с кодом 200
	// или с кодом 401 - (возможно отклеелся сервис авторизации)
	m.MyCustomMetric = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "osoc_some_count",
			Help: "Пример бизнес метрики",
		},
		[]string{"someInfo"},
	)

	prometheus.MustRegister(m.MyCustomMetric)
}
