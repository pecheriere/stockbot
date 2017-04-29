package apiai

import (
	"testing"
)

func TestFulfill(t *testing.T) {
	if err, _ := Fulfill(nil); err == nil {
		t.Fatalf("Should have returned an error")
	}

	request :=  new(FulfillmentRequest)
	if err, _ := Fulfill(request); err == nil {
		t.Fatalf("Should have returned an error")
	}

	request.Result.Action = "test"
	if err, response := Fulfill(request); err != nil {
		t.Fatalf("Should not have returned an error")
	} else if response.Speech != "Hello my dear friend. Your test is well received !" {
		t.Fatalf("Not the correct answer")
	}

	request.Result.Action = "actionTest"
	if err, response := Fulfill(request); err != nil {
		t.Fatalf("Should not have returned an error")
	} else if response.Speech != "I don't underdstand what you want to do" {
		t.Fatalf("Not the correct answer")
	}

}