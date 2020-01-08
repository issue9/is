// SPDX-License-Identifier: MIT

// Package gb11643 解析身分证详情
package gb11643

import (
	"errors"
	"time"
)

// 我国现行的身份证号码有两种标准：GB11643-1989、GB11643-1999：
//
// GB11643-1989为一代身份证，从左至右分别为：
//  ------------------------------------------------------------
//  | 6位行政区域代码 | 6位出生年日期（不含世纪数）| 3位顺序码 |
//  ------------------------------------------------------------
//
// GB11643-1999为二代身份证，从左至右分别为：
//  ------------------------------------------------------------
//  | 6位行政区域代码 |  8位出生日期 |  3位顺序码 |  1位检验码 |
//  ------------------------------------------------------------

const layout = "20060102"

// GB11643 身份证信息
type GB11643 struct {
	Raw    string    // 原始数据
	Region string    // 区域代码
	Date   time.Time // 出生年月
	IsMale bool      // 是否为男姓，否为女生
}

// Parse 分析身份证信息
//
// 不作正确性检测，如有需求，请使用 is.GB11643
func Parse(bs string) (*GB11643, error) {
	switch len(bs) {
	case 15:
		return parse15(bs)
	case 18:
		return parse18(bs)
	default:
		return nil, errors.New("长度错误")
	}
}

func parse15(bs string) (*GB11643, error) {
	date, err := time.Parse(layout, "19"+bs[6:12])
	if err != nil {
		return nil, err
	}

	return &GB11643{
		Raw:    string(bs),
		Region: bs[:6],
		Date:   date,
		IsMale: (bs[14]-'0')%2 == 1,
	}, nil
}

func parse18(bs string) (*GB11643, error) {
	date, err := time.Parse(layout, bs[6:14])
	if err != nil {
		return nil, err
	}

	return &GB11643{
		Raw:    string(bs),
		Region: bs[:6],
		Date:   date,
		IsMale: (bs[16]-'0')%2 == 1,
	}, nil
}
