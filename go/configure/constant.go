package configure
const (
	// logFormat
	LOGFORMAT_JSON    = "json"
	LOGFORMAT_CONSOLE = "console"

	// EncoderConfig
	TIME_KEY       = "time"
	LEVLE_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "caller"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"

	// 日志归档配置项
	// 每个日志文件保存的最大尺寸 单位：M
	MAX_SIZE = 1
	// 文件最多保存多少天
	MAX_BACKUPS = 5
	// 日志文件最多保存多少个备份
	MAX_AGE = 7
)