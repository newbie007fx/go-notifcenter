package notif

type Retry struct {
	Count   uint
	Backoff Backoff
}
