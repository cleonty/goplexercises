package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	saved := notifyUser
	defer func() { notifyUser = saved }()
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	const user = "joe@example.com"
	checkQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notify user not called")
	}
	if notifiedUser != user {
		t.Errorf("Уведомлен (%s) вместо %s", notifiedUser, user)
	}
	const wantSubstring = "98% вашей квоты"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("неожиданное уведомление <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}

}
