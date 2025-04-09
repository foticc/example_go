package main

import (
	"fmt"
	"io"
	"math"
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

// 获取瓦片的x和y编号
func latLngToTile(lat float64, lon float64, zoom int) (int, int) {
	n := math.Pow(2, float64(zoom))
	x := int((lon + 180.0) / 360.0 * n)
	y := int((1.0 - (math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0)) / math.Pi)) / 2.0 * n)
	return x, y
}

// 获取给定范围的所有瓦片
func getTilesInRange(minLat, minLon, maxLat, maxLon float64, zoom int) []Tile {
	var tiles []Tile

	xMin, yMin := latLngToTile(maxLat, minLon, zoom)
	xMax, yMax := latLngToTile(minLat, maxLon, zoom)

	for x := xMin; x <= xMax; x++ {
		for y := yMin; y <= yMax; y++ {
			tile := Tile{
				x:    x,
				y:    y,
				zoom: zoom,
			}
			tiles = append(tiles, tile)
		}
	}

	return tiles
}

type Tile struct {
	x    int
	y    int
	zoom int
}

type URI struct {
	filename string
	uri      string
	typ      string
	zoom     int
	x        int
	y        int
}

func format(col int, row int, zoom int) []URI {
	var arr = []URI{}
	vec_w := "https://t0.tianditu.gov.cn/vec_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=vec&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILEMATRIX=%d&TILECOL=%d&TILEROW=%d&tk=f8cc3dcbfc3665361e844046c4e05709"
	cia_w := "https://t0.tianditu.gov.cn/cia_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=cia&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILEMATRIX=%d&TILECOL=%d&TILEROW=%d&tk=f8cc3dcbfc3665361e844046c4e05709"
	img_w := "https://t0.tianditu.gov.cn/img_w/wmts?SERVICE=WMTS&REQUEST=GetTile&VERSION=1.0.0&LAYER=img&STYLE=default&TILEMATRIXSET=w&FORMAT=tiles&TILEMATRIX=%d&TILECOL=%d&TILEROW=%d&tk=f8cc3dcbfc3665361e844046c4e05709"

	file_path := "D:/html/hi-rust/tianmap-downloader/map/%s/%d/%d-%d.png"

	arr = append(arr, URI{filename: fmt.Sprintf(file_path, "vec_w", zoom, col, row), uri: fmt.Sprintf(vec_w, zoom, col, row), typ: "vec_w", zoom: zoom, x: col, y: row})
	arr = append(arr, URI{filename: fmt.Sprintf(file_path, "cia_w", zoom, col, row), uri: fmt.Sprintf(cia_w, zoom, col, row), typ: "cia_w", zoom: zoom, x: col, y: row})
	arr = append(arr, URI{filename: fmt.Sprintf(file_path, "img_w", zoom, col, row), uri: fmt.Sprintf(img_w, zoom, col, row), typ: "img_w", zoom: zoom, x: col, y: row})
	return arr
}

func downloads(urls []URI) {
	// 创建下载器（限制最大并发数为30）
	downloader := NewDownloader(6)

	// 使用等待组跟踪所有下载
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go func(index int, url URI) {
			defer wg.Done()

			// 生成保存路径（根据URL自动生成文件名）
			// filename := path.Base(url)
			filename := url.filename

			// 执行下载
			if err := downloader.DownloadImage(url.uri, filename); err != nil {
				fmt.Printf("下载失败 [%s]: %v\n", url.uri, err)
				os.Exit(1)
			} else {
				fmt.Printf("成功下载: %s\n", filename)
			}
		}(i, url)
	}

	// 等待所有下载完成
	wg.Wait()
	fmt.Println("所有下载任务完成")
}

func main() {
	// 示例范围：西南角(34.0, -118.0)到东北角(34.5, -117.5)
	minLat, minLon := 27.36, 118.50
	maxLat, maxLon := 29.05, 120.66

	var urls = []URI{}

	// 计算1到18层级的瓦片
	for zoom := 4; zoom <= 11; zoom++ {
		tiles := getTilesInRange(minLat, minLon, maxLat, maxLon, zoom)
		fmt.Printf("Zoom level %d:\n", zoom)
		for _, tile := range tiles {
			urls = append(urls, format(tile.x, tile.y, zoom)...)
		}
	}
	downloads(urls)
	// fmt.Println("len", len(urls))
	// var real_urls = []URI{}
	// for _, number := range urls {

	// 	if number.zoom >= 14 && number.x >= 13650 {
	// 		real_urls = append(real_urls, number)
	// 	}
	// }
	// fmt.Println("{real_urls}", len(real_urls))
	// downloads(real_urls)

}
