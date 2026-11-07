# Changelog

## [0.2.0](https://github.com/mnizarzr/wol/compare/v0.1.2...v0.2.0) (2026-11-07)


### Features

* add --config flag to set config file path (Trugamr/wol[#26](https://github.com/mnizarzr/wol/issues/26)) ([b0cab27](https://github.com/mnizarzr/wol/commit/b0cab2754b71546adc156e2f340af5c969993dd5))
* add scheduled wake-ups via cron (Trugamr/wol[#28](https://github.com/mnizarzr/wol/issues/28)) ([7dbd173](https://github.com/mnizarzr/wol/commit/7dbd173508fdc1ed0eca0daba9cacc79070a7fd7))
* configurable broadcast address and port (Trugamr/wol[#29](https://github.com/mnizarzr/wol/issues/29)) ([392aeb2](https://github.com/mnizarzr/wol/commit/392aeb21060ea8be281eaa2c83a2a36b26982631))
* **ui:** redesign with light/dark themes, single-machine hero, per-action flash colors ([5719f2f](https://github.com/mnizarzr/wol/commit/5719f2f32318c0a8a891e8c97022f7d4a42c8df6))


### Bug Fixes

* cap serveApi request body size ([56f545d](https://github.com/mnizarzr/wol/commit/56f545df66f15e06eefa0b6c9ad8658ca7d99fe7))
* don't log on each status poll for unresolvable hosts (Trugamr/wol[#27](https://github.com/mnizarzr/wol/issues/27)) ([732282d](https://github.com/mnizarzr/wol/commit/732282dfa0587e0d48e6052ca5f11e02efa1dec3))
* keep machines in status map when ping fails ([38565b0](https://github.com/mnizarzr/wol/commit/38565b0cbc2fb1b8efbead68fe46d1b1711d088e))
* preserve non-ASCII machine names in flash cookie (Trugamr/wol[#37](https://github.com/mnizarzr/wol/issues/37)) ([16da539](https://github.com/mnizarzr/wol/commit/16da5397f54f430276cd1d97cbf9da94fe25e906))
* reject send actions other than wake or sleep ([5cfa665](https://github.com/mnizarzr/wol/commit/5cfa665245ba9b203da66c7ab793bfdd78e43396))
* satisfy golangci-lint gofmt and errcheck ([c591798](https://github.com/mnizarzr/wol/commit/c591798e968a57a628c3b9d5508650b29f04fd8f))


### Performance Improvements

* parse index template once at startup ([a7dd342](https://github.com/mnizarzr/wol/commit/a7dd342c31a1a6d06c046566a240a27e6c12fcf7))
