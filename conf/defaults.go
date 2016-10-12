package conf

var defaultOptions = Options{
	RestEnabled:  false,
	PollingTime:  1,
	PgwTolerance: 60,
	SgwTolerance: 60,
	PgwScaleStep: 1,
	SgwScaleStep: 1,
	MarathonURL:  "master.mesos:8080",
	PgwJSON:      "pgw.json",
	SgwJSON:      "sgw.json",
}
