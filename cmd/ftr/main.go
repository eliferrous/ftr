package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/eliferrous/ftr/execmtr"
)

func main() {
	runner := execmtr.New()

	rep, err := runner.Run(context.Background(), "ipv6.google.com", 5)
	if err != nil {
		fmt.Println(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(rep) // pprint

}
