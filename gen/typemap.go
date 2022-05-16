package gen

var TypeCHDicMp = map[string]string{
	"UInt":       "uint",
	"UInt8":      "uint8",
	"UInt16":     "uint16",
	"UInt32":     "uint32",
	"UInt64":     "uint64",
	"Int":        "int",
	"Int8":       "int8",
	"Int16":      "int16",
	"Int32":      "int32",
	"Int64":      "int64",
	"Float32":    "float32",
	"Float64":    "float64",
	"Date":       "time.Time",
	"Date32 ":    "time.Time",
	"Datetime":   "time.Time",
	"Datetime64": "time.Time",
	"Boolean":    "bool",
	"String":     "string",
	"Array":      "array",
}

var TypeCHMatchList = []struct {
	Key   string
	Value string
}{
	{`^VARCHAR()[(]\d+[)]`, "string"},
	{`^VARCHAR2()[(]\d+[)]`, "string"},
	{`^Date()[(]\d+[)]`, "time.Time"},
	{`^Date32()[(]\d+[)]`, "time.Time"},
	{`^Datetime()[(]\d+[)]`, "time.Time"},
	{`^DateTime64()[(]\d+[)]`, "time.Time"},
	{`^UInt()[(]\d+[)]`, "uint"},
	{`^UInt8()[(]\d+[)]`, "uint8"},
	{`^UInt16()[(]\d+[)]`, "uint16"},
	{`^UInt32()[(]\d+[)]`, "uint32"},
	{`^UInt64()[(]\d+[)]`, "uint64"},
	{`^Int()[(]\d+[)]`, "int"},
	{`^Int8()[(]\d+[)]`, "int8"},
	{`^Int16()[(]\d+[)]`, "int16"},
	{`^Int32()[(]\d+[)]`, "int32"},
	{`^Int64()[(]\d+[)]`, "int64"},
	{`^Float32()[(]\d+[)]`, "float32"},
	{`^Float64()[(]\d+[)]`, "float64"},
}
