package main
import "os"
import "io"
import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func outputFile(path string) {
	buf := make([]byte, 4096)
	f, err := os.Open(path)
	check(err)

	for {
		n, err := f.Read(buf)
		if n > 0 {
			os.Stdout.Write(buf[:n])
		}
		if (err == io.EOF) {
			break
		}
		check(err)
	}
	f.Close()
}

func readstdio() {
	buf := make([]byte, 4096)
	for {
		n, err := os.Stdin.Read(buf)
		if (n > 0) {
			os.Stdout.Write(buf[:n])
		}
		if (err == io.EOF) {
			break
		}
		check(err)
		if (err != nil) {
			panic(err)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		readstdio()
	}
	fmt.Println(os.Args)
	if os.Args[1:] != nil {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i] == "-" {
				readstdio()
			} else {
				outputFile(os.Args[i])
			}
			
		}
	}
}