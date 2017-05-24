package scout_apm_go

import (
	"os"
	"time"
)

type ScoutTransaction struct {
	StartTime time.Time
	StopTime  time.Time
	Context   map[string]string
	OrgKey    string
}

func NewTransaction() *ScoutTransaction {
	t := &ScoutTransaction{}
	t.Context = make(map[string]string)

	t.AddDefaultContext()

	return t
}

func (transaction *ScoutTransaction) Start() {
	transaction.StartTime = time.Now()
}

// Marks stop time, and then reports
func (transaction *ScoutTransaction) Stop() {
	transaction.StopTime = time.Now()
	ShipTransaction(transaction)
}

func (transaction *ScoutTransaction) AddContext(key string, value string) {
	transaction.Context[key] = value
}

func (transaction *ScoutTransaction) AddDefaultContext() {
	hostname, err := os.Hostname()
	if err == nil {
		transaction.AddContext("hostname", hostname)
	}
}

func (transaction *ScoutTransaction) Duration() time.Duration {
	return transaction.StopTime.Sub(transaction.StartTime)
}

func (transaction *ScoutTransaction) EncodeToProtobuf() *Event {
	return &Event{
		Timestamp: transaction.StartTime.Format(time.RFC3339),
		Duration:  int64(transaction.Duration()),
		Context:   transaction.Context,
		OrgKey:    scoutOrgKey,
	}
}
