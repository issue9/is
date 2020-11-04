is
[![Go](https://github.com/issue9/is/workflows/Go/badge.svg)](https://github.com/issue9/is/actions?query=workflow%3AGo)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/issue9/is/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/is)
======

**不再更新，相关函数可采用 <github.com/issue9/validation/is> 下的内容**

一些常用的验证函数：

```go
// 判断是否为数值
is.Number("123")

// 判断是否为ISBN序列号
is.ISBN("1-919876-03-0")

// 身份证验证
is.GB11643("33232333211233432")

// 银行卡验证
is.BankCard("33232123234")
```

安装
----

```shell
go get github.com/issue9/is
```

文档
----

[![Go Walker](https://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/issue9/is)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/issue9/is)

版权
----

本项目采用 [MIT](https://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
