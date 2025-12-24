package vhtml

import (
	"bytes"
	"context"
	"crypto/rand"
	"math/big"

	"github.com/volts-dev/lexer"
	"github.com/volts-dev/vertex/core/console"
)

const (
	HTML_RESULT ResultType = iota
	SVG_RESULT
	MATHML_RESULT
)

type (
	TemplatePartType int // TemplatePart types
	ResultType       int // TemplateResult types

	ExpressionPart struct {
		Name  string
		Value string
		Type  TemplatePartType
	}

	IfConditonPart struct {
		value       string
		contentIdx1 int
		contentIdx2 int
		endIdx      int
	}

	RangeConditonPart struct {
		value      string
		contentIdx int
		endIdx     int
	}

	TemplateParser struct {
		id   string
		html bytes.Buffer
		//attributeNames []string
		parts []*ExpressionPart
		ctx   context.Context
	}
)

const (
	// Base62 字符集 (去掉了容易混淆的符号)
	letters              = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	boundAttributePrefix = "$vtx"
)

var marker string

func init() {
	marker, _ = GenerateShortUID(8)
	marker = "$vtx" + marker
}

func GenerateShortUID(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}
	return string(ret), nil
}

func getTemplateHtml(content string, ctx ...context.Context) (string, []*ExpressionPart) {
	var c context.Context
	if len(ctx) == 0 {
		c = context.Background()
	} else {
		c = ctx[0]
	}

	parser := newParser(content)
	tmpl := newTemplateParser(c)
	err := tmpl.parseHtml(parser)
	if err != nil {
		console.Error("HTML parse error: %v", err)
		return "", nil
	}

	return tmpl.html.String(), tmpl.parts
}

func (self *EventPart) String() string {
	return "@" + self.name + "=\"" + self.method + "\""
}

func newTemplateParser(ctx context.Context) *TemplateParser {
	return &TemplateParser{ctx: ctx}
}

func (self *TemplateParser) parseExpr(parser *Parser, inElement bool) error {
	// TODO 处理表达式

	// TODO 处理<Element {{ 变量
	parser.SkipSpace()
	parser.Next()

	switch v := parser.Token(); v.Type {
	// TODO 处理 .属性变量 接收一个参数
	case lexer.PERIOD:
		parser.Next()
		attrName := parser.Token().Val
		parser.SkipSpace()
		parser.Next()

		part := &ExpressionPart{
			Type:  ATTRIBUTE_PART,
			Name:  attrName,
			Value: parser.Token().Val,
		}

		attrName = "." + attrName

		//self.attributeNames = append(self.attributeNames, attrName)
		if inElement {
			self.html.WriteString(attrName)
			self.html.WriteString(boundAttributePrefix)
			self.html.WriteString("=")
			self.html.WriteString("\"" + marker + "\"")
		} else {
			self.html.WriteString("<!--" + marker + ">")
		}

		self.parts = append(self.parts, part)
	// TODO 处理 @变量
	case lexer.AT, lexer.HOLDER:
		symbol := parser.Token().Val
		parser.Next()
		attrName := parser.Token().Val

		parser.SkipSpace()
		parser.Next()

		part := &ExpressionPart{
			Type:  EVENT_PART,
			Name:  attrName,
			Value: parser.Token().Val,
		}

		if v.Type == lexer.HOLDER {
			part.Type = BOOLEAN_ATTRIBUTE_PART
		}

		attrName = symbol + attrName
		//self.attributeNames = append(self.attributeNames, attrName)
		self.html.WriteString(attrName + boundAttributePrefix + "=\"" + marker + "\"")
		self.parts = append(self.parts, part)

	//  TODO 处理 变量
	case lexer.IDENT:
		switch v.Val {
		case "if":
			parser.SkipSpace()
			parser.Next()

			part := &ExpressionPart{
				Type:  IF_PART,
				Name:  "if",
				Value: parser.Token().Val,
			}

			self.html.WriteString("<!--if")
			self.html.WriteString(marker)
			self.html.WriteString(">")
			parser.AcceptUntil(func(t *lexer.TToken) bool { return t.Type == lexer.RBRACE }) //}
			parser.Next()
			parser.Next()
			/*ifpart, err := self.parseIf(parser)
			if err != nil {
				return err
			}
			 self.ctx = context.WithValue(self.ctx, "expr", ifpart)
			*/
			self.parts = append(self.parts, part)

			return nil
		case "else":
			self.html.WriteString("<!--else")
			self.html.WriteString(marker)
			self.html.WriteString(">")

			parser.AcceptUntil(func(t *lexer.TToken) bool { return t.Type == lexer.RBRACE }) //}
			parser.Next()
			parser.Next()
			parser.Next()
			//contentIdx := len(self.parts)
			if err := self.parseHtml(parser); err != nil {
				return err
			}

			//ifpart := self.ctx.Value("expr").(*IfConditonPart)
			//ifpart.contentIdx2 = contentIdx

			return nil
			//	case "range":
			//	case "break":
			//	case "continue":
		case "end":
			self.html.WriteString("<!--end")
			self.html.WriteString(marker)
			self.html.WriteString(">")

			parser.AcceptUntil(func(t *lexer.TToken) bool { return t.Type == lexer.RBRACE }) //}
			parser.Next()
			parser.Next()

			/*
				switch part := self.ctx.Value("expr").(type) {
				case *IfConditonPart:
					part.endIdx = len(self.parts)
				case *RangeConditonPart:
					part.endIdx = len(self.parts)
				}*/

			return nil
		default:
			self.html.WriteString(marker)
			//logger.Dbg("Unknown expr:", v.Val)
		}
	}

	// 直接跳到}}
	parser.AcceptUntil(func(t *lexer.TToken) bool { return t.Type == lexer.RBRACE }) //}
	parser.Next()
	parser.Next()
	return nil
}

func (self *TemplateParser) parseElement(parser *Parser) error {
	var (
		elementName string
		attrCentent bytes.Buffer
	)

	for !parser.IsEnd() {
		token := parser.Token()

		switch token.Type {
		case lexer.IDENT:
			// Tag Name
			elementName = parser.Token().Val
			attrCentent.WriteString(elementName)
			//self.html.WriteString(elementName)
			//attrCentent.WriteByte(' ')

		// 结束标签 />
		case lexer.QUO:
			attrCentent.WriteByte('/')
			//self.html.WriteByte('/')

			parser.Next()
			if token.Type == lexer.GTR {
				attrCentent.WriteByte('>')
				//self.html.WriteByte('>')
				goto RETURN
			}
			continue

		// 解析元素里的内容，这是另一个 html 片段
		case lexer.GTR:
			attrCentent.WriteByte('>')
			//self.html.WriteByte('>')
			goto RETURN

			// {{符号+变量 参数...}
		case lexer.LBRACE:
			parser.Next()
			if token.Type == lexer.LBRACE {
				// 保存之前的文本内容
				if attrCentent.Len() > 0 {
					self.html.Write(attrCentent.Bytes())
					//self.push(attrCentent.Bytes())
					attrCentent.Reset()
				}

				self.parseExpr(parser, true)

				goto NEXT
			}
			fallthrough
		default:
			attrCentent.WriteString(token.Val)
		}

	NEXT:
		parser.Next()
	}

RETURN:
	if attrCentent.Len() > 0 {
		self.html.Write(attrCentent.Bytes())
		//self.push(attrCentent.Bytes())
	}
	return nil
}

func (self *TemplateParser) parseHtml(parser *Parser) error {
	var htmlText bytes.Buffer

	for !parser.IsEnd() {
		item := parser.Token()

		switch item.Type {
		case lexer.LSS:
			if htmlText.Len() > 0 {
				self.html.Write(htmlText.Bytes())
				//self.push(htmlText.Bytes())
				htmlText.Reset()
			}

			self.parseElement(parser)

		case lexer.LBRACE:
			parser.Next()
			if parser.Token().Type == lexer.LBRACE {
				if htmlText.Len() > 0 {
					self.html.Write(htmlText.Bytes())
					//self.push(htmlText.Bytes())
					htmlText.Reset()
				}

				self.parseExpr(parser, false)
			}
		default:
			htmlText.WriteString(item.Val)
		}

		parser.Next()
	}

	if htmlText.Len() > 0 {
		self.html.Write(htmlText.Bytes())
		//self.push(htmlText.Bytes())
		htmlText.Reset()
	}

	return nil
}
func (self *TemplateParser) Id() string {
	return self.id
}
func (self *TemplateParser) ToBytes() []byte {
	return self.html.Bytes()
}
func (self *TemplateParser) ToString() string {
	return self.html.String()
}
