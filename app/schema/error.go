package schema

type DecodeErr struct {
	Err error
}

func (DecodeErr) ResError() string {
	return "Could not decode request: JSON parsing failed"
}
