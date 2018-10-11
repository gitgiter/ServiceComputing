echo "test 1 start"
selpg -s=1 -e=1 selpg_input.txt
echo "test 1 complete"
# 1. 该命令将把“selpg_input.txt”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。

echo "test 2 start"
selpg -s=1 -e=1 < selpg_input.txt
echo "test 2 complete"
# 2. 该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“selpg_input.txt”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。

echo "test 3 start"
selpg -s=2 -e=5 selpg_input.txt | selpg -s=1 -e=2
echo "test 3 complete"
# 3. “selpg -s=2 -e=5”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 1 页到第 2 页写至 selpg 的标准输出（屏幕）。

echo "test 4 start"
selpg -s=10 -e=20 selpg_input.txt > selpg_output.txt
echo "test 4 complete"
# 4. selpg 将第 10 页到第 20 页写至标准输出；标准输出被 shell／内核重定向至“selpg_output.txt”。

echo "test 5 start"
selpg -s=10 -e=20 selpg_input.txt 2> selpg_error.txt
echo "test 5 complete"
# 5. selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“selpg_error.txt”。请注意：在“2”和“>”之间不能有空格；这是 shell 语法的一部分（请参阅“man bash”或“man sh”）。

echo "test 6 start"
selpg -s=10 -e=20 selpg_input.txt > selpg_output.txt 2> selpg_error.txt
echo "test 6 complete"
# 6. selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“selpg_output.txt”；selpg 写至标准错误的所有内容都被重定向至“selpg_error.txt”。当“selpg_input.txt”很大时可使用这种调用；您不会想坐在那里等着 selpg 完成工作，并且您希望对输出和错误都进行保存。

echo "test 7 start"
selpg -s=10 -e=20 selpg_input.txt > selpg_output.txt 2> /dev/null
echo "test 7 complete"
# 7. selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“selpg_output.txt”；selpg 写至标准错误的所有内容都被重定向至 /dev/null（空设备），这意味着错误消息被丢弃了。设备文件 /dev/null 废弃所有写至它的输出，当从该设备文件读取时，会立即返回 EOF。

echo "test 8 start"
selpg -s=10 -e=20 selpg_input.txt > /dev/null
echo "test 8 complete"
# 8. selpg 将第 10 页到第 20 页写至标准输出，标准输出被丢弃；错误消息在屏幕出现。这可作为测试 selpg 的用途，此时您也许只想（对一些测试情况）检查错误消息，而不想看到正常输出。

echo "test 9 start"
selpg -s=10 -e=20 selpg_input.txt | wc
echo "test 9 complete"
# 9. selpg 的标准输出透明地被 shell／内核重定向，成为“other_command”的标准输入，第 10 页到第 20 页被写至该标准输入。“other_command”的示例可以是 lp，它使输出在系统缺省打印机上打印。“other_command”的示例也可以 wc，它会显示选定范围的页中包含的行数、字数和字数。“other_command”可以是任何其它能从其标准输入读取的命令。错误消息仍在屏幕显示。

echo "test 10 start"
selpg -s=10 -e=20 selpg_input.txt 2> selpg_error.txt | lp
echo "test 10 complete"
# 10. 与上面的示例 9 相似，只有一点不同：错误消息被写至“selpg_error.txt”。

echo "test 11 start"
selpg -s=10 -e=20 -l=66 selpg_input.txt
echo "test 11 complete"
# 11. 该命令将页长设置为 66 行，这样 selpg 就可以把输入当作被定界为该长度的页那样处理。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。

echo "test 12 start"
selpg -s=10 -e=20 -f selpg_input.txt 2> selpg_error.txt
echo "test 12 complete"
# 12. 假定页由换页符定界。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。selpg 写至标准错误的所有内容都被重定向至“selpg_error.txt”。

