package util

import (
	"oc_res_control/conf"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var GlobalStat  *Statistics

type Statistics struct {
	prefix      string
	counterTags map[string]*prometheus.CounterVec
	counterChan chan *counterValue
	summaryTags map[string]*prometheus.SummaryVec
	summaryChan chan *summaryValue
}

const (
	defaultStatsPrefix     = "filter_"
	StatQueryGcid          = "query_gcid_status_api"
	StatQueryKeyword       = "query_keyword_status_api"
	StatQueryUncheck       = "query_uncheck_api"
	StatAddHumanCheck      = "add_human_check_api"
)

func (st *Statistics) Init(cfg *conf.Config) {
	st.prefix = defaultStatsPrefix
	st.counterTags = make(map[string]*prometheus.CounterVec)
	st.counterChan = make(chan *counterValue, 1024)
	st.summaryTags = make(map[string]*prometheus.SummaryVec)
	st.summaryChan = make(chan *summaryValue, 1024)

	go st.collect()
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		panic(http.ListenAndServe(cfg.Stat.PrometheusHttpAddr, nil))
	}()
}

func (st *Statistics) collect() {
	for {
		select {
		case c := <-st.counterChan:
			counter, exist := st.counterTags[c.tag]
			if !exist {
				counter = prometheus.NewCounterVec(
					prometheus.CounterOpts{
						Name: st.prefix + c.tag + "_total",
						Help: st.prefix + c.tag + "_total"},
					[]string{"code"})
				st.counterTags[c.tag] = counter
				prometheus.MustRegister(counter)
			}
			counter.WithLabelValues(c.code).Add(float64(c.num))
		case s := <-st.summaryChan:
			summary, exist := st.summaryTags[s.tag]
			if !exist {
				summary = prometheus.NewSummaryVec(
					prometheus.SummaryOpts{
						Name:       st.prefix + s.tag + "_duration_seconds",
						Help:       st.prefix + s.tag + "_duration_seconds",
						Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001}},
					[]string{"code"})
				st.summaryTags[s.tag] = summary
				prometheus.MustRegister(summary)
			}
			summary.WithLabelValues(s.code).Observe(float64(s.num))
		}
	}
}

type counterValue struct {
	tag  string
	num  int
	code string
}

type summaryValue struct {
	tag  string
	num  float64
	code string
}

func (s *Statistics) AddCount(statsTag, code string, num int) {
	s.counterChan <- &counterValue{tag: statsTag, num: num, code: code}
}

func (s *Statistics) MonitorCost(statsTag, code string, costMs float64) {
	s.summaryChan <- &summaryValue{tag: statsTag, num: costMs, code: code}
}
