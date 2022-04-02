package main

import (
	"fmt"
)

// 参考 fs.FileMode的用法
// 从而理解golang中常量和iota的使用方式
func main() {
	var c = Mode(Chinese, TraditionalChinese)
	fmt.Println(c.Support(Chinese))
	fmt.Println(c.Support(TraditionalChinese))
	fmt.Println(c.Support(English))
	fmt.Println(c.Support(Japannes))
}

// 测试使用type别名
const (
	Chinese Language = 1 << iota
	TraditionalChinese
	English
	Japannes
)

type Language uint64

type LanguageMode uint64

// SupportedLanguages 支持的语言
func SupportedLanguages() []string {
	return []string{
		"Chinese",
		"TraditionalChinese",
		"English",
		"Japannes",
	}
}

func (m LanguageMode) Support(language Language) bool {
	return (uint64(m) & uint64(language)) > 0
}

func (m LanguageMode) SupportChinese() bool {
	return false
}

func Mode(ls ...Language) LanguageMode {
	var rtn LanguageMode
	// for range 会帮我们判断ls是否为nil

	for _, l := range ls {
		rtn = rtn | LanguageMode(l)
	}

	return rtn
}
