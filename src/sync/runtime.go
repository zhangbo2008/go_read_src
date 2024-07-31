// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import "unsafe"

// defined in package runtime

// Semacquire waits until *s > 0 and then atomically decrements it.
// It is intended as a simple sleep primitive for use by the synchronization
// library and should not be used directly.
func runtime_Semacquire(s *uint32) //信号量表示的是可以被利用的资源数,所以我们一直等到大于0,才会进行获取. 获取锁之后信号量--.

// Semacquire(RW)Mutex(R) is like Semacquire, but for profiling contended
// Mutexes and RWMutexes.
// If lifo is true, queue waiter at the head of wait queue.
// skipframes is the number of frames to omit during tracing, counting from
// runtime_SemacquireMutex's caller.
// The different forms of this function just tell the runtime how to present
// the reason for waiting in a backtrace, and is used to compute some metrics.
// Otherwise they're functionally identical.
func runtime_SemacquireMutex(s *uint32, lifo bool, skipframes int)
func runtime_SemacquireRWMutexR(s *uint32, lifo bool, skipframes int)
func runtime_SemacquireRWMutex(s *uint32, lifo bool, skipframes int)

// Semrelease atomically increments *s and notifies a waiting goroutine
// if one is blocked in Semacquire.
// It is intended as a simple wakeup primitive for use by the synchronization
// library and should not be used directly.
// If handoff is true, pass count directly to the first waiter.
// skipframes is the number of frames to omit during tracing, counting from
// runtime_Semrelease's caller.  //增加s 信号量. 然后通知一个等待进程可以拿锁了.  // ps:如果信号量的值大于 0,则进程可以继续访问资源,并将信号量的值减 1;如果信号量的值等于 0,则进程会被阻塞,直到信号量的值变为正数。
func runtime_Semrelease(s *uint32, handoff bool, skipframes int)

// See runtime/sema.go for documentation.
func runtime_notifyListAdd(l *notifyList) uint32

// See runtime/sema.go for documentation.
func runtime_notifyListWait(l *notifyList, t uint32)

// See runtime/sema.go for documentation.
func runtime_notifyListNotifyAll(l *notifyList)

// See runtime/sema.go for documentation.
func runtime_notifyListNotifyOne(l *notifyList)

// Ensure that sync and runtime agree on size of notifyList.
func runtime_notifyListCheck(size uintptr)
func init() {
	var n notifyList
	runtime_notifyListCheck(unsafe.Sizeof(n))
}

// Active spinning runtime support.
// runtime_canSpin reports whether spinning makes sense at the moment.
func runtime_canSpin(i int) bool

// runtime_doSpin does active spinning.
func runtime_doSpin()

func runtime_nanotime() int64
