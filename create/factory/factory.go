package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (J jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// 简单工厂
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type yamlRuleConfigParserFactory struct {
}

func (y yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

// 工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "yaml":
		return yamlRuleConfigParserFactory{}
	}
	return nil
}
