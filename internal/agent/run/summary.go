package run

import (
	"bufio"
	_ "embed"
	"fmt"
	builtin "github.com/aaronchen2k/deeptest/internal/pkg/buildin"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
)

func newOutSummary() *Summary {
	platForm := &Platform{
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%v-%v", runtime.GOOS, runtime.GOARCH),
	}
	return &Summary{
		Success: true,
		Stat:    &Stat{},
		Time: &TestScenarioTime{
			StartAt: time.Now(),
		},
		Platform: platForm,
	}
}

// Summary stores tests summary for current task execution, maybe include one or multiple testcases
type Summary struct {
	Success  bool                   `json:"success" yaml:"success"`
	Stat     *Stat                  `json:"stat" yaml:"stat"`
	Time     *TestScenarioTime      `json:"time" yaml:"time"`
	Platform *Platform              `json:"platform" yaml:"platform"`
	Details  []*TestScenarioSummary `json:"details" yaml:"details"`
	rootDir  string
}

func (s *Summary) appendScenarioSummary(scenarioSummary *TestScenarioSummary) {
	s.Success = s.Success && scenarioSummary.Success
	s.Stat.TestScenarios.Total += 1
	s.Stat.TestStages.Total += len(scenarioSummary.Records)

	if scenarioSummary.Success {
		s.Stat.TestScenarios.Success += 1
	} else {
		s.Stat.TestScenarios.Fail += 1
	}

	s.Stat.TestStages.Successes += scenarioSummary.Stat.Successes
	s.Stat.TestStages.Failures += scenarioSummary.Stat.Failures
	s.Details = append(s.Details, scenarioSummary)
	s.Success = s.Success && scenarioSummary.Success

	// specify output reports dir
	if len(s.Details) == 1 {
		s.rootDir = scenarioSummary.RootDir
	} else if s.rootDir != scenarioSummary.RootDir {
		// if multiple testscenarios have different root path, use current working dir
		s.rootDir, _ = os.Getwd()
	}
}

func (s *Summary) genHTMLReport() error {
	reportsDir := filepath.Join(s.rootDir, resultsDir)
	err := builtin.EnsureFolderExists(reportsDir)
	if err != nil {
		return err
	}

	reportPath := filepath.Join(reportsDir, fmt.Sprintf("report-%v.html", s.Time.StartAt.Unix()))
	file, err := os.OpenFile(reportPath, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		log.Error().Err(err).Msg("open file failed")
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	tmpl := template.Must(template.New("report").Parse(reportTemplate))
	err = tmpl.Execute(writer, s)
	if err != nil {
		log.Error().Err(err).Msg("execute applies a parsed template to the specified data object failed")
		return err
	}
	err = writer.Flush()
	if err == nil {
		log.Info().Str("path", reportPath).Msg("generate HTML report")
	} else {
		log.Error().Str("path", reportPath).Msg("generate HTML report failed")
	}
	return err
}

func (s *Summary) genSummary() error {
	reportsDir := filepath.Join(s.rootDir, resultsDir)
	err := builtin.EnsureFolderExists(reportsDir)
	if err != nil {
		return err
	}

	summaryPath := filepath.Join(reportsDir, fmt.Sprintf("summary-%v.json", s.Time.StartAt.Unix()))
	err = builtin.Dump2JSON(s, summaryPath)
	if err != nil {
		return err
	}
	return nil
}

// // go:embed internal/scaffold/templates/report/template.html
var reportTemplate string

const resultsDir = "reports"

type Stat struct {
	TestScenarios TestScenarioStat `json:"testScenarios" yaml:"testScenarios"`
	TestStages    TestStageStat    `json:"testStages" yaml:"testStages"`
}

type TestScenarioStat struct {
	Total   int `json:"total" yaml:"total"`
	Success int `json:"success" yaml:"success"`
	Fail    int `json:"fail" yaml:"fail"`
}

type TestStageStat struct {
	Total     int `json:"total" yaml:"total"`
	Successes int `json:"successes" yaml:"successes"`
	Failures  int `json:"failures" yaml:"failures"`
}

type TestScenarioTime struct {
	StartAt  time.Time `json:"startAt,omitempty" yaml:"startAt,omitempty"`
	Duration float64   `json:"duration,omitempty" yaml:"duration,omitempty"`
}

type Platform struct {
	DeepTestVersion string `json:"deepTestVersion" yaml:"deepTestVersion"`
	GoVersion       string `json:"goVersion" yaml:"goVersion"`
	Platform        string `json:"platform" yaml:"platform"`
}

// TestScenarioSummary stores tests summary for one testScenario
type TestScenarioSummary struct {
	Name       string             `json:"name" yaml:"name"`
	Success    bool               `json:"success" yaml:"success"`
	ScenarioId string             `json:"scenarioId,omitempty" yaml:"scenarioId,omitempty"` // TODO
	Stat       *TestStageStat     `json:"stat" yaml:"stat"`
	Time       *TestScenarioTime  `json:"time" yaml:"time"`
	InOut      *TestScenarioInOut `json:"inOut" yaml:"inOut"`
	Log        string             `json:"log,omitempty" yaml:"log,omitempty"` // TODO
	Records    []*StageResult     `json:"records" yaml:"records"`
	RootDir    string             `json:"rootDir" yaml:"rootDir"`
}

type TestScenarioInOut struct {
	ConfigVars map[string]interface{} `json:"configVars" yaml:"configVars"`
	ExportVars map[string]interface{} `json:"exportVars" yaml:"exportVars"`
}

func newSessionData() *SessionData {
	return &SessionData{
		Success: false,
		ReqResp: &ReqResp{},
	}
}

type SessionData struct {
	Success    bool                `json:"success" yaml:"success"`
	ReqResp    *ReqResp            `json:"reqResp" yaml:"reqResp"`
	Address    *Address            `json:"address,omitempty" yaml:"address,omitempty"` // TODO
	Validators []*ValidationResult `json:"validators,omitempty" yaml:"validators,omitempty"`
}

type ReqResp struct {
	Request  interface{} `json:"request" yaml:"request"`
	Response interface{} `json:"response" yaml:"response"`
}

type Address struct {
	ClientIP   string `json:"clientIp,omitempty" yaml:"clientIp,omitempty"`
	ClientPort string `json:"clientPort,omitempty" yaml:"clientIport,omitempty"`
	ServerIP   string `json:"serverIp,omitempty" yaml:"serverIp,omitempty"`
	ServerPort string `json:"serverPort,omitempty" yaml:"serverPort,omitempty"`
}

type ValidationResult struct {
	Validator
	CheckValue  interface{} `json:"checkValue" yaml:"checkValue"`
	CheckResult string      `json:"checkResult" yaml:"checkResult"`
}

func newSummary() *TestScenarioSummary {
	return &TestScenarioSummary{
		Success: true,
		Stat:    &TestStageStat{},
		Time:    &TestScenarioTime{},
		InOut:   &TestScenarioInOut{},
		Records: []*StageResult{},
	}
}
