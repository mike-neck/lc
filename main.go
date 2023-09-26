package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/stoewer/go-strcase"
	"os"
	"strings"
	"time"
)

func main() {
	option := LcOption{CamelCase}
	flag.Var(&option, "case", "output letter case, default 'camel', available values are 'camel'(=lower camel case) 'ucamel'(=upper camel case) 'kebab' 'ukebab' 'snake' 'usnake'")
	flag.Parse()
	ch := runReader()
	s := option.RunWriter(ch)
	os.Exit(s)
}

type Case int

const (
	CamelCase Case = iota + 1
	UpperCamelCase
	KebabCase
	UpperKebabCase
	SnakeCase
	UpperSnakeCase
)

func (c *Case) String() string {
	switch *c {
	case CamelCase:
		return "CamelCase"
	case UpperCamelCase:
		return "UpperCamelCase"
	case KebabCase:
		return "KebabCase"
	case UpperKebabCase:
		return "UpperKebabCase"
	case SnakeCase:
		return "SnakeCase"
	case UpperSnakeCase:
		return "UpperSnakeCase"
	}
	return fmt.Sprintf("<UNKNOWN CASE=%d>", *c)
}

func (c *Case) Convert(str string) (string, error) {
	switch *c {
	case CamelCase:
		return strcase.LowerCamelCase(str), nil
	case UpperCamelCase:
		return strcase.UpperCamelCase(str), nil
	case KebabCase:
		return strcase.KebabCase(str), nil
	case UpperKebabCase:
		return strcase.UpperKebabCase(str), nil
	case SnakeCase:
		return strcase.SnakeCase(str), nil
	case UpperSnakeCase:
		return strcase.UpperSnakeCase(str), nil
	}
	return "", fmt.Errorf("unknown format %s", c.String())
}

type LcOption struct {
	Case
}

func (o *LcOption) String() string {
	if o == nil {
		return "NIL"
	}
	return o.Case.String()
}

func (o *LcOption) Set(s string) error {
	l := strings.ToLower(s)
	switch l {
	case "camel":
		o.Case = CamelCase
		return nil
	case "ucamel":
		o.Case = UpperCamelCase
		return nil
	case "kebab":
		o.Case = KebabCase
		return nil
	case "ukebab":
		o.Case = UpperKebabCase
		return nil
	case "snake":
		o.Case = SnakeCase
		return nil
	case "usnake":
		o.Case = UpperSnakeCase
		return nil
	default:
		return fmt.Errorf("unknown case[%s]", s)
	}
}

func (o LcOption) RunWriter(ch <-chan string) int {
	t := time.NewTimer(1 * time.Second)
	for {
		select {
		case <-t.C:
			return 3
		case input, hasMore := <-ch:
			output, err := o.Convert(input)
			if err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
				return 2
			} else if hasMore {
				fmt.Println(output)
			} else {
				fmt.Println(output)
				return 0
			}
		}
	}
}

func runReader() <-chan string {
	ch := make(chan string)
	go func() {
		defer func() { close(ch) }()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()
	return ch
}
