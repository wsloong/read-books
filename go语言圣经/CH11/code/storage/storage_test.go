package main

import (
	"strings"
	"testing"
)

func TestCheckQuota(t *testing.T) {
	// 使用伪邮件发送会替换掉真正的notifyUsers
	// 所以测试函数结束后要恢复
	saved := notifyUser
	defer func() { notifyUser = saved }()

	var notifiedUser, notifiedMsg string

	// 使用伪邮件发送替换真实的邮件发送
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.org"
	CheckQuota(user)

	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}

	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, want substring %q", notifiedMsg, wantSubstring)
	}
}
