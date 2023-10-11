## 变量声明

**不能放在单独的一行**
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
#### go语言不接受隐式类型转换
字符串转整形：

num, _ = strconv.Atoi(str)

整形转字符串：

str := strconv.Itoa(num)

### 变量
变量声明用空格隔开

var identifier type

变量声明也可以“分组”成一个语法块。

var identifier1, identifier2 type

指定变量类型，**如果没有初始化，则变量默认为零值**

默认值   0 false “”

:=赋值操作符 
**只能被用在函数体内**

### 常量

const identifier [type] = value

Go 在不同类型的项之间赋值时需要**显式转换**

常量不能用 := 语法声明  用**const**关键字

**iota**  const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )

### 格式化字符串

字符串形式，格式化符号以 % 开头， **%s** 字符串格式，**%d** 十进制的整数格式 **%v** 按值的本来值输出
**%#v**	输出 Go 语言语法格式的值 **%T**	输出 Go 语言语法格式的类型和值

### 循环
Go中只有一种循环：**for循环**

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

**goto** 转移到过程中的行

### 条件

if 语句可以在条件表达式前执行一个简单的语句。
eg：if v := math.Pow(x, n); v < lim {
		return v
	}
该语句声明的变量作用域仅在 if 之内 

**switch语句**

**select语句**

### 函数

func function_name( [parameter list] ) [return_types] {
   函数体
}
defer会将函数推迟到外层函数返回之后执行
推迟的函数调用会被压入一个**栈**中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用
eg：for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}//从9到0输出
#### 参数
值传递
引用传递（地址），会影响实际参数
#### 函数作为实参
func main(){
   /* 声明函数变量 */
   getSquareRoot := func(x float64) float64 {
      return math.Sqrt(x)
   }

   /* 使用函数 */
   fmt.Println(getSquareRoot(9))

}
#### 闭包
匿名函数是一种没有函数名的函数，通常用于在函数内部定义函数，或者作为函数参数进行传递。
func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}

multiply := func(x, y int) int {
        return x * y
    }

calculate := func(operation func(int, int) int, x, y int) int {
        return operation(x, y)
    }
sum := calculate(multiply, 2, 8)
### 方法
方法一个**包含了接受者的函数**，接受者可以是命名类型或者结构体类型的一个值或者是一个指针

该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}

**没有类。不过可以为结构体类型定义方法**
*方法只是个带接收者参数的函数*

**指针接受者**
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
会修改接受者指向的值
**函数不能接受指针，但是方法可以**

### 指针
var var_name *var-type

*T 是指向 T类型的指针。其零值为nil
var *p int
i := 5
p = &i

#### 指向指针的指针
var ptr **int
#### 指针作为函数参数
func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址的值 */
   *x = *y      /* 将 y 赋值给 x */
   *y = temp    /* 将 temp 赋值给 y */
}
实参会发生改变


### 数组
var arrayName [size]dataType
var a [10]int

var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
数量不确定  由...来确定

//  将索引为 1 和 3 的元素初始化
balance := [5]float32{1:2.0,3:7.0}
**多维数组**
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type

### 切片Slice
切片可以理解为“动态数组”，
var identifier []type//通过未指定大小的数组来定义切片
使用**make()函数**创建切片
var slice1 []type = make([]type, len)
slice1 := make([]type, len)

make([]T, length, capacity)// capacity 容量为可选参数。
**增加切片的容量**，**必须创建一个新的更大的切片**并把原分片的内容都拷贝过来


[]T  a[low:high](同python)
但是更改切片会修改底层数组
零值切片的值为 nil


### Range范围

range在for中可以对数组、切片、字符串等进行迭代循环
for key, value := range oldMap{

}

### Map集合
map
var m map[type]type 键key 和值value
或者map_variable := make(map[KeyType]ValueType, initialCapacity)
m := make(map[string]int)
v2, ok := m["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值

### 递归函数

func recursion(){
	recursion()
}
如斐波那契数列


### 结构体

type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}

struct是一组字段
eg：type Books struct {
   title string
   author string
   subject string
   book_id int
}
var Book1 Books
访问 Book1.title

可以指针访问结构体



### 接口

**接口类型** 是由一组方法签名定义的集合。

type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}
type struct_name struct {
   /* variables */
}
func (struct_name_variable struct_name) method_name1() [return_type] {
   /* 方法实现 */
}

**隐式实现** 
*接口也是值*

空接口
var i interface{}
空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

### 类型断言 

提供了访问接口值底层具体值的方式 t := i.(T)

t, ok := i.(T)
若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。

**类型选择**

类似switch

### Stringer

### error错误处理

error类型是一个接口类型
type error interface {
    Error() string
}


### io.Reader

### 并发
go关键字开启 goroutine
goroutine 是轻量级线程
#### 通道
用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯
通道在使用前必须先创建，使用chan关键字
ch := make(chan int, x)
tips：x可以定义缓冲区大小，异步
通过range关键字来遍历数据
