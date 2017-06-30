// +build !darwin,!windows

package notify

func newNotifier(dir string) (Notifier, error) {
	return &emptyNotifier{}, nil
}
