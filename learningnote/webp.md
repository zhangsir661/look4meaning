## 渗透？
### goglehacking
### 基本搜索
逻辑与: and
逻辑或: or
逻辑非:
完整匹配:“关键词”
通配符:*?
### 高级搜索
intext:正文包括
intitle:标题包括
allintitle:
inurl:url中包含
allinurl:
site:指定站点
info:返回网站信息
### 工具
在线工具类：fofa、shodan、whois、域名端口扫描、CMS指纹识别等
其他常见工具：御剑、Layer、K8PortScan、PScan、Nmap、Whatweb、Dirsearch、Dirb等等
### 注入
#### sql
单引号'，and 1=1测试
**盲注**：length()字符串长度
substr(string,start,num)：截取string的值，从第start个字符开始，每次返回num个
ascii()：查询ascii码中对应的值
count()：统计记录的数量
if()：条件判断函数
**时间**：if(condition,a,b):如果condition为true，返回a；否则返回b
sleep()函数，延迟响应时间（秒）
benchmark()函数，调试函数，用来测试某些特定操作的执行速度

#### 命令
system()、exec()、shell_exec()、passthru()等

**拼接符**A;B先执行A，再执行B；
A&B简单拼接，AB之间无制约关系；
A|B显示B的执行结果；
A&&BA执行成功，然后才会执行B；
A||BA执行失败，然后才会执行B

#### XSS

#### 反序列化

shiro通过默认秘钥构造cookie
Log4j2

#### 文件上传
0x00截断 %00

#### 逻辑

### WAF拦截
白名单+规则

### 网站木马


### 应急响应
#### 日志分析
Apache日志：
access_log访问日志 error_log错误日志
**192.168.16.211--[06/Jan/2022:17:03:29+0800]"GET/favicon.icoHTTP/1.1"404209**
192.168.16.211 -访问来源
\- 空白（E-Mmail）
\- 空白（登录名）
[]  请求时间
方法+资源+协议

Nginx日志
**127.0.0.1--[28/Oct/2016:09:03:56+0800]"GET/phpinfo.phpHTTP/1.1"20024914"-""Mozilla/5.0(WindowsNT6.1;WOW64;Trident/7.0;rv:11.0)likeGecko"**
客户端IP 客户端用户名称 时间 请求方式 请求URL地址（去掉host部分）/phpinfo.php 请求状态 请求页面大小（Byte） 用户浏览器其他信息

IIS日志分析
**2019-07-07 05:59:26 W3SVC1 127.0.0.1 GET /iisstart.htm - 80 -127.0.0.1 Mozilla/4.0(conpatible;+msie+6.0;+Windows+NT+5.2;+SV1;+.NET+CLR+1.1.4322) 200 0 0**
请求发生时间 客户端IP访问了服务端IP的那个端口 客户端工具类型，版本 请求URL以及查询字符串参数 请求方法 请求处理结果：HTTP状态码，操作系统底层的状态码 上传数据 发送数据 总共占用服务器时间

## 实际

### sql
判断类型
### xss
从代码分析
### 文件上传
判断拦截类型
黑名单？后缀？文件头？拦截规则？
### WAF


## APP渗透
