// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package executor

func New(host, user, pass, os string) Executor {
	// TODO: Decide struct type by os, And create struct object.
	return nil
}
