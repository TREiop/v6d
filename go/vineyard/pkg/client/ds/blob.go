/** Copyright 2020-2021 Alibaba Group Holding Limited.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vineyard

import (
	"errors"
	"fmt"
	"github.com/v6d-io/v6d/go/vineyard/pkg/common"
	//x	"github.com/apache/arrow/go/arrow"
)

type Blob struct {
	id int
	size int
	buffer []byte
}

func (b *Blob) Size() int {
	return b.size
}

func (b *Blob) Data() ([]byte, error) {
	if b.size > 0 && len(b.buffer) == 0 {
		return nil, errors.New(fmt.Sprintf("The object might be a (partially) remote object " +
		                    	"and the payload data is not locally available: %d", b.id))
	}
	return b.buffer, nil
}

type BufferSet struct {
	//buffer arrow.
}

func (b *BufferSet) EmplaceBuffer(id common.ObjectID) {

}

type BlobWriter struct {

}




