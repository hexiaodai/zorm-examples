package zormexamples

import (
	"context"
	"errors"
	"log"
	"sync/atomic"
	"testing"
	"time"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

//TestCtx 事务传播。即事务由ctx传递下去，其他事务复用源事务
func TestCtx(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	// 事务a
	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		_, err := zorm.Delete(ctx, &models.DemoStruct{Id: "20210630163227149563000042432429"})
		if err != nil {
			return nil, err
		}

		// 事务b
		txB(ctx, "20210630163227149563000042432430", "20210630163227149563000042432431")

		return nil, err
	})
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}

func txB(ctx context.Context, argv ...interface{}) {

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		for _, parameter := range argv {
			log.Println(parameter.(string))
		}

		return nil, errors.New("测试事务传播") //模拟触发b事务失败，b回滚后，a也会回滚
	})
	if err != nil { //标记测试失败
		log.Printf("错误:%v", err)
	}
}

// TestTimeOut 设置事务执行超时
func TestTimeOut(t *testing.T) {
	var (
		pattern = 1
		t1, t2  = 0, 0
	)

	switch pattern {
	case 0:
		//正常模式
		t1, t2 = 10, 8
	case 1:
		//超时模式
		t1, t2 = 8, 10
	default:
		t1, t2 = 10, 8
	}

	ctxTimeOut, _ := context.WithTimeout(context.TODO(), time.Duration(time.Duration(t1)*time.Second))
	_, err := zorm.Transaction(ctxTimeOut, func(ctx context.Context) (interface{}, error) {
		var (
			stop int64
			done = make(chan int)
		)

		go func() {
			for i := 0; i < t2; i++ {
				if atomic.LoadInt64(&stop) == 1 {
					t.Log("收到停止作业的指令")
					return
				}

				time.Sleep(1 * time.Second)
				t.Logf("------sleep %d------", i)
			}

			done <- 1
		}()

		for {
			select {

			case <-done:
				t.Log("作业已正常处理完成!")
				return nil, nil

			case <-ctxTimeOut.Done():
				atomic.StoreInt64(&stop, 1)
				t.Log("发送停止作业的指令")
				return nil, errors.New("作业执行超时")
			}
		}
	})
	if err != nil { //标记测试失败
		log.Printf("错误:%v", err)
	}

}
