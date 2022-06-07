package libavutil

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"unsafe"
)

/*
 * Copyright (c) 2014 Tim Walker <tdskywalker@gmail.com>
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

//#ifndef AVUTIL_DOWNMIX_INFO_H
//#define AVUTIL_DOWNMIX_INFO_H
//
//#include "frame.h"

/**
 * @file
 * audio downmix medatata
 */

/**
 * @addtogroup lavu_audio
 * @{
 */

/**
 * @defgroup downmix_info Audio downmix metadata
 * @{
 */

/**
 * Possible downmix types.
 */
type AVDownmixType = int32

const (
	AV_DOWNMIX_TYPE_UNKNOWN = iota /**< Not indicated. */
	AV_DOWNMIX_TYPE_LORO           /**< Lo/Ro 2-channel downmix (Stereo). */
	AV_DOWNMIX_TYPE_LTRT           /**< Lt/Rt 2-channel downmix, Dolby Surround compatible. */
	AV_DOWNMIX_TYPE_DPLII          /**< Lt/Rt 2-channel downmix, Dolby Pro Logic II compatible. */
	AV_DOWNMIX_TYPE_NB             /**< Number of downmix types. Not part of ABI. */
)

/**
 * This structure describes optional metadata relevant to a downmix procedure.
 *
 * All fields are set by the decoder to the value indicated in the audio
 * bitstream (if present), or to a "sane" default otherwise.
 */
type AVDownmixInfo struct {

	/**
	 * Type of downmix preferred by the mastering engineer.
	 */
	PreferredDownmixType AVDownmixType

	/**
	 * Absolute scale factor representing the nominal level of the center
	 * channel during a regular downmix.
	 */
	CenterMixLevel ffcommon.FDouble

	/**
	 * Absolute scale factor representing the nominal level of the center
	 * channel during an Lt/Rt compatible downmix.
	 */
	CenterMixLevelLtrt ffcommon.FDouble

	/**
	 * Absolute scale factor representing the nominal level of the surround
	 * channels during a regular downmix.
	 */
	SurroundMixLevel ffcommon.FDouble

	/**
	 * Absolute scale factor representing the nominal level of the surround
	 * channels during an Lt/Rt compatible downmix.
	 */
	SurroundMixLevelLtrt ffcommon.FDouble

	/**
	 * Absolute scale factor representing the level at which the LFE data is
	 * mixed into L/R channels during downmixing.
	 */
	LfeMixLevel ffcommon.FDouble
}

/**
 * Get a frame's AV_FRAME_DATA_DOWNMIX_INFO side data for editing.
 *
 * If the side data is absent, it is created and added to the frame.
 *
 * @param frame the frame for which the side data is to be obtained or created
 *
 * @return the AVDownmixInfo structure to be edited by the caller, or NULL if
 *         the structure cannot be allocated.
 */
//AVDownmixInfo *av_downmix_info_update_side_data(AVFrame *frame);
func (frame *AVFrame) AvDownmixInfoUpdateSideData() (res *AVDownmixInfo) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_downmix_info_update_side_data").Call(
		uintptr(unsafe.Pointer(frame)),
	)
	if t == 0 {

	}
	res = (*AVDownmixInfo)(unsafe.Pointer(t))
	return
}

/**
 * @}
 */

/**
 * @}
 */

//#endif /* AVUTIL_DOWNMIX_INFO_H */
