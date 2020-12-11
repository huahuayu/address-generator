package redis

const (
	// user session key, param is sid
	KEY_USER_SESSION = "app:session:%s"
	// key for insert user lock to prevent concurrent insert, first param is email, second is username
	KEY_INSERT_USER_LOCK = "app:i:u:lock:%s:%s"
)
