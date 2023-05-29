// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package merge

import (
	"bytes"
)

type TextScanner interface {
	Scan() bool
	Text() string
	Bytes() []byte
}

type Merged struct {
	src     []TextScanner
	cur     [][]byte
	ord     []int
	compare func(a, b []byte) int
}

func New(in ...TextScanner) TextScanner {
	return &Merged{src: in, compare: bytes.Compare}
}

func NewWithCompare(compare func(a, b []byte) int, in ...TextScanner) TextScanner {
	return &Merged{src: in, compare: compare}
}

func (m *Merged) Scan() bool {
	if len(m.src) == 0 {
		return false
	}

	// Build first ordered list
	if len(m.ord) == 0 {
		m.cur = make([][]byte, len(m.src))
		for i := range m.src {
			if !m.src[i].Scan() {
				continue
			}

			if len(m.ord) == 0 {
				// First pass on reader
				m.cur[i] = m.src[i].Bytes()
				m.ord = append(m.ord, i)
				continue
			}

			for {
				if t := m.insert(i); t < 0 {
					break
				} else {
					i = t
				}
				if !m.src[i].Scan() {
					break
				}
			}
		}
		return len(m.ord) > 0
	}

	i := m.ord[0]
	m.ord = m.ord[1:]
	if !m.src[i].Scan() {
		// If we hit the end of the first file, just eliminate it
		return len(m.ord) > 0
	}

	for {
		if t := m.insert(i); t < 0 {
			break
		} else {
			i = t
		}
		if !m.src[i].Scan() {
			break
		}
	}
	return len(m.ord) > 0
}

func (m *Merged) insert(i int) int {
	next := m.src[i].Bytes()
	for j := 0; j < len(m.ord); j++ {
		if cmp := m.compare(m.cur[m.ord[j]], next); cmp == 0 {
			if i < m.ord[j] {
				m.cur[i] = next
				i, m.ord[j] = m.ord[j], i
			}
			return i
		} else if cmp > 0 {
			m.ord = append(m.ord[:j], append([]int{i}, m.ord[j:]...)...)
			m.cur[i] = next
			return -1
		}
	}
	m.ord = append(m.ord, i)
	m.cur[i] = next
	return -1
}

func (m *Merged) Text() string {
	return string(m.Bytes())
}
func (m *Merged) Bytes() []byte {
	if len(m.ord) == 0 {
		return nil
	}
	return m.cur[m.ord[0]]
}
