package log

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"sort"
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
var latencies []float64
var kubectl Kubectl

type Kubectl struct {
	ClientVersion    Version
	KustomizeVersion string
}

type Version struct {
	Major        string
	Minor        string
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

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

// Logger writes k6 metric samples to stdout.
type Logger struct {
	out io.Writer
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
		// fmt.Fprintf(l.out, "%s [%s]\n", all[0].GetTime().Format(time.RFC3339Nano), metricKeyValues(all))
		metricKeyValues(all)
	}
}

// Stop finalizes any tasks in progress, closes network connections, etc.
func (*Logger) Stop() error {
	stopTime = time.Now()
	duration := stopTime.Sub(startTime)
	createYamlFile(duration)
	return nil
}

// init is called by the Go runtime at application startup.
func init() {
	output.RegisterExtension("smp", new)
}

// metricKeyValues returns a string of key-value pairs for all metrics in the sample.
func metricKeyValues(samples []metrics.Sample) string {
	names := make([]string, 0, len(samples))
	for _, sample := range samples {
		// names = append(names, fmt.Sprintf("%s=%v", sample.Metric.Name, sample.Value))
		if sample.Metric.Name == "iterations" {
			requests++
		}
		if sample.Metric.Name == "http_req_duration" {
			latencies = append(latencies, sample.Value)
		}
	}
	return strings.Join(names, ", ")
}

// New returns a new instance of Logger.
func new(params output.Params) (output.Output, error) {
	return &Logger{params.StdOut}, nil
}

func kubectlVersion() string {

	out, err := exec.Command("kubectl", "version", "--output=json").Output()

	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(out, &kubectl)

	return kubectl.ClientVersion.GitVersion
}

func latenciesInMilliseconds(p90Indexlatencies []float64) (min float64, max float64, average float64, p50 float64, p90 float64, p99 float64) {
	min = latencies[0]
	max = latencies[0]
	sum := 0.0
	sort.Float64s(latencies)
	for _, latency := range latencies {
		if latency > max {
			max = latency
		}
		if latency < min {
			min = latency
		}
		sum += latency
	}

	p50Index := math.Round(0.5 * (float64)(len(latencies)))
	p90Index := math.Round(0.9 * (float64)(len(latencies)))
	p99Index := math.Round(0.99 * (float64)(len(latencies)))

	if int(p90Index) == len(latencies) {
		p90Index -= 1
	}

	if int(p99Index) == len(latencies) {
		p99Index -= 1
	}

	return min, max, sum / (float64)(len(latencies)), latencies[int(p50Index)], latencies[int(p90Index)], latencies[int(p99Index)]
}

func createYamlFile(duration time.Duration) {

	id := uuid.New()
	requestsPerSeconds := float64(requests) / duration.Seconds()

	min, max, average, p50, p90, p99 := latenciesInMilliseconds(latencies)

	smp := ServiceMeshPerformance{
		Start_time:     startTime,
		End_time:       stopTime,
		Exp_uuid:       id,
		Endpoint_url:   "http://localhost:8080",
		Load_generator: "k6",
		Env: EnvStruct{
			Kubernetes: kubectlVersion(),
		},
		Client: ClientStruct{
			Connections: 1,
			Rps:         requestsPerSeconds,
			Latencies_ms: LatencyStruct{
				Min:     min,
				Average: average,
				P50:     p50,
				P90:     p90,
				P99:     p99,
				Max:     max,
			},
		},
		Metrics: map[string]string{},
	}

	yamlData, err := yaml.Marshal(&smp)

	if err != nil {
		fmt.Printf("Error while Marshaling. %v", err)
	}

	fileName := createFileNameAsIso8601()
	err = os.WriteFile(fileName, yamlData, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}

}

func createFileNameAsIso8601() string {
	now := time.Now()

	day := strconv.Itoa(now.Day())
	month := strconv.Itoa(int(now.Month()))

	if len(day) == 1 {
		day = "0" + day
	}

	if len(month) == 1 {
		month = "0" + month
	}

	return strconv.Itoa(now.Year()) + "-" + month + "-" + day + "T" + strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute()) + ":" + strconv.Itoa(now.Second()) + "+01:00.yaml"
}
