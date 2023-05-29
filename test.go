package go_test

import (
	"fmt"
	"sync"
	"time"
)

type Test struct {
	wait sync.WaitGroup
	num int
	fun func()
	startTime time.Time
	endTime time.Time
}
type Result struct {
	num int
	StartTime time.Time
	EndTime time.Time
	SumTime time.Duration
	AvgTime time.Duration
}

func NewResult(num int,startTime,endTime time.Time) *Result  {
	var res = new(Result)
	res.num = num
	res.StartTime = startTime
	res.EndTime = endTime
	res.SumTime = endTime.Sub(startTime)
	res.AvgTime = time.Duration(res.SumTime.Nanoseconds()/int64(res.num))
	return res
}

func (r *Result) String() string {
	str := fmt.Sprintf("测试结果：\n运行次数：%v 总计耗时：%v 平均耗时：%v\n开始时间：%v 结束时间：%v",r.num,r.SumTime,r.AvgTime,r.StartTime,r.EndTime)
	return str
}

func (r *Result) Debug()  {
	fmt.Println(r.String())
}

func NewTest(num int,f func(),isSync ...bool) *Test {
	return &Test{
		wait:      sync.WaitGroup{},
		num:       num,
		fun: f,
	}
}

func (t *Test) Sync() *Result  {
	t.startTime = time.Now()
	for i:=0;i<t.num;i++{
		t.fun()
	}
	t.endTime = time.Now()

	return NewResult(t.num,t.startTime,t.endTime)
}

func (t *Test) Async() *Result {
	t.wait.Add(t.num)
	t.startTime = time.Now()
	for i:=0;i<t.num;i++{
		go func() {
			t.fun()
			t.wait.Done()
		}()
	}
	t.wait.Wait()
	t.endTime = time.Now()

	return NewResult(t.num,t.startTime,t.endTime)
}