package agent

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MediaCrush/DataCrush/link"
)

type CPUWatch struct {
	link      link.Link
	threshold float64
	source    string
	interval  time.Duration
}

var (
	NO_LINK = errors.New("Link is not ready")
)

func NewCPUWatchAgent(link link.Link, host string, maxload float64, interval time.Duration) (*CPUWatch, error) {
	if !link.Ready() {
		return nil, NO_LINK
	}

	agent := &CPUWatch{
		link:     link,
		source:   host,
		interval: interval,
	}

	cpus, err := link.Run("grep -c processor /proc/cpuinfo")
	if err != nil {
		return nil, err
	}

	cores, err := strconv.ParseFloat(clean(cpus), 64)
	agent.threshold = cores * maxload
	return agent, err
}

func (s *CPUWatch) Run(events chan<- Event) {
	for {
		result, _ := s.link.Run("cat /proc/loadavg | awk '{print $1}'")
		load, _ := strconv.ParseFloat(clean(result), 64)

		if load < s.threshold {
			events <- Event{
				Source:  s.source,
				Payload: fmt.Sprintf("%f below %f", load, s.threshold),
			}
		} else {
			events <- Event{
				Source:  s.source,
				Payload: fmt.Sprintf("%f >above< %f!", load, s.threshold),
			}
		}

		time.Sleep(time.Second * s.interval)
	}
}

func (s *CPUWatch) Stop() {
}
