package gst

/*
#include "gst.go.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// MpegtsSection is a go representation of a GstMpegtsSection
type MpegtsSection struct {
	section *C.GstMpegtsSection
}

// FromGstMpegtsSectionUnsafeFull wraps the given unsafe.Pointer in a MpegtsSection. No ref is taken
// and a finalizer is placed on the resulting object.
func FromGstMpegtsSectionUnsafeFull(section unsafe.Pointer) *MpegtsSection {
	gosection := ToGstMpegtsSection(section)
	runtime.SetFinalizer(gosection, (*MpegtsSection).Unref)
	return gosection
}

// ToGstMpegtsSection converts the given pointer into a MpegtsSection without affecting the ref count or
// placing finalizers.
func ToGstMpegtsSection(section unsafe.Pointer) *MpegtsSection {
	return wrapMpagtsSection((*C.GstMpegtsSection)(section))
}

// Instance returns the underlying GstMpegtsSection instance.
func (m *MpegtsSection) Instance() *C.GstMpegtsSection {
	return C.toGstMpegtsSection(unsafe.Pointer(m.section))
}

// Unref will call `gst_mpegts_section_unref` on the underlying GstMpegtsSection, freeing it from memory.
func (m *MpegtsSection) Unref() { C.mpegtsSectionUnref(m.Instance()) }

// Ref will increase the ref count on this MpegtsSection. This increases the total amount of times
// Unref needs to be called before the object is freed from memory. It returns the underlying
// MpegtsSection object for convenience.
func (m *MpegtsSection) Ref() *MpegtsSection {
	C.mpegtsSectionRef(m.Instance())
	return m
}
