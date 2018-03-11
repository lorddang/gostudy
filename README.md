# day 2

## 字符串

### 字符串的格式化
### 字符串的切片
### 字符串的翻转

6

### 包 strings

```
fmt.Println("Hello World")
```


# day 3

## strings

### strings.HasPrefix(s string, prefix string) bool
作用:判断字符串s是否以prefix开头
练习：判断一个url是否以http:// 开头，如果不是,则加上http://.
### strings.HasSuffix(s string, suffix string) bool:
作用:判断字符串s是否以suffix结尾
练习2：判断一个路径是否以/结尾,如果不是加上/。

### strings.Index(s string, str string) int

作用:判断str 在s中首次出现的位置如果没有出现则返回-1
4. strings.LastIndex(s string, str string) int：判断str在s中最后出现的位置，如果没有出现，则返回-1

练习3：写一个函数返回一个字符串在另一个字符串的首次出现和最后出现位置
func StrIndex(str string, substr string)(int, int){}

5. strings.Replace(str string, old string, new string, n int)：字符串替换

6. strings.Count(str string, substr string)int：字符串计数

7. strings.Repeat(str string, count int)string：重复count次str

8. strings.ToLower(str string)string：转为小写

9. strings.ToUpper(str string)string：转为大写

