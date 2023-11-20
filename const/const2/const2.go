package main

type ForwardStatus int

const (
	Initial  ForwardStatus = iota // 入库
	Normal                        // 审核通过
	UserDel                       // 用户删除
	AuditDel                      // 审核删除
)
