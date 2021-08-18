package task

//
//import (
//	"time"
//
//	"pmo-test4.yz-intelligence.com/kit/component/gins/logger"
//	"pmo-test4.yz-intelligence.com/kit/component/task"
//)
//
//func init() {
//	demoTask := &demo{}
//	err := Manager.Add(demoTask.new())
//	if err != nil {
//		panic(err)
//	}
//}
//
//type demo struct {
//	interval time.Duration
//}
//
//func (c *demo) new() *task.Task {
//
//	// 间隔时间
//	c.interval = 10 * time.Second
//
//	return &task.Task{
//		Name:     "示例任务",
//		IsSingle: true,
//		OnInterval: func(t *task.Task) time.Duration {
//			return c.interval
//		},
//		OnRun:   c.onRun,
//		OnPanic: c.onPanic,
//	}
//}
//
//// onPanic 异常处理
//func (c *demo) onPanic(t *task.Task, err error) {
//	logger.Errorf("异常：%s", err)
//}
//
//// onRun 执行任务
//func (c *demo) onRun(t *task.Task) {
//
//	logger.Infof("任务：%s 执行了", t.Name)
//
//	return
//}
