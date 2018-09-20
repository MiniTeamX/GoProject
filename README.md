# Go后台API接口
### 说明
- 采用Restful的API格式，http请求的Get、Post、Delete、Put方法分别表示获取资源，提交资源，删除资源，修改资源。
可以参考[理解RESTful架构](http://www.ruanyifeng.com/blog/2011/09/restful.html)
- 项目host：192.168.26.163  port:8088
### User API
```
- 【Get】host:port/user/(uid)  
  Get请求:http://192.168.26.193:8088/user/1
  返回结果：
  {
  "UserId": 1,
  "Username": "11111111111",
  "Password": "123",
  "NickName": "张三1",
  "Gender": 1,
  "PhotoUrl": "",
  "Introduction": "哈哈，我是张三1"
 }
```
