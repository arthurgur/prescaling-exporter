package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"github.com/arthurgur/prescaling-exporter/pkg/prescaling"
	"github.com/arthurgur/prescaling-exporter/pkg/utils"
)

type prescalingCollector struct {
	prescaleMetrics   *prometheus.Desc
	minMetrics        *prometheus.Desc
	multiplierMetrics *prometheus.Desc
	prescaling        prescaling.IPrescaling
}

func NewPrescalingCollector(p prescaling.IPrescaling) prometheus.Collector {
	return &prescalingCollector{
		prescaleMetrics: prometheus.NewDesc(
			"prescaling_metric",
			"Number used for prescale application",
			[]string{"project", "deployment", "namespace"},
			nil,
		), minMetrics: prometheus.NewDesc(
			"prescaling_min_replica",
			"Number of pod desired for prescale",
			[]string{"project", "deployment", "namespace"},
			nil,
		), multiplierMetrics: prometheus.NewDesc(
			"prescaling_multiplier",
			"Multiplying factor of min replica",
			[]string{},
			nil,
		),
		prescaling: p,
	}
}

func (collector *prescalingCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.prescaleMetrics
	ch <- collector.minMetrics
	ch <- collector.multiplierMetrics
}

func (collector *prescalingCollector) Collect(ch chan<- prometheus.Metric) {
	var multiplier int = 1
	log.Info("Collect")
	hpaList := collector.prescaling.GetHpa()
	if len(hpaList) == 0 {
		log.Error("error - no prescaling hpa configuration found")
		return
	}

	currentPrescalingEvent, err := collector.prescaling.GetEventService().Current()

	now := collector.prescaling.GetEventService().GetClock().Now()
	for _, hpa := range hpaList {
		if err == nil && currentPrescalingEvent.StartTime != "" && currentPrescalingEvent.EndTime != "" {
			hpa.Start, _ = utils.SetTime(currentPrescalingEvent.StartTime, now)
			hpa.End, _ = utils.SetTime(currentPrescalingEvent.EndTime, now)
			multiplier = currentPrescalingEvent.Multiplier
		}

		collector.addHpaDataToMetrics(ch, multiplier, hpa)
	}

	collector.addDataToMetrics(ch, multiplier)

}

func (collector *prescalingCollector) addHpaDataToMetrics(ch chan<- prometheus.Metric, multiplier int, hpa prescaling.Hpa) {
	eventInRangeTime := utils.InRangeTime(hpa.Start, hpa.End, collector.prescaling.GetEventService().GetClock().Now())
	desiredScalingType := prescaling.DesiredScaling(eventInRangeTime, multiplier, hpa.Replica, hpa.CurrentReplicas)
	prescaleMetric := prometheus.MustNewConstMetric(collector.prescaleMetrics, prometheus.GaugeValue, float64(desiredScalingType), hpa.Project, hpa.Deployment, hpa.Namespace)
	minMetric := prometheus.MustNewConstMetric(collector.minMetrics, prometheus.GaugeValue, float64(hpa.Replica), hpa.Project, hpa.Deployment, hpa.Namespace)
	ch <- prescaleMetric
	ch <- minMetric
}

func (collector *prescalingCollector) addDataToMetrics(ch chan<- prometheus.Metric, multiplier int) {
	multiplierMetric := prometheus.MustNewConstMetric(collector.multiplierMetrics, prometheus.GaugeValue, float64(multiplier))
	ch <- multiplierMetric
}
