package environment

import (
	"flag"
)

const (
	AgentDefaultSendingAdressURL = "localhost:8080"
	AgentDefaultPollInterval     = 2
	AgentDefaultReportInterval   = 10

	AgentEnvSendingAdressURL = "ADDRESS"
	AgentEnvPollInterval     = "POLL_INTERVAL"
	AgentEnvReportInterval   = "REPORT_INTERVAL"
)

type AgentConfig struct {
	SendingAdress   string
	PollIntervalS   int
	ReportIntervalS int
}

func NewAgentConfig() AgentConfig {

	sendingAdressFlag := flag.String("a", AgentDefaultSendingAdressURL, "SENDING_ADRESS")
	pollIntervalFlag := flag.Int("p", AgentDefaultPollInterval, "POLL_INTERVAL")
	reportIntervalFlag := flag.Int("r", AgentDefaultReportInterval, "REPORT_INTERVAL")

	flag.Parse()

	ac := AgentConfig{
		SendingAdress:   GetEnvString(AgentEnvSendingAdressURL, *sendingAdressFlag),
		PollIntervalS:   GetEnvDuration(AgentEnvPollInterval, *pollIntervalFlag),
		ReportIntervalS: GetEnvDuration(AgentEnvReportInterval, *reportIntervalFlag),
	}

	return ac
}
