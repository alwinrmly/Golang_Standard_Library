package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"alwin", "ramli", "sasmita"})
	_ = writer.Write([]string{"alvionita", "mirwanthy", "rogaye"})
	_ = writer.Write([]string{"bejo", "tejo", "kronjo"})

	writer.Flush()
}
