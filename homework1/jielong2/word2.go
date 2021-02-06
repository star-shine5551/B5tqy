package main

import "fmt"

func main() {
	var word byte                       //第一单词首字母
	var tmp string                      //获取单词的中间变量
	var re []string = make([]string, 0) //用于结果输出的切片
	var str map[byte]string             //存储单词的字典
	str = make(map[byte]string)         //首字母为key,单词为value

	fmt.Println("输入单词,以nil结尾：") //获取单词
	for tmp != "nil" {
		fmt.Scan(&tmp)
		str[tmp[0]] = tmp
	}
	fmt.Println("请输入第一个单词的首字母：")
	fmt.Scanf("%c", &word)
	tmp = str[word]      //字典键值对应值
	re = append(re, tmp) //切片追加

	for { //生成所需的词组
		tmp = str[tmp[len(tmp)-1]] //获取tmp尾字母开头的单词，长度减1，下标从零开始
		//fmt.Println(tmp)
		if tmp != "" {
			re = append(re, tmp)
		} else {
			break
		}

	}
	fmt.Println(re)
}
