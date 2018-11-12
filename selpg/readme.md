# selpg

---

## 程序说明
selpg是一个命令行工具，允许用户指定该程序从标准输入或从作为命令行参数给出的文件名读取文本输入，并允许用户指定标准输出位置和输入输出的页范围。可以有选择性的查看或打印某个文档的部分内容，简单高效节约。

---

本程序使用go语言实现，具体的用法以及c语言实现参考[开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html)

---

关于selpg的实现思路和测试以及用到的相关Linux命令行准则的知识可以参考我的博客
- [csdn：selpg-CLI实用程序](https://blog.csdn.net/Wonderful_sky/article/details/83021169)
- [gitgiter：CLI实用程序之selpg](https://gitgiter.github.io/2018/10/11/SC-hw2-selpg/)

## 文件说明
- selpg.go：核心功能文件，selpg功能的所有实现，包括错误处理
- selpg_input_generator.sh：批处理文件，用于生成selpg_input.txt
- selpg_test.sh：批处理文件，用于批量测试selpg
- selpg_input.txt：可选输入文件，内容为1~10000的递增数列，用于对selpg进行一系列测试
- selpg_output.txt：可选输出文件，用于测试selpg的输出重定向
- selpg_error.txt：可选错误输出文件，用于测试selpg的错误输出重定向

## 参数格式说明
- -s Num：指定起始页码为Num，必选项，默认为-1（不合法，强制要求重新指定）
- -e Num：指定结束页码为Num，必选项，默认为-1（不合法，强制要求重新指定）
- filename：指定输入文件，可选项，默认为空
- -l Num：指定每页的行数为Num，可选项，默认为每页72行，并且为默认解读模式
- -f：指定解读模式为，'\f'作为页分隔符，不可与-l同时使用
- -d destination：指定打印机设备地址