package main

import (
	"fmt"
	"os"
)

var migong [10][10]int
var flag = 0

func Go(e, f, g, h, n, m int) {
	fmt.Printf("(%d,%d)\n", e, f) //每次移动打印位置
	migong[e][f] = 2
	if e == g && f == h {
		flag = 1
		fmt.Printf("这就是外面的世界")
		os.Exit(0) //终止程序
	}
	if flag != 1 && e >= 1 && migong[e-1][f] == 0 {
		Go(e-1, f, g, h, n, m) //判断向上是否有路
	}
	if flag != 1 && migong[e+1][f] == 0 {
		if e+1 < n { //限制规模，不要超出[N][M]数组迷宫
			Go(e+1, f, g, h, n, m) //判断向下是否有路
		}
	}
	if flag != 1 && migong[e][f+1] == 0 {
		if f+1 < m { //限制规模，不要超出[N][M]数组迷宫
			Go(e, f+1, g, h, n, m) //判断向右是否有路
		}
	}
	if flag != 1 && f >= 1 && migong[e][f-1] == 0 {
		Go(e, f-1, g, h, n, m) //判断向左是否有路
	}

}

func main() {
	var N, M, T, i, j int
	var a, b, c, d int
	fmt.Printf("输入迷宫规模和障碍个数\n")
	fmt.Scanf("%d %d %d", &N, &M, &T)
	fmt.Printf("请输入起点坐标和终点坐标:\n")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	fmt.Printf("输入障碍坐标:\n")
	for k := 0; k < T; k++ {
		fmt.Scanf("%d %d", &i, &j)
		migong[i][j] = 1 //迷宫障碍定为1
	}
	Go(a, b, c, d, N, M) //传入起点终点坐标和迷宫规模大小
}
