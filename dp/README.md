# 动态规划

## 原理

动态规划问题的基本思想与分治法相似，讲一个大的问题分解为若干个小问题，通过顺序解答后后面的问题解答可以利用前面的计算结果，即后面的计算结果依赖前面的计算结果。相比于分治法，最大的差别在于划分的子过程是否依赖原有的计算。

## 问题

在一个有向无环图（directed acyclic graph）中，找到一个结点（源点）到另外一个终点之间的最短路径。要求给出20个城市，随机生成20个城市的之间的路径长度（无回路），找到源点到相应的结点间最短路径。

## 算法

这个问题可以使用动态规划的思路解决。wo men