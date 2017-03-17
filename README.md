# 语法

注意：

1. 实际语法中不能有空格（空格会当作数据的一部分），以下所出现的空格仅为方便阅读
2. 文档中采用 # 符号后接注释内容，但实际中不能使用，原因同上

```
{Key} LF
{Key} : {Value} LF
{Key,MetaTag} : {Value} LF
{Key,MetaTag;MetaParam} : {Value} LF
...
LF
```

解释：

1. HBCP 为 Key-value 制协议，每一对 Key 与 Value 称作 Row，设计上便于人类阅读
2. 没有转义机制
3. 如只有 Key，Value 默认为空字符串
4. Key 不区分大小写
5. Key 只能由英文与数字组成
6. 在一次报文中，同一个 Key 若出现多次，最后只会保留该 Key 的最后一个 Row
7. Meta 影响 Value 的处理方式
8. 所有实现都应该支持的 MetaTag：
   * b64：Value 为一个 Base64 编码的字符串，读出时进行解码。没有 param
9. 若出现没有响应处理方法的 Meta，实现中应舍弃该 Row
10. 报文结尾有额外一个 LF 代表报文结束
11. LF 即换行符 \n

示例：

```
date:2006-01-02 15:04:05 CST -070
wxdao
binary,b64:SGVsbG8sIFdvcmxkIQ==
greeting,prefix;Hello! :How is it going?

```

# 使用示例

```go
package main

import (
	"fmt"
	"github.com/wxdao/hbcp"
)

func main() {
	office := hbcp.NewOffice(hbcp.Handler{
		OnMsg: func (context *hbcp.Context, msg hbcp.Msg) {
			fmt.Println("MSG: ", msg)
			context.Respond(msg)
		},
		OnJoin: func (context *hbcp.Context) {
			fmt.Println("JOIN: ", fmt.Sprint(context.RemoteAddr()))
		},
		OnClose: func (context *hbcp.Context) {
			fmt.Println("CLOSE: ", fmt.Sprint(context.RemoteAddr()))
		},
	}, map[string]hbcp.MetaHandler{
		"prefix": func (param string, value string) ([]byte, error) {
			return []byte(param + value), nil
		},
	})
	fmt.Println(office.Serve(":1234"))
}
```
