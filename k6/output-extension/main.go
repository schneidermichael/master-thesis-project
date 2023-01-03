package log

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.k6.io/k6/metrics"
	"go.k6.io/k6/output"
	"gopkg.in/yaml.v3"
)

var requests int = 0
var startTime time.Time
var stopTime time.Time

type LatencyStruct struct {
	Min     float64 `yaml:"min"`
	Average float64 `yaml:"average"`
	P50     float64 `yaml:"p50"`
	P90     float64 `yaml:"p90"`
	P99     float64 `yaml:"p99"`
	Max     float64 `yaml:"max"`
}

type EnvStruct struct {
	Kubernetes string `yaml:"kubernetes"`
}

type ClientStruct struct {
	Connections  int           `yaml:"connections"`
	Rps          float64       `yaml:"rps"`
	Latencies_ms LatencyStruct `yaml:"latencies_ms"`
}

type ServiceMeshPerformance struct {
	Start_time     time.Time         `yaml:"start_time"`
	End_time       time.Time         `yaml:"end_time"`
	Exp_uuid       uuid.UUID         `yaml:"exp_uuid"`
	Endpoint_url   string            `yaml:"endpoint_url"`
	Load_generator string            `yaml:"load_generator"`
	Env            EnvStruct         `yaml:"env"`
	Client         ClientStruct      `yaml:"client"`
	Metrics        map[string]string `yaml:"metrics"`
}

// init is called by the Go runtime at application startup.
func init() {
	output.RegisterExtension("smp", New)
}

// Logger writes k6 metric samples to stdout.
type Logger struct {
	out io.Writer
}

// New returns a new instance of Logger.
func New(params output.Params) (output.Output, error) {
	return &Logger{params.StdOut}, nil
}

// Description returns a short human-readable description of the output.
func (*Logger) Description() string {
	return "smp"
}

// Start initializes any state needed for the output, establishes network
// connections, etc.
func (o *Logger) Start() error {
	startTime = time.Now()
	return nil
}

// AddMetricSamples receives metric samples from the k6 Engine as they're emitted.
func (l *Logger) AddMetricSamples(samples []metrics.SampleContainer) {

	for _, sample := range samples {
		all := sample.GetSamples()
		fmt.Fprintf(l.out, "%s [%s]\n", all[0].GetTime().Format(time.RFC3339Nano), metricKeyValues(all))
	}
}

// metricKeyValues returns a string of key-value pairs for all metrics in the sample.
func metricKeyValues(samples []metrics.Sample) string {
	names := make([]string, 0, len(samples))
	for _, sample := range samples {
		names = append(names, fmt.Sprintf("%s=%v", sample.Metric.Name, sample.Value))
		if sample.Metric.Name == "iterations" {
			requests++
		}
	}
	return strings.Join(names, ", ")
}

// Stop finalizes any tasks in progress, closes network connections, etc.
func (*Logger) Stop() error {
	stopTime = time.Now()
	duration := stopTime.Sub(startTime)
	createYamlFile(duration)
	return nil
}

func createYamlFile(duration time.Duration) {

	id := uuid.New()
	requestsPerSeconds := float64(requests) / duration.Seconds()

	smp := ServiceMeshPerformance{
		Start_time:     startTime,
		End_time:       stopTime,
		Exp_uuid:       id,
		Endpoint_url:   "http://localhost:8080",
		Load_generator: "k6",
		Env: EnvStruct{
			Kubernetes: "v1.25.2",
		},
		Client: ClientStruct{
			Connections: 1,
			Rps:         requestsPerSeconds,
			Latencies_ms: LatencyStruct{
				Min:     0.00,
				Average: 0.00,
				P50:     0.00,
				P90:     0.00,
				P99:     0.00,
				Max:     0.00,
			},
		},
		Metrics: map[string]string{},
	}

	yamlData, err := yaml.Marshal(&smp)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fileName := createFileName()
	err = os.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}

}

func createFileName() string {
	now := time.Now()

	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	minute := strconv.Itoa(now.Minute())
	second := strconv.Itoa(now.Second())

	return year + "-" + month + "-" + day + " " + hour + ":" + minute + ":" + second + ".yaml"
}
