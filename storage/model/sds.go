package model

import "errors"

// Sds 模拟redis 数据结构，golang中对[]byte 已有足够优化，故不做其他处理
type Sds struct {
	buf []byte
}

// SdsCmp 对比两个sds是否相同
func SdsCmp(s1, s2 *Sds) bool {
	return false
}

//
func SdsNew(bs []byte) *Sds {
	s := new(Sds)
	s.buf = bs
	return s
}

func SdsEmpty() *Sds {
	bs := make([]byte, 0)
	return SdsNew(bs)
}

// SdsLen 字符串长度
func (s *Sds) SdsLen() int {
	return len(s.buf)
}

// SdsAvail 字符串中空闲长度
func (s *Sds) SdsAvail() int {
	return cap(s.buf) - len(s.buf)
}

// SdsDup 创建sds的副本
func (s *Sds) SdsDup() *Sds {
	return nil
}

// SdsClear 惰性删除
func (s *Sds) SdsClear() {
	bs := make([]byte, 0)
	s.buf = bs
}

// SdsCat 将给定的字节数组拼接到后面
func (s *Sds) SdsCat(bs []byte) {
	s.buf = append(s.buf, bs...)
}

// 将字节数组覆盖sds原有的字符串
func (s *Sds) SdsCpy(bs []byte) {
	s.buf = bs
}

// 用空字符将sds扩展至给定长度
func (s *Sds) SdsGrowZero() {

}

// 保留区间的数组，不在区间内的覆盖
func (s *Sds) SdsRange(start, end int) error {
	if start < 0 || end >= s.SdsLen() {
		return errors.New("array out of bounds")
	}
	s.buf = s.buf[start:end]
	return nil
}

// 从sds左右两端分别起初所有在bs字符串中出现过的字符
func (s *Sds) SdsTrim(bs []byte) {

}

func (s *Sds) ToString() string {
	if s == nil {
		return ""
	}
	return string(s.buf)
}
