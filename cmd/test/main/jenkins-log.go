package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"time"
)

func main() {
	ctx := context.Background()
	jenkins := gojenkins.CreateJenkins(nil, "http://localhost:58080/", "admin", "P2ssw0rd")
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert

	_, err := jenkins.Init(ctx)
	if err != nil {
		panic(err.Error())
	}

	//job, err := jenkins.GetJob(ctx, "#jobname")
	//if err != nil {
	//	panic(err)
	//}
	//queueId, err := job.InvokeSimple(ctx, params)

	params := map[string]string{}
	queueId, err := jenkins.BuildJob(ctx, "ci_test_testng", params)
	if err != nil {
		panic(err)
	}
	build, err := jenkins.GetBuildFromQueueID(ctx, queueId)
	if err != nil {
		panic(err)
	}

	offset := int64(0)
	for build.IsRunning(context.TODO()) {
		time.Sleep(time.Second * 1)

		resp, _ := build.GetConsoleOutputFromIndex(context.TODO(), offset)
		offset = resp.Offset
		if len(resp.Content) > 0 {
			fmt.Print(resp.Content)
		}
	}

	fmt.Printf("build number %d with result: %v\n", build.GetBuildNumber(), build.GetResult())
}
