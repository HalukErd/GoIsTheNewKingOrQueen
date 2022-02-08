This method solves mixed math operations with O(n)

For example:
```text
"1+22*(3+1)(35-(33*2)+33)-2*(3)"
```

Firstly edits
* puts * for hidden * operation, it becomes like that  )( ---> )*(
* puts given str in brackets, it becomes like that str ---> (str)
* and by changing - to +- removes - operation

After transformation it is like that:
```text
(1+22*(3+1)*(35+-(33*2)+33)+-2*(3))
```

Result:
```text
171
```
