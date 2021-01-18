// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

var (
	IID_NULL                      = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
	IID_IUnknown                  = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IDispatch                 = &GUID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
	IID_IConnectionPointContainer = &GUID{0xB196B284, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
	IID_IConnectionPoint          = &GUID{0xB196B286, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
)

var (
	CSIDL_DESKTOP                 int32 = 0x0000         // <desktop>
	CSIDL_INTERNET                int32 = 0x0001         // Internet Explorer (icon on desktop)
	CSIDL_PROGRAMS                int32 = 0x0002         // Start Menu\Programs
	CSIDL_CONTROLS                int32 = 0x0003         // My Computer\Control Panel
	CSIDL_PRINTERS                int32 = 0x0004         // My Computer\Printers
	CSIDL_PERSONAL                int32 = 0x0005         // My Documents
	CSIDL_FAVORITES               int32 = 0x0006         // <user name>\Favorites
	CSIDL_STARTUP                 int32 = 0x0007         // Start Menu\Programs\Startup
	CSIDL_RECENT                  int32 = 0x0008         // <user name>\Recent
	CSIDL_SENDTO                  int32 = 0x0009         // <user name>\SendTo
	CSIDL_BITBUCKET               int32 = 0x000a         // <desktop>\Recycle Bin
	CSIDL_STARTMENU               int32 = 0x000b         // <user name>\Start Menu
	CSIDL_MYDOCUMENTS             int32 = CSIDL_PERSONAL //  Personal was just a silly name for My Documents
	CSIDL_MYMUSIC                 int32 = 0x000d         // "My Music" folder
	CSIDL_MYVIDEO                 int32 = 0x000e         // "My Videos" folder
	CSIDL_DESKTOPDIRECTORY        int32 = 0x0010         // <user name>\Desktop
	CSIDL_DRIVES                  int32 = 0x0011         // My Computer
	CSIDL_NETWORK                 int32 = 0x0012         // Network Neighborhood (My Network Places)
	CSIDL_NETHOOD                 int32 = 0x0013         // <user name>\nethood
	CSIDL_FONTS                   int32 = 0x0014         // windows\fonts
	CSIDL_TEMPLATES               int32 = 0x0015
	CSIDL_COMMON_STARTMENU        int32 = 0x0016 // All Users\Start Menu
	CSIDL_COMMON_PROGRAMS         int32 = 0x0017 // All Users\Start Menu\Programs
	CSIDL_COMMON_STARTUP          int32 = 0x0018 // All Users\Startup
	CSIDL_COMMON_DESKTOPDIRECTORY int32 = 0x0019 // All Users\Desktop
	CSIDL_APPDATA                 int32 = 0x001a // <user name>\Application Data
	CSIDL_PRINTHOOD               int32 = 0x001b // <user name>\PrintHood
	CSIDL_LOCAL_APPDATA           int32 = 0x001c // <user name>\Local Settings\Applicaiton Data (non roaming)
	CSIDL_ALTSTARTUP              int32 = 0x001d // non localized startup
	CSIDL_COMMON_ALTSTARTUP       int32 = 0x001e // non localized common startup
	CSIDL_COMMON_FAVORITES        int32 = 0x001f
)

