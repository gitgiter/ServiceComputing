package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
)

type selpgArgs struct {
	start    int    // index of start page
	end      int    // index of end page
	filename string // name of input file
	length   int    // page length --- number of lines, default 72
	pageType bool   // -f: true, form-feed-delimited;	-l: false, lines-delimited --- default l
	des      string // device destination of printer
}

func inputArgs(args *selpgArgs) {
	flag.IntVar(&(args.start), "s", -1, "start page")
	flag.IntVar(&(args.end), "e", -1, "end page")
	flag.IntVar(&(args.length), "l", -1, "page length")
	flag.BoolVar(&(args.pageType), "f", false, "page type")
	flag.StringVar(&(args.des), "d", "", "print destination")

	flag.Parse()

	// deal with -l and -f, especially -l 72 and -f
	if args.length != -1 && args.pageType == true {
		fmt.Fprintf(os.Stderr, "-l Num and -f cannot be used together\n")
		os.Exit(1)
	}
	if args.length == -1 {
		args.length = 72
	}

	others := flag.Args()
	if len(others) == 1 {
		args.filename = others[0]
	} else if len(others) == 0 {
		args.filename = ""
	} else {
		fmt.Fprintf(os.Stderr, "too many arguments\n")
		os.Exit(2)
	}
}

func checkArgs(args *selpgArgs) {
	if args.start < 1 || args.end < 1 {
		fmt.Fprintf(os.Stderr, "the number of start page and end page should start from 1\n")
		os.Exit(3)
	}
	if args.start > math.MaxInt32-1 || args.end > math.MaxInt32-1 {
		fmt.Fprintf(os.Stderr, "the number of start page and end page should not larger than MAX INT\n")
		os.Exit(4)
	}
	if args.start > args.end {
		fmt.Fprintf(os.Stderr, "the number of start page should not larger than the number of end page\n")
		os.Exit(5)
	}
	if args.length < 1 || args.length > math.MaxInt32-1 {
		fmt.Fprintf(os.Stderr, "invalid page length\n")
		os.Exit(6)
	}
}

func processData(args *selpgArgs) {

	var reader *bufio.Reader
	if args.filename == "" {
		reader = bufio.NewReader(os.Stdin)
	} else {
		filein, err := os.Open(args.filename)
		defer filein.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file %s\n", args.filename)
			os.Exit(7)
		}
		reader = bufio.NewReader(filein)
	}

	pageCount := 1

	if args.des == "" {
		writer := bufio.NewWriter(os.Stdout)
		if args.pageType {
			pageCount = optionF(args, reader, writer)
		} else {
			pageCount = optionL(args, reader, writer)
		}
	} else {
		cmd := exec.Command("lp", "-d", args.des)
		writer, err := cmd.StdinPipe()
		defer writer.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "stdin pipe error\n")
			os.Exit(8)
		}

		err = cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "command start error\n")
			os.Exit(9)
		}

		if args.pageType {
			pageCount = optionFD(args, reader, writer)
		} else {
			pageCount = optionLD(args, reader, writer)
		}

		err = cmd.Wait()
		if err != nil {
			fmt.Fprintf(os.Stderr, "command wait error\n")
			os.Exit(10)
		}
	}

	if pageCount < args.start {
		fmt.Fprintf(os.Stderr, "the number of start page does not exist\n")
		os.Exit(11)
	}
	if pageCount < args.end {
		fmt.Fprintf(os.Stderr, "the number of end page does not exist\n")
		os.Exit(12)
	}
}

func optionF(args *selpgArgs, reader *bufio.Reader, writer *bufio.Writer) (pageCount int) {
	pageCount = 1
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "failed to read byte\n")
			os.Exit(13)
		}
		if pageCount >= args.start && pageCount <= args.end {
			writeErr := writer.WriteByte(ch)
			if writeErr != nil {
				fmt.Fprintf(os.Stderr, "failed to write byte\n")
				os.Exit(14)
			}
			writer.Flush()
		}
		if ch == '\f' {
			pageCount++
		}
	}
	return pageCount
}

func optionL(args *selpgArgs, reader *bufio.Reader, writer *bufio.Writer) (pageCount int) {
	pageCount = 1
	lineCount := 1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "failed to read bytes, %s\n", err.Error())
			os.Exit(15)
		}

		if pageCount >= args.start && pageCount <= args.end {
			_, writeErr := writer.Write(line)
			if writeErr != nil {
				fmt.Fprintf(os.Stderr, "failed to write bytes\n")
				os.Exit(16)
			}
			writer.Flush()
		}
		if lineCount >= args.length {
			pageCount++
			lineCount = 1
		} else {
			lineCount++
		}
	}
	return pageCount
}

func optionFD(args *selpgArgs, reader *bufio.Reader, writer io.WriteCloser) (pageCount int) {
	pageCount = 1
	for {
		ch, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "failed to read byte\n")
			os.Exit(13)
		}
		if pageCount >= args.start && pageCount <= args.end {
			_, writeErr := writer.Write([]byte{ch})
			if writeErr != nil {
				fmt.Fprintf(os.Stderr, "failed to write bytes\n")
				os.Exit(16)
			}
		}
		if ch == '\f' {
			pageCount++
		}
	}
	return pageCount
}

func optionLD(args *selpgArgs, reader *bufio.Reader, writer io.WriteCloser) (pageCount int) {
	pageCount = 1
	lineCount := 1
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "failed to read bytes\n")
			os.Exit(15)
		}
		if pageCount >= args.start && pageCount <= args.end {
			_, writeErr := writer.Write(line)
			if writeErr != nil {
				fmt.Fprintf(os.Stderr, "failed to write bytes\n")
				os.Exit(16)
			}
		}
		if lineCount >= args.length {
			pageCount++
			lineCount = 1
		} else {
			lineCount++
		}
	}
	return pageCount
}

func main() {
	args := new(selpgArgs)
	inputArgs(args)
	checkArgs(args)
	processData(args)
	fmt.Println("================done!================")
	//fmt.Printf("%+v\n", args)
}
