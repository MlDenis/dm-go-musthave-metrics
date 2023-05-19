package environment

import (
	"flag"
)

const (
	agentDefaultSendingAdressURL = "localhost:8080"
	agentDefaultPollInterval     = 2
	agentDefaultReportInterval   = 10

	agentEnvSendingAdressURL = "ADDRESS"
	agentEnvPollInterval     = "POLL_INTERVAL"
	agentEnvReportInterval   = "REPORT_INTERVAL"
)

type AgentConfig struct {
	SendingAdress   string
	PollIntervalS   int
	ReportIntervalS int
}

func NewAgentConfig() AgentConfig {

	sendingAdressFlag := flag.String("a", agentDefaultSendingAdressURL, "SENDING_ADRESS")
	pollIntervalFlag := flag.Int("p", agentDefaultPollInterval, "POLL_INTERVAL")
	reportIntervalFlag := flag.Int("r", agentDefaultReportInterval, "REPORT_INTERVAL")

	flag.Parse()

	ac := AgentConfig{
		SendingAdress:   GetEnvString(agentEnvSendingAdressURL, *sendingAdressFlag),
		PollIntervalS:   GetEnvDuration(agentEnvPollInterval, *pollIntervalFlag),
		ReportIntervalS: GetEnvDuration(agentEnvReportInterval, *reportIntervalFlag),
	}

	return ac
}
