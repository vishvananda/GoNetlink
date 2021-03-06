package route

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

type Origin byte

const (
	RTPROT_UNSPEC Origin = iota
	RTPROT_REDIRECT
	RTPROT_KERNEL
	RTPROT_BOOT
	RTPROT_STATIC
)

var OriginStrings = map[Origin]string{
	RTPROT_UNSPEC:   "RTPROT_UNSPEC",
	RTPROT_REDIRECT: "RTPROT_REDIRECT",
	RTPROT_KERNEL:   "RTPROT_KERNEL",
	RTPROT_BOOT:     "RTPROT_BOOT",
	RTPROT_STATIC:   "RTPROT_STATIC",
}

func (self Origin) String() string {
	return OriginStrings[self]
}
