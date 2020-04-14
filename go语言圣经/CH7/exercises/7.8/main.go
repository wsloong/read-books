// 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：
// 主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
// 定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方法
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("4m24s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("3m37s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	// 使用 tabwriter 来生成一个列是整齐对齐和隔开的表格
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	//tw = tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)	// 和上面的一行相等
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	//  会格式化整个表格并且将它写向os.Stdout（标准输出）
	tw.Flush() // calculate column widths and print table
}

type lessFunc func(i, j int) bool

type tableSort struct {
	t         []*Track
	lessFuncs []lessFunc
}

func (x tableSort) Len() int { return len(x.t) }
func (x tableSort) Less(i, j int) bool {
	for k := len(x.lessFuncs) - 1; k >= 0; k-- {
		// 如果最后一个直接可以确定
		if x.lessFuncs[k](i, j) {
			return true
		} else if !x.lessFuncs[k](j, i) { // 判断是否相等，相等就执行下一个less函数
			continue
		}

		return false
	}
	return false
}

func (x tableSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func (x tableSort) byTitle(i, j int) bool {
	return x.t[i].Title < x.t[j].Title
}

func (x tableSort) byArtist(i, j int) bool {
	return x.t[i].Artist < x.t[j].Artist
}

func (x tableSort) byYear(i, j int) bool {
	return x.t[i].Year < x.t[j].Year
}

func (x tableSort) byLength(i, j int) bool {
	return x.t[i].Length < (x.t[j].Length)
}

func sortByTable() {
	t := tableSort{t: tracks}
	t.lessFuncs = append(t.lessFuncs, t.byLength, t.byTitle)
	sort.Sort(t)

	printTracks(tracks)
}

func sortOneByOne() {
	t := tableSort{t: tracks}
	t.lessFuncs = append(t.lessFuncs, t.byTitle)
	sort.Sort(t)

	t.lessFuncs[0] = t.byLength
	sort.Sort(t)

	printTracks(tracks)
}

func main() {
	// table调用是第一个不确定才会使用第二个键
	sortByTable()

	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println()

	// 一个一个调用是后面的在前一个的基础上再次排序，两次结果不一样
	sortOneByOne()
}
