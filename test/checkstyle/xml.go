package checkstyle

type CheckStyle struct {
	Files []CheckStyleFile `xml:"file,omitempty"`
}

type CheckStyleFile struct {
	Name   string            `xml:"name,attr"`
	Errors []CheckStyleError `xml:"error"`
}

type CheckStyleError struct {
	Column   int    `xml:"column,attr,omitempty"`
	Line     int    `xml:"line,attr"`
	Message  string `xml:"message,attr"`
	Severity string `xml:"severity,attr,omitempty"`
	Source   string `xml:"source,attr,omitempty"`
}
