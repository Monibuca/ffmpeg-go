package libavutil

import (
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
)

/*
 * An implementation of the CAST128 algorithm as mentioned in RFC2144
 * Copyright (c) 2014 Supraja Meedinti
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVUTIL_CAST5_H
//#define AVUTIL_CAST5_H
//
//#include <stdint.h>

/**
 * @file
 * @brief Public header for libavutil CAST5 algorithm
 * @defgroup lavu_cast5 CAST5
 * @ingroup lavu_crypto
 * @{
 */

//extern const int av_cast5_size;

// struct AVCAST5;
type AVCAST5 struct {
}

/**
 * Allocate an AVCAST5 context
 * To free the struct: av_free(ptr)
 */
//struct AVCAST5 *av_cast5_alloc(void);
func AvCast5Alloc() (res *AVCAST5) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_cast5_alloc").Call()
	res = (*AVCAST5)(unsafe.Pointer(t))
	return
}

/**
 * Initialize an AVCAST5 context.
 *
 * @param ctx an AVCAST5 context
 * @param key a key of 5,6,...16 bytes used for encryption/decryption
 * @param key_bits number of keybits: possible are 40,48,...,128
 * @return 0 on success, less than 0 on failure
 */
//int av_cast5_init(struct AVCAST5 *ctx, const uint8_t *key, int key_bits);
func (ctx *AVCAST5) AvCast5Init(key *ffcommon.FUint8T, key_bits ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_cast5_init").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(key)),
		uintptr(key_bits),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context, ECB mode only
 *
 * @param ctx an AVCAST5 context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_cast5_crypt(struct AVCAST5 *ctx, uint8_t *dst, const uint8_t *src, int count, int decrypt);
func (ctx *AVCAST5) AvCast5Crypt(dst, src *ffcommon.FUint8T, count, decrypt ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_cast5_crypt").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
		uintptr(decrypt),
	)
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context
 *
 * @param ctx an AVCAST5 context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, NULL for ECB mode
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_cast5_crypt2(struct AVCAST5 *ctx, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
func (ctx *AVCAST5) AvCast5Crypt2(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_cast5_crypt2").Call(
		uintptr(unsafe.Pointer(ctx)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(count),
		uintptr(unsafe.Pointer(iv)),
		uintptr(decrypt),
	)
}

/**
 * @}
 */
//#endif /* AVUTIL_CAST5_H */