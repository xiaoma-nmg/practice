package user

import (
	"errors"
	"log"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

// 静态设置返回值
func TestReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockUserRepository(ctrl)
	// 期望FindOne(1) 返回 "张三" 用户
	repo.EXPECT().FindOne(1).Return(&User{Name: "张三"}, nil)
	// 期望FindOne(2) 返回 "李四" 用户
	repo.EXPECT().FindOne(2).Return(&User{Name: "李四"}, nil)
	// 期望FindOne(3) 返回 找不到用户的错误
	repo.EXPECT().FindOne(3).Return(nil, errors.New("user not found"))

	// 验证结果
	log.Println(repo.FindOne(1))
	log.Println(repo.FindOne(2))
	log.Println(repo.FindOne(3))
	//log.Println(repo.FindOne(4)) // 没有设置4的返回值，却执行了调用，测试不通过
}

func TestTime(t *testing.T) {
	timeStr := time.Now().Format("2006-01-02")
	today, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	beginTimeNum := today.Unix()
	endTimeNum := beginTimeNum + 86400
	t.Log(beginTimeNum)
	t.Log(endTimeNum)
}
