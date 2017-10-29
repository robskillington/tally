// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tally

import (
	"sync"
	"sync/atomic"
)

const (
	offset64 = 14695981039346656037
	prime64  = 1099511628211
)

func hashNew() uint64 {
	return offset64
}

func hashAdd(h uint64, s string) uint64 {
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime64
	}
	return h
}

func hashAddByte(h uint64, b byte) uint64 {
	h ^= uint64(b)
	h *= prime64
	return h
}

type CounterVector struct {
	sync.RWMutex
	tags      []string
	help      string
	instances map[uint64]*resolvedCounterVector
}

type resolvedCounterVector struct {
	v int64
}

type VectorOptions struct {
	Help string
}

func NewCounterVector(opts VectorOptions, tags ...string) *CounterVector {
	return &CounterVector{
		help:      opts.Help,
		tags:      tags,
		instances: make(map[uint64]*resolvedCounterVector),
	}
}

func (v *CounterVector) With(tagName, tagValue string) CounterVectorQuery {
	hash := hashNew()
	valid := false
	if v.tags[0] == tagName {
		hash = hashAdd(hash, tagName)
		hash = hashAddByte(hash, byte(255))
		valid = true
	}
	return CounterVectorQuery{
		v:       v,
		hash:    hash,
		valid:   valid,
		idxNext: 1,
	}
}

func (v *CounterVector) value(hash uint64) int64 {
	v.RLock()
	r, ok := v.instances[hash]
	v.RUnlock()
	if !ok {
		return 0
	}
	return atomic.LoadInt64(&r.v)
}

func (v *CounterVector) inc(hash uint64, delta int64) {
	v.RLock()
	r, ok := v.instances[hash]
	v.RUnlock()
	if !ok {
		v.Lock()
		r, ok = v.instances[hash]
		if !ok {
			r = new(resolvedCounterVector)
			v.instances[hash] = r
		}
		v.Unlock()
	}
	atomic.AddInt64(&r.v, delta)
}

type CounterVectorQuery struct {
	v       *CounterVector
	hash    uint64
	valid   bool
	idxNext int
}

func (q CounterVectorQuery) Valid() bool {
	return q.valid
}

func (q CounterVectorQuery) With(tagName, tagValue string) CounterVectorQuery {
	if !q.valid {
		return q
	}
	if q.idxNext >= len(q.v.tags) {
		r := q
		r.valid = false
		return q
	}
	if q.v.tags[q.idxNext] != tagName {
		r := q
		r.valid = false
		return q
	}
	r := q
	r.idxNext++
	r.hash = hashAdd(r.hash, tagValue)
	r.hash = hashAddByte(r.hash, byte(255))
	return r
}

func (q CounterVectorQuery) Resolved() bool {
	return q.idxNext == len(q.v.tags)
}

func (q CounterVectorQuery) Value() int64 {
	if !q.Resolved() {
		return 0
	}
	return q.v.value(q.hash)
}

func (q CounterVectorQuery) Inc(delta int64) {
	if !q.Resolved() {
		return
	}
	q.v.inc(q.hash, delta)
}

type Vectors struct {
	sync.Mutex
	counters []*CounterVector
}

func (v *Vectors) RegisterCounter(c *CounterVector) {
	v.Lock()
	v.counters = append(v.counters, c)
	v.Unlock()
}
