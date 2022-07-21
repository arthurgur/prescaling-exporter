package exporter

import (
	"github.com/bedrockstreaming/prescaling-exporter/pkg/prescaling"
	"github.com/bedrockstreaming/prescaling-exporter/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type prescalingCollector struct {
	desiredMetrics *prometheus.Desc
	prescaling     prescaling.IPrescaling
}

func NewPrescalingCollector(p prescaling.IPrescaling) prometheus.Collector {
	return &prescalingCollector{
		desiredMetrics: prometheus.NewDesc(
			"min_replica",
			"Number of pod desired for prescale",
			[]string{"project", "deployment", "namespace"},
			nil,
		),
		prescaling: p,
	}
}

func (collector *prescalingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desiredMetrics
}

func (collector *prescalingCollector) Collect(ch chan<- prometheus.Metric) {
	log.Info("Collect")
	hpaList := collector.prescaling.GetHpa()
	if len(hpaList) == 0 {
		log.Error("error - no prescaling hpa configuration found")
		return
	}

	currentPrescalingEvent, err := collector.prescaling.GetEventService().Current()

	now := collector.prescaling.GetEventService().GetClock().Now()
	for _, hpa := range hpaList {
		multiplier := 1
		if err == nil && currentPrescalingEvent.StartTime != "" && currentPrescalingEvent.EndTime != "" {
			hpa.Start, _ = utils.SetTime(currentPrescalingEvent.StartTime, now)
			hpa.End, _ = utils.SetTime(currentPrescalingEvent.EndTime, now)
			multiplier = currentPrescalingEvent.Multiplier
		}

		collector.addDataToMetrics(ch, multiplier, hpa)
	}

}

func (collector *prescalingCollector) addDataToMetrics(ch chan<- prometheus.Metric, multiplier int, hpa prescaling.Hpa) {
	eventInRangeTime := utils.InRangeTime(hpa.Start, hpa.End, collector.prescaling.GetEventService().GetClock().Now())
	desiredScalingType := prescaling.DesiredScaling(eventInRangeTime, multiplier, hpa.Replica, hpa.CurrentReplicas)
	metric := prometheus.MustNewConstMetric(collector.desiredMetrics, prometheus.GaugeValue, float64(desiredScalingType), hpa.Project, hpa.Deployment, hpa.Namespace)
	ch <- metric
}
