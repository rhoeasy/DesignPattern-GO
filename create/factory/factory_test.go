package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParser{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParser{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactory
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParserFactory{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}

}
