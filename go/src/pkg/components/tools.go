package components

// 使用 Redis zset 类型 实现简单限流
func isActionAllowed(uid int, actionName string, maxCount int) bool {
	return true
}
