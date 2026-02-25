load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "af_inet_netaddr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "inet.af/netaddr",
        sum = "h1:B4dC8ySKTQXasnjDTMsoCMf1sQG4WsMej0WXaHxunmU=",
        version = "v0.0.0-20220617031823-097006376321",
    )
    go_repository(
        name = "cat_dario_mergo",
        build_file_proto_mode = "disable",
        importpath = "dario.cat/mergo",
        sum = "h1:AGCNq9Evsj31mOgNPcLyXc+4PNABt905YmuqPYYpBWk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "co_honnef_go_gotraceui",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "honnef.co/go/gotraceui",
        sum = "h1:dmNsfQ9Vl3GwbiVD7Z8d/osC6WtGGrasyrC2suc4ZIQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "co_honnef_go_tools",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "honnef.co/go/tools",
        sum = "h1:qTakTkI6ni6LFD5sBwwsdSO+AQqbSIxOauHTTQKZ/7o=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_99designs_gqlgen",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/99designs/gqlgen",
        sum = "h1:u/o/rv2SZ9s5280dyUOOrkpIIkr/7kITMXYD3rkJ9go=",
        version = "v0.17.36",
    )
    go_repository(
        name = "com_github_actgardner_gogen_avro_v10",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/actgardner/gogen-avro/v10",
        sum = "h1:z3pOGblRjAJCYpkIJ8CmbMJdksi4rAhaygw0dyXZ930=",
        version = "v10.2.1",
    )
    go_repository(
        name = "com_github_actgardner_gogen_avro_v9",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/actgardner/gogen-avro/v9",
        sum = "h1:YZ5tCwV5xnDZrG4uRDQYT2VAWZCRAG3eyQH/WYR2T6Q=",
        version = "v9.1.0",
    )
    go_repository(
        name = "com_github_adalogics_go_fuzz_headers",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/AdaLogics/go-fuzz-headers",
        sum = "h1:V8krnnfGj4pV65YLUm3C0/8bl7V5Nry2Pwvy3ru/wLc=",
        version = "v0.0.0-20210715213245-6c3934b029d8",
    )
    go_repository(
        name = "com_github_agiledragon_gomonkey_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/agiledragon/gomonkey/v2",
        sum = "h1:k+UnUY0EMNYUFUAQVETGY9uUTxjMdnUkP0ARyJS1zzs=",
        version = "v2.3.1",
    )
    go_repository(
        name = "com_github_agnivade_levenshtein",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/agnivade/levenshtein",
        sum = "h1:QY8M92nrzkmr798gCo3kmMyqXFzdQVpxLlGPRBij0P8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_ajg_form",
        build_file_proto_mode = "disable",
        importpath = "github.com/ajg/form",
        sum = "h1:t9c7v8JUKu/XxOGBU0yjNpaMloxGEJhUkqFRq0ibGeU=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_ajstarks_deck",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ajstarks/deck",
        sum = "h1:7kQgkwGRoLzC9K0oyXdJo7nve/bynv/KwUsxbiTlzAM=",
        version = "v0.0.0-20200831202436-30c9fc6549a9",
    )
    go_repository(
        name = "com_github_ajstarks_deck_generate",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ajstarks/deck/generate",
        sum = "h1:iXUgAaqDcIUGbRoy2TdeofRG/j1zpGRSEmNK05T+bi8=",
        version = "v0.0.0-20210309230005-c3f852c02e19",
    )
    go_repository(
        name = "com_github_ajstarks_svgo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ajstarks/svgo",
        sum = "h1:slYM766cy2nI3BwyRiyQj/Ud48djTMtMebDqepE95rw=",
        version = "v0.0.0-20211024235047-1546f124cd8b",
    )
    go_repository(
        name = "com_github_alecthomas_assert_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/assert/v2",
        sum = "h1:mAsH2wmvjsuvyBvAmCtm7zFsBlb8mIHx5ySLVdDZXL0=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_participle_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/participle/v2",
        sum = "h1:z7dElHRrOEEq45F2TG5cbQihMtNTv8vwldytDj7Wrz4=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_alecthomas_repr",
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/repr",
        sum = "h1:HAzS41CIzNW5syS8Mf9UwXhNH1J9aix/BvDRf1Ml2Yk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_alecthomas_template",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/template",
        sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
        version = "v0.0.0-20190718012654-fb15b899a751",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/alecthomas/units",
        sum = "h1:t3eaIm0rUkzbrIewtiFmMK5RXHej2XnoXNhxVsAYUfg=",
        version = "v0.0.0-20240626203959-61d1e3462e30",
    )
    go_repository(
        name = "com_github_alexflint_go_filemutex",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/alexflint/go-filemutex",
        sum = "h1:IAWuUuRYL2hETx5b8vCgwnD+xSdlsTQY6s2JjBsqLdg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_alicebob_miniredis_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/alicebob/miniredis/v2",
        sum = "h1:QwLphYqCEAo1eu1TqPRN2jgVMPBweeQcR21jeqDCONI=",
        version = "v2.35.0",
    )
    go_repository(
        name = "com_github_andreyvit_diff",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/andreyvit/diff",
        sum = "h1:bvNMNQO63//z+xNgfBlViaCIJKLlCJ6/fmUseuG0wVQ=",
        version = "v0.0.0-20170406064948-c7f18ee00883",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:Yf9fFpf49Zrxb9NlQaluyE92/+X7UVHlhMNJN2sxfOI=",
        version = "v1.0.6",
    )
    go_repository(
        name = "com_github_andybalholm_cascadia",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/andybalholm/cascadia",
        sum = "h1:BuuO6sSfQNFRu1LppgbD25Hr2vLYW25JvxHs5zzsLTo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_ansel1_merry",
        build_file_proto_mode = "disable",
        importpath = "github.com/ansel1/merry",
        sum = "h1:0xr40haRrfVzmOH/JVOu7KOKGEI1c/7q5EmgTEbn+Ng=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_ansel1_merry_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/ansel1/merry/v2",
        sum = "h1:WeiKZdslHPAPFYxTtgX7clC2Vh75NCoWs5OjCZbIA0A=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_ansel1_vespucci_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/ansel1/vespucci/v4",
        sum = "h1:CrKIMZgkGBdjF1ekTQvxFkcrH3tt1W2He2+tg3qXXbU=",
        version = "v4.1.1",
    )
    go_repository(
        name = "com_github_antihax_optional",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v15",
        build_file_proto_mode = "disable",
        importpath = "github.com/apache/arrow/go/v15",
        sum = "h1:60IliRbiyTWCWjERBCkO1W4Qun9svcYoZrSLcyOsMLE=",
        version = "v15.0.2",
    )
    go_repository(
        name = "com_github_apache_thrift",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/apache/thrift",
        sum = "h1:cMd2aj52n+8VoAtvSvLn4kDC3aZ6IAkBuqWQ2IDu7wo=",
        version = "v0.17.0",
    )
    go_repository(
        name = "com_github_arbovm_levenshtein",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/arbovm/levenshtein",
        sum = "h1:jfIu9sQUG6Ig+0+Ap1h4unLjW6YQJpKZVmUzxsD4E/Q=",
        version = "v0.0.0-20160628152529-48b4e1c0c4d0",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/circbuf",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version = "v0.0.0-20150827004946-bbbad097214e",
    )
    go_repository(
        name = "com_github_armon_consul_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/consul-api",
        sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
        version = "v0.0.0-20180202201655-eb2c6b5be1b6",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:hR91U9KYmb6bLBYLQjyM+3j+rcd/UhE+G78SFnF8gJA=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-radix",
        sum = "h1:F4z6KzEeeQIMeLFa97iZU6vupzoecKdU5TX24SNppXI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_armon_go_socks5",
        build_file_proto_mode = "disable",
        importpath = "github.com/armon/go-socks5",
        sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
        version = "v0.0.0-20160902184237-e75332964ef5",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:DklsrG3dyBCFEj5IhUbnKptjxatkF07cF2ak3yi77so=",
        version = "v0.0.0-20230301143203-a9d515a09cc2",
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
        name = "com_github_auth0_go_jwt_middleware_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/auth0/go-jwt-middleware/v2",
        sum = "h1:pqxEIwlCztD0T9ZygGfOrw4NK/F9iotnCnPJVADKbkE=",
        version = "v2.2.1",
    )
    go_repository(
        name = "com_github_avast_retry_go_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/avast/retry-go/v4",
        sum = "h1:cqI48aXx0BExKoM7XPklDpoHAg7/srPPLAfWG5z62jo=",
        version = "v4.3.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:KKUZBfBoyqy5d3swXyiC7Q76ic40rYcbqH7qjh59kzU=",
        version = "v1.55.5",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:0OF9RiEMEdDdZEMqF9MRjevyxAQcf6gY+E7vwBILFj0=",
        version = "v1.36.5",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
        sum = "h1:12SpdwU8Djs+YGklkinSSlcrPyj3H4VifVsKf78KbwA=",
        version = "v1.6.11",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:r+X1x8QI6FEPdJDWCNBDZHyAcyFwSjHN8q8uuus+Axs=",
        version = "v1.25.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:rv1V3kIJ14pdmTu01hwcMJ0WAERensSiD9rEWEBb1Tk=",
        version = "v1.17.63",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:x793wxmUWVDhshP8WW2mlnXuFrO4cOd3HLBroh1paFw=",
        version = "v1.16.30",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:SsytQyTMHMDPspp+spo7XwXTP44aJZZAC7fBV2C5+5s=",
        version = "v1.3.36",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:i2vNHQiXUvKhs3quBR6aqlgJaiaexz/aNvdCktW/kAM=",
        version = "v2.6.36",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_ini",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/ini",
        sum = "h1:n3GDfwqF2tzEkXlv5cuy4iy7LpKDtqDMcNLfZDu9rls=",
        version = "v1.7.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
        sum = "h1:GMYy2EOWfzdP3wfVAGXBNKY5vK4K8vMET4sYOYltmqs=",
        version = "v1.3.36",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_dynamodb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/dynamodb",
        sum = "h1:x3V1JRHq7q9RUbDpaeNpLH7QoipGpCo3fdnMMuSeABU=",
        version = "v1.21.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ec2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ec2",
        sum = "h1:c6a19AjfhEXKlEX63cnlWtSQ4nzENihHZOG0I3wH6BE=",
        version = "v1.93.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_eventbridge",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/eventbridge",
        sum = "h1:G18wotYZxZ0A5tkqKv6FHCjsF86UQrqNHy5LS+T7JWM=",
        version = "v1.20.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
        sum = "h1:CXV68E2dNqhuynZJPB80bhPQwAKqBWVer887figW6Jc=",
        version = "v1.12.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
        sum = "h1:nAP2GYbfh8dd2zGZqFRSMlq+/F6cMPBUuCsGAMkN074=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_endpoint_discovery",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery",
        sum = "h1:JlxVMFDHivlhNOIxd2O/9z4O0wC2zIC4lRB71lejVHU=",
        version = "v1.7.34",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:t0E6FzREdtCsiLIoLCWsYliNsRBgyGD/MCK571qk4MI=",
        version = "v1.12.17",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
        sum = "h1:qcLWgdhq45sDM9na4cvXax9dyLitn8EYBRl8Ak4XtG4=",
        version = "v1.18.17",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_kinesis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/kinesis",
        sum = "h1:UohaQds+Puk9BEbvncXkZduIGYImxohbFpVmSoymXck=",
        version = "v1.18.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_s3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
        sum = "h1:5Y75q0RPQoAbieyOuGLhjV9P3txvYgXv2lg0UwJOfmE=",
        version = "v1.83.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sesv2",
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sesv2",
        sum = "h1:uNAn3m1yFv+7j+tbsAh36kG8JvZlUgZbzdQPSC6W0m4=",
        version = "v1.46.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sfn",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sfn",
        sum = "h1:yIyFY2kbCOoHvuivf9minqnP2RLYJgmvQRYxakIb2oI=",
        version = "v1.19.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sns",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sns",
        sum = "h1:Asj098jPfIZYzAbk4xVFwVBGij5hgMcli0d+5Pe4aZA=",
        version = "v1.21.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sqs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sqs",
        sum = "h1:bp8KUUx15mnLMe8SSJqO/kYEn0C2kKfWq/M9SRK9i1E=",
        version = "v1.24.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:8JdC7Gr9NROg1Rusk25IcZeTO59zLxsKgE0gkh5O6h0=",
        version = "v1.25.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ssooidc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ssooidc",
        sum = "h1:wK8O+j2dOolmpNVY1EWIbLgxrGCHJKVPm08Hv/u80M8=",
        version = "v1.29.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:PZV5W8yk4OtH1JAuhV2PXwwO9v5G5Aoj+eMCn4T+1Kc=",
        version = "v1.33.17",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:uqXzVZNuNexwc/xrh6Tb56u89WDlJY6HS+KC0S4QSjw=",
        version = "v1.22.4",
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
        name = "com_github_aymanbagabas_go_udiff",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/aymanbagabas/go-udiff",
        sum = "h1:TK0fH4MteXUDspT88n8CKzvK0X9O2xu9yQjWpi6yML8=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_aymerick_douceur",
        build_file_proto_mode = "disable",
        importpath = "github.com/aymerick/douceur",
        sum = "h1:Mv+mAeH1Q+n9Fr+oyamOlAkUNPWPlA8PPGR0QAaYuPk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go",
        sum = "h1:KnPIugL51v3N3WwvaSmZbxukD1WuWXOiE9fRdu32f2I=",
        version = "v16.2.1+incompatible",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azcore",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azcore",
        sum = "h1:nyQWyZvwGTvunIMxi1Y9uXkcyr+I7TeNrr/foo4Kpk8=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_azidentity",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/azidentity",
        sum = "h1:tfLQ34V6F7tVSwoTf/4lH5sE0o6eCJuNDTmH09nDpbc=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_internal",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/internal",
        sum = "h1:ywEEhmNahHBihViHepv3xPBn1663uRv2t2q/ESv9seY=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_resourcemanager_compute_armcompute_v5",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5",
        sum = "h1:LkHbJbgF3YyvC53aqYGR+wWQDn2Rdp9AQdGndf9QvY4=",
        version = "v5.7.0",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go_sdk_resourcemanager_network_armnetwork_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4",
        sum = "h1:bXwSugBiSbgtz7rOtbfGf+woewp4f06orW9OP5BjHLA=",
        version = "v4.3.0",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:L/gRVlceqvL25UVaW/CKtUDjefjrs0SPonmDGUVOYP0=",
        version = "v0.0.0-20230124172434-306776ec8161",
    )
    go_repository(
        name = "com_github_azure_go_autorest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest",
        sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
        version = "v14.2.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest",
        sum = "h1:90Y4srNYrwOtAgVo3ndrQkTYn6kf1Eg/AjTFJ8Is2aM=",
        version = "v0.11.18",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_adal",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/adal",
        sum = "h1:Mp5hbtOePIzM8pJVRa3YLrWWmZtoxRXqUEzCfJt3+/Q=",
        version = "v0.9.13",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_date",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/date",
        sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_mocks",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/autorest/mocks",
        sum = "h1:K0laFcLE6VLTOwNgSxaGbUcLPuGXlNkbVvq4cW4nIHk=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_logger",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/logger",
        sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_tracing",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-autorest/tracing",
        sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_azure_go_ntlmssp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Azure/go-ntlmssp",
        sum = "h1:/IBSNwUN8+eKzUzbJPqhK839ygXJ82sde8x3ogr6R28=",
        version = "v0.0.0-20200615164410-66371956d46c",
    )
    go_repository(
        name = "com_github_azuread_microsoft_authentication_library_for_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/AzureAD/microsoft-authentication-library-for-go",
        sum = "h1:XHOnouVk1mxXfQidrMEnLlPk9UMeRtyBTnEFtxkV0kU=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_bazelbuild_rules_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/bazelbuild/rules_go",
        sum = "h1:5vCbuvy8Q11g41lseGJDc5vxhDjJtfxr6nM/IC4VmqM=",
        version = "v0.49.0",
    )
    go_repository(
        name = "com_github_bboreham_go_loser",
        build_file_proto_mode = "disable",
        importpath = "github.com/bboreham/go-loser",
        sum = "h1:6df1vn4bBlDDo4tARvBm7l6KA9iVMnE3NWizDeWSrps=",
        version = "v0.0.0-20230920113527-fcc2c21820a3",
    )
    go_repository(
        name = "com_github_benbjohnson_clock",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:Q92kusRqC1XV2MjkWETPvjJVqKetz1OzxZB7mHJLju8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bgentry/speakeasy",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_bitly_go_hostpool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bitly/go-hostpool",
        sum = "h1:mXoPYz/Ul5HYEDvkta6I8/rnYM5gSdSV2tJ6XbZuEtY=",
        version = "v0.0.0-20171023180738-a3a6125de932",
    )
    go_repository(
        name = "com_github_bitly_go_simplejson",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bitly/go-simplejson",
        sum = "h1:6IH+V8/tVMab511d5bn4M7EwGXZf9Hj6i2xSwkNEM+Y=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_bits_and_blooms_bitset",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bits-and-blooms/bitset",
        sum = "h1:Kn4yilvwNtMACtf1eYDlG8H77R07mZSPbMjLyS07ChA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_bketelsen_crypt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bketelsen/crypt",
        sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
        version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
    )
    go_repository(
        name = "com_github_blang_semver",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/blang/semver",
        sum = "h1:cQNTCjp13qL8KC3Nbxr/y2Bqb63oX6wdnnjpJbkM4JQ=",
        version = "v3.5.1+incompatible",
    )
    go_repository(
        name = "com_github_bluele_gcache",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bluele/gcache",
        sum = "h1:WcbfdXICg7G/DGBh1PFfcirkWOQV+v077yF1pSy3DGw=",
        version = "v0.0.2",
    )
    go_repository(
        name = "com_github_bmatcuk_doublestar_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/bmatcuk/doublestar/v2",
        sum = "h1:6I6oUiT/sU27eE2OFcWqBhL1SwjyvQuOssxT4a1yidI=",
        version = "v2.0.4",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_boombuler_barcode",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/boombuler/barcode",
        sum = "h1:NDBbPmhS+EqABEs5Kg3n/5ZNjy73Pz7SIV+KCeqyXcs=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bradfitz_gomemcache",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bradfitz/gomemcache",
        sum = "h1:Dr+ezPI5ivhMn/3WOoB86XzMhie146DNaBbhaQWZHMY=",
        version = "v0.0.0-20230611145640-acc696258285",
    )
    go_repository(
        name = "com_github_bradleyjkemp_cupaloy_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/bradleyjkemp/cupaloy/v2",
        sum = "h1:any4BmKE+jGIaMpnU8YgH/I2LPiLBufr6oMMlVBbn9M=",
        version = "v2.8.0",
    )
    go_repository(
        name = "com_github_bshuster_repo_logrus_logstash_hook",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bshuster-repo/logrus-logstash-hook",
        sum = "h1:pgAtgj+A31JBVtEHu2uHuEx0n+2ukqUJnS2vVe5pQNA=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_bsm_ginkgo_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bsm/ginkgo/v2",
        sum = "h1:Ny8MWAHyOepLGlLKYmXG4IEkioBysk6GpaRTLC8zwWs=",
        version = "v2.12.0",
    )
    go_repository(
        name = "com_github_bsm_gomega",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bsm/gomega",
        sum = "h1:yeMWxP2pV2fG3FgAODIY8EiRE3dy0aeFYt4l7wh6yKA=",
        version = "v1.27.10",
    )
    go_repository(
        name = "com_github_buger_jsonparser",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/buger/jsonparser",
        sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_bugsnag_bugsnag_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bugsnag/bugsnag-go",
        sum = "h1:rFt+Y/IK1aEZkEHchZRSq9OQbsSzIT/OrI8YFFmRIng=",
        version = "v0.0.0-20141110184014-b1d153021fcd",
    )
    go_repository(
        name = "com_github_bugsnag_osext",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bugsnag/osext",
        sum = "h1:otBG+dV+YK+Soembjv71DPz3uX/V/6MMlSyD9JBQ6kQ=",
        version = "v0.0.0-20130617224835-0dd3f918b21b",
    )
    go_repository(
        name = "com_github_bugsnag_panicwrap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bugsnag/panicwrap",
        sum = "h1:nvj0OLI3YqYXer/kZD8Ri1aaunCxIEsOst1BVJswV0o=",
        version = "v0.0.0-20151223152923-e2c28503fcd0",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:W5quZX/G/csjUnuI8SUYlsHs9M38FC7znL0lIO+DvMg=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_burntsushi_xgb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/BurntSushi/xgb",
        sum = "h1:1BDTz0u9nC3//pOCMdNH+CiXJVYJh5UQNCOBG7jbELc=",
        version = "v0.0.0-20160522181843-27f122750802",
    )
    go_repository(
        name = "com_github_bxcodec_faker_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bxcodec/faker/v4",
        sum = "h1:gqYNBvN72QtzKkYohNDKQlm+pg+uwBDVMN28nWHS18k=",
        version = "v4.0.0-beta.3",
    )
    go_repository(
        name = "com_github_bytedance_sonic",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/bytedance/sonic",
        sum = "h1:MS8gmaH16Gtirygw7jV91pDCN33NyMrPbN7qiYhEsF0=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_github_bytedance_sonic_loader",
        build_file_proto_mode = "disable",
        importpath = "github.com/bytedance/sonic/loader",
        sum = "h1:ZWCw4stuXUsn1/+zQDqeE7JKP+QO47tz7QCNan80NzY=",
        version = "v0.2.4",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff/v3",
        sum = "h1:cfUAAO3yvKMYKPrvhDuHSwQnhZNk/RMHKdZqKTxfm6M=",
        version = "v3.2.2",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:MyRJ/UdXutAwSAT+s3wNd7MfTIcy71VQueUuFK343L8=",
        version = "v4.3.0",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v5",
        build_file_proto_mode = "disable",
        importpath = "github.com/cenkalti/backoff/v5",
        sum = "h1:rIfFVxEf1QsI7E1ZHfp/B4DF/6QBAUhmgkxc0H7Zss8=",
        version = "v5.0.2",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_certifi_gocertifi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/certifi/gocertifi",
        sum = "h1:uH66TXeswKn5PW5zdZ39xEwfS9an067BirqA+P4QaLI=",
        version = "v0.0.0-20200922220541-2c3bb06c6054",
    )
    go_repository(
        name = "com_github_cespare_xxhash",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:UL815xU9SqsFlibzuggzjXhog7bL6oX9BbNZnL2UFvs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_charmbracelet_bubbles",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/bubbles",
        sum = "h1:jSZu6qD8cRQ6k9OMfR1WlM+ruM8fkPWkHvQWD9LIutE=",
        version = "v0.20.0",
    )
    go_repository(
        name = "com_github_charmbracelet_bubbletea",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/bubbletea",
        sum = "h1:KJ2/DnmpfqFtDNVTvYZ6zpPFL9iRCRr0qqKOCvppbPY=",
        version = "v1.1.1",
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
        sum = "h1:4X3PPeoWEDCMvzDvGmTajSyYPcZM4+y8sCA/SsA3cjw=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_charmbracelet_x_ansi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/x/ansi",
        sum = "h1:VfFN0NUpcjBRd4DnKfRaIRo53KRgey/nhOoEqosGDEY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_charmbracelet_x_exp_golden",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/x/exp/golden",
        sum = "h1:MnAMdlwSltxJyULnrYbkZpp4k58Co7Tah3ciKhSNo0Q=",
        version = "v0.0.0-20240815200342-61de596daa2b",
    )
    go_repository(
        name = "com_github_charmbracelet_x_term",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/charmbracelet/x/term",
        sum = "h1:cNB9Ot9q8I711MyZ7myUR5HFWL/lc3OpU8jZ4hwm0x0=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/checkpoint-restore/go-criu/v4",
        sum = "h1:WW2B2uxx9KWF6bGlHqhm8Okiafwwx7Y2kcpn8lCpjgo=",
        version = "v4.1.0",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/checkpoint-restore/go-criu/v5",
        sum = "h1:wpFFOoomK3389ue2lAb0Boag6XPht5QYpipxmSNL4d8=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_chenzhuoyu_base64x",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chenzhuoyu/base64x",
        sum = "h1:77cEq6EriyTZ0g/qfRdp61a3Uu/AWrgIq2s0ClJV1g0=",
        version = "v0.0.0-20230717121745-296ad89f973d",
    )
    go_repository(
        name = "com_github_chenzhuoyu_iasm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chenzhuoyu/iasm",
        sum = "h1:tUHQJXo3NhBqw6s33wkGn9SP3bvrWLdlVIJ3hQBL7P0=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_chromedp_cdproto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chromedp/cdproto",
        sum = "h1:aPflPkRFkVwbW6dmcVqfgwp1i+UWGFH6VgR1Jim5Ygc=",
        version = "v0.0.0-20230802225258-3cf4e6d46a89",
    )
    go_repository(
        name = "com_github_chromedp_chromedp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chromedp/chromedp",
        sum = "h1:dKtNz4kApb06KuSXoTQIyUC2TrA0fhGDwNZf3bcgfKw=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_chromedp_sysutil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chromedp/sysutil",
        sum = "h1:+ZxhTpfpZlmchB58ih/LBHX52ky7w2VhQVKQMucy3Ic=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/logex",
        sum = "h1:XHDu3E6q+gdHgsdTPH6ImJMIp436vR6MPtH8gP05QzM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/readline",
        sum = "h1:upd/6fQk4src78LMRzh5vItIt361/o4uq553V8B5sGI=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_chzyer_test",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/chzyer/test",
        sum = "h1:p3BQDXSxOhOG0P9z6/hGnII4LGiEPOYBhs8asl/fC04=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_cilium_ebpf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cilium/ebpf",
        sum = "h1:V8gS/bTCCjX9uUnkUFUpPsksM8n1lXBAvHcpiFk1X2Y=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_github_circonus_labs_circonus_gometrics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/circonus-labs/circonus-gometrics",
        sum = "h1:C29Ae4G5GtYyYMm1aztcyj/J5ckgJm2zwdDajFbx1NY=",
        version = "v2.3.1+incompatible",
    )
    go_repository(
        name = "com_github_circonus_labs_circonusllhist",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/circonus-labs/circonusllhist",
        sum = "h1:TJH+oke8D16535+jHExHj4nQvzlZrj7ug5D7I/orNUA=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_client9_misspell",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudwego_base64x",
        build_file_proto_mode = "disable",
        importpath = "github.com/cloudwego/base64x",
        sum = "h1:XPciSp1xaq2VCSt6lF0phncD4koWyULpl5bUxbfCyP4=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_cloudwego_iasm",
        build_file_proto_mode = "disable",
        importpath = "github.com/cloudwego/iasm",
        sum = "h1:1KNIy1I1H9hNNFEEH3DVnI4UujN+1zjpuk6gwHLTssg=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:aQ3y1lwWyqYPiWZThqv1aFbZMiM9vblcSArJRf2Irls=",
        version = "v0.0.0-20250501225837-2ac532fd4443",
    )
    go_repository(
        name = "com_github_cockroachdb_apd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/apd",
        sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cockroachdb_cockroach_go_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/cockroach-go/v2",
        sum = "h1:Khtm8K6fTTz/ZCWPzU9Ne3aOW9VyAnj4qIPCJgKtwK0=",
        version = "v2.3.5",
    )
    go_repository(
        name = "com_github_cockroachdb_datadriven",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/datadriven",
        sum = "h1:xD/lrqdvwsc+O2bjSSi3YqY73Ke3LAiSCx49aCesA0E=",
        version = "v0.0.0-20200714090401-bf6692d28da5",
    )
    go_repository(
        name = "com_github_cockroachdb_errors",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/errors",
        sum = "h1:Lap807SXTH5tri2TivECb/4abUkMZC9zRoLarvcKDqs=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_cockroachdb_logtags",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cockroachdb/logtags",
        sum = "h1:o/kfcElHqOiXqcou5a3rIlMc7oJbMQkeLk0VQJ7zgqY=",
        version = "v0.0.0-20190617123548-eb05cc24525f",
    )
    go_repository(
        name = "com_github_codahale_rfc6979",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/codahale/rfc6979",
        sum = "h1:EDmT6Q9Zs+SbUoc7Ik9EfrFqcylYqgPZ9ANSbTAntnE=",
        version = "v0.0.0-20141003034818-6a90f24967eb",
    )
    go_repository(
        name = "com_github_code_hex_go_generics_cache",
        build_file_proto_mode = "disable",
        importpath = "github.com/Code-Hex/go-generics-cache",
        sum = "h1:6vhZGc5M7Y/YD8cIUcY8kcuQLB4cHR7U+0KMqAA0KcU=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_confluentinc_confluent_kafka_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/confluentinc/confluent-kafka-go",
        sum = "h1:gV/GxhMBUb03tFWkN+7kdhg+zf+QUM+wVkI9zwh770Q=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_github_confluentinc_confluent_kafka_go_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/confluentinc/confluent-kafka-go/v2",
        sum = "h1:qy+SfqDauR/TX2qH2VuZqA1rcEAqApBYtHpI6rcqM0U=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_containerd_aufs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/aufs",
        sum = "h1:2oeJiwX5HstO7shSrPZjrohJZLzK36wvpdmzDRkL/LY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_btrfs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/btrfs",
        sum = "h1:osn1exbzdub9L5SouXO5swW4ea/xVdJZ3wokxN5GrnA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cgroups",
        sum = "h1:jN/mbWBEaz+T1pi5OFtnkQ+8qnmEbAr1Oo1FRm5B0dA=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_containerd_cgroups_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cgroups/v3",
        sum = "h1:S5ByHZ/h9PMe5IOQoN7E+nMc2UcLEM/V48DGDJ9kip0=",
        version = "v3.0.3",
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
        name = "com_github_containerd_containerd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/containerd",
        sum = "h1:G/ZQr3gMZs6ZT0qPUZ15znx5QSdQdASW11nXTLTM2Pg=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_containerd_continuity",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/continuity",
        sum = "h1:6HVkalIp+2u1ZLH1J/pYX2oBVXlJZvh1X1A7bEZ9Su8=",
        version = "v0.4.3",
    )
    go_repository(
        name = "com_github_containerd_fifo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/fifo",
        sum = "h1:6PirWBr9/L7GDamKr+XM0IeUFXu5mf3M/BPpH9gaLBU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_go_cni",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/go-cni",
        sum = "h1:el5WPymG5nRRLQF1EfB97FWob4Tdc8INg8RZMaXWZlo=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_github_containerd_go_runc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/go-runc",
        sum = "h1:oU+lLv1ULm5taqgV/CJivypVODI4SUz1znWjv3nNYS0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_imgcrypt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/imgcrypt",
        sum = "h1:iKTstFebwy3Ak5UF0RHSeuCTahC5OIrPJa6vjMAM81s=",
        version = "v1.1.4",
    )
    go_repository(
        name = "com_github_containerd_log",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/log",
        sum = "h1:TCJt7ioM2cr/tfR8GPbGf9/VRAX8D2B4PjzCpfX540I=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_containerd_nri",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/nri",
        sum = "h1:6QioHRlThlKh2RkRTR4kIT3PKAcrLo3gIWnjkM4dQmQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_containerd_stargz_snapshotter_estargz",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/stargz-snapshotter/estargz",
        sum = "h1:5e7heayhB7CcgdTkqfZqrNaNv15gABwr3Q2jBTbLlt4=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_containerd_ttrpc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/ttrpc",
        sum = "h1:GbtyLRxb0gOLR0TYQWt3O6B0NvT8tMdorEHqIQo/lWI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_containerd_typeurl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/typeurl",
        sum = "h1:Chlt8zIieDbzQFzXzAeBEF92KhExuE4p9p92/QmY7aY=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_containerd_zfs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/zfs",
        sum = "h1:cXLJbx+4Jj7rNsTiqVfm6i+RNLx6FFA2fMmDlEf+Wm8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containernetworking_cni",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containernetworking/cni",
        sum = "h1:ky20T7c0MvKvbMOwS/FrlbNwjEoqJEUUYfsL4b0mc4k=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_containernetworking_plugins",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containernetworking/plugins",
        sum = "h1:+AGfFigZ5TiQH00vhR8qPeSatj53eNGz0C1d3wVYlHE=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_containers_ocicrypt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/containers/ocicrypt",
        sum = "h1:uMxn2wTb4nDR7GqG3rnZSfpJXqWURfzZ7nKydzIeKpA=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_coreos_bbolt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/bbolt",
        sum = "h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_coreos_etcd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/etcd",
        sum = "h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
        version = "v3.3.13+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_iptables",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-iptables",
        sum = "h1:is9qnZMPYjLd8LYqmm/qlE+wwEgJIkTYdhV3rfZo4jk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-oidc",
        sum = "h1:mh48q/BqXqgjVHpy2ZY7WnWAbenxRjsz9N1i1YxjHAk=",
        version = "v2.2.1+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:wkHLiw0WNATZnSG7epLsujiMCgPAc9xhjJ4tgnAxmfM=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-systemd",
        sum = "h1:JOrtw2xFKzlg+cbHpyrpLDmnN1HqhBfnX7WDiW7eG2c=",
        version = "v0.0.0-20190719114852-fd7a80b32e1f",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
        version = "v22.5.0",
    )
    go_repository(
        name = "com_github_coreos_pkg",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/coreos/pkg",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:XJtiaUW6dEEqVuZiMTn1ldk455QWwEIsMIJlo5vtkx0=",
        version = "v2.0.6",
    )
    go_repository(
        name = "com_github_creack_pty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/creack/pty",
        sum = "h1:07n33Z8lZxZ2qwegKbObQohDhXDQxiMMz1NOUGYlesw=",
        version = "v1.1.11",
    )
    go_repository(
        name = "com_github_cristalhq_jwt_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/cristalhq/jwt/v4",
        sum = "h1:g/AD3h0VicDamtlM70GWGElp8kssQEv+5wYd7L9WOhU=",
        version = "v4.0.2",
    )
    go_repository(
        name = "com_github_cyphar_filepath_securejoin",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/cyphar/filepath-securejoin",
        sum = "h1:YX6ebbZCZP7VkM3scTTokDgBL2TY741X51MTk3ycuNI=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_d2g_dhcp4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/dhcp4",
        sum = "h1:Xo2rK1pzOm0jO6abTPIQwbAmqBIOj132otexc1mmzFc=",
        version = "v0.0.0-20170904100407-a1d1b6c41b1c",
    )
    go_repository(
        name = "com_github_d2g_dhcp4client",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/dhcp4client",
        sum = "h1:suYBsYZIkSlUMEz4TAYCczKf62IA2UWC+O8+KtdOhCo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_d2g_dhcp4server",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/dhcp4server",
        sum = "h1:+CpLbZIeUn94m02LdEKPcgErLJ347NUwxPKs5u8ieiY=",
        version = "v0.0.0-20181031114812-7d4a0a7f59a5",
    )
    go_repository(
        name = "com_github_d2g_hardwareaddr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/d2g/hardwareaddr",
        sum = "h1:itqmmf1PFpC4n5JW+j4BU7X4MTfVurhYRTjODoPb2Y8=",
        version = "v0.0.0-20190221164911-e7d9fbe030e4",
    )
    go_repository(
        name = "com_github_datadog_appsec_internal_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/appsec-internal-go",
        sum = "h1:xpAS/hBo429pVh7rngquAK2DezUaJjfsX7Wd8cw0aIk=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_datadog_datadog_agent_pkg_obfuscate",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-agent/pkg/obfuscate",
        sum = "h1:bUMSNsw1iofWiju9yc1f+kBd33E3hMJtq9GuU602Iy8=",
        version = "v0.48.0",
    )
    go_repository(
        name = "com_github_datadog_datadog_agent_pkg_remoteconfig_state",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-agent/pkg/remoteconfig/state",
        sum = "h1:5nE6N3JSs2IG3xzMthNFhXfOaXlrsdgqmJ73lndFf8c=",
        version = "v0.48.1",
    )
    go_repository(
        name = "com_github_datadog_datadog_api_client_go_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-api-client-go/v2",
        sum = "h1:9Zq42D6M3U///VDxjx2SS1g+EW55WhZYZFHtzM+cO4k=",
        version = "v2.25.0",
    )
    go_repository(
        name = "com_github_datadog_datadog_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-go",
        sum = "h1:qSG2N4FghB1He/r2mFrWKCaL7dXCilEuNEeAn20fdD4=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_datadog_datadog_go_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/datadog-go/v5",
        sum = "h1:2q2qjFOb3RwAZNU+ez27ZVDwErJv5/VpbBPprz7Z+s8=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_datadog_go_libddwaf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/go-libddwaf",
        sum = "h1:C0cHE++wMFWf5/BDO8r/3dTDCj21U/UmPIT0PiFMvsA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_datadog_go_libddwaf_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/go-libddwaf/v2",
        sum = "h1:bujaT5+KnLDFQqVA5ilvVvW+evUSHow9FrTHRgUwN4A=",
        version = "v2.3.1",
    )
    go_repository(
        name = "com_github_datadog_go_tuf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/go-tuf",
        sum = "h1:EeZr937eKAWPxJ26IykAdWA4A0jQXJgkhUjqEI/w7+I=",
        version = "v1.0.2-0.5.2",
    )
    go_repository(
        name = "com_github_datadog_gostackparse",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/gostackparse",
        sum = "h1:i7dLkXHvYzHV308hnkvVGDL3BR4FWl7IsXNPz/IGQh4=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_datadog_sketches_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/sketches-go",
        sum = "h1:gppNudE9d19cQ98RYABOetxIhpTCl4m7CnbRZjvVA/o=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_datadog_zstd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/DataDog/zstd",
        sum = "h1:vUG4lAyuPCXO0TLbXvPv7EB7cNK1QV/luu55UHLrrn8=",
        version = "v1.5.2",
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
        name = "com_github_decred_dcrd_dcrec_secp256k1_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/decred/dcrd/dcrec/secp256k1/v4",
        sum = "h1:8UrgZ3GkP4i/CLijOJx79Yu+etlyjdBU4sfcs2WYQMs=",
        version = "v4.2.0",
    )
    go_repository(
        name = "com_github_denisenkom_go_mssqldb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/denisenkom/go-mssqldb",
        sum = "h1:9rHa233rhdOyrz2GcP9NM+gi2psgJZ4GWDpL/7ND8HI=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_github_dennwc_varint",
        build_file_proto_mode = "disable",
        importpath = "github.com/dennwc/varint",
        sum = "h1:kGNFFSSw8ToIy3obO/kKr8U9GZYUAxQEVuix4zfDWzE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_denverdino_aliyungo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/denverdino/aliyungo",
        sum = "h1:p6poVbjHDkKa+wtC8frBMwQtT3BmqGYBjzMwJ63tuR4=",
        version = "v0.0.0-20190125010748-a747050bb1ba",
    )
    go_repository(
        name = "com_github_dgraph_io_ristretto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dgraph-io/ristretto",
        sum = "h1:SYG07bONKMlFDUYu5pEu3DGAh8c2OFNzKm6G9J4Si84=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_farm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-farm",
        sum = "h1:fAjc9m62+UWV/WAFKLNi6ZS0675eEUC9y3AlwSbQu1Y=",
        version = "v0.0.0-20200201041132-a6ae2369ad13",
    )
    go_repository(
        name = "com_github_dgryski_go_rendezvous",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-rendezvous",
        sum = "h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=",
        version = "v0.0.0-20200823014737-9f7001d12a5f",
    )
    go_repository(
        name = "com_github_dgryski_go_sip13",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/go-sip13",
        sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
        version = "v0.0.0-20181026042036-e10d5fee7954",
    )
    go_repository(
        name = "com_github_dgryski_trifles",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dgryski/trifles",
        sum = "h1:fRzb/w+pyskVMQ+UbP35JkH8yB7MYb4q/qhBarqZE6g=",
        version = "v0.0.0-20200323201526-dd97f9abfb48",
    )
    go_repository(
        name = "com_github_digitalocean_godo",
        build_file_proto_mode = "disable",
        importpath = "github.com/digitalocean/godo",
        sum = "h1:+Znh7VMQj/E8ArbjWnc7OKGjWfzC+I8OCSRp7r1MdD8=",
        version = "v1.126.0",
    )
    go_repository(
        name = "com_github_dimfeld_httptreemux_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dimfeld/httptreemux/v5",
        sum = "h1:p8jkiMrCuZ0CmhwYLcbNbl7DDo21fozhKHQ2PccwOFQ=",
        version = "v5.5.0",
    )
    go_repository(
        name = "com_github_disintegration_imaging",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/disintegration/imaging",
        sum = "h1:w1LecBlG2Lnp8B3jk5zSuNqd7b4DXhcjwek1ei82L+c=",
        version = "v1.6.2",
    )
    go_repository(
        name = "com_github_distribution_reference",
        build_file_proto_mode = "disable",
        importpath = "github.com/distribution/reference",
        sum = "h1:0IXCQ5g4/QMHHkarYzh5l+u8T3t73zM5QvfrDyIgxBk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_dnaeon_go_vcr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dnaeon/go-vcr",
        sum = "h1:zHCHvJYTMh1N7xnV7zf1m1GPBF9Ad0Jk/whtQ1663qI=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_dnephin_pflag",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dnephin/pflag",
        sum = "h1:oxONGlWxhmUct0YzKTgrpQv9AUA1wtPBn7zuSjJqptk=",
        version = "v1.0.7",
    )
    go_repository(
        name = "com_github_docker_cli",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/cli",
        sum = "h1:I8PHdc0MtxEADqYJZvhBrW9bo8gawKwwenxRM7/rLu8=",
        version = "v26.1.4+incompatible",
    )
    go_repository(
        name = "com_github_docker_distribution",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/distribution",
        sum = "h1:Q50tZOPR6T/hjNsyc9g8/syEs6bk8XXApsHjKukMl68=",
        version = "v2.8.1+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/docker",
        sum = "h1:KttF0XoteNTicmUtBO0L2tP+J7FGRFTjaEF4k6WdhfI=",
        version = "v27.3.1+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker_credential_helpers",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/docker-credential-helpers",
        sum = "h1:zI2p9+1NQYdnG6sMU26EX4aVGlqbInSQxQXLvzJ4RPQ=",
        version = "v0.6.3",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-connections",
        sum = "h1:USnMq7hx7gwdVZq1L49hLXaFtUdTADjXGp+uj1Br63c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_docker_go_events",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-events",
        sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
        version = "v0.0.0-20190806004212-e31b211e4f1c",
    )
    go_repository(
        name = "com_github_docker_go_metrics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-metrics",
        sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_docker_go_units",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/go-units",
        sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_docker_libtrust",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/libtrust",
        sum = "h1:ZClxb8laGDf5arXfYcAtECDFgAgHklGI8CxgjHnXKJ4=",
        version = "v0.0.0-20150114040149-fa567046d9b1",
    )
    go_repository(
        name = "com_github_docker_spdystream",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docker/spdystream",
        sum = "h1:cenwrSVm+Z7QLSV/BsnenAOcDXdX4cMv4wP0B/5QbPg=",
        version = "v0.0.0-20160310174837-449fdfce4d96",
    )
    go_repository(
        name = "com_github_docopt_docopt_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/docopt/docopt-go",
        sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
        version = "v0.0.0-20180111231733-ee0de3bc6815",
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
        name = "com_github_dvyukov_go_fuzz",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/dvyukov/go-fuzz",
        sum = "h1:q1oJaUPdmpDm/VyXosjgPgr6wS7c5iV2p0PwJD73bUI=",
        version = "v0.0.0-20210103155950-6a8e9d1f2415",
    )
    go_repository(
        name = "com_github_eapache_go_resiliency",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/go-resiliency",
        sum = "h1:3OK9bWpPk5q6pbFAaYSEwD9CLUSHG8bnZuqX2yMt3B0=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_eapache_go_xerial_snappy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/go-xerial-snappy",
        sum = "h1:Oy0F4ALJ04o5Qqpdz8XLIpNA3WM/iSIXqxtqo7UGVws=",
        version = "v0.0.0-20230731223053-c322873962e3",
    )
    go_repository(
        name = "com_github_eapache_queue",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/eapache/queue",
        sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_ebitengine_purego",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ebitengine/purego",
        sum = "h1:EYID3JOAdmQ4SNZYJHu9V6IqOeRQDBYxqKAg9PyoHFY=",
        version = "v0.6.0-alpha.5",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:6EUwBLQ/Mcr1EYLE4Tn1VdW1A4ckqCQWZBw8Hr0kjpQ=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_elastic_elastic_transport_go_v8",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/elastic/elastic-transport-go/v8",
        sum = "h1:NeqEz1ty4RQz+TVbUrpSU7pZ48XkzGWQj02k5koahIE=",
        version = "v8.1.0",
    )
    go_repository(
        name = "com_github_elastic_go_elasticsearch_v6",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/elastic/go-elasticsearch/v6",
        sum = "h1:U2HtkBseC1FNBmDr0TR2tKltL6FxoY+niDAlj5M8TK8=",
        version = "v6.8.5",
    )
    go_repository(
        name = "com_github_elastic_go_elasticsearch_v7",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/elastic/go-elasticsearch/v7",
        sum = "h1:49mHcHx7lpCL8cW1aioEwSEVKQF3s+Igi4Ye/QTWwmk=",
        version = "v7.17.1",
    )
    go_repository(
        name = "com_github_elastic_go_elasticsearch_v8",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/elastic/go-elasticsearch/v8",
        sum = "h1:Rn1mcqaIMcNT43hnx2H62cIFZ+B6mjWtzj85BDKrvCE=",
        version = "v8.4.0",
    )
    go_repository(
        name = "com_github_elazarl_goproxy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/elazarl/goproxy",
        sum = "h1:RIB4cRk+lBqKK3Oy0r2gRX4ui7tuhiZq2SuTtTCi0/0=",
        version = "v0.0.0-20221015165544-a0805db90819",
    )
    go_repository(
        name = "com_github_elliotchance_sshtunnel",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/elliotchance/sshtunnel",
        sum = "h1:uunvEbhtzDqEyl58E1qC7j2sDFXhtEcj0sEsc33e/Gw=",
        version = "v1.6.1",
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/emicklei/go-restful",
        sum = "h1:rgqiKNjTnFQA6kkhFe16D8epTksy9HQ1MyrbDXSdYhM=",
        version = "v2.16.0+incompatible",
    )
    go_repository(
        name = "com_github_emicklei_go_restful_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/emicklei/go-restful/v3",
        sum = "h1:rAQeMHw1c7zTmncogyy8VvRZwtkmkZ4FxERmMY4rD+g=",
        version = "v3.11.0",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:zEqyPVyku6IvWCFwux4x9RxkLOMUL+1vC9xUFv5l2/M=",
        version = "v0.13.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane_envoy",
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/go-control-plane/envoy",
        sum = "h1:jb83lalDRZSpPWW2Z7Mck/8kXZ5CQAFYVjQcdVIr83A=",
        version = "v1.32.4",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane_ratelimit",
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/go-control-plane/ratelimit",
        sum = "h1:/G9QYbddjL25KvtKTv3an9lx6VBE2cnb8wp1vEGNYGI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:DEo3O99U8j4hBFwbJfrz9VtgcDfUKS7KJ7spH3d86P8=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_erikgeiser_coninput",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/erikgeiser/coninput",
        sum = "h1:Y/CXytFA4m6baUTXGLOoWe4PQhGxaX0KpnayAqC48p4=",
        version = "v0.0.0-20211004153227-1c3628e74d0f",
    )
    go_repository(
        name = "com_github_erikstmartin_go_testdb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/erikstmartin/go-testdb",
        sum = "h1:Yzb9+7DPaBjB8zlTR87/ElzFsnQfuHnVUVqpZZIcV5Y=",
        version = "v0.0.0-20160219214506-8d10e4a1bae5",
    )
    go_repository(
        name = "com_github_ettle_strcase",
        build_file_proto_mode = "disable",
        importpath = "github.com/ettle/strcase",
        sum = "h1:htFueZyVeE1XNnMEfbqp5r67qAN/4r6ya1ysq8Q+Zcw=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/evanphx/json-patch",
        sum = "h1:4onqiflcdA9EOZ4RxV643DvftH5pOlLGNtQ5lPWQu84=",
        version = "v4.12.0+incompatible",
    )
    go_repository(
        name = "com_github_evanphx_json_patch_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/evanphx/json-patch/v5",
        sum = "h1:b91NhWfaz02IuVxO9faSllyAtNXHMPkC5J8sJCLunww=",
        version = "v5.6.0",
    )
    go_repository(
        name = "com_github_facette_natsort",
        build_file_proto_mode = "disable",
        importpath = "github.com/facette/natsort",
        sum = "h1:IT4JYU7k4ikYg1SCxNI1/Tieq/NFvh6dzLdgi7eu0tM=",
        version = "v0.0.0-20181210072756-2cd4dd1e2dcb",
    )
    go_repository(
        name = "com_github_fasthttp_websocket",
        build_file_proto_mode = "disable",
        importpath = "github.com/fasthttp/websocket",
        sum = "h1:omHqsl8j+KXpmzRjF8bmzOSYJ8GnS0E3efi1wYT+niY=",
        version = "v1.4.3-rc.6",
    )
    go_repository(
        name = "com_github_fatih_color",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/color",
        sum = "h1:zmkK9Ngbjj+K0yRhTVONQh1p/HknKYSlNT+vZCzyokM=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_fatih_structs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/fatih/structs",
        sum = "h1:Q7juDM0QtcnhCpeyLGQKyg4TOIghuNXrkL32pHAUMxo=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_felixge_fgprof",
        build_file_proto_mode = "disable",
        importpath = "github.com/felixge/fgprof",
        sum = "h1:VvyZxILNuCiUCSXtPtYmmtGvb65nqXh2QFWc0Wpf2/g=",
        version = "v0.9.3",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_flynn_go_docopt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/flynn/go-docopt",
        sum = "h1:Ss/B3/5wWRh8+emnK0++g5zQzwDTi30W10pKxKc4JXI=",
        version = "v0.0.0-20140912013429-f6dd2ebbb31e",
    )
    go_repository(
        name = "com_github_fogleman_gg",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/fogleman/gg",
        sum = "h1:/7zJX8F6AaYQc57WQCyN9cAIz+4bCJGO9B+dyW29am8=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_form3tech_oss_jwt_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/form3tech-oss/jwt-go",
        sum = "h1:7ZaBxOI7TMoYBfyA3cQHErNNyAWIKUMIwqxEtgHOs5c=",
        version = "v3.2.3+incompatible",
    )
    go_repository(
        name = "com_github_fortytw2_leaktest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/fortytw2/leaktest",
        sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:dfYrrRyLtiqT9GyKXgdh+k4inNeTvmGbuSgZ3lx3GhA=",
        version = "v1.14.5",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_fullsailor_pkcs7",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/fullsailor/pkcs7",
        sum = "h1:RDBNVkRviHZtvDvId8XSGPu3rmpmSe+wKRcEWNgsfWU=",
        version = "v0.0.0-20190404230743-d7302db945fa",
    )
    go_repository(
        name = "com_github_fxamacker_cbor_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/fxamacker/cbor/v2",
        sum = "h1:iM5WgngdRBanHcxugY4JySA0nk1wZorNOpTgCMedv5E=",
        version = "v2.7.0",
    )
    go_repository(
        name = "com_github_gabriel_vasile_mimetype",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gabriel-vasile/mimetype",
        sum = "h1:5k+WDwEsD9eTLL8Tz3L0VnmVh9QxGjRmjBvAG7U/oYY=",
        version = "v1.4.9",
    )
    go_repository(
        name = "com_github_garyburd_redigo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/garyburd/redigo",
        sum = "h1:LFu2R3+ZOPgSMWMOL+saa/zXRjw0ID2G8FepO53BGlg=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_github_gavv_httpexpect_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/gavv/httpexpect/v2",
        sum = "h1:nIJqt5v5e4P7/0jODpX2gtSw+pHXUqdP28YcjqwDZmE=",
        version = "v2.17.0",
    )
    go_repository(
        name = "com_github_gemalto_flume",
        build_file_proto_mode = "disable",
        importpath = "github.com/gemalto/flume",
        sum = "h1:EEeQvAxyFys3BH8IxEU7ZpM6Kr1sYn20HuZq6dgyMR8=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_gemalto_kmip_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/gemalto/kmip-go",
        sum = "h1:jAAZejUdRrspKigLoA62MTmIj0T7DDDOzdxHi1cDjoU=",
        version = "v0.0.10",
    )
    go_repository(
        name = "com_github_getkin_kin_openapi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/getkin/kin-openapi",
        sum = "h1:j77zg3Ec+k+r+GA3d8hBoXpAc6KX9TbBPrwQGBIy2sY=",
        version = "v0.76.0",
    )
    go_repository(
        name = "com_github_getsentry_raven_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/getsentry/raven-go",
        sum = "h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gin_contrib_cors",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gin-contrib/cors",
        sum = "h1:0Z7D/bVhE6ja07lI8CTjTonp6SB07o8bNuFyRbsBUQg=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_gin_contrib_gzip",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gin-contrib/gzip",
        sum = "h1:NjcunTcGAj5CO1gn4N8jHOSIeRFHIbn51z6K+xaN4d4=",
        version = "v0.0.6",
    )
    go_repository(
        name = "com_github_gin_contrib_sse",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gin-contrib/sse",
        sum = "h1:n0w2GMuUpWDVp7qSpvze6fAu9iRxJY4Hmj6AmBOU05w=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_gin_gonic_gin",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gin-gonic/gin",
        sum = "h1:T0ujvqyCSqRopADpgPgiTT63DUQVSfojyME59Ei63pQ=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_globalsign_mgo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/globalsign/mgo",
        sum = "h1:DujepqpGd1hyOd7aW59XpK7Qymp8iy83xq74fLr21is=",
        version = "v0.0.0-20181015135952-eeefdecb41b8",
    )
    go_repository(
        name = "com_github_gnboorse_centipede",
        build_file_proto_mode = "disable",
        importpath = "github.com/gnboorse/centipede",
        sum = "h1:lpGjsG4b2bPc+GoHozJZis0u4JTvVyUBpg+KADTXE9U=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_go_asn1_ber_asn1_ber",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-asn1-ber/asn1-ber",
        sum = "h1:pDbRAunXzIUXfx4CB2QJFv5IuPiuoW+sWvr/Us009o8=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_go_bindata_go_bindata",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-bindata/go-bindata",
        sum = "h1:5vjJMVhowQdPzjE1LdxyFF7YFTXg5IgGVW4gBr5IbvE=",
        version = "v3.1.2+incompatible",
    )
    go_repository(
        name = "com_github_go_chi_chi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chi/chi",
        sum = "h1:QHdzF2szwjqVV4wmByUnTcsbIg7UGaQ0tPF2t5GcAIs=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_go_chi_chi_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-chi/chi/v5",
        sum = "h1:rLz5avzKpjqxrYwXNfmjkrYYXOyLJd37pz53UFHC6vk=",
        version = "v5.0.10",
    )
    go_repository(
        name = "com_github_go_errors_errors",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-errors/errors",
        sum = "h1:ljK/pL5ltg3qoN+OtN6yCv9HWSfMwxSx90GJCZQxYNg=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_go_fonts_dejavu",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-fonts/dejavu",
        sum = "h1:JSajPXURYqpr+Cu8U9bt8K+XcACIHWqWrvWCKyeFmVQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_go_fonts_latin_modern",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-fonts/latin-modern",
        sum = "h1:5/Tv1Ek/QCr20C6ZOz15vw3g7GELYL98KWr8Hgo+3vk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_fonts_liberation",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-fonts/liberation",
        sum = "h1:jAkAWJP4S+OsrPLZM4/eC9iW7CtHy+HBXrEwZXWo5VM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_fonts_stix",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-fonts/stix",
        sum = "h1:UlZlgrvvmT/58o573ot7NFw0vZasZ5I6bcIft/oMdgg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_go_gl_glfw",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gl/glfw",
        sum = "h1:QbL/5oDUmRBzO9/Z7Seo6zf912W/a6Sr4Eu0G/3Jho0=",
        version = "v0.0.0-20190409004039-e6da0acd62b1",
    )
    go_repository(
        name = "com_github_go_gl_glfw_v3_3_glfw",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-gl/glfw/v3.3/glfw",
        sum = "h1:WtGNWLvXpe6ZudgnXrq0barxBImvnnJoMEhXAzcbM0I=",
        version = "v0.0.0-20200222043503-6f7a984d4dc4",
    )
    go_repository(
        name = "com_github_go_ini_ini",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-ini/ini",
        sum = "h1:Mujh4R/dH6YL8bxuISne3xX2+qcQ9p0IxKAP6ExWoUo=",
        version = "v1.25.4",
    )
    go_repository(
        name = "com_github_go_jose_go_jose_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-jose/go-jose/v3",
        sum = "h1:fFKWeig/irsp7XD2zBxvnmA/XaRWp5V3CBsZXJF7G7k=",
        version = "v3.0.3",
    )
    go_repository(
        name = "com_github_go_jose_go_jose_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-jose/go-jose/v4",
        sum = "h1:M6T8+mKZl/+fNNuFHvGIzDz7BTLQPIounk/b9dw3AaE=",
        version = "v4.0.5",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/kit",
        sum = "h1:wDJmvq38kDhkVxi50ni9ykkdUr1PKgqKOoi01fa0Mdk=",
        version = "v0.9.0",
    )
    go_repository(
        name = "com_github_go_kit_log",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_latex_latex",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-latex/latex",
        sum = "h1:6zl3BbBhdnMkpSj2YY30qV3gDcVBGtFgVsV3+/i+mKQ=",
        version = "v0.0.0-20210823091927-c0d11ff05a81",
    )
    go_repository(
        name = "com_github_go_ldap_ldap_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-ldap/ldap/v3",
        sum = "h1:fU/0xli6HY02ocbMuozHAYsaHLcnkLjvho2r5a34BUU=",
        version = "v3.4.1",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/logr",
        sum = "h1:CjnDlHq8ikf6E492q6eKboGOC0T8CDaOvkHCIg8idEI=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_openapi_analysis",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/analysis",
        sum = "h1:ZBmNoP2h5omLKr/srIC9bfqrUGzT6g6gNv03HE9Vpj0=",
        version = "v0.22.2",
    )
    go_repository(
        name = "com_github_go_openapi_errors",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/errors",
        sum = "h1:c4xY/OLxUBSTiepAg3j/MHuAv5mJhnf53LLMWFB+u/w=",
        version = "v0.22.0",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:mQc3nmndL8ZBzStEo3JYF8wzmeWffDH4VbXz58sAx6Q=",
        version = "v0.20.2",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:bKlDxQxQJgwpUSgOENiMPzCTBVuc7vTdXSSgNeAhojU=",
        version = "v0.20.4",
    )
    go_repository(
        name = "com_github_go_openapi_loads",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/loads",
        sum = "h1:jDzF4dSoHw6ZFADCGltDb2lE4F6De7aWSpe+IcsRzT0=",
        version = "v0.21.5",
    )
    go_repository(
        name = "com_github_go_openapi_runtime",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/runtime",
        sum = "h1:yX9HMGQbz32M87ECaAhGpJjBmErO3QLcgdZj9BzGx7c=",
        version = "v0.24.2",
    )
    go_repository(
        name = "com_github_go_openapi_spec",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/spec",
        sum = "h1:7CBlRnw+mtjFGlPDRZmAMnq35cRzI91xj03HVyUi/Do=",
        version = "v0.20.14",
    )
    go_repository(
        name = "com_github_go_openapi_strfmt",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/strfmt",
        sum = "h1:nlUS6BCqcnAk0pyhi9Y+kdDVZdZMHfEKQiS4HaMgO/c=",
        version = "v0.23.0",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:XX2DssF+mQKM2DHsbgZK74y/zj4mo9I99+89xUmuZCE=",
        version = "v0.22.9",
    )
    go_repository(
        name = "com_github_go_openapi_validate",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-openapi/validate",
        sum = "h1:2l7PJLzCis4YUGEoW6eoQw3WhyM65WSIcjX6SQnlfDw=",
        version = "v0.23.0",
    )
    go_repository(
        name = "com_github_go_pdf_fpdf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-pdf/fpdf",
        sum = "h1:MlgtGIfsdMEEQJr2le6b/HNr1ZlQwxyWr77r2aj2U/8=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_go_pg_pg_v10",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-pg/pg/v10",
        sum = "h1:vYwbFpqoMpTDphnzIPshPPepdy3VpzD8qo29OFKp4vo=",
        version = "v10.11.1",
    )
    go_repository(
        name = "com_github_go_pg_zerochecker",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-pg/zerochecker",
        sum = "h1:pp7f72c3DobMWOb2ErtZsnrPaSvHd2W4o9//8HtF4mU=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_playground_assert_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/assert/v2",
        sum = "h1:JvknZsQTYeFEAhQwI4qEt9cyV5ONwRHC+lYKSsYSR8s=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/locales",
        sum = "h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_go_playground_validator_v10",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-playground/validator/v10",
        sum = "h1:SP05Nqhjcvz81uJaRfEV0YBSSSGMc/iMaVtFbr3Sw2k=",
        version = "v10.26.0",
    )
    go_repository(
        name = "com_github_go_redis_redis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis",
        sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
        version = "v6.15.9+incompatible",
    )
    go_repository(
        name = "com_github_go_redis_redis_extra_rediscmd_v8",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/extra/rediscmd/v8",
        sum = "h1:ftG8tp8SG81xyuL2woNEx5t2RZ8mOJuC2+tumi+/NR8=",
        version = "v8.11.5",
    )
    go_repository(
        name = "com_github_go_redis_redis_extra_redisotel_v8",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/extra/redisotel/v8",
        sum = "h1:BqyYJgvdSr2S/6O2l7zmCj26ocUTxDLgagsGIRfkS+Q=",
        version = "v8.11.5",
    )
    go_repository(
        name = "com_github_go_redis_redis_v7",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/v7",
        sum = "h1:PASvf36gyUpr2zdOUS/9Zqc80GbM+9BDyiJSJDDOrTI=",
        version = "v7.4.1",
    )
    go_repository(
        name = "com_github_go_redis_redis_v8",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redis/redis/v8",
        sum = "h1:AcZZR7igkdvfVmQTPnu9WE37LRrO/YrBH5zWyjDC0oI=",
        version = "v8.11.5",
    )
    go_repository(
        name = "com_github_go_redsync_redsync_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-redsync/redsync/v4",
        sum = "h1:vRmYusI+qF95XSpApHAdeu+RjyDvxBXbMthbc/x148c=",
        version = "v4.9.4",
    )
    go_repository(
        name = "com_github_go_resty_resty_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-resty/resty/v2",
        sum = "h1:x+LHXBI2nMB1vqndymf26quycC4aggYJ7DECYbiz03g=",
        version = "v2.13.1",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:LedoTUt/eveggdHS9qUFC1EFSa8bU2+1pZjSRpvNJ1Y=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-task/slim-sprig",
        sum = "h1:tfuBGBXKqDEevZMzYi5KSi8KkcZtzBcTgAUUtapy0OI=",
        version = "v0.0.0-20230315185526-52ccab3ef572",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-task/slim-sprig/v3",
        sum = "h1:sUs3vkvUymDpBKi3qH1YSqBQk9+9D/8M2mN1vB6EwHI=",
        version = "v3.0.0",
    )
    go_repository(
        name = "com_github_go_test_deep",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-test/deep",
        sum = "h1:WOcxcdHcvdgThNXjw0t76K42FXTU7HpNQWHpA2HHNlg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_go_text_typesetting",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/go-text/typesetting",
        sum = "h1:yV4rFdcvwZXE0lZZ3EoBWjVysHyVo8DLY8VihDciNN0=",
        version = "v0.0.0-20230329143336-a38d00edd832",
    )
    go_repository(
        name = "com_github_go_zookeeper_zk",
        build_file_proto_mode = "disable",
        importpath = "github.com/go-zookeeper/zk",
        sum = "h1:DPzxraQx7OrPyXq2phlGlNSIyWEsAox0RJmjTseMV6I=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_gobuffalo_attrs",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/attrs",
        sum = "h1:LPXvtWcVpt6KJh31cK4MDyWlgehutmHpvjaWvHDRI4o=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_gobuffalo_envy",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/envy",
        sum = "h1:EIi03p9c3yeuRCFPOKcSfajzkLb3hrRjEpHGI8I2Wo4=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_github_gobuffalo_fizz",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/fizz",
        sum = "h1:8uume7joF6niTNWN582IQ2jhGTUoa9g1fiV/tIoGdBs=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_github_gobuffalo_flect",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/flect",
        sum = "h1:eBFmskjXZgAOagiTXJH25Nt5sdFwNRcb8DKZsIsAUQI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gobuffalo_genny_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/genny/v2",
        sum = "h1:cCRBbqzo3GfNvj3UetD16zRgUvWFEyyl0qTqquuIqOM=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_gobuffalo_github_flavored_markdown",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/github_flavored_markdown",
        sum = "h1:rSMPtx9ePkFB22vJ+dH+m/EUBS8doQ3S8LeEXcdwZHk=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_gobuffalo_helpers",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/helpers",
        sum = "h1:C9CedoRSfgWg2ZoIkVXgjI5kgmSpL34Z3qdnzpfNVd8=",
        version = "v0.6.7",
    )
    go_repository(
        name = "com_github_gobuffalo_here",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/here",
        sum = "h1:hpfhh+kt2y9JLDfhYUxxCRxQol540jsVfKUZzjlbp8o=",
        version = "v0.6.7",
    )
    go_repository(
        name = "com_github_gobuffalo_httptest",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/httptest",
        sum = "h1:GpGy520SfY1QEmyPvaqmznTpG4gEQqQ82HtHqyNEreM=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_gobuffalo_logger",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/logger",
        sum = "h1:LTLwWelETXDYyqF/ASf0nxaIcdEOIJNxRokPcfI/xbU=",
        version = "v1.0.7",
    )
    go_repository(
        name = "com_github_gobuffalo_nulls",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/nulls",
        sum = "h1:GAqBR29R3oPY+WCC7JL9KKk9erchaNuV6unsOSZGQkw=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_gobuffalo_packd",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/packd",
        sum = "h1:Yg523YqnOxGIWCp69W12yYBKsoChwI7mtu6ceM9Bwfw=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_gobuffalo_plush_v4",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/plush/v4",
        sum = "h1:bnPjdMTEUQHqj9TNX2Ck3mxEXYZa+0nrFMNM07kpX9g=",
        version = "v4.1.18",
    )
    go_repository(
        name = "com_github_gobuffalo_pop_v6",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/pop/v6",
        sum = "h1:eUDBaZcb0gYrmFnKwpuTEUA7t5ZHqNfvS4POqJYXDZY=",
        version = "v6.1.1",
    )
    go_repository(
        name = "com_github_gobuffalo_tags_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/tags/v3",
        sum = "h1:X/ydLLPhgXV4h04Hp2xlbI2oc5MDaa7eub6zw8oHjsM=",
        version = "v3.1.4",
    )
    go_repository(
        name = "com_github_gobuffalo_validate_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobuffalo/validate/v3",
        sum = "h1:o7wkIGSvZBYBd6ChQoLxkz2y1pfmhbI4jNJYh6PuNJ4=",
        version = "v3.3.3",
    )
    go_repository(
        name = "com_github_gobwas_glob",
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/glob",
        sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_gobwas_httphead",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:exrUm0f4YX0L7EBwZHuCF4GDp8aJfVeBrlLQrs6NqWU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/pool",
        sum = "h1:xfeeEhW7pwmX8nuLVlqbzVc7udMDrwetjEv+TZIz1og=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gobwas/ws",
        sum = "h1:F2aeBZrm2NDsc7vbovKrWSogd4wvfAxg0FQ89/iqOTk=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_goccy_go_json",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/goccy/go-json",
        sum = "h1:Fq85nIqj+gXn/S5ahsiTlK3TmC85qgirsdTP/+DeaC4=",
        version = "v0.10.5",
    )
    go_repository(
        name = "com_github_goccy_go_yaml",
        build_file_proto_mode = "disable",
        importpath = "github.com/goccy/go-yaml",
        sum = "h1:n7Z+zx8S9f9KgzG6KtQKf+kwqXZlLNR2F6018Dgau54=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_gocql_gocql",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gocql/gocql",
        sum = "h1:6ImvI6U901e1ezn/8u2z3bh1DZIvMOia0yTSBxhy4Ao=",
        version = "v0.0.0-20220224095938-0eacd3183625",
    )
    go_repository(
        name = "com_github_godbus_dbus",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/godbus/dbus",
        sum = "h1:BWhy2j3IXJhjCbC68FptL43tDKIq8FladmaTs3Xs7Z8=",
        version = "v0.0.0-20190422162347-ade71ed3457e",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:mkgN1ofwASrYnJ5W6U/BxG15eXXXjirgZc7CLqkcaro=",
        version = "v5.0.6",
    )
    go_repository(
        name = "com_github_gofiber_fiber_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gofiber/fiber/v2",
        sum = "h1:ia0JaB+uw3GpNSCR5nvC5dsaxXjRU5OEu36aytx+zGw=",
        version = "v2.50.0",
    )
    go_repository(
        name = "com_github_gofrs_flock",
        build_file_proto_mode = "disable",
        importpath = "github.com/gofrs/flock",
        sum = "h1:+gYjHKf32LDeiEEFhQaotPbLuUXjY5ZqxKgXy7n59aw=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_github_gofrs_uuid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gofrs/uuid",
        sum = "h1:3qXRTX8/NbyulANqlc0lchS1gqAVxRgsuW1YrTJupqA=",
        version = "v4.4.0+incompatible",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/googleapis",
        sum = "h1:zgVt4UpGxcqVOw97aRGxT4svlcmdK35fynLNctY32zI=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_golang_freetype",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/freetype",
        sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
        version = "v0.0.0-20170609003504-e2365dfdc4a0",
    )
    go_repository(
        name = "com_github_golang_glog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/glog",
        sum = "h1:DrW6hGnjIhtvhOIiAKT6Psh/Kd/ldepEa81DKeiRJ5I=",
        version = "v1.2.5",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-jwt/jwt",
        sum = "h1:IfV12K8xAKAnZqdXVzCZ+TOjboZ2keLg81eXfW3O+oY=",
        version = "v3.2.2+incompatible",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:rcc4lwaZgFMCZ5jxF9ABolDcIHdBytAFgqFPbSJQAYs=",
        version = "v4.4.2",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-jwt/jwt/v5",
        sum = "h1:OuVbFODueb089Lh128TAcimifWaLhJwVflnrgM17wHk=",
        version = "v5.2.1",
    )
    go_repository(
        name = "com_github_golang_mock",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/mock",
        sum = "h1:YojYx61/OLFsiv6Rw1Z96LpldJIy31o+UHmwAUMJ6/U=",
        version = "v1.7.0-rc.1",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_golang_snappy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-sql/civil",
        sum = "h1:au07oEsX2xN0ktxqI+Sida1w446QrXBRJ0nee3SNZlA=",
        version = "v0.0.0-20220223132316-b832511892a9",
    )
    go_repository(
        name = "com_github_golang_sql_sqlexp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/golang-sql/sqlexp",
        sum = "h1:ZCD6MBpcuOVfGVqsEmY5/4FtYiKz6tSyUv9LPEDei6A=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gomodule_redigo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gomodule/redigo",
        sum = "h1:Sl3u+2BI/kk+VEatbj0scLdrFhjPmbxOc1myhDP41ws=",
        version = "v1.8.9",
    )
    go_repository(
        name = "com_github_google_btree",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/btree",
        sum = "h1:CVpQJjYgC4VbzxeGVHfvZrv1ctoYCAI8vbl07Fcxlyg=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_google_flatbuffers",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/flatbuffers",
        sum = "h1:M9dgRyhJemaM4Sw8+66GHBu8ioaQmyPLg1b8VwK5WJg=",
        version = "v23.5.26+incompatible",
    )
    go_repository(
        name = "com_github_google_gnostic_models",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gnostic-models",
        sum = "h1:yo/ABAfM5IMRsS1VnXjTBvUb61tFIHozhlYvRgGre9I=",
        version = "v0.6.8",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-cmp",
        sum = "h1:wk8382ETsv4JYUZwIsn6YpYiWiBsYLSJiTsyBybVuN8=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_google_go_containerregistry",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-containerregistry",
        sum = "h1:/+mFTs4AlwsJ/mJe8NDtKb7BxLtbZFpcn8vDsneEkwQ=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_google_go_jsonnet",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-jsonnet",
        sum = "h1:WG4TTSARuV7bSm4PMB4ohjxe33IHT5WVTrJSU33uT4g=",
        version = "v0.20.0",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:PVRnTgtArZ3QQqTGtbtjtnIkzl2iY2kt24yqbrf7td8=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/go-querystring",
        sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gofuzz",
        sum = "h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_google_martian",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/martian",
        sum = "h1:/CP5g8u/VJHijgedC/Legn3BAbAaWPgecwXBIDzw5no=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/martian/v3",
        sum = "h1:DIhPTQrbPkgs2yJYdXU/eNACCG5DVQjySNRNlflZ9Fc=",
        version = "v3.3.3",
    )
    go_repository(
        name = "com_github_google_pprof",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/pprof",
        sum = "h1:xRmpO92tb8y+Z85iUOMOicpCfaYcv7o3Cg3wKrIpg8g=",
        version = "v0.0.0-20240711041743-f6c9dda6c6da",
    )
    go_repository(
        name = "com_github_google_renameio",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/s2a-go",
        sum = "h1:LGD7gtMgezd8a/Xak7mEWL0PjoTQFvpRudN895yqKW0=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_google_shlex",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/shlex",
        sum = "h1:El6M4kTTCOh6aBiKaUGG7oYTSPP8MxqL4YI3kZKwcP4=",
        version = "v0.0.0-20191202100458-e7afc7fbc510",
    )
    go_repository(
        name = "com_github_google_tink_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/tink/go",
        sum = "h1:6Eox8zONGebBFcCBqkVmt60LaWZa6xg1cl/DwAh/J1w=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_google_uuid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_googleapis_cloud_bigtable_clients_test",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/cloud-bigtable-clients-test",
        sum = "h1:afMKTvA/jc6jSTMkeHBZGFDTt8Cc+kb1ATFzqMK85hw=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:GW/XbdyBFQ8Qe+YAmFU9uHLo7OnF5tL52HFAgMmyrf4=",
        version = "v0.3.6",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:eBLnkZ9635krYIPD+ag1USrOAI0Nr0QYF3+/3GqO0k0=",
        version = "v2.14.2",
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gnostic",
        sum = "h1:9fHAtK0uDfpveeqqo1hkEZJcFvYXAiCN3UutL8F9xHw=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_googleapis_go_type_adapters",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/go-type-adapters",
        sum = "h1:9XdMn+d/G57qq1s8dNc5IesGCXHf6V2HZ2JwRxfA2tA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:tlyzajkF3030q6M8SvmJSemC9DTHL/xaMa18b65+JM4=",
        version = "v0.0.0-20200911160855-bcd43fbb19e8",
    )
    go_repository(
        name = "com_github_googlecloudplatform_grpc_gcp_go_grpcgcp",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/grpc-gcp-go/grpcgcp",
        sum = "h1:DBjmt6/otSdULyJdVg2BlG0qGZO5tKL4VzOs0jpvw5Q=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_detectors_gcp",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp",
        sum = "h1:ErKg/3iS1AKcTkf3yixlZ54f9U1rljCkQyEXWUnIUxc=",
        version = "v1.27.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_exporter_metric",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric",
        sum = "h1:fYE9p3esPxA/C0rQ0AHhP0drtPXDRhaWiwg1DPqO7IU=",
        version = "v0.51.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_internal_cloudmock",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/cloudmock",
        sum = "h1:OqVGm6Ei3x5+yZmSJG1Mh2NwHvpVmZ08CB5qJhT9Nuk=",
        version = "v0.51.0",
    )
    go_repository(
        name = "com_github_googlecloudplatform_opentelemetry_operations_go_internal_resourcemapping",
        build_file_proto_mode = "disable",
        importpath = "github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping",
        sum = "h1:6/0iUd0xrnX7qt+mLNRwg5c0PGv8wpE8K90ryANQwMI=",
        version = "v0.51.0",
    )
    go_repository(
        name = "com_github_gophercloud_gophercloud",
        build_file_proto_mode = "disable",
        importpath = "github.com/gophercloud/gophercloud",
        sum = "h1:DTCNaTVGl8/cFu58O1JwWgis9gtISAFONqpMKNg/Vpw=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_gopherjs_gopherjs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gopherjs/gopherjs",
        sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
        version = "v0.0.0-20181017120253-0766667cb4d1",
    )
    go_repository(
        name = "com_github_gorilla_css",
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/css",
        sum = "h1:BQqNyPTi50JCFMTw/b67hByjMVXZRwGha6wxVGkeihY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:893HsJqtxp9z1SF76gg6hY70hRY1wVlTSnC/h1yUDCo=",
        version = "v0.0.0-20150720190736-60c7bfde3e33",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/mux",
        sum = "h1:TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_gorilla_securecookie",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/securecookie",
        sum = "h1:miw7JPhV+b/lAHSXz4qd/nN9jRiAFV5FwjeKyCS8BvQ=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_sessions",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/sessions",
        sum = "h1:DHd3rPN5lE3Ts3D8rKkQ8x/0kqfeNmBAaiSi+o7FsgI=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_grafana_otel_profiling_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/grafana/otel-profiling-go",
        sum = "h1:stVPKAFZSa7eGiqbYuG25VcqYksR6iWvF3YH66t4qL8=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_grafana_pyroscope_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/grafana/pyroscope-go",
        sum = "h1:VWBBlqxjyR0Cwk2W6UrE8CdcdD80GOFNutj0Kb1T8ac=",
        version = "v1.2.7",
    )
    go_repository(
        name = "com_github_grafana_pyroscope_go_godeltaprof",
        build_file_proto_mode = "disable",
        importpath = "github.com/grafana/pyroscope-go/godeltaprof",
        sum = "h1:c1Us8i6eSmkW+Ez05d3co8kasnuOY813tbMN8i/a3Og=",
        version = "v0.1.9",
    )
    go_repository(
        name = "com_github_grafana_regexp",
        build_file_proto_mode = "disable",
        importpath = "github.com/grafana/regexp",
        sum = "h1:GN2Lv3MGO7AS6PrRoT6yV5+wkrOpcszoIsO4+4ds248=",
        version = "v0.0.0-20240518133315-a468a5bfb3bc",
    )
    go_repository(
        name = "com_github_graph_gophers_graphql_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/graph-gophers/graphql-go",
        sum = "h1:fDqblo50TEpD0LY7RXk/LFVYEVqo3+tXMNMPSVXA1yc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_graphql_go_graphql",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/graphql-go/graphql",
        sum = "h1:p7/Ou/WpmulocJeEx7wjQy611rtXGQaAcXGqanuMMgc=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_github_graphql_go_handler",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/graphql-go/handler",
        sum = "h1:CANh8WPnl5M9uA25c2GBhPqJhE53Fg0Iue/fRNla71E=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/gregjones/httpcache",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:X5VWvz21y3gzm9Nw/kaUeku/1+uBhcekkmy4IkffJww=",
        version = "v2.27.1",
    )
    go_repository(
        name = "com_github_hailocab_go_hostpool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hailocab/go-hostpool",
        sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
        version = "v0.0.0-20160125115350-e80d13ce29ed",
    )
    go_repository(
        name = "com_github_hamba_avro",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hamba/avro",
        sum = "h1:/UBljlJ9hLjkcY7PhpI/bFYb4RMEXHEwHr17gAm/+l8=",
        version = "v1.5.6",
    )
    go_repository(
        name = "com_github_hamba_avro_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/hamba/avro/v2",
        sum = "h1:6PKpEWzJfNnvBgn7m2/8WYaDOUASxfDU+Jyb4ojDgFY=",
        version = "v2.17.2",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:P6slzxDLBOxUSj3fWo2o65VuKtbtOXFi7TSSgtXutuE=",
        version = "v1.29.4",
    )
    go_repository(
        name = "com_github_hashicorp_consul_proto_public",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/proto-public",
        sum = "h1:+DA/3g/IiKlJZb88NBn0ZgXrxJp2NlvCZdEyl+qxvL0=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_hashicorp_consul_sdk",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/consul/sdk",
        sum = "h1:V8TxTnImoPD5cj0U9Spl0TUxcytjcbbJeADFF07KdHg=",
        version = "v0.16.1",
    )
    go_repository(
        name = "com_github_hashicorp_cronexpr",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/cronexpr",
        sum = "h1:wG/ZYIKT+RT3QkOdgYc+xsKWVRgnxJ1OJtjjy84fJ9A=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:Qr2kF+eVWjTiYmU7Y31tYlP1h0q/X3Nl3tPGdaB11/k=",
        version = "v1.6.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_kms_wrapping_entropy_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-kms-wrapping/entropy/v2",
        sum = "h1:pSjQfW3vPtrOTcasTUKgCTQT7OGPPTTMVRrOfU6FJD8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_kms_wrapping_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-kms-wrapping/v2",
        sum = "h1:9Q2lu1YbbmiAgvYZ7Pr31RdlVonUpX+mmDL7Z7qTA2U=",
        version = "v2.0.8",
    )
    go_repository(
        name = "com_github_hashicorp_go_msgpack",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-msgpack",
        sum = "h1:i9R9JSrqIz0QVLz3sz+i3YJdT7TTSLcfLLzJi9aZTuI=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_net",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go.net",
        sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_plugin",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-plugin",
        sum = "h1:CHGwpxYDOttQOY7HOWgETU9dyVjOXzniXDqJcYJE1zM=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_hashicorp_go_retryablehttp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-retryablehttp",
        sum = "h1:C8hUCYzor8PIfXHa4UrZkU4VvK8o9ISHxT2Q8+VepXU=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_base62",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/base62",
        sum = "h1:ET4pqyjiGmY09R5y+rSd70J2w45CtbWDNvGqWp/R3Ng=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_mlock",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/mlock",
        sum = "h1:p4AKXPPS24tO8Wc8i1gLvSKdmkiSY5xuju57czJ/IJQ=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_parseutil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/parseutil",
        sum = "h1:UpiO20jno/eV1eVZcxqWnUohyKRe1g8FPV/xH1s/2qs=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_password",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/password",
        sum = "h1:6JzmBqXprakgFEHwBgdchsjaA9x3GyjdI568bXKxa60=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_strutil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/strutil",
        sum = "h1:kes8mmyCpxJsI7FTwtzRqEy9CdjCtrXrXGuOpxEA7Ts=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_secure_stdlib_tlsutil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-secure-stdlib/tlsutil",
        sum = "h1:phcbL8urUzF/kxA/Oj6awENaRwfWsjP59GW7u2qlDyY=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_sockaddr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-sockaddr",
        sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_syslog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-syslog",
        sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_version",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/go-version",
        sum = "h1:feTTfFNnjP967rlCxM/I9g701jU+RN74YKx2mOkIeek=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:dV3g9Z/unq5DpblPpw+Oqcv4dU/1omnb4Ok8iPY6p1c=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/golang-lru/v2",
        sum = "h1:a+bsQ5rvGLjzHuww6tVxozPZFVghXaHOwFs4luLUK2k=",
        version = "v2.0.7",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:kI3hhbbyzr4dldA8UdTb7ZlVVlI2DACdCfz31RPDgJM=",
        version = "v1.0.1-vault-5",
    )
    go_repository(
        name = "com_github_hashicorp_logutils",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/logutils",
        sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_mdns",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/mdns",
        sum = "h1:sY0CMhFmjIPDMlTB+HfymFHCaYLhgifZ0QhjaYKD/UQ=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_hashicorp_memberlist",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/memberlist",
        sum = "h1:EtYPN8DpAURiapus508I4n9CzHs2W+8NZGbmmR/prTM=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_hashicorp_nomad_api",
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/nomad/api",
        sum = "h1:fgVfQ4AC1avVOnu2cfms8VAiD8lUq3vWI8mTocOXN/w=",
        version = "v0.0.0-20240717122358-3d93bd3778f3",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:Z1H2J60yRKvfDYAOZLd2MU0ND4AH/WDz7xYHDWQsIPY=",
        version = "v0.10.1",
    )
    go_repository(
        name = "com_github_hashicorp_vault_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/vault/api",
        sum = "h1:YjkZLJ7K3inKgMZ0wzCU9OHqc+UqMQyXsPXnf3Cl2as=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_github_hashicorp_vault_sdk",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/vault/sdk",
        sum = "h1:H1kitfl1rG2SHbeGEyvhEqmIjVKE3E6c2q3ViKOs6HA=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_github_hashicorp_yamux",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hashicorp/yamux",
        sum = "h1:xixZ2bWeofWV68J+x6AzmKuVM/JWCQwkWm6GW/MUR6I=",
        version = "v0.0.0-20211028200310-0bc27b27de87",
    )
    go_repository(
        name = "com_github_heetch_avro",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/heetch/avro",
        sum = "h1:5PmgDy1cX/MegMy6btJ4bUFHgT5GLfSYfc5U7+JUQzg=",
        version = "v0.4.4",
    )
    go_repository(
        name = "com_github_hetznercloud_hcloud_go_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/hetznercloud/hcloud-go/v2",
        sum = "h1:jq0GP4QaYE5d8xR/Zw17s9qoaESRJMXfGmtD1a/qckQ=",
        version = "v2.13.1",
    )
    go_repository(
        name = "com_github_hexops_gotextdiff",
        build_file_proto_mode = "disable",
        importpath = "github.com/hexops/gotextdiff",
        sum = "h1:gitA9+qJrrTCsiCl7+kh75nPqQt1cx4ZkudSTLoUqJM=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hokaccha_go_prettyjson",
        build_file_proto_mode = "disable",
        importpath = "github.com/hokaccha/go-prettyjson",
        sum = "h1:7LYC+Yfkj3CTRcShK0KOL/w6iTiKyqqBA9a41Wnggw8=",
        version = "v0.0.0-20211117102719-0474bc63780f",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_iancoleman_orderedmap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/iancoleman/orderedmap",
        sum = "h1:i462o439ZjprVSFSZLZxcsoAe592sZB1rci2Z8j4wdk=",
        version = "v0.0.0-20190318233801-ac98e3ecb4b0",
    )
    go_repository(
        name = "com_github_iancoleman_strcase",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:nTXanmYxhfFAMjZL34Ov6gkzEsSJZ5DbhxWjvSASxEI=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:KwWnWVWCNtNq/ewIX7HIKnELmEx2nDP42yskD/pi7QE=",
        version = "v0.0.0-20240312041847-bd984b5ce465",
    )
    go_repository(
        name = "com_github_ibm_sarama",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/IBM/sarama",
        sum = "h1:lL01NNg/iBeigUbT+wpPysuTYW6roHo6kc1QrffRf0k=",
        version = "v1.40.1",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/imdario/mergo",
        sum = "h1:wwQJbIsHYGMUyLSPrEq1CT16AhnhNJQ51+4fdHUnCl4=",
        version = "v0.3.16",
    )
    go_repository(
        name = "com_github_imkira_go_interpol",
        build_file_proto_mode = "disable",
        importpath = "github.com/imkira/go-interpol",
        sum = "h1:KIiKr0VSG2CUW1hl1jpiyuzuJeKUUpC8iM1AIE7N1Vk=",
        version = "v1.1.0",
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
        name = "com_github_inhies_go_bytesize",
        build_file_proto_mode = "disable",
        importpath = "github.com/inhies/go-bytesize",
        sum = "h1:FtEj8sfIcaaBfAKrE1Cwb61YDtYq9JxChK1c7AKce7s=",
        version = "v0.0.0-20220417184213-4913239db9cf",
    )
    go_repository(
        name = "com_github_intel_goresctrl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/intel/goresctrl",
        sum = "h1:JyZjdMQu9Kl/wLXe9xA6s1X+tF6BWsQPFGJMEeCfWzE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_invopop_jsonschema",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/invopop/jsonschema",
        sum = "h1:2vgQcBz1n256N+FpX3Jq7Y17AjYt46Ig3zIWyy770So=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_ionos_cloud_sdk_go_v6",
        build_file_proto_mode = "disable",
        importpath = "github.com/ionos-cloud/sdk-go/v6",
        sum = "h1:mxxN+frNVmbFrmmFfXnBC3g2USYJrl6mc1LW2iNYbFY=",
        version = "v6.2.1",
    )
    go_repository(
        name = "com_github_j_keck_arping",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/j-keck/arping",
        sum = "h1:hlLhuXgQkzIJTZuhMigvG/CuSkaspeaD9hRDk2zuiMI=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_jackc_chunkreader",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/chunkreader",
        sum = "h1:4s39bBR8ByfqH+DKm8rQA3E1LHZWB9XWcrz8fqaZbe0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_chunkreader_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/chunkreader/v2",
        sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_jackc_pgconn",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgconn",
        sum = "h1:bVoTr12EGANZz66nZPkMInAV/KHD2TxH9npjXXgiB3w=",
        version = "v1.14.3",
    )
    go_repository(
        name = "com_github_jackc_pgio",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgio",
        sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgmock",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgmock",
        sum = "h1:DadwsjnMwFjfWc9y5Wi/+Zz7xoE5ALHsRQlOctkOiHc=",
        version = "v0.0.0-20210724152146-4ad1a8207f65",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgproto3",
        sum = "h1:FYYE4yRw+AgI8wXIinMlNjBbp/UitDJwfj5LqqewP1A=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgproto3/v2",
        sum = "h1:1HLSx5H+tXR9pW3in3zaztoEwQYRC9SQaYUHjTSUOag=",
        version = "v2.3.3",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:iCEnooe7UlwOQYpKFhBabPMi4aNAfoODPEFNiAnClxo=",
        version = "v0.0.0-20240606120523-5a60cdf6a761",
    )
    go_repository(
        name = "com_github_jackc_pgtype",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgtype",
        sum = "h1:y+xUdabmyMkJLyApYuPj38mW+aAIqCe5uuBB51rH3Vw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_jackc_pgx_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgx/v4",
        sum = "h1:xVpYkNR5pk5bMCZGfClbO962UIqVABcAGt7ha1s/FeU=",
        version = "v4.18.2",
    )
    go_repository(
        name = "com_github_jackc_pgx_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/pgx/v5",
        sum = "h1:rWQc5FwZSPX58r1OQmkuaNicxdmExaEz5A2DO2hUuTk=",
        version = "v5.7.6",
    )
    go_repository(
        name = "com_github_jackc_puddle",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/puddle",
        sum = "h1:eHK/5clGOatcjX3oWGBO/MpxpbHzSwud5EWTSCI+MX0=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jackc_puddle_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jackc/puddle/v2",
        sum = "h1:PR8nw+E/1w0GLuRFSmiioY6UooMp6KJv0/61nB7icHo=",
        version = "v2.2.2",
    )
    go_repository(
        name = "com_github_jandelgado_gcov2lcov",
        build_file_proto_mode = "disable",
        importpath = "github.com/jandelgado/gcov2lcov",
        sum = "h1:rkBt40h0CVK4oCb8Dps950gvfd1rYvQ8+cWa346lVU0=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_jbenet_go_base58",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jbenet/go-base58",
        sum = "h1:4zOlv2my+vf98jT1nQt4bT/yKWUImevYPJ2H344CloE=",
        version = "v0.0.0-20150317085156-6237cf65f3a6",
    )
    go_repository(
        name = "com_github_jcmturner_aescts_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/aescts/v2",
        sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_dnsutils_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/dnsutils/v2",
        sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_gofork",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/gofork",
        sum = "h1:QH0l3hzAU1tfT3rZCnW5zXl+orbkNMMRGJfdJjHVETg=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_github_jcmturner_goidentity_v6",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/goidentity/v6",
        sum = "h1:VKnZd2oEIMorCTsFBnJWbExfNN7yZr3EhJAxwOkZg6o=",
        version = "v6.0.1",
    )
    go_repository(
        name = "com_github_jcmturner_gokrb5_v8",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/gokrb5/v8",
        sum = "h1:x1Sv4HaTpepFkXbt2IkL29DXRf8sOfZXo8eRKh687T8=",
        version = "v8.4.4",
    )
    go_repository(
        name = "com_github_jcmturner_rpc_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jcmturner/rpc/v2",
        sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_jhump_gopoet",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jhump/gopoet",
        sum = "h1:gYjOPnzHd2nzB37xYQZxj4EIQNpBrBskRqQQ3q4ZgSg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_jhump_goprotoc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jhump/goprotoc",
        sum = "h1:Y1UgUX+txUznfqcGdDef8ZOVlyQvnV0pKWZH08RmZuo=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jhump/protoreflect",
        sum = "h1:N88q7JkxTHWFEqReuTsYH1dPIwXxA0ITNQp7avLY10s=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_jinzhu_gorm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jinzhu/gorm",
        sum = "h1:+IyIjPEABKRpsu/F8OvDPy9fyQlgsg2luMV2ZIH5i5o=",
        version = "v1.9.16",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jinzhu/now",
        sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath_internal_testify",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jmespath/go-jmespath/internal/testify",
        sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_jmoiron_sqlx",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jmoiron/sqlx",
        sum = "h1:vFFPA71p1o5gAeqtEAwLU4dnX2napprKtHr7PYIcN3g=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_joefitzgerald_rainbow_reporter",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/joefitzgerald/rainbow-reporter",
        sum = "h1:AuMG652zjdzI0YCCnXAqATtRBpGXMcAnrajcaTrSeuo=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_johncgriffin_overflow",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/JohnCGriffin/overflow",
        sum = "h1:RGWPOewvKIROun94nF7v2cua9qP+thov/7M50KEoeSU=",
        version = "v0.0.0-20211019200055-46fa312c352c",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/joho/godotenv",
        sum = "h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jonboulle/clockwork",
        sum = "h1:UOGuzwb1PwsrDAObMuhUnj0p5ULPj8V/xJ7Kx9qUBdQ=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_josharian_intern",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_jtolds_gls",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jtolds/gls",
        sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
        version = "v4.20.0+incompatible",
    )
    go_repository(
        name = "com_github_juju_qthttptest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/juju/qthttptest",
        sum = "h1:JPju5P5CDMCy8jmBJV2wGLjDItUsx2KKL514EfOYueM=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jung_kurt_gofpdf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/jung-kurt/gofpdf",
        sum = "h1:jgbatWHfRlPYiK85qgevsZTHviWXKwB1TTiKdz5PtRc=",
        version = "v1.16.2",
    )
    go_repository(
        name = "com_github_k0kubun_colorstring",
        build_file_proto_mode = "disable",
        importpath = "github.com/k0kubun/colorstring",
        sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
        version = "v0.0.0-20150214042306-9440f1994b88",
    )
    go_repository(
        name = "com_github_k0kubun_pp",
        build_file_proto_mode = "disable",
        importpath = "github.com/k0kubun/pp",
        sum = "h1:EKhKbi34VQDWJtq+zpsKSEhkHHs9w2P8Izbq8IhLVSo=",
        version = "v2.3.0+incompatible",
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
        name = "com_github_kevinmbeaulieu_eq_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kevinmbeaulieu/eq-go",
        sum = "h1:AQgYHURDOmnVJ62jnEk0W/7yFKEn+Lv8RHN6t7mB0Zo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_kimmachinegun_automemlimit",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/KimMachineGun/automemlimit",
        sum = "h1:RkbaC0MwhjL1ZuBKunGDjE/ggwAX43DwZrJqVwyveTk=",
        version = "v0.7.5",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_asmfmt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/asmfmt",
        sum = "h1:4Ri7ox3EwapiOjCki+hw14RyKk201CN4rzyCJRFLpK4=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/compress",
        sum = "h1:c/Cqfb0r+Yi+JtIEq73FWXVkRonBlf0CRNYc8Zttxdo=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_klauspost_cpuid_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/cpuid/v2",
        sum = "h1:0OwqZRYI2rFrjS4kvkDnqJkKHdHaRnCm68/DY4OxRzU=",
        version = "v2.2.11",
    )
    go_repository(
        name = "com_github_klauspost_reedsolomon",
        build_file_proto_mode = "disable",
        importpath = "github.com/klauspost/reedsolomon",
        sum = "h1:4cJuyH926If33BeDgiZpI5OU0pE+wUHZvMSyNGqN73Y=",
        version = "v1.12.5",
    )
    go_repository(
        name = "com_github_knadh_koanf_maps",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/maps",
        sum = "h1:G5TjmUh2D7G2YWf5SQQqSiHRJEjaicvU0KpypqB3NIs=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_knadh_koanf_parsers_json",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/parsers/json",
        sum = "h1:dzSZl5pf5bBcW0Acnu20Djleto19T0CfHcvZ14NJ6fU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_knadh_koanf_parsers_toml",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/parsers/toml",
        sum = "h1:S2hLqS4TgWZYj4/7mI5m1CQQcWurxUz6ODgOub/6LCI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_knadh_koanf_parsers_yaml",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/parsers/yaml",
        sum = "h1:ZZ8/iGfRLvKSaMEECEBPM1HQslrZADk8fP1XFUxVI5w=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_knadh_koanf_providers_posflag",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/providers/posflag",
        sum = "h1:mKJlLrKPcAP7Ootf4pBZWJ6J+4wHYujwipe7Ie3qW6U=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_knadh_koanf_providers_rawbytes",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/providers/rawbytes",
        sum = "h1:dpzgu2KO6uf6oCb4aP05KDmKmAmI51k5pe8RYKQ0qME=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_knadh_koanf_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/knadh/koanf/v2",
        sum = "h1:1dYGITt1I23x8cfx8ZnldtezdyaZtfAuRtIFOiRzK7g=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_knz_go_libedit",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/knz/go-libedit",
        sum = "h1:0pHpWtx9vcvC0xGZqEQlQdfSQs7WRlAjuPvk3fOZDCo=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_kolo_xmlrpc",
        build_file_proto_mode = "disable",
        importpath = "github.com/kolo/xmlrpc",
        sum = "h1:udzkj9S/zlT5X367kqJis0QP7YMxobob6zhzq6Yre00=",
        version = "v0.0.0-20220921171641-a4b6fa1dd06b",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_kr_fs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_logfmt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )
    go_repository(
        name = "com_github_kr_pretty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/pty",
        sum = "h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_kr_text",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_kylebanks_depth",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/KyleBanks/depth",
        sum = "h1:5h8fQADFrWtarTdtDudMmGsC7GPbOAu6RVB3ffsVFHc=",
        version = "v1.2.1",
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
        name = "com_github_labstack_echo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/labstack/echo",
        sum = "h1:pGRcYk231ExFAyoAjAfD85kQzRJCRI8bbnE7CX5OEgg=",
        version = "v3.3.10+incompatible",
    )
    go_repository(
        name = "com_github_labstack_echo_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/labstack/echo/v4",
        sum = "h1:dEpLU2FLg4UVmvCGPuk/APjlH6GDpbEPti61srUUUs4=",
        version = "v4.11.1",
    )
    go_repository(
        name = "com_github_labstack_gommon",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/labstack/gommon",
        sum = "h1:y7cvthEAEbU0yHOf4axH8ZG2NH8knB9iNSoTO8dyIk8=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_laher_mergefs",
        build_file_proto_mode = "disable",
        importpath = "github.com/laher/mergefs",
        sum = "h1:nV2bTS57vrmbMxeR6uvJpI8LyGl3QHj4bLBZO3aUV58=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_lann_builder",
        build_file_proto_mode = "disable",
        importpath = "github.com/lann/builder",
        sum = "h1:SOEGU9fKiNWd/HOJuq6+3iTQz8KNCLtVX6idSoTLdUw=",
        version = "v0.0.0-20180802200727-47ae307949d0",
    )
    go_repository(
        name = "com_github_lann_ps",
        build_file_proto_mode = "disable",
        importpath = "github.com/lann/ps",
        sum = "h1:P6pPBnrTSX3DEVR4fDembhRWSsG5rVo6hYhAB/ADZrk=",
        version = "v0.0.0-20150810152359-62de8c46ede0",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:WT9HwE9SGECu3lg4d/dIA+jxlljEa1/ffXKmRjqdmIQ=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_lestrrat_go_backoff_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat-go/backoff/v2",
        sum = "h1:oNb5E5isby2kiro9AgdHLv5N5tint1AnDVVf2E2un5A=",
        version = "v2.0.8",
    )
    go_repository(
        name = "com_github_lestrrat_go_blackmagic",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat-go/blackmagic",
        sum = "h1:Cg2gVSc9h7sz9NOByczrbUvLopQmXrfFx//N+AkAr5k=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_lestrrat_go_httpcc",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat-go/httpcc",
        sum = "h1:ydWCStUeJLkpYyjLDHihupbn2tYmZ7m22BGkcvZZrIE=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lestrrat_go_iter",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat-go/iter",
        sum = "h1:gMXo1q4c2pHmC3dn8LzRhJfP1ceCbgSiT9lUydIzltI=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_lestrrat_go_jwx",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat-go/jwx",
        sum = "h1:QT0utmUJ4/12rmsVQrJ3u55bycPkKqGYuGT4tyRhxSQ=",
        version = "v1.2.29",
    )
    go_repository(
        name = "com_github_lestrrat_go_option",
        build_file_proto_mode = "disable",
        importpath = "github.com/lestrrat-go/option",
        sum = "h1:oAzP2fvZGQKWkvHa1/SAcFolBEca1oN+mQ7eooNBEYU=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lib_pq",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/lib/pq",
        sum = "h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=",
        version = "v1.10.9",
    )
    go_repository(
        name = "com_github_linkedin_goavro",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/linkedin/goavro",
        sum = "h1:DV2aUlj2xZiuxQyvag8Dy7zjY69ENjS66bWkSfdpddY=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_linkedin_goavro_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/linkedin/goavro/v2",
        sum = "h1:4cuAtbDfqkKnBXp9E+tRkIJGa6W6iAjwonwt8O1f4U0=",
        version = "v2.11.1",
    )
    go_repository(
        name = "com_github_linode_linodego",
        build_file_proto_mode = "disable",
        importpath = "github.com/linode/linodego",
        sum = "h1:GcP7JIBr9iLRJ9FwAtb9/WCT1DuPJS/xUApapfdjtiY=",
        version = "v1.41.0",
    )
    go_repository(
        name = "com_github_linuxkit_virtsock",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/linuxkit/virtsock",
        sum = "h1:jUp75lepDg0phMUJBCmvaeFDldD2N3S1lBuPwUTszio=",
        version = "v0.0.0-20201010232012-f8cee7dfc7a3",
    )
    go_repository(
        name = "com_github_linxgnu_grocksdb",
        build_file_proto_mode = "disable",
        importpath = "github.com/linxGnu/grocksdb",
        sum = "h1:CzSS/vHLtVIdxdrjvqWR/sm93u/0eB6UcoO14YqObnw=",
        version = "v1.9.9",
    )
    go_repository(
        name = "com_github_logrusorgru_aurora_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/logrusorgru/aurora/v3",
        sum = "h1:R6zcoZZbvVcGMvDCKo45A9U/lzYyzl5NfYIvznmDfE4=",
        version = "v3.0.0",
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
        name = "com_github_luna_duclos_instrumentedsql",
        build_file_proto_mode = "disable",
        importpath = "github.com/luna-duclos/instrumentedsql",
        sum = "h1:t7mvC0z1jUt5A0UQ6I/0H31ryymuQRnJcWCiqV3lSAA=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/lyft/protoc-gen-star",
        sum = "h1:erE0rdztuaDq3bpGifD95wfoPrSZc95nGA6tbiNYh6M=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/lyft/protoc-gen-star/v2",
        sum = "h1:sIXJOMrYnQZJu7OB7ANSF4MYri2fTEGIsRLz6LwI4xE=",
        version = "v2.0.4-0.20230330145011-496ad1ac90a4",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/magiconair/properties",
        sum = "h1:IeQXZAiQcpL9mgcAe1Nu6cX9LLw6ExEHKjN0VQdvPDY=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/MakeNowJust/heredoc",
        sum = "h1:cXCdzVdstXyiTqTvfqk9SDHpKNjxuom+DOlyEeQ4pzQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_markbates_pkger",
        build_file_proto_mode = "disable",
        importpath = "github.com/markbates/pkger",
        sum = "h1:/MKEtWqtc0mZvu9OinB9UzVN9iYCwLWuyUv4Bw+PCno=",
        version = "v0.17.1",
    )
    go_repository(
        name = "com_github_marstr_guid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/marstr/guid",
        sum = "h1:/M4H/1G4avsieL6BbUwCOBzulmoeKVP5ux/3mQNnbyI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_masterminds_semver_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/semver/v3",
        sum = "h1:hLg3sBzpNErnxhQtUy/mmLR2I9foDujNK030IGemrRc=",
        version = "v3.1.1",
    )
    go_repository(
        name = "com_github_masterminds_squirrel",
        build_file_proto_mode = "disable",
        importpath = "github.com/Masterminds/squirrel",
        sum = "h1:uUcX/aBc8O7Fg9kaISIUsHXdKuqehiXAMQTYX8afzqM=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_matryer_moq",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/matryer/moq",
        sum = "h1:RtpiPUM8L7ZSCbSwK+QcZH/E9tgqAkFjKQxsRs25b4w=",
        version = "v0.2.7",
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
        sum = "h1:E5ScNMtiwvlvB5paMFdw9p4kSQzbXFikJ5SQO6TULQc=",
        version = "v0.0.16",
    )
    go_repository(
        name = "com_github_mattn_go_shellwords",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-shellwords",
        sum = "h1:M2zGm7EW6UQJvDeQxo4T51eKPurbeFbe8WtebGE2xrk=",
        version = "v1.0.12",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:yOQRA0RpS5PFz/oikGwBEqvAWhWg5ufRz4ETLjwpU1Y=",
        version = "v1.14.16",
    )
    go_repository(
        name = "com_github_mattn_goveralls",
        build_file_proto_mode = "disable",
        importpath = "github.com/mattn/goveralls",
        sum = "h1:PEEeF0k1SsTjOBQ8FOmrOAoCu4ytuMaWCnWe94zxbCg=",
        version = "v0.0.12",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_maxbrunsfeld_counterfeiter_v6",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/maxbrunsfeld/counterfeiter/v6",
        sum = "h1:g+4J5sZg6osfvEfkRZxJ1em0VT95/UOZgi/l7zi1/oE=",
        version = "v6.2.2",
    )
    go_repository(
        name = "com_github_mdlayher_socket",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/socket",
        sum = "h1:eM9y2/jlbs1M615oshPQOHZzj6R6wMT7bX5NPiQvn2U=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_mdlayher_vsock",
        build_file_proto_mode = "disable",
        importpath = "github.com/mdlayher/vsock",
        sum = "h1:pC1mTJTvjo1r9n9fbm7S1j04rCgCzhCOS5DY0zqHlnQ=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:5PJl274Y63IEHC+7izoQE9x6ikvDFZS2mDVS3drnohI=",
        version = "v0.0.0-20200706080929-d51e80ef957d",
    )
    go_repository(
        name = "com_github_microcosm_cc_bluemonday",
        build_file_proto_mode = "disable",
        importpath = "github.com/microcosm-cc/bluemonday",
        sum = "h1:xbqSvqzQMeEHCqMi64VAs4d8uy6Mequs3rQ0k/Khz58=",
        version = "v1.0.26",
    )
    go_repository(
        name = "com_github_microsoft_go_mssqldb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/microsoft/go-mssqldb",
        sum = "h1:p2rpHIL7TlSv1QrbXJUAcbyRKnIT0C9rRkH2E4OjLn8=",
        version = "v0.21.0",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:F2VQgta7ecxGYO8k3ZZz3RS8fVIXVxONVUPlNERoyfY=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/hcsshim",
        sum = "h1:mnUj0ivWy6UzbB1uLFqKR6F+ZyiDc7j4iGgHTpO+5+I=",
        version = "v0.9.4",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim_test",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/hcsshim/test",
        sum = "h1:4FA+QBaydEHlwxg0lMN3rhwoDaQy6LKhVWR4qvq4BuA=",
        version = "v0.0.0-20210227013316-43a75bb4edd3",
    )
    go_repository(
        name = "com_github_miekg_dns",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/miekg/dns",
        sum = "h1:cN8OuEF1/x5Rq6Np+h1epln8OiyPWV+lROx9LxcGgIQ=",
        version = "v1.1.62",
    )
    go_repository(
        name = "com_github_miekg_pkcs11",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/miekg/pkcs11",
        sum = "h1:Ugu9pdy6vAYku5DEpVWVFPYnzV+bxB+iRdbuFSu7TvU=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_mileusna_useragent",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mileusna/useragent",
        sum = "h1:p3RJWhi3LfuI6BHdddojREyK3p6qX67vIfOVMnUIVr0=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_minio_asm2plan9s",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/minio/asm2plan9s",
        sum = "h1:AMFGa4R4MiIpspGNG7Z948v4n35fFGB3RR3G/ry4FWs=",
        version = "v0.0.0-20200509001527-cdd76441f9d8",
    )
    go_repository(
        name = "com_github_minio_c2goasm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/minio/c2goasm",
        sum = "h1:+n/aFZefKZp7spd8DFdX7uMikMLXX4oubIzJF4kv/wI=",
        version = "v0.0.0-20190812172519-36a3d3bbc4f3",
    )
    go_repository(
        name = "com_github_mistifyio_go_zfs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mistifyio/go-zfs",
        sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
        version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
    )
    go_repository(
        name = "com_github_mitchellh_cli",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/cli",
        sum = "h1:tEElEatulEHDeedTxwckzyYMA5c86fbmNIUL1hBIiTg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_copystructure",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/copystructure",
        sum = "h1:vpKXTN4ewci03Vljg/q9QvCGUDttBOGBIa15WveJJGw=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_testing_interface",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-testing-interface",
        sum = "h1:jrgshOhYAUVNMAJiKbEu7EqAwgJJ2JqpQmpLJOu07cU=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mitchellh_gox",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/gox",
        sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_mitchellh_iochan",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/iochan",
        sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:jeMsZIYE/09sWLaz43PL7Gy6RuMjD2eJVyuac5Z2hdY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_mitchellh_osext",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/osext",
        sum = "h1:2+myh5ml7lgEU/51gbeLHfKGNfgEQQIWrlbdaOsidbQ=",
        version = "v0.0.0-20151018003038-5e2d6d41470f",
    )
    go_repository(
        name = "com_github_mitchellh_reflectwalk",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mitchellh/reflectwalk",
        sum = "h1:G2LzWKi524PWgd3mLHV8Y5k7s6XUvT0Gef6zxSIeXaQ=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_moby_docker_image_spec",
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/docker-image-spec",
        sum = "h1:jMKff3w6PgbfSa69GfNg+zN/XLhfXJGnEx3Nl2EsFP0=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_moby_locker",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/locker",
        sum = "h1:fOXqR41zeveg4fFODix+1Ch4mj/gT0NE1XJbp/epuBg=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_moby_patternmatcher",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/patternmatcher",
        sum = "h1:YCZgJOeULcxLw1Q+sVR636pmS7sPEn1Qo2iAN6M7DBo=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_moby_spdystream",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/spdystream",
        sum = "h1:Vy79D6mHeJJjiPdFEL2yku1kl0chZpJfZcPpb16BRl8=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_moby_sys_mount",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/mount",
        sum = "h1:fX1SVkXFJ47XWDoeFW4Sq7PdQJnV2QIDZAqjNqgEjUs=",
        version = "v0.3.3",
    )
    go_repository(
        name = "com_github_moby_sys_mountinfo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/mountinfo",
        sum = "h1:BzJjoreD5BMFNmD9Rus6gdd1pLuecOFPt8wC+Vygl78=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_moby_sys_sequential",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/sequential",
        sum = "h1:OPvI35Lzn9K04PBbCLW0g4LcFAJgHsvXsRyewg5lXtc=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_moby_sys_signal",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/signal",
        sum = "h1:aDpY94H8VlhTGa9sNYUFCFsMZIUh5wm0B6XkIoJj/iY=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_moby_sys_symlink",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/sys/symlink",
        sum = "h1:tk1rOM+Ljp0nFmfOIBtlV3rTDlWOwFRhjEeAhZB0nZc=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_moby_term",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/moby/term",
        sum = "h1:xt8Q1nalod/v7BqbG21f8mQPqH+xAaC9C3N3wfWbVP0=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_modocache_gover",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/modocache/gover",
        sum = "h1:8Q0qkMVC/MmWkpIdlvZgcv2o2jrlF6zqVOh7W5YHdMA=",
        version = "v0.0.0-20171022184752-b58185e213c5",
    )
    go_repository(
        name = "com_github_mohae_deepcopy",
        build_file_proto_mode = "disable",
        importpath = "github.com/mohae/deepcopy",
        sum = "h1:RWengNIwukTxcDr9M+97sNutRR1RKhG96O6jWumTTnw=",
        version = "v0.0.0-20170929034955-c48cc78d4826",
    )
    go_repository(
        name = "com_github_montanaflynn_stats",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/montanaflynn/stats",
        sum = "h1:Duep6KMIDpY4Yo11iFsvyqJDyfzLF9+sndUKT+v64GQ=",
        version = "v0.6.6",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_moul_http2curl",
        build_file_proto_mode = "disable",
        importpath = "github.com/moul/http2curl",
        sum = "h1:Pip12xNtMvEFUBF4f8/b5yRXj94LLrNdLWELfOr2KcY=",
        version = "v0.0.0-20170919181001-9ac6cf4d929b",
    )
    go_repository(
        name = "com_github_mr_tron_base58",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mr-tron/base58",
        sum = "h1:T/HDJBh4ZCPbU39/+c3rRvE0uKBQlU27+QI8LJ4t64o=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_mrunalp_fileutils",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mrunalp/fileutils",
        sum = "h1:NKzVxiH7eSk+OQ4M+ZYW1K6h27RUV3MI6NUTsHhU6Z4=",
        version = "v0.5.0",
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
        name = "com_github_munnerz_goautoneg",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/mxk/go-flowrate",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
        version = "v0.0.0-20140419014527-cca7078d478f",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:pSFyXApG+yWU/TgbKCjmm5K4wrHu86231/w84qRVR+U=",
        version = "v1.48.0",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:q44qGV008kYd9W1b1nEBkNzvnWxtRSQ7A8BoqRrcfa0=",
        version = "v0.4.11",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_ncw_swift",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ncw/swift",
        sum = "h1:4DQRPj35Y41WogBxyhOXlrI37nzGlyEcsforeudyYPQ=",
        version = "v1.0.47",
    )
    go_repository(
        name = "com_github_networkplumbing_go_nft",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/networkplumbing/go-nft",
        sum = "h1:eKapmyVUt/3VGfhYaDos5yeprm+LPt881UeksmKKZHY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )
    go_repository(
        name = "com_github_nrwiersma_avro_benchmarks",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/nrwiersma/avro-benchmarks",
        sum = "h1:wDbc54qVQ+C5oQZ8Q5VlMbqEt2hrnev2bC/gIGL3Ksk=",
        version = "v0.0.0-20210913175520-21aec48c8f76",
    )
    go_repository(
        name = "com_github_nsf_jsondiff",
        build_file_proto_mode = "disable",
        importpath = "github.com/nsf/jsondiff",
        sum = "h1:dOYG7LS/WK00RWZc8XGgcUTlTxpp3mKhdR2Q9z9HbXM=",
        version = "v0.0.0-20230430225905-43f6cf3098c1",
    )
    go_repository(
        name = "com_github_nvveen_gotty",
        build_file_proto_mode = "disable",
        importpath = "github.com/Nvveen/Gotty",
        sum = "h1:TngWCqHvy9oXAN6lEVMRuU21PR1EtLVZJmdB18Gu3Rw=",
        version = "v0.0.0-20120604004816-cd527374f1e5",
    )
    go_repository(
        name = "com_github_nxadm_tail",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/nxadm/tail",
        sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_nyaruka_phonenumbers",
        build_file_proto_mode = "disable",
        importpath = "github.com/nyaruka/phonenumbers",
        sum = "h1:DcueYq7QrOArAprAYNoQfDgp0KetO4LqtnBtQC6Wyes=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/NYTimes/gziphandler",
        sum = "h1:ZUDjpQae29j0ryrS0u/B8HZfJBtBQHjqw2rQ2cqUQ3I=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_oklog_run",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/run",
        sum = "h1:GEenZ1cK0+q0+wsJew9qUg/DyD8k3JzYsZAi5gYi2mA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_oklog_ulid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/ulid",
        sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_oklog_ulid_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/oklog/ulid/v2",
        sum = "h1:+9lhoxAP56we25tyYETBBY1YLA2SaoLvUFgrP2miPJU=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_oleiade_reflections",
        build_file_proto_mode = "disable",
        importpath = "github.com/oleiade/reflections",
        sum = "h1:D1XO3LVEYroYskEsoSiGItp9RUxG6jWnCVvrqH0HHQM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:58+kh9C6jJVXYjt8IE48G2eWl6BjwU5Gj0gqY84fy78=",
        version = "v0.0.0-20170122224234-a0225b3f23b5",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:KMrpdQIwFcEqXDklaen+P1axHaj9BSKzvpUUfnHldSE=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
        version = "v1.16.5",
    )
    go_repository(
        name = "com_github_onsi_ginkgo_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/ginkgo/v2",
        sum = "h1:9Cnnf7UHo57Hy3k6/m5k3dRfGTMXGvxhHFvkDTCTpvA=",
        version = "v2.19.0",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/onsi/gomega",
        sum = "h1:dsYjIxxSR755MDmKVsaFQTE22ChNBcuuTWgkUDSubOk=",
        version = "v1.33.1",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:8SG7/vwALn54lVB/0yZ/MMwhFrPYtpEHQb2IpWsCzug=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/runc",
        sum = "h1:rgSuzbmgz5DUJjeSnw337TxDbRuqjs6iqQck/2weR6w=",
        version = "v1.1.14",
    )
    go_repository(
        name = "com_github_opencontainers_runtime_spec",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/runtime-spec",
        sum = "h1:3snG66yBm59tKhhSPQrQ/0bCrv1LQbKt40LnUPiUxdc=",
        version = "v1.0.3-0.20210326190908-1c3f411f0417",
    )
    go_repository(
        name = "com_github_opencontainers_runtime_tools",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/runtime-tools",
        sum = "h1:H7DMc6FAjgwZZi8BRqjrAAHWoqEr5e5L6pS4V0ezet4=",
        version = "v0.0.0-20181011054405-1d69bd0f9c39",
    )
    go_repository(
        name = "com_github_opencontainers_selinux",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opencontainers/selinux",
        sum = "h1:09LIPVRP3uuZGQvgR+SgMSNBd1Eb3vlRbGqQpoHsF8w=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_openzipkin_zipkin_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/openzipkin/zipkin-go",
        sum = "h1:zjqfqHjUpPmB3c1GlCvvgsM1G4LkvqQbBDueDOCg/jA=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_ory_analytics_go_v5",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/analytics-go/v5",
        sum = "h1:LX8T5B9FN8KZXOtxgN+R3I4THRRVB6+28IKgKBpXmAM=",
        version = "v5.0.1",
    )
    go_repository(
        name = "com_github_ory_dockertest_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/dockertest/v3",
        sum = "h1:By96ZSVuH5LyjXLVVMfvJoLVGHaT96LdOnwgFSLVf0E=",
        version = "v3.10.1-0.20240704115616-d229e74b748d",
    )
    go_repository(
        name = "com_github_ory_fosite",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/fosite",
        sum = "h1:KNqO7RVt/1X8F08/UI0Y+GRvcpscCWgjqvpLBQPRovo=",
        version = "v0.49.0",
    )
    go_repository(
        name = "com_github_ory_go_acc",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/go-acc",
        sum = "h1:rvu4obdvqR0fkSIJ8IfgzKOWwZ5kOT2UNfLq81Qk7rc=",
        version = "v0.2.9-0.20230103102148-6b1c9a70dbbe",
    )
    go_repository(
        name = "com_github_ory_go_convenience",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/go-convenience",
        sum = "h1:zouLKfF2GoSGnJwGq+PE/nJAE6dj2Zj5QlTgmMTsTS8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_ory_herodot",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/herodot",
        sum = "h1:gGvNMHgAwWzdP/eo+roSiT5CGssygHSjDU7MSQNlJ4E=",
        version = "v0.10.2",
    )
    go_repository(
        name = "com_github_ory_jsonschema_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/jsonschema/v3",
        sum = "h1:Ssdb3eJ4lDZ/+XnGkvQS/te0p+EkolqwTsDOCxr/FmU=",
        version = "v3.0.8",
    )
    go_repository(
        name = "com_github_ory_x",
        build_file_proto_mode = "disable",
        importpath = "github.com/ory/x",
        sum = "h1:61vv0ObCDSX1vOQYbxBeqDiv4YiPmMT91lYxDaaKX08=",
        version = "v0.0.665",
    )
    go_repository(
        name = "com_github_otiai10_copy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/otiai10/copy",
        sum = "h1:hVoPiN+t+7d2nzzwMiDHPSOogsWAStewq3TwU05+clE=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_otiai10_curr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/otiai10/curr",
        sum = "h1:TJIWdbX0B+kpNagQrjgq8bCMrbhiuX73M2XwgtDMoOI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_otiai10_mint",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/otiai10/mint",
        sum = "h1:7JgpsBaN0uMkyju4tbYHu0mnM55hNKVYLsXmwr15NQI=",
        version = "v1.3.3",
    )
    go_repository(
        name = "com_github_outcaste_io_ristretto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/outcaste-io/ristretto",
        sum = "h1:AK4zt/fJ76kjlYObOeNwh4T3asEuaCmp26pOvUOL9w0=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_ovh_go_ovh",
        build_file_proto_mode = "disable",
        importpath = "github.com/ovh/go-ovh",
        sum = "h1:ixLOwxQdzYDx296sXcgS35TOPEahJkpjMGtzPadCjQI=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_parnurzeal_gorequest",
        build_file_proto_mode = "disable",
        importpath = "github.com/parnurzeal/gorequest",
        sum = "h1:oPjDCsF5IkD4gUk6vIgsxYNaSgvAnIh1EJeROn3HdJU=",
        version = "v0.2.15",
    )
    go_repository(
        name = "com_github_pascaldekloe_goe",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pascaldekloe/goe",
        sum = "h1:cBOtyMzM9HTpWjXfbbunk26uA6nG3a8n06Wieeh0MwY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_pbnjay_memory",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pbnjay/memory",
        sum = "h1:onHthvaw9LFnH4t2DcNVpwGmV9E1BkGknEliJkfwQj0=",
        version = "v0.0.0-20210728143218-7b4eea64cf58",
    )
    go_repository(
        name = "com_github_pborman_getopt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pborman/getopt",
        sum = "h1:BHT1/DKsYDGkUgQ2jmMaozVcdk+sVfz0+1ZJq4zkWgw=",
        version = "v0.0.0-20170112200414-7148bc3a4c30",
    )
    go_repository(
        name = "com_github_pelletier_go_toml",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:4yBQzkHv+7BHq2PQUZF3Mx0IYxG7LsP222s7Agd3ve8=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:mye9XuhQ6gvn5h28+VilKrrPoQVanw5PMw/TB0t5Ec4=",
        version = "v2.2.4",
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/peterbourgon/diskv",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_peterhellberg_link",
        build_file_proto_mode = "disable",
        importpath = "github.com/peterhellberg/link",
        sum = "h1:UA5pg3Gp/E0F2WdX7GERiNrPQrM1K6CVJUUWfHa4t6c=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_philhofer_fwd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/philhofer/fwd",
        sum = "h1:bnDivRJ1EWPjUIRXV5KfORO897HTbpFAQddBdE8t7Gw=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_phpdave11_gofpdf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/phpdave11/gofpdf",
        sum = "h1:KPKiIbfwbvC/wOncwhrpRdXVj2CZTCFlw4wnoyjtHfQ=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_phpdave11_gofpdi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/phpdave11/gofpdi",
        sum = "h1:o61duiW8M9sMlkVXWlvP92sZJtGKENvW3VExs6dZukQ=",
        version = "v1.0.13",
    )
    go_repository(
        name = "com_github_pierrec_lz4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pierrec/lz4",
        sum = "h1:9UY3+iC23yxF0UfGaYrGplQ+79Rg+h/q9FV9ix19jjM=",
        version = "v2.6.1+incompatible",
    )
    go_repository(
        name = "com_github_pierrec_lz4_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pierrec/lz4/v4",
        sum = "h1:xaKrnTkyoqfh1YItXl56+6KJNVYWlEEPuAQW9xsplYQ=",
        version = "v4.1.18",
    )
    go_repository(
        name = "com_github_pkg_browser",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/browser",
        sum = "h1:+mdjkGKdHQG3305AYmdv1U2eRNDiU2ErMBj1gwrq8eQ=",
        version = "v0.0.0-20240102092130-5ac0b6a4141c",
    )
    go_repository(
        name = "com_github_pkg_diff",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_pkg_errors",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_profile",
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/profile",
        sum = "h1:hnbDkaNWPCLMO9wGLdBFTIZvzDrDfBM2072E1S9gJkA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pkg/sftp",
        sum = "h1:I2qBYMChEhIjOgazfJmV3/mZM256btk6wkCDRmW7JYs=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_planetscale_vtprotobuf",
        build_file_proto_mode = "disable",
        importpath = "github.com/planetscale/vtprotobuf",
        sum = "h1:GFCKgmp0tecUJ0sJuv4pzYCqS9+RGSn52M3FUwPs+uo=",
        version = "v0.6.1-0.20240319094008-0393e58bdf10",
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
        name = "com_github_posener_complete",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/posener/complete",
        sum = "h1:NP0eAhjcjImqslEwo/1hq7gpajME0fTLTezBKDqfXqo=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pquerna/cachecontrol",
        sum = "h1:vBXSNuE5MYP9IJ5kjsdo8uq+w41jSPgvba2DEnkRx9k=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_pquerna_otp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/pquerna/otp",
        sum = "h1:oJV/SkzR33anKXwQU3Of42rL4wbrffP4uvUf1SvS5Xs=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_prometheus_alertmanager",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/alertmanager",
        sum = "h1:V6nTa2J5V4s8TG4C4HtrBP/WNSebCCTYGGv4qecA/+I=",
        version = "v0.27.0",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:rb93p9lokFEsctTys46VnV1kLCDpVZ0a/Y92Vm0Zc6Q=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:ZKSh/rekM+n3CeS952MLRAdFwIKqeY8b62p8ais2e9E=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_prometheus_common",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/common",
        sum = "h1:xasJaQlnWAeyHdUBeGjXmutelfJHWMRr+Fg4QszZ2Io=",
        version = "v0.62.0",
    )
    go_repository(
        name = "com_github_prometheus_common_assets",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/common/assets",
        sum = "h1:0P5OrzoHrYBOSM1OigWL3mY8ZvV2N4zIE/5AahrSrfM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_prometheus_common_sigv4",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/common/sigv4",
        sum = "h1:qoVebwtwwEhS85Czm2dSROY5fTo2PAPEVdDeppTwGX4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_prometheus_exporter_toolkit",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/exporter-toolkit",
        sum = "h1:lmA0Q+8IaXgmFRKw09RldZmZdnvu9wwcDLIXGmTPw1c=",
        version = "v0.13.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:YagwOFzUgYfKKHX6Dr+sHT7km/hxC76UB0learggepc=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_prometheus_prometheus",
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/prometheus",
        sum = "h1:9KKcTTq80gkzmXW0Et/QCFSrBPgmwiS3Hlcxc6o8KlM=",
        version = "v0.300.1",
    )
    go_repository(
        name = "com_github_prometheus_tsdb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/prometheus/tsdb",
        sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
        version = "v0.7.1",
    )
    go_repository(
        name = "com_github_puerkitobio_goquery",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/goquery",
        sum = "h1:PSPBGne8NIUWw+/7vFBV+kG2J/5MOjbzc7154OaKCSE=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_puerkitobio_purell",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/purell",
        sum = "h1:WEQqlqaGbrPkxLJWfBwQmfEAE1Z7ONdDLqrN38tNFfI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_puerkitobio_urlesc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/PuerkitoBio/urlesc",
        sum = "h1:d+Bc7a5rLufV/sSk/8dngufqelfh6jnri85riMAaF/M=",
        version = "v0.0.0-20170810143723-de5bf2ad4578",
    )
    go_repository(
        name = "com_github_rabbitmq_amqp091_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rabbitmq/amqp091-go",
        sum = "h1:STpn5XsHlHGcecLmMFCtg7mqq0RnD+zFr4uzukfVhBw=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_rakutentech_jwk_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/rakutentech/jwk-go",
        sum = "h1:PiLwepKyUaW+QFG3ki78DIO2+b4IVK3nMhlxM70zrQ4=",
        version = "v1.1.3",
    )
    go_repository(
        name = "com_github_rcrowley_go_metrics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rcrowley/go-metrics",
        sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
        version = "v0.0.0-20201227073835-cf1acfcdf475",
    )
    go_repository(
        name = "com_github_redis_go_redis_v9",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/redis/go-redis/v9",
        sum = "h1:u4tNCjXOyzfgeLN+vAZaW1xUooqWDqVEsZN0U01jfAE=",
        version = "v9.14.0",
    )
    go_repository(
        name = "com_github_remyoudompheng_bigfft",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/remyoudompheng/bigfft",
        sum = "h1:W09IVJc94icq4NjY3clb7Lk8O1qJ8BdBEF8z0ibU0rE=",
        version = "v0.0.0-20230129092748-24d4a6f8daec",
    )
    go_repository(
        name = "com_github_richardartoul_molecule",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/richardartoul/molecule",
        sum = "h1:Qp27Idfgi6ACvFQat5+VJvlYToylpM/hcyLBI3WaKPA=",
        version = "v1.0.1-0.20221107223329-32cfee06a052",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:WUdvkW8uEhrYfLC4ZzdpI2ztxP1I582+49Oc5Mq64VQ=",
        version = "v0.4.7",
    )
    go_repository(
        name = "com_github_rogpeppe_clock",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/clock",
        sum = "h1:3QH7VyOaaiUHNrA9Se4YQIRkDTCw1EJls9xTUCaCeRM=",
        version = "v0.0.0-20190514195947-2896927a307a",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:KvO1DLK/DRN07sQ1LQKScxyZJuNnedQ5/wKSR38lUII=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_rs_cors",
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/cors",
        sum = "h1:0B9GE/r9Bc2UxRMMtymBkHTenPkHDv0CW4Y98GBY+po=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/xid",
        sum = "h1:fV591PaemRlL6JfRxGDEPl69wICngIQ3shQtzfy2gxU=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rs/zerolog",
        sum = "h1:k43nTLIwcTVQAncfCw4KZ2VY6ukYoZaBPNOE8txlOeY=",
        version = "v1.34.0",
    )
    go_repository(
        name = "com_github_rueian_rueidis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/rueian/rueidis",
        sum = "h1:cG905akj2+QyHx0x9y4mN0K8vLi6M94QiyoLulXS3l0=",
        version = "v0.0.93",
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
        name = "com_github_ruudk_golang_pdf417",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ruudk/golang-pdf417",
        sum = "h1:K1Xf3bKttbF+koVGaX5xngRIZ5bVjbmPnaxE/dR08uY=",
        version = "v0.0.0-20201230142125-a7e3863a1245",
    )
    go_repository(
        name = "com_github_ryanuber_columnize",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ryanuber/columnize",
        sum = "h1:j1Wcmh8OrK4Q7GXY+V7SVSY8nUWQxHW5TkBe7YUl+2s=",
        version = "v2.1.0+incompatible",
    )
    go_repository(
        name = "com_github_ryanuber_go_glob",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ryanuber/go-glob",
        sum = "h1:iQh3xXAumdQ+4Ufa5b25cRpC5TYKlno6hsv6Cb3pkBk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_safchain_ethtool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/safchain/ethtool",
        sum = "h1:ZFfeKAhIQiiOrQaI3/znw0gOmYpO28Tcu1YaqMa/jtQ=",
        version = "v0.0.0-20210803160452-9aa261dae9b1",
    )
    go_repository(
        name = "com_github_sagikazarmark_crypt",
        build_file_proto_mode = "disable",
        importpath = "github.com/sagikazarmark/crypt",
        sum = "h1:96E1qrToLBU6fGzo+PRRz7KGOc9FkYFiPnR3/zf8Smg=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_sahilm_fuzzy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sahilm/fuzzy",
        sum = "h1:ceu5RHF8DGgoi+/dR5PsECjCDH1BE3Fnmpo7aVXOdRA=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_sanity_io_litter",
        build_file_proto_mode = "disable",
        importpath = "github.com/sanity-io/litter",
        sum = "h1:iE+sBxPBzoK6uaEP5Lt3fHNgpKcHXc/A2HGETy0uJQo=",
        version = "v1.5.5",
    )
    go_repository(
        name = "com_github_santhosh_tekuri_jsonschema_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/santhosh-tekuri/jsonschema/v5",
        sum = "h1:WCcC4vZDS1tYNxjWlwRJZQy28r8CMoggKnxNzxsVDMQ=",
        version = "v5.2.0",
    )
    go_repository(
        name = "com_github_satori_go_uuid",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/satori/go.uuid",
        sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_savsgio_gotils",
        build_file_proto_mode = "disable",
        importpath = "github.com/savsgio/gotils",
        sum = "h1:N3Af8f13ooDKcIhsmFT7Z05CStZWu4C7Md0uDEy4q6o=",
        version = "v0.0.0-20210617111740-97865ed5a873",
    )
    go_repository(
        name = "com_github_scaleway_scaleway_sdk_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/scaleway/scaleway-sdk-go",
        sum = "h1:yoKAVkEVwAqbGbR8n87rHQ1dulL25rKloGadb3vm770=",
        version = "v1.0.0-beta.30",
    )
    go_repository(
        name = "com_github_sclevine_agouti",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sclevine/agouti",
        sum = "h1:8IBJS6PWz3uTlMP3YBIR5f+KAldcGuOeFkFbUWfBgK4=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_sclevine_spec",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sclevine/spec",
        sum = "h1:1Jwdf9jSfDl9NVmt8ndHqbTZ7XCCPbh1jI3hkDBHVYA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sean_seed",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sean-/seed",
        sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
        version = "v0.0.0-20170313163322-e2103e2c3529",
    )
    go_repository(
        name = "com_github_seatgeek_logrus_gelf_formatter",
        build_file_proto_mode = "disable",
        importpath = "github.com/seatgeek/logrus-gelf-formatter",
        sum = "h1:0b8DF5kR0PhRoRXDiEEdzrgBc8UqVY4JWLkQJCRsLME=",
        version = "v0.0.0-20210414080842-5b05eb8ff761",
    )
    go_repository(
        name = "com_github_seccomp_libseccomp_golang",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/seccomp/libseccomp-golang",
        sum = "h1:RpforrEYXWkmGwJHIGnLZ3tTWStkjVVstwzNGqxX2Ds=",
        version = "v0.9.2-0.20220502022130-f33da4d89646",
    )
    go_repository(
        name = "com_github_secure_systems_lab_go_securesystemslib",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/secure-systems-lab/go-securesystemslib",
        sum = "h1:OwvJ5jQf9LnIAS83waAjPbcMsODrTQUpJ02eNLUoxBg=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_segmentio_backo_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/segmentio/backo-go",
        sum = "h1:68RQccglxZeyURy93ASB/2kc9QudzgIDexJ927N++y4=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_segmentio_kafka_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/segmentio/kafka-go",
        sum = "h1:mcyC3tT5WeyWzrFbd6O374t+hmcu1NKt2Pu1L3QaXmc=",
        version = "v0.4.50",
    )
    go_repository(
        name = "com_github_sergi_go_diff",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sergi/go-diff",
        sum = "h1:xkr+Oxo4BOQKmkn/B9eMK0g5Kg/983T9DqqPHwYqD+8=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_shopify_logrus_bugsnag",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/logrus-bugsnag",
        sum = "h1:UrqY+r/OJnIp5u0s1SbQ8dVfLCZJsnvazdBP5hS4iRs=",
        version = "v0.0.0-20171204204709-577dee27f20d",
    )
    go_repository(
        name = "com_github_shopify_sarama",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/sarama",
        sum = "h1:lqqPUPQZ7zPqYlWpTh+LQ9bhYNu2xJL6k1SJN4WVe2A=",
        version = "v1.38.1",
    )
    go_repository(
        name = "com_github_shopify_toxiproxy_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/Shopify/toxiproxy/v2",
        sum = "h1:i4LPT+qrSlKNtQf5QliVjdP08GyAH8+BUIc9gT0eahc=",
        version = "v2.5.0",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:abSATXmQEYyShuxI4/vyW3tV1MrKAJzCZ/0zLUXYbsQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_shurcool_httpfs",
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/httpfs",
        sum = "h1:aqg5Vm5dwtvL+YgDpBcK1ITf3o96N/K7/wsRXQnUTEs=",
        version = "v0.0.0-20230704072500-f1e31cf0ba5c",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_smartystreets_assertions",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/assertions",
        sum = "h1:zE9ykElWQ6/NYmHa3jpm/yHnI4xSofP+UP6SpjHcSeM=",
        version = "v0.0.0-20180927180507-b2de0cb4f26d",
    )
    go_repository(
        name = "com_github_smartystreets_go_aws_auth",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/go-aws-auth",
        sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
        version = "v0.0.0-20180515143844-0c1422d1fdb9",
    )
    go_repository(
        name = "com_github_smartystreets_goconvey",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/smartystreets/goconvey",
        sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/soheilhy/cmux",
        sum = "h1:jjzc5WVemNEDTLwv9tlmemhC73tI08BNOIGwBOo10Js=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_sourcegraph_annotate",
        build_file_proto_mode = "disable",
        importpath = "github.com/sourcegraph/annotate",
        sum = "h1:yKm7XZV6j9Ev6lojP2XaIshpT4ymkqhMeSghO5Ps00E=",
        version = "v0.0.0-20160123013949-f4cad6c6324d",
    )
    go_repository(
        name = "com_github_sourcegraph_syntaxhighlight",
        build_file_proto_mode = "disable",
        importpath = "github.com/sourcegraph/syntaxhighlight",
        sum = "h1:qpG93cPwA5f7s/ZPBJnGOYQNK/vKsaDaseuKT5Asee8=",
        version = "v0.0.0-20170531221838-bd320f5d308e",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:7c1g84S4BPRrfL5Xrdp6fOJ206sU9y293DDHaoy0bLI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_afero",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/afero",
        sum = "h1:EaGW2JJh15aKOejeuJ+wpFSHnbd7GE6Wvp3TsNhb6LY=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_spf13_cast",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cast",
        sum = "h1:R+kOtfhWQE6TVQzY+4D7wJLBgkdVasCEFxSUBYBYIlA=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/cobra",
        sum = "h1:lJeBwCfmrnXthfAupyUTzJ/J4Nc1RsHC/mSRU2dll/s=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:ue6voC5bR5F8YxI5S67j9i582FU4Qvo2bmqnqMYADFk=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/pflag",
        sum = "h1:9exaQaMOCwffKiiiYk6/BndUBv+iRViNW+4lEMi0PvY=",
        version = "v1.0.9",
    )
    go_repository(
        name = "com_github_spf13_viper",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/spf13/viper",
        sum = "h1:rGGH0XDZhdUOryiDWjmIvUSWpbNqisK8Wk0Vyefw8hc=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_spiffe_go_spiffe_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/spiffe/go-spiffe/v2",
        sum = "h1:N2I01KCUkv1FAjZXJMwh95KK1ZIQLYbPfhaxw8WS0hE=",
        version = "v2.5.0",
    )
    go_repository(
        name = "com_github_stefanberger_go_pkcs11uri",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stefanberger/go-pkcs11uri",
        sum = "h1:lIOOHPEbXzO3vnmx2gok1Tfs31Q8GQqKLc8vVqyQq/I=",
        version = "v0.0.0-20201008174630-78d3cae3a980",
    )
    go_repository(
        name = "com_github_stoewer_go_strcase",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stoewer/go-strcase",
        sum = "h1:g0eASXYtp+yvN9fK8sH94oCIk0fau9uV1/ZdJ0AVEzs=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/objx",
        sum = "h1:xuMeJ0Sdp5ZMRXx/aWO6RZxdr3beISkG5/G/aIRr3pY=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stretchr/testify",
        sum = "h1:7s2iGBzp5EwR7/aIZr8ao5+dra3wiQyKjjFuvgVKu7U=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_github_stvp_tempredis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/stvp/tempredis",
        sum = "h1:QVqDTf3h2WHt08YuiTGPZLls0Wq99X9bWd0Q5ZSBesM=",
        version = "v0.0.0-20181119212430-b82af8480203",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:X1TuBLAMDFbaTAChgCBLu3DU3UPyELpnF2jjJ2cz/S8=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_substrait_io_substrait_go",
        build_file_proto_mode = "disable",
        importpath = "github.com/substrait-io/substrait-go",
        sum = "h1:buDnjsb3qAqTaNbOR7VKmNgXf4lYQxWEcnSGUWBtmN8=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_swaggo_files",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/swaggo/files",
        sum = "h1:kAe4YSu0O0UFn1DowNo2MY5p6xzqtJ/wQ7LZynSvGaY=",
        version = "v0.0.0-20220728132757-551d4a08d97a",
    )
    go_repository(
        name = "com_github_swaggo_gin_swagger",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/swaggo/gin-swagger",
        sum = "h1:8mWmHLolIbrhJJTflsaFoZzRBYVmEE7JZGIq08EiC0Q=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_github_swaggo_swag",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/swaggo/swag",
        sum = "h1:28Pp+8DkQoV+HLzLx8RGJZXNGKbFqnuvSbAAtoxiY04=",
        version = "v1.16.2",
    )
    go_repository(
        name = "com_github_syndtr_gocapability",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/syndtr/gocapability",
        sum = "h1:kdXcSzyDtseVEc4yCz2qF8ZrQvIDBJLl4S1c3GCXmoI=",
        version = "v0.0.0-20200815063812-42c35b437635",
    )
    go_repository(
        name = "com_github_syndtr_goleveldb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/syndtr/goleveldb",
        sum = "h1:vfofYNRScrDdvS342BElfbETmL1Aiz3i2t0zfRj16Hs=",
        version = "v1.0.1-0.20220721030215-126854af5e6d",
    )
    go_repository(
        name = "com_github_tailscale_depaware",
        build_file_proto_mode = "disable",
        importpath = "github.com/tailscale/depaware",
        sum = "h1:34icjjmqJ2HPjrSuJYEkdZ+0ItmGQAQ75cRHIiftIyE=",
        version = "v0.0.0-20210622194025-720c4b409502",
    )
    go_repository(
        name = "com_github_tchap_go_patricia",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tchap/go-patricia",
        sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "com_github_testcontainers_testcontainers_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/testcontainers/testcontainers-go",
        sum = "h1:h0D5GaYG9mhOWr2qHdEKDXpkce/VlvaYOCzTRi6UBi8=",
        version = "v0.14.0",
    )
    go_repository(
        name = "com_github_tidwall_assert",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/assert",
        sum = "h1:aWcKyRBUAdLoVebxo95N7+YZVTFF/ASTr7BN4sLP6XI=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_tidwall_btree",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/btree",
        sum = "h1:LDZfKfQIBHGHWSwckhXI0RPSXzlo+KYdjK7FWSqOzzg=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_tidwall_buntdb",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/buntdb",
        sum = "h1:gdhWO+/YwoB2qZMeAU9JcWWsHSYU3OvcieYgFRS0zwA=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_tidwall_gjson",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/gjson",
        sum = "h1:SyXa+dsSPpUlcwEDuKuEBJEz5vzTvOea+9rjyYodQFg=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_tidwall_grect",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/grect",
        sum = "h1:dA3oIgNgWdSspFzn1kS4S/RDpZFLrIxAZOdJKjYapOg=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_tidwall_lotsa",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/lotsa",
        sum = "h1:dNVBH5MErdaQ/xd9s769R31/n2dXavsQ0Yf4TMEHHw8=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_tidwall_match",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/match",
        sum = "h1:+Ho715JplO36QYgwN9PGYNhgZvoUSc9X2c80KVTi+GA=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_tidwall_pretty",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/pretty",
        sum = "h1:qjsOFOWWQl+N3RsoF5/ssm1pHmJJwhjlSbZ51I6wMl4=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_tidwall_rtred",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/rtred",
        sum = "h1:exmoQtOLvDoO8ud++6LwVsAMTu0KPzLTUrMln8u1yu8=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_tidwall_sjson",
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/sjson",
        sum = "h1:kLy8mja+1c9jlljvWTlSazM7cKDRfJuR/bOJhcY5NcY=",
        version = "v1.2.5",
    )
    go_repository(
        name = "com_github_tidwall_tinyqueue",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tidwall/tinyqueue",
        sum = "h1:SpNEvEggbpyN5DIReaJ2/1ndroY8iyEGxPYxoSaymYE=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_tinylib_msgp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tinylib/msgp",
        sum = "h1:FCXC1xanKO4I8plpHGH2P7koL/RzZs12l/+r7vakfm0=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        sum = "h1:uruHq4dN7GR16kFc5fp3d1RIYzJW5onx8Ybykw2YQFA=",
        version = "v0.0.0-20201229170055-e5319fda7802",
    )
    go_repository(
        name = "com_github_tmthrgd_go_hex",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tmthrgd/go-hex",
        sum = "h1:9lRDQMhESg+zvGYmW5DyG0UqvY96Bu5QYsTLvCHdrgo=",
        version = "v0.0.0-20190904060850-447a3041c3bc",
    )
    go_repository(
        name = "com_github_tv42_httpunix",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/tv42/httpunix",
        sum = "h1:u6SKchux2yDvFQnDHS3lPnIRmfVJ5Sxy3ao2SIdysLQ=",
        version = "v0.0.0-20191220191345-2ba4b9c3382c",
    )
    go_repository(
        name = "com_github_twitchtv_twirp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/twitchtv/twirp",
        sum = "h1:+F4TdErPgSUbMZMwp13Q/KgDVuI7HJXP61mNV3/7iuU=",
        version = "v8.1.3+incompatible",
    )
    go_repository(
        name = "com_github_twitchyliquid64_golang_asm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/twitchyliquid64/golang-asm",
        sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_tylerbrock_colorjson",
        build_file_proto_mode = "disable",
        importpath = "github.com/TylerBrock/colorjson",
        sum = "h1:ZBbLwSJqkHBuFDA6DUhhse0IGJ7T5bemHyNILUjvOq4=",
        version = "v0.0.0-20200706003622-8a50f05110d2",
    )
    go_repository(
        name = "com_github_ugorji_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ugorji/go",
        sum = "h1:qYhyWUUd6WbiM+C6JZAUkIJt/1WrjzNHY9+KCIjVqTo=",
        version = "v1.2.7",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:Qd2W2sQawAfG8XSvzwhBeoGq71zXOC/Q1E9y/wUcsUA=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_urfave_cli",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/urfave/cli",
        sum = "h1:gsqYFH8bb9ekPA12kRo0hfjngWQjkJPlN9R0N78BoUo=",
        version = "v1.22.2",
    )
    go_repository(
        name = "com_github_urfave_cli_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/urfave/cli/v2",
        sum = "h1:VAzn5oq403l5pHjc4OhD54+XGO9cdKVL/7lDjF+iKUs=",
        version = "v2.25.7",
    )
    go_repository(
        name = "com_github_urfave_negroni",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/urfave/negroni",
        sum = "h1:kIimOitoypq34K7TG7DUaJ9kq/N4Ofuwi1sjz0KipXc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_bytebufferpool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/bytebufferpool",
        sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_fasthttp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/fasthttp",
        sum = "h1:H7fweIlBm0rXLs2q0XbalvJ6r0CUPFWK3/bB4N13e9M=",
        version = "v1.50.0",
    )
    go_repository(
        name = "com_github_valyala_fasttemplate",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/fasttemplate",
        sum = "h1:lxLXG0uE3Qnshl9QyaK6XJxMXlQZELvChBOCmQD0Loo=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_valyala_tcplisten",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/valyala/tcplisten",
        sum = "h1:rBHj/Xf+E1tRGZyWIWwJDiRY0zc1Js+CV5DqwacVSA8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_vektah_gqlparser_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vektah/gqlparser/v2",
        sum = "h1:pm6WOnGdzFOCfcQo9L3+xzW51mKrlwTEg4Wr7AH1JW4=",
        version = "v2.5.8",
    )
    go_repository(
        name = "com_github_vishvananda_netlink",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vishvananda/netlink",
        sum = "h1:+UB2BJA852UkGH42H+Oee69djmxS3ANzl2b/JtT1YiA=",
        version = "v1.1.1-0.20210330154013-f5de75959ad5",
    )
    go_repository(
        name = "com_github_vishvananda_netns",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vishvananda/netns",
        sum = "h1:p4VB7kIXpOQvVn1ZaTIVp+3vuYAXFe3OJEvjbUYJLaA=",
        version = "v0.0.0-20210104183010-2eb08e3e575f",
    )
    go_repository(
        name = "com_github_vmihailenco_bufpool",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/bufpool",
        sum = "h1:gOq2WmBrq0i2yW5QJ16ykccQ4wH9UyEsgLm6czKAd94=",
        version = "v0.1.11",
    )
    go_repository(
        name = "com_github_vmihailenco_msgpack_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/msgpack/v5",
        sum = "h1:5gO0H1iULLWGhs2H5tbAHIZTV8/cYafcFOr9znI5mJU=",
        version = "v5.3.5",
    )
    go_repository(
        name = "com_github_vmihailenco_tagparser",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/tagparser",
        sum = "h1:gnjoVuB/kljJ5wICEEOpx98oXMWPLj22G67Vbd1qPqc=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_vmihailenco_tagparser_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/vmihailenco/tagparser/v2",
        sum = "h1:y09buUbR+b5aycVFQs/g70pqKVZNBmxwAhO7/IwNM9g=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_vultr_govultr_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/vultr/govultr/v2",
        sum = "h1:gej/rwr91Puc/tgh+j33p/BLR16UrIPnSr+AIwYWZQs=",
        version = "v2.17.2",
    )
    go_repository(
        name = "com_github_willf_bitset",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/willf/bitset",
        sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
        version = "v1.1.11",
    )
    go_repository(
        name = "com_github_x448_float16",
        build_file_proto_mode = "disable",
        importpath = "github.com/x448/float16",
        sum = "h1:qLwI1I70+NjRFUR3zs1JPUCgaCXSh3SW62uAKT1mSBM=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_github_xdg_go_pbkdf2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg-go/pbkdf2",
        sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xdg_go_scram",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg-go/scram",
        sum = "h1:FHX5I5B4i4hKRVRBCFRxq1iQRej7WO3hhBuJf+UUySY=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_xdg_go_stringprep",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg-go/stringprep",
        sum = "h1:XLI/Ng3O1Atzq0oBs3TWm+5ZVgkq2aqdlvP9JtoZ6c8=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_xdg_scram",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg/scram",
        sum = "h1:u40Z8hqBAAQyv+vATcGgV0YCnDjqSL7/q/JyPhhJSPk=",
        version = "v0.0.0-20180814205039-7eeb5667e42c",
    )
    go_repository(
        name = "com_github_xdg_stringprep",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xdg/stringprep",
        sum = "h1:d9X0esnoa3dFsV0FG35rAT0RIhYFlPq7MiP+DW89La0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:zGWFAtiMcyryUHoUjUJX0/lt1H2+i2Ka2n+D3DImSNo=",
        version = "v0.0.0-20190905194746-02993c407bfb",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xiang90/probing",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        version = "v0.0.0-20190116061207-43a291ad63a2",
    )
    go_repository(
        name = "com_github_xordataexchange_crypt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xordataexchange/crypt",
        sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
        version = "v0.0.3-0.20170626215501-b2862e3d0a77",
    )
    go_repository(
        name = "com_github_xrash_smetrics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/xrash/smetrics",
        sum = "h1:bAn7/zixMGCfxrRTfdpNzjtPYqr8smhKouy9mxVdGPU=",
        version = "v0.0.0-20201216005158-039620a65673",
    )
    go_repository(
        name = "com_github_xtgo_uuid",
        build_file_proto_mode = "disable",
        importpath = "github.com/xtgo/uuid",
        sum = "h1:3lbZUMbMiGUW/LMkfsEABsc5zNT9+b1CvsJx47JzJ8g=",
        version = "v0.0.0-20140804021211-a0b114877d4c",
    )
    go_repository(
        name = "com_github_yalp_jsonpath",
        build_file_proto_mode = "disable",
        importpath = "github.com/yalp/jsonpath",
        sum = "h1:6fRhSjgLCkTD3JnJxvaJ4Sj+TYblw757bqYgZaOq5ZY=",
        version = "v0.0.0-20180802001716-5cc68e5049a0",
    )
    go_repository(
        name = "com_github_youmark_pkcs8",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/youmark/pkcs8",
        sum = "h1:splanxYIlg+5LfHAM6xpdFEAYOk8iySO56hMFq6uLyA=",
        version = "v0.0.0-20181117223130-1be2e3e5546d",
    )
    go_repository(
        name = "com_github_yudai_gojsondiff",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/gojsondiff",
        sum = "h1:27cbfqXLVEJ1o8I6v3y9lg8Ydm53EKqHXAOMxEGlCOA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_yudai_golcs",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/golcs",
        sum = "h1:BHyfKlQyqbsFN5p3IfnEUduWvb9is428/nNb5L3U01M=",
        version = "v0.0.0-20170316035057-ecda9a501e82",
    )
    go_repository(
        name = "com_github_yudai_pp",
        build_file_proto_mode = "disable",
        importpath = "github.com/yudai/pp",
        sum = "h1:Q4//iY4pNF6yPLZIigmvcl7k/bPgrcTPIFIcmawg5bI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_github_yuin_gopher_lua",
        build_file_proto_mode = "disable",
        importpath = "github.com/yuin/gopher-lua",
        sum = "h1:kYKnWBjvbNP4XLT3+bPEwAXJx262OhaHDWDVOPjL46M=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_yvasiyarov_go_metrics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/yvasiyarov/go-metrics",
        sum = "h1:+lm10QQTNSBd8DVTNGHx7o/IKu9HYDvLMffDhbyLccI=",
        version = "v0.0.0-20140926110328-57bccd1ccd43",
    )
    go_repository(
        name = "com_github_yvasiyarov_gorelic",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/yvasiyarov/gorelic",
        sum = "h1:hlE8//ciYMztlGpl/VA+Zm1AcTPHYkHJPbHqE6WJUXE=",
        version = "v0.0.0-20141212073537-a9bba5b9ab50",
    )
    go_repository(
        name = "com_github_yvasiyarov_newrelic_platform_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/yvasiyarov/newrelic_platform_go",
        sum = "h1:ERexzlUfuTvpE74urLSbIQW0Z/6hF9t8U4NsJLaioAY=",
        version = "v0.0.0-20140908184405-b21fdbd4370f",
    )
    go_repository(
        name = "com_github_zeebo_assert",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/zeebo/assert",
        sum = "h1:g7C04CbJuIDKNPFHmsk4hwZDO5O+kntRxzaUoNXj+IQ=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_zeebo_errs",
        build_file_proto_mode = "disable",
        importpath = "github.com/zeebo/errs",
        sum = "h1:XNdoD/RRMKP7HD0UhJnIzUy74ISdGGxURlYG8HSWSfM=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_zeebo_xxh3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/zeebo/xxh3",
        sum = "h1:xZmwmqxHZA8AI603jOQ0tMqmBr9lPeFwGg6d+xy9DC0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_zenazn_goji",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "github.com/zenazn/goji",
        sum = "h1:4lbD8Mx2h7IvloP7r2C0D6ltZP6Ufip8Hn0wmSK5LR8=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_google_cloud_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go",
        sum = "h1:S3kTQSydxmu1JfLRLpKtxRPA7rSrYPRPEUmL/PavVUw=",
        version = "v0.121.1",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:UkmDPCKvj24bkGVrvgJPcgSDkmIPw/bAmOiDb9avOiE=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:2LnncRqfYB8NEdh9+FeYxAt9POTW/0zVboktnRlO11w=",
        version = "v1.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:niSJYc6ldWWVM9faXPo1Et1MVSQoLvVGriD7fwbJdtE=",
        version = "v1.89.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:W2ft49J/LeEj9A07Jsd5Q2kAzajK0j0IffOyyzbxw04=",
        version = "v0.28.1",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:do+u3rjDYuTxD2ypRfv4uwTMoy/VHFLclvaYcb5Mv6I=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:ijEJSni5xROOn1YyiHgqcW0B0TWr0di9VgIi2gvyNjY=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:TgdjAoGoRY81DEc2LYsYvi/OqCFImMzAk/TVKiSRsQw=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:JJyY8icMmQeWfQ+d36IhkGvd3Guzvw0UAkvxT0wmUx8=",
        version = "v1.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:iJrZ6AleZr4l+q0/fWVANFOhs90KiSB1Ccait5OYyNg=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:A20kj2S2HO9vlyBVyVFHPxArjxkXvLP5LjcdE7NhaPc=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:i55wWC/EwVdHMyJgRfbLp/L6ez4nQuOpZwSxkuqN9ek=",
        version = "v1.21.1",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:ip/shfJYx6lrHBWYADjrrrubcm7uZzy50TTF5tPG7ek=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_auth",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/auth",
        sum = "h1:XrXauHMd30LhQYVRHLGvJiYeczweKQXZxsTbV9TiguU=",
        version = "v0.16.1",
    )
    go_repository(
        name = "com_google_cloud_go_auth_oauth2adapt",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/auth/oauth2adapt",
        sum = "h1:keo8NaayQZ6wimpNSmW5OPc283g65QNIiLpZnkHRbnc=",
        version = "v0.2.8",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:ZLj48Ur2Qcso4M3bgOtjsOmeV5Ee92N14wuOc8OW+L0=",
        version = "v1.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:9bdGlpY1LgLONQjFsDwrkjLzdPTlROpfU+GhA97YpOk=",
        version = "v1.3.6",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:gWQdvdPplptpvrkqF6ibtxZkOsYKLTFbxYawHa/TvCg=",
        version = "v1.12.2",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:4FcR+4QmcNGkhVij6TrYS4AQVNLBo7PBXKxNrKzpclQ=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:rZvHnjSUs5sHK3F9awiuFk2PeOaB8suqNuim21GbaTc=",
        version = "v1.69.0",
    )
    go_repository(
        name = "com_google_cloud_go_bigtable",
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/bigtable",
        sum = "h1:Q+x7y04lQ0B+WXp03wc1/FLhFt4CwcQdkwWT0M4Jp3w=",
        version = "v1.37.0",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:pqM5/c9UGydB9H90IPCxSvfCNLUPazAOSMsZkz5q5P4=",
        version = "v1.20.4",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:T0zYEroXT+y0O/x/yZd5SwQdFv4UbUINjvJyJKzDm0Q=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:+ZPglfDurCcsv4azizDFpBucD1IkRjWjbnU7zceyjfY=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:UI+ZsRkS15hi9DRF+WAvTVLVuSeZiRmvCU8cjkjOwUU=",
        version = "v1.19.5",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:4LlrIFa3IFLgD1mGEXmUE4cm9fYoU71OLwTvjM7Dg3c=",
        version = "v1.22.2",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:IWJbQBEECTaNanDRN1XdR7FU53MJ1nylTl3s9T3MuyI=",
        version = "v1.8.7",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:Fwan19UiNoFD+3KY0MnNHE5DyixOxNzS1mZ4ChOdpy0=",
        version = "v1.13.6",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:MilCLYQW2m7Dku8hRIIKo4r0oKastlD74sSu16riYKs=",
        version = "v1.38.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:PBWF+iiAerVNe8UCHxdOt6eHLVc3ydFeOCw78U8ytSU=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:lenyU3uzHwKDveCwmpfNxHYvLS3uEBWdn+O7+rSxy+Q=",
        version = "v1.17.3",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/container",
        sum = "h1:A6J92FJPfxTvyX7MHF+w4t2W9WCqvHOi9UB5SAeSy3w=",
        version = "v1.43.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:1SoHlNqL3XrhqcoozB+3eoHif2sRUFtp/JeASQTtGKo=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:eFgygb3DTufTWWUB8ARk+dSuXz+aefNJXTlkWlQcWwE=",
        version = "v1.26.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:AdhB4cAkMOC9NtrHJxpKOVvO/VqBLaIyk0tEEhbGjYM=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:0eCPTPUC/RZ863aVfXTJLkg0tEpdpn62VD6ywSmmzxM=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:GZ6J+CR8CEeWAj8luRCtr8GvImSQRkArIIqGiZOnzBA=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:VOZ5U+78ttnhNCEID7qdeogqZQzK5N+LPHIQ9Q3YDsc=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:Xr0Toh6wyBlmL3H4EPu1YKwxUtkDSzzq+IP0iLc88kk=",
        version = "v1.25.3",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:KhC8wdLILpAs17yeTG6Miwg1v0nOP/OXD+9QNg3w6AQ=",
        version = "v2.11.2",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:qTRAG/E3T63Xj1orefRlwupfwH9c9ERUAnWSRGp75so=",
        version = "v0.9.7",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:NNpXoyEqIJmZFc0ACcwBEaXnmscUpcG4NkKnbCePmiM=",
        version = "v1.20.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:j+y0lUKm9pbDjJn0YcWxPI/hXNGUQ80GE6yrFuJC/JA=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:C0VqBhFyQFp6+xgPHZAD7LeRA4XGy5YLzGmPQ2NhlLk=",
        version = "v1.27.2",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:bXpoqPRf37KKxB79PKr20B/TAU/Z5iA0FnB6C5N2jrA=",
        version = "v1.68.2",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:3xWRKylXxhysaQaV+DLev1YcIywFUCc7yJEE6R7ZGDQ=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:7fla8GcarupO15eatRTUveXCob6DOSW1Wa+1i63CM3Q=",
        version = "v1.37.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:TI+Aavwc31KD8huOquJz0ISchCq1zSEWc9M+JcPJyxc=",
        version = "v0.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:9tfGCicvrki927T+hGMB0yYmwIbRuZY6JR1/awrKiZ0=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:isaoPwWX8kbAOea4qahcmttoS79+gQhvKsfg5L5AgH8=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:ysHZ4gr4plW1CL1Ur/AucUUfh20hDjSFbfjxSK0q/sk=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:bZW7ZMM+XXNErg6rOZcgxUzAgz4vpReRDP3ZiGf7/sI=",
        version = "v1.15.5",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:LjoAyp9TvVNBns3sUUzPaNsQiGpR2BReGmTS3bUCuBE=",
        version = "v1.10.2",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:cuydCaLS7Vl2SatAeivXyhbhDEIR8BDmtn4egDhIn2s=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:vJgWlvxtJG6p/JrbXAkz83DbgwOyFhZZI1Y32vUddjY=",
        version = "v1.19.6",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:eBqOt61yEChvj7I/GDPBbdCCRdUPudD1qrQYfYWV3Ok=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:67/rnPmF/I1Wmf7jWyKH+z4OWjU8ZUI0Vmzxvmzf3KY=",
        version = "v0.12.4",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:9iogrmNNa+drDPf/zkLH/6KGgUf7FuuyokmithoGwMQ=",
        version = "v0.15.6",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:334aZmOzIt3LVBpguCof8IHaLaftcZlx+L0TGBukYkY=",
        version = "v1.5.3",
    )
    go_repository(
        name = "com_google_cloud_go_grafeas",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/grafeas",
        sum = "h1:lBjwKmhpiqOAFaE0xdqF8CqO74a99s8tUT5mCkBBxPs=",
        version = "v0.3.15",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:sk0SxpCGIA7tIO//XdiiG29f2vrF6Pq/dsxxyBGiRBY=",
        version = "v1.7.7",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:qgFRAGEmd8z6dJ/qyEchAuL9jpswyODjA2lS+w234g8=",
        version = "v1.5.2",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:VIioCrYsyWiRGx7Y8RDNylpI6d4t1Qx5ZgSLUVmWWPo=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:uKGuaWozDcjg3wyf54Gd7tCH2YK8BFeH9qo1xBNiPKE=",
        version = "v1.5.6",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:A3AhugnIViAZkC3/lHAQDaXBIk2ZOPBZS0XQCyZsjjc=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:dBRIj7+GDeeEvatJeTB19oYZNV0aj6wEqSIT/7gLqtk=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/language",
        sum = "h1:BVJ/POtlnJ55LElvnQY19UOxpMVtHoHHkFJW2uHJsVU=",
        version = "v1.14.5",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:Vu7XF4s5KJ8+mSLIL4eaQM6JTyWXvSB54oqC+CUZH20=",
        version = "v0.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:7j0HgAp0B94o1YRDqiqm26w4q1rDMH7XNRU34lJXHYc=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:IGtfDWHhQCgCjwQjV9iiLnUta9LBCo8R9QmAFsS/PrE=",
        version = "v0.6.7",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:zrZVWXZJlmHnfpyCrTQIbDBGUBHrcOOvrsjMjoXRxrk=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:El61AfMxC1sU/RU8Wzs9dkZEgltyunKM86aKF9aDlaE=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:SDGatA73TgZ8iCvILVXpk/1qhTK5DJyufUDEWgbmbV8=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:33IVqQEmFiITsBXwGHeTkUhWz0kLNKr90nV3e22uLPs=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:dLm59AHHZCorveCylj7c2iWhkQsmMIeWTsV+tG/BXtY=",
        version = "v1.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:5OTsoJ1dXYIiMiuL+sYscLc9BumrL3CarVLL7dd7lHM=",
        version = "v1.24.2",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:RQcG1rZNCNV5Dn3tnINs4TYswDXk2hKH+85eh+JvoWU=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:ecukgArkYCVcK5w2h7WDDd+nHgmBAp9Bst7ClmVKz5A=",
        version = "v1.19.1",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:6b6fcCG9BFNcmtNO+VuPE04vkZb5TKNX9+7ZhYMgstE=",
        version = "v0.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:nCfZwVihArMPP2atRoxRrXOXJ/aC9rAgpBQGCc2zpYw=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:jDvIuSxDsXI2P7l2sYXm6CoX1YBIIT6Khm5m0hq0/KQ=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:PnlZ/O4R/eiounpxUkhI9ZXRMWbG7vFqxc6L6sR+31k=",
        version = "v1.11.9",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:uQziDu3UKYk9ZwUgneZAW5aWxZFKgOXXsuVKFKh0z7Y=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:4uJrA1obzMBp1I+DF15y/MvsXKIODevuANpq3QhvX30=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:BDKVcxo1OO4ZT+PbuFchZjnbrlUGfChilt6+pITY1VI=",
        version = "v1.14.6",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:yl572bBQbPjflX250SOflN6gwO2uYoddN2uRp36fDTo=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:Z8+tO2z21MY1arBBuJjwrOjbw8fbZb13AZTHXdzkl2U=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:R951ikhxIanXEijBCu0xnoUAOteS5m/Xplek0YvsNTE=",
        version = "v0.10.7",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:5054IkbslnrMCgA2MAEPcsN3Ky+AyMpEZcii/DoySPo=",
        version = "v1.49.0",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:jLQozsEVr+c6tOU13vDugtnaBSUy/PD5zK6mhm+uF1Y=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:P4QMryKcWdi4LIe1Sx0b2ZOAQv5gVfdzPt2peXcN32Y=",
        version = "v2.20.4",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:slN7h23vswGccW8x3f+xUXCu9Yo18/GNkazH93LJbFk=",
        version = "v0.9.6",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:cIsyRKGNw4LpCfY5c8CCQadhlp54jP4fHtP+d5Sy2xE=",
        version = "v1.13.5",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:JlHLceAOILEmbn+NIS7l+vmUKkFuobLToCWTxL7NGcQ=",
        version = "v1.18.2",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:LIa8kKE8HF71zm976oHMqpWFiaDHVw/H1YMO71lrGmo=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:13HOFU7v4cEvIHXSAQbinF4wp2Baybbq7q9FMctg1Ek=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:8jgWgtAg1mk91WmaoWRTlL9CcvazPwqZ3YT9n6Gva9U=",
        version = "v1.21.0",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/run",
        sum = "h1:CDhz0PPzI/cVpmNFyHe3Yp21jNpiAqtkfRxuoLi+JU0=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:zkMEJ0UbEJ3O7NwEUlKLIp6eXYv1L7wHjbxyxznajKM=",
        version = "v1.11.7",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:VkscIRzj7GcmZyO4z9y1EH7Xf81PcoiAo7MtlD+0O80=",
        version = "v1.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/security",
        sum = "h1:6hqzvuwC8za9jyCTxygmEHnp4vZ8hfhwKVArxSCAVCo=",
        version = "v1.18.5",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:hLA58IBYmWrNiXDIONvuCUQ4sHLVPy8JvDo2j1wSYCw=",
        version = "v1.36.2",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:pl/KUNvFzlXpxgnPgzQjyTQQcv5WsQ97zCHaPrLQlYA=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:jLWyztGlNWBx55QXBM4HbWvfv7aiRjPzRKTUkZA8dXk=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:w9uO8RqEoBooBLX4nqV1RtgudyU2ZX780KTLRgeVg60=",
        version = "v1.82.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:+OktATNlQc+4WH78OrQadIP4CzXb9mBucdDGCO1NrlI=",
        version = "v1.27.1",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:gg0ERZwL17pJ+Cz3cD2qS60w1WMDnwcm5YPAIQBHUAw=",
        version = "v1.53.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:uqKX3OgcYzR1W1YI943ZZ45id0RqA2eXXoCBSPstlbw=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:wDP+++O/P1cTJBMkYlSY46k0a6atSoyO+UkBGuU9+Ao=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:oWWFQp0yFl4EJOr3opDkKH9304wUsZjgPjrTDS6S1a8=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:S4Ptq+yFIPNLEzQ/OQwiIYDNzk5I2vYmhf0SmFQOmWo=",
        version = "v1.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:2O2zjPzqPYAHrn3OKl029qlqG6W8ZdYaOWRyr8NgMT4=",
        version = "v1.11.6",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:QPMNi4WCtHwc2PPfxbyUMwdN/0+cyCGLaKi2tig41J8=",
        version = "v1.12.5",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/video",
        sum = "h1:KTB2BEXjGm2K/JcKxQXEgx3nSoMTByepnPZa4kln064=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:heq7jEO39sH5TycBh8TGFJ827XCxK0tIWatmBY/n0jI=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:UJZ0H6UlOaYKgCn6lWG2iMAOJIsJZLnseEfzBR8yIqQ=",
        version = "v2.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:68hOQDhs1DOITrCrhritrwr8xy6s8QMdwDyMzMiFleU=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:OsGd1SB91y9fDuzdzFngMv4UcT4cqmRxjsCsS4Xmcu8=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:RYtUB9rQEijX9Tc6lQcGst58ZOzPgaYTkz6+2pyPQTM=",
        version = "v1.8.6",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:yZKNB7zRxOMriLrhP5WDE+BjxXVl0wJHHZSdaYzbdVU=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:cIPKJKZA3l7D8DfL4nxce8HGOWXBw3WAUBF0ymOW9GQ=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:phBz5TOAES0YGogxZ6Q7ISSudaf618lRhE3euzBpE9U=",
        version = "v1.14.2",
    )
    go_repository(
        name = "com_lukechampine_uint128",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "lukechampine.com/uint128",
        sum = "h1:cDdUVfRwDUDovz610ABgFD17nXD4/uDgVHl2sC3+sbo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_nullprogram_x_optparse",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "nullprogram.com/x/optparse",
        sum = "h1:xGFgVi5ZaWOnYdac2foDT3vg0ZZC9ErXFV57mr4OHrI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_shuralyov_dmitri_gpu_mtl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "dmitri.shuralyov.com/gpu/mtl",
        sum = "h1:VpgP7xuJadIUuKccphEpTJnWhS2jkQyMt6Y7pJCD7fY=",
        version = "v0.0.0-20190408044501-666a987793e9",
    )
    go_repository(
        name = "dev_cel_expr",
        build_file_proto_mode = "disable",
        importpath = "cel.dev/expr",
        sum = "h1:56OvJKSH3hDGL0ml5uSxZmz3/3Pq4tJ+fb1unVLAFcY=",
        version = "v0.24.0",
    )
    go_repository(
        name = "dev_dny_code_ssrf",
        build_file_proto_mode = "disable",
        importpath = "code.dny.dev/ssrf",
        sum = "h1:wCBP990rQQ1CYfRpW+YK1+8xhwUjv189AQ3WMo1jQaI=",
        version = "v0.2.0",
    )
    go_repository(
        name = "ht_sr_git_sbinet_gg",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "git.sr.ht/~sbinet/gg",
        sum = "h1:LNhjNn8DerC8f9DHLz6lS0YYul/b602DUxDgGkd/Aik=",
        version = "v0.3.1",
    )
    go_repository(
        name = "im_mellium_sasl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "mellium.im/sasl",
        sum = "h1:wE0LW6g7U83vhvxjC1IY8DnXM+EU095yeo8XClvCdfo=",
        version = "v0.3.1",
    )
    go_repository(
        name = "in_gopkg_airbrake_gobrake_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/airbrake/gobrake.v2",
        sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
        version = "v2.0.9",
    )
    go_repository(
        name = "in_gopkg_alecthomas_kingpin_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/alecthomas/kingpin.v2",
        sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
        version = "v2.2.6",
    )
    go_repository(
        name = "in_gopkg_avro_v0",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/avro.v0",
        sum = "h1:PGIdqvwfpMUyUP+QAlAnKTSWQ671SmYjoou2/5j7HXk=",
        version = "v0.0.0-20171217001914-a730b5802183",
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
        name = "in_gopkg_cheggaaa_pb_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/cheggaaa/pb.v1",
        sum = "h1:Ev7yu1/f6+d+b3pi5vPdRPc6nNtP1umSfcWiEfRqv6I=",
        version = "v1.0.25",
    )
    go_repository(
        name = "in_gopkg_datadog_dd_trace_go_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/DataDog/dd-trace-go.v1",
        sum = "h1:XKO91GwTjpIRhd56Xif/BZ2YgHkQufVTOvtkbRYSPi8=",
        version = "v1.61.0",
    )
    go_repository(
        name = "in_gopkg_errgo_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/errgo.v1",
        sum = "h1:n+7XfCyygBFb8sEjg6692xjC6Us50TFRO54+xYUEwjE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_evanphx_json_patch_v4",
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/evanphx/json-patch.v4",
        sum = "h1:n6jtcsulIzXPJaxegRbvFNNrZDjbij7ny3gmSPG+6V4=",
        version = "v4.12.0",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
        sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
        version = "v2.1.2",
    )
    go_repository(
        name = "in_gopkg_httprequest_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/httprequest.v1",
        sum = "h1:pEPLMdF/gjWHnKxLpuCYaHFjc8vAB2wrYjXrqDVC16E=",
        version = "v1.2.1",
    )
    go_repository(
        name = "in_gopkg_inconshreveable_log15_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/inconshreveable/log15.v2",
        sum = "h1:RlWgLqCMMIYYEVcAR5MDsuHlVkaIPDAF+5Dehzg8L5A=",
        version = "v2.0.0-20180818164646-67afb5ed74ec",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:Dgnx+6+nfE+IfzjUEISNeydPJh9AXNNsWbGP9KzCsOA=",
        version = "v1.67.0",
    )
    go_repository(
        name = "in_gopkg_jinzhu_gorm_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/jinzhu/gorm.v1",
        sum = "h1:sTqyEcgrxG68jdeUXA9syQHNdeRhhfaYZ+vcL3x730I=",
        version = "v1.9.2",
    )
    go_repository(
        name = "in_gopkg_mgo_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/mgo.v2",
        sum = "h1:VpOs+IwYnYBaFnrNAeB8UUWtL3vEUnzSCL1nVjPhqrw=",
        version = "v2.0.0-20190816093944-a6b53ec6cb22",
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        sum = "h1:1Lc07Kr7qY4U2YPouBjpCLxpiyxIVoxqXgkXLknAOE8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "in_gopkg_natefinch_npipe_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/natefinch/npipe.v2",
        sum = "h1:+JknDZhAj8YMt7GC73Ei8pv4MzjDUNPHgQWJdtMAaDU=",
        version = "v2.0.0-20160621034901-c1b8fa8bdcce",
    )
    go_repository(
        name = "in_gopkg_olivere_elastic_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/olivere/elastic.v3",
        sum = "h1:u3B8p1VlHF3yNLVOlhIWFT3F1ICcHfM5V6FFJe6pPSo=",
        version = "v3.0.75",
    )
    go_repository(
        name = "in_gopkg_olivere_elastic_v5",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/olivere/elastic.v5",
        sum = "h1:acF/tRSg5geZpE3rqLglkS79CQMIMzOpWZE7hRXIkjs=",
        version = "v5.0.84",
    )
    go_repository(
        name = "in_gopkg_resty_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/resty.v1",
        sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
        version = "v1.12.0",
    )
    go_repository(
        name = "in_gopkg_retry_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/retry.v1",
        sum = "h1:a9CArYczAVv6Qs6VGoLMio99GEs7kY9UzSF9+LD+iGs=",
        version = "v1.0.3",
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/square/go-jose.v2",
        sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
        version = "v2.6.0",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
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
        name = "io_etcd_go_bbolt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/bbolt",
        sum = "h1:/ecaJf0sk1l4l6V4awd65v2C3ILy7MSj+s/x1ADCIMU=",
        version = "v1.3.6",
    )
    go_repository(
        name = "io_etcd_go_etcd",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd",
        sum = "h1:1JFLBqwIgdyHN1ZtgjTBwO+blA6gVOmZurpiMEsETKo=",
        version = "v0.5.0-alpha.5.0.20200910180754-dd1b699fc489",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:4wSsluwyTbGGmyjJktOf3wFQoTBIURXHnq9n/G/JQHs=",
        version = "v3.5.9",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:oidDC4+YEuSIQbsR94rY9gur91UPL6DnxDCIYd2IGsE=",
        version = "v3.5.9",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:AELPkjNR3/igjbO7CjyF1fPuVPjrblliiKj+Y6xSGOU=",
        version = "v2.305.7",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:r5xghnU7CwbUxD/fbUtRyJGaYNfDun8sp/gTr1hew6E=",
        version = "v3.5.9",
    )
    go_repository(
        name = "io_etcd_go_etcd_pkg_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/pkg/v3",
        sum = "h1:ntrg6vvKRW26JRmHTE0iNlDgYK6JX3hg/4cD62X0ixk=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_raft_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/raft/v3",
        sum = "h1:kw2TmO3yFTgE+F0mdKkG7xMxkit2duBDa2Hu6D/HMlw=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_server_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/server/v3",
        sum = "h1:jk8D/lwGEDlQU9kZXUFMSANkE22Sg5+mW27ip8xcF9E=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_filippo_edwards25519",
        build_file_proto_mode = "disable",
        importpath = "filippo.io/edwards25519",
        sum = "h1:FNf4tywRC1HmFuKW5xopWpigGjJKiJSV0Cqo0cJWDaA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "io_gorm_driver_mysql",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gorm.io/driver/mysql",
        sum = "h1:omJoilUzyrAp0xNoio88lGJCroGdIOen9hq2A/+3ifw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "io_gorm_driver_postgres",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gorm.io/driver/postgres",
        sum = "h1:1FPESNXqIKG5JmraaH2bfCVlMQ7paLoCreFxDtqzwdc=",
        version = "v1.4.6",
    )
    go_repository(
        name = "io_gorm_driver_sqlserver",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gorm.io/driver/sqlserver",
        sum = "h1:nMtEeKqv2R/vv9FoHUFWfXfP6SskAgRar0TPlZV1stk=",
        version = "v1.4.2",
    )
    go_repository(
        name = "io_gorm_gorm",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gorm.io/gorm",
        sum = "h1:zi4rHZj1anhZS2EuEODMhDisGy+Daq9jtPrNGgbQYD8=",
        version = "v1.25.3",
    )
    go_repository(
        name = "io_k8s_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/api",
        sum = "h1:Xe1hX/fPW3PXYYv8BlozYqw63ytA92snr96zMW9gWTU=",
        version = "v0.31.1",
    )
    go_repository(
        name = "io_k8s_apimachinery",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apimachinery",
        sum = "h1:mhcUBbj7KUjaVhyXILglcVjuS4nYXiwC+KKFBgIVy7U=",
        version = "v0.31.1",
    )
    go_repository(
        name = "io_k8s_apiserver",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apiserver",
        sum = "h1:71krQxCUz218ecb+nPhfDsNB6QgP1/4EMvi1a2uYBlg=",
        version = "v0.22.5",
    )
    go_repository(
        name = "io_k8s_client_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/client-go",
        sum = "h1:f0ugtWSbWpxHR7sjVpQwuvw9a3ZKLXX0u0itkFXufb0=",
        version = "v0.31.1",
    )
    go_repository(
        name = "io_k8s_code_generator",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/code-generator",
        sum = "h1:kM/68Y26Z/u//TFc1ggVVcg62te8A2yQh57jBfD0FWQ=",
        version = "v0.19.7",
    )
    go_repository(
        name = "io_k8s_component_base",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/component-base",
        sum = "h1:U0eHqZm7mAFE42hFwYhY6ze/MmVaW00JpMrzVsQmzYE=",
        version = "v0.22.5",
    )
    go_repository(
        name = "io_k8s_cri_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/cri-api",
        sum = "h1:0DHL/hpTf4Fp+QkUXFefWcp1fhjXr9OlNdY9X99c+O8=",
        version = "v0.23.1",
    )
    go_repository(
        name = "io_k8s_gengo",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/gengo",
        sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
        version = "v0.0.0-20210813121822-485abfe95c7c",
    )
    go_repository(
        name = "io_k8s_gengo_v2",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/gengo/v2",
        sum = "h1:NGrVE502P0s0/1hudf8zjgwki1X/TByhmAoILTarmzo=",
        version = "v2.0.0-20240228010128-51d4e06bde70",
    )
    go_repository(
        name = "io_k8s_klog",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/klog",
        sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_k8s_klog_v2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/klog/v2",
        sum = "h1:n9Xl7H1Xvksem4KFG4PYbdQCQxqc/tTUyrgXaOhHSzk=",
        version = "v2.130.1",
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kube-openapi",
        sum = "h1:BZqlfIlq5YbRMFko6/PM7FjZpUb45WallggurYhKGag=",
        version = "v0.0.0-20240228011516-70dd3763d340",
    )
    go_repository(
        name = "io_k8s_kubernetes",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kubernetes",
        sum = "h1:qTfB+u5M92k2fCCCVP2iuhgwwSOv1EkAkvQY1tQODD8=",
        version = "v1.13.0",
    )
    go_repository(
        name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
        sum = "h1:fmRfl9WJ4ApJn7LxNuED4m0t18qivVQOxP6aAYG9J6c=",
        version = "v0.0.22",
    )
    go_repository(
        name = "io_k8s_sigs_json",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/json",
        sum = "h1:EDPBXCAspyGV4jQlpZSudPeMmr1bNJefnuqLsRAsHZo=",
        version = "v0.0.0-20221116044647-bc3834ca7abd",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v4",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/structured-merge-diff/v4",
        sum = "h1:150L+0vs/8DA78h1u02ooW1/fFq/Lwr+sGiqlzvrtq4=",
        version = "v4.4.1",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:Mk1wCc2gy/F0THH0TAp1QYyJNzRm2KCLy3o5ASXVI5E=",
        version = "v1.4.0",
    )
    go_repository(
        name = "io_k8s_utils",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "k8s.io/utils",
        sum = "h1:pUdcCO1Lk/tbT5ztQWOBi5HBgbBP1J8+AsQnQCKsi8A=",
        version = "v0.0.0-20240711033017-18e509b52bc8",
    )
    go_repository(
        name = "io_moul_http2curl_v2",
        build_file_proto_mode = "disable",
        importpath = "moul.io/http2curl/v2",
        sum = "h1:9r3JfDzWPcbIklMOs2TnIFzDYvfAZvjeavG6EzP7jYs=",
        version = "v2.3.0",
    )
    go_repository(
        name = "io_opencensus_go",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opentelemetry_go_auto_sdk",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/auto/sdk",
        sum = "h1:cH53jehLUN6UFLY71z+NDOiNJqDdPRaXzTel0sJySYA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "io_opentelemetry_go_collector_pdata",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/collector/pdata",
        sum = "h1:g02K8jlRnmQ7TQDuXpdgVL6vIxIVqr5Gbb1qIR27rto=",
        version = "v1.16.0",
    )
    go_repository(
        name = "io_opentelemetry_go_collector_semconv",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/collector/semconv",
        sum = "h1:KHQnOHe3gUz0zsxe8ph9kN5OTypCFD4V+06AiBTfeNk=",
        version = "v0.110.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib",
        sum = "h1:ubFQUn0VCZ0gPwIoJfBJVpeBlyRMxu8Mm/huKWYd9p0=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_detectors_gcp",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/detectors/gcp",
        sum = "h1:F7q2tNlCaHY9nMKHR6XH9/qkp8FktLnIcy6jJNyOCQw=",
        version = "v1.36.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_github_com_gin_gonic_gin_otelgin",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin",
        sum = "h1:fZNpsQuTwFFSGC96aJexNOBrCD7PjD9Tm/HyHtXhmnk=",
        version = "v0.62.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_github_com_gorilla_mux_otelmux",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux",
        sum = "h1:HxJLvY878W39Q/yHlZW//4TXCPNth9t1MV1DcpoXzs0=",
        version = "v0.46.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:x7wzEgXfnzJcHDwStJT+mxOz4etr2EcexjqhBvmoakw=",
        version = "v0.60.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_httptrace_otelhttptrace",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace",
        sum = "h1:4BZHA+B1wXEQoGNHxW8mURaLhcdGwvRnmhGbm+odRbc=",
        version = "v0.56.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:Hf9xI/XLML9ElpiHVDNwvqI0hIFlzV8dgIr35kV1kRU=",
        version = "v0.62.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_runtime",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/instrumentation/runtime",
        sum = "h1:ZIt0ya9/y4WyRIzfLC8hQRRsWg0J9M9GyaGtIMiElZI=",
        version = "v0.62.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_propagators_b3",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/propagators/b3",
        sum = "h1:0aGKdIuVhy5l4GClAjl72ntkZJhijf2wg1S7b5oLoYA=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_propagators_jaeger",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/propagators/jaeger",
        sum = "h1:f4beMGDKiVzg9IcX7/VuWVy+oGdjx3dNJ72YehmtY5k=",
        version = "v1.21.1",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_samplers_jaegerremote",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/contrib/samplers/jaegerremote",
        sum = "h1:Qb+5A+JbIjXwO7l4HkRUhgIn4Bzz0GNS2q+qdmSx+0c=",
        version = "v0.15.1",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:9zhNfelUvx0KBfu/gb+ZgeAfAgtWrfHJZcAqFC228wQ=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_jaeger",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/jaeger",
        sum = "h1:D7UpUy2Xc2wsi1Ras6V40q806WM07rqoCWzXu7Sqy+4=",
        version = "v1.17.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp",
        sum = "h1:PTNgq9MRmQqqJY0REVbZFvwkYOA85vbdQU/nVfxDyqg=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
        sum = "h1:R/OBkMoGgfy2fLhs2QhkCI1w4HLEQX92GCcJB6SSdNk=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlpmetric_otlpmetricgrpc",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc",
        sum = "h1:zG8GlgXCJQd5BU98C0hZnBbElszTmUgCNCfYneaDL0A=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlpmetric_otlpmetrichttp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp",
        sum = "h1:9PgnL3QNlj10uGxExowIDIZu66aVBwWhXmbOp1pa6RA=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        sum = "h1:Ahq7pZmv87yiyn3jeFz/LekZmPLLdKejuO3NcK9MssM=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
        sum = "h1:EtFWSnwW9hGObjkIdmlnWSydO+Qs8OwzfzXLUPg4xOc=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracehttp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp",
        sum = "h1:bDMKF3RUSxshZ5OjOTi8rsHGaPKsAt76FaqgvIUySLc=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_stdout_stdoutmetric",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric",
        sum = "h1:PB3Zrjs1sG1GBX51SXyTSoOTqcDglmsk7nT6tkKPb/k=",
        version = "v1.35.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_stdout_stdouttrace",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/stdout/stdouttrace",
        sum = "h1:SNhVp/9q4Go/XHBkQ1/d5u9P/U+L1yaGPoi0x+mStaI=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_zipkin",
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/zipkin",
        sum = "h1:D+Gv6lSfrFBWmQYyxKjDd0Zuld9SRXpIrEsKZvE4DO4=",
        version = "v1.21.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:mvwbQS5m0tbmqML4NqK+e3aDiO02vsf/WgbsdpcPoZE=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_oteltest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/oteltest",
        sum = "h1:HiITxCawalo5vQzdHfKeZurV8x7ljcqAgiWzF6Vaeaw=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:ItB0QUqnjesGRvNcmAcU0LyvkVyGJ2xftD29bWdDvKI=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_export_metric",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk/export/metric",
        sum = "h1:c5VRjxCXdQlx1HjzwGdQHzZaVI82b5EbBgOu2ljD92g=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:90lI228XrB9jCMuSdA0673aubgRobVZFhbjxHHspCPc=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:HLdcFNbRQBE2imdSEgm/kwqmQj1Or1l/7bW6mxVK7z4=",
        version = "v1.37.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:jX1VolD6nHuFzOYso2E73H85i92Mv8JQYk0K9vz09os=",
        version = "v1.7.0",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "io_rsc_pdf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "rsc.io/pdf",
        sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_rsc_quote_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "rsc.io/quote/v3",
        sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
        version = "v3.1.0",
    )
    go_repository(
        name = "io_rsc_sampler",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "rsc.io/sampler",
        sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_bazil_fuse",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "bazil.org/fuse",
        sum = "h1:SRsZGA7aFnCZETmov57jwPrWuTmaZK6+4R4v5FUe1/c=",
        version = "v0.0.0-20200407214033-5883e5a4b512",
    )
    go_repository(
        name = "org_gioui",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gioui.org",
        sum = "h1:+wYws3ydNacEpOr3j473gLFMF7yMfj5xdRg2teuC51g=",
        version = "v0.0.0-20230404125508-ad3db5212d10",
    )
    go_repository(
        name = "org_gioui_cpu",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gioui.org/cpu",
        sum = "h1:AGDDxsJE1RpcXTAxPG2B4jrwVUJGFDjINIPi1jtO6pc=",
        version = "v0.0.0-20210817075930-8d6a761490d2",
    )
    go_repository(
        name = "org_gioui_shader",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gioui.org/shader",
        sum = "h1:cvZmU+eODFR2545X+/8XucgZdTtEjR3QWW6W65b0q5Y=",
        version = "v1.0.6",
    )
    go_repository(
        name = "org_gioui_x",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gioui.org/x",
        sum = "h1:Ks31mE4kQd8nWfMVv+KxAL0tO/veB2822tdqnNnDNBs=",
        version = "v0.0.0-20230403130642-fd712aa4daf5",
    )
    go_repository(
        name = "org_go4_intern",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go4.org/intern",
        sum = "h1:UXLjNohABv4S58tHmeuIZDO6e3mHpW2Dx33gaNt03LE=",
        version = "v0.0.0-20211027215823-ae77deb06f29",
    )
    go_repository(
        name = "org_go4_unsafe_assume_no_moving_gc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go4.org/unsafe/assume-no-moving-gc",
        sum = "h1:FyBZqvoA/jbNzuAWLQE2kG820zMAkcilx6BMjGbL/E4=",
        version = "v0.0.0-20220617031537-928513b29760",
    )
    go_repository(
        name = "org_golang_google_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/api",
        sum = "h1:iGZfjXAJiUFSSaekVB7LzXl6tRfEKhUN7FkZN++07tI=",
        version = "v0.233.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_cloud",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/cloud",
        sum = "h1:Cpp2P6TPjujNoC5M2KHY6g7wfyLYfIWRZaSdIKfDasA=",
        version = "v0.0.0-20151119220103-975617b05ea8",
    )
    go_repository(
        name = "org_golang_google_genproto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto",
        replace = "google.golang.org/genproto",
        sum = "h1:rHWScKit0gvAPuOnu87KpaYtjK5zBMLcULh7gxkCXu4=",
        version = "v0.0.0-20250603155806-513f23925822",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:mVXdvnmR3S3BQOqHECm9NGMjYiRtEvDYcqAqedTXY6s=",
        version = "v0.0.0-20250721164621-a45f3dfb1074",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:DbpkGFGRkd4GORg+IWQW2EhxUaa/My/PM8d1CGyTDMY=",
        version = "v0.0.0-20250505200425-f936aa4a68b2",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:MAKi5q709QWfnkkpNQ0M12hYJ1+e8qYVDyowc4U1XZM=",
        version = "v0.0.0-20250728155136-f173205681a0",
    )
    go_repository(
        name = "org_golang_google_grpc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:WoosgB65DlWVC9FqI82dGsZhWFNBSLjQ84bjROOpMu4=",
        version = "v1.74.2",
    )
    go_repository(
        name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        sum = "h1:rNBFJjBCOgVr9pWD7rs/knKL4FRTKgpZmsRfV214zcA=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_golang_google_grpc_examples",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc/examples",
        sum = "h1:MLBCGN1O7GzIx+cBiwfYPwtmZ41U3Mn/cotLJciaArI=",
        version = "v0.0.0-20230224211313-3775f633ce20",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/protobuf",
        sum = "h1:z1NpPI8ku2WgiWnf+t9wTPsn6eP1L7ksHUlkfLvd9xY=",
        version = "v1.36.6",
    )
    go_repository(
        name = "org_golang_x_arch",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/arch",
        sum = "h1:WN9poc33zL4AzGxqf8VtpKUnGvMi8O9lhNyBMF/85qc=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/crypto",
        sum = "h1:cKRW/pmt1pKAfetfu+RCEvjvZkA9RimPbh7bhFjGVBU=",
        version = "v0.46.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/exp",
        sum = "h1:2dVuKD2vS7b0QIHQbpyTISPd0LeHDbnYEryqj5Q1ug8=",
        version = "v0.0.0-20240719175910-8a7402abbf56",
    )
    go_repository(
        name = "org_golang_x_exp_shiny",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/exp/shiny",
        sum = "h1:ryT6Nf0R83ZgD8WnFFdfI8wCeyqgdXWN4+CkFVNPAT0=",
        version = "v0.0.0-20220827204233-334a2380cb91",
    )
    go_repository(
        name = "org_golang_x_image",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/image",
        sum = "h1:jGzIakQa/ZXI1I0Fxvaa9W7yP25TqT6cHIHn+6CqvSQ=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_lint",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mobile",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mobile",
        sum = "h1:4+4C/Iv2U4fMZBiMCc98MG1In4gJY5YRhtpDNeDeHWs=",
        version = "v0.0.0-20190719004257-d2bd2a29d028",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/mod",
        sum = "h1:HaW9xtz0+kOcWKwli0ZXy79Ix+UW/vOfmWI5QVd2tgI=",
        version = "v0.31.0",
    )
    go_repository(
        name = "org_golang_x_net",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/net",
        sum = "h1:zyQRTTrjc33Lhh0fBgT/H3oZq9WuvRR5gPC70xpDiQU=",
        version = "v0.48.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/oauth2",
        sum = "h1:dnDm7JmhM45NNpd8FDDeLhK6FwqbOf4MLCM9zb1BOHI=",
        version = "v0.30.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sync",
        sum = "h1:vV+1eWNmZ5geRlYjzm2adRgW2/mcpevXNg50YZtPCE4=",
        version = "v0.19.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/sys",
        sum = "h1:Ivj+2Cp/ylzLiEU89QhWblYnOE9zerudt9Ftecq2C6k=",
        version = "v0.41.0",
    )
    go_repository(
        name = "org_golang_x_telemetry",
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/telemetry",
        sum = "h1:bH6xUXay0AIFMElXG2rQ4uiE+7ncwtiOdPfYK1NK2XA=",
        version = "v0.0.0-20251203150158-8fff8a5912fc",
    )
    go_repository(
        name = "org_golang_x_term",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/term",
        sum = "h1:PQ5pkm/rLO6HnxFR7N2lJHOZX6Kez5Y1gDSJla6jo7Q=",
        version = "v0.38.0",
    )
    go_repository(
        name = "org_golang_x_text",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/text",
        sum = "h1:B3njUFyqtHDUI5jMn1YIr5B0IE2U0qck04r6d4KPAxE=",
        version = "v0.33.0",
    )
    go_repository(
        name = "org_golang_x_time",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/time",
        sum = "h1:/bpjEDfN9tkoN/ryeYHnv5hcMlc8ncjMcM4XBk5NWV0=",
        version = "v0.11.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/tools",
        sum = "h1:yLkxfA+Qnul4cs9QA3KnlFu0lVmd8JJfoq+E41uSutA=",
        version = "v0.40.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "golang.org/x/xerrors",
        sum = "h1:noIWHXmPHxILtqtCOPIhSt0ABwskkZKjD3bXGnZGpNY=",
        version = "v0.0.0-20240903120638-7835f813f4da",
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gonum.org/v1/gonum",
        sum = "h1:xKuo6hzt+gMav00meVPUlXwSdoEJP46BR+wdxQEFK2o=",
        version = "v0.12.0",
    )
    go_repository(
        name = "org_gonum_v1_netlib",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gonum.org/v1/netlib",
        sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
        version = "v0.0.0-20190313105609-8cb42192e0e0",
    )
    go_repository(
        name = "org_gonum_v1_plot",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gonum.org/v1/plot",
        sum = "h1:y1ZNmfz/xHuHvtgFe8USZVyykQo5ERXPnspQNVK15Og=",
        version = "v0.12.0",
    )
    go_repository(
        name = "org_modernc_cc_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/cc/v3",
        sum = "h1:P3g79IUS/93SYhtoeaHW+kRCIrYaxJ27MFPv+7kaTOw=",
        version = "v3.40.0",
    )
    go_repository(
        name = "org_modernc_ccgo_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/ccgo/v3",
        sum = "h1:Mkgdzl46i5F/CNR/Kj80Ri59hC8TKAhZrYSaqvkwzUw=",
        version = "v3.16.13",
    )
    go_repository(
        name = "org_modernc_ccorpus",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/ccorpus",
        sum = "h1:J16RXiiqiCgua6+ZvQot4yUuUy8zxgqbqEEUuGPlISk=",
        version = "v1.11.6",
    )
    go_repository(
        name = "org_modernc_httpfs",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/httpfs",
        sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
        version = "v1.0.6",
    )
    go_repository(
        name = "org_modernc_libc",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/libc",
        sum = "h1:wymSbZb0AlrjdAVX3cjreCHTPCpPARbQXNz6BHPzdwQ=",
        version = "v1.22.4",
    )
    go_repository(
        name = "org_modernc_mathutil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/mathutil",
        sum = "h1:rV0Ko/6SfM+8G+yKiyI830l3Wuz1zRutdslNoQ0kfiQ=",
        version = "v1.5.0",
    )
    go_repository(
        name = "org_modernc_memory",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/memory",
        sum = "h1:N+/8c5rE6EqugZwHii4IFsaJ7MUhoWX07J5tC/iI5Ds=",
        version = "v1.5.0",
    )
    go_repository(
        name = "org_modernc_opt",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/opt",
        sum = "h1:3XOZf2yznlhC+ibLltsDGzABUGVx8J6pnFMS3E4dcq4=",
        version = "v0.1.3",
    )
    go_repository(
        name = "org_modernc_sqlite",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/sqlite",
        sum = "h1:ixuUG0QS413Vfzyx6FWx6PYTmHaOegTY+hjzhn7L+a0=",
        version = "v1.21.2",
    )
    go_repository(
        name = "org_modernc_strutil",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/strutil",
        sum = "h1:fNMm+oJklMGYfU9Ylcywl0CO5O6nTfaowNsh2wpPjzY=",
        version = "v1.1.3",
    )
    go_repository(
        name = "org_modernc_tcl",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/tcl",
        sum = "h1:mOQwiEK4p7HruMZcwKTZPw/aqtGM4aY00uzWhlKKYws=",
        version = "v1.15.1",
    )
    go_repository(
        name = "org_modernc_token",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/token",
        sum = "h1:Xl7Ap9dKaEs5kLoOQeQmPWevfnk/DM5qcLcYlA8ys6Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_modernc_z",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "modernc.org/z",
        sum = "h1:xkDw/KepgEjeizO2sNco+hqYkU12taxQFqPEmgm1GWE=",
        version = "v1.7.0",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:P98w8egYRjYe3XDjxhYJagTokP/H6HzlsnojRgZRd80=",
        version = "v1.14.0",
    )
    go_repository(
        name = "org_mozilla_go_pkcs7",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.mozilla.org/pkcs7",
        sum = "h1:A/5uWzF44DlIgdm/PQFwfMkW0JX+cIcQi/SwLAmZP5M=",
        version = "v0.0.0-20200128120323-432b2356ecb1",
    )
    go_repository(
        name = "org_uber_go_atomic",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/atomic",
        sum = "h1:ZvwS0R+56ePWxUNi+Atn9dWONBPp/AUETXlHW0DxSjE=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_automaxprocs",
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/automaxprocs",
        sum = "h1:O3y2/QNTOdbF+e/dpXNNW7Rx2hZ4sTIPyybbxyNqTUs=",
        version = "v1.6.0",
    )
    go_repository(
        name = "org_uber_go_goleak",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/goleak",
        sum = "h1:2K3zAYmnTNqV73imy9J1T3WC+gmCePx2hEGkimedGto=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/multierr",
        sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_tools",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/tools",
        sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
        version = "v0.0.0-20190618225709-2cfd321de3ee",
    )
    go_repository(
        name = "org_uber_go_zap",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "go.uber.org/zap",
        sum = "h1:WefMeulhovoZ2sYXz7st6K0sLj7bBhpiFaud4r4zST8=",
        version = "v1.21.0",
    )
    go_repository(
        name = "st_wow_git_gmp_jni",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "git.wow.st/gmp/jni",
        sum = "h1:bGG/g4ypjrCJoSvFrP5hafr9PPB5aw8SjcOWWila7ZI=",
        version = "v0.0.0-20210610011705-34026c7e22d0",
    )
    go_repository(
        name = "tech_einride_go_aip",
        build_file_proto_mode = "disable",
        importpath = "go.einride.tech/aip",
        sum = "h1:16/AfSxcQISGN5z9C5lM+0mLYXihrHbQ1onvYTr93aQ=",
        version = "v0.68.1",
    )
    go_repository(
        name = "tools_gotest",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gotest.tools",
        sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
        version = "v2.2.0+incompatible",
    )
    go_repository(
        name = "tools_gotest_gotestsum",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gotest.tools/gotestsum",
        sum = "h1:szU3TaSz8wMx/uG+w/A2+4JUPwH903YYaMI9yOOYAyI=",
        version = "v1.8.2",
    )
    go_repository(
        name = "tools_gotest_v3",
        build_directives = ["\"gazelle:default_visibility //visibility:public\""],
        build_file_proto_mode = "disable",
        importpath = "gotest.tools/v3",
        sum = "h1:EENdUnS3pdur5nybKYIh2Vfgc8IUNBjxDPSjtiJcOzU=",
        version = "v3.5.1",
    )
