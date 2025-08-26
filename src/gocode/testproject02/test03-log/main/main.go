package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LogStats struct {
	TotalRequests int            // 计数总请求数量
	IpCounts      map[string]int // 存ip和访问次数的对应关系
	StatusCode    map[string]int // 存code码和访问次数的对应关系
	TotalBytes    float64        // 访问产生的总流量
	ErrLines      int            // 不符合日志格式规范的行计数
}
type KeyValue struct { // 将IpCounts转换为当前结构体的切片，用于访问量从大到小的排序
	ipKey string
	ipNum int
}
type CodeNum struct { // 将StatusCode转换为当前结构体的切片，用于访问量从大到小的排序
	Code string
	Num  int
}

func main() {
	// 将本脚本编译好之后，判断用户输入参数是否符合规范
	if len(os.Args) < 3 {
		fmt.Printf("使用方法: %s <日志文件名> <top数量>\n", os.Args[0])
		fmt.Printf("示例: %s access.log 10\n", os.Args[0])
		os.Exit(1)
	}
	// 将第二个参数转为int64格式，因为下边getTopistIp需要将该参数传入
	TopNum64, strconvErr := strconv.ParseInt(os.Args[2], 10, 0)
	if strconvErr != nil {
		fmt.Println("输入的第二个参数有问题，请检查，请输入一个正整数代表获取ip访问数量排序的前N名")
		os.Exit(2)
	}
	TopNum := int(TopNum64)
	logFile := os.Args[1]
	//TopNum := os.Args[2]
	//logFile := "src/gocode/testproject02/test03-log/test.log"

	// 调用函数分析日志
	stats, err := analyzeLogFile(logFile)
	if err != nil {
		fmt.Printf("分析日志文件失败: %v\n", err)
		return
	}
	if len(stats.IpCounts) < TopNum || TopNum < 0 {
		fmt.Printf("第二个参数输入有误，根据日志文件分析得出，第二个参数输入应小于%d,并且大于0\n", len(stats.IpCounts))
		os.Exit(3)

	}
	s1 := getTopistIp(stats.IpCounts, TopNum)
	codeS1 := getTopistCode(stats.StatusCode)
	fmt.Println("日志分析报告：")
	fmt.Printf("总请求数%d\n", stats.TotalRequests)
	fmt.Printf("最频繁的访问者IP: %s (%d 次请求)", s1[0].ipKey, s1[0].ipNum)
	fmt.Println("HTTP状态码统计:")
	for _, v := range codeS1 {
		fmt.Printf("%s : %d\n", v.Code, v.Num)
	}
	fmt.Printf("总流量 %1f KB\n", stats.TotalBytes)
	fmt.Printf("总流量 %1f Mb\n", stats.TotalBytes/1024)
	fmt.Printf("总流量 %1f Gb\n", stats.TotalBytes/1024/1024)

}

func analyzeLogFile(filename string) (*LogStats, error) {
	// 打开文件
	file, err := os.Open(filename)
	// 如果文件打开失败，直接退出
	if err != nil {
		fmt.Println("文件打开失败。。。", err)
		return nil, err
	}
	// 创建缓冲区
	reader := bufio.NewReader(file)
	// 初始化结构体
	stats := &LogStats{
		IpCounts:   make(map[string]int),
		StatusCode: make(map[string]int),
	}
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {

				break
			}
			stats.ErrLines++
			return nil, fmt.Errorf("读取文件时出错: %w", err)
		}
		stats.TotalRequests++
		if stats.TotalRequests%10000 == 0 {
			fmt.Printf("日志已经处理%d行\n", stats.TotalRequests)
		}
		processLogLine(str, stats)
	}

	return stats, nil
}
func processLogLine(line string, stats *LogStats) {
	line = strings.TrimSpace(line)
	if line == "" {
		stats.ErrLines++
		return
	}
	parts := strings.Fields(line)
	if len(parts) < 7 {
		stats.ErrLines++
		return
	}
	f, err := strconv.ParseFloat(parts[9], 64)
	if err != nil {
		fmt.Println("转换错误，该行下标为9的访问字节数无法转为float64，下面将打印出该行具体内容，本行记作errlines+1")
		fmt.Println(line)
		stats.ErrLines++
		return
	}
	stats.TotalBytes += f
	ip := parts[0]
	code := parts[8]
	stats.IpCounts[ip]++
	stats.StatusCode[code]++

}

func getTopistIp(ipCounts map[string]int, TopNum int) []KeyValue {
	if TopNum > len(ipCounts) {

	}
	s1 := make([]KeyValue, 0, len(ipCounts))
	for k, v := range ipCounts {
		s1 = append(s1, KeyValue{k, v})
	}
	sort.Slice(s1, func(i, j int) bool {
		return s1[i].ipNum > s1[j].ipNum
	})
	return s1[:TopNum]
}
func getTopistCode(StatusCode map[string]int) []CodeNum {

	s1 := make([]CodeNum, 0, len(StatusCode))
	for k, v := range StatusCode {
		s1 = append(s1, CodeNum{k, v})
	}
	sort.Slice(s1, func(i, j int) bool {
		return s1[i].Num > s1[j].Num
	})
	return s1
}
