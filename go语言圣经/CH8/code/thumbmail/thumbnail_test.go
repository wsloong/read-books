// 书中的代码示例
package main

import (
	"log"
	"os"
	"sync"
)

// 制作缩略图
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 不正确的！
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		// 相当于只是启动了多个goroutine
		// 但是还没有执行程序就退出了
		go ImageFile(f)
	}
}

// 并发的生成缩略图
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f) // 这里忽略的错误处理
			ch <- struct{}{}
		}(f) // 这里很重要！！
	}

	// 等待goroutine完成
	for range filenames {
		<-ch
	}
}

// 并发执行，如果有一个失败返回错误
func makeThumbnails4(filenames []string) {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	// 书上说如果遇到第一个非nil的error时候会发生goroutine泄露
	// 我测试发现并不会
	// <-errors是从里面消费数据，哪怕是nil也会拿出来消费掉，只是nil的时候不return
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

// 并发执行
// 这里定义了一个struct，用于接收ImageFile的返回
// 并且使用了有缓存的channel
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// makeThumbnails6为从通道接收的生成的缩略图。
//它返回它创建的文件所占用的字节数
func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()

			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// 等待所有的goroutine执行完毕，关闭channel
	// 这个操作不能放到main goroutine中
	// 如果放到main goroutine中，因为使用了无缓冲的channel(sizes := make(chan int64))
	// 从无缓冲的chann接收值操作在wait()之后，所以发送会阻塞(就是sizes <- info.Size())，
	// 导致`defer wg.Done()`代码无法执行，wait()就会阻塞
	// 如果要放到main goroutine中，可以使用有缓冲的channel来解决
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
