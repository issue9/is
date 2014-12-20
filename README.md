validation [![Build Status](https://travis-ci.org/issue9/validation.svg?branch=master)](https://travis-ci.org/issue9/validation)
======

一个简单的验证管理器及一系列验证函数：
```go
v := validation.New()
// 判断是否为数值
v.Apply(validator.IsNumber("123"), "必须为数值", "num")

// 判断是否为ISBN序列号
v.Apply(validator.IsISBN("1-919876-03-0"), "必须为一个ISBN序列号", "isbn")
```

### 安装

```shell
go get github.com/issue9/validation
```


### 文档

[![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/issue9/validation)
[![GoDoc](https://godoc.org/github.com/issue9/validation?status.svg)](https://godoc.org/github.com/issue9/validation)


### 版权

[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/issue9/validation/blob/master/LICENSE)
