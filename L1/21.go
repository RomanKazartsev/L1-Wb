type LoggerAdapter struct {
	oldLogger *OldLogger
}

func (a *LoggerAdapter) Info(msg string) {
	a.oldLogger.LogMessage("Info: " + msg)
}

func (a *LoggerAdapter) Error(msg string) {
	a.oldLogger.LogMessage("Error: " + msg)
}
