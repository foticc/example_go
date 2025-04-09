package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

// 下载器结构体
type Downloader struct {
	concurrencyLimit chan struct{} // 并发控制通道
	wg               sync.WaitGroup
	client           *http.Client
}

func NewDownloader(maxConcurrent int) *Downloader {
	return &Downloader{
		concurrencyLimit: make(chan struct{}, maxConcurrent),
		client: &http.Client{
			Timeout: 30 * time.Second, // 设置超时时间
		},
	}
}

func (d *Downloader) DownloadImage(url, savePath string) error {
	// 控制并发：发送空结构体到通道（如果通道满会阻塞）
	d.concurrencyLimit <- struct{}{}
	defer func() {
		<-d.concurrencyLimit // 完成后释放通道位置
	}()

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头（模拟浏览器）
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	// 执行请求
	resp, err := d.client.Do(req)
	if err != nil {
		return fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("非200状态码: %d", resp.StatusCode)
	}

	// 创建保存目录
	if err := os.MkdirAll(path.Dir(savePath), 0755); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}

	// 创建文件
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer file.Close()

	// 复制响应内容到文件
	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	return nil
}

func main() {
	// 示例图片URL列表
	urls := []string{
		"https://t0.tianditu.gov.cn/vec_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=vec&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILEMATRIX=17&TILECOL=109334&TILEROW=54026&tk=f8cc3dcbfc3665361e844046c4e05709",
		// 添加更多URL...
	}

	// 创建下载器（限制最大并发数为30）
	downloader := NewDownloader(30)

	// 使用等待组跟踪所有下载
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(index int, url string) {
			defer wg.Done()

			// 生成保存路径（根据URL自动生成文件名）
			// filename := path.Base(url)
			filename := "123.jpg"
			savePath := fmt.Sprintf("./downloads/%d_%s", index, filename)

			// 执行下载
			if err := downloader.DownloadImage(url, savePath); err != nil {
				fmt.Printf("下载失败 [%s]: %v\n", url, err)
			} else {
				fmt.Printf("成功下载: %s\n", savePath)
			}
		}(i, url)
	}

	// 等待所有下载完成
	wg.Wait()
	fmt.Println("所有下载任务完成")
}
