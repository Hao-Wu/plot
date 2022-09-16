// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !minimal
// +build !minimal

package plot // import "github.com/Hao-Wu/plot"

import (
	_ "github.com/Hao-Wu/plot/vg/vgeps"
	_ "github.com/Hao-Wu/plot/vg/vgimg"
	_ "github.com/Hao-Wu/plot/vg/vgpdf"
	_ "github.com/Hao-Wu/plot/vg/vgsvg"
	_ "github.com/Hao-Wu/plot/vg/vgtex"
)
