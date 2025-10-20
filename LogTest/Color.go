package LogTest

import "fmt"

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见

const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func (*Log) Black(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextBlack, msg, 0x1B)
}

func (*Log) Red(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextRed, msg, 0x1B)
}

func (*Log) Green(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextGreen, msg, 0x1B)
}

func (*Log) Yellow(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextYellow, msg, 0x1B)
}

func (*Log) Blue(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextBlue, msg, 0x1B)
}

func (*Log) Magenta(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextMagenta, msg, 0x1B)
}

func (*Log) Cyan(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextCyan, msg, 0x1B)
}

func (*Log) White(msg string) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 0, 0, TextWhite, msg, 0x1B)
}

func (*Log) SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
