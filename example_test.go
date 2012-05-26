package stathat_test

import (
        "github.com/stathat/go"
        "log"
        "time"
)

func ExamplePostCount() {
        err := stathat.PostEZCountOne("go example test run", "info@stathat.com")
        if err != nil {
                log.Printf("error posting ez count one: %v", err)
                return
        }
        stathat.WaitUntilFinished(5 * time.Second)
}

