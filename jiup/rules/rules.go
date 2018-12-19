package rules

import (
	"errors"
	"fmt"
	"strings"

	"github.com/just-install/just-install-updater-go/jiup/rules/d"
	"github.com/just-install/just-install-updater-go/jiup/rules/h"
	"github.com/just-install/just-install-updater-go/jiup/rules/v"
	"github.com/just-install/just-install-updater-go/jiup/rules/w"
)

func init() {
	Rule("7zip",
		v.Regexp(
			"https://7-zip.org/download.html",
			h.Re("Download 7-Zip ([0-9][0-9].[0-9][0-9])"),
		),
		d.Template(
			"https://www.7-zip.org/a/7z{{.VersionN}}.msi",
			"https://www.7-zip.org/a/7z{{.VersionN}}-x64.msi",
		),
	)
	Rule("anaconda",
		v.Regexp(
			"https://www.anaconda.com/download/",
			h.Re("Anaconda3-([0-9.]+)-"),
		),
		d.HTMLA(
			"https://www.anaconda.com/download/",
			"a[href*='Anaconda3-'][href$='-Windows-x86.exe']",
			"a[href*='Anaconda3-'][href$='-Windows-x86_64.exe']",
		),
	)
	Rule("android-studio-ide",
		v.Regexp(
			"https://developer.android.com/studio/",
			h.Re("install/([0-9.]+)/android-studio-ide-"),
		),
		d.HTMLA(
			"https://developer.android.com/studio/",
			"a[href*='android-studio-ide'][href$='-windows.exe'].button.devsite-dialog-button",
			"",
		),
	)
	Rule("arduino",
		v.Regexp(
			"https://www.arduino.cc/en/Main/Software",
			h.Re("arduino-([0-9.]+)-"),
		),
		d.Template(
			"https://downloads.arduino.cc/arduino-{{.Version}}-windows.exe",
			"",
		),
	)
	Rule("audacity",
		v.Regexp(
			"http://www.oldfoss.com/Audacity.html",
			h.Re("audacity-win-([0-9.]+).exe"),
		),
		func(version string) (*string, *string, error) {
			return d.Regexp(
				"http://www.oldfoss.com/Audacity.html",
				h.Re("\"(http.+audacity-win-"+version+".exe)\""),
				nil,
			)(version)
		},
	)
	Rule("bleachbit",
		v.Regexp(
			"https://www.bleachbit.org/download/windows",
			h.Re("BleachBit-([0-9.]+)-setup.exe"),
		),
		d.Template(
			"https://download.bleachbit.org/BleachBit-{{.Version}}-setup.exe",
			"",
		),
	)
	Rule("bcc",
		v.GitHubRelease(
			"wormt/bcc",
			h.Re("(.+)"),
		),
		d.GitHubRelease(
			"wormt/bcc",
			h.Re("bcc-.+-32bit.zip"),
			h.Re("bcc-.+-64bit.zip"),
		),
	)
	Rule("bcuninstaller",
		v.GitHubRelease(
			"Klocman/Bulk-Crap-Uninstaller",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"Klocman/Bulk-Crap-Uninstaller",
			h.Re(".*setup.exe"),
			nil,
		),
	)
	Rule("bitpay",
		v.GitHubRelease(
			"bitpay/copay",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"bitpay/copay",
			h.Re("BitPay.exe"),
			nil,
		),
	)
	Rule("brackets",
		v.GitHubRelease(
			"adobe/brackets",
			h.Re("release-(.+)"),
		),
		d.GitHubRelease(
			"adobe/brackets",
			h.Re("Brackets.Release.*.msi"),
			nil,
		),
	)
	Rule("ccleaner",
		func() (string, error) {
			version, err := v.Regexp(
				"https://www.ccleaner.com/ccleaner/download/standard",
				h.Re("ccsetup([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}
			return string(version[0]) + "." + string(version[1:]), nil
		},
		d.HTMLA(
			"https://www.ccleaner.com/ccleaner/download/standard",
			"a[href$='.exe']:contains('start the download')",
			"",
		),
	)
	Rule("cdburnerxp",
		v.Regexp(
			"https://download.cdburnerxp.se/msi/",
			h.Re("_([0-9.]+).msi"),
		),
		d.HTMLA(
			"https://download.cdburnerxp.se/msi/",
			"a[href^='cdbxp_setup_'][href$='msi']:not([href~='x64'])",
			"a[href^='cdbxp_setup_x64'][href$='msi']",
		),
	)
	Rule("classic-shell",
		w.UnderscoreToDot(v.Regexp(
			"http://www.oldfoss.com/Classic-Shell.html",
			h.Re("ClassicShellSetup_([0-9_]+)"),
		)),
		func(version string) (*string, *string, error) {
			return d.Regexp(
				"http://www.oldfoss.com/Classic-Shell.html",
				h.Re("\"(http.+ClassicShellSetup_"+strings.Replace(version, ".", "_", -1)+".exe)\""),
				nil,
			)(version)
		},
	)
	Rule("clementine-player",
		v.GitHubRelease(
			"clementine-player/Clementine",
			h.Re("(.+)"),
		),
		d.GitHubRelease(
			"clementine-player/Clementine",
			h.Re("ClementineSetup-.*.exe"),
			nil,
		),
	)
	Rule("cmake",
		v.Regexp(
			"https://cmake.org/download/",
			h.Re("Latest Release \\(([0-9.]+)\\)"),
		),
		d.HTMLA(
			"https://cmake.org/download/",
			"a[href$='-win32-x86.msi']",
			"a[href$='-win64-x64.msi']",
		),
	)
	Rule("colemak",
		v.Regexp(
			"https://colemak.com/Windows",
			h.Re("Colemak-([0-9.]+)"),
		),
		d.HTMLA(
			"https://colemak.com/Windows",
			"a:contains('Download now')",
			"",
		),
	)
	Rule("conemu",
		v.GitHubRelease(
			"Maximus5/ConEmu",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"Maximus5/ConEmu",
			h.Re("ConEmuSetup.*.exe"),
			nil,
		),
	)
	Rule("cpu-z",
		v.Regexp(
			"https://www.cpuid.com/softwares/cpu-z.html",
			h.Re("Version ([0-9.]+) for [Ww]indows"),
		),
		d.Template(
			"http://download.cpuid.com/cpu-z/cpu-z_{{.Version}}-en.exe",
			"",
		),
	)
	Rule("cryptomator",
		v.HTML(
			"https://cryptomator.org/downloads",
			"meta[itemprop='softwareVersion']",
			"content",
			nil,
		),
		d.HTMLA(
			"https://cryptomator.org/downloads",
			"",
			"#winDownload a[href$='.exe']:contains('64 Bit')",
		),
	)
	Rule("crystaldisk-info",
		w.UnderscoreToDot(v.HTML(
			"https://osdn.net/projects/crystaldiskinfo/releases/",
			"a.pref-download-btn.pref-download-btn-win32[href]",
			"href",
			h.Re("CrystalDiskInfo([0-9_]+).zip"),
		)),
		func(version string) (*string, *string, error) {
			vu := strings.Replace(version, ".", "_", -1)
			dlp, err := v.HTML(
				"https://osdn.net/projects/crystaldiskinfo/releases/",
				"a.pref-download-btn.pref-download-btn-win32[href]",
				"href",
				h.Re("downloads/([0-9]+/CrystalDiskInfo"+vu+").zip"),
			)()
			if err != nil {
				return nil, nil, err
			}
			return h.StrPtr("http://osdn.dl.osdn.jp/crystaldiskinfo/" + dlp + ".exe"), nil, nil
		},
	)
	Rule("crystaldisk-mark",
		v.HTML(
			"https://osdn.net/projects/crystaldiskmark/releases/",
			".release-item-title a[href*='crystaldiskmark/releases']",
			"innerText",
			h.Re("([0-9.]+)"),
		),
		func(version string) (*string, *string, error) {
			vu := strings.Replace(version, ".", "_", -1)
			x86, x64, err := d.HTMLA(
				"https://osdn.net/dl/crystaldiskmark/CrystalDiskMark"+vu+".exe",
				"a.mirror_link",
				"",
			)(version)
			return x86, x64, err
		},
	)
	Rule("cyberduck",
		v.Regexp(
			"https://update.cyberduck.io/windows/?C=M;O=D",
			h.Re("Cyberduck-Installer-([0-9.]+).msi"),
		),
		d.HTMLA(
			"https://update.cyberduck.io/windows/?C=M;O=D",
			"a[href$='.msi']",
			"",
		),
	)
	Rule("dbeaver",
		v.GitHubRelease(
			"dbeaver/dbeaver",
			h.Re("(.+)"),
		),
		d.GitHubRelease(
			"dbeaver/dbeaver",
			h.Re("dbeaver-ce-.+-x86-setup.exe"),
			h.Re("dbeaver-ce-.+-x86_64-setup.exe"),
		),
	)
	Rule("defraggler",
		func() (string, error) {
			version, err := v.Regexp(
				"https://www.ccleaner.com/defraggler/download/standard",
				h.Re("dfsetup([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}
			return string(version[0]) + "." + string(version[1:]), nil
		},
		d.HTMLA(
			"https://www.ccleaner.com/defraggler/download/standard",
			"a[href$='.exe']:contains('start the download')",
			"",
		),
	)
	Rule("deluge",
		v.Regexp(
			"https://dev.deluge-torrent.org/wiki/Download",
			h.Re("Latest Release: <strong>([0-9.]+)"),
		),
		d.Template(
			"http://download.deluge-torrent.org/windows/deluge-{{.Version}}-win32-py2.7.exe",
			"",
		),
	)
	Rule("dependency-walker",
		v.Regexp(
			"http://www.dependencywalker.com",
			h.Re("Dependency Walker ([0-9.]+)"),
		),
		d.Template(
			"http://www.dependencywalker.com/depends{{.VersionN}}_x86.zip",
			"http://www.dependencywalker.com/depends{{.VersionN}}_x64.zip",
		),
	)
	Rule("deskpins",
		v.Regexp(
			"https://efotinis.neocities.org/deskpins/",
			h.Re("v([0-9.]+)"),
		),
		d.HTMLA(
			"https://efotinis.neocities.org/deskpins/",
			"a[href*='DeskPins-'][href$='-setup.exe']",
			"",
		),
	)
	Rule("ditto",
		v.Regexp(
			"http://ditto-cp.sourceforge.net/index.php",
			h.Re("versionDots ?= ?\"([0-9.]+)\""),
		),
		d.Template(
			"https://sourceforge.net/projects/ditto-cp/files/Ditto/{{.Version}}/DittoSetup_{{.VersionU}}.exe/download",
			"https://sourceforge.net/projects/ditto-cp/files/Ditto/{{.Version}}/DittoSetup_64bit_{{.VersionU}}.exe/download",
		),
	)
	Rule("doublecmd",
		v.Regexp(
			"https://sourceforge.net/p/doublecmd/wiki/Download/",
			h.Re("doublecmd-([0-9.]+)\\."),
		),
		d.HTMLA(
			"https://sourceforge.net/p/doublecmd/wiki/Download/",
			"a[href$='i386-win32.msi/download']",
			"a[href$='x86_64-win64.msi/download']",
		),
	)
	Rule("duck",
		v.Regexp(
			"https://dist.duck.sh/?C=M;O=D",
			h.Re("duck-([0-9.]+).msi"),
		),
		d.HTMLA(
			"https://dist.duck.sh/?C=M;O=D",
			"a[href$='.msi']",
			"",
		),
	)
	Rule("eac",
		v.Regexp(
			"http://www.exactaudiocopy.de/en/index.php/weitere-seiten/download-from-alternative-servers-2/",
			h.Re("Exact Audio Copy V([0-9.]+)"),
		),
		d.HTMLA(
			"http://www.exactaudiocopy.de/en/index.php/weitere-seiten/download-from-alternative-servers-2/",
			"a[href*='eac'][href$='.exe']:contains('Download Installer')",
			"",
		),
	)
	for _, edition := range []string{
		"committers",
		"cpp",
		"java",
		"jee",
		"php",
	} {
		Rule("eclipse-"+edition,
			v.Regexp(
				"https://eclipse.org/downloads/eclipse-packages/",
				h.Re("eclipse-"+edition+"-([a-zA-Z0-9]+-[a-zA-Z0-9]+)-win32"), // e.g. photon-R
			),
			func(version string) (*string, *string, error) {
				x86, x64, err := d.HTMLA(
					"https://eclipse.org/downloads/eclipse-packages/",
					".downloadLink-content a[href*='?file='][href*='eclipse-"+edition+"-'][href$='-win32.zip']",
					".downloadLink-content a[href*='?file='][href*='eclipse-"+edition+"-'][href$='-win32-x86_64.zip']",
				)(version)

				if err != nil {
					return nil, nil, err
				}

				*x86 = "http://ftp.osuosl.org/pub/eclipse" + strings.SplitN(strings.SplitN(*x86, "?file=", 2)[1], "&", 2)[0]
				*x64 = "http://ftp.osuosl.org/pub/eclipse" + strings.SplitN(strings.SplitN(*x64, "?file=", 2)[1], "&", 2)[0]

				return x86, x64, nil
			},
		)
	}
	Rule("eig",
		v.GitHubRelease(
			"EvilInsultGenerator/c-sharp-desktop",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"EvilInsultGenerator/c-sharp-desktop",
			h.Re("EvilInsultGenerator_Setup.exe"),
			nil,
		),
	)
	Rule("emacs",
		func() (string, error) {
			majorVersion, err := v.HTML(
				"https://ftp.gnu.org/gnu/emacs/windows/?C=M;O=D",
				"a[href*='emacs-']",
				"href",
				h.Re("emacs-([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}

			version, err := v.HTML(
				"https://ftp.gnu.org/gnu/emacs/windows/emacs-"+majorVersion+"/?C=N;O=D",
				"a[href*='emacs-']",
				"href",
				h.Re("emacs-([0-9.]+)"),
			)()
			if err != nil {
				return "", err
			}

			if strings.Split(version, ".")[0] != majorVersion {
				return "", errors.New("emacs rule needs to be updated")
			}

			return version, nil
		},
		func(version string) (*string, *string, error) {
			majorVersion := strings.Split(version, ".")[0]
			x86 := fmt.Sprintf("https://ftp.gnu.org/gnu/emacs/windows/emacs-%s/emacs-%s-i686.zip", majorVersion, version)
			x64 := fmt.Sprintf("https://ftp.gnu.org/gnu/emacs/windows/emacs-%s/emacs-%s-x86_64.zip", majorVersion, version)
			return &x86, &x64, nil
		},
	)
	Rule("enpass",
		v.HTML(
			"https://www.enpass.io/downloads/",
			"a[href*='Enpass_'][href$='_Setup.exe']",
			"href",
			h.Re("Enpass_([0-9.]+)_"),
		),
		d.HTMLA(
			"https://www.enpass.io/downloads/",
			"a[href*='Enpass_'][href$='_Setup.exe']",
			"",
		),
	)
	Rule("erlang",
		v.Regexp(
			"https://www.erlang.org/downloads/",
			h.Re("DOWNLOAD\\s+OTP\\s+([0-9.]+)"),
		),
		d.HTMLA(
			"https://www.erlang.org/downloads/",
			"a[href*='win32'][href$='exe']:contains('Windows 32-bit Binary File')",
			"a[href*='win64'][href$='exe']:contains('Windows 64-bit Binary File')",
		),
	)
	Rule("etcher",
		v.GitHubRelease(
			"resin-io/etcher",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"resin-io/etcher",
			h.Re("Etcher-Setup-.+-x86.exe"),
			h.Re("Etcher-Setup-.+-x64.exe"),
		),
	)
	Rule("everything-search",
		v.Regexp(
			"https://www.voidtools.com/downloads/",
			h.Re("Download Everything ([0-9.]+)"),
		),
		d.HTMLA(
			"https://www.voidtools.com/downloads/",
			"a[href$='x86-Setup.exe']:contains('Download Installer')",
			"a[href$='x64-Setup.exe']:contains('Download Installer 64-bit')",
		),
	)
	Rule("filezilla-server",
		w.UnderscoreToDot(v.HTML(
			"https://download.filezilla-project.org/server/?C=M;O=D",
			"a[href*='FileZilla_Server-'][href$='.exe']",
			"href",
			h.Re("FileZilla_Server-([0-9_]+)"),
		)),
		d.HTMLA(
			"https://download.filezilla-project.org/server/?C=M;O=D",
			"a[href*='FileZilla_Server-'][href$='.exe']",
			"",
		),
	)
	Rule("flash-player",
		v.HTML(
			"http://get.adobe.com/flashplayer/about/",
			"td:contains('Opera, Chromium-based browsers - PPAPI')+td",
			"innerText",
			h.Re("([0-9.]+)"),
		),
		d.Template(
			"https://fpdownload.macromedia.com/get/flashplayer/pdc/{{.Version}}/install_flash_player_{{.Version0}}_plugin.msi",
			"",
		),
	)
	Rule("flash-player-ie",
		v.HTML(
			"http://get.adobe.com/flashplayer/about/",
			"td:contains('Internet Explorer - ActiveX')+td",
			"innerText",
			h.Re("([0-9.]+)"),
		),
		d.Template(
			"https://fpdownload.macromedia.com/get/flashplayer/pdc/{{.Version}}/install_flash_player_{{.Version0}}_active_x.msi",
			"",
		),
	)
	Rule("freefilesync",
		v.Regexp(
			"https://www.freefilesync.org/download.php",
			h.Re("Download FreeFileSync ([0-9.]+)"),
		),
		d.HTMLA(
			"https://www.freefilesync.org/download.php",
			"a.direct-download-link[href$='.exe']",
			"",
		),
	)
	Rule("freeplane",
		v.Regexp(
			"https://sourceforge.net/projects/freeplane/files/freeplane%20stable/",
			h.Re("Freeplane-Setup-([0-9.]+).exe"),
		),
		d.Template(
			"https://sourceforge.net/projects/freeplane/files/freeplane%20stable/Freeplane-Setup-{{.Version}}.exe/download",
			"",
		),
	)
	Rule("geforce-experience",
		v.Regexp(
			"https://www.nvidia.com/en-us/geforce/geforce-experience/",
			h.Re("https://us.download.nvidia.com/GFE/GFEClient/([0-9.]+)/GeForce_Experience"),
		),
		d.HTMLA(
			"https://www.nvidia.com/en-us/geforce/geforce-experience/",
			"a[href^='https://us.download.nvidia.com/GFE/GFEClient/'][href$='.exe']",
			"",
		),
	)
	Rule("gimp",
		v.Regexp(
			"https://www.gimp.org/downloads/",
			h.Re("current stable release of GIMP is <b>([0-9.]+)"),
		),
		d.HTMLA(
			"https://www.gimp.org/downloads/",
			"#win a[href*='-setup'][href$='.exe']",
			"",
		),
	)
	Rule("git",
		v.GitHubRelease(
			"git-for-windows/git",
			h.Re("v([0-9.]+)\\.windows.+"),
		),
		d.GitHubRelease(
			"git-for-windows/git",
			h.Re("Git-.+-32-bit.exe"),
			h.Re("Git-.+-64-bit.exe"),
		),
	)
	Rule("gitextensions",
		v.GitHubRelease(
			"gitextensions/gitextensions",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"gitextensions/gitextensions",
			h.Re("GitExtensions-.*.msi"),
			nil,
		),
	)
	Rule("git-credential-manager-for-windows",
		v.GitHubRelease(
			"Microsoft/Git-Credential-Manager-for-Windows",
			h.Re("(.+)"),
		),
		d.GitHubRelease(
			"Microsoft/Git-Credential-Manager-for-Windows",
			h.Re("GCMW-.+.exe"),
			nil,
		),
	)
	Rule("git-lfs",
		v.GitHubRelease(
			"git-lfs/git-lfs",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"git-lfs/git-lfs",
			h.Re("git-lfs-windows-v.+.exe"),
			nil,
		),
	)
	Rule("go",
		v.Regexp(
			"https://golang.org/dl/",
			h.Re("go([0-9.]+)\\.windows"),
		),
		d.Template(
			"https://dl.google.com/go/go{{.Version}}.windows-386.msi",
			"https://dl.google.com/go/go{{.Version}}.windows-amd64.msi",
		),
	)
	Rule("gog-galaxy",
		v.Regexp(
			"https://www.gog.com/galaxy",
			h.Re("setup_galaxy_([0-9.]+).exe"),
		),
		d.HTMLA(
			"https://www.gog.com/galaxy",
			"a[href*='setup'][href$='.exe']:contains('Windows')",
			"",
		),
	)
	Rule("gow",
		v.GitHubRelease(
			"bmatzelle/gow",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"bmatzelle/gow",
			h.Re("Gow-.+.exe"),
			nil,
		),
	)
	Rule("greenshot",
		v.GitHubRelease(
			"greenshot/greenshot",
			h.Re("Greenshot-RELEASE-([0-9.]+)"),
		),
		d.GitHubRelease(
			"greenshot/greenshot",
			h.Re("Greenshot-INSTALLER-.+-RELEASE.exe"),
			nil,
		),
	)
	Rule("gvim",
		v.Regexp(
			"https://www.vim.org/download.php",
			h.Re("latest version \\(currently ([0-9.]+)\\)"),
		),
		d.HTMLA(
			"http://ftp.vim.org/pub/vim/pc/?C=M;O=D",
			"a[href*='gvim'][href$='.exe']",
			"",
		),
	)
	Rule("hashcheck",
		v.GitHubRelease(
			"gurnec/HashCheck",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"gurnec/HashCheck",
			h.Re("HashCheckSetup-.+.exe"),
			nil,
		),
	)
	Rule("heidisql",
		v.Regexp(
			"https://www.heidisql.com/download.php",
			h.Re("HeidiSQL_([0-9.]+)_"),
		),
		d.HTMLA(
			"https://www.heidisql.com/download.php",
			"a[href$='Setup.exe']:contains('Installer')",
			"",
		),
	)
	Rule("hexchat",
		v.Regexp(
			"https://hexchat.github.io/downloads.html",
			h.Re("HexChat ([0-9.]+)"),
		),
		d.HTMLA(
			"https://hexchat.github.io/downloads.html",
			"a[href$='x86.exe']",
			"a[href$='x64.exe']",
		),
	)
	Rule("hugo",
		v.GitHubRelease(
			"gohugoio/hugo",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"gohugoio/hugo",
			h.Re("hugo_.+_Windows-32bit.zip"),
			h.Re("hugo_.+_Windows-64bit.zip"),
		),
	)
	Rule("imageglass",
		v.GitHubRelease(
			"d2phap/ImageGlass",
			h.Re("([0-9.]+)"),
		),
		d.GitHubRelease(
			"d2phap/ImageGlass",
			h.Re("ImageGlass_.+.exe"),
			nil,
		),
	)
	Rule("inkscape",
		v.Regexp(
			"https://inkscape.org/en/release/",
			h.Re("Download Inkscape ([0-9.]+)"),
		),
		d.Template(
			"https://media.inkscape.org/dl/resources/file/inkscape-{{.Version}}-x86.msi",
			"https://media.inkscape.org/dl/resources/file/inkscape-{{.Version}}-x64.msi",
		),
	)
	Rule("jdk",
		v.Regexp(
			"https://lv.binarybabel.org/catalog-api/java/jdk8.txt?p=version",
			h.Re("([0-9a-zA-Z.-]+)"),
		),
		func(version string) (*string, *string, error) {
			buf, _, ok, err := h.GetURL(nil, "https://lv.binarybabel.org/catalog-api/java/jdk8.txt?p=downloads.exe", map[string]string{}, []int{200})
			if err != nil {
				return nil, nil, err
			}
			if !ok {
				return nil, nil, errors.New("unexpected response code")
			}
			x64 := string(buf)
			x86 := strings.Replace(x64, "x64", "i586", -1)
			return &x86, &x64, nil
		},
	)
	Rule("jre",
		func() (string, error) {
			version, err := v.Regexp("https://www.java.com/en/download/manual.jsp", h.Re("Recommended Version ([0-9]* Update [0-9]*)"))()
			if err != nil {
				return "", err
			}
			return strings.Replace(version, " Update ", ".", 1), nil
		},
		d.HTMLA(
			"https://www.java.com/en/download/manual.jsp",
			"a[title='Download Java software for Windows Offline']",
			"a[title='Download Java software for Windows (64-bit)']",
		),
	)
	Rule("keepass",
		v.Regexp(
			"https://sourceforge.net/projects/keepass/files/",
			h.Re("KeePass-([0-9.]+)\\."),
		),
		d.Template(
			"https://sourceforge.net/projects/keepass/files/KeePass%202.x/{{.Version}}/KeePass-{{.Version}}.msi/download",
			"",
		),
	)
	Rule("keepassxc",
		v.GitHubRelease(
			"keepassxreboot/keepassxc",
			h.Re("([0-9.]+)"),
		),
		d.GitHubRelease(
			"keepassxreboot/keepassxc",
			h.Re("KeePassXC-.+-Win32.msi"),
			h.Re("KeePassXC-.+-Win64.msi"),
		),
	)
	Rule("keeweb",
		v.GitHubRelease(
			"keeweb/keeweb",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"keeweb/keeweb",
			h.Re("KeeWeb-.+.win.ia32.exe"),
			h.Re("KeeWeb-.+.win.x64.exe"),
		),
	)
	Rule("kicad",
		v.Regexp(
			"http://kicad-pcb.org/download/windows/",
			h.Re("Stable Release Current Version: ([0-9.]+)"),
		),
		d.HTMLA(
			"http://kicad-pcb.org/download/windows/",
			"a[href$='.exe']:contains('Windows 32-bit')",
			"a[href$='.exe']:contains('Windows 64-bit')",
		),
	)
	Rule("kodi",
		v.Regexp(
			"http://mirrors.kodi.tv/releases/windows/win32/?C=M&O=D",
			h.Re("kodi-([0-9.]+)-"),
		),
		d.HTMLA(
			"http://mirrors.kodi.tv/releases/windows/win32/?C=M&O=D",
			"a[href*='kodi'][href$='-x86.exe']",
			"",
		),
	)
	Rule("krita",
		v.Regexp(
			"https://krita.org/en/download/krita-desktop/",
			h.Re("krita-x86-([0-9.]+)-"),
		),
		d.HTMLA(
			"https://krita.org/en/download/krita-desktop/",
			"a[href*='krita-x86'][href$='.exe']",
			"a[href*='krita-x64'][href$='.exe']",
		),
	)
	Rule("libreoffice",
		v.Regexp(
			"https://www.libreoffice.org/download/libreoffice-fresh/?type=win-x86&lang=en-US",
			h.Re("LibreOffice ([0-9.]+) "),
		),
		d.Template(
			"https://download.documentfoundation.org/libreoffice/stable/{{.Version}}/win/x86/LibreOffice_{{.Version}}_Win_x86.msi",
			"https://download.documentfoundation.org/libreoffice/stable/{{.Version}}/win/x86_64/LibreOffice_{{.Version}}_Win_x64.msi",
		),
	)
	Rule("lockhunter",
		v.Regexp(
			"http://lockhunter.com/download.htm",
			h.Re("Version: ([0-9.]+)"),
		),
		d.Template(
			"http://lockhunter.com/exe/lockhuntersetup_{{.VersionD}}.exe",
			"",
		),
	)
	Rule("marktext",
		v.GitHubRelease(
			"marktext/marktext",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"marktext/marktext",
			h.Re("marktext-setup-.+.exe"),
			nil,
		),
	)
	Rule("mercurial",
		v.Regexp(
			"https://www.mercurial-scm.org/sources.js",
			h.Re("windows/mercurial-([0-9.]+)-"),
		),
		d.Regexp(
			"https://www.mercurial-scm.org/sources.js",
			h.Re("(https://www.mercurial-scm.org/release/windows/mercurial-[0-9.]+-x86.msi)"),
			h.Re("(https://www.mercurial-scm.org/release/windows/mercurial-[0-9.]+-x64.msi)"),
		),
	)
	Rule("mono",
		v.Regexp(
			"http://www.mono-project.com/download/stable/",
			h.Re("[0-9.]+ Stable \\(([0-9.]+)\\)"),
		),
		d.HTMLA(
			"http://www.mono-project.com/download/stable/",
			"a[href*='download.mono-project.com'][href*='windows-installer'][href$='.msi']:not([href*='gtksharp'])",
			"",
		),
	)
	Rule("mountainduck",
		v.Regexp(
			"https://mountainduck.io/",
			h.Re("Installer-([0-9.]+).exe"),
		),
		d.HTMLA(
			"https://mountainduck.io/",
			"a[href*='Installer'][href$='.msi']",
			"",
		),
	)
	Rule("mp3tag",
		v.Regexp(
			"https://www.mp3tag.de/en/download.html",
			h.Re("Mp3tag v([0-9.a-z]+)"),
		),
		d.HTMLA(
			"https://www.mp3tag.de/en/dodownload.html",
			"a[href*='download'][href$='.exe']:contains('here')",
			"",
		),
	)
	Rule("mpc-hc",
		v.Regexp(
			"https://mpc-hc.org/downloads/",
			h.Re("latest stable build is v([0-9.]+)"),
		),
		d.HTMLA(
			"https://mpc-hc.org/downloads/",
			"a[href$='.x86.exe']",
			"a[href$='.x64.exe']",
		),
	)
	Rule("mumble",
		v.GitHubRelease(
			"mumble-voip/mumble",
			h.Re("([0-9.]+)"),
		),
		d.GitHubRelease(
			"mumble-voip/mumble",
			h.Re("mumble-.+.msi"),
			nil,
		),
	)
	Rule("naps2",
		v.GitHubRelease(
			"cyanfish/naps2",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"cyanfish/naps2",
			h.Re("naps2-.+-setup.msi"),
			nil,
		),
	)
	Rule("nextcloud",
		v.Regexp(
			"https://download.nextcloud.com/desktop/releases/Windows/?C=M;O=D",
			h.Re("Nextcloud-([0-9.]+)-"),
		),
		d.HTMLA(
			"https://download.nextcloud.com/desktop/releases/Windows/?C=M;O=D",
			"a[href$='setup.exe']",
			"",
		),
	)
	Rule("node",
		v.Regexp(
			"https://nodejs.org/en/download/current/",
			h.Re("Latest Current Version: <strong>([0-9.]+)"),
		),
		d.HTMLA(
			"https://nodejs.org/en/download/current/",
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('32-bit')",
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('64-bit')",
		),
	)
	Rule("node-lts",
		v.Regexp(
			"https://nodejs.org/en/download/",
			h.Re("Latest LTS Version: <strong>([0-9.]+)"),
		),
		d.HTMLA(
			"https://nodejs.org/en/download/",
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('32-bit')",
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('64-bit')",
		),
	)
	Rule("notepad++",
		v.Regexp(
			"https://notepad-plus-plus.org/download/",
			h.Re("Download Notepad\\+\\+ ([0-9.]+)"),
		),
		d.HTMLA(
			"https://notepad-plus-plus.org/download/",
			"a[href*='npp'][href$='nstaller.exe']",
			"a[href*='npp'][href$='nstaller.x64.exe']",
		),
	)
	Rule("notepad2-mod",
		v.GitHubRelease(
			"XhmikosR/notepad2-mod",
			h.Re("([0-9.]+)"),
		),
		d.GitHubRelease(
			"XhmikosR/notepad2-mod",
			h.Re("Notepad2-mod..+.exe"),
			nil,
		),
	)
	Rule("npackd",
		v.GitHubRelease(
			"tim-lebedkov/npackd-cpp",
			h.Re("version_([0-9.]+)"),
		),
		d.GitHubRelease(
			"tim-lebedkov/npackd-cpp",
			h.Re("Npackd32-.+.msi"),
			h.Re("Npackd64-.+.msi"),
		),
	)
	Rule("npackdcl",
		v.GitHubRelease(
			"tim-lebedkov/npackd-cpp",
			h.Re("version_([0-9.]+)"),
		),
		d.GitHubRelease(
			"tim-lebedkov/npackd-cpp",
			h.Re("NpackdCL-.+.msi"),
			nil,
		),
	)
	Rule("nxlog",
		v.Regexp(
			"https://nxlog.co/products/nxlog-community-edition/download",
			h.Re("nxlog-ce-([0-9.]+)\\.msi"),
		),
		d.HTMLA(
			"https://nxlog.co/products/nxlog-community-edition/download",
			"a[href*='nxlog-ce-'][href$='.msi']",
			"",
		),
	)
	Rule("obs-studio",
		v.Regexp(
			"https://obsproject.com/download",
			h.Re("Version: ([0-9.]+)"),
		),
		d.HTMLA(
			"https://obsproject.com/download",
			"a[href*='OBS-Studio-'][href$='Full-Installer-x64.exe']",
			"",
		),
	)
	Rule("octave",
		v.Regexp(
			"https://ftp.gnu.org/gnu/octave/windows/?C=M;O=D",
			h.Re("octave-([0-9.]+)-w32-installer.exe"),
		),
		d.HTMLA(
			"https://ftp.gnu.org/gnu/octave/windows/?C=M;O=D",
			"a[href*='octave-'][href$='-w32-installer.exe']",
			"a[href*='octave-'][href$='-w64-installer.exe']",
		),
	)
	Rule("open-hardware-monitor",
		v.Regexp(
			"http://openhardwaremonitor.org/downloads/",
			h.Re("Open Hardware Monitor ([0-9.]+)"),
		),
		d.HTMLA(
			"http://openhardwaremonitor.org/downloads/",
			"a[href*='openhardwaremonitor-'][href$='.zip']",
			"",
		),
	)
	Rule("openssh",
		v.Regexp(
			"https://www.mls-software.com/opensshd.html",
			h.Re("OpenSSH ([0-9.]+)p"),
		),
		d.HTMLA(
			"https://www.mls-software.com/opensshd.html",
			"a[href*='setupssh-'][href$='.exe']",
			"",
		),
	)
	Rule("perl",
		v.HTML(
			"http://strawberryperl.com/releases.html",
			"a[href*='strawberry-perl-'][href$='32bit.msi']",
			"href",
			h.Re("strawberry-perl-([0-9.]+)-"),
		),
		d.HTMLA(
			"http://strawberryperl.com/releases.html",
			"a[href*='strawberry-perl-'][href$='32bit.msi']",
			"a[href*='strawberry-perl-'][href$='64bit.msi']",
		),
	)
	Rule("php",
	     v.Regexp(
		        "https://windows.php.net/download",
		        h.Re("PHP [0-9.]+ \(([0-9.]+)\)"),
	     ),
        )
	Rule("pia",
		v.Regexp(
			"https://www.privateinternetaccess.com/pages/downloads",
			h.Re("v([0-9]+)"),
		),
		d.HTMLA(
			"https://www.privateinternetaccess.com/pages/downloads",
			"a[href*='pia-'][href$='installer-win.exe']",
			"",
		),
	)
	Rule("plex-media-server",
		v.Regexp(
			"https://plex.tv/api/downloads/1.json",
			h.Re("version\":\"([0-9.]+)"),
		),
		d.Regexp(
			"https://plex.tv/api/downloads/1.json",
			h.Re("\"(https://downloads.plex.tv/plex-media-server/[0-9a-z.-]+?/Plex-Media-Server-[0-9a-z.-]+?.exe)\""),
			nil,
		),
	)
	Rule("powershell-core",
		v.GitHubRelease(
			"PowerShell/PowerShell",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"PowerShell/PowerShell",
			h.Re("PowerShell-.+-win-x86.msi"),
			h.Re("PowerShell-.+-win-x64.msi"),
		),
	)
	Rule("processhacker",
		v.GitHubRelease(
			"processhacker/processhacker",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"processhacker/processhacker",
			h.Re("processhacker-.+-setup.exe"),
			nil,
		),
	)
	Rule("putty",
		v.Regexp(
			"http://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html",
			h.Re("latest release \\(([0-9.]+)\\)"),
		),
		d.HTMLA(
			"http://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html",
			"span.downloadfile a[href^='https'][href*='w32/putty'][href$='.msi']",
			"span.downloadfile a[href^='https'][href*='w64/putty'][href$='.msi']",
		),
	)
	Rule("pycharm-community",
		v.Regexp(
			"https://data.services.jetbrains.com/products/releases?code=PCP%2CPCC&latest=true",
			h.Re("version\":\"([0-9.]+)"),
		),
		d.Regexp(
			"https://data.services.jetbrains.com/products/releases?code=PCP%2CPCC&latest=true",
			h.Re("\"(https://download.jetbrains.com/python/pycharm-community-[0-9.]+.exe)\""),
			nil,
		),
	)
	Rule("python2",
		v.HTML(
			"https://www.python.org/downloads/windows",
			"a:contains('Python 2')",
			"innerText",
			h.Re("Python (2\\.[0-9.]+)"),
		),
		d.Template(
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}.msi",
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}.amd64.msi",
		),
	)
	Rule("python2-win32",
		v.GitHubRelease(
			"mhammond/pywin32",
			h.Re("b(.+)"),
		),
		d.GitHubRelease(
			"mhammond/pywin32",
			h.Re("pywin32-.+.win32-py2.7.exe"),
			h.Re("pywin32-.+.win-amd64-py2.7.exe"),
		),
	)
	Rule("python3",
		v.HTML(
			"https://www.python.org/downloads/",
			".download-for-current-os .download-os-windows a[href*='python-3']",
			"innerText",
			h.Re("Download Python ([0-9.]+)"),
		),
		d.Template(
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}.exe",
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}-amd64.exe",
		),
	)
	Rule("qbittorrent",
		v.Regexp(
			"http://www.oldfoss.com/qBittorrent.html",
			h.Re("qbittorrent_([0-9.]+)_setup.exe"),
		),
		func(version string) (*string, *string, error) {
			return d.Regexp(
				"http://www.oldfoss.com/qBittorrent.html",
				h.Re("\"(http.+qbittorrent_"+version+"_setup.exe)\""),
				h.Re("\"(http.+qbittorrent_"+version+"_x64_setup.exe)\""),
			)(version)
		},
	)
	Rule("qtox",
		v.GitHubRelease(
			"qTox/qTox",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"qTox/qTox",
			h.Re("setup-qtox-i686-release.exe"),
			h.Re("setup-qtox-x86_64-release.exe"),
		),
	)
	Rule("recuva",
		func() (string, error) {
			version, err := v.Regexp(
				"https://www.ccleaner.com/recuva/download/standard",
				h.Re("rcsetup([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}
			return string(version[0]) + "." + string(version[1:]), nil
		},
		d.HTMLA(
			"https://www.ccleaner.com/recuva/download/standard",
			"a[href$='.exe']:contains('start the download')",
			"",
		),
	)
	Rule("retroarch",
		v.Regexp(
			"https://www.retroarch.com/?page=platforms",
			h.Re("https://buildbot.libretro.com/stable/([0-9.]+)/windows/"),
		),
		d.HTMLA(
			"https://www.retroarch.com/?page=platforms",
			"a[href$='.exe']:contains('Installer (32bit)')",
			"a[href$='.exe']:contains('Installer (64bit)')",
		),
	)
	Rule("ruby",
		v.GitHubRelease(
			"oneclick/rubyinstaller2",
			h.Re("rubyinstaller-([0-9.]+)"),
		),
		d.GitHubRelease(
			"oneclick/rubyinstaller2",
			h.Re("rubyinstaller-[0-9.]+-.+-x86.exe"),
			h.Re("rubyinstaller-[0-9.]+-.+-x64.exe"),
		),
	)
	Rule("seafile-client",
		w.NoHTTPSForVersionExtractor(v.HTML(
			"https://www.seafile.com/en/download/",
			".txt > h3:contains('Client for Windows')~a[href*='seafile'][href$='en.msi'].download-op",
			"innerText",
			h.Re("([0-9.]+)"),
		)),
		w.NoHTTPSForDownloadExtractor(d.HTML(
			"https://www.seafile.com/en/download/",
			".txt > h3:contains('Client for Windows')~a[href*='seafile'][href$='en.msi'].download-op",
			"",
			"href",
			"",
			nil,
			nil,
		)),
	)
	Rule("sharex",
		v.GitHubRelease(
			"ShareX/ShareX",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"ShareX/ShareX",
			h.Re("ShareX-.+-setup.exe"),
			nil,
		),
	)
	Rule("signal",
		v.Regexp(
			"https://updates.signal.org/desktop/latest.yml",
			h.Re("version: ([0-9.]+)"),
		),
		d.Template(
			"https://updates.signal.org/desktop/signal-desktop-win-{{.Version}}.exe",
			"",
		),
	)
	Rule("simplenote",
		v.GitHubRelease(
			"Automattic/simplenote-electron",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"Automattic/simplenote-electron",
			h.Re("Simplenote-win-[0-9.]+.exe"),
			nil,
		),
	)
	Rule("sharpkeys",
		v.GitHubRelease(
			"randyrants/sharpkeys",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"randyrants/sharpkeys",
			h.Re("sharpkeys.+.msi"),
			nil,
		),
	)
	Rule("smplayer",
		v.Regexp(
			"https://sourceforge.net/projects/smplayer/files/",
			h.Re("smplayer-([0-9.]+)\\."),
		),
		d.Template(
			"https://sourceforge.net/projects/smplayer/files/SMPlayer/{{.Version}}/smplayer-{{.Version}}-win32.exe/download",
			"https://sourceforge.net/projects/smplayer/files/SMPlayer/{{.Version}}/smplayer-{{.Version}}-x64.exe/download",
		),
	)
	Rule("sourcetree",
		v.Regexp(
			"https://www.sourcetreeapp.com",
			h.Re("SourceTreeSetup-([0-9.]+)\\.exe"),
		),
		d.HTMLA(
			"https://www.sourcetreeapp.com",
			"a[href*='SourceTreeSetup'][href$='exe']",
			"",
		),
	)
	Rule("sublime-text",
		v.Regexp(
			"https://www.sublimetext.com/2",
			h.Re("Version:</i> ([0-9.]+)"),
		),
		d.HTMLA(
			"https://www.sublimetext.com/2",
			"#dl_win_32 a[href$='exe']",
			"#dl_win_64 a[href$='exe']",
		),
	)
	Rule("sublime-text-3",
		v.Regexp(
			"https://www.sublimetext.com/3",
			h.Re("Version:</i> Build ([0-9]+)"),
		),
		d.HTMLA(
			"https://www.sublimetext.com/3",
			"#dl_win_32 a[href$='exe']",
			"#dl_win_64 a[href$='exe']",
		),
	)
	Rule("sublime-text-dev",
		v.Regexp(
			"https://www.sublimetext.com/3dev",
			h.Re("Version:</i> Build ([0-9]+)"),
		),
		d.HTMLA(
			"https://www.sublimetext.com/3dev",
			"#dl_win_32 a[href$='exe']",
			"#dl_win_64 a[href$='exe']",
		),
	)
	Rule("subversion",
		v.Regexp(
			"https://sliksvn.com/download/",
			h.Re("Subversion-([0-9.]+)-"),
		),
		d.HTMLA(
			"https://sliksvn.com/download/",
			".client a[href$='zip']:contains('32 bit')",
			".client a[href$='zip']:contains('64 bit')",
		),
	)
	Rule("sumatrapdf",
		v.Regexp(
			"https://www.sumatrapdfreader.org/news.html",
			h.Re(">([0-9.]+) \\(20"),
		),
		d.Template(
			"https://www.sumatrapdfreader.org/dl/SumatraPDF-{{.Version}}-install.exe",
			"https://www.sumatrapdfreader.org/dl/SumatraPDF-{{.Version}}-64-install.exe",
		),
	)
	Rule("syncthing",
		v.GitHubRelease(
			"syncthing/syncthing",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"syncthing/syncthing",
			h.Re(".*windows-386.*.zip"),
			h.Re(".*windows-amd64.*.zip"),
		),
	)
	Rule("teamspeak",
		v.Regexp(
			"https://www.teamspeak.com/en/downloads",
			h.Re("Client-win32-([0-9.]+)\\.exe"),
		),
		d.HTML(
			"https://www.teamspeak.com/en/downloads",
			"option[value*='win32'][value$='.exe']",
			"option[value*='win64'][value$='.exe']",
			"value",
			"value",
			nil,
			nil,
		),
	)
	Rule("tightvnc",
		v.Regexp(
			"https://tightvnc.com/download.php",
			h.Re("Version ([0-9.]+)"),
		),
		d.HTMLA(
			"https://tightvnc.com/download.php",
			"a[href*='tightvnc-'][href$='-setup-32bit.msi']",
			"a[href*='tightvnc-'][href$='-setup-64bit.msi']",
		),
	)
	Rule("tor-browser",
		v.Regexp(
			"https://www.torproject.org/download/download.html.en",
			h.Re("torbrowser-install-([0-9.]+)_en"),
		),
		d.HTMLA(
			"https://www.torproject.org/download/download.html.en",
			"a.button.win-tbb",
			"a.button.win-tbb64",
		),
	)
	Rule("tortoisegit",
		v.Regexp(
			"https://tortoisegit.org/download/",
			h.Re("TortoiseGit-([0-9.]+)"),
		),
		d.HTMLA(
			"https://tortoisegit.org/download/",
			"a[href$='32bit.msi']",
			"a[href$='64bit.msi']",
		),
	)
	Rule("tortoisesvn",
		v.Regexp(
			"https://tortoisesvn.net/downloads.html",
			h.Re("The current version is ([0-9.]+)"),
		),
		func(version string) (*string, *string, error) {
			// Layer 1: Link to OSDN
			x86, x64, err := d.HTMLA(
				"https://tortoisesvn.net/downloads.html",
				"a[href^='https://osdn.net'][href*='win32-svn']",
				"a[href^='https://osdn.net'][href*='x64-svn']",
			)(version)
			if err != nil {
				return nil, nil, err
			}
			if x64 == nil {
				return nil, nil, errors.New("x64 link empty")
			}
			// Layer 2: OSDN to redir link
			x86, _, err = d.HTMLA(
				*x86,
				"a.mirror_link[href*='/frs/redir'][href*='win32-svn']",
				"",
			)(version)
			if err != nil {
				return nil, nil, err
			}
			_, x64, err = d.HTMLA(
				*x64,
				"a",
				"a.mirror_link[href*='/frs/redir'][href*='x64-svn']",
			)(version)
			if err != nil {
				return nil, nil, err
			}
			return x86, x64, nil
		},
	)
	Rule("upx",
		v.GitHubRelease(
			"upx/upx",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"upx/upx",
			h.Re("upx-[0-9.]+-win32.zip"),
			h.Re("upx-[0-9.]+-win64.zip"),
		),
	)
	Rule("vagrant",
		v.Regexp(
			"https://www.vagrantup.com/downloads.html",
			h.Re("vagrant_([0-9.]+)_"),
		),
		d.HTMLA(
			"https://www.vagrantup.com/downloads.html",
			"a[href*='vagrant_'][href$='_i686.msi']",
			"a[href*='vagrant_'][href$='_x86_64.msi']",
		),
	)
	Rule("veracrypt",
		v.Regexp(
			"https://sourceforge.net/projects/veracrypt/files/",
			h.Re("VeraCrypt_([0-9.]+)_"),
		),
		d.Template(
			"https://sourceforge.net/projects/veracrypt/files/VeraCrypt%20{{.Version}}/VeraCrypt%20Setup%20{{.Version}}.exe/download",
			"",
		),
	)
	Rule("virtualbox",
		v.Regexp(
			"https://www.virtualbox.org/wiki/Downloads",
			h.Re("VirtualBox-([0-9.]+)-"),
		),
		d.HTMLA(
			"https://www.virtualbox.org/wiki/Downloads",
			"a[href$='.exe']:contains('Windows')",
			"",
		),
	)
	Rule("virtualbox-extpack",
		v.Regexp(
			"https://www.virtualbox.org/wiki/Downloads",
			h.Re("VirtualBox_Extension_Pack-([0-9.]+)\\."),
		),
		d.HTMLA(
			"https://www.virtualbox.org/wiki/Downloads",
			"a[href$='.vbox-extpack']",
			"",
		),
	)
	Rule("vivaldi",
		v.Regexp(
			"https://vivaldi.com/download/",
			h.Re("Vivaldi\\.([0-9.]+)\\.exe"),
		),
		d.HTMLA(
			"https://vivaldi.com/download/",
			"a[href*='Vivaldi.'][href$='.exe']:not([href$='.x64.exe'])",
			"a[href*='Vivaldi.'][href$='.x64.exe']",
		),
	)
	Rule("vlc",
		v.Regexp(
			"https://download.videolan.org/pub/videolan/vlc/last/win32/",
			h.Re("vlc-([0-9.]+)-win32.msi"),
		),
		d.Template(
			"https://download.videolan.org/pub/videolan/vlc/last/win32/vlc-{{.Version}}-win32.msi",
			"https://download.videolan.org/pub/videolan/vlc/last/win64/vlc-{{.Version}}-win64.msi",
		),
	)
	Rule("vpnunlimited",
		v.Regexp(
			"https://www.vpnunlimitedapp.com/en/downloads/windows",
			h.Re("_v([0-9.]+)\\."),
		),
		d.HTMLA(
			"https://www.vpnunlimitedapp.com/en/downloads/windows",
			"a[href*='VPN_Unlimited_'][href$='.exe']",
			"",
		),
	)
	Rule("webtorrent",
		v.GitHubRelease(
			"webtorrent/webtorrent-desktop",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"webtorrent/webtorrent-desktop",
			h.Re("WebTorrentSetup-v[0-9.]+-ia32.exe"),
			h.Re("WebTorrentSetup-v[0-9.]+.exe"),
		),
	)
	Rule("winrar",
		v.Regexp(
			"https://www.win-rar.com/download.html",
			h.Re("WinRAR ([0-9.]+) "),
		),
		d.Template(
			"https://rarlab.com/rar/wrar{{.VersionN}}.exe",
			"https://rarlab.com/rar/winrar-x64-{{.VersionN}}.exe",
		),
	)
	Rule("winscp",
		v.Regexp(
			"https://sourceforge.net/projects/winscp/files/",
			h.Re("WinSCP-([0-9.]+)-"),
		),
		d.Template(
			"https://sourceforge.net/projects/winscp/files/WinSCP/{{.Version}}/WinSCP-{{.Version}}-Setup.exe/download",
			"",
		),
	)
	Rule("wireshark",
		v.Regexp(
			"https://www.wireshark.org/download.html",
			h.Re("Stable Release \\(([0-9.]+)\\)"),
		),
		d.HTMLA(
			"https://www.wireshark.org/download.html",
			"a[href*='Wireshark-win32-'][href$='.exe']",
			"a[href*='Wireshark-win64-'][href$='.exe']",
		),
	)
	Rule("wixedit",
		v.GitHubRelease(
			"WixEdit/WixEdit",
			h.Re("v([0-9]+\\.[0-9]+\\.[0-9]+)"),
		),
		d.GitHubRelease(
			"WixEdit/WixEdit",
			h.Re("wixedit-.+.msi"),
			nil,
		),
	)
	Rule("workflowy",
		v.Regexp(
			"https://workflowy.com/downloads/windows/",
			h.Re("download/v([0-9.]+)/WorkFlowy"),
		),
		d.HTMLA(
			"https://workflowy.com/downloads/windows/",
			".js--start-download[href*='Installer.exe']",
			"",
		),
	)
	Rule("wox",
		v.GitHubRelease(
			"Wox-launcher/Wox",
			h.Re("v(.+)"),
		),
		d.GitHubRelease(
			"Wox-launcher/Wox",
			h.Re("Wox-[0-9.]+.exe"),
			nil,
		),
	)
	Rule("ynab",
		v.Regexp(
			"http://classic.youneedabudget.com/download",
			h.Re("_([0-9.]+)_Setup.exe"),
		),
		d.HTMLA(
			"http://classic.youneedabudget.com/download",
			"a[href*='YNAB'][href$='Setup.exe']",
			"",
		),
	)
	Rule("youtube-dl",
		v.GitHubRelease(
			"rg3/youtube-dl",
			h.Re("([0-9.]+)"),
		),
		d.GitHubRelease(
			"rg3/youtube-dl",
			h.Re("youtube-dl.exe"),
			nil,
		),
	)
}
