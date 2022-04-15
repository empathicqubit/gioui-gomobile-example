package piet

import (
	_ "embed"
	"runtime"

	"gioui.org/shader"
)

var (
	Shader_backdrop_comp = shader.Sources{
		Name:           "backdrop.comp",
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{128, 1, 1},
	}
	zbackdrop_comp_0_spirv                string
	zbackdrop_comp_0_dxbc                 string
	zbackdrop_comp_0_metallibmacos        string
	zbackdrop_comp_0_metallibios          string
	zbackdrop_comp_0_metallibiossimulator string
	Shader_binning_comp                   = shader.Sources{
		Name:           "binning.comp",
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{128, 1, 1},
	}
	zbinning_comp_0_spirv                string
	zbinning_comp_0_dxbc                 string
	zbinning_comp_0_metallibmacos        string
	zbinning_comp_0_metallibios          string
	zbinning_comp_0_metallibiossimulator string
	Shader_coarse_comp                   = shader.Sources{
		Name:           "coarse.comp",
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{128, 1, 1},
	}
	zcoarse_comp_0_spirv                string
	zcoarse_comp_0_dxbc                 string
	zcoarse_comp_0_metallibmacos        string
	zcoarse_comp_0_metallibios          string
	zcoarse_comp_0_metallibiossimulator string
	Shader_elements_comp                = shader.Sources{
		Name:           "elements.comp",
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "SceneBuf", Binding: 2}, {Name: "StateBuf", Binding: 3}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{32, 1, 1},
	}
	zelements_comp_0_spirv                string
	zelements_comp_0_dxbc                 string
	zelements_comp_0_metallibmacos        string
	zelements_comp_0_metallibios          string
	zelements_comp_0_metallibiossimulator string
	Shader_kernel4_comp                   = shader.Sources{
		Name:           "kernel4.comp",
		Images:         []shader.ImageBinding{{Name: "images", Binding: 3}, {Name: "image", Binding: 2}},
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{16, 8, 1},
	}
	zkernel4_comp_0_spirv                string
	zkernel4_comp_0_dxbc                 string
	zkernel4_comp_0_metallibmacos        string
	zkernel4_comp_0_metallibios          string
	zkernel4_comp_0_metallibiossimulator string
	Shader_path_coarse_comp              = shader.Sources{
		Name:           "path_coarse.comp",
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{32, 1, 1},
	}
	zpath_coarse_comp_0_spirv                string
	zpath_coarse_comp_0_dxbc                 string
	zpath_coarse_comp_0_metallibmacos        string
	zpath_coarse_comp_0_metallibios          string
	zpath_coarse_comp_0_metallibiossimulator string
	Shader_tile_alloc_comp                   = shader.Sources{
		Name:           "tile_alloc.comp",
		StorageBuffers: []shader.BufferBinding{{Name: "Memory", Binding: 0}, {Name: "ConfigBuf", Binding: 1}},
		WorkgroupSize:  [3]int{128, 1, 1},
	}
	ztile_alloc_comp_0_spirv                string
	ztile_alloc_comp_0_dxbc                 string
	ztile_alloc_comp_0_metallibmacos        string
	ztile_alloc_comp_0_metallibios          string
	ztile_alloc_comp_0_metallibiossimulator string
)

func init() {
	const (
		opengles = runtime.GOOS == "linux" || runtime.GOOS == "freebsd" || runtime.GOOS == "openbsd" || runtime.GOOS == "windows" || runtime.GOOS == "js" || runtime.GOOS == "android" || runtime.GOOS == "darwin" || runtime.GOOS == "ios"
		opengl   = runtime.GOOS == "darwin"
		d3d11    = runtime.GOOS == "windows"
		vulkan   = runtime.GOOS == "linux" || runtime.GOOS == "android"
	)
	if vulkan {
		Shader_backdrop_comp.SPIRV = zbackdrop_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_backdrop_comp.DXBC = zbackdrop_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_backdrop_comp.MetalLib = zbackdrop_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_backdrop_comp.MetalLib = zbackdrop_comp_0_metallibiossimulator
		} else {
			Shader_backdrop_comp.MetalLib = zbackdrop_comp_0_metallibios
		}
	}
	if vulkan {
		Shader_binning_comp.SPIRV = zbinning_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_binning_comp.DXBC = zbinning_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_binning_comp.MetalLib = zbinning_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_binning_comp.MetalLib = zbinning_comp_0_metallibiossimulator
		} else {
			Shader_binning_comp.MetalLib = zbinning_comp_0_metallibios
		}
	}
	if vulkan {
		Shader_coarse_comp.SPIRV = zcoarse_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_coarse_comp.DXBC = zcoarse_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_coarse_comp.MetalLib = zcoarse_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_coarse_comp.MetalLib = zcoarse_comp_0_metallibiossimulator
		} else {
			Shader_coarse_comp.MetalLib = zcoarse_comp_0_metallibios
		}
	}
	if vulkan {
		Shader_elements_comp.SPIRV = zelements_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_elements_comp.DXBC = zelements_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_elements_comp.MetalLib = zelements_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_elements_comp.MetalLib = zelements_comp_0_metallibiossimulator
		} else {
			Shader_elements_comp.MetalLib = zelements_comp_0_metallibios
		}
	}
	if vulkan {
		Shader_kernel4_comp.SPIRV = zkernel4_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_kernel4_comp.DXBC = zkernel4_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_kernel4_comp.MetalLib = zkernel4_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_kernel4_comp.MetalLib = zkernel4_comp_0_metallibiossimulator
		} else {
			Shader_kernel4_comp.MetalLib = zkernel4_comp_0_metallibios
		}
	}
	if vulkan {
		Shader_path_coarse_comp.SPIRV = zpath_coarse_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_path_coarse_comp.DXBC = zpath_coarse_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_path_coarse_comp.MetalLib = zpath_coarse_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_path_coarse_comp.MetalLib = zpath_coarse_comp_0_metallibiossimulator
		} else {
			Shader_path_coarse_comp.MetalLib = zpath_coarse_comp_0_metallibios
		}
	}
	if vulkan {
		Shader_tile_alloc_comp.SPIRV = ztile_alloc_comp_0_spirv
	}
	if opengles {
	}
	if opengl {
	}
	if d3d11 {
		Shader_tile_alloc_comp.DXBC = ztile_alloc_comp_0_dxbc
	}
	if runtime.GOOS == "darwin" {
		Shader_tile_alloc_comp.MetalLib = ztile_alloc_comp_0_metallibmacos
	}
	if runtime.GOOS == "ios" {
		if runtime.GOARCH == "amd64" {
			Shader_tile_alloc_comp.MetalLib = ztile_alloc_comp_0_metallibiossimulator
		} else {
			Shader_tile_alloc_comp.MetalLib = ztile_alloc_comp_0_metallibios
		}
	}
}
