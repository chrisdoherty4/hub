package grub

import (
	"reflect"
	"testing"
)

var workingConfig = `#
# DO NOT EDIT THIS FILE
#
# It is automatically generated by grub-mkconfig using templates
# from /etc/grub.d and settings from /etc/default/grub
#

### BEGIN /etc/grub.d/00_header ###
if [ -s $prefix/grubenv ]; then
  set have_grubenv=true
  load_env
fi
if [ "${next_entry}" ] ; then
   set default="${next_entry}"
   set next_entry=
   save_env next_entry
   set boot_once=true
else
   set default="0"
fi

if [ x"${feature_menuentry_id}" = xy ]; then
  menuentry_id_option="--id"
else
  menuentry_id_option=""
fi

export menuentry_id_option

if [ "${prev_saved_entry}" ]; then
  set saved_entry="${prev_saved_entry}"
  save_env saved_entry
  set prev_saved_entry=
  save_env prev_saved_entry
  set boot_once=true
fi

function savedefault {
  if [ -z "${boot_once}" ]; then
    saved_entry="${chosen}"
    save_env saved_entry
  fi
}
function recordfail {
  set recordfail=1
  # GRUB lacks write support for lvm, so recordfail support is disabled.
}
function load_video {
  if [ x$feature_all_video_module = xy ]; then
    insmod all_video
  else
    insmod efi_gop
    insmod efi_uga
    insmod ieee1275_fb
    insmod vbe
    insmod vga
    insmod video_bochs
    insmod video_cirrus
  fi
}

if [ x$feature_default_font_path = xy ] ; then
   font=unicode
else
insmod part_msdos
insmod lvm
insmod ext2
set root='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'
if [ x$feature_platform_search_hint = xy ]; then
  search --no-floppy --fs-uuid --set=root --hint='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'  afad0891-8f09-4093-9520-d02a1fdf5726
else
  search --no-floppy --fs-uuid --set=root afad0891-8f09-4093-9520-d02a1fdf5726
fi
    font="/usr/share/grub/unicode.pf2"
fi

if loadfont $font ; then
  set gfxmode=auto
  load_video
  insmod gfxterm
  set locale_dir=$prefix/locale
  set lang=en_GB
  insmod gettext
fi
terminal_output gfxterm
if [ "${recordfail}" = 1 ] ; then
  set timeout=30
else
  if [ x$feature_timeout_style = xy ] ; then
    set timeout_style=hidden
    set timeout=2
  # Fallback hidden-timeout code in case the timeout_style feature is
  # unavailable.
  elif sleep --interruptible 2 ; then
    set timeout=0
  fi
fi
if [ $grub_platform = efi ]; then
  set timeout=30
  if [ x$feature_timeout_style = xy ] ; then
    set timeout_style=menu
  fi
fi
### END /etc/grub.d/00_header ###

### BEGIN /etc/grub.d/05_debian_theme ###
set menu_color_normal=white/black
set menu_color_highlight=black/light-gray
### END /etc/grub.d/05_debian_theme ###

### BEGIN /etc/grub.d/10_linux ###
function gfxmode {
        set gfxpayload="${1}"
        if [ "${1}" = "keep" ]; then
                set vt_handoff=vt.handoff=1
        else
                set vt_handoff=
        fi
}
if [ "${recordfail}" != 1 ]; then
  if [ -e ${prefix}/gfxblacklist.txt ]; then
    if hwmatch ${prefix}/gfxblacklist.txt 3; then
      if [ ${match} = 0 ]; then
        set linux_gfx_mode=keep
      else
        set linux_gfx_mode=text
      fi
    else
      set linux_gfx_mode=text
    fi
  else
    set linux_gfx_mode=keep
  fi
else
  set linux_gfx_mode=text
fi
export linux_gfx_mode
menuentry 'Ubuntu' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-simple-afad0891-8f09-4093-9520-d02a1fdf5726' {
        recordfail
        load_video
        gfxmode $linux_gfx_mode
        insmod gzio
        if [ x$grub_platform = xxen ]; then insmod xzio; insmod lzopio; fi
        insmod part_msdos
        insmod lvm
        insmod ext2
        set root='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'
        if [ x$feature_platform_search_hint = xy ]; then
          search --no-floppy --fs-uuid --set=root --hint='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'  afad0891-8f09-4093-9520-d02a1fdf5726
        else
          search --no-floppy --fs-uuid --set=root afad0891-8f09-4093-9520-d02a1fdf5726
        fi
        linux   /boot/vmlinuz-4.15.0-135-generic root=/dev/mapper/code--vg-root ro  
        initrd  /boot/initrd.img-4.15.0-135-generic
}
submenu 'Advanced options for Ubuntu' $menuentry_id_option 'gnulinux-advanced-afad0891-8f09-4093-9520-d02a1fdf5726' {
        menuentry 'Ubuntu, with Linux 4.15.0-135-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-4.15.0-135-generic-advanced-afad0891-8f09-4093-9520-d02a1fdf5726' {
                recordfail
                load_video
                gfxmode $linux_gfx_mode
                insmod gzio
                if [ x$grub_platform = xxen ]; then insmod xzio; insmod lzopio; fi
                insmod part_msdos
                insmod lvm
                insmod ext2
                set root='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root --hint='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'  afad0891-8f09-4093-9520-d02a1fdf5726
                else
                  search --no-floppy --fs-uuid --set=root afad0891-8f09-4093-9520-d02a1fdf5726
                fi
                echo    'Loading Linux 4.15.0-135-generic ...'
                linux   /boot/vmlinuz-4.15.0-135-generic root=/dev/mapper/code--vg-root ro  
                echo    'Loading initial ramdisk ...'
                initrd  /boot/initrd.img-4.15.0-135-generic
        }
        menuentry 'Ubuntu, with Linux 4.15.0-135-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-4.15.0-135-generic-recovery-afad0891-8f09-4093-9520-d02a1fdf5726' {
                recordfail
                load_video
                insmod gzio
                if [ x$grub_platform = xxen ]; then insmod xzio; insmod lzopio; fi
                insmod part_msdos
                insmod lvm
                insmod ext2
                set root='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root --hint='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'  afad0891-8f09-4093-9520-d02a1fdf5726
                else
                  search --no-floppy --fs-uuid --set=root afad0891-8f09-4093-9520-d02a1fdf5726
                fi
                echo    'Loading Linux 4.15.0-135-generic ...'
                linux   /boot/vmlinuz-4.15.0-135-generic root=/dev/mapper/code--vg-root ro recovery nomodeset dis_ucode_ldr 
                echo    'Loading initial ramdisk ...'
                initrd  /boot/initrd.img-4.15.0-135-generic
        }
        menuentry 'Ubuntu, with Linux 4.15.0-132-generic' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-4.15.0-132-generic-advanced-afad0891-8f09-4093-9520-d02a1fdf5726' {
                recordfail
                load_video
                gfxmode $linux_gfx_mode
                insmod gzio
                if [ x$grub_platform = xxen ]; then insmod xzio; insmod lzopio; fi
                insmod part_msdos
                insmod lvm
                insmod ext2
                set root='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root --hint='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'  afad0891-8f09-4093-9520-d02a1fdf5726
                else
                  search --no-floppy --fs-uuid --set=root afad0891-8f09-4093-9520-d02a1fdf5726
                fi
                echo    'Loading Linux 4.15.0-132-generic ...'
                linux   /boot/vmlinuz-4.15.0-132-generic root=/dev/mapper/code--vg-root ro  
                echo    'Loading initial ramdisk ...'
                initrd  /boot/initrd.img-4.15.0-132-generic
        }
        menuentry 'Ubuntu, with Linux 4.15.0-132-generic (recovery mode)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'gnulinux-4.15.0-132-generic-recovery-afad0891-8f09-4093-9520-d02a1fdf5726' {
                recordfail
                load_video
                insmod gzio
                if [ x$grub_platform = xxen ]; then insmod xzio; insmod lzopio; fi
                insmod part_msdos
                insmod lvm
                insmod ext2
                set root='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root --hint='lvmid/EK2Sae-O7Mr-LuAq-cVlI-8aBN-3ov1-Huas9B/GXj3DU-Lxkc-LdcD-2nv6-NdnY-LUuJ-g5uaHO'  afad0891-8f09-4093-9520-d02a1fdf5726
                else
                  search --no-floppy --fs-uuid --set=root afad0891-8f09-4093-9520-d02a1fdf5726
                fi
                echo    'Loading Linux 4.15.0-132-generic ...'
                linux   /boot/vmlinuz-4.15.0-132-generic root=/dev/mapper/code--vg-root ro recovery nomodeset dis_ucode_ldr 
                echo    'Loading initial ramdisk ...'
                initrd  /boot/initrd.img-4.15.0-132-generic
        }
}

### END /etc/grub.d/10_linux ###

### BEGIN /etc/grub.d/20_linux_xen ###

### END /etc/grub.d/20_linux_xen ###

### BEGIN /etc/grub.d/30_os-prober ###
menuentry 'Ubuntu 20.04.1 LTS (20.04) (on /dev/nbd0p1)' --class ubuntu --class gnu-linux --class gnu --class os $menuentry_id_option 'osprober-gnulinux-simple-f90a6910-6dba-4817-884d-fe54732c4e04' {
        insmod part_gpt
        insmod ext2
        if [ x$feature_platform_search_hint = xy ]; then
          search --no-floppy --fs-uuid --set=root  f90a6910-6dba-4817-884d-fe54732c4e04
        else
          search --no-floppy --fs-uuid --set=root f90a6910-6dba-4817-884d-fe54732c4e04
        fi
        linux /boot/vmlinuz-5.4.0-64-generic root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro console=tty1 console=ttyS0 panic=-1
        initrd /boot/initrd.img-5.4.0-64-generic
}
submenu 'Advanced options for Ubuntu 20.04.1 LTS (20.04) (on /dev/nbd0p1)' $menuentry_id_option 'osprober-gnulinux-advanced-f90a6910-6dba-4817-884d-fe54732c4e04' {
        menuentry 'Ubuntu (on /dev/nbd0p1)' --class gnu-linux --class gnu --class os $menuentry_id_option 'osprober-gnulinux-/boot/vmlinuz-5.4.0-64-generic--f90a6910-6dba-4817-884d-fe54732c4e04' {
                insmod part_gpt
                insmod ext2
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root  f90a6910-6dba-4817-884d-fe54732c4e04
                else
                  search --no-floppy --fs-uuid --set=root f90a6910-6dba-4817-884d-fe54732c4e04
                fi
                linux /boot/vmlinuz-5.4.0-64-generic root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro console=tty1 console=ttyS0 panic=-1
                initrd /boot/initrd.img-5.4.0-64-generic
        }
        menuentry 'Ubuntu, with Linux 5.4.0-64-generic (on /dev/nbd0p1)' --class gnu-linux --class gnu --class os $menuentry_id_option 'osprober-gnulinux-/boot/vmlinuz-5.4.0-64-generic--f90a6910-6dba-4817-884d-fe54732c4e04' {
                insmod part_gpt
                insmod ext2
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root  f90a6910-6dba-4817-884d-fe54732c4e04
                else
                  search --no-floppy --fs-uuid --set=root f90a6910-6dba-4817-884d-fe54732c4e04
                fi
                linux /boot/vmlinuz-5.4.0-64-generic root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro console=tty1 console=ttyS0 panic=-1
                initrd /boot/initrd.img-5.4.0-64-generic
        }
        menuentry 'Ubuntu, with Linux 5.4.0-64-generic (recovery mode) (on /dev/nbd0p1)' --class gnu-linux --class gnu --class os $menuentry_id_option 'osprober-gnulinux-/boot/vmlinuz-5.4.0-64-generic-root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro recovery nomodeset dis_ucode_ldr panic=-1-f90a6910-6dba-4817-884d-fe54732c4e04' {
                insmod part_gpt
                insmod ext2
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root  f90a6910-6dba-4817-884d-fe54732c4e04
                else
                  search --no-floppy --fs-uuid --set=root f90a6910-6dba-4817-884d-fe54732c4e04
                fi
                linux /boot/vmlinuz-5.4.0-64-generic root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro recovery nomodeset dis_ucode_ldr panic=-1
                initrd /boot/initrd.img-5.4.0-64-generic
        }
        menuentry 'Ubuntu, with Linux 5.4.0-62-generic (on /dev/nbd0p1)' --class gnu-linux --class gnu --class os $menuentry_id_option 'osprober-gnulinux-/boot/vmlinuz-5.4.0-62-generic--f90a6910-6dba-4817-884d-fe54732c4e04' {
                insmod part_gpt
                insmod ext2
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root  f90a6910-6dba-4817-884d-fe54732c4e04
                else
                  search --no-floppy --fs-uuid --set=root f90a6910-6dba-4817-884d-fe54732c4e04
                fi
                linux /boot/vmlinuz-5.4.0-62-generic root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro console=tty1 console=ttyS0 panic=-1
                initrd /boot/initrd.img-5.4.0-62-generic
        }
        menuentry 'Ubuntu, with Linux 5.4.0-62-generic (recovery mode) (on /dev/nbd0p1)' --class gnu-linux --class gnu --class os $menuentry_id_option 'osprober-gnulinux-/boot/vmlinuz-5.4.0-62-generic-root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro recovery nomodeset dis_ucode_ldr panic=-1-f90a6910-6dba-4817-884d-fe54732c4e04' {
                insmod part_gpt
                insmod ext2
                if [ x$feature_platform_search_hint = xy ]; then
                  search --no-floppy --fs-uuid --set=root  f90a6910-6dba-4817-884d-fe54732c4e04
                else
                  search --no-floppy --fs-uuid --set=root f90a6910-6dba-4817-884d-fe54732c4e04
                fi
                linux /boot/vmlinuz-5.4.0-62-generic root=PARTUUID=70cd2749-4cff-4ad7-aac4-5fbdd09b87d9 ro recovery nomodeset dis_ucode_ldr panic=-1
                initrd /boot/initrd.img-5.4.0-62-generic
        }
}

set timeout_style=menu
if [ "${timeout}" = 0 ]; then
  set timeout=10
fi
### END /etc/grub.d/30_os-prober ###

### BEGIN /etc/grub.d/30_uefi-firmware ###
### END /etc/grub.d/30_uefi-firmware ###

### BEGIN /etc/grub.d/40_custom ###
# This file provides an easy way to add custom menu entries.  Simply type the
# menu entries you want to add after this comment.  Be careful not to change
# the 'exec tail' line above.
### END /etc/grub.d/40_custom ###

### BEGIN /etc/grub.d/41_custom ###
if [ -f  ${config_directory}/custom.cfg ]; then
  source ${config_directory}/custom.cfg
elif [ -z "${config_directory}" -a -f  $prefix/custom.cfg ]; then
  source $prefix/custom.cfg;
fi
### END /etc/grub.d/41_custom ###`

func TestParseGrubCfg(t *testing.T) {
	type args struct {
		grubcfg string
	}
	tests := []struct {
		name              string
		args              args
		wantConfigs       []Config
		wantDefaultConfig int64
	}{
		{
			"working",
			args{workingConfig},
			nil,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotConfigs, gotDefaultConfig := ParseGrubCfg(tt.args.grubcfg)

			// Should be able to parse 11 configurations
			if len(gotConfigs) != 11 {
				t.Errorf("ParseGrubCfg() gotConfigs = %v, want %v", gotConfigs, tt.wantConfigs)
			}
			if gotDefaultConfig != tt.wantDefaultConfig {
				t.Errorf("ParseGrubCfg() gotDefaultConfig = %v, want %v", gotDefaultConfig, tt.wantDefaultConfig)
			}
		})
	}
}

func TestGetDefaultCconfig(t *testing.T) {
	type args struct {
		grubcfg string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *Config
	}{
		{
			"working",
			args{workingConfig},
			&Config{
				"'Ubuntu' ",
				"/boot/vmlinuz-4.15.0-135-generic",
				"/boot/initrd.img-4.15.0-135-generic",
				"root=/dev/mapper/code--vg-root ro",
				"",
				"",
				nil,
			},
		},
	}
	// Should determine our "default configuration"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCfg := GetDefaultConfig(tt.args.grubcfg); !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("GetDefaultConfig() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
