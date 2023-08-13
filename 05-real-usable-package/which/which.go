package which

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Witch(filename string) {
	now := time.Now()

	result := false
	fmt.Println("Environment PATH IS")
	fmt.Println(os.Getenv("PATH"))
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, filename)
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() {
				if mode&0111 != 0 {
					result = true
					break
				}
			}
		}
	}
	fmt.Println("result is : ", result)
	fmt.Println("time : ", time.Since(now)) //75~100 ms
}
