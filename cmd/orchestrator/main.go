package main

import (
	"context"
	"fmt"
	"os"

	"github.com/hashim715/mini-container-orchestrator/internal/runtime"
);

func main() {
	rt, err := runtime.NewDockerRuntime();

	if (err != nil) {
		os.Exit(1);
	};

	defer rt.Close();

	ctx := context.Background();

	containers,err := rt.ListContainers(ctx);

	if err != nil {
		fmt.Printf("error in listing conatiners and the error is: %v\n",err);
		return;
	};

	for _,id := range containers {
		fmt.Println(id);
	};
};