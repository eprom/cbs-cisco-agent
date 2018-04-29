package status

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type tuple struct {
	field string
	value interface{}
}

type Status struct {
	Allocated  int
	Stored     int
	Used       int
	Deleted    int
	Sent       int
	Pending    int
	Failed     int
	LastStatus string
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Status",
			pos:  position{line: 22, col: 1, offset: 293},
			expr: &actionExpr{
				pos: position{line: 22, col: 11, offset: 303},
				run: (*parser).callonStatus1,
				expr: &seqExpr{
					pos: position{line: 22, col: 11, offset: 303},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 22, col: 11, offset: 303},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 22, col: 14, offset: 306},
							label: "infos",
							expr: &oneOrMoreExpr{
								pos: position{line: 22, col: 20, offset: 312},
								expr: &choiceExpr{
									pos: position{line: 22, col: 22, offset: 314},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 22, col: 22, offset: 314},
											name: "IncomingInfo",
										},
										&ruleRefExpr{
											pos:  position{line: 22, col: 37, offset: 329},
											name: "OutgoingInfo",
										},
										&ruleRefExpr{
											pos:  position{line: 22, col: 52, offset: 344},
											name: "UnhandledInfo",
										},
										&ruleRefExpr{
											pos:  position{line: 22, col: 68, offset: 360},
											name: "EmptyLine",
										},
									},
								},
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 22, col: 81, offset: 373},
							expr: &anyMatcher{
								line: 22, col: 81, offset: 373,
							},
						},
						&ruleRefExpr{
							pos:  position{line: 22, col: 84, offset: 376},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "IncomingInfo",
			pos:  position{line: 56, col: 1, offset: 1306},
			expr: &actionExpr{
				pos: position{line: 57, col: 3, offset: 1325},
				run: (*parser).callonIncomingInfo1,
				expr: &seqExpr{
					pos: position{line: 57, col: 3, offset: 1325},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 57, col: 3, offset: 1325},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 57, col: 5, offset: 1327},
							val:        "incoming",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 17, offset: 1339},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 57, col: 19, offset: 1341},
							val:        "message",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 30, offset: 1352},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 57, col: 32, offset: 1354},
							val:        "information",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 47, offset: 1369},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 49, offset: 1371},
							name: "NL",
						},
						&ruleRefExpr{
							pos:  position{line: 58, col: 3, offset: 1377},
							name: "SeparatorLine",
						},
						&labeledExpr{
							pos:   position{line: 59, col: 3, offset: 1394},
							label: "data",
							expr: &zeroOrMoreExpr{
								pos: position{line: 59, col: 8, offset: 1399},
								expr: &choiceExpr{
									pos: position{line: 59, col: 10, offset: 1401},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 59, col: 10, offset: 1401},
											name: "Stored",
										},
										&ruleRefExpr{
											pos:  position{line: 59, col: 19, offset: 1410},
											name: "Deleted",
										},
										&ruleRefExpr{
											pos:  position{line: 59, col: 29, offset: 1420},
											name: "Allocated",
										},
										&ruleRefExpr{
											pos:  position{line: 59, col: 41, offset: 1432},
											name: "Used",
										},
										&ruleRefExpr{
											pos:  position{line: 59, col: 48, offset: 1439},
											name: "AnyLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 60, col: 3, offset: 1453},
							name: "EmptyLine",
						},
					},
				},
			},
		},
		{
			name: "Stored",
			pos:  position{line: 63, col: 1, offset: 1490},
			expr: &actionExpr{
				pos: position{line: 63, col: 11, offset: 1500},
				run: (*parser).callonStored1,
				expr: &seqExpr{
					pos: position{line: 63, col: 11, offset: 1500},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 63, col: 11, offset: 1500},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 63, col: 13, offset: 1502},
							val:        "sms",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 20, offset: 1509},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 63, col: 22, offset: 1511},
							val:        "stored",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 32, offset: 1521},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 63, col: 34, offset: 1523},
							val:        "in",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 40, offset: 1529},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 63, col: 42, offset: 1531},
							val:        "modem",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 51, offset: 1540},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 63, col: 53, offset: 1542},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 57, offset: 1546},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 63, col: 59, offset: 1548},
							label: "stored",
							expr: &ruleRefExpr{
								pos:  position{line: 63, col: 66, offset: 1555},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 74, offset: 1563},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Deleted",
			pos:  position{line: 66, col: 1, offset: 1633},
			expr: &actionExpr{
				pos: position{line: 66, col: 12, offset: 1644},
				run: (*parser).callonDeleted1,
				expr: &seqExpr{
					pos: position{line: 66, col: 12, offset: 1644},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 66, col: 12, offset: 1644},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 14, offset: 1646},
							val:        "total",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 23, offset: 1655},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 25, offset: 1657},
							val:        "sms",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 32, offset: 1664},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 34, offset: 1666},
							val:        "deleted",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 45, offset: 1677},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 47, offset: 1679},
							val:        "since",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 56, offset: 1688},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 58, offset: 1690},
							val:        "booting",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 69, offset: 1701},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 71, offset: 1703},
							val:        "up",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 77, offset: 1709},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 66, col: 79, offset: 1711},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 83, offset: 1715},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 66, col: 85, offset: 1717},
							label: "deleted",
							expr: &ruleRefExpr{
								pos:  position{line: 66, col: 93, offset: 1725},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 66, col: 101, offset: 1733},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Allocated",
			pos:  position{line: 69, col: 1, offset: 1805},
			expr: &actionExpr{
				pos: position{line: 69, col: 14, offset: 1818},
				run: (*parser).callonAllocated1,
				expr: &seqExpr{
					pos: position{line: 69, col: 14, offset: 1818},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 69, col: 14, offset: 1818},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 69, col: 16, offset: 1820},
							val:        "storage",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 27, offset: 1831},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 69, col: 29, offset: 1833},
							val:        "records",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 40, offset: 1844},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 69, col: 42, offset: 1846},
							val:        "allocated",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 55, offset: 1859},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 69, col: 57, offset: 1861},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 61, offset: 1865},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 69, col: 63, offset: 1867},
							label: "allocated",
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 73, offset: 1877},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 81, offset: 1885},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Used",
			pos:  position{line: 72, col: 1, offset: 1961},
			expr: &actionExpr{
				pos: position{line: 72, col: 9, offset: 1969},
				run: (*parser).callonUsed1,
				expr: &seqExpr{
					pos: position{line: 72, col: 9, offset: 1969},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 72, col: 9, offset: 1969},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 72, col: 11, offset: 1971},
							val:        "storage",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 72, col: 22, offset: 1982},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 72, col: 24, offset: 1984},
							val:        "records",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 72, col: 35, offset: 1995},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 72, col: 37, offset: 1997},
							val:        "used",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 72, col: 45, offset: 2005},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 72, col: 47, offset: 2007},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 72, col: 51, offset: 2011},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 72, col: 53, offset: 2013},
							label: "used",
							expr: &ruleRefExpr{
								pos:  position{line: 72, col: 58, offset: 2018},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 72, col: 66, offset: 2026},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "OutgoingInfo",
			pos:  position{line: 78, col: 1, offset: 2098},
			expr: &actionExpr{
				pos: position{line: 79, col: 3, offset: 2117},
				run: (*parser).callonOutgoingInfo1,
				expr: &seqExpr{
					pos: position{line: 79, col: 3, offset: 2117},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 79, col: 3, offset: 2117},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 79, col: 5, offset: 2119},
							val:        "outgoing",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 79, col: 17, offset: 2131},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 79, col: 19, offset: 2133},
							val:        "message",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 79, col: 30, offset: 2144},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 79, col: 32, offset: 2146},
							val:        "information",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 79, col: 47, offset: 2161},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 79, col: 49, offset: 2163},
							name: "NL",
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 3, offset: 2169},
							name: "SeparatorLine",
						},
						&labeledExpr{
							pos:   position{line: 81, col: 3, offset: 2186},
							label: "data",
							expr: &zeroOrMoreExpr{
								pos: position{line: 81, col: 8, offset: 2191},
								expr: &choiceExpr{
									pos: position{line: 81, col: 10, offset: 2193},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 81, col: 10, offset: 2193},
											name: "Sent",
										},
										&ruleRefExpr{
											pos:  position{line: 81, col: 17, offset: 2200},
											name: "Pending",
										},
										&ruleRefExpr{
											pos:  position{line: 81, col: 27, offset: 2210},
											name: "Failed",
										},
										&ruleRefExpr{
											pos:  position{line: 81, col: 36, offset: 2219},
											name: "LastStatus",
										},
										&ruleRefExpr{
											pos:  position{line: 81, col: 49, offset: 2232},
											name: "AnyLine",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 3, offset: 2246},
							name: "EmptyLine",
						},
					},
				},
			},
		},
		{
			name: "Sent",
			pos:  position{line: 85, col: 1, offset: 2283},
			expr: &actionExpr{
				pos: position{line: 85, col: 9, offset: 2291},
				run: (*parser).callonSent1,
				expr: &seqExpr{
					pos: position{line: 85, col: 9, offset: 2291},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 85, col: 9, offset: 2291},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 85, col: 11, offset: 2293},
							val:        "total",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 20, offset: 2302},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 85, col: 22, offset: 2304},
							val:        "sms",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 29, offset: 2311},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 85, col: 31, offset: 2313},
							val:        "sent",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 39, offset: 2321},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 85, col: 41, offset: 2323},
							val:        "successfully",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 57, offset: 2339},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 85, col: 59, offset: 2341},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 63, offset: 2345},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 85, col: 65, offset: 2347},
							label: "sent",
							expr: &ruleRefExpr{
								pos:  position{line: 85, col: 70, offset: 2352},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 78, offset: 2360},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Pending",
			pos:  position{line: 88, col: 1, offset: 2426},
			expr: &actionExpr{
				pos: position{line: 88, col: 12, offset: 2437},
				run: (*parser).callonPending1,
				expr: &seqExpr{
					pos: position{line: 88, col: 12, offset: 2437},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 88, col: 12, offset: 2437},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 14, offset: 2439},
							val:        "number",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 24, offset: 2449},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 26, offset: 2451},
							val:        "of",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 32, offset: 2457},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 34, offset: 2459},
							val:        "outgoing",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 46, offset: 2471},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 48, offset: 2473},
							val:        "sms",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 55, offset: 2480},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 57, offset: 2482},
							val:        "pending",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 68, offset: 2493},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 88, col: 70, offset: 2495},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 74, offset: 2499},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 88, col: 76, offset: 2501},
							label: "pending",
							expr: &ruleRefExpr{
								pos:  position{line: 88, col: 84, offset: 2509},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 88, col: 92, offset: 2517},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Failed",
			pos:  position{line: 91, col: 1, offset: 2589},
			expr: &actionExpr{
				pos: position{line: 91, col: 11, offset: 2599},
				run: (*parser).callonFailed1,
				expr: &seqExpr{
					pos: position{line: 91, col: 11, offset: 2599},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 91, col: 11, offset: 2599},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 91, col: 13, offset: 2601},
							val:        "total",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 91, col: 22, offset: 2610},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 91, col: 24, offset: 2612},
							val:        "sms",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 91, col: 31, offset: 2619},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 91, col: 33, offset: 2621},
							val:        "send",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 91, col: 41, offset: 2629},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 91, col: 43, offset: 2631},
							val:        "failure",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 91, col: 54, offset: 2642},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 91, col: 56, offset: 2644},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 91, col: 60, offset: 2648},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 91, col: 62, offset: 2650},
							label: "failed",
							expr: &ruleRefExpr{
								pos:  position{line: 91, col: 69, offset: 2657},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 91, col: 77, offset: 2665},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "LastStatus",
			pos:  position{line: 94, col: 1, offset: 2735},
			expr: &actionExpr{
				pos: position{line: 94, col: 15, offset: 2749},
				run: (*parser).callonLastStatus1,
				expr: &seqExpr{
					pos: position{line: 94, col: 15, offset: 2749},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 94, col: 15, offset: 2749},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 17, offset: 2751},
							val:        "last",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 25, offset: 2759},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 27, offset: 2761},
							val:        "outgoing",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 39, offset: 2773},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 41, offset: 2775},
							val:        "sms",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 48, offset: 2782},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 50, offset: 2784},
							val:        "status",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 60, offset: 2794},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 94, col: 62, offset: 2796},
							val:        "=",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 66, offset: 2800},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 94, col: 68, offset: 2802},
							label: "status",
							expr: &ruleRefExpr{
								pos:  position{line: 94, col: 75, offset: 2809},
								name: "LineString",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 94, col: 86, offset: 2820},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "UnhandledInfo",
			pos:  position{line: 100, col: 1, offset: 2903},
			expr: &actionExpr{
				pos: position{line: 100, col: 18, offset: 2920},
				run: (*parser).callonUnhandledInfo1,
				expr: &seqExpr{
					pos: position{line: 100, col: 18, offset: 2920},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 100, col: 18, offset: 2920},
							name: "AnyLine",
						},
						&zeroOrOneExpr{
							pos: position{line: 100, col: 26, offset: 2928},
							expr: &ruleRefExpr{
								pos:  position{line: 100, col: 26, offset: 2928},
								name: "SeparatorLine",
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 100, col: 41, offset: 2943},
							expr: &ruleRefExpr{
								pos:  position{line: 100, col: 41, offset: 2943},
								name: "AnyLine",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 100, col: 50, offset: 2952},
							name: "EmptyLine",
						},
					},
				},
			},
		},
		{
			name: "SeparatorLine",
			pos:  position{line: 104, col: 1, offset: 2989},
			expr: &actionExpr{
				pos: position{line: 104, col: 18, offset: 3006},
				run: (*parser).callonSeparatorLine1,
				expr: &seqExpr{
					pos: position{line: 104, col: 18, offset: 3006},
					exprs: []interface{}{
						&choiceExpr{
							pos: position{line: 104, col: 19, offset: 3007},
							alternatives: []interface{}{
								&oneOrMoreExpr{
									pos: position{line: 104, col: 19, offset: 3007},
									expr: &litMatcher{
										pos:        position{line: 104, col: 19, offset: 3007},
										val:        "=",
										ignoreCase: false,
									},
								},
								&oneOrMoreExpr{
									pos: position{line: 104, col: 26, offset: 3014},
									expr: &litMatcher{
										pos:        position{line: 104, col: 26, offset: 3014},
										val:        "-",
										ignoreCase: false,
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 104, col: 32, offset: 3020},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "EmptyLine",
			pos:  position{line: 108, col: 1, offset: 3050},
			expr: &actionExpr{
				pos: position{line: 108, col: 14, offset: 3063},
				run: (*parser).callonEmptyLine1,
				expr: &seqExpr{
					pos: position{line: 108, col: 14, offset: 3063},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 108, col: 14, offset: 3063},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 108, col: 16, offset: 3065},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "AnyLine",
			pos:  position{line: 112, col: 1, offset: 3095},
			expr: &actionExpr{
				pos: position{line: 112, col: 12, offset: 3106},
				run: (*parser).callonAnyLine1,
				expr: &ruleRefExpr{
					pos:  position{line: 112, col: 12, offset: 3106},
					name: "NonEmptyLine",
				},
			},
		},
		{
			name: "NonEmptyLine",
			pos:  position{line: 116, col: 1, offset: 3146},
			expr: &seqExpr{
				pos: position{line: 116, col: 17, offset: 3162},
				exprs: []interface{}{
					&labeledExpr{
						pos:   position{line: 116, col: 17, offset: 3162},
						label: "str",
						expr: &ruleRefExpr{
							pos:  position{line: 116, col: 21, offset: 3166},
							name: "LineString",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 32, offset: 3177},
						name: "NL",
					},
					&andCodeExpr{
						pos: position{line: 116, col: 35, offset: 3180},
						run: (*parser).callonNonEmptyLine5,
					},
				},
			},
		},
		{
			name: "LineString",
			pos:  position{line: 125, col: 1, offset: 3312},
			expr: &actionExpr{
				pos: position{line: 125, col: 15, offset: 3326},
				run: (*parser).callonLineString1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 125, col: 15, offset: 3326},
					expr: &charClassMatcher{
						pos:        position{line: 125, col: 15, offset: 3326},
						val:        "[^\\n]",
						chars:      []rune{'\n'},
						ignoreCase: false,
						inverted:   true,
					},
				},
			},
		},
		{
			name: "Integer",
			pos:  position{line: 131, col: 1, offset: 3375},
			expr: &actionExpr{
				pos: position{line: 131, col: 12, offset: 3386},
				run: (*parser).callonInteger1,
				expr: &seqExpr{
					pos: position{line: 131, col: 12, offset: 3386},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 131, col: 12, offset: 3386},
							expr: &litMatcher{
								pos:        position{line: 131, col: 12, offset: 3386},
								val:        "-",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 131, col: 17, offset: 3391},
							expr: &ruleRefExpr{
								pos:  position{line: 131, col: 17, offset: 3391},
								name: "DecimalDigit",
							},
						},
					},
				},
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 138, col: 1, offset: 3492},
			expr: &charClassMatcher{
				pos:        position{line: 138, col: 17, offset: 3508},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name:        "__",
			displayName: "\"WSNL\"",
			pos:         position{line: 139, col: 1, offset: 3515},
			expr: &zeroOrMoreExpr{
				pos: position{line: 139, col: 14, offset: 3528},
				expr: &charClassMatcher{
					pos:        position{line: 139, col: 14, offset: 3528},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "_",
			displayName: "\"WS\"",
			pos:         position{line: 140, col: 1, offset: 3540},
			expr: &zeroOrMoreExpr{
				pos: position{line: 140, col: 11, offset: 3550},
				expr: &charClassMatcher{
					pos:        position{line: 140, col: 11, offset: 3550},
					val:        "[ \\t\\r]",
					chars:      []rune{' ', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "NL",
			pos:  position{line: 141, col: 1, offset: 3560},
			expr: &litMatcher{
				pos:        position{line: 141, col: 7, offset: 3566},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 142, col: 1, offset: 3572},
			expr: &notExpr{
				pos: position{line: 142, col: 8, offset: 3579},
				expr: &anyMatcher{
					line: 142, col: 9, offset: 3580,
				},
			},
		},
	},
}

func (c *current) onStatus1(infos interface{}) (interface{}, error) {

	status := Status{}
	for _, datas := range infos.([]interface{}) {
		if datas != nil {
			for _, data := range datas.([]interface{}) {
				switch tup := data.(type) {
				case tuple:
					switch tup.field {
					case "allocated":
						status.Allocated = tup.value.(int)
					case "stored":
						status.Stored = tup.value.(int)
					case "used":
						status.Used = tup.value.(int)
					case "deleted":
						status.Deleted = tup.value.(int)
					case "sent":
						status.Sent = tup.value.(int)
					case "pending":
						status.Pending = tup.value.(int)
					case "failed":
						status.Failed = tup.value.(int)
					case "laststatus":
						status.LastStatus = tup.value.(string)
					}
				}
			}
		}
	}
	return status, nil
}

func (p *parser) callonStatus1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStatus1(stack["infos"])
}

func (c *current) onIncomingInfo1(data interface{}) (interface{}, error) {
	return data, nil
}

func (p *parser) callonIncomingInfo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIncomingInfo1(stack["data"])
}

func (c *current) onStored1(stored interface{}) (interface{}, error) {

	return tuple{field: "stored", value: stored.(int)}, nil
}

func (p *parser) callonStored1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onStored1(stack["stored"])
}

func (c *current) onDeleted1(deleted interface{}) (interface{}, error) {

	return tuple{field: "deleted", value: deleted.(int)}, nil
}

func (p *parser) callonDeleted1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDeleted1(stack["deleted"])
}

func (c *current) onAllocated1(allocated interface{}) (interface{}, error) {

	return tuple{field: "allocated", value: allocated.(int)}, nil
}

func (p *parser) callonAllocated1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAllocated1(stack["allocated"])
}

func (c *current) onUsed1(used interface{}) (interface{}, error) {

	return tuple{field: "used", value: used.(int)}, nil
}

func (p *parser) callonUsed1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUsed1(stack["used"])
}

func (c *current) onOutgoingInfo1(data interface{}) (interface{}, error) {
	return data, nil
}

func (p *parser) callonOutgoingInfo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onOutgoingInfo1(stack["data"])
}

func (c *current) onSent1(sent interface{}) (interface{}, error) {

	return tuple{field: "sent", value: sent.(int)}, nil
}

func (p *parser) callonSent1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSent1(stack["sent"])
}

func (c *current) onPending1(pending interface{}) (interface{}, error) {

	return tuple{field: "pending", value: pending.(int)}, nil
}

func (p *parser) callonPending1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPending1(stack["pending"])
}

func (c *current) onFailed1(failed interface{}) (interface{}, error) {

	return tuple{field: "failed", value: failed.(int)}, nil
}

func (p *parser) callonFailed1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFailed1(stack["failed"])
}

func (c *current) onLastStatus1(status interface{}) (interface{}, error) {

	return tuple{field: "laststatus", value: status.(string)}, nil
}

func (p *parser) callonLastStatus1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLastStatus1(stack["status"])
}

func (c *current) onUnhandledInfo1() (interface{}, error) {

	return nil, nil
}

func (p *parser) callonUnhandledInfo1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUnhandledInfo1()
}

func (c *current) onSeparatorLine1() (interface{}, error) {

	return nil, nil
}

func (p *parser) callonSeparatorLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSeparatorLine1()
}

func (c *current) onEmptyLine1() (interface{}, error) {

	return nil, nil
}

func (p *parser) callonEmptyLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onEmptyLine1()
}

func (c *current) onAnyLine1() (interface{}, error) {

	return nil, nil
}

func (p *parser) callonAnyLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onAnyLine1()
}

func (c *current) onNonEmptyLine5(str interface{}) (bool, error) {

	for _, r := range str.(string) {
		if !unicode.IsSpace(r) {
			return true, nil
		}
	}
	return false, nil
}

func (p *parser) callonNonEmptyLine5() (bool, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNonEmptyLine5(stack["str"])
}

func (c *current) onLineString1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonLineString1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLineString1()
}

func (c *current) onInteger1() (interface{}, error) {

	n, err := strconv.ParseInt(string(c.text), 10, 0)
	return int(n), err
}

func (p *parser) callonInteger1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInteger1()
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEntrypoint is returned when the specified entrypoint rule
	// does not exit.
	errInvalidEntrypoint = errors.New("invalid entrypoint")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errMaxExprCnt is used to signal that the maximum number of
	// expressions have been parsed.
	errMaxExprCnt = errors.New("max number of expresssions parsed")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// MaxExpressions creates an Option to stop parsing after the provided
// number of expressions have been parsed, if the value is 0 then the parser will
// parse for as many steps as needed (possibly an infinite number).
//
// The default for maxExprCnt is 0.
func MaxExpressions(maxExprCnt uint64) Option {
	return func(p *parser) Option {
		oldMaxExprCnt := p.maxExprCnt
		p.maxExprCnt = maxExprCnt
		return MaxExpressions(oldMaxExprCnt)
	}
}

// Entrypoint creates an Option to set the rule name to use as entrypoint.
// The rule name must have been specified in the -alternate-entrypoints
// if generating the parser with the -optimize-grammar flag, otherwise
// it may have been optimized out. Passing an empty string sets the
// entrypoint to the first rule in the grammar.
//
// The default is to start parsing at the first rule in the grammar.
func Entrypoint(ruleName string) Option {
	return func(p *parser) Option {
		oldEntrypoint := p.entrypoint
		p.entrypoint = ruleName
		if ruleName == "" {
			p.entrypoint = g.rules[0].name
		}
		return Entrypoint(oldEntrypoint)
	}
}

// Statistics adds a user provided Stats struct to the parser to allow
// the user to process the results after the parsing has finished.
// Also the key for the "no match" counter is set.
//
// Example usage:
//
//     input := "input"
//     stats := Stats{}
//     _, err := Parse("input-file", []byte(input), Statistics(&stats, "no match"))
//     if err != nil {
//         log.Panicln(err)
//     }
//     b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
//     if err != nil {
//         log.Panicln(err)
//     }
//     fmt.Println(string(b))
//
func Statistics(stats *Stats, choiceNoMatch string) Option {
	return func(p *parser) Option {
		oldStats := p.Stats
		p.Stats = stats
		oldChoiceNoMatch := p.choiceNoMatch
		p.choiceNoMatch = choiceNoMatch
		if p.Stats.ChoiceAltCnt == nil {
			p.Stats.ChoiceAltCnt = make(map[string]map[string]int)
		}
		return Statistics(oldStats, oldChoiceNoMatch)
	}
}

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// AllowInvalidUTF8 creates an Option to allow invalid UTF-8 bytes.
// Every invalid UTF-8 byte is treated as a utf8.RuneError (U+FFFD)
// by character class matchers and is matched by the any matcher.
// The returned matched value, c.text and c.offset are NOT affected.
//
// The default is false.
func AllowInvalidUTF8(b bool) Option {
	return func(p *parser) Option {
		old := p.allowInvalidUTF8
		p.allowInvalidUTF8 = b
		return AllowInvalidUTF8(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// GlobalStore creates an Option to set a key to a certain value in
// the globalStore.
func GlobalStore(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.globalStore[key]
		p.cur.globalStore[key] = value
		return GlobalStore(key, old)
	}
}

// InitState creates an Option to set a key to a certain value in
// the global "state" store.
func InitState(key string, value interface{}) Option {
	return func(p *parser) Option {
		old := p.cur.state[key]
		p.cur.state[key] = value
		return InitState(key, old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (i interface{}, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match

	// state is a store for arbitrary key,value pairs that the user wants to be
	// tied to the backtracking of the parser.
	// This is always rolled back if a parsing rule fails.
	state storeDict

	// globalStore is a general store for the user to store arbitrary key-value
	// pairs that they need to manage and that they do not want tied to the
	// backtracking of the parser. This is only modified by the user and never
	// rolled back by the parser. It is always up to the user to keep this in a
	// consistent state.
	globalStore storeDict
}

type storeDict map[string]interface{}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type recoveryExpr struct {
	pos          position
	expr         interface{}
	recoverExpr  interface{}
	failureLabel []string
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type throwExpr struct {
	pos   position
	label string
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type stateCodeExpr struct {
	pos position
	run func(*parser) error
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos             position
	val             string
	basicLatinChars [128]bool
	chars           []rune
	ranges          []rune
	classes         []*unicode.RangeTable
	ignoreCase      bool
	inverted        bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner    error
	pos      position
	prefix   string
	expected []string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	stats := Stats{
		ChoiceAltCnt: make(map[string]map[string]int),
	}

	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
		cur: current{
			state:       make(storeDict),
			globalStore: make(storeDict),
		},
		maxFailPos:      position{col: 1, line: 1},
		maxFailExpected: make([]string, 0, 20),
		Stats:           &stats,
		// start rule is rule [0] unless an alternate entrypoint is specified
		entrypoint: g.rules[0].name,
		emptyState: make(storeDict),
	}
	p.setOptions(opts)

	if p.maxExprCnt == 0 {
		p.maxExprCnt = math.MaxUint64
	}

	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

const choiceNoMatch = -1

// Stats stores some statistics, gathered during parsing
type Stats struct {
	// ExprCnt counts the number of expressions processed during parsing
	// This value is compared to the maximum number of expressions allowed
	// (set by the MaxExpressions option).
	ExprCnt uint64

	// ChoiceAltCnt is used to count for each ordered choice expression,
	// which alternative is used how may times.
	// These numbers allow to optimize the order of the ordered choice expression
	// to increase the performance of the parser
	//
	// The outer key of ChoiceAltCnt is composed of the name of the rule as well
	// as the line and the column of the ordered choice.
	// The inner key of ChoiceAltCnt is the number (one-based) of the matching alternative.
	// For each alternative the number of matches are counted. If an ordered choice does not
	// match, a special counter is incremented. The name of this counter is set with
	// the parser option Statistics.
	// For an alternative to be included in ChoiceAltCnt, it has to match at least once.
	ChoiceAltCnt map[string]map[string]int
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	depth   int
	recover bool
	debug   bool

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// parse fail
	maxFailPos            position
	maxFailExpected       []string
	maxFailInvertExpected bool

	// max number of expressions to be parsed
	maxExprCnt uint64
	// entrypoint for the parser
	entrypoint string

	allowInvalidUTF8 bool

	*Stats

	choiceNoMatch string
	// recovery expression stack, keeps track of the currently available recovery expression, these are traversed in reverse
	recoveryStack []map[string]interface{}

	// emptyState contains an empty storeDict, which is used to optimize cloneState if global "state" store is not used.
	emptyState storeDict
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

// push a recovery expression with its labels to the recoveryStack
func (p *parser) pushRecovery(labels []string, expr interface{}) {
	if cap(p.recoveryStack) == len(p.recoveryStack) {
		// create new empty slot in the stack
		p.recoveryStack = append(p.recoveryStack, nil)
	} else {
		// slice to 1 more
		p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)+1]
	}

	m := make(map[string]interface{}, len(labels))
	for _, fl := range labels {
		m[fl] = expr
	}
	p.recoveryStack[len(p.recoveryStack)-1] = m
}

// pop a recovery expression from the recoveryStack
func (p *parser) popRecovery() {
	// GC that map
	p.recoveryStack[len(p.recoveryStack)-1] = nil

	p.recoveryStack = p.recoveryStack[:len(p.recoveryStack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position, []string{})
}

func (p *parser) addErrAt(err error, pos position, expected []string) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String(), expected: expected}
	p.errs.add(pe)
}

func (p *parser) failAt(fail bool, pos position, want string) {
	// process fail if parsing fails and not inverted or parsing succeeds and invert is set
	if fail == p.maxFailInvertExpected {
		if pos.offset < p.maxFailPos.offset {
			return
		}

		if pos.offset > p.maxFailPos.offset {
			p.maxFailPos = pos
			p.maxFailExpected = p.maxFailExpected[:0]
		}

		if p.maxFailInvertExpected {
			want = "!" + want
		}
		p.maxFailExpected = append(p.maxFailExpected, want)
	}
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError && n == 1 { // see utf8.DecodeRune
		if !p.allowInvalidUTF8 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// Cloner is implemented by any value that has a Clone method, which returns a
// copy of the value. This is mainly used for types which are not passed by
// value (e.g map, slice, chan) or structs that contain such types.
//
// This is used in conjunction with the global state feature to create proper
// copies of the state to allow the parser to properly restore the state in
// the case of backtracking.
type Cloner interface {
	Clone() interface{}
}

// clone and return parser current state.
func (p *parser) cloneState() storeDict {
	if p.debug {
		defer p.out(p.in("cloneState"))
	}

	if len(p.cur.state) == 0 {
		if len(p.emptyState) > 0 {
			p.emptyState = make(storeDict)
		}
		return p.emptyState
	}

	state := make(storeDict, len(p.cur.state))
	for k, v := range p.cur.state {
		if c, ok := v.(Cloner); ok {
			state[k] = c.Clone()
		} else {
			state[k] = v
		}
	}
	return state
}

// restore parser current state to the state storeDict.
// every restoreState should applied only one time for every cloned state
func (p *parser) restoreState(state storeDict) {
	if p.debug {
		defer p.out(p.in("restoreState"))
	}
	p.cur.state = state
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	startRule, ok := p.rules[p.entrypoint]
	if !ok {
		p.addErr(errInvalidEntrypoint)
		return nil, p.errs.err()
	}

	p.read() // advance to first rune
	val, ok = p.parseRule(startRule)
	if !ok {
		if len(*p.errs) == 0 {
			// If parsing fails, but no errors have been recorded, the expected values
			// for the farthest parser position are returned as error.
			maxFailExpectedMap := make(map[string]struct{}, len(p.maxFailExpected))
			for _, v := range p.maxFailExpected {
				maxFailExpectedMap[v] = struct{}{}
			}
			expected := make([]string, 0, len(maxFailExpectedMap))
			eof := false
			if _, ok := maxFailExpectedMap["!."]; ok {
				delete(maxFailExpectedMap, "!.")
				eof = true
			}
			for k := range maxFailExpectedMap {
				expected = append(expected, k)
			}
			sort.Strings(expected)
			if eof {
				expected = append(expected, "EOF")
			}
			p.addErrAt(errors.New("no match found, expected: "+listJoin(expected, ", ", "or")), p.maxFailPos, expected)
		}

		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func listJoin(list []string, sep string, lastSep string) string {
	switch len(list) {
	case 0:
		return ""
	case 1:
		return list[0]
	default:
		return fmt.Sprintf("%s %s %s", strings.Join(list[:len(list)-1], sep), lastSep, list[len(list)-1])
	}
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.ExprCnt++
	if p.ExprCnt > p.maxExprCnt {
		panic(errMaxExprCnt)
	}

	var val interface{}
	var ok bool
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *recoveryExpr:
		val, ok = p.parseRecoveryExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *stateCodeExpr:
		val, ok = p.parseStateCodeExpr(expr)
	case *throwExpr:
		val, ok = p.parseThrowExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		state := p.cloneState()
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position, []string{})
		}
		p.restoreState(state)

		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	state := p.cloneState()

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn == utf8.RuneError && p.pt.w == 0 {
		// EOF - see utf8.DecodeRune
		p.failAt(false, p.pt.position, ".")
		return nil, false
	}
	start := p.pt
	p.read()
	p.failAt(true, start.position, ".")
	return p.sliceFrom(start), true
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	start := p.pt

	// can't match EOF
	if cur == utf8.RuneError && p.pt.w == 0 { // see utf8.DecodeRune
		p.failAt(false, start.position, chr.val)
		return nil, false
	}

	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				p.failAt(false, start.position, chr.val)
				return nil, false
			}
			p.read()
			p.failAt(true, start.position, chr.val)
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		p.failAt(true, start.position, chr.val)
		return p.sliceFrom(start), true
	}
	p.failAt(false, start.position, chr.val)
	return nil, false
}

func (p *parser) incChoiceAltCnt(ch *choiceExpr, altI int) {
	choiceIdent := fmt.Sprintf("%s %d:%d", p.rstack[len(p.rstack)-1].name, ch.pos.line, ch.pos.col)
	m := p.ChoiceAltCnt[choiceIdent]
	if m == nil {
		m = make(map[string]int)
		p.ChoiceAltCnt[choiceIdent] = m
	}
	// We increment altI by 1, so the keys do not start at 0
	alt := strconv.Itoa(altI + 1)
	if altI == choiceNoMatch {
		alt = p.choiceNoMatch
	}
	m[alt]++
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for altI, alt := range ch.alternatives {
		// dummy assignment to prevent compile error if optimized
		_ = altI

		state := p.cloneState()

		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			p.incChoiceAltCnt(ch, altI)
			return val, ok
		}
		p.restoreState(state)
	}
	p.incChoiceAltCnt(ch, choiceNoMatch)
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	ignoreCase := ""
	if lit.ignoreCase {
		ignoreCase = "i"
	}
	val := fmt.Sprintf("%q%s", lit.val, ignoreCase)
	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.failAt(false, start.position, val)
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	p.failAt(true, start.position, val)
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	state := p.cloneState()

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	p.restoreState(state)

	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	state := p.cloneState()
	p.pushV()
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	_, ok := p.parseExpr(not.expr)
	p.maxFailInvertExpected = !p.maxFailInvertExpected
	p.popV()
	p.restoreState(state)
	p.restore(pt)

	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRecoveryExpr(recover *recoveryExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRecoveryExpr (" + strings.Join(recover.failureLabel, ",") + ")"))
	}

	p.pushRecovery(recover.failureLabel, recover.recoverExpr)
	val, ok := p.parseExpr(recover.expr)
	p.popRecovery()

	return val, ok
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	vals := make([]interface{}, 0, len(seq.exprs))

	pt := p.pt
	state := p.cloneState()
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restoreState(state)
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseStateCodeExpr(state *stateCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseStateCodeExpr"))
	}

	err := state.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, true
}

func (p *parser) parseThrowExpr(expr *throwExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseThrowExpr"))
	}

	for i := len(p.recoveryStack) - 1; i >= 0; i-- {
		if recoverExpr, ok := p.recoveryStack[i][expr.label]; ok {
			if val, ok := p.parseExpr(recoverExpr); ok {
				return val, ok
			}
		}
	}

	return nil, false
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}
