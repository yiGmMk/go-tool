# go-tool

## package

### convert

- Number2ChineseYUAN 数字到中文大写金额转换,精度仅支持到分

## serverless

netlify 和 vercel都支持部署到aws lambda,参考样例 api 和 functions

### netlify

项目根目录下放配置文件 netlify.toml,函数位置可配置

#### 限制

- 1.仅支持js/ts/go运行时(2024.9.24),python不支持

### vercel

项目根目录下放配置文件 vercel.json,接口放api目录下

#### 特性

- 1. 支持js/ts/go/python运行时,[python样例,fastapi](https://github.com/yiGmMk/free-unoficial-gpt4o-mini-api)

#### 限制

