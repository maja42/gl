package gl

import "strconv"

// Source: https://www.khronos.org/registry/OpenGL-Refpages/gl4/html/glGetError.xhtml
var errMapping = map[Enum]string{
	NO_ERROR:                      "No error has been recorded.",
	INVALID_ENUM:                  "An unacceptable value is specified for an enumerated argument. The offending command is ignored and has no other side effect than to set the error flag.",
	INVALID_VALUE:                 "A numeric argument is out of range. The offending command is ignored and has no other side effect than to set the error flag.",
	INVALID_OPERATION:             "The specified operation is not allowed in the current state. The offending command is ignored and has no other side effect than to set the error flag.",
	INVALID_FRAMEBUFFER_OPERATION: "The command is trying to render to or read from the framebuffer while the currently bound framebuffer is not framebuffer complete. The offending command is ignored and has no other side effect than to set the error flag.",
	OUT_OF_MEMORY:                 "There is not enough memory left to execute the command. The state of the GL is undefined, except for the state of the error flags, after this error is recorded.",
	// STACK_UNDERFLOW:            "An attempt has been made to perform an operation that would cause an internal stack to underflow.",
	// STACK_OVERFLOW:             "An attempt has been made to perform an operation that would cause an internal stack to overflow.",
}

// GLError represents an OpenGL error.
type GLError struct {
	Code Enum
}

func (e GLError) Error() string {
	errStr, ok := errMapping[e.Code]
	if ok {
		return errStr
	}
	return "Error " + strconv.Itoa(int(e.Code))
}

// CheckError checks the current gl error flag and returns an error if available
func CheckError() error {
	if glError := GetError(); glError != NO_ERROR {
		return GLError{Code: glError}
	}
	return nil
}
