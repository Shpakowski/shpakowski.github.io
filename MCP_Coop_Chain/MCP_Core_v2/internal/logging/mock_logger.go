package logging

import "go.uber.org/zap"

// MockLogger — мок-реализация Logger для тестов.
type MockLogger struct {
	Entries []string
}

func (m *MockLogger) Debug(msg string, fields ...zap.Field) {
	m.Entries = append(m.Entries, "DEBUG: "+msg)
}
func (m *MockLogger) Info(msg string, fields ...zap.Field) {
	m.Entries = append(m.Entries, "INFO: "+msg)
}
func (m *MockLogger) Warn(msg string, fields ...zap.Field) {
	m.Entries = append(m.Entries, "WARN: "+msg)
}
func (m *MockLogger) Error(msg string, err error, fields ...zap.Field) {
	m.Entries = append(m.Entries, "ERROR: "+msg+" err="+err.Error())
}
