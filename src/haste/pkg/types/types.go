package types

import (
	"fmt"
	"time"
)

// Key is Used for context keys.
type Key string

// Data is used for input and output.
type Data map[string]interface{}

// IDKeyData is used for getting db rows with id as key.
type IDKeyData map[int32]Data

// DataArray is used for getting db rows.
type DataArray []Data

// ValidationError is used to hold a validation error for a field.
type ValidationError struct {
	Key                string `json:"key"`
	ValidationErrorMsg string `json:"validation_error_msg"`
}

// Errors is used to hold all types of errors including validation errors.
type Errors struct {
	ErrorMsg         string            `json:"error_msg"`
	ValidationErrors []ValidationError `json:"validation_errors"`
}

// AddValidationError adds a validation error to the validation errors array.
func (e *Errors) AddValidationError(key, msg string) {
	if e == nil {
		return
	}
	e.ValidationErrors = append(e.ValidationErrors, ValidationError{
		Key:                key,
		ValidationErrorMsg: msg,
	})
}

// LogMessage.
type LogMessage struct {
	UUID      string    `json:"uuid"`
	Seq       int       `json:"seq"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
	Filename  string    `json:"filename"`
	Linenum   string    `json:"linenum"`
	Carrier   string    `json:"carrier"`
	Timestamp time.Time `json:"ts"`
}

// APICall.
type APICall struct {
	UUID            string    `json:"uuid"`
	UserID          int32     `json:"user_id"`
	SessionID       string    `json:"sessionid"`
	URL             string    `json:"url"`
	Method          string    `json:"method"`
	Headers         string    `json:"headers"`
	Params          string    `json:"params"`
	Body            string    `json:"body"`
	Files           string    `json:"string"`
	Status          int       `json:"status"`
	ResponseHeaders string    `json:"response_headers"`
	Response        string    `json:"response"`
	Error           string    `json:"error"`
	IPAddr          string    `json:"ip_addr"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
}

// Log.
type Log struct {
	TraceID   string    `json:"trace_id"`
	Message   string    `json:"message"`
	Level     int32     `json:"level"`
	LevelName string    `json:"level_name"`
	TimeStamp time.Time `json:"timestamp"`
	Context   Data      `json:"context"`
	Extra     LogExtra  `json:"extra"`
}

// LogExtra.
type LogExtra struct {
	Location   *LogLocation   `json:"location"`
	APILogging *LogAPILogging `json:"api_logging,omitempty"`
	StackTrace string         `json:"stacktrace,omitempty"`
	DBQuery    string         `json:"db_query,omitempty"`
	RedisQuery string         `json:"redis_query,omitempty"`
	SrvName    string         `json:"srv_name,omitempty"`
	Extra      Data           `json:"extra,omitempty"`
}

// LogLocation.
type LogLocation struct {
	FilePath     string `json:"file_path,omitempty"`
	LineNumber   int32  `json:"line_number,omitempty"`
	FunctionName string `json:"function_name,omitempty"`
}

// LogAPILogging.
type LogAPILogging struct {
	ID           int64      `json:"id,omitempty"`
	Duration     float64    `json:"duration,omitempty"`
	ResponseSize int64      `json:"response_size,omitempty"`
	StartTime    *time.Time `json:"start_time,omitempty"`
	EndTime      *time.Time `json:"end_time,omitempty"`
	StatusCode   int32      `json:"status_code,omitempty"`
}

// LogClientContext.
type LogClientDetails struct {
	IPAddress string `json:"ip_address,omitempty"`
	OS        string `json:"os,omitempty"`
	Browser   string `json:"browser,omitempty"`
	Platform  string `json:"platform,omitempty"`
	Referer   string `json:"referer,omitempty"`
}

func init() {
	fmt.Println("init __ types")
}
