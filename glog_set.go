package log

import "fmt"

func SetLevel(i int32){
	threshold := severity(i)
	logging.stderrThreshold.set(threshold)
	fmt.Println("level: ",logging.stderrThreshold)
}

func SetOutput(path string){
	*logDir = path
}

func LogToStderr(v bool){
	logging.alsoToStderr = v
}