load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "com_github_alecaivazis_survey_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/AlecAivazis/survey/v2",
        sum = "h1:NvTuVHISgTHEHeBFqt6BHOe4Ny/NwGZr7w+F8S9ziyw=",
        version = "v2.3.6",
    )
    go_repository(
        name = "com_github_atotto_clipboard",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/atotto/clipboard",
        sum = "h1:EH0zSVneZPSuFR11BlR9YppQTVDbh5+16AmcJi4g1z4=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_aymanbagabas_go_osc52_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aymanbagabas/go-osc52/v2",
        sum = "h1:HwpRHbFMcZLEVr42D4p7XBqjyuxQH5SMiErDT4WkJ2k=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_charmbracelet_bubbles",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/bubbles",
        sum = "h1:PYv1A036luoBGroX6VWjQIE9Syf2Wby2oOl/39KLfy0=",
        version = "v0.18.0",
    )
    go_repository(
        name = "com_github_charmbracelet_bubbletea",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/bubbletea",
        sum = "h1:9iVLaZMztMsxYga94vzAd0xclMSSF+A7I1Dq6LFP7O4=",
        version = "v0.25.1-0.20240319021810-0ebbb194aae5",
    )
    go_repository(
        name = "com_github_charmbracelet_harmonica",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/harmonica",
        sum = "h1:8NxJWRWg/bzKqqEaaeFNipOu77YR5t8aSwG4pgaUBiQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_charmbracelet_lipgloss",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/lipgloss",
        sum = "h1:PNyd3jvaJbg4jRHKWXnCj1akQm4rh8dbEzN1p/u1KWg=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_containerd_console",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/console",
        sum = "h1:q2hJAaP1k2wIvVRd/hEHD7lacgqrCPS+k8g1MndzfWY=",
        version = "v1.0.4-0.20230313162750-1ae8d489ac81",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:p1EgwI/C7NhT0JmVkwCD2ZBK8j4aeHQX2pMHHBfMQ6w=",
        version = "v2.0.2",
    )
    go_repository(
        name = "com_github_creack_pty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/creack/pty",
        sum = "h1:QeVUsEDNrLBW4tMgZHvxy18sKtr6VI492kBhUfhDJNI=",
        version = "v1.1.17",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:U9qPSI2PIWSS1VwoXQT9A3Wy9MM3WgvqSxFWenqJduM=",
        version = "v1.1.2-0.20180830191138-d8f796af33cc",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_hinshun_vt10x",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hinshun/vt10x",
        sum = "h1:qv2VnGeEQHchGaZ/u7lxST/RaJw+cv273q79D81Xbog=",
        version = "v0.0.0-20220119200601-820417d04eec",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_kballard_go_shellquote",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kballard/go-shellquote",
        sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
        version = "v0.0.0-20180428030007-95032a82bc51",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_lucasb_eyer_go_colorful",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/lucasb-eyer/go-colorful",
        sum = "h1:1nnpGOrhyZZuNyfu1QjKiUICQ74+3FNCN69Aj6K7nkY=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=",
        version = "v0.0.20",
    )
    go_repository(
        name = "com_github_mattn_go_localereader",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-localereader",
        sum = "h1:ygSAOl7ZXTx4RdPYinUpg6W99U8jWvWi9Ye2JC/oIi4=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:UNAjwbU9l54TA3KzvqLGxwWjHmMgBUVhBiTjelZgg3U=",
        version = "v0.0.15",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:j7+1HpAFS1zy5+Q4qx1fWh90gTKwiN4QCGoY9TWyyO4=",
        version = "v0.0.0-20170206155736-9520e82c474b",
    )
    go_repository(
        name = "com_github_muesli_ansi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/muesli/ansi",
        sum = "h1:ZK8zHtRHOkbHy6Mmr5D264iyp3TiX5OmNcI5cIARiQI=",
        version = "v0.0.0-20230316100256-276c6243b2f6",
    )
    go_repository(
        name = "com_github_muesli_cancelreader",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/muesli/cancelreader",
        sum = "h1:3I4Kt4BQjOR54NavqnDogx/MIoWBFa0StPA8ELUXHmA=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_muesli_reflow",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/muesli/reflow",
        sum = "h1:IFsN6K9NfGtjeggFP+68I4chLZV2yIKsXJFNZ+eWh6s=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_muesli_termenv",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/muesli/termenv",
        sum = "h1:GohcuySI0QmI3wN8Ok9PtKGkgkFIk7y6Vpb5PvrY+Wo=",
        version = "v0.15.2",
    )
    go_repository(
        name = "com_github_netflix_go_expect",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Netflix/go-expect",
        sum = "h1:+vx7roKuyA63nhn5WAunQHLTznkw5W8b1Xc0dNjp83s=",
        version = "v0.0.0-20220104043353-73e0943537d2",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:Jamvg5psRIccs7FGNTlIRMkT8wgtp5eCXdBlqhYGL6U=",
        version = "v1.0.1-0.20181226105442-5d4384ee4fb2",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:Sovz9sDSwbOz9tgUy8JpT+KgCkPYJEN/oYzlJiYTNLg=",
        version = "v0.4.6",
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_sahilm_fuzzy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sahilm/fuzzy",
        sum = "h1:MvTmaQdww/z0Q4wrYjDSCcZ78NoftLQyHBSLW/Cx79Y=",
        version = "v0.1.1-0.20230530133925-c48e322e2a8f",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cobra",
        sum = "h1:hyqWnYt1ZQShIddO5kBpj3vu05/++x6tJ6dg8EC572I=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/objx",
        sum = "h1:4VhoImhV/Bm0ToFkXFi8hXNXwpDRZ/ynw3amt82mzq0=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/testify",
        sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
        version = "v1.8.4",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "org_golang_x_crypto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/crypto",
        sum = "h1:X31++rzVUdKhX5sWmSOFZxx8UW/ldWx55cbf08iNAMA=",
        version = "v0.21.0",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mod",
        sum = "h1:dGoOF9QVLYng8IHTm7BAyWqCqSheQ5pYWGhzW00YJr0=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_net",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/net",
        sum = "h1:AQyQV4dYCvJ7vGmJyKki9+PBdyvhkSd8EIx/qb0AYv4=",
        version = "v0.21.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sync",
        sum = "h1:5BMeUDZ7vkXGfEr1x9B4bRcTH4lpkTkpdh0T/J+qjbQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sys",
        sum = "h1:DBdB3niSjOA/O0blCZBqDefyWNYveAYMNF1Wum0DYQ4=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_term",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/term",
        sum = "h1:FcHjZXDMxI8mM3nwhX9HlKop4C0YQvCVCdwYl2wOtE8=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_text",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/tools",
        sum = "h1:GO788SKMRunPIBCXiQyo2AaexLstOrVhuAL5YwsckQM=",
        version = "v0.16.0",
    )
