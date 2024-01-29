package digcloudlog

// dig cloud log
func DCL_Info(format string, args ...any) {
	GetLogAdapter().Info(format, args...)
}
func DCL_Error(format string, args ...any) {
	GetLogAdapter().Error(format, args...)
}
func DCL_addStdout() {
	GetLogAdapter().AddStdout()
}
func DCL_addLocalFileDefault() error {
	return GetLogAdapter().AddLocalFile("", "", "")
}
func DCL_addLocalFile(basePath, infoFileName, errorFileName string) error {
	return GetLogAdapter().AddLocalFile(basePath, infoFileName, errorFileName)
}
func DCL_addLogflare(sourceid, apiKey string) error {
	return GetLogAdapter().AddLogflare(sourceid, apiKey)
}

/////// third part log service
func DCL_AddNewLogService(service interface{}) error {
	return GetLogAdapter().AddNewLogService(service)
}
