package environment

import (
	"flag"
	"time"
)

const (
	AgentDefaultSendingAdressURL = "localhost:8080"
	AgentDefaultPollInterval     = 2 * time.Second
	AgentDefaultReportInterval   = 10 * time.Second

	AgentEnvSendingAdressURL = "ADDRESS"
	AgentEnvPollInterval     = "POLL_INTERVAL"
	AgentEnvReportInterval   = "REPORT_INTERVAL"
)

type AgentConfig struct {
	SendingAdress   string
	PollIntervalS   time.Duration
	ReportIntervalS time.Duration
}

func NewAgentConfig() AgentConfig {

	sendingAdressFlag := flag.String("a", AgentDefaultSendingAdressURL, "SENDING_ADRESS")
	pollIntervalFlag := flag.Duration("p", AgentDefaultPollInterval, "POLL_INTERVAL")
	reportIntervalFlag := flag.Duration("r", AgentDefaultReportInterval, "REPORT_INTERVAL")

	flag.Parse()

	ac := AgentConfig{
		SendingAdress:   GetEnvString(AgentEnvSendingAdressURL, *sendingAdressFlag),
		PollIntervalS:   GetEnvDuration(AgentEnvPollInterval, *pollIntervalFlag),
		ReportIntervalS: GetEnvDuration(AgentEnvReportInterval, *reportIntervalFlag),
	}

	return ac
}
