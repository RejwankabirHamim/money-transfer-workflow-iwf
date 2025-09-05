package iwf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"net/http"
	"sync"
)

func BuildCLI() *cli.App {
	app := cli.NewApp()
	app.Name = "iwf money transfer"
	app.Usage = "money transfer workflow using iWF"
	app.Version = "v0.1"

	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "start iWF worker",
			Action: start,
		},
	}
	return app
}

func start(c *cli.Context) {
	fmt.Println("starting money transfer workflow worker...")
	closeFn := startWorkflowWorker()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
	closeFn()
}

var client = iwf.NewClient(workflows.GetRegistry(), nil)
var workerService = iwf.NewWorkerService(workflows.GetRegistry(), nil)

func startWorkflowWorker() (closeFunc func()) {
	router := gin.Default()

	// iWF internal APIs
	router.POST(iwf.WorkflowStateWaitUntilApi, apiV1WorkflowStateStart)
	router.POST(iwf.WorkflowStateExecuteApi, apiV1WorkflowStateDecide)
	router.POST(iwf.WorkflowWorkerRPCAPI, apiV1WorkflowWorkerRpc)

	// money transfer workflow APIs
	router.GET("/moneytransfer/start", workflows.StartMoneyTransferWorkflow)
	router.GET("/moneytransfer/describe", workflows.DescribeMoneyTransferWorkflow)

	server := &http.Server{
		Addr:    ":" + iwf.DefaultWorkerPort,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	return func() { server.Close() }
}
