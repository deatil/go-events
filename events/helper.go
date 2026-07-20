package events

// 默认事件
// default new events
var defaultEvent = New()

// 默认事件
// default new events
func Default() *Events {
	return defaultEvent
}

// 注册操作
// Add Action
func AddAction(event any, listener any, sort ...int) {
	useSort := DefaultSort
	if len(sort) > 0 {
		useSort = sort[0]
	}

	defaultEvent.Action().Listen(event, listener, useSort)
}

// 触发操作
// Do Action
func DoAction(event any, params ...any) {
	defaultEvent.Action().Trigger(event, params...)
}

// 移除操作
// Remove Action
func RemoveAction(event string, listener any, sort ...int) bool {
	useSort := DefaultSort
	if len(sort) > 0 {
		useSort = sort[0]
	}

	return defaultEvent.Action().RemoveListener(event, listener, useSort)
}

// 是否有操作
// Has Action
func HasAction(event string, listener any) bool {
	return defaultEvent.Action().HasListener(event, listener)
}

// 注册过滤器
// Add Filter
func AddFilter(event any, listener any, sort ...int) {
	useSort := DefaultSort
	if len(sort) > 0 {
		useSort = sort[0]
	}

	defaultEvent.Filter().Listen(event, listener, useSort)
}

// 触发过滤器
// Apply Filters
func ApplyFilters(event any, params ...any) any {
	return defaultEvent.Filter().Trigger(event, params...)
}

// 移除过滤器
// Remove Filter
func RemoveFilter(event string, listener any, sort ...int) bool {
	useSort := DefaultSort
	if len(sort) > 0 {
		useSort = sort[0]
	}

	return defaultEvent.Filter().RemoveListener(event, listener, useSort)
}

// 是否有过滤器
// Has Filter
func HasFilter(event string, listener any) bool {
	return defaultEvent.Filter().HasListener(event, listener)
}
