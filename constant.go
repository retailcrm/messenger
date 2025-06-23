package messenger

type ThreadControlType string

const (
	PassThreadControl    ThreadControlType = "pass_thread_control"
	ReleaseThreadControl ThreadControlType = "release_thread_control"
)
