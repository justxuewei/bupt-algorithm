# 搜索

## 原理

求解纵横字谜游戏需要通过回溯法来完成，但是在有多条带匹配的词语时，如果不借助特殊的数据结构，需要对每一个单词执行一次回溯法，那么时间复杂度就会很高了，所以需要建立一种前缀树来避免陷入过大的时间复杂度中。

回溯法实际上是一种特殊的DFS方法。它尝试分步的去解决一个问题。在分步解决问题的过程中，当它通过尝试发现，现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。

## 问题

设计算法求解纵横字谜（Crossword）游戏的搜索算法（求解时可以先给出一个小规模的字典），再对不同规模的问题进行求解，对算法的性能进行评价。

## 算法

首先是建立前缀树，它本质上来说是一个