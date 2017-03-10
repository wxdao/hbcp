# 语法

注意：实际语法中不能有空格（空格会当作数据的一部分），以下所出现的空格仅为方便阅读

```
{Key} LF
{Key} : {Value} LF
...
LF
```

解释：

1. HBCP 为 Key-value 制协议，每一对 Key 与 Value 称作 Row，设计上便于人类阅读
2. 也可以没有 Value，Key 占用一行作为一个数据
3. Key 不区分大小写
4. Key 只能由英文与数字组成
5. 在一次报文中，同一个 Key 只能出现一次
6. 对于 Value 为简单数据（无不可见字符，无换行符），LF 作为一条 Row 的结束标志
7. 对于 Value 为二进制数据等非简单数据，应转换为 base64 编码
8. 报文结尾有额外一个 LF 代表报文结束
9. LF 即换行符 \n

示例：

```
date:2006-01-02 15:04:05 CST -070
binary:InFactImBase64Encoded
author:wxdao


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
		HandleMsgFunc: func (context *hbcp.Context, msg hbcp.Msg) {
			fmt.Println("MSG: ", msg)
			context.Respond(hbcp.Msg{"request": fmt.Sprint(msg)})
		},
		HandleJoinFunc: func (context *hbcp.Context) {
			fmt.Println("JOIN: ", fmt.Sprint(context.RemoteAddr()))
		},
		HandleCloseFunc: func (context *hbcp.Context) {
			fmt.Println("CLOSE: ", fmt.Sprint(context.RemoteAddr()))
		},
	})
	fmt.Println(office.Serve(":1234"))
}
```
