package gLog

const (
	OUTPUT_STDOUT uint8 = 1
	OUTPUT_FILE   uint8 = 2
)

var outputType uint8 = OUTPUT_FILE

func SetOutputType(output uint8) {
	outputType = output
}

func OutputType() uint8 {
	return outputType
}
