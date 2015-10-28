package executor

import (
	"reflect"
	"testing"
)

func TestNew_リモートホストのosがWindowsの場合はWindowsの構造体が作成(t *testing.T) {
	exec := New("host", "user", "password", "windows")
	if reflect.TypeOf(exec).String() != "*executor.Windows" {
		t.Errorf("Windows構造体が設定されていない。 [%s]", reflect.TypeOf(exec).String())
	}

	win, result := exec.(*Windows)
	if !result {
		t.Errorf("型アサーションに失敗")
	}

	if win.HostName != "host" {
		t.Errorf("期待するホスト名が設定されていない。[%s]", win.HostName)
	}

	if win.UserName != "user" {
		t.Errorf("期待するユーザ名が設定されていない。[%s]", win.UserName)
	}

	if win.Password != "password" {
		t.Errorf("期待するパスワードが設定されていない。[%s]", win.Password)
	}

}

func TestNew_リモートホストのosがLinuxの場合はLinuxの構造体が作成(t *testing.T) {

}
