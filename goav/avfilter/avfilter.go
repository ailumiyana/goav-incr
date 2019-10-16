// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package avfilter contains methods that deal with ffmpeg filters
//filters in the same linear chain are separated by commas, and distinct linear chains of filters are separated by semicolons.
//FFmpeg is enabled through the "C" libavfilter library
package avfilter

/*
	#cgo pkg-config: libavfilter
	#include <libavfilter/avfilter.h>
	#include <libavfilter/buffersrc.h>
	#include <libavfilter/buffersink.h>
	#include <libavutil/avutil.h>
*/
import "C"
import (
	"unsafe"
	"github.com/ailumiyana/goav-incr/goav/avutil"
)

type (
	Filter     			C.struct_AVFilter
	Context    			C.struct_AVFilterContext
	Link       			C.struct_AVFilterLink
	Graph      			C.struct_AVFilterGraph
	Input      			C.struct_AVFilterInOut
	Pad        			C.struct_AVFilterPad
	Dictionary 			C.struct_AVDictionary
	BufferSrcParameters	C.struct_AVBufferSrcParameters
	Class      			C.struct_AVClass
	Frame      			C.struct_AVFrame
	MediaType 		 	C.enum_AVMediaType
)

//Return the LIBAvFILTER_VERSION_INT constant.
func AvfilterVersion() uint {
	return uint(C.avfilter_version())
}

//Return the libavfilter build-time configuration.
func AvfilterConfiguration() string {
	return C.GoString(C.avfilter_configuration())
}

//Return the libavfilter license.
func AvfilterLicense() string {
	return C.GoString(C.avfilter_license())
}

//Get the number of elements in a NULL-terminated array of Pads (e.g.
func AvfilterPadCount(p *Pad) int {
	return int(C.avfilter_pad_count((*C.struct_AVFilterPad)(p)))
}

//Get the name of an Pad.
func AvfilterPadGetName(p *Pad, pi int) string {
	return C.GoString(C.avfilter_pad_get_name((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Get the type of an Pad.
func AvfilterPadGetType(p *Pad, pi int) MediaType {
	return (MediaType)(C.avfilter_pad_get_type((*C.struct_AVFilterPad)(p), C.int(pi)))
}

//Link two filters together.
func AvfilterLink(s *Context, sp uint, d *Context, dp uint) int {
	return int(C.avfilter_link((*C.struct_AVFilterContext)(s), C.uint(sp), (*C.struct_AVFilterContext)(d), C.uint(dp)))
}

//Free the link in *link, and set its pointer to NULL.
func AvfilterLinkFree(l **Link) {
	C.avfilter_link_free((**C.struct_AVFilterLink)(unsafe.Pointer(l)))
}

//Get the number of channels of a link.
func AvfilterLinkGetChannels(l *Link) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

//Set the closed field of a link.
func AvfilterLinkSetClosed(l *Link, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

//Negotiate the media format, dimensions, etc of all inputs to a filter.
func AvfilterConfigLinks(f *Context) int {
	return int(C.avfilter_config_links((*C.struct_AVFilterContext)(f)))
}

//Make the filter instance process a command.
func AvfilterProcessCommand(f *Context, cmd, arg, res string, l, fl int) int {
	return int(C.avfilter_process_command((*C.struct_AVFilterContext)(f), C.CString(cmd), C.CString(arg), C.CString(res), C.int(l), C.int(fl)))
}

//Initialize the filter system.
func AvfilterRegisterAll() {
	C.avfilter_register_all()
}

//Initialize a filter with the supplied parameters.
func (ctx *Context) AvfilterInitStr(args string) int {
	return int(C.avfilter_init_str((*C.struct_AVFilterContext)(ctx), C.CString(args)))
}

//Initialize a filter with the supplied dictionary of options.
func (ctx *Context) AvfilterInitDict(o **Dictionary) int {
	return int(C.avfilter_init_dict((*C.struct_AVFilterContext)(ctx), (**C.struct_AVDictionary)(unsafe.Pointer(o))))
}

//Free a filter context.
func (ctx *Context) AvfilterFree() {
	C.avfilter_free((*C.struct_AVFilterContext)(ctx))
}

//Insert a filter in the middle of an existing link.
func AvfilterInsertFilter(l *Link, f *Context, fsi, fdi uint) int {
	return int(C.avfilter_insert_filter((*C.struct_AVFilterLink)(l), (*C.struct_AVFilterContext)(f), C.uint(fsi), C.uint(fdi)))
}

//avfilter_get_class
func AvfilterGetClass() *Class {
	return (*Class)(C.avfilter_get_class())
}

//Allocate a single Input entry.
func AvfilterInoutAlloc() *Input {
	return (*Input)(C.avfilter_inout_alloc())
}

//Free the supplied list of Input and set *inout to NULL.
func AvfilterInoutFree(i *Input) {
	C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(i)))
}

func (i *Input)FilterContext() *Context {
	return (*Context)(i.filter_ctx)
}

func (i *Input)Next() *Input {
	return (*Input)(i.next)
}

func (i *Input)PadIdx() uint {
	return (uint)(i.pad_idx)
}

func AvBufferSinkGetFrame(c *Context, f *Frame) int {
	return (int)(C.av_buffersink_get_frame((*C.struct_AVFilterContext)(c), (*C.struct_AVFrame)(f)))
}

func AvBufferSinkGetFrameFlags(c *Context, f *Frame, flag int) int {
	return (int)(C.av_buffersink_get_frame_flags((*C.struct_AVFilterContext)(c), (*C.struct_AVFrame)(f), C.int(flag)))
}

// keep ref
func AvBuffersrcWriteFrame(c *Context, f *Frame) int {
	return (int)(C.av_buffersrc_write_frame((*C.struct_AVFilterContext)(c), (*C.struct_AVFrame)(f)))
}

// copy
func AvBuffersrcAddFrame(c *Context, f *Frame) int {
	return (int)(C.av_buffersrc_add_frame((*C.struct_AVFilterContext)(c), (*C.struct_AVFrame)(f)))
}

func (ctx *Context) SetHWDeviceCtx(hw_device_ctx *avutil.BufferRef) {
	ctx.hw_device_ctx = (*C.struct_AVBufferRef)(unsafe.Pointer(hw_device_ctx))
}

func (ctx *Context) HWDeviceCtx() *avutil.BufferRef {
	return (*avutil.BufferRef)(unsafe.Pointer(ctx.hw_device_ctx))
}

func (ctx *Context) AvBuffersrcParametersSet(par *BufferSrcParameters) int {
	return (int)(C.av_buffersrc_parameters_set((*C.struct_AVFilterContext)(ctx), (*C.struct_AVBufferSrcParameters)(par)))
}

func AvBuffersrcParametersAlloc() *BufferSrcParameters{
	return (*BufferSrcParameters)(C.av_buffersrc_parameters_alloc())
}

func (p *BufferSrcParameters)SetHwFramesCtx(hw_frames_ctx *avutil.BufferRef) {
	p.hw_frames_ctx = (*C.struct_AVBufferRef)(unsafe.Pointer(hw_frames_ctx))
}

func (ctx *Context) AvBuffersinkGetHwFramesCtx() *avutil.BufferRef {
	return (*avutil.BufferRef)(unsafe.Pointer(C.av_buffersink_get_hw_frames_ctx((*C.struct_AVFilterContext)(ctx))))
}
