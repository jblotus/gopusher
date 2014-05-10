package gopusher

import (
    "errors"
    "fmt"
)

const (
    MAX_PUSHER_CHANNELS = 100
    TOO_MANY_CHANNELS_ERROR_MSG = "an event can be triggered on a maximum of 100 channels in a single call."
)

func Trigger(channels []string, appid int) (string, error) {
    
    if (len(channels) > MAX_PUSHER_CHANNELS) {
        err := errors.New(TOO_MANY_CHANNELS_ERROR_MSG)
        return "", err;
    }
    
    pusherurl := fmt.Sprintf("/apps/%s", appid);
    
    
    return pusherurl, nil
}