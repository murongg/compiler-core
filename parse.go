package main

import (
	"errors"
	"regexp"
	"strings"

	"github.com/gookit/goutil/strutil"
)

type Parse struct{}

type ParseContext struct {
	Source string
}

type NodeContent struct {
	Type    string
	content string
}

type Node struct {
	Type     string
	Tag      string
	TagType  string
	Content  NodeContent
	Children []Node
	Helpers  map[string]interface{}
}

func (parse *Parse) BaseParse(content string) (Node, error) {
	context := parse.CreateParseContext(content)
	node, err := parse.ParseChildren(&context, []*Node{})
	return *parse.CreateRoot(node), err
}

func (parse *Parse) CreateRoot(children []Node) *Node {
	return &Node{
		Type:     NodeTypesRoot,
		Children: children,
	}
}

func (parse *Parse) CreateParseContext(content string) ParseContext {
	return ParseContext{
		Source: content,
	}
}

func (parse *Parse) ParseChildren(context *ParseContext, ancestors []*Node) ([]Node, error) {
	var nodes []Node
	reg, err := regexp.Compile(`/[a-z]/i`)
	if err != nil {
		return []Node{}, err
	}
	for !parse.IsEnd(context, ancestors) {
		var node Node
		s := context.Source
		strSlice := strutil.ToArray(s)

		if strings.HasPrefix(s, "{{") {
			node = parse.ParseInterpolation(context)
		} else if strSlice[0] == "<" {
			if strSlice[1] == "/" && reg.MatchString(strSlice[2]) {
				parse.ParseTag(context, TagTypeEnd)
				continue
			}
		} else if reg.MatchString(strSlice[1]) {
			node, err = parse.ParseElement(context, ancestors)
		}

		if node.Type != "" {
			node = parse.ParseText(context)
		}
		nodes = append(nodes, node)
	}
	return nodes, err
}

func (parse *Parse) IsEnd(context *ParseContext, ancestors []*Node) bool {
	s := context.Source
	if strings.HasPrefix(context.Source, "</") {
		for _, v := range ancestors {
			if parse.StartsWithEndTagOpen(s, v.Tag) {
				return true
			}
		}
	}
	return context.Source == ""
}

func (parse *Parse) StartsWithEndTagOpen(source string, tag string) bool {
	if strings.HasPrefix(source, "</") {
		return ToLowerCase(StringSlice(source, 2, 2+len([]byte(tag)))) == ToLowerCase(tag)
	}
	return false
}

func (parse *Parse) ParseInterpolation(context *ParseContext) Node {
	openDelimiter := "{{"
	closeDelimiter := "}}"
	closeIndex := strings.Index(context.Source, closeDelimiter)
	parse.advanceBy(context, 2)

	rawContentLength := closeIndex - strutil.Len(openDelimiter)
	rawContent := StringSlice(context.Source, 0, rawContentLength)

	preTrimContent := parse.ParseTextData(context, strutil.Len(rawContent))
	content := strutil.Trim(preTrimContent)
	parse.advanceBy(context, strutil.Len(closeDelimiter))

	return Node{
		Type: NodeTypesInterpolation,
		Content: NodeContent{
			Type:    NodeTypesSimpleExpression,
			content: content,
		},
	}
}

func (parse *Parse) ParseTag(context *ParseContext, tagType string) Node {
	reg := regexp.MustCompile(`/^<\/?([a-z][^\r\n\t\f />]*)/i`)
	match := reg.FindAllString(context.Source, -1)
	tag := match[1]
	parse.advanceBy(context, strutil.Len(match[0]))
	parse.advanceBy(context, 1)
	if tagType == TagTypeEnd {
		return Node{}
	}
	return Node{
		Tag:     tag,
		TagType: ElementTypesElement,
		Type:    NodeTypesElement,
	}
}

func (parse *Parse) ParseElement(context *ParseContext, ancestors []*Node) (Node, error) {
	element := parse.ParseTag(context, TagTypeStart)

	ancestors = append(ancestors, &element)
	children, err := parse.ParseChildren(context, ancestors)
	ancestors = ancestors[:len(ancestors)-1]

	if parse.StartsWithEndTagOpen(context.Source, element.Tag) {
		parse.ParseTag(context, TagTypeEnd)
	} else {
		err = errors.New(`Missing close tag: ` + element.Tag)
	}
	element.Children = children
	return element, err
}

func (parse *Parse) ParseText(context *ParseContext) Node {
	endTokens := []string{"<", "{{"}
	endIndex := strutil.Len(context.Source)

	for i := 0; i < len(endTokens); i++ {
		index := strings.Index(context.Source, endTokens[i])
		if index != -1 && endIndex > index {
			endIndex = index
		}
	}
	content := parse.ParseTextData(context, endIndex)
	return Node{
		Type: NodeTypesText,
		Content: NodeContent{
			content: content,
		},
	}
}

func (parse *Parse) advanceBy(context *ParseContext, numberOfCharacters int) {
	context.Source = StringSlice(context.Source, numberOfCharacters, strutil.Len(context.Source)-1)
}

func (parse *Parse) ParseTextData(context *ParseContext, length int) string {
	rawText := StringSlice(context.Source, 0, length)
	parse.advanceBy(context, length)
	return rawText
}
