package request

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
)

// Run starts to run API test with default configs.
func Run(testcases ...ITestCase) error {
	t := &testing.T{}
	return NewRunner(t).SetRequestsLogOn().Run(testcases...)
}

// NewRunner constructs a new runner instance.
func NewRunner(t *testing.T) *MainRunner {
	if t == nil {
		t = &testing.T{}
	}
	jar, _ := cookiejar.New(nil)
	return &MainRunner{
		t:             t,
		failfast:      true,
		genHTMLReport: false,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Jar:     jar, // insert response cookies into request
			Timeout: 120 * time.Second,
		},
		http2Client: &http.Client{
			Transport: &http2.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			Timeout: 120 * time.Second,
		},
	}
}

// SetClientTransport configures transport of http client for high concurrency load testing
func (r *MainRunner) SetClientTransport(maxConns int, disableKeepAlive bool, disableCompression bool) *MainRunner {
	log.Info().
		Int("maxConns", maxConns).
		Bool("disableKeepAlive", disableKeepAlive).
		Bool("disableCompression", disableCompression).
		Msg("[init] SetClientTransport")

	r.httpClient.Transport = &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		DialContext:         (&net.Dialer{}).DialContext,
		MaxIdleConns:        0,
		MaxIdleConnsPerHost: maxConns,
		DisableKeepAlives:   disableKeepAlive,
		DisableCompression:  disableCompression,
	}
	r.http2Client.Transport = &http2.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: disableCompression,
	}
	r.wsDialer.EnableCompression = !disableCompression
	return r
}

// SetFailfast configures whether to stop running when one step fails.
func (r *MainRunner) SetFailfast(failfast bool) *MainRunner {
	log.Info().Bool("failfast", failfast).Msg("[init] SetFailfast")
	r.failfast = failfast
	return r
}

// SetRequestsLogOn turns on request & response details logging.
func (r *MainRunner) SetRequestsLogOn() *MainRunner {
	log.Info().Msg("[init] SetRequestsLogOn")
	r.requestsLogOn = true
	return r
}

// SetHTTPStatOn turns on HTTP latency stat.
func (r *MainRunner) SetHTTPStatOn() *MainRunner {
	log.Info().Msg("[init] SetHTTPStatOn")
	r.httpStatOn = true
	return r
}

// SetPluginLogOn turns on plugin logging.
func (r *MainRunner) SetPluginLogOn() *MainRunner {
	log.Info().Msg("[init] SetPluginLogOn")
	r.pluginLogOn = true
	return r
}

// SetPython3Venv specifies python3 venv.
func (r *MainRunner) SetPython3Venv(venv string) *MainRunner {
	log.Info().Str("venv", venv).Msg("[init] SetPython3Venv")
	r.venv = venv
	return r
}

// SetProxyUrl configures the proxy URL, which is usually used to capture HTTP packets for debugging.
func (r *MainRunner) SetProxyUrl(proxyUrl string) *MainRunner {
	log.Info().Str("proxyUrl", proxyUrl).Msg("[init] SetProxyUrl")
	p, err := url.Parse(proxyUrl)
	if err != nil {
		log.Error().Err(err).Str("proxyUrl", proxyUrl).Msg("[init] invalid proxyUrl")
		return r
	}
	r.httpClient.Transport = &http.Transport{
		Proxy:           http.ProxyURL(p),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	r.wsDialer.Proxy = http.ProxyURL(p)
	return r
}

// SetTimeout configures global timeout in seconds.
func (r *MainRunner) SetTimeout(timeout time.Duration) *MainRunner {
	log.Info().Float64("timeout(seconds)", timeout.Seconds()).Msg("[init] SetTimeout")
	r.httpClient.Timeout = timeout
	return r
}

// SetSaveTests configures whether to save summary of tests.
func (r *MainRunner) SetSaveTests(saveTests bool) *MainRunner {
	log.Info().Bool("saveTests", saveTests).Msg("[init] SetSaveTests")
	r.saveTests = saveTests
	return r
}

// GenHTMLReport configures whether to gen html report of api tests.
func (r *MainRunner) GenHTMLReport() *MainRunner {
	log.Info().Bool("genHTMLReport", true).Msg("[init] SetgenHTMLReport")
	r.genHTMLReport = true
	return r
}

// Run starts to execute one or multiple testcases.
func (r *MainRunner) Run(testcases ...ITestCase) (err error) {
	return
}

// NewSessionRunner creates a new session runner for testcase.
// each testcase has its own session runner
func (r *MainRunner) NewSessionRunner(testcase *TestCase) (*SessionRunner, error) {
	runner, err := r.newCaseRunner(testcase)
	if err != nil {
		return nil, err
	}

	sessionRunner := &SessionRunner{
		ScenarioRunner: runner,
	}

	return sessionRunner, nil
}

func (r *MainRunner) newCaseRunner(testcase *TestCase) (*ScenarioRunner, error) {
	runner := &ScenarioRunner{
		testCase:  testcase,
		hrpRunner: r,
		parser:    newParser(),
	}

	return runner, nil
}

type MainRunner struct {
	t             *testing.T
	failfast      bool
	httpStatOn    bool
	requestsLogOn bool
	pluginLogOn   bool
	venv          string
	saveTests     bool
	genHTMLReport bool
	httpClient    *http.Client
	http2Client   *http.Client
	wsDialer      *websocket.Dialer
}

type ScenarioRunner struct {
	testCase           *TestCase
	hrpRunner          *MainRunner
	parser             *Parser
	parsedConfig       *TConfig
	parametersIterator *ParametersIterator
	rootDir            string // project root dir
}
