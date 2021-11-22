#cat test.log | tr ' ' '\n' | sort | uniq -c|sort -r |awk '{print $2" "$1 }'
cat test.log | xargs -n1 echo | sort | uniq -c | sort -r | awk '{print $2" "$1}'

#solution 1 cannot trim head spaces

#question refer: https://leetcode-cn.com/problems/word-frequency/submissions/

#solution refer : https://blog.csdn.net/qq_29232943/article/details/116099401