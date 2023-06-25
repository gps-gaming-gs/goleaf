package internal

import (
    {{.imports}}
)

func init() {
	{{.handlers}}
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
