package convert

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	simpleDigits      = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
	traditionalDigits = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}
	simpleUnits       = []string{"", "十", "百", "千"}
	traditionalUnits  = []string{"", "拾", "佰", "仟"}

	max = decimal.NewFromFloat(9999_9999_9999_9999.99)
	min = decimal.NewFromFloat(-9999_9999_9999_9999.99)
)

/*Number2ChineseYUAN
 * @description: 数字转换成大写中文数字金额
 * @param {decimal.Decimal} val 需要转换的数字
 * @param {*} isTraditional 是否繁体
 * @return {*}
 * 注意浮点数精度问题: 单精度浮点数float32的精度为6~8位，双精度浮点数float64的精度为15~17位
 * 输入的val若使用decimal.NewFromFloat()初始化且长度大于等于16位将会丢失精度,这是因为赋值后进入函数转换前精度已丢失
 * 如2345_9999_9999_1232.23(18位),请使用字符串初始化decimal.NewFromString("2345999999991232.23")
 */
func Number2ChineseYUAN(val decimal.Decimal, isTraditional bool) (string, error) {
	if val.IsZero() {
		return "零元整", nil
	}

	if (val.IsPositive() && val.GreaterThan(max)) || (val.IsNegative() && val.LessThan(min)) {
		return "", errors.Errorf("数字超出范围,不支持转换,最大值:%v,最小值:%v", max, min)
	}

	negative := val.IsNegative()
	if negative {
		val = val.Abs()
	}

	amount := val.Mul(decimal.NewFromInt(100)).Round(2).IntPart()
	分 := amount % 10
	amount /= 10
	角 := amount % 10
	amount /= 10

	//将数字以万为单位分为多份
	parts := make([]int64, 10)
	numParts := 0
	for i := 0; amount != 0; i++ {
		part := amount % 10000
		parts[i] = part
		numParts++
		amount = amount / 10000
	}

	res := ""
	nextIsZero := true // '万'下一级是否为0
	for i := 0; i < numParts; i++ {
		part := toChinese(parts[i], isTraditional)
		if i%2 == 0 {
			nextIsZero = part == ""
		}
		if i == 0 { //万以内,直接加
			goto addEnd
		}
		if i%2 == 0 {
			part += "亿"
			goto addEnd
		}
		// 如果'万'对应的part为0,而'万'下面一级不为0,则不加'万',而加“零'
		if part == "" && !nextIsZero {
			part += "零"
			goto addEnd
		}
		part += "万"
		// 如果'万'的部分不为 0,而"万"前面的部分小于1000大于0,则万后面应该跟“零'
		if parts[i-1] > 0 && parts[i-1] < 1000 {
			part += "零"
		}
	addEnd:
		res = part + res
	}

	nums := traditionalDigits
	if !isTraditional {
		nums = simpleDigits
	}

	// 整数部分为 0, 则表达为"零"
	if res == "" {
		res = nums[0]
	}
	if negative {
		res = "负" + res
	}
	res = formatFractionalPart(res, 角, 分, nums)
	return res, nil
}

/**
 * @description: 金额小数部分,仅保留转换至分
 * @param {string} intPart 元整数部分
 * @param {*} tenths 角
 * @param {int64} percentile 分
 * @param {[]string} nums 中文数字
 * @return {*} 转换结果:中文大写金额
 */
func formatFractionalPart(intPart string, tenths, percentile int64, nums []string) string {
	res := intPart
	角 := tenths
	分 := percentile
	if 分 == 0 && 角 == 0 {
		res += "元整"
		return res
	}
	if 分 == 0 {
		res += "元" + nums[角] + "角整"
		return res
	}
	if 角 == 0 {
		res += "元零" + nums[分] + "分"
		return res
	}
	res += "元" + nums[角] + "角" + nums[分] + "分"
	return res
}

/**
 * @description: 0~9999之间的数字转换为中文数字,为0返回空字符串
 * @param {int} amount 需要转换的数字
 * @param {bool} isTraditional 是否是繁体
 * @return {*} 转换后的中文数字
 */
func toChinese(amount int64, isTraditional bool) string {
	res := ""
	if amount == 0 {
		return res
	}
	nums := traditionalDigits
	units := traditionalUnits
	if !isTraditional {
		nums = simpleDigits
		units = simpleUnits
	}

	lastIsZero := true // 低位往高位循环,记录上一位是否为0(百位的上一位指十位)

	for i, temp := 0, amount; temp > 0; i++ {
		digit := temp % 10
		temp /= 10
		// 当前为0上一位不为零要加上'零'
		if digit == 0 && !lastIsZero {
			lastIsZero = true
			res = "零" + res
			continue
		}
		// 不为零,直接加上当前位
		res = nums[digit] + units[i] + res
		lastIsZero = false
	}
	return res
}
