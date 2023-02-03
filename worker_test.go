package worker

import (
	"fmt"
	"testing"
)

// TestJob is a simple implementation of the Job interface
type TestJob struct {
	val int
}

func (tj *TestJob) Do(errCh chan error) {
	if tj.val == 0 {
		errCh <- fmt.Errorf("TestJob with value 0")
	}
}

func TestPool(t *testing.T) {
	p := NewPool(2)
	p.Enqueue(&TestJob{val: 1})
	p.Enqueue(&TestJob{val: 0})
	p.Enqueue(&TestJob{val: 1})
	p.StopQueueingJob()
	errCount := 0
	for err := range p.Errors() {
		errCount++
		if err.Error() != "TestJob with value 0" {
			t.Fatalf("Unexpected error: %v", err)
		}
	}
	if errCount != 1 {
		t.Fatalf("Expected 1 error, got %d", errCount)
	}
}
