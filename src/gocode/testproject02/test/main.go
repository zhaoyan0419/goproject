package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 统计结果结构体
type LogStats struct {
	TotalLines int            // 总行数
	IPCounts   map[string]int // IP访问次数统计
	ErrorCount int            // 错误数量
}

func main() {
	// 要分析的日志文件路径
	logFile := "./access.log"

	// 调用函数分析日志
	stats, err := analyzeLogFile(logFile)
	if err != nil {
		fmt.Printf("分析日志文件失败: %v\n", err)
		return
	}
	fmt.Println(stats.TotalLines)
	// 输出统计结果
	fmt.Printf("日志分析完成！\n")
	fmt.Printf("总行数: %d\n", stats.TotalLines)
	fmt.Printf("错误行数: %d\n", stats.ErrorCount)
	fmt.Printf("独立IP数量: %d\n", len(stats.IPCounts))

	// 打印访问最频繁的5个IP
	fmt.Println("\n访问最频繁的IP地址:")
	for i, ip := range getTopIPs(stats.IPCounts, 5) {
		fmt.Printf("%d. %s (%d次访问)\n", i+1, ip, stats.IPCounts[ip])
	}
}

// 分析日志文件的主要函数
func analyzeLogFile(filename string) (*LogStats, error) {
	// 1. 打开日志文件
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close() // 确保函数结束时关闭文件

	// 2. 创建带缓冲的Reader（核心！）
	reader := bufio.NewReader(file)

	// 3. 初始化统计结果
	stats := &LogStats{
		IPCounts: make(map[string]int),
	}

	fmt.Println("开始分析日志文件...")

	// 4. 逐行读取日志文件（核心循环！）
	for {
		// 读取一行内容（直到遇到换行符）
		line, err := reader.ReadString('\n')

		// 处理读取错误或文件结束
		if err != nil {
			// 如果已经读到文件末尾，正常退出
			if err.Error() == "EOF" {
				break
			}
			// 如果是其他错误，返回错误
			return nil, fmt.Errorf("读取文件时出错: %w", err)
		}

		// 统计行数
		stats.TotalLines++

		// 每处理1000行输出一次进度（可选，对于大文件很有用）
		if stats.TotalLines%1000 == 0 {
			fmt.Printf("已处理 %d 行...\n", stats.TotalLines)
		}

		// 5. 处理每一行日志
		processLogLine(line, stats)
	}

	return stats, nil
}

// 处理单行日志的函数
func processLogLine(line string, stats *LogStats) {
	// 移除行首尾的空白字符
	line = strings.TrimSpace(line)

	// 跳过空行
	if line == "" {
		return
	}

	// 示例：简单的Nginx日志格式解析
	// 真实日志格式：127.0.0.1 - - [10/Oct/2023:14:32:01 +0800] "GET /api/user HTTP/1.1" 200 1234
	parts := strings.Fields(line) // 按空格分割行
	fmt.Println(parts)
	if len(parts) < 7 {
		stats.ErrorCount++
		return // 格式错误，跳过
	}

	// 提取IP地址（通常是第一个字段）
	ip := parts[0]

	// 统计IP访问次数
	stats.IPCounts[ip]++

	// 这里可以添加更多的日志分析逻辑，例如：
	// - 分析HTTP状态码（parts[8]）
	// - 分析请求路径（parts[6]）
	// - 分析响应时间（parts[9]）
}

// 获取访问最频繁的IP（辅助函数）
func getTopIPs(ipCounts map[string]int, topN int) []string {
	// 这里简化实现，实际应该用排序算法
	var topIPs []string
	for ip := range ipCounts {
		if len(topIPs) < topN {
			topIPs = append(topIPs, ip)
		}
	}
	return topIPs
}
