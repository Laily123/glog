package log

func SetLevel(i int32){
	threshold := severity(i)
	logging.stderrThreshold.set(threshold)
}

func SetOutput(path string){
	*logDir = path
}

func LogToStderr(v bool){
	logging.alsoToStderr = v
}