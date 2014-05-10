package gopusher

import "errors"

const (
    MAX_PUSHER_CHANNELS = 100
    TOO_MANY_CHANNELS_ERROR_MSG = "an event can be triggered on a maximum of 100 channels in a single call."
)

func Trigger(channels []string) (string, error) {
    
    if (len(channels) > MAX_PUSHER_CHANNELS) {
        err := errors.New(TOO_MANY_CHANNELS_ERROR_MSG)
        return "", err;
    }
    return "this will return a pusher response eventually", nil
}