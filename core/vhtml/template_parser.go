package vhtml

import (
	"context"
	"fmt"
	"strings"

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

	_ExpressionPart struct {
		Name     string
		Value    string
		Children []*_ExpressionPart
		Type     TemplatePartType
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
		html strings.Builder
		//attributeNames []string
		parts   []*TemplatePart
		context context.Context
		marker  string
	}
)

func parseTemplateHtml(content string, ctx ...context.Context) (string, []*TemplatePart) {
	var c context.Context
	if len(ctx) == 0 {
		c = context.Background()
	} else {
		c = ctx[0]
	}

	parser := newParser(content)
	tmpl := newTemplateParser(c)
	err := tmpl.parseHtml(nil, parser, false)
	if err != nil {
		console.Error("HTML parse error: %v", err)
		return "", nil
	}

	return tmpl.html.String(), tmpl.parts
}

func (self *EventPart) String() string {
	return fmt.Sprintf(`@%s = %v`, self.name, self.value)
}

func newTemplateParser(ctx context.Context) *TemplateParser {
	marker := ctx.Value("_vertex_marker").(string)
	return &TemplateParser{context: ctx, marker: marker}
}

func (self *TemplateParser) parseExprBody(parser *Parser, typ TemplatePartType) *TemplatePart {
	var buf strings.Builder
	exprs := make([]string, 0)
	for !parser.IsEnd() && parser.Token().Type != lexer.RBRACE {
		if parser.Token().Type == lexer.IDENT {
			exprs = append(exprs, parser.Token().Val)
		}
		buf.WriteString(parser.Token().Val)
		parser.Next()
	}

	if !parser.IsEnd() && parser.Token().Type == lexer.RBRACE {
		parser.Next()
	}

	part := &TemplatePart{Type: typ}

	switch typ {
	case CHILD_IF_PART:
		part.Name = strings.TrimSpace(buf.String())
		part.Strings = exprs
	default:
		//exprs := strings.Fields(strings.TrimSpace(buf.String()))
		if len(exprs) == 0 {
			return part
		}

		part.Name = exprs[0]
		if len(exprs) > 1 {
			part.Value = exprs[1]
		}
	}

	return part
}

func (self *TemplateParser) parseExpr(parent *TemplatePart, parser *Parser, parts *[]*TemplatePart, htmlText *strings.Builder, inElement bool) (bool, error) {
	parser.SkipSpace()
	parser.Next()
	var part *TemplatePart
	switch v := parser.Token(); v.Type {
	case lexer.PERIOD: // 处理 .属性变量
		parser.Next()

		if inElement {
			part = self.parseExprBody(parser, ATTRIBUTE_PART)
			htmlText.WriteString(part.Name)
			htmlText.WriteString(vertexPrefix)
			htmlText.WriteString(`="` + self.marker + `"`)
		} else {
			part = self.parseExprBody(parser, CHILD_VARIANT_PART)
			htmlText.WriteString("<!--" + self.marker + "-->")
		}

		*parts = append(*parts, part)

	case lexer.AT, lexer.HOLDER: // TODO 处理 @变量
		symbol := v.Val
		parser.Next()

		partType := BOOLEAN_ATTRIBUTE_PART
		if v.Type == lexer.AT {
			partType = EVENT_PART
		}

		part = self.parseExprBody(parser, partType)
		attrName := symbol + part.Name
		htmlText.WriteString(attrName + vertexPrefix + `="` + self.marker + `"`)
		*parts = append(*parts, part)

	case lexer.IDENT: //  TODO 处理 变量
		switch v.Val {
		case "if":
			htmlText.WriteString("<!--if" + self.marker + "-->")
			parser.Next() // 跳过当前“if”
			parser.SkipSpace()
			part = self.parseExprBody(parser, CHILD_IF_PART)
			parser.Next()
			self.parseHtml(part, parser, false)
			*parts = append(*parts, part)

		case "else":
			//if parent != nil && parent.Type == CHILD_IF_PART && len(*parts) == 0 {
			ln := len(*parts)
			if ln == 0 || (*parts)[ln-1].Type != CHILD_IF_PART {
				parser.Backup(3)
				return true, nil
			}
			//}

			htmlText.WriteString("<!--else" + self.marker + "-->")
			parser.Next() // 跳过当前“else”
			parser.SkipSpace()
			part = self.parseExprBody(parser, CHILD_ELSE_PART)
			if len(part.Name) == 0 && (*parts)[ln-1].Type == CHILD_IF_PART {
				part.Name = (*parts)[ln-1].Name
				part.Strings = (*parts)[ln-1].Strings

			}

			parser.Next() // 跳过当前“}”
			self.parseHtml(part, parser, false)
			*parts = append(*parts, part)

		case "end":
			parser.Next()
			parser.Next()
			return true, nil

		default:
			htmlText.WriteString(self.marker)
		}
	}

	return false, nil
}

func (self *TemplateParser) parseElement(parser *Parser, parts *[]*TemplatePart, htmlText *strings.Builder) error {
	for !parser.IsEnd() {
		token := parser.Token()
		switch token.Type {
		case lexer.IDENT:
			// Tag Name
			htmlText.WriteString(parser.Token().Val)

		// 结束标签 />
		case lexer.QUO:
			htmlText.WriteByte('/')
			parser.Next()
			if token.Type == lexer.GTR {
				htmlText.WriteByte('>')
				goto RETURN
			}

			continue

		// 解析元素里的内容，这是另一个 html 片段
		case lexer.GTR:
			htmlText.WriteByte('>')
			goto RETURN

			// {{符号+变量 参数...}
		case lexer.LBRACE:
			parser.Next()
			if token.Type == lexer.LBRACE {
				_, err := self.parseExpr(nil, parser, parts, htmlText, true)
				if err != nil {
					return err
				}

				goto NEXT
			}
			fallthrough
		default:
			htmlText.WriteString(token.Val)
		}

	NEXT:
		parser.Next()
	}

RETURN:
	return nil
}

func (self *TemplateParser) parseHtml(parent *TemplatePart, parser *Parser, inElement bool) error {
	var htmlText strings.Builder
	var parts []*TemplatePart

	for !parser.IsEnd() {
		item := parser.Token()
		switch item.Type {
		case lexer.LSS:
			if err := self.parseElement(parser, &parts, &htmlText); err != nil {
				return err
			}

		case lexer.LBRACE:
			parser.Next()
			if parser.Token().Type == lexer.LBRACE { // {{符号
				mustReturn, err := self.parseExpr(parent, parser, &parts, &htmlText, inElement)
				if err != nil {
					return err
				}

				if mustReturn {
					goto RETURN
				}
			}

		case lexer.UNKNOWN, lexer.SAPCE:
			htmlText.WriteString(" ")

		default:
			htmlText.WriteString(item.Val)
		}

		parser.Next()
	}

RETURN:
	if htmlText.Len() > 0 {
		if parent != nil {
			parent.Value = htmlText.String()
			parent.Children = parts
		} else {
			self.parts = parts
			self.html.WriteString(htmlText.String())
		}
	}

	return nil
}

func (self *TemplateParser) Id() string {
	return self.id
}

func (self *TemplateParser) ToBytes() []byte {
	return []byte(self.html.String())
}

func (self *TemplateParser) ToString() string {
	return self.html.String()
}
