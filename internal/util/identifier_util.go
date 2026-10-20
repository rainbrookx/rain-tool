package util

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

type Identifier struct {
	CamelCase      string `json:"camelCase"`      // 小驼峰
	PascalCase     string `json:"pascalCase"`     // 大驼峰
	SnakeCase      string `json:"snakeCase"`      // 下划线小写
	KebabCase      string `json:"kebabCase"`      // 中划线
	UpperSnake     string `json:"upperSnake"`     // 全大写下划线
	VarNameSpace   string `json:"varNameSpace"`   // 全小写空格
	TitleNameSpace string `json:"titleNameSpace"` // 每个单词首字母大写空格
}

type PrefixAndSuffix struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

/*
IdentifierConvert 输入其中任意一种，输出所有的命名类型

【标准标识符】
- camelCase（小驼峰）
- PascalCase（大驼峰）
- snake_case（下划线）
- kebab-case（中划线）
- UPPER_SNAKE（全大写下划线）

【空格间隔】
输入：不区分大小写
输出：
- variable name（全小写，空格间隔）
- Variable Name（首字母大写，空格间隔）
*/
func IdentifierConvert(originalText string) (*Identifier, error) {
	originalText = strings.TrimSpace(originalText)
	if originalText == "" {
		return nil, errors.New("originalText cannot be empty")
	}

	// 1. 先拆成单词列表（核心步骤）
	words, err := splitToWords(originalText)
	if err != nil {
		return nil, err
	}
	if len(words) == 0 {
		return nil, errors.New("parse identifier failed, no words extracted")
	}

	// 2. 统一转小写单词数组
	lowerWords := make([]string, len(words))
	for i, w := range words {
		lowerWords[i] = strings.ToLower(w)
	}

	// 3. 逐个生成格式
	id := &Identifier{}

	// variable name（全小写空格）
	id.VarNameSpace = strings.Join(lowerWords, " ")

	// Variable Name（每个单词首字母大写，空格）
	titleWords := make([]string, len(lowerWords))
	for i, w := range lowerWords {
		titleWords[i] = firstUpper(w)
	}
	id.TitleNameSpace = strings.Join(titleWords, " ")

	// snake_case
	id.SnakeCase = strings.Join(lowerWords, "_")
	// UPPER_SNAKE
	id.UpperSnake = strings.ToUpper(id.SnakeCase)
	// kebab-case
	id.KebabCase = strings.Join(lowerWords, "-")

	// camelCase：第一个小写，后面每个首字母大写
	camel := make([]string, len(lowerWords))
	camel[0] = lowerWords[0]
	for i := 1; i < len(lowerWords); i++ {
		camel[i] = firstUpper(lowerWords[i])
	}
	id.CamelCase = strings.Join(camel, "")

	// PascalCase：全部首字母大写
	pascal := make([]string, len(lowerWords))
	for i, w := range lowerWords {
		pascal[i] = firstUpper(w)
	}
	id.PascalCase = strings.Join(pascal, "")

	return id, nil
}

// splitToWords 自动识别各种命名风格，切割成单词
// 支持：驼峰、下划线、中划线、空格
func splitToWords(s string) ([]string, error) {
	// 先统一把 _ - 空格 替换成空格，再处理驼峰
	repl := regexp.MustCompile(`[-_\t\n\f\r ]+`)
	s = repl.ReplaceAllString(s, " ")

	var words []string
	var current strings.Builder

	for _, r := range s {
		if unicode.IsSpace(r) {
			if current.Len() > 0 {
				words = append(words, current.String())
				current.Reset()
			}
			continue
		}

		// 驼峰分割：遇到大写，且前面是小写字母，断开
		if unicode.IsUpper(r) {
			if current.Len() > 0 {
				lastRune := lastRune(current.String())
				if unicode.IsLower(lastRune) {
					words = append(words, current.String())
					current.Reset()
				}
			}
		}
		current.WriteRune(r)
	}
	if current.Len() > 0 {
		words = append(words, current.String())
	}

	// 去空 + 去重（顺序不变）
	var res []string
	seen := make(map[string]bool)
	for _, w := range words {
		w = strings.TrimSpace(w)
		if w != "" && !seen[w] {
			seen[w] = true
			res = append(res, w)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("no valid words")
	}
	return res, nil
}

func lastRune(s string) rune {
	runes := []rune(s)
	if len(runes) == 0 {
		return 0
	}
	return runes[len(runes)-1]
}

// firstUpper 首字母大写，其余小写
func firstUpper(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// IdentifierPrefixAndSuffix identifier 添加前缀或后缀
func IdentifierPrefixAndSuffix(originalText string, prefixAndSuffix []PrefixAndSuffix) []string {
	var list []string

	for _, ps := range prefixAndSuffix {
		list = append(list, ps.Prefix+originalText+ps.Suffix)
	}
	return list
}
