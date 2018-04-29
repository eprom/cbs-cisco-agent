package main

import (
	//"github.com/jakhog/cbs-cisco-agent/cisco/sms"
	//"github.com/jakhog/cbs-cisco-agent/cisco/status"
	"github.com/jakhog/cbs-cisco-agent/config"
	"github.com/jakhog/cbs-cisco-agent/sms"
	"github.com/jakhog/cbs-cisco-agent/storage"
	//"errors"
	"github.com/jakhog/cbs-cisco-agent/log"
	"github.com/jakhog/cbs-cisco-agent/task"
	"sync"
	"time"
)

func main() {
	logger := log.NewLogger("Agent")
	logger.Info("Starting Agent tasks")

	wg := &sync.WaitGroup{}

	// Config reader task
	configTask := config.NewReaderTask()
	configRunner := task.NewRunner(configTask, time.Second*30)
	configRunner.RunImmediately()
	configRunner.Start(wg)

	// Incoming SMS reader
	incomingTask := sms.NewIncomingTask()
	incomingRunner := task.NewRunnerWithMax(incomingTask, time.Second*10, time.Second*30)
	incomingRunner.Start(wg)

	// Persistent storage writer
	storageTask := storage.NewTask()
	storageRunner := task.NewRunner(storageTask, time.Second*10)
	storageRunner.RunImmediately()
	storageRunner.Start(wg)

	logger.Info("Agent tasks started")
	wg.Wait()

	/*
		logger.Info("Loading config...")

		cfg, err := config.Read("package_config.ini")
		if err != nil {
			logger.Error(err)
		}

		logger.Info("Getting status...")
		stat, err := status.GetGeneralStatus(&cfg.Cisco)
		if err != nil {
			logger.Error("Status error", err)
		}
		logger.Info("Got status!")
		fmt.Println(stat)
	*/

	/*
		fmt.Println("Getting smses...")
		smses, err := sms.GetAllSMSes(&cfg.Cisco)
		if err != nil {
			fmt.Println("Error", err)
		}
		for _, sms := range smses {
			fmt.Println("SMS", sms.From, sms.Received, sms.Size)
		}
	*/
	/*
		fmt.Println("Getting SMS status...")
		status, err := sms.GetSMSStatus(&cfg.Cisco)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			fmt.Println("Status", status)
		}

		fmt.Println("Sending SMS")
		err = sms.SendSMS(&cfg.Cisco, sms.OutgoingSMS{
			To:   "+4741645532",
			Text: "I'm still here\n Are you?",
		})
		if err != nil {
			fmt.Println("Error", err)
		}
	*/
}
