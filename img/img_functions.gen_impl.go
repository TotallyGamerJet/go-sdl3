//go:build unix || windows

package img

import (
	sdl "github.com/Zyko0/go-sdl3"
	puregogen "github.com/Zyko0/purego-gen"
	purego "github.com/ebitengine/purego"
	"unsafe"
)

// File generated by github.com/Zyko0/purego-gen. DO NOT EDIT.

var (
	// Library handles
	_hnd_img uintptr
	// Symbols
	// img
	_addr_IMG_Version                  uintptr
	_addr_IMG_LoadTyped_IO             uintptr
	_addr_IMG_Load                     uintptr
	_addr_IMG_Load_IO                  uintptr
	_addr_IMG_LoadTexture              uintptr
	_addr_IMG_LoadTexture_IO           uintptr
	_addr_IMG_LoadTextureTyped_IO      uintptr
	_addr_IMG_LoadAVIF_IO              uintptr
	_addr_IMG_LoadICO_IO               uintptr
	_addr_IMG_LoadCUR_IO               uintptr
	_addr_IMG_LoadBMP_IO               uintptr
	_addr_IMG_LoadGIF_IO               uintptr
	_addr_IMG_LoadJPG_IO               uintptr
	_addr_IMG_LoadJXL_IO               uintptr
	_addr_IMG_LoadLBM_IO               uintptr
	_addr_IMG_LoadPCX_IO               uintptr
	_addr_IMG_LoadPNG_IO               uintptr
	_addr_IMG_LoadPNM_IO               uintptr
	_addr_IMG_LoadSVG_IO               uintptr
	_addr_IMG_LoadQOI_IO               uintptr
	_addr_IMG_LoadTGA_IO               uintptr
	_addr_IMG_LoadTIF_IO               uintptr
	_addr_IMG_LoadXCF_IO               uintptr
	_addr_IMG_LoadXPM_IO               uintptr
	_addr_IMG_LoadXV_IO                uintptr
	_addr_IMG_LoadWEBP_IO              uintptr
	_addr_IMG_LoadSizedSVG_IO          uintptr
	_addr_IMG_ReadXPMFromArray         uintptr
	_addr_IMG_ReadXPMFromArrayToRGB888 uintptr
	_addr_IMG_SaveAVIF                 uintptr
	_addr_IMG_SaveAVIF_IO              uintptr
	_addr_IMG_SavePNG                  uintptr
	_addr_IMG_SavePNG_IO               uintptr
	_addr_IMG_SaveJPG                  uintptr
	_addr_IMG_SaveJPG_IO               uintptr
	_addr_IMG_LoadAnimation            uintptr
	_addr_IMG_LoadAnimation_IO         uintptr
	_addr_IMG_LoadAnimationTyped_IO    uintptr
	_addr_IMG_FreeAnimation            uintptr
	_addr_IMG_LoadGIFAnimation_IO      uintptr
	_addr_IMG_LoadWEBPAnimation_IO     uintptr
)

func initialize() {
	var err error

	// Symbols img
	_addr_IMG_Version, err = puregogen.OpenSymbol(_hnd_img, "IMG_Version")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_Version")
	}
	_addr_IMG_LoadTyped_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTyped_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTyped_IO")
	}
	_addr_IMG_Load, err = puregogen.OpenSymbol(_hnd_img, "IMG_Load")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_Load")
	}
	_addr_IMG_Load_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_Load_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_Load_IO")
	}
	_addr_IMG_LoadTexture, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTexture")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTexture")
	}
	_addr_IMG_LoadTexture_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTexture_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTexture_IO")
	}
	_addr_IMG_LoadTextureTyped_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTextureTyped_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTextureTyped_IO")
	}
	_addr_IMG_LoadAVIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAVIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAVIF_IO")
	}
	_addr_IMG_LoadICO_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadICO_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadICO_IO")
	}
	_addr_IMG_LoadCUR_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadCUR_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadCUR_IO")
	}
	_addr_IMG_LoadBMP_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadBMP_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadBMP_IO")
	}
	_addr_IMG_LoadGIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadGIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadGIF_IO")
	}
	_addr_IMG_LoadJPG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadJPG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadJPG_IO")
	}
	_addr_IMG_LoadJXL_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadJXL_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadJXL_IO")
	}
	_addr_IMG_LoadLBM_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadLBM_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadLBM_IO")
	}
	_addr_IMG_LoadPCX_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadPCX_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadPCX_IO")
	}
	_addr_IMG_LoadPNG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadPNG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadPNG_IO")
	}
	_addr_IMG_LoadPNM_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadPNM_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadPNM_IO")
	}
	_addr_IMG_LoadSVG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadSVG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadSVG_IO")
	}
	_addr_IMG_LoadQOI_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadQOI_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadQOI_IO")
	}
	_addr_IMG_LoadTGA_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTGA_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTGA_IO")
	}
	_addr_IMG_LoadTIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTIF_IO")
	}
	_addr_IMG_LoadXCF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadXCF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadXCF_IO")
	}
	_addr_IMG_LoadXPM_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadXPM_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadXPM_IO")
	}
	_addr_IMG_LoadXV_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadXV_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadXV_IO")
	}
	_addr_IMG_LoadWEBP_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadWEBP_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadWEBP_IO")
	}
	_addr_IMG_LoadSizedSVG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadSizedSVG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadSizedSVG_IO")
	}
	_addr_IMG_ReadXPMFromArray, err = puregogen.OpenSymbol(_hnd_img, "IMG_ReadXPMFromArray")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_ReadXPMFromArray")
	}
	_addr_IMG_ReadXPMFromArrayToRGB888, err = puregogen.OpenSymbol(_hnd_img, "IMG_ReadXPMFromArrayToRGB888")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_ReadXPMFromArrayToRGB888")
	}
	_addr_IMG_SaveAVIF, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveAVIF")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveAVIF")
	}
	_addr_IMG_SaveAVIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveAVIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveAVIF_IO")
	}
	_addr_IMG_SavePNG, err = puregogen.OpenSymbol(_hnd_img, "IMG_SavePNG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SavePNG")
	}
	_addr_IMG_SavePNG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_SavePNG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SavePNG_IO")
	}
	_addr_IMG_SaveJPG, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveJPG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveJPG")
	}
	_addr_IMG_SaveJPG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveJPG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveJPG_IO")
	}
	_addr_IMG_LoadAnimation, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAnimation")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAnimation")
	}
	_addr_IMG_LoadAnimation_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAnimation_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAnimation_IO")
	}
	_addr_IMG_LoadAnimationTyped_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAnimationTyped_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAnimationTyped_IO")
	}
	_addr_IMG_FreeAnimation, err = puregogen.OpenSymbol(_hnd_img, "IMG_FreeAnimation")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_FreeAnimation")
	}
	_addr_IMG_LoadGIFAnimation_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadGIFAnimation_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadGIFAnimation_IO")
	}
	_addr_IMG_LoadWEBPAnimation_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadWEBPAnimation_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadWEBPAnimation_IO")
	}

	iVersion = func() int32 {
		_r0, _, _ := purego.SyscallN(_addr_IMG_Version)
		__r0 := int32(_r0)
		return __r0
	}
	iLoadTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTyped_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(typ))))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoad = func(file string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_Load, uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoad_IO = func(src *sdl.IOStream, closeio bool) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_Load_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadTexture = func(renderer *sdl.Renderer, file string) *sdl.Texture {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTexture, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := (*sdl.Texture)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadTexture_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTexture_IO, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio))
		__r0 := (*sdl.Texture)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadTextureTyped_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, typ string) *sdl.Texture {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTextureTyped_IO, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(typ))))
		__r0 := (*sdl.Texture)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadAVIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAVIF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadICO_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadICO_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadCUR_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadCUR_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadBMP_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadBMP_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadGIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadGIF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadJPG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadJPG_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadJXL_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadJXL_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadLBM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadLBM_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadPCX_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadPCX_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadPNG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadPNG_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadPNM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadPNM_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadSVG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadSVG_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadQOI_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadQOI_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadTGA_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTGA_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadTIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTIF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadXCF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadXCF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadXPM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadXPM_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadXV_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadXV_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadWEBP_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadWEBP_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadSizedSVG_IO = func(src *sdl.IOStream, width int32, height int32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadSizedSVG_IO, uintptr(unsafe.Pointer(src)), uintptr(width), uintptr(height))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iReadXPMFromArray = func(xpm *string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_ReadXPMFromArray, uintptr(unsafe.Pointer(xpm)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iReadXPMFromArrayToRGB888 = func(xpm *string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_ReadXPMFromArrayToRGB888, uintptr(unsafe.Pointer(xpm)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iSaveAVIF = func(surface *sdl.Surface, file string, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveAVIF, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))), uintptr(quality))
		__r0 := _r0 != 0
		return __r0
	}
	iSaveAVIF_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveAVIF_IO, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), puregogen.BoolToUintptr(closeio), uintptr(quality))
		__r0 := _r0 != 0
		return __r0
	}
	iSavePNG = func(surface *sdl.Surface, file string) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SavePNG, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := _r0 != 0
		return __r0
	}
	iSavePNG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SavePNG_IO, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), puregogen.BoolToUintptr(closeio))
		__r0 := _r0 != 0
		return __r0
	}
	iSaveJPG = func(surface *sdl.Surface, file string, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveJPG, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))), uintptr(quality))
		__r0 := _r0 != 0
		return __r0
	}
	iSaveJPG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveJPG_IO, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), puregogen.BoolToUintptr(closeio), uintptr(quality))
		__r0 := _r0 != 0
		return __r0
	}
	iLoadAnimation = func(file string) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAnimation, uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadAnimation_IO = func(src *sdl.IOStream, closeio bool) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAnimation_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadAnimationTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAnimationTyped_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(typ))))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iFreeAnimation = func(anim *Animation) {
		purego.SyscallN(_addr_IMG_FreeAnimation, uintptr(unsafe.Pointer(anim)))
	}
	iLoadGIFAnimation_IO = func(src *sdl.IOStream) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadGIFAnimation_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iLoadWEBPAnimation_IO = func(src *sdl.IOStream) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadWEBPAnimation_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
}
