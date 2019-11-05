package common

import (
	"fmt"
	"log/syslog"
	"net/url"
)

// Logger interface
type Logger interface {
	Alert(m string) error
	Crit(m string) error
	Debug(m string) error
	Emerg(m string) error
	Err(m string) error
	Info(m string) error
	Notice(m string) error
	Warning(m string) error
}

type establishConnectionFn = func(network, raddr string, priority syslog.Priority, tag string) (Logger, error)

// SyslogEstablishConnectionFn establish a connection with syslog dialler
func SyslogEstablishConnectionFn() func(network, raddr string, priority syslog.Priority, tag string) (Logger, error) {
	return func(network, raddr string, priority syslog.Priority, tag string) (Logger, error) {
		return syslog.Dial(network, raddr, priority, tag)
	}
}

// NewUnifiedLogger returns a Logger connected to a centralised log agent
func NewUnifiedLogger(c establishConnectionFn, u *url.URL, name string) (Logger, error) {
	sysLog, err := c(u.Scheme, u.Host, syslog.LOG_NOTICE|syslog.LOG_DAEMON, name)
	if err != nil {
		return nil, err
	}

	return sysLog, nil
}

type testConnection struct {
	Name       string
	RemoteAddr string
}

func testEstablishConnectionFn(tc *testConnection) func(network, raddr string, priority syslog.Priority, tag string) (Logger, error) {
	return func(network, raddr string, priority syslog.Priority, tag string) (Logger, error) {
		tc.Name = tag
		tc.RemoteAddr = raddr
		return &testLogger{}, nil
	}
}

type testLogger struct {
	Log string
}

func (l *testLogger) Alert(m string) error {
	l.Log = fmt.Sprintf("Alert: %s", m)
	return nil
}

func (l *testLogger) Crit(m string) error {
	l.Log = fmt.Sprintf("Crit: %s", m)
	return nil
}

func (l *testLogger) Debug(m string) error {
	l.Log = fmt.Sprintf("Debug: %s", m)
	return nil
}

func (l *testLogger) Emerg(m string) error {
	l.Log = fmt.Sprintf("Emerg: %s", m)
	return nil
}

func (l *testLogger) Err(m string) error {
	l.Log = fmt.Sprintf("Err: %s", m)
	return nil
}

func (l *testLogger) Info(m string) error {
	l.Log = fmt.Sprintf("Info: %s", m)
	return nil
}

func (l *testLogger) Notice(m string) error {
	l.Log = fmt.Sprintf("Notice: %s", m)
	return nil
}

func (l *testLogger) Warning(m string) error {
	l.Log = fmt.Sprintf("Warning: %s", m)
	return nil
}
