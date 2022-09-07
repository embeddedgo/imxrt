// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package dma provides interface to the eDMA controller. The interface is based
// on two main types: Controller and Channel.
//
// Controller represents an instance of eDMA engine together with the
// corresponding DMAMUX. Each controller provides 32 channels.
//
// Channel represents a DMA channel. You can select a specific channel using the
// Controller.Channel method but the prefered way to obtain a channel is to use
// Controller.AllocChannel which arbitrarily allocate an unused one.
//
// When this package is imported it alters the default configuration of all
// available controllers to use round robin arbitration and to halt on error.
// The default fixed priority arbitration with its requirement of unique channel
// prioritiesis does not work well with the Controller.AllocChannel method.
// Additionally, there is a problem with canceling a transfer in fixed priority
// mode if channel preemption is enabled.
package dma
