package gopusher

import "errors"

const (
    MAX_PUSHER_CHANNELS = 100
)

func Trigger(channels []string) (string, error) {
    
    if (len(channels) > MAX_PUSHER_CHANNELS) {
        err := errors.New("error message")
        return "", err;
    }
    return "this will return a pusher response eventually", nil
}