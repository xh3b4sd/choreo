package backoff

// Interface describes a backoff implementation for executing an error function
// and retrying its execution on failure according to the underlying backoff
// settings.
type Interface interface {
	Backoff(fnc func() error) error
}
