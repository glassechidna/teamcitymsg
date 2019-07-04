package teamcitymsg

import (
	"fmt"
	"time"
)

func NewMsgMessage(input, errorDetails, status string) ServiceMessage {
	attrs := map[string]string{
		"text": input,
	}
	if len(status) > 0 {
		attrs["status"] = status
	}
	if len(errorDetails) > 0 {
		attrs["errorDetails"] = errorDetails
	}
	return &multi{
		msgType: "message",
		attrs:   attrs,
	}
}

func NewMsgBlockOpened(blockName string) ServiceMessage {
	return &multi{
		msgType: "blockOpened",
		attrs:   map[string]string{"name": blockName},
	}
}

func NewMsgBlockClosed(blockName string) ServiceMessage {
	return &multi{
		msgType: "blockClosed",
		attrs:   map[string]string{"name": blockName},
	}
}

func NewMsgCompilationStarted(compiler string) ServiceMessage {
	return &multi{
		msgType: "compilationStarted",
		attrs:   map[string]string{"compiler": compiler},
	}
}

func NewMsgCompilationFinished(compiler string) ServiceMessage {
	return &multi{
		msgType: "compilationFinished",
		attrs:   map[string]string{"compiler": compiler},
	}
}

func NewMsgTestSuiteStarted(suiteName string) ServiceMessage {
	return &multi{
		msgType: "testSuiteStarted",
		attrs:   map[string]string{"name": suiteName},
	}
}

func NewMsgTestSuiteFinished(suiteName string) ServiceMessage {
	return &multi{
		msgType: "testSuiteFinished",
		attrs:   map[string]string{"name": suiteName},
	}
}

func NewMsgTestStarted(testName string, captureStdout bool) ServiceMessage {
	cap := "false"
	if captureStdout {
		cap = "true"
	}

	return &multi{
		msgType: "testStarted",
		attrs: map[string]string{
			"name":                  testName,
			"captureStandardOutput": cap,
		},
	}
}

func NewMsgTestFinished(testName string, duration time.Duration) ServiceMessage {
	return &multi{
		msgType: "testFinished",
		attrs: map[string]string{
			"name":     testName,
			"duration": fmt.Sprintf("%d", int(duration.Nanoseconds()/1000)),
		},
	}
}

func NewMsgTestIgnored(testName, message string) ServiceMessage {
	return &multi{
		msgType: "testIgnored",
		attrs: map[string]string{
			"name":    testName,
			"message": message,
		},
	}
}

func NewMsgTestStdout(testName, data string) ServiceMessage {
	return &multi{
		msgType: "testStdOut",
		attrs: map[string]string{
			"name": testName,
			"out":  data,
		},
	}
}

func NewMsgTestStdErr(testName, data string) ServiceMessage {
	return &multi{
		msgType: "testStdErr",
		attrs: map[string]string{
			"name": testName,
			"out":  data,
		},
	}
}

// TODO: another with type,expected,actual
func NewMsgTestFailed(testName, message, details string) ServiceMessage {
	return &multi{
		msgType: "testFailed",
		attrs: map[string]string{
			"name":    testName,
			"message": message,
			"details": details,
		},
	}
}

func NewMsgInspectionType(id, name, description, category string) ServiceMessage {
	return &multi{
		msgType: "inspectionType",
		attrs: map[string]string{
			"id":          id,
			"name":        name,
			"description": description,
			"category":    category,
		},
	}
}

func NewMsgInspection(typeId, file, message string, line int) ServiceMessage {
	attrs := map[string]string{
		"typeId": typeId,
		"file":   file,
	}
	if len(message) > 0 {
		attrs["message"] = message
	}
	if line >= 0 {
		attrs["line"] = fmt.Sprintf("%d", line)
	}
	return &multi{
		msgType: "inspection",
		attrs:   attrs,
	}
}

func NewMsgPublishArtifacts(path string) ServiceMessage {
	return &single{
		msgType: "publishArtifacts",
		value:   path,
	}
}

func NewMsgProgressMessage(input string) ServiceMessage {
	return &single{
		msgType: "progressMessage",
		value:   input,
	}
}

func NewMsgProgressStart(input string) ServiceMessage {
	return &single{
		msgType: "progressStart",
		value:   input,
	}
}

func NewMsgProgressFinish(input string) ServiceMessage {
	return &single{
		msgType: "progressFinish",
		value:   input,
	}
}

func NewMsgBuildProblem(description, identity string) ServiceMessage {
	attrs := map[string]string{"description": description}
	if len(identity) > 0 {
		attrs["identity"] = identity
	}
	return &multi{
		msgType: "buildProblem",
		attrs:   attrs,
	}
}

func NewMsgBuildStatus(text, status string) ServiceMessage {
	attrs := map[string]string{"text": text}
	if len(status) > 0 {
		attrs["status"] = status
	}
	return &multi{
		msgType: "buildStatus",
		attrs:   attrs,
	}
}

func NewMsgBuildNumber(input string) ServiceMessage {
	return &single{
		msgType: "buildNumber",
		value:   input,
	}
}

func NewMsgBuildStatisticValue(key string, value int) ServiceMessage {
	return &multi{
		msgType: "buildStatisticValue",
		attrs: map[string]string{
			"key":   key,
			"value": fmt.Sprintf("%d", value),
		},
	}
}

func NewMsgSetParameter(name, value string) ServiceMessage {
	return &multi{
		msgType: "setParameter",
		attrs: map[string]string{
			"name":  name,
			"value": value,
		},
	}
}

//func NewMsgEnableServiceMessages(input string) ServiceMessage {
//	return &multi{
//		msgType: "enableServiceMessages",
//		attrs:   map[string]string{},
//	}
//}
//
//func NewMsgDisableServiceMessages(input string) ServiceMessage {
//	return &multi{
//		msgType: "disableServiceMessages",
//		attrs:   map[string]string{},
//	}
//}

// TODO: enum for typeId
func NewMsgImportData(typeId, path string) ServiceMessage {
	return &multi{
		msgType: "importData",
		attrs: map[string]string{
			"type": typeId,
			"path": path,
		},
	}
}
