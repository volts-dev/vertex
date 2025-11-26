package vhtml

import (
	"strings"

	"github.com/volts-dev/lexer"
	"github.com/volts-dev/logger"
)

const (
	// 默认的标签
	DLBRACE = 991 // {{
	DRBRACE = 992 // }}
)

var printToken bool = true // print token

// Tree is the representation of a single parsed template.
type (
	// Parser 提供解析的逻辑 和数据的变换
	Parser struct {
		lexer     *lexer.TLexer
		pos       int
		len       int
		isEnd     bool
		Tokens    []*lexer.TToken
		LastToken *lexer.TToken
	}

	ValidatorFn func(*lexer.TToken) bool
)

// TODO 支持<a/>格式
// Creates a new parser to parse tokens.
// Used inside pongo2 to parse documents and to provide an easy-to-use
// parser for tag authors
func newParser(content string) *Parser {
	lexer.RegisterState("<", func(l *lexer.TLexer) lexer.StateFn {
		l.Next()
		l.Emit(lexer.LSS)
		return lexer.NextState
	})
	lexer.RegisterState(">", func(l *lexer.TLexer) lexer.StateFn {
		l.Next()
		l.Emit(lexer.GTR)
		return lexer.NextState
	})
	lexer.RegisterState("~{{", func(l *lexer.TLexer) lexer.StateFn {
		l.Next()
		l.Emit(DLBRACE)
		return lexer.NextState
	})
	lexer.RegisterState("~}}", func(l *lexer.TLexer) lexer.StateFn {
		l.Next()
		l.Emit(DRBRACE)
		return lexer.NextState
	})

	lex, err := lexer.NewLexer(strings.NewReader(content))
	if err != nil {
		logger.Err(err.Error())
	}

	parser := &Parser{
		lexer: lex,
		pos:   0,
		len:   0,
	}

	for {
		item, ok := <-lex.Tokens
		if !ok {
			break
		}

		// print token
		if printToken {
			logger.Info(lexer.PrintToken(item))
		}

		parser.Tokens = append(parser.Tokens, &item)
	}

	parser.len = len(parser.Tokens)
	if parser.len > 0 {
		parser.LastToken = parser.Tokens[parser.len-1]
	}
	return parser
}

// 主要-略过特殊字符移动
// 并返回不符合条件的Item
// 回退Pos 到空白Item处,保持下一个有效字符
func (self *Parser) SkipSpace() (item *lexer.TToken) {
	for {
		self.Next()
		if self.isEnd {
			break
		}
		//fmt.Println("consume_whitespace", self.Item().Val)
		switch self.Token().Type {
		case lexer.SAPCE: //lexer.TokenWhitespace,
			continue
		default:
			item = self.Token()
			self.Backup()
			goto exit
		}
	}
exit:
	//fmt.Println("exit consume_whitespace", self.Item().Val)
	return
}

func (self *Parser) AcceptUntil(fn ValidatorFn) int {
	self.Next()
	count := 0
	for !fn(self.Token()) && !self.IsEnd() {
		self.Next()
		count++
	}
	self.Backup()
	return count
}

func (self *Parser) Backup(cnt ...int) {
	count := 1
	if len(cnt) > 0 {
		count = cnt[0]
	}

	//fmt.Println("Backup", (self.len-count) > 0)
	if (self.len - count) > 0 {
		self.pos = self.pos - count
	}
}

func (self *Parser) Next() {
	if self.pos >= self.len-1 { //如果大于Buf 则停止
		self.isEnd = true
		return
	}

	self.pos++
}

func (self *Parser) IsEnd() bool {
	return self.isEnd
}

func (self *Parser) Token() *lexer.TToken {
	return self.Tokens[self.pos]
}

func (self *Parser) Find(i lexer.TToken) int {
	return 0
}

func (self *Parser) Length() int {
	return self.len
}

// Returns the UNCONSUMED token count.
func (self *Parser) Remaining() int {
	return len(self.Tokens) - self.pos
}
