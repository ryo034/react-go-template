package multi_factor

import "sort"

type MultiFactors interface {
	Size() int
	AsSlice() []MultiFactor
	IsEmpty() bool
	IsNotEmpty() bool
	Latest() *MultiFactor
}

type multiFactors struct {
	wrapped []MultiFactor
}

func NewMultiFactors(wrapped []MultiFactor) MultiFactors {
	return &multiFactors{wrapped}
}

func (mfs *multiFactors) IsEmpty() bool {
	return len(mfs.wrapped) == 0
}

func (mfs *multiFactors) IsNotEmpty() bool {
	return !mfs.IsEmpty()
}

func (mfs *multiFactors) Size() int {
	return len(mfs.wrapped)
}

func (mfs *multiFactors) AsSlice() []MultiFactor {
	return append(make([]MultiFactor, 0, mfs.Size()), mfs.wrapped...)
}

func (mfs *multiFactors) Latest() *MultiFactor {
	if mfs.IsEmpty() {
		return nil
	}
	sort.Slice(mfs.wrapped, func(i, j int) bool {
		return mfs.wrapped[i].EnrollmentTimestamp().ToTime().Before(mfs.wrapped[j].EnrollmentTimestamp().ToTime())
	})
	return &mfs.wrapped[len(mfs.wrapped)-1]
}
