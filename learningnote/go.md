## 变量声明

**{**不能放在单独的一行
var 声明

简洁赋值语句 := 可在类型明确的地方代替 var 声明
函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用

#### 类型
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128


int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。

### 变量
变量声明用空格隔开
变量声明也可以“分组”成一个语法块。
默认值
0 false “”

Go 在不同类型的项之间赋值时需要**显式转换**
常量不能用 := 语法声明  用**const**关键字

### 格式化字符串

字符串形式，格式化符号以 % 开头， **%s** 字符串格式，**%d** 十进制的整数格式 **%v** 按值的本来值输出
**%#v**	输出 Go 语言语法格式的值 **%T**	输出 Go 语言语法格式的类型和值

### 循环
Go中只有一种循环：for循环

基本的 for 循环由三部分组成，它们用分号隔开：

初始化语句：在第一次迭代前执行 **（可选**
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行 **（可选**
初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。
eg：for i := 0; i < 10; i++ {
		sum += i
	}

去掉分号 就变成了**while** 
eg：sum := 1
	for sum < 1000 {
		sum += sum
	}
再去掉条件 无限循环

if 语句可以在条件表达式前执行一个简单的语句。
eg：if v := math.Pow(x, n); v < lim {
		return v
	}
该语句声明的变量作用域仅在 if 之内 

switch语句

### 函数
defer会将函数推迟到外层函数返回之后执行
推迟的函数调用会被压入一个**栈**中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
eg：for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}//从9到0输出

### 指针

*T 是指向 T类型的指针。其零值为nil
var *p int
i := 5
p = &i

### 结构体

struct是一组字段
eg：type Vertex struct {
	X int
	Y int
    }
可以指针访问结构体

### 数组

var a [10]int
**切片**
[]T  a[low:high](同python)
但是更改切片会修改底层数组
零值切片的值为 nil

make函数 

### 映射
map
var m map[type]type 键key 和值value
m := make(map[string]int)
v2, ok := m["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值

### 函数

**闭包**

### 方法

**没有类。不过可以为结构体类型定义方法**
*方法只是个带接收者参数的函数*

**指针接受者**
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
会修改接受者指向的值
函数不能接受指针，但是方法可以

### 接口

**接口类型** 是由一组方法签名定义的集合。






