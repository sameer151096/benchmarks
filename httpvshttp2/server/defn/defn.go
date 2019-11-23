package defn

import "fmt"

type Request struct {
	Key int
}

type Response struct {
	Value string
}

type LoggerDefn struct {
	Error          string
	Message        string
	ParentFuncName string
}

func (logger *LoggerDefn) SetLogggerDefn(ErrorOccured error, MessageForError string, FuncIn string) string {
	logger.Error = ErrorOccured.Error()
	logger.Message = MessageForError
	logger.ParentFuncName = FuncIn

	loggerstring := fmt.Sprintf(" Error: %v ; Message: %v ; ParentFuncName: %v", logger.Error, logger.Message, logger.ParentFuncName)
	return loggerstring
}
