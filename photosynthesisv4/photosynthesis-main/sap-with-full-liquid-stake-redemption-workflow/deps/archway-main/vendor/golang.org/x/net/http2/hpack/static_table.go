// go generate gen.go
// Code generated by the command above; DO NOT EDIT.

package hpack

var staticTable = &headerFieldTable{
	evictCount: 0,
	byName: map[string]uint64{
		":authority":                  1,
		":method":                     3,
		":path":                       5,
		":scheme":                     7,
		":status":                     14,
		"accept-charset":              15,
		"accept-encoding":             16,
		"accept-language":             17,
		"accept-ranges":               18,
		"accept":                      19,
		"access-control-allow-origin": 20,
		"age":                         21,
		"allow":                       22,
		"authorization":               23,
		"cache-control":               24,
		"content-disposition":         25,
		"content-encoding":            26,
		"content-language":            27,
		"content-length":              28,
		"content-location":            29,
		"content-range":               30,
		"content-type":                31,
		"cookie":                      32,
		"date":                        33,
		"etag":                        34,
		"expect":                      35,
		"expires":                     36,
		"from":                        37,
		"host":                        38,
		"if-match":                    39,
		"if-modified-since":           40,
		"if-none-match":               41,
		"if-range":                    42,
		"if-unmodified-since":         43,
		"last-modified":               44,
		"link":                        45,
		"location":                    46,
		"max-forwards":                47,
		"proxy-authenticate":          48,
		"proxy-authorization":         49,
		"range":                       50,
		"referer":                     51,
		"refresh":                     52,
		"retry-after":                 53,
		"server":                      54,
		"set-cookie":                  55,
		"strict-transport-security":   56,
		"transfer-encoding":           57,
		"user-agent":                  58,
		"vary":                        59,
		"via":                         60,
		"www-authenticate":            61,
	},
	byNameValue: map[pairNameValue]uint64{
		{name: ":authority", value: ""}:                   1,
		{name: ":method", value: "GET"}:                   2,
		{name: ":method", value: "POST"}:                  3,
		{name: ":path", value: "/"}:                       4,
		{name: ":path", value: "/index.html"}:             5,
		{name: ":scheme", value: "http"}:                  6,
		{name: ":scheme", value: "https"}:                 7,
		{name: ":status", value: "200"}:                   8,
		{name: ":status", value: "204"}:                   9,
		{name: ":status", value: "206"}:                   10,
		{name: ":status", value: "304"}:                   11,
		{name: ":status", value: "400"}:                   12,
		{name: ":status", value: "404"}:                   13,
		{name: ":status", value: "500"}:                   14,
		{name: "accept-charset", value: ""}:               15,
		{name: "accept-encoding", value: "gzip, deflate"}: 16,
		{name: "accept-language", value: ""}:              17,
		{name: "accept-ranges", value: ""}:                18,
		{name: "accept", value: ""}:                       19,
		{name: "access-control-allow-origin", value: ""}:  20,
		{name: "age", value: ""}:                          21,
		{name: "allow", value: ""}:                        22,
		{name: "authorization", value: ""}:                23,
		{name: "cache-control", value: ""}:                24,
		{name: "content-disposition", value: ""}:          25,
		{name: "content-encoding", value: ""}:             26,
		{name: "content-language", value: ""}:             27,
		{name: "content-length", value: ""}:               28,
		{name: "content-location", value: ""}:             29,
		{name: "content-range", value: ""}:                30,
		{name: "content-type", value: ""}:                 31,
		{name: "cookie", value: ""}:                       32,
		{name: "date", value: ""}:                         33,
		{name: "etag", value: ""}:                         34,
		{name: "expect", value: ""}:                       35,
		{name: "expires", value: ""}:                      36,
		{name: "from", value: ""}:                         37,
		{name: "host", value: ""}:                         38,
		{name: "if-match", value: ""}:                     39,
		{name: "if-modified-since", value: ""}:            40,
		{name: "if-none-match", value: ""}:                41,
		{name: "if-range", value: ""}:                     42,
		{name: "if-unmodified-since", value: ""}:          43,
		{name: "last-modified", value: ""}:                44,
		{name: "link", value: ""}:                         45,
		{name: "location", value: ""}:                     46,
		{name: "max-forwards", value: ""}:                 47,
		{name: "proxy-authenticate", value: ""}:           48,
		{name: "proxy-authorization", value: ""}:          49,
		{name: "range", value: ""}:                        50,
		{name: "referer", value: ""}:                      51,
		{name: "refresh", value: ""}:                      52,
		{name: "retry-after", value: ""}:                  53,
		{name: "server", value: ""}:                       54,
		{name: "set-cookie", value: ""}:                   55,
		{name: "strict-transport-security", value: ""}:    56,
		{name: "transfer-encoding", value: ""}:            57,
		{name: "user-agent", value: ""}:                   58,
		{name: "vary", value: ""}:                         59,
		{name: "via", value: ""}:                          60,
		{name: "www-authenticate", value: ""}:             61,
	},
	ents: []HeaderField{
		{Name: ":authority", Value: "", Sensitive: false},
		{Name: ":method", Value: "GET", Sensitive: false},
		{Name: ":method", Value: "POST", Sensitive: false},
		{Name: ":path", Value: "/", Sensitive: false},
		{Name: ":path", Value: "/index.html", Sensitive: false},
		{Name: ":scheme", Value: "http", Sensitive: false},
		{Name: ":scheme", Value: "https", Sensitive: false},
		{Name: ":status", Value: "200", Sensitive: false},
		{Name: ":status", Value: "204", Sensitive: false},
		{Name: ":status", Value: "206", Sensitive: false},
		{Name: ":status", Value: "304", Sensitive: false},
		{Name: ":status", Value: "400", Sensitive: false},
		{Name: ":status", Value: "404", Sensitive: false},
		{Name: ":status", Value: "500", Sensitive: false},
		{Name: "accept-charset", Value: "", Sensitive: false},
		{Name: "accept-encoding", Value: "gzip, deflate", Sensitive: false},
		{Name: "accept-language", Value: "", Sensitive: false},
		{Name: "accept-ranges", Value: "", Sensitive: false},
		{Name: "accept", Value: "", Sensitive: false},
		{Name: "access-control-allow-origin", Value: "", Sensitive: false},
		{Name: "age", Value: "", Sensitive: false},
		{Name: "allow", Value: "", Sensitive: false},
		{Name: "authorization", Value: "", Sensitive: false},
		{Name: "cache-control", Value: "", Sensitive: false},
		{Name: "content-disposition", Value: "", Sensitive: false},
		{Name: "content-encoding", Value: "", Sensitive: false},
		{Name: "content-language", Value: "", Sensitive: false},
		{Name: "content-length", Value: "", Sensitive: false},
		{Name: "content-location", Value: "", Sensitive: false},
		{Name: "content-range", Value: "", Sensitive: false},
		{Name: "content-type", Value: "", Sensitive: false},
		{Name: "cookie", Value: "", Sensitive: false},
		{Name: "date", Value: "", Sensitive: false},
		{Name: "etag", Value: "", Sensitive: false},
		{Name: "expect", Value: "", Sensitive: false},
		{Name: "expires", Value: "", Sensitive: false},
		{Name: "from", Value: "", Sensitive: false},
		{Name: "host", Value: "", Sensitive: false},
		{Name: "if-match", Value: "", Sensitive: false},
		{Name: "if-modified-since", Value: "", Sensitive: false},
		{Name: "if-none-match", Value: "", Sensitive: false},
		{Name: "if-range", Value: "", Sensitive: false},
		{Name: "if-unmodified-since", Value: "", Sensitive: false},
		{Name: "last-modified", Value: "", Sensitive: false},
		{Name: "link", Value: "", Sensitive: false},
		{Name: "location", Value: "", Sensitive: false},
		{Name: "max-forwards", Value: "", Sensitive: false},
		{Name: "proxy-authenticate", Value: "", Sensitive: false},
		{Name: "proxy-authorization", Value: "", Sensitive: false},
		{Name: "range", Value: "", Sensitive: false},
		{Name: "referer", Value: "", Sensitive: false},
		{Name: "refresh", Value: "", Sensitive: false},
		{Name: "retry-after", Value: "", Sensitive: false},
		{Name: "server", Value: "", Sensitive: false},
		{Name: "set-cookie", Value: "", Sensitive: false},
		{Name: "strict-transport-security", Value: "", Sensitive: false},
		{Name: "transfer-encoding", Value: "", Sensitive: false},
		{Name: "user-agent", Value: "", Sensitive: false},
		{Name: "vary", Value: "", Sensitive: false},
		{Name: "via", Value: "", Sensitive: false},
		{Name: "www-authenticate", Value: "", Sensitive: false},
	},
}
