package columns_test

import (
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	if unsafe.Pointer(&s) != unsafe.Pointer(&s[0]) {
		t.Errorf("&s=%p != &s[0]=%p", &s, &s[0]) // fail: pointer to slice points to slice header
	}
}

func TestArray(t *testing.T) {
	a := [7]int{0, 1, 2, 3, 4, 5, 6}
	if unsafe.Pointer(&a) != unsafe.Pointer(&a[0]) {
		t.Errorf("&a=%p != &a[0]=%p", &a, &a[0]) // ok: pointer to array points to array element
	}
}

func TestSliceOfSlice(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	ss := s[1:4]
	if unsafe.Pointer(&s[1]) != unsafe.Pointer(&ss[0]) {
		t.Errorf("&s[1]=%p != &ss[0]=%p", &s[1], &ss[0]) // ok: slice of slice shares underlying array
	}
}

func TestSliceToArray(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	a := [7]int(s)
	if unsafe.Pointer(&s[0]) != unsafe.Pointer(&a[0]) {
		t.Errorf("&s[0]=%p != &a[0]=%p", &s[0], &a[0]) // fail: conversion from slice to array copies elements
	}
}

func TestArrayToSlice(t *testing.T) {
	a := [7]int{0, 1, 2, 3, 4, 5, 6}
	s := []int(a[:])
	if unsafe.Pointer(&a[0]) != unsafe.Pointer(&s[0]) {
		t.Errorf("&a[0]=%p != &s[0]=%p", &a[0], &s[0]) // ok: conversion from array to slice reuses array
	}
}

func TestSliceToArrayPtr(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	p := (*[5]int)(s[2:])
	if unsafe.Pointer(&s[2]) != unsafe.Pointer(&p[0]) {
		t.Errorf("&a[2]=%p != &p[0]=%p", &s[2], &p[0]) // ok: conversion from slice to array pointer reuses array
	}
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("panic: %v", err) // fail: conversion from slice to array pointer panics if slice length is too short
		}
	}()
	_ = (*[6]int)(s[2:]) // panic: runtime error: slice bounds out of range
}
