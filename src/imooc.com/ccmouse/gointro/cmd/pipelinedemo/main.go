package main

import (
	"imooc.com/ccmouse/gointro/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const filename = "large.in"
	const n = 100000
	file, err:= os.Create(filename)
	if err!= nil {
		panic(err)
	}
	defer  file.Close()

	p:= pipeline.RandomSource(n)

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p=pipeline.ReaderSource(bufio.NewReader(file))
	count := 0
	for v:=range p{
		fmt.Println(v)
		count ++
		if count>10 {
			break
		}
	}

}

func mergeDemo()  {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(1, 2, 6, 7, 14)),
		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8)))

	/*for {
		if num, ok := <-p; ok {
			fmt.Println(num)
		}else{
			break
		}
	}*/
	for v := range p {
		fmt.Println(v)
	}
}