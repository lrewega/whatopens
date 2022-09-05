package main

import (
	"net/url"
	"unicode/utf16"
	"unsafe"
)

// NSString conversion lifted from Gio (gioui.org).

/*
#cgo CFLAGS: -Werror -Wno-deprecated-declarations -fobjc-arc -x objective-c
#cgo LDFLAGS: -framework AppKit -framework Foundation
#import <AppKit/AppKit.h>
#import <Foundation/Foundation.h>

static NSUInteger nsstringLength(CFTypeRef cstr) {
	NSString *str = (__bridge NSString *)cstr;
	return [str length];
}

static void nsstringGetCharacters(CFTypeRef cstr, unichar *chars, NSUInteger loc, NSUInteger length) {
	NSString *str = (__bridge NSString *)cstr;
	[str getCharacters:chars range:NSMakeRange(loc, length)];
}

static CFTypeRef newNSString(unichar *chars, NSUInteger length) {
	@autoreleasepool {
		NSString *s = [NSString string];
		if (length > 0) {
			s = [NSString stringWithCharacters:chars length:length];
		}
		return CFBridgingRetain(s);
	}
}

static CFTypeRef urlForApplicationToOpenURL(CFTypeRef str) {
	@autoreleasepool {
		NSString *s = (__bridge NSString *)str;
		NSURL *nsurl = [NSURL URLWithString:s];
		nsurl = [[NSWorkspace sharedWorkspace] URLForApplicationToOpenURL:nsurl];
		if (nsurl == 0) {
			return 0;
		}
		return CFBridgingRetain([nsurl absoluteString]);
	}
}
*/
import "C"

// nsstringToString converts a NSString to a Go string.
func nsstringToString(str C.CFTypeRef) string {
	if str == 0 {
		return ""
	}
	n := C.nsstringLength(str)
	if n == 0 {
		return ""
	}
	chars := make([]uint16, n)
	C.nsstringGetCharacters(str, (*C.unichar)(unsafe.Pointer(&chars[0])), 0, n)
	utf8 := utf16.Decode(chars)
	return string(utf8)
}

// stringToNSString converts a Go string to a retained NSString.
func stringToNSString(str string) C.CFTypeRef {
	u16 := utf16.Encode([]rune(str))
	var chars *C.unichar
	if len(u16) > 0 {
		chars = (*C.unichar)(unsafe.Pointer(&u16[0]))
	}
	return C.newNSString(chars, C.NSUInteger(len(u16)))
}

// urlForApplicationToOpenURL wraps NSWorkspace.URLForApplicationToOpenURL
func urlForApplicationToOpenURL(u *url.URL) (*url.URL, error) {
	str := stringToNSString(u.String())
	defer C.CFRelease(str)
	str = C.urlForApplicationToOpenURL(str)
	if str != 0 {
		defer C.CFRelease(str)
	}
	appURL := nsstringToString(str)
	return url.Parse(appURL)
}
