package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

	if(len(os.Args) != 2) {
		fmt.Printf("Usage: %s <file name>\n",os.Args[0])
		return
	}
	filename := os.Args[1]

	f, err := os.Open(filename)
	check(err)

	pre_s := "ffmpeg -i "
	mid_s := " -codec:a libmp3lame -qscale:a 5 "
	post_s := ""

	bf := bufio.NewReader(f)
	for {
		switch line, err := bf.ReadString('\n'); err {
		case nil:
			fn := strings.TrimSpace(line)
			fn1 := "\"D:\\put_here\\eric collection" + fn + "\""
			fn2 := "\"D:\\mp3\\eric collection 1" + fn + "\""
			fmt.Printf("%s%-70s%s%-70s%s\n", pre_s, fn1, mid_s, fn2, post_s)
		case io.EOF:
			if line > "" {
			// last line of file missing \n, but still valid
				fmt.Println(line)
			}
			return
		default:
		log.Fatal(err)
        }
    }
    f.Close()
}

