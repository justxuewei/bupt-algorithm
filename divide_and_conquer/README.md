# 分治法

## 原理

该算法为线性时间查找，用于在无序数组中找第k个数字，该算法使用的解决思路是分治法，通过找中位数来划分更小的数据集以降低时间复杂度。

## 问题

无序数组的（确定的）线性时间查找算法的实现，要求查找过程中修改所给序列，研究并分析在进行大量的查找过程中同一关键字的查找次数变化情况，注意所给列逐步有序的过程。要求随机生成10000个数据，随机取200个数做20次查询，随机取其他1000个数做5次以下查询（合计2000次）。给出20次查询的数字后的相关统计结果，记录最长递增子列长度的变化。

## 算法

该算法的步骤为：

- 将数组分为n/5(向上取整)个组
- 通过插入排序找到每个组的中位数
- 递归查找中位数的中位数，直到最后只有一个中位数，记为pivot
- 与快速排序类似，通过pivot将当前数组划分为三个部分：比pivot小的、和pivot相等的以及比pivot大的
- 根据情况舍弃上面的一部分，在进行下一步的递归

时间复杂度为O(N)，在实验中对200个随机数进行查询后，可以很明显的观察到数组的有有序性提高了，下面记录了最长有序数据的长度：

```
linear_time_selection_test.go:43: maximum continuous length == 10
    linear_time_selection_test.go:43: maximum continuous length == 10
    linear_time_selection_test.go:43: maximum continuous length == 15
    linear_time_selection_test.go:43: maximum continuous length == 15
    linear_time_selection_test.go:43: maximum continuous length == 22
    linear_time_selection_test.go:43: maximum continuous length == 25
    linear_time_selection_test.go:43: maximum continuous length == 25
    linear_time_selection_test.go:43: maximum continuous length == 35
    linear_time_selection_test.go:43: maximum continuous length == 35
    linear_time_selection_test.go:43: maximum continuous length == 36
    linear_time_selection_test.go:43: maximum continuous length == 36
    linear_time_selection_test.go:43: maximum continuous length == 36
    linear_time_selection_test.go:43: maximum continuous length == 36
    linear_time_selection_test.go:43: maximum continuous length == 36
    linear_time_selection_test.go:43: maximum continuous length == 36
    linear_time_selection_test.go:43: maximum continuous length == 45
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 70
    linear_time_selection_test.go:43: maximum continuous length == 71
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 72
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 80
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 81
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 85
    linear_time_selection_test.go:43: maximum continuous length == 90
    linear_time_selection_test.go:43: maximum continuous length == 90
    linear_time_selection_test.go:43: maximum continuous length == 92
    linear_time_selection_test.go:43: maximum continuous length == 93
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 94
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 109
    linear_time_selection_test.go:43: maximum continuous length == 110
    linear_time_selection_test.go:43: maximum continuous length == 110
    linear_time_selection_test.go:43: maximum continuous length == 113
    linear_time_selection_test.go:43: maximum continuous length == 117
    linear_time_selection_test.go:43: maximum continuous length == 121
    linear_time_selection_test.go:43: maximum continuous length == 123
    linear_time_selection_test.go:43: maximum continuous length == 125
    linear_time_selection_test.go:43: maximum continuous length == 125
    linear_time_selection_test.go:43: maximum continuous length == 129
    linear_time_selection_test.go:43: maximum continuous length == 133
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
    linear_time_selection_test.go:43: maximum continuous length == 137
```

## 总结

这个题是一个典型的分治法，把一个大问题划分为若干个小问题，同时通过中位数这个机制可以，很方便的忽略一部分数组，使得整个算法的时间复杂度保持在线性范围内。