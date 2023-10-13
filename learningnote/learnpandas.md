## Pandas笔记
Pandas主要有两种数据结构 Series和DataFrame
### Series
Series类似于表格中的一个列column，类似于一维数组

pandas.Series(data, inedex, dtype, name, copy)

data:数据 index:索引标签，默认0开始 dtype:数据类型 name:设置名称 copy:拷贝数据

### DataFrame

DataFrame是一个表格型的数据结构，一组有序的列

pandas.DataFrame(data, index, columns, dtype, copy)

columns:列标签 

### CSV
清洗空值

DataFrame.dropna(axis=0, how='any', thresh=None, subset=None, inplace=False)

inplace为ture是会修改源数据DataFrame

fillna 替换空字段


### JSON
json模块

glom模块


