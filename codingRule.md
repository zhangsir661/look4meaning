常见的编码

散列hash加密

一、MD5

test@123456
32位小写
39c3f82b0f479fa61bf46f78e8b4df35 

32位大写
39C3F82B0F479FA61BF46F78E8B4DF35 

16位小写
0f479fa61bf46f78 

16位大写
0F479FA61BF46F78 

一般MD5值是32位由数字“0-9”和字母“a-f”所组成的字符串。
MD5的三个特征：1.确定性 2.碰撞性（即多对一的映射关系）3.不可逆

二、sha1

特征与MD5差不多，但是是40位

三、HMAC

示例：test@123456 秘钥123 7cc2437b404060fedfc68f11e8b27b85
即在上述过程中加入了秘钥

四、NTLM

windows的哈希密码，32位
test@123456 47c61a0fa8738ba77308a8a600f88e4b

Base系列编码

base系列编码的不同，在于用于编码的字符数量的多少

一、Base64

码表 ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ 

另有“=”作为补充

base64和URL传参问题：

标准的Base64并不适合直接放在URL里传输，因为URL编码器会把标准Base64中的“/”和“+”字符变为形如“%XX”的形式，
用于URL的改进Base64编码，它在末尾填充’='号，并将标准Base64中的“+”和“/”分别改成了“-”和“_”（还有一些其他形式）

二、Base58

Base58不使用数字"0"，字母大写"O"，字母大写"I"，和字母小写"l"，以及"+“和”/"符号
字母表：123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz

三、Base32

Base32使用了ASCII编码中可打印的32个字符：大写字母A-Z和数字2-7
由于传输数据的单位是字节(即8个二进制位).所以分割之前的二进制位数是40的倍数(40是5和8的最小公倍数).如果不足40位，则在编码后数据补充"="，一个"="相当于一个组(5个二进制位)，编码后的数据是原先的8/5倍.
明文超过10个后就会有很多等号

四、Base16（HEX，即16进制编码）

将数据(根据ASCII编码，UTF-8编码等)转成对应的二进制数，不足8比特位高位补0。然后将所有的二进制全部串起来，4个二进制位为一组，转化成对应十进制数。
0-9 A-F

编码

一、Unicode

汉字示例&#36825;、字母示例&#116;、数字符号示例&#53

二、16进制Unicode

\u8fd9\u662f\u4e00

三、URL编码

编码规则：对 String 编码时，使用以下规则：

字母、数字和字符， “a” 到 “z”、”A” 到 “Z” 和 “0” 到 “9” 保持不变；

特殊字符 “.”、”-“、”*” 和 “_” 保持不变；

空格字符 ” ” 转换为一个加号 “+”；

url编码时会将 + 编码成空格；

除此之外，所有的其他字符都是不安全的，会变成以 % 开头的字符；

对于一个中文，会使用3个字节表示。

encodeURIComponent() 函数 与 encodeURI() 函数的区别
请注意 encodeURIComponent() 函数 与 encodeURI() 函数的区别之处，前者假定它的参数是 URI 的一部分（比如协议、主机名、路径或查询字符串）。
因此 encodeURIComponent() 函数将转义用于分隔 URI 各个部分的标点符号

补充：

\x开头编码

```
\xe4\xbd\xa0\xe5\xa5\xbd\xe4\xb8\x96\xe7\x95\x8c
中文是： 你好世界
```

utf-8编码，但数据类型是字符串类型

只要将字符串中的 " \x " 改为 " % " 利用urllib中的unquote方法解码就可以得到中文了，因为url中的中文utf-8编码和这里的区别就是url中编码是%开头

```python
s = '\xe4\xbd\xa0\xe5\xa5\xbd\xe4\xb8\x96\xe7\x95\x8c'
s = s.encode('unicode_escape')
ss = s.decode('utf-8').replace('\\x', '%')
un = parse.unquote(ss)
```

