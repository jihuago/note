package readAndWriteData

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadCsv()  {
	inputFile, err := os.Open("./public/products.txt")

	if err != nil {
		panic(err)
	}

	r := csv.NewReader(inputFile)

	// 修改分隔符
	r.Comma = ';'

	for {
		record, err := r.Read()

		// 一旦读取到文件末尾，err的值将变成非空（其值常量io.EOF）
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
