package table

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
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/satori/go.uuid"
)

// NOTE: remember to update the Equals if fields change
type SMS struct {
	UUID     uuid.UUID `json:"uuid"`
	Id       int       `json:"id"`
	Received time.Time `json:"received"`
	From     string    `json:"from"`
	Size     int       `json:"size"`
	Text     string    `json:"text"`
}

func linesToString(line interface{}, lines interface{}) string {
	var text string
	switch first := line.(type) {
	case string:
		text = first
	case []interface{}:
		text = first[0].(string)
	}
	switch extras := lines.(type) {
	case []interface{}:
		for _, extra := range extras {
			switch other := extra.(type) {
			case string:
				text += "\n" + other
			case []interface{}:
				text += "\n" + other[0].(string)
			}
		}
	}
	return text
}

var g = &grammar{
	rules: []*rule{
		{
			name: "Table",
			pos:  position{line: 44, col: 1, offset: 954},
			expr: &actionExpr{
				pos: position{line: 44, col: 10, offset: 963},
				run: (*parser).callonTable1,
				expr: &seqExpr{
					pos: position{line: 44, col: 10, offset: 963},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 44, col: 10, offset: 963},
							name: "__",
						},
						&labeledExpr{
							pos:   position{line: 44, col: 13, offset: 966},
							label: "result",
							expr: &oneOrMoreExpr{
								pos: position{line: 44, col: 20, offset: 973},
								expr: &choiceExpr{
									pos: position{line: 44, col: 21, offset: 974},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 44, col: 21, offset: 974},
											name: "SMS",
										},
										&ruleRefExpr{
											pos:  position{line: 44, col: 27, offset: 980},
											name: "Line",
										},
									},
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 44, col: 34, offset: 987},
							name: "__",
						},
						&ruleRefExpr{
							pos:  position{line: 44, col: 37, offset: 990},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "SMS",
			pos:  position{line: 59, col: 1, offset: 1310},
			expr: &actionExpr{
				pos: position{line: 59, col: 8, offset: 1317},
				run: (*parser).callonSMS1,
				expr: &seqExpr{
					pos: position{line: 59, col: 8, offset: 1317},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 59, col: 8, offset: 1317},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 11, offset: 1320},
								name: "ID",
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 14, offset: 1323},
							label: "received",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 23, offset: 1332},
								name: "Received",
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 32, offset: 1341},
							label: "from",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 37, offset: 1346},
								name: "From",
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 42, offset: 1351},
							label: "size",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 47, offset: 1356},
								name: "Size",
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 52, offset: 1361},
							label: "line",
							expr: &ruleRefExpr{
								pos:  position{line: 59, col: 57, offset: 1366},
								name: "Line",
							},
						},
						&labeledExpr{
							pos:   position{line: 59, col: 62, offset: 1371},
							label: "lines",
							expr: &zeroOrMoreExpr{
								pos: position{line: 59, col: 68, offset: 1377},
								expr: &ruleRefExpr{
									pos:  position{line: 59, col: 68, offset: 1377},
									name: "NonSeparatorLine",
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 86, offset: 1395},
							name: "SeparatorLine",
						},
					},
				},
			},
		},
		{
			name: "ID",
			pos:  position{line: 69, col: 1, offset: 1586},
			expr: &actionExpr{
				pos: position{line: 69, col: 7, offset: 1592},
				run: (*parser).callonID1,
				expr: &seqExpr{
					pos: position{line: 69, col: 7, offset: 1592},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 69, col: 7, offset: 1592},
							val:        "sms id",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 17, offset: 1602},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 69, col: 19, offset: 1604},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 23, offset: 1608},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 69, col: 25, offset: 1610},
							label: "id",
							expr: &ruleRefExpr{
								pos:  position{line: 69, col: 28, offset: 1613},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 36, offset: 1621},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 38, offset: 1623},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Received",
			pos:  position{line: 73, col: 1, offset: 1652},
			expr: &actionExpr{
				pos: position{line: 73, col: 13, offset: 1664},
				run: (*parser).callonReceived1,
				expr: &seqExpr{
					pos: position{line: 73, col: 13, offset: 1664},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 73, col: 13, offset: 1664},
							val:        "time",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 73, col: 21, offset: 1672},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 73, col: 23, offset: 1674},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 73, col: 27, offset: 1678},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 73, col: 29, offset: 1680},
							label: "received",
							expr: &ruleRefExpr{
								pos:  position{line: 73, col: 38, offset: 1689},
								name: "DateTime",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 73, col: 47, offset: 1698},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 73, col: 49, offset: 1700},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "From",
			pos:  position{line: 77, col: 1, offset: 1735},
			expr: &actionExpr{
				pos: position{line: 77, col: 9, offset: 1743},
				run: (*parser).callonFrom1,
				expr: &seqExpr{
					pos: position{line: 77, col: 9, offset: 1743},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 77, col: 9, offset: 1743},
							val:        "from",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 17, offset: 1751},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 77, col: 19, offset: 1753},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 23, offset: 1757},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 77, col: 25, offset: 1759},
							label: "from",
							expr: &ruleRefExpr{
								pos:  position{line: 77, col: 30, offset: 1764},
								name: "PhoneNumber",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 42, offset: 1776},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 44, offset: 1778},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Size",
			pos:  position{line: 81, col: 1, offset: 1809},
			expr: &actionExpr{
				pos: position{line: 81, col: 9, offset: 1817},
				run: (*parser).callonSize1,
				expr: &seqExpr{
					pos: position{line: 81, col: 9, offset: 1817},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 81, col: 9, offset: 1817},
							val:        "size",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 17, offset: 1825},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 81, col: 19, offset: 1827},
							val:        ":",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 23, offset: 1831},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 81, col: 25, offset: 1833},
							label: "size",
							expr: &ruleRefExpr{
								pos:  position{line: 81, col: 30, offset: 1838},
								name: "Integer",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 38, offset: 1846},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 81, col: 40, offset: 1848},
							name: "NL",
						},
					},
				},
			},
		},
		{
			name: "Line",
			pos:  position{line: 85, col: 1, offset: 1879},
			expr: &actionExpr{
				pos: position{line: 85, col: 9, offset: 1887},
				run: (*parser).callonLine1,
				expr: &labeledExpr{
					pos:   position{line: 85, col: 9, offset: 1887},
					label: "line",
					expr: &choiceExpr{
						pos: position{line: 85, col: 15, offset: 1893},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 85, col: 15, offset: 1893},
								name: "NonSeparatorLine",
							},
							&ruleRefExpr{
								pos:  position{line: 85, col: 34, offset: 1912},
								name: "SeparatorLine",
							},
						},
					},
				},
			},
		},
		{
			name: "NonSeparatorLine",
			pos:  position{line: 89, col: 1, offset: 1955},
			expr: &seqExpr{
				pos: position{line: 89, col: 21, offset: 1975},
				exprs: []interface{}{
					&labeledExpr{
						pos:   position{line: 89, col: 21, offset: 1975},
						label: "line",
						expr: &ruleRefExpr{
							pos:  position{line: 89, col: 26, offset: 1980},
							name: "LineString",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 37, offset: 1991},
						name: "NL",
					},
					&andCodeExpr{
						pos: position{line: 89, col: 40, offset: 1994},
						run: (*parser).callonNonSeparatorLine5,
					},
				},
			},
		},
		{
			name: "LineString",
			pos:  position{line: 102, col: 1, offset: 2178},
			expr: &actionExpr{
				pos: position{line: 102, col: 15, offset: 2192},
				run: (*parser).callonLineString1,
				expr: &zeroOrMoreExpr{
					pos: position{line: 102, col: 15, offset: 2192},
					expr: &charClassMatcher{
						pos:        position{line: 102, col: 15, offset: 2192},
						val:        "[^\\n]",
						chars:      []rune{'\n'},
						ignoreCase: false,
						inverted:   true,
					},
				},
			},
		},
		{
			name: "SeparatorLine",
			pos:  position{line: 106, col: 1, offset: 2237},
			expr: &actionExpr{
				pos: position{line: 106, col: 18, offset: 2254},
				run: (*parser).callonSeparatorLine1,
				expr: &seqExpr{
					pos: position{line: 106, col: 18, offset: 2254},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 106, col: 18, offset: 2254},
							expr: &litMatcher{
								pos:        position{line: 106, col: 18, offset: 2254},
								val:        "-",
								ignoreCase: false,
							},
						},
						&zeroOrOneExpr{
							pos: position{line: 106, col: 23, offset: 2259},
							expr: &ruleRefExpr{
								pos:  position{line: 106, col: 23, offset: 2259},
								name: "NL",
							},
						},
					},
				},
			},
		},
		{
			name: "Integer",
			pos:  position{line: 110, col: 1, offset: 2301},
			expr: &actionExpr{
				pos: position{line: 110, col: 12, offset: 2312},
				run: (*parser).callonInteger1,
				expr: &oneOrMoreExpr{
					pos: position{line: 110, col: 12, offset: 2312},
					expr: &ruleRefExpr{
						pos:  position{line: 110, col: 12, offset: 2312},
						name: "DecimalDigit",
					},
				},
			},
		},
		{
			name: "DateTime",
			pos:  position{line: 114, col: 1, offset: 2373},
			expr: &actionExpr{
				pos: position{line: 114, col: 13, offset: 2385},
				run: (*parser).callonDateTime1,
				expr: &seqExpr{
					pos: position{line: 114, col: 13, offset: 2385},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 114, col: 13, offset: 2385},
							name: "Date",
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 18, offset: 2390},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 114, col: 20, offset: 2392},
							name: "Time",
						},
					},
				},
			},
		},
		{
			name: "Date",
			pos:  position{line: 118, col: 1, offset: 2463},
			expr: &seqExpr{
				pos: position{line: 118, col: 9, offset: 2471},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 118, col: 9, offset: 2471},
						name: "Integer",
					},
					&litMatcher{
						pos:        position{line: 118, col: 17, offset: 2479},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 21, offset: 2483},
						name: "Integer",
					},
					&litMatcher{
						pos:        position{line: 118, col: 29, offset: 2491},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 33, offset: 2495},
						name: "Integer",
					},
				},
			},
		},
		{
			name: "Time",
			pos:  position{line: 120, col: 1, offset: 2506},
			expr: &seqExpr{
				pos: position{line: 120, col: 9, offset: 2514},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 120, col: 9, offset: 2514},
						name: "Integer",
					},
					&litMatcher{
						pos:        position{line: 120, col: 17, offset: 2522},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 120, col: 21, offset: 2526},
						name: "Integer",
					},
					&litMatcher{
						pos:        position{line: 120, col: 29, offset: 2534},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 120, col: 33, offset: 2538},
						name: "Integer",
					},
				},
			},
		},
		{
			name: "PhoneNumber",
			pos:  position{line: 122, col: 1, offset: 2549},
			expr: &actionExpr{
				pos: position{line: 122, col: 16, offset: 2564},
				run: (*parser).callonPhoneNumber1,
				expr: &seqExpr{
					pos: position{line: 122, col: 16, offset: 2564},
					exprs: []interface{}{
						&zeroOrOneExpr{
							pos: position{line: 122, col: 16, offset: 2564},
							expr: &litMatcher{
								pos:        position{line: 122, col: 16, offset: 2564},
								val:        "+",
								ignoreCase: false,
							},
						},
						&oneOrMoreExpr{
							pos: position{line: 122, col: 21, offset: 2569},
							expr: &ruleRefExpr{
								pos:  position{line: 122, col: 21, offset: 2569},
								name: "DecimalDigit",
							},
						},
					},
				},
			},
		},
		{
			name: "DecimalDigit",
			pos:  position{line: 128, col: 1, offset: 2625},
			expr: &charClassMatcher{
				pos:        position{line: 128, col: 17, offset: 2641},
				val:        "[0-9]",
				ranges:     []rune{'0', '9'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name:        "_",
			displayName: "\"WS\"",
			pos:         position{line: 129, col: 1, offset: 2648},
			expr: &zeroOrMoreExpr{
				pos: position{line: 129, col: 11, offset: 2658},
				expr: &charClassMatcher{
					pos:        position{line: 129, col: 11, offset: 2658},
					val:        "[ \\t\\r]",
					chars:      []rune{' ', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name:        "__",
			displayName: "\"WSNL\"",
			pos:         position{line: 130, col: 1, offset: 2668},
			expr: &zeroOrMoreExpr{
				pos: position{line: 130, col: 14, offset: 2681},
				expr: &charClassMatcher{
					pos:        position{line: 130, col: 14, offset: 2681},
					val:        "[ \\n\\t\\r]",
					chars:      []rune{' ', '\n', '\t', '\r'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "NL",
			pos:  position{line: 131, col: 1, offset: 2693},
			expr: &litMatcher{
				pos:        position{line: 131, col: 7, offset: 2699},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 132, col: 1, offset: 2705},
			expr: &notExpr{
				pos: position{line: 132, col: 8, offset: 2712},
				expr: &anyMatcher{
					line: 132, col: 9, offset: 2713,
				},
			},
		},
	},
}

func (c *current) onTable1(result interface{}) (interface{}, error) {

	sections := result.([]interface{})
	smses := make([]SMS, 0, len(sections))
	for _, section := range sections {
		switch sms := section.(type) {
		case SMS:
			smses = append(smses, sms)
		}
	}
	if len(smses) < 1 {
		return nil, errors.New("no sms found")
	}
	return smses, nil
}

func (p *parser) callonTable1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onTable1(stack["result"])
}

func (c *current) onSMS1(id, received, from, size, line, lines interface{}) (interface{}, error) {

	return SMS{
		Id:       id.(int),
		Received: received.(time.Time),
		From:     from.(string),
		Size:     size.(int),
		Text:     linesToString(line, lines),
	}, nil
}

func (p *parser) callonSMS1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSMS1(stack["id"], stack["received"], stack["from"], stack["size"], stack["line"], stack["lines"])
}

func (c *current) onID1(id interface{}) (interface{}, error) {

	return id, nil
}

func (p *parser) callonID1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onID1(stack["id"])
}

func (c *current) onReceived1(received interface{}) (interface{}, error) {

	return received, nil
}

func (p *parser) callonReceived1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onReceived1(stack["received"])
}

func (c *current) onFrom1(from interface{}) (interface{}, error) {

	return from, nil
}

func (p *parser) callonFrom1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFrom1(stack["from"])
}

func (c *current) onSize1(size interface{}) (interface{}, error) {

	return size, nil
}

func (p *parser) callonSize1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSize1(stack["size"])
}

func (c *current) onLine1(line interface{}) (interface{}, error) {

	return line, nil
}

func (p *parser) callonLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLine1(stack["line"])
}

func (c *current) onNonSeparatorLine5(line interface{}) (bool, error) {

	str := line.(string)
	if len(str) < 1 {
		return true, nil
	}
	for _, r := range str {
		if r != '-' {
			return true, nil
		}
	}
	return false, nil
}

func (p *parser) callonNonSeparatorLine5() (bool, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNonSeparatorLine5(stack["line"])
}

func (c *current) onLineString1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonLineString1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLineString1()
}

func (c *current) onSeparatorLine1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonSeparatorLine1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSeparatorLine1()
}

func (c *current) onInteger1() (interface{}, error) {

	return strconv.Atoi(string(c.text))
}

func (p *parser) callonInteger1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onInteger1()
}

func (c *current) onDateTime1() (interface{}, error) {

	return time.Parse("06-01-02 15:04:05", string(c.text))
}

func (p *parser) callonDateTime1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDateTime1()
}

func (c *current) onPhoneNumber1() (interface{}, error) {

	return string(c.text), nil
}

func (p *parser) callonPhoneNumber1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPhoneNumber1()
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
