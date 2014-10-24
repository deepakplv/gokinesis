package main

import (
	"fmt"
	"kinesis"
)

type EchoConsumer struct {
	numProcessed	int64
}

func (ec *EchoConsumer) Init(shardId string) error {
	fmt.Printf("init: %s\n", shardId)
	return nil
}

func (ec *EchoConsumer) ProcessRecords(records []*kinesis.KclRecord,
	checkpointer *kinesis.Checkpointer) error {
	for i := range records {
		fmt.Printf("process: %s\n", records[i].DataB64)
		ec.numProcessed++
	}

	// just for fun
	if ec.numProcessed % 3 == 0 {
		checkpointer.CheckpointAll()
	}
	return nil
}

func (ec *EchoConsumer) Shutdown(shutdownType kinesis.ShutdownType,
	checkpointer *kinesis.Checkpointer) error {
	fmt.Printf("shutdown: %s\n", shutdownType)
	if shutdownType == kinesis.GracefulShutdown {
		checkpointer.CheckpointAll()
	}
	return nil
}

func main() {
	var ec EchoConsumer
	kinesis.Run(&ec)
}
