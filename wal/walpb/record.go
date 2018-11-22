// Copyright 2015 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package walpb

import "errors"

var (
	ErrCRCMismatch = errors.New("walpb: crc mismatch")
)

func (rec *Record) Validate(crc uint32) error {
	// CRC(循环冗余校验码): https://zh.wikipedia.org/wiki/%E5%BE%AA%E7%92%B0%E5%86%97%E9%A4%98%E6%A0%A1%E9%A9%97
	if rec.Crc == crc {
		return nil
	}
	rec.Reset()
	return ErrCRCMismatch
}
