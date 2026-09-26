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

	for _,container := range containers {
		fmt.Println(container.ID);
	};

	resp, err := rt.CreateContainer(ctx, "alpine",[]string{"echo","hello world"},[]string{},"my-test-container");

	if err != nil {
		fmt.Printf("error in creating conatiners and the error is: %v\n",err);
		return;
	};

	if (resp != "") {
		fmt.Printf("the response from creating container is: %v\n",resp);
	};
};