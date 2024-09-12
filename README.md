# TimeFinder

[![Go Report Card](https://goreportcard.com/badge/github.com/threadshare/timefinder)](https://goreportcard.com/report/github.com/threadshare/timefinder)
[![GoDoc](https://godoc.org/github.com/threadshare/timefinder?status.svg)](https://godoc.org/github.com/threadshare/timefinder)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

TimeFinder is a Golang library for finding and extracting time information from Chinese text. It provides a simple way to parse text and identify time expressions.

时间信息提取 | 中文自然语言处理 | NLP | 日期解析 | 时间解析

## 简介

TimeFinder 是一个用于在中文文本中查找和提取时间信息的 Golang 库。它提供了一种简单的方式来解析文本并识别其中的时间表达。

主要特点：
- 基于 SeGo 进行分词
- 专门针对中文自然语言进行时间提取
- 快速准确的时间信息提取
- 支持多种日期和时间格式
- 提供灵活的时间范围识别功能
- 易于集成和使用

## 安装

确保您已安装 Golang，然后使用以下命令安装 TimeFinder：

```
$ go get github.com/threadshare/timefinder

```

## 使用

```
go get github.com/threadshare/timefinder
```

```
var msg string
var extract []time.Time

// 初始化timefinder
segmenter := timefinder.New()

msg = " 6月9日有一场show要去观看"
// 解析话语词汇
extract = segmenter.TimeExtract(msg)
fmt.Println(msg)
fmt.Println(extract[0].Format(timeFormat))
```
上述代码会在文本中查找时间信息，并将结果打印输出。您可以根据需要自定义输出格式和进一步处理提取的时间信息。

## 支持的时间格式

TimeFinder 支持多种日期和时间格式的识别和提取，包括但不限于：

- 年月日：2023年6月1日、2023-06-01、6/1/2023
- 时间：15:30、15点30分、下午3点30分
- 相对时间：明天、下周五、三天后
- 时间范围：6月1日至6月5日、上午9点到下午5点

## 更多示例

为了展示 TimeFinder 的更多用法和支持的时间格式，我们在 `tests/finder_test.go` 文件中提供了详细的测试用例。这些测试用例涵盖了各种常见的时间表达方式，包括：

- 具体日期和时间
- 相对时间表达（如"明天"、"后天"）
- 时间范围
- 模糊时间表达

我们强烈建议您查看 `tests/finder_test.go` 文件以了解更多使用场景和预期输出。这将帮助您更好地理解 TimeFinder 的功能和灵活性。

## 支持的语言

- 中文 (Chinese)

## 关键词

时间提取, 日期解析, 中文NLP, 自然语言处理, Golang, 时间识别, 文本分析

## 相关项目

- [SeGo](https://github.com/huichen/sego) - Go 中文分词

## 贡献和反馈

TimeFinder 是一个开源项目，欢迎您的贡献和反馈。如果您发现问题、有改进建议或者想要贡献代码，请参阅项目的贡献指南并提交 Issue 或 Pull Request。

## 许可证

TimeFinder 使用 [MIT 许可证](LICENSE)。
