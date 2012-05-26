package stathat_test

import (
        "github.com/stathat/go"
        "log"
        "time"
)

func ExamplePostEZCountOne() {
        log.Printf("starting example")
        stathat.Verbose = true
        err := stathat.PostEZCountOne("go example test run", "patrick@stathat.com")
        if err != nil {
                log.Printf("error posting ez count one: %v", err)
                return
        }
        stathat.WaitUntilFinished(5 * time.Second)
        // Output: something
}

