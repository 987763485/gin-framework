/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  snowflake.go
 * @Time: 2020/9/17 10:06 上午
 */

package lib

import idworker "github.com/gitstliu/go-id-worker"

func GenerateGUID() (uint64, error) {
	var currWork = idworker.IdWorker{}
	er := currWork.InitIdWorker(1000, 1)
	if er != nil {
		return 0, er
	}
	GUID, err := currWork.NextId()
	if err != nil {
		return 0, err
	} else {
		return uint64(GUID), nil
	}
}
