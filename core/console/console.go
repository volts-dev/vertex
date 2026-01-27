package console

import (
	"github.com/volts-dev/vertex/js"
)

var value js.Value

func init() {
	value = js.Global().Get("console")
}

func AssertErr(err error) bool {
	if err != nil {
		Assert(err == nil, err.Error())
	}

	return err == nil
}

func Assert(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("assert", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Clear(args ...interface{}) {
	if v := value.Call("clear", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Count(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("count", args...); v.Error() != nil {
		Error(v.Error())
	}
}
func CountReset(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("countReset", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Debug(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("debug", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Dir(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("dir", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Dirxml(args ...interface{}) {
	if v := value.Call("dirxml", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Error(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("error", args...); v.Error() != nil {
		value.Call("error", v.Error().Error())
	}
}

func Group(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("group", args...); v.Error() != nil {
		Error(v.Error())
	}
}
func GroupCollapsed(args ...interface{}) {
	if v := value.Call("groupCollapsed", args...); v.Error() != nil {
		Error(v.Error())
	}
}
func GroupEnd(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("groupEnd", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Info(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("info", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Log(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("log", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Table(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("table", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Time(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("time", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func TimeEnd(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("timeEnd", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func TimeLog(args ...interface{}) {
	if v := value.Call("timeLog", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Trace(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("trace", args...); v.Error() != nil {
		Error(v.Error())
	}
}

func Warn(args ...interface{}) {
	args = js.ToValues(args...)
	if v := value.Call("warn", args...); v.Error() != nil {
		Error(v.Error())
	}
}
