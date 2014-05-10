package gopusher

import (
    "testing"
)

func Test_Trigger(t *testing.T) {
    r := Trigger();
    
    if (r != "this will return a pusher response eventually") {
        t.Errorf("expected `%s` to be `foo`", r);    
    }
    
}