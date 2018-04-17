package rules

import (
	"errors"
	"strings"

	. "github.com/just-install/just-install-updater-go/jiup/rules/helpers"
)

func init() {
	AddRule(
		"7zip",
		RegexpVersionExtractor(
			"https://7-zip.org/download.html",
			Re("Download 7-Zip ([0-9][0-9].[0-9][0-9])"),
		),
		TemplateDownloadExtractor(
			"https://www.7-zip.org/a/7z{{.VersionN}}.msi",
			"https://www.7-zip.org/a/7z{{.VersionN}}-x64.msi",
		),
	)
	AddRule(
		"anaconda",
		RegexpVersionExtractor(
			"https://www.anaconda.com/download/",
			Re("Anaconda3-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"https://www.anaconda.com/download/",
			true,
			"a[href*='Windows-x86.exe']",
			"a[href*='Windows-x86_64.exe']",
			"href",
			"href",
			Re("(.+Anaconda3-[0-9.]+-Windows-x86.exe)"),
			Re("(.+Anaconda3-[0-9.]+-Windows-x86_64.exe)"),
		),
	)
	AddRule(
		"android-studio-ide",
		RegexpVersionExtractor(
			"https://developer.android.com/studio/index.html",
			Re("Version: ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://developer.android.com/studio/index.html",
			false,
			"a#win-bundle",
			"",
			"href",
			"",
			Re("(.+android-studio-ide-[0-9.]+-windows.exe)"),
			nil,
		),
	)
	AddRule(
		"arduino",
		RegexpVersionExtractor(
			"https://www.arduino.cc/en/Main/Software",
			Re("arduino-([0-9.]+)-"),
		),
		TemplateDownloadExtractor(
			"https://downloads.arduino.cc/arduino-{{.Version}}-windows.exe",
			"",
		),
	)
	AddRule(
		"audacity",
		RegexpVersionExtractor(
			"http://www.oldfoss.com/Audacity.html",
			Re("audacity-win-([0-9.]+).exe"),
		),
		func(version string) (string, *string, error) {
			return RegexpDownloadExtractor(
				"http://www.oldfoss.com/Audacity.html",
				Re("\"(http.+audacity-win-"+version+".exe)\""),
				nil,
			)(version)
		},
	)
	AddRule(
		"bleachbit",
		RegexpVersionExtractor(
			"https://www.bleachbit.org/download/windows",
			Re("BleachBit-([0-9.]+)-setup.exe"),
		),
		TemplateDownloadExtractor(
			"https://download.bleachbit.org/BleachBit-{{.Version}}-setup.exe",
			"",
		),
	)
	AddRule(
		"bcuninstaller",
		GitHubReleaseVersionExtractor(
			"Klocman",
			"Bulk-Crap-Uninstaller",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Klocman",
			"Bulk-Crap-Uninstaller",
			Re(".*setup.exe"),
			nil,
		),
	)
	AddRule(
		"bitpay",
		GitHubReleaseVersionExtractor(
			"bitpay",
			"copay",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"bitpay",
			"copay",
			Re("BitPay.exe"),
			nil,
		),
	)
	AddRule(
		"brackets",
		GitHubReleaseVersionExtractor(
			"adobe",
			"brackets",
			Re("release-(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"adobe",
			"brackets",
			Re("Brackets.Release.*.msi"),
			nil,
		),
	)
	AddRule(
		"ccleaner",
		func() (string, error) {
			version, err := RegexpVersionExtractor(
				"https://www.ccleaner.com/ccleaner/download/standard",
				Re("ccsetup([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}
			return string(version[0]) + "." + string(version[1:]), nil
		},
		HTMLDownloadExtractor(
			"https://www.ccleaner.com/ccleaner/download/standard",
			false,
			"a[href$='.exe']:contains('start the download')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"cdburnerxp",
		RegexpVersionExtractor(
			"https://download.cdburnerxp.se/msi/",
			Re("_([0-9.]+).msi"),
		),
		HTMLDownloadExtractor(
			"https://download.cdburnerxp.se/msi/",
			true,
			"a[href^='cdbxp_setup_'][href$='msi']:not([href~='x64'])",
			"a[href^='cdbxp_setup_x64'][href$='msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"classic-shell",
		UnderscoreToDot(RegexpVersionExtractor(
			"http://www.oldfoss.com/Classic-Shell.html",
			Re("ClassicShellSetup_([0-9_]+)"),
		)),
		func(version string) (string, *string, error) {
			return RegexpDownloadExtractor(
				"http://www.oldfoss.com/Classic-Shell.html",
				Re("\"(http.+ClassicShellSetup_"+strings.Replace(version, ".", "_", -1)+".exe)\""),
				nil,
			)(version)
		},
	)
	AddRule(
		"clementine-player",
		GitHubReleaseVersionExtractor(
			"clementine-player",
			"Clementine",
			Re("(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"clementine-player",
			"Clementine",
			Re("ClementineSetup-.*.exe"),
			nil,
		),
	)
	AddRule(
		"cmake",
		RegexpVersionExtractor(
			"https://cmake.org/download/",
			Re("Latest Release \\(([0-9.]+)\\)"),
		),
		HTMLDownloadExtractor(
			"https://cmake.org/download/",
			true,
			"a[href$='-win32-x86.msi']",
			"a[href$='-win64-x64.msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"colemak",
		RegexpVersionExtractor(
			"https://colemak.com/Windows",
			Re("Colemak-([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://colemak.com/Windows",
			false,
			"a:contains('Download now')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"conemu",
		GitHubReleaseVersionExtractor(
			"Maximus5",
			"ConEmu",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Maximus5",
			"ConEmu",
			Re("ConEmuSetup.*.exe"),
			nil,
		),
	)
	AddRule(
		"cpu-z",
		RegexpVersionExtractor(
			"https://www.cpuid.com/softwares/cpu-z.html",
			Re("Version ([0-9.]+) for [Ww]indows"),
		),
		TemplateDownloadExtractor(
			"http://download.cpuid.com/cpu-z/cpu-z_{{.Version}}-en.exe",
			"",
		),
	)
	AddRule(
		"crashplan",
		RegexpVersionExtractor(
			"https://www.crashplan.com/shared/js/cp.download.js",
			Re("CPC_CLIENT_VERSION ?= ?'([0-9.]+)'"),
		),
		TemplateDownloadExtractor(
			"https://download.code42.com/installs/win/install/CrashPlan/jre/CrashPlan_{{.Version}}_Win.msi",
			"https://download.code42.com/installs/win/install/CrashPlan/jre/CrashPlan_{{.Version}}_Win64.msi",
		),
	)
	AddRule(
		"cryptomator",
		HTMLVersionExtractor(
			"https://cryptomator.org/downloads",
			"meta[itemprop='softwareVersion']",
			"content",
			nil,
		),
		HTMLDownloadExtractor(
			"https://cryptomator.org/downloads",
			true,
			"#winDownload a[href$='.exe']:contains('32 Bit')",
			"#winDownload a[href$='.exe']:contains('64 Bit')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"crystaldisk-info",
		UnderscoreToDot(HTMLVersionExtractor(
			"https://osdn.net/projects/crystaldiskinfo/releases/",
			"a.pref-download-btn.pref-download-btn-win32[href]",
			"href",
			Re("CrystalDiskInfo([0-9_]+).zip"),
		)),
		func(version string) (string, *string, error) {
			vu := strings.Replace(version, ".", "_", -1)
			dlp, err := HTMLVersionExtractor(
				"https://osdn.net/projects/crystaldiskinfo/releases/",
				"a.pref-download-btn.pref-download-btn-win32[href]",
				"href",
				Re("downloads/([0-9]+/CrystalDiskInfo"+vu+").zip"),
			)()
			if err != nil {
				return "", nil, err
			}
			return "http://osdn.dl.osdn.jp/crystaldiskinfo/" + dlp + ".exe", nil, nil
		},
	)
	AddRule(
		"crystaldisk-mark",
		UnderscoreToDot(HTMLVersionExtractor(
			"https://osdn.net/projects/crystaldiskmark/releases/",
			"a.pref-download-btn.pref-download-btn-win32[href]",
			"href",
			Re("CrystalDiskMark([0-9_]+).zip"),
		)),
		func(version string) (string, *string, error) {
			vu := strings.Replace(version, ".", "_", -1)
			dlp, err := HTMLVersionExtractor(
				"https://osdn.net/projects/crystaldiskmark/releases/",
				"a.pref-download-btn.pref-download-btn-win32[href]",
				"href",
				Re("downloads/([0-9]+/CrystalDiskMark"+vu+").zip"),
			)()
			if err != nil {
				return "", nil, err
			}
			return "http://osdn.dl.osdn.jp/crystaldiskmark/" + dlp + ".exe", nil, nil
		},
	)
	AddRule(
		"dbeaver",
		GitHubReleaseVersionExtractor(
			"dbeaver",
			"dbeaver",
			Re("(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"dbeaver",
			"dbeaver",
			Re("dbeaver-ce-.+-x86-setup.exe"),
			Re("dbeaver-ce-.+-x86_64-setup.exe"),
		),
	)
	AddRule(
		"defraggler",
		func() (string, error) {
			version, err := RegexpVersionExtractor(
				"https://www.ccleaner.com/defraggler/download/standard",
				Re("dfsetup([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}
			return string(version[0]) + "." + string(version[1:]), nil
		},
		HTMLDownloadExtractor(
			"https://www.ccleaner.com/defraggler/download/standard",
			false,
			"a[href$='.exe']:contains('start the download')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"deluge",
		RegexpVersionExtractor(
			"https://dev.deluge-torrent.org/wiki/Download",
			Re("Latest Release: <strong>([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"http://download.deluge-torrent.org/windows/deluge-{{.Version}}-win32-py2.7.exe",
			"",
		),
	)
	AddRule(
		"dependency-walker",
		RegexpVersionExtractor(
			"http://www.dependencywalker.com",
			Re("Dependency Walker ([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"http://www.dependencywalker.com/depends{{.VersionN}}_x86.zip",
			"http://www.dependencywalker.com/depends{{.VersionN}}_x64.zip",
		),
	)
	AddRule(
		"deskpins",
		RegexpVersionExtractor(
			"https://efotinis.neocities.org/deskpins/",
			Re("v([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://efotinis.neocities.org/deskpins/",
			false,
			"a[href*='DeskPins-'][href$='-setup.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"ditto",
		RegexpVersionExtractor(
			"http://ditto-cp.sourceforge.net/index.php",
			Re("versionDots ?= ?\"([0-9.]+)\""),
		),
		TemplateDownloadExtractor(
			"https://sourceforge.net/projects/ditto-cp/files/Ditto/{{.Version}}/DittoSetup_{{.VersionU}}.exe/download",
			"https://sourceforge.net/projects/ditto-cp/files/Ditto/{{.Version}}/DittoSetup_64bit_{{.VersionU}}.exe/download",
		),
	)
	AddRule(
		"doublecmd",
		RegexpVersionExtractor(
			"https://sourceforge.net/p/doublecmd/wiki/Download/",
			Re("doublecmd-([0-9.]+)\\."),
		),
		HTMLDownloadExtractor(
			"https://sourceforge.net/p/doublecmd/wiki/Download/",
			true,
			"a[href$='i386-win32.msi/download']",
			"a[href$='x86_64-win64.msi/download']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"eac",
		RegexpVersionExtractor(
			"http://www.exactaudiocopy.de/en/index.php/weitere-seiten/download-from-alternative-servers-2/",
			Re("Exact Audio Copy V([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"http://www.exactaudiocopy.de/en/index.php/weitere-seiten/download-from-alternative-servers-2/",
			false,
			"a[href*='eac'][href$='.exe']:contains('Download Installer')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	for _, edition := range []string{
		"committers",
		"cpp",
		"java",
		"jee",
		"php",
	} {
		AddRule(
			"eclipse-"+edition,
			RegexpVersionExtractor(
				"https://eclipse.org/downloads/eclipse-packages/",
				Re("\\(([0-9a-z.]+)\\) +Release"),
			),
			AppendToURL(
				"&r=1",
				HTMLDownloadExtractor(
					"https://eclipse.org/downloads/eclipse-packages/",
					true,
					"a.downloadLink[href*='eclipse-"+edition+"-'][href$='-win32.zip']",
					"a.downloadLink[href*='eclipse-"+edition+"-'][href$='-win32-x86_64.zip']",
					"href",
					"href",
					nil,
					nil,
				),
			),
		)
	}
	AddRule(
		"eig",
		GitHubReleaseVersionExtractor(
			"EvilInsultGenerator",
			"c-sharp-desktop",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"EvilInsultGenerator",
			"c-sharp-desktop",
			Re("EvilInsultGenerator_Setup.exe"),
			nil,
		),
	)
	AddRule(
		"emacs",
		HTMLVersionExtractor(
			"https://ftp.gnu.org/gnu/emacs/windows/?C=M;O=D",
			"a[href*='emacs'][href$='-i686.zip']",
			"href",
			Re("emacs-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"https://ftp.gnu.org/gnu/emacs/windows/?C=M;O=D",
			true,
			"a[href*='emacs'][href$='-i686.zip']",
			"a[href*='emacs'][href$='-x86_64.zip']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"enpass",
		HTMLVersionExtractor(
			"https://www.enpass.io/downloads/",
			"a[href*='Enpass_'][href$='_Setup.exe']",
			"href",
			Re("Enpass_([0-9.]+)_"),
		),
		HTMLDownloadExtractor(
			"https://www.enpass.io/downloads/",
			false,
			"a[href*='Enpass_'][href$='_Setup.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"erlang",
		RegexpVersionExtractor(
			"https://www.erlang.org/downloads/",
			Re("DOWNLOAD\\s+OTP\\s+([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.erlang.org/downloads/",
			true,
			"a[href*='win32'][href$='exe']:contains('Windows 32-bit Binary File')",
			"a[href*='win64'][href$='exe']:contains('Windows 64-bit Binary File')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"etcher",
		GitHubReleaseVersionExtractor(
			"resin-io",
			"etcher",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"resin-io",
			"etcher",
			Re("Etcher-Setup-.+-x86.exe"),
			Re("Etcher-Setup-.+-x64.exe"),
		),
	)
	AddRule(
		"everything-search",
		RegexpVersionExtractor(
			"https://www.voidtools.com/downloads/",
			Re("Download Everything ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.voidtools.com/downloads/",
			true,
			"a[href$='x86-Setup.exe']:contains('Download Installer')",
			"a[href$='x64-Setup.exe']:contains('Download Installer 64-bit')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"filezilla-server",
		UnderscoreToDot(HTMLVersionExtractor(
			"https://download.filezilla-project.org/server/?C=M;O=D",
			"a[href*='FileZilla_Server-'][href$='.exe']",
			"href",
			Re("FileZilla_Server-([0-9_]+)"),
		)),
		HTMLDownloadExtractor(
			"https://download.filezilla-project.org/server/?C=M;O=D",
			false,
			"a[href*='FileZilla_Server-'][href$='.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"flash-player",
		HTMLVersionExtractor(
			"http://get.adobe.com/flashplayer/about/",
			"td:contains('Opera, Chromium-based browsers - PPAPI')+td",
			"innerText",
			Re("([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://fpdownload.macromedia.com/get/flashplayer/pdc/{{.Version}}/install_flash_player_{{.Version0}}_plugin.msi",
			"",
		),
	)
	AddRule(
		"flash-player-ie",
		HTMLVersionExtractor(
			"http://get.adobe.com/flashplayer/about/",
			"td:contains('Internet Explorer - ActiveX')+td",
			"innerText",
			Re("([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://fpdownload.macromedia.com/get/flashplayer/pdc/{{.Version}}/install_flash_player_{{.Version0}}_active_x.msi",
			"",
		),
	)
	AddRule(
		"freefilesync",
		RegexpVersionExtractor(
			"https://www.freefilesync.org/download.php",
			Re("Download FreeFileSync ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.freefilesync.org/download.php",
			false,
			"a.direct-download-link[href$='.exe']:contains('Windows Setup')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"geforce-experience",
		RegexpVersionExtractor(
			"https://www.nvidia.com/en-us/geforce/geforce-experience/",
			Re("https://us.download.nvidia.com/GFE/GFEClient/([0-9.]+)/GeForce_Experience"),
		),
		HTMLDownloadExtractor(
			"https://www.nvidia.com/en-us/geforce/geforce-experience/",
			false,
			"a[href^='https://us.download.nvidia.com/GFE/GFEClient/'][href$='.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"gimp",
		RegexpVersionExtractor(
			"https://www.gimp.org/downloads/",
			Re("current stable release of GIMP is <b>([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.gimp.org/downloads/",
			false,
			"#win a[href$='-setup.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"git",
		GitHubReleaseVersionExtractor(
			"git-for-windows",
			"git",
			Re("v([0-9.]+)\\.windows.+"),
		),
		GitHubReleaseDownloadExtractor(
			"git-for-windows",
			"git",
			Re("Git-.+-32-bit.exe"),
			Re("Git-.+-64-bit.exe"),
		),
	)
	AddRule(
		"gitextensions",
		GitHubReleaseVersionExtractor(
			"gitextensions",
			"gitextensions",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"gitextensions",
			"gitextensions",
			Re("GitExtensions-.*-Setup.msi"),
			nil,
		),
	)
	AddRule(
		"git-lfs",
		GitHubReleaseVersionExtractor(
			"git-lfs",
			"git-lfs",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"git-lfs",
			"git-lfs",
			Re("git-lfs-windows-.+.exe"),
			nil,
		),
	)
	AddRule(
		"go",
		RegexpVersionExtractor(
			"https://golang.org/dl/",
			Re("go([0-9.]+)\\.windows"),
		),
		TemplateDownloadExtractor(
			"https://dl.google.com/go/go{{.Version}}.windows-386.msi",
			"https://dl.google.com/go/go{{.Version}}.windows-amd64.msi",
		),
	)
	AddRule(
		"gog-galaxy",
		RegexpVersionExtractor(
			"https://www.gog.com/galaxy",
			Re("setup_galaxy_([0-9.]+).exe"),
		),
		HTMLDownloadExtractor(
			"https://www.gog.com/galaxy",
			false,
			"a[href*='setup'][href$='.exe']:contains('Windows')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"gow",
		GitHubReleaseVersionExtractor(
			"bmatzelle",
			"gow",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"bmatzelle",
			"gow",
			Re("Gow-.+.exe"),
			nil,
		),
	)
	AddRule(
		"greenshot",
		GitHubReleaseVersionExtractor(
			"greenshot",
			"greenshot",
			Re("Greenshot-RELEASE-([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"greenshot",
			"greenshot",
			Re("Greenshot-INSTALLER-.+-RELEASE.exe"),
			nil,
		),
	)
	AddRule(
		"gvim",
		RegexpVersionExtractor(
			"https://www.vim.org/download.php",
			Re("latest version \\(currently ([0-9.]+)\\)"),
		),
		HTMLDownloadExtractor(
			"http://ftp.vim.org/pub/vim/pc/?C=M;O=D",
			false,
			"a[href*='gvim'][href$='.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"hashcheck",
		GitHubReleaseVersionExtractor(
			"gurnec",
			"HashCheck",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"gurnec",
			"HashCheck",
			Re("HashCheckSetup-.+.exe"),
			nil,
		),
	)
	AddRule(
		"heidisql",
		RegexpVersionExtractor(
			"https://www.heidisql.com/download.php",
			Re("HeidiSQL_([0-9.]+)_"),
		),
		HTMLDownloadExtractor(
			"https://www.heidisql.com/download.php",
			false,
			"a[href$='Setup.exe']:contains('Installer')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"hugo",
		GitHubReleaseVersionExtractor(
			"gohugoio",
			"hugo",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"gohugoio",
			"hugo",
			Re("hugo_.+_Windows-32bit.zip"),
			Re("hugo_.+_Windows-64bit.zip"),
		),
	)
	AddRule(
		"imageglass",
		GitHubReleaseVersionExtractor(
			"d2phap",
			"ImageGlass",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"d2phap",
			"ImageGlass",
			Re("ImageGlass_.+.exe"),
			nil,
		),
	)
	AddRule(
		"inkscape",
		RegexpVersionExtractor(
			"https://inkscape.org/en/release/",
			Re("Download Inkscape ([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://media.inkscape.org/dl/resources/file/inkscape-{{.Version}}-x86.msi",
			"https://media.inkscape.org/dl/resources/file/inkscape-{{.Version}}-x64.msi",
		),
	)
	AddRule(
		"jdk",
		RegexpVersionExtractor(
			"https://lv.binarybabel.org/catalog-api/java/jdk8.txt?p=version",
			Re("([0-9a-zA-Z.-]+)"),
		),
		func(version string) (string, *string, error) {
			buf, _, ok, err := GetURL(nil, "https://lv.binarybabel.org/catalog-api/java/jdk8.txt?p=downloads.exe", map[string]string{}, []int{200})
			if err != nil {
				return "", nil, err
			}
			if !ok {
				return "", nil, errors.New("unexpected response code")
			}
			x64 := string(buf)
			x86 := strings.Replace(x64, "x64", "i586", -1)
			return x86, &x64, nil
		},
	)
	AddRule(
		"jre",
		func() (string, error) {
			version, err := RegexpVersionExtractor("https://www.java.com/en/download/manual.jsp", Re("Recommended Version ([0-9]* Update [0-9]*)"))()
			if err != nil {
				return "", err
			}
			return strings.Replace(version, " Update ", ".", 1), nil
		},
		HTMLDownloadExtractor(
			"https://www.java.com/en/download/manual.jsp",
			true,
			"a[title='Download Java software for Windows Offline']",
			"a[title='Download Java software for Windows (64-bit)']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"keepass",
		RegexpVersionExtractor(
			"https://sourceforge.net/projects/keepass/files/",
			Re("KeePass-([0-9.]+)\\.zip"),
		),
		TemplateDownloadExtractor(
			"https://sourceforge.net/projects/keepass/files/KeePass%202.x/{{.Version}}/KeePass-{{.Version}}.msi/download",
			"",
		),
	)
	AddRule(
		"keepassxc",
		GitHubReleaseVersionExtractor(
			"keepassxreboot",
			"keepassxc",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"keepassxreboot",
			"keepassxc",
			Re("KeePassXC-.+-Win32.exe"),
			Re("KeePassXC-.+-Win64.exe"),
		),
	)
	AddRule(
		"keeweb",
		GitHubReleaseVersionExtractor(
			"keeweb",
			"keeweb",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"keeweb",
			"keeweb",
			Re("KeeWeb-.+.win.ia32.exe"),
			Re("KeeWeb-.+.win.x64.exe"),
		),
	)
	AddRule(
		"kicad",
		RegexpVersionExtractor(
			"http://kicad-pcb.org/download/windows/",
			Re("Stable Release Current Version: ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"http://kicad-pcb.org/download/windows/",
			true,
			"a[href$='.exe']:contains('Windows 32-bit')",
			"a[href$='.exe']:contains('Windows 64-bit')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"kodi",
		RegexpVersionExtractor(
			"http://mirrors.kodi.tv/releases/windows/win32/?C=M&O=D",
			Re("kodi-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"http://mirrors.kodi.tv/releases/windows/win32/?C=M&O=D",
			false,
			"a[href*='kodi'][href$='-x86.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"libreoffice",
		RegexpVersionExtractor(
			"https://www.libreoffice.org/download/libreoffice-fresh/?type=win-x86&lang=en-US",
			Re("LibreOffice ([0-9.]+) "),
		),
		TemplateDownloadExtractor(
			"https://download.documentfoundation.org/libreoffice/stable/{{.Version}}/win/x86/LibreOffice_{{.Version}}_Win_x86.msi",
			"https://download.documentfoundation.org/libreoffice/stable/{{.Version}}/win/x86_64/LibreOffice_{{.Version}}_Win_x64.msi",
		),
	)
	AddRule(
		"lockhunter",
		RegexpVersionExtractor(
			"http://lockhunter.com/download.htm",
			Re("Version: ([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"http://lockhunter.com/exe/lockhuntersetup_{{.VersionD}}.exe",
			"",
		),
	)
	AddRule(
		"mercurial",
		RegexpVersionExtractor(
			"https://www.mercurial-scm.org/sources.js",
			Re("windows/mercurial-([0-9.]+)-"),
		),
		RegexpDownloadExtractor(
			"https://www.mercurial-scm.org/sources.js",
			Re("(https://www.mercurial-scm.org/release/windows/mercurial-[0-9.]+-x86.msi)"),
			Re("(https://www.mercurial-scm.org/release/windows/mercurial-[0-9.]+-x64.msi)"),
		),
	)
	AddRule(
		"mono",
		RegexpVersionExtractor(
			"http://www.mono-project.com/download/stable/",
			Re("[0-9.]+ Stable \\(([0-9.]+)\\)"),
		),
		HTMLDownloadExtractor(
			"http://www.mono-project.com/download/stable/",
			false,
			"a[href*='download.mono-project.com'][href*='windows-installer'][href$='.msi']:not([href*='gtksharp'])",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"mountainduck",
		RegexpVersionExtractor(
			"https://mountainduck.io/",
			Re("Installer-([0-9.]+).exe"),
		),
		HTMLDownloadExtractor(
			"https://mountainduck.io/",
			false,
			"a[href*='Installer'][href$='.msi']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"mp3tag",
		RegexpVersionExtractor(
			"https://www.mp3tag.de/en/download.html",
			Re("Mp3tag v([0-9.a-z]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.mp3tag.de/en/dodownload.html",
			false,
			"a[href*='download'][href$='.exe']:contains('here')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"mpc-hc",
		RegexpVersionExtractor(
			"https://mpc-hc.org/downloads/",
			Re("latest stable build is v([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://mpc-hc.org/downloads/",
			true,
			"a[href$='.x86.exe']",
			"a[href$='.x64.exe']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"mumble",
		GitHubReleaseVersionExtractor(
			"mumble-voip",
			"mumble",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"mumble-voip",
			"mumble",
			Re("mumble-.+.msi"),
			nil,
		),
	)
	AddRule(
		"naps2",
		GitHubReleaseVersionExtractor(
			"cyanfish",
			"naps2",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"cyanfish",
			"naps2",
			Re("naps2-.+-setup.msi"),
			nil,
		),
	)
	AddRule(
		"nextcloud",
		HTMLVersionExtractor(
			"https://nextcloud.com/install/",
			"#tab-desktop a[href*='desktop/releases/Windows'][href$='setup.exe']",
			"href",
			Re("Nextcloud-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"https://nextcloud.com/install/",
			false,
			"#tab-desktop a[href*='desktop/releases/Windows'][href$='setup.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"node",
		RegexpVersionExtractor(
			"https://nodejs.org/en/download/current/",
			Re("Latest Current Version: <strong>([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://nodejs.org/en/download/current/",
			true,
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('32-bit')",
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('64-bit')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"node-lts",
		RegexpVersionExtractor(
			"https://nodejs.org/en/download/",
			Re("Latest LTS Version: <strong>([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://nodejs.org/en/download/",
			true,
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('32-bit')",
			"th:contains('Windows Installer (.msi)') ~ td>a:contains('64-bit')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"notepad++",
		RegexpVersionExtractor(
			"https://notepad-plus-plus.org/download/",
			Re("Download Notepad\\+\\+ ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://notepad-plus-plus.org/download/",
			true,
			"a[href*='npp'][href$='nstaller.exe']",
			"a[href*='npp'][href$='nstaller.x64.exe']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"notepad2-mod",
		GitHubReleaseVersionExtractor(
			"XhmikosR",
			"notepad2-mod",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"XhmikosR",
			"notepad2-mod",
			Re("Notepad2-mod..+.exe"),
			nil,
		),
	)
	AddRule(
		"npackd",
		GitHubReleaseVersionExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("version_([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("Npackd32-.+.msi"),
			Re("Npackd64-.+.msi"),
		),
	)
	AddRule(
		"npackdcl",
		GitHubReleaseVersionExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("version_([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"tim-lebedkov",
			"npackd-cpp",
			Re("NpackdCL-.+.msi"),
			nil,
		),
	)
	AddRule(
		"nxlog",
		RegexpVersionExtractor(
			"https://nxlog.co/products/nxlog-community-edition/download",
			Re("nxlog-ce-([0-9.]+)\\.msi"),
		),
		HTMLDownloadExtractor(
			"https://nxlog.co/products/nxlog-community-edition/download",
			false,
			"a[href*='nxlog-ce-'][href$='.msi']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"obs-studio",
		RegexpVersionExtractor(
			"https://obsproject.com/download",
			Re("download/([0-9.]+)/OBS"),
		),
		HTMLDownloadExtractor(
			"https://obsproject.com/download",
			false,
			"a[href*='OBS-Studio-'][href$='Full-Installer.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"octave",
		RegexpVersionExtractor(
			"https://ftp.gnu.org/gnu/octave/windows/?C=M;O=D",
			Re("octave-([0-9.]+)-w32-installer.exe"),
		),
		HTMLDownloadExtractor(
			"https://ftp.gnu.org/gnu/octave/windows/?C=M;O=D",
			true,
			"a[href*='octave-'][href$='-w32-installer.exe']",
			"a[href*='octave-'][href$='-w64-installer.exe']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"open-hardware-monitor",
		RegexpVersionExtractor(
			"http://openhardwaremonitor.org/downloads/",
			Re("Open Hardware Monitor ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"http://openhardwaremonitor.org/downloads/",
			false,
			"a[href*='openhardwaremonitor-'][href$='.zip']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"openssh",
		RegexpVersionExtractor(
			"https://www.mls-software.com/opensshd.html",
			Re("OpenSSH ([0-9.]+)p"),
		),
		HTMLDownloadExtractor(
			"https://www.mls-software.com/opensshd.html",
			false,
			"a[href*='setupssh-'][href$='.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"perl",
		HTMLVersionExtractor(
			"http://strawberryperl.com/releases.html",
			"a[href*='strawberry-perl-'][href$='32bit.msi']",
			"href",
			Re("strawberry-perl-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"http://strawberryperl.com/releases.html",
			true,
			"a[href*='strawberry-perl-'][href$='32bit.msi']",
			"a[href*='strawberry-perl-'][href$='64bit.msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"pia",
		RegexpVersionExtractor(
			"https://www.privateinternetaccess.com/pages/downloads",
			Re("Clients v([0-9]+) Released"),
		),
		HTMLDownloadExtractor(
			"https://www.privateinternetaccess.com/pages/downloads",
			false,
			"a[href*='pia-'][href$='installer-win.exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"plex-media-server",
		RegexpVersionExtractor(
			"https://plex.tv/api/downloads/1.json",
			Re("version\":\"([0-9.]+)"),
		),
		RegexpDownloadExtractor(
			"https://plex.tv/api/downloads/1.json",
			Re("\"(https://downloads.plex.tv/plex-media-server/[0-9a-z.-]+?/Plex-Media-Server-[0-9a-z.-]+?.exe)\""),
			nil,
		),
	)
	AddRule(
		"processhacker",
		GitHubReleaseVersionExtractor(
			"processhacker",
			"processhacker",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"processhacker",
			"processhacker",
			Re("processhacker-.+-setup.exe"),
			nil,
		),
	)
	AddRule(
		"putty",
		RegexpVersionExtractor(
			"http://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html",
			Re("latest release \\(([0-9.]+)\\)"),
		),
		HTMLDownloadExtractor(
			"http://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html",
			true,
			"span.downloadfile a[href^='https'][href*='w32/putty'][href$='.msi']",
			"span.downloadfile a[href^='https'][href*='w64/putty'][href$='.msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"pycharm-community",
		RegexpVersionExtractor(
			"https://data.services.jetbrains.com/products/releases?code=PCP%2CPCC&latest=true",
			Re("version\":\"([0-9.]+)"),
		),
		RegexpDownloadExtractor(
			"https://data.services.jetbrains.com/products/releases?code=PCP%2CPCC&latest=true",
			Re("\"(https://download.jetbrains.com/python/pycharm-community-[0-9.]+.exe)\""),
			nil,
		),
	)
	AddRule(
		"python2",
		HTMLVersionExtractor(
			"https://www.python.org/downloads/",
			".download-for-current-os .download-os-windows a[href*='python-2']",
			"innerText",
			Re("Download Python ([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}.msi",
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}.amd64.msi",
		),
	)
	AddRule(
		"python2-yaml",
		GitHubTagVersionExtractor(
			"yaml",
			"pyyaml",
			Re("([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://pyyaml.org/download/pyyaml/PyYAML-{{.Version}}.win32-py2.7.exe",
			"https://pyyaml.org/download/pyyaml/PyYAML-{{.Version}}.win-amd64-py2.7.exe",
		),
	)
	AddRule(
		"python2-win32",
		GitHubReleaseVersionExtractor(
			"mhammond",
			"pywin32",
			Re("b(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"mhammond",
			"pywin32",
			Re("pywin32-.+.win32-py2.7.exe"),
			Re("pywin32-.+.win-amd64-py2.7.exe"),
		),
	)
	AddRule(
		"python3",
		HTMLVersionExtractor(
			"https://www.python.org/downloads/",
			".download-for-current-os .download-os-windows a[href*='python-3']",
			"innerText",
			Re("Download Python ([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}.exe",
			"https://www.python.org/ftp/python/{{.Version}}/python-{{.Version}}-amd64.exe",
		),
	)
	AddRule(
		"qbittorrent",
		RegexpVersionExtractor(
			"http://www.oldfoss.com/qBittorrent.html",
			Re("qbittorrent_([0-9.]+)_setup.exe"),
		),
		func(version string) (string, *string, error) {
			return RegexpDownloadExtractor(
				"http://www.oldfoss.com/qBittorrent.html",
				Re("\"(http.+qbittorrent_"+version+"_setup.exe)\""),
				Re("\"(http.+qbittorrent_"+version+"_x64_setup.exe)\""),
			)(version)
		},
	)
	AddRule(
		"recuva",
		func() (string, error) {
			version, err := RegexpVersionExtractor(
				"https://www.ccleaner.com/recuva/download/standard",
				Re("rcsetup([0-9]+)"),
			)()
			if err != nil {
				return "", err
			}
			return string(version[0]) + "." + string(version[1:]), nil
		},
		HTMLDownloadExtractor(
			"https://www.ccleaner.com/recuva/download/standard",
			false,
			"a[href$='.exe']:contains('start the download')",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"ruby",
		GitHubReleaseVersionExtractor(
			"oneclick",
			"rubyinstaller2",
			Re("rubyinstaller-([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"oneclick",
			"rubyinstaller2",
			Re("rubyinstaller-[0-9.]+-.+-x86.exe"),
			Re("rubyinstaller-[0-9.]+-.+-x64.exe"),
		),
	)
	AddRule(
		"seafile-client",
		NoHTTPSForVersionExtractor(HTMLVersionExtractor(
			"https://www.seafile.com/en/download/",
			".txt > h3:contains('Client for Windows')~a[href*='seafile'][href$='en.msi'].download-op",
			"innerText",
			Re("([0-9.]+)"),
		)),
		NoHTTPSForDownloadExtractor(HTMLDownloadExtractor(
			"https://www.seafile.com/en/download/",
			false,
			".txt > h3:contains('Client for Windows')~a[href*='seafile'][href$='en.msi'].download-op",
			"",
			"href",
			"",
			nil,
			nil,
		)),
	)
	AddRule(
		"sharex",
		GitHubReleaseVersionExtractor(
			"ShareX",
			"ShareX",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"ShareX",
			"ShareX",
			Re("ShareX-.+-setup.exe"),
			nil,
		),
	)
	AddRule(
		"signal",
		RegexpVersionExtractor(
			"https://updates.signal.org/desktop/latest.yml",
			Re("version: ([0-9.]+)"),
		),
		TemplateDownloadExtractor(
			"https://updates.signal.org/desktop/signal-desktop-win-{{.Version}}.exe",
			"",
		),
	)
	AddRule(
		"simplenote",
		GitHubReleaseVersionExtractor(
			"Automattic",
			"simplenote-electron",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Automattic",
			"simplenote-electron",
			Re("Simplenote-windows-.+.exe"),
			nil,
		),
	)
	AddRule(
		"sharpkeys",
		GitHubReleaseVersionExtractor(
			"randyrants",
			"sharpkeys",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"randyrants",
			"sharpkeys",
			Re("sharpkeys.+.msi"),
			nil,
		),
	)
	AddRule(
		"smplayer",
		RegexpVersionExtractor(
			"https://sourceforge.net/projects/smplayer/files/",
			Re("smplayer-([0-9.]+)\\.tar"),
		),
		TemplateDownloadExtractor(
			"https://sourceforge.net/projects/smplayer/files/SMPlayer/{{.Version}}/smplayer-{{.Version}}-win32.exe/download",
			"https://sourceforge.net/projects/smplayer/files/SMPlayer/{{.Version}}/smplayer-{{.Version}}-x64.exe/download",
		),
	)
	AddRule(
		"sourcetree",
		RegexpVersionExtractor(
			"https://www.sourcetreeapp.com",
			Re("SourceTreeSetup-([0-9.]+)\\.exe"),
		),
		HTMLDownloadExtractor(
			"https://www.sourcetreeapp.com",
			false,
			"a[href*='SourceTreeSetup'][href$='exe']",
			"",
			"href",
			"",
			nil,
			nil,
		),
	)
	AddRule(
		"sublime-text",
		RegexpVersionExtractor(
			"https://www.sublimetext.com/2",
			Re("Version:</i> ([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.sublimetext.com/2",
			true,
			"#dl_win_32 a[href$='exe']",
			"#dl_win_64 a[href$='exe']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"sublime-text-3",
		RegexpVersionExtractor(
			"https://www.sublimetext.com/3",
			Re("Version:</i> Build ([0-9]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.sublimetext.com/3",
			true,
			"#dl_win_32 a[href$='exe']",
			"#dl_win_64 a[href$='exe']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"sublime-text-dev",
		RegexpVersionExtractor(
			"https://www.sublimetext.com/3dev",
			Re("Version:</i> Build ([0-9]+)"),
		),
		HTMLDownloadExtractor(
			"https://www.sublimetext.com/3dev",
			true,
			"#dl_win_32 a[href$='exe']",
			"#dl_win_64 a[href$='exe']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"subversion",
		RegexpVersionExtractor(
			"https://sliksvn.com/download/",
			Re("Subversion-([0-9.]+)-"),
		),
		HTMLDownloadExtractor(
			"https://sliksvn.com/download/",
			true,
			".client a[href$='zip']:contains('32 bit')",
			".client a[href$='zip']:contains('64 bit')",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"sumatrapdf",
		RegexpVersionExtractor(
			"https://www.sumatrapdfreader.org/news.html",
			Re(">([0-9.]+) \\(20"),
		),
		TemplateDownloadExtractor(
			"https://www.sumatrapdfreader.org/dl/SumatraPDF-{{.Version}}-install.exe",
			"https://www.sumatrapdfreader.org/dl/SumatraPDF-{{.Version}}-64-install.exe",
		),
	)
	AddRule(
		"syncthing",
		GitHubReleaseVersionExtractor(
			"syncthing",
			"syncthing",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"syncthing",
			"syncthing",
			Re(".*windows-386.*.zip"),
			Re(".*windows-amd64.*.zip"),
		),
	)
	AddRule(
		"tortoisegit",
		RegexpVersionExtractor(
			"https://tortoisegit.org/download/",
			Re("TortoiseGit-([0-9.]+)"),
		),
		HTMLDownloadExtractor(
			"https://tortoisegit.org/download/",
			true,
			"a[href$='32bit.msi']",
			"a[href$='64bit.msi']",
			"href",
			"href",
			nil,
			nil,
		),
	)
	AddRule(
		"tortoisesvn",
		RegexpVersionExtractor(
			"https://tortoisesvn.net/downloads.html",
			Re("The current version is ([0-9.]+)"),
		),
		func(version string) (string, *string, error) {
			// Layer 1: Link to OSDN
			x86, x64, err := HTMLDownloadExtractor(
				"https://tortoisesvn.net/downloads.html",
				true,
				"a[href^='https://osdn.net'][href*='win32-svn']",
				"a[href^='https://osdn.net'][href*='x64-svn']",
				"href",
				"href",
				nil,
				nil,
			)(version)
			if err != nil {
				return "", nil, err
			}
			if x64 == nil {
				return "", nil, errors.New("x64 link empty")
			}
			// Layer 2: OSDN to redir link
			x86, _, err = HTMLDownloadExtractor(
				x86,
				false,
				"a.mirror_link[href*='/frs/redir'][href*='win32-svn']",
				"",
				"href",
				"",
				nil,
				nil,
			)(version)
			if err != nil {
				return "", nil, err
			}
			_, x64, err = HTMLDownloadExtractor(
				*x64,
				true,
				"a",
				"a.mirror_link[href*='/frs/redir'][href*='x64-svn']",
				"href",
				"href",
				nil,
				nil,
			)(version)
			if err != nil {
				return "", nil, err
			}
			return x86, x64, nil
		},
	)
	AddRule(
		"upx",
		GitHubReleaseVersionExtractor(
			"upx",
			"upx",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"upx",
			"upx",
			Re("upx[0-9]+w.zip"),
			nil,
		),
	)
	AddRule(
		"webtorrent",
		GitHubReleaseVersionExtractor(
			"webtorrent",
			"webtorrent-desktop",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"webtorrent",
			"webtorrent-desktop",
			Re("WebTorrentSetup-v[0-9.]+-ia32.exe"),
			Re("WebTorrentSetup-v[0-9.]+.exe"),
		),
	)
	AddRule(
		"wixedit",
		GitHubReleaseVersionExtractor(
			"WixEdit",
			"WixEdit",
			Re("v([0-9]+\\.[0-9]+\\.[0-9]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"WixEdit",
			"WixEdit",
			Re("wixedit-.+.msi"),
			nil,
		),
	)
	AddRule(
		"workflowy",
		GitHubReleaseVersionExtractor(
			"workflowy",
			"desktop",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"workflowy",
			"desktop",
			Re("WorkFlowy.exe"),
			nil,
		),
	)
	AddRule(
		"wox",
		GitHubReleaseVersionExtractor(
			"Wox-launcher",
			"Wox",
			Re("v(.+)"),
		),
		GitHubReleaseDownloadExtractor(
			"Wox-launcher",
			"Wox",
			Re("Wox-[0-9.]+.exe"),
			nil,
		),
	)
	AddRule(
		"youtube-dl",
		GitHubReleaseVersionExtractor(
			"rg3",
			"youtube-dl",
			Re("([0-9.]+)"),
		),
		GitHubReleaseDownloadExtractor(
			"rg3",
			"youtube-dl",
			Re("youtube-dl.exe"),
			nil,
		),
	)
}
