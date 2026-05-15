# Changelog

## 0.8.0 (2026-05-15)

Full Changelog: [v0.7.1...v0.8.0](https://github.com/y2-intel/y2-go/compare/v0.7.1...v0.8.0)

### Features

* **api:** api update ([b61a20f](https://github.com/y2-intel/y2-go/commit/b61a20f0098ed182705b0e6984aa312378a4168d))
* **client:** optimize json encoder for internal types ([43225ba](https://github.com/y2-intel/y2-go/commit/43225ba0d2ae26fa2d068e39b2ec04cca1dfdf6d))
* **go:** add default http client with timeout ([62b74a3](https://github.com/y2-intel/y2-go/commit/62b74a3739f89e9552d08c00d218a42903f28c2c))
* **internal:** support comma format in multipart form encoding ([eeb52df](https://github.com/y2-intel/y2-go/commit/eeb52dff148b2c9966abeafa3bcce17bf3d1fcc2))
* support setting headers via env ([10bd63e](https://github.com/y2-intel/y2-go/commit/10bd63e51c2b9d4b40ddeec14fd0c817b0eb3b3f))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([55727f0](https://github.com/y2-intel/y2-go/commit/55727f07fbe9357ebf72c2f43e435c221ccca628))
* prevent duplicate ? in query params ([cce9ca3](https://github.com/y2-intel/y2-go/commit/cce9ca3827594184831e9856a6485d40ee23e754))


### Chores

* avoid embedding reflect.Type for dead code elimination ([fb83fc5](https://github.com/y2-intel/y2-go/commit/fb83fc5f263ad5c67f51c3be4cdd21a7d9ea74fc))
* **ci:** skip lint on metadata-only changes ([0f65e40](https://github.com/y2-intel/y2-go/commit/0f65e40d71014e222df2f9e699942ea381cf8849))
* **ci:** support opting out of skipping builds on metadata-only commits ([6a38c9e](https://github.com/y2-intel/y2-go/commit/6a38c9eba437ef726c5f0465915a8e1dc4921ae5))
* **client:** fix multipart serialisation of Default() fields ([a8c9e98](https://github.com/y2-intel/y2-go/commit/a8c9e98a4fe30a86ffd14c35606ba85109fc985d))
* **internal:** more robust bootstrap script ([3014456](https://github.com/y2-intel/y2-go/commit/3014456bab5fd1e4e020656b9fb10b4615509511))
* **internal:** support default value struct tag ([ca31221](https://github.com/y2-intel/y2-go/commit/ca31221a96f0253b3a6de947de9dd8dac0d9a151))
* redact api-key headers in debug logs ([d722e02](https://github.com/y2-intel/y2-go/commit/d722e021e8e69db97ef9c3ff87c909cfa8576761))
* remove unnecessary error check for url parsing ([286386f](https://github.com/y2-intel/y2-go/commit/286386f98c59536baec42477a3cb7d3835374a62))
* update docs for api:"required" ([029c823](https://github.com/y2-intel/y2-go/commit/029c823e9d19b197ee15631579b422825cffa4f8))

## 0.7.1 (2026-03-24)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/y2-intel/y2-go/compare/v0.7.0...v0.7.1)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([d1d4a22](https://github.com/y2-intel/y2-go/commit/d1d4a22eb90423a77ab21a196095673dc4dba1dc))
* **internal:** minor cleanup ([3c95082](https://github.com/y2-intel/y2-go/commit/3c9508290479a1a99c1ad5450a03aed8ba8e12eb))
* **internal:** tweak CI branches ([ad9cb29](https://github.com/y2-intel/y2-go/commit/ad9cb2903ec851b67d2fd65adc3057ab4f47353d))
* **internal:** update gitignore ([784fe11](https://github.com/y2-intel/y2-go/commit/784fe113bbc341d022674b37d297cb7a419b8298))
* **internal:** use explicit returns ([434d72f](https://github.com/y2-intel/y2-go/commit/434d72fe5324935ee7ee2c5ea672a621a60b1a4f))
* **internal:** use explicit returns in more places ([191a625](https://github.com/y2-intel/y2-go/commit/191a62536a7d427351049b6e1e8e90f0020a8e14))

## 0.7.0 (2026-03-07)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/y2-intel/y2-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** manual updates ([bf61517](https://github.com/y2-intel/y2-go/commit/bf61517270f3ef78918d595362a985c255348857))


### Chores

* **internal:** codegen related update ([cb6715b](https://github.com/y2-intel/y2-go/commit/cb6715b249f4e9cd2365bada11d59548857c3b80))

## 0.6.0 (2026-03-06)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/y2-intel/y2-go/compare/v0.5.0...v0.6.0)

### Features

* **api:** api update ([2417241](https://github.com/y2-intel/y2-go/commit/2417241a6b619216e95c7bb5c4bcea417a593912))

## 0.5.0 (2026-03-04)

Full Changelog: [v0.4.1...v0.5.0](https://github.com/y2-intel/y2-go/compare/v0.4.1...v0.5.0)

### Features

* **api:** api update ([10b2082](https://github.com/y2-intel/y2-go/commit/10b2082599e8ac35a5b787d6d8f7a8f23b9791ad))


### Chores

* **internal:** codegen related update ([069dade](https://github.com/y2-intel/y2-go/commit/069dadec4aa98726f88e5e079fefae28f6b2fdd7))

## 0.4.1 (2026-02-25)

Full Changelog: [v0.4.0...v0.4.1](https://github.com/y2-intel/y2-go/compare/v0.4.0...v0.4.1)

### Bug Fixes

* allow canceling a request while it is waiting to retry ([acfea7f](https://github.com/y2-intel/y2-go/commit/acfea7fe2b684991018554459dc631b72aedd3f2))
* **encoder:** correctly serialize NullStruct ([4fccf82](https://github.com/y2-intel/y2-go/commit/4fccf8243eaaba81920f67f1c19be9769e2515a5))


### Chores

* **internal:** move custom custom `json` tags to `api` ([841a070](https://github.com/y2-intel/y2-go/commit/841a070f734d6e19d0b5419fb64477228489d57c))
* **internal:** remove mock server code ([a12e1b6](https://github.com/y2-intel/y2-go/commit/a12e1b67299bd61882d3fa3097b48d0d10a42bcd))
* update mock server docs ([5e195ab](https://github.com/y2-intel/y2-go/commit/5e195ab603c42c5f8b56f80f031b3aa60c322d36))

## 0.4.0 (2026-02-10)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/y2-intel/y2-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** manual updates ([0e106d4](https://github.com/y2-intel/y2-go/commit/0e106d478fa5c6b5602fe7f50d58891cd56cff30))

## 0.3.0 (2026-01-24)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/y2-intel/y2-go/compare/v0.2.0...v0.3.0)

### Features

* **client:** add a convenient param.SetJSON helper ([5dab84f](https://github.com/y2-intel/y2-go/commit/5dab84fc6e56a84945055f1fc8817cec12ad9a61))


### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([3bc46ee](https://github.com/y2-intel/y2-go/commit/3bc46eec5724f9f80f81a4b16502055b9c0178f1))


### Chores

* **internal:** codegen related update ([2751d60](https://github.com/y2-intel/y2-go/commit/2751d603054448532851e6dfad066bef56fc4d64))
* **internal:** update `actions/checkout` version ([8474f68](https://github.com/y2-intel/y2-go/commit/8474f684274591f952ef3985582e2669c745d6c2))

## 0.2.0 (2025-12-31)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/y2-intel/y2-go/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([84b49f9](https://github.com/y2-intel/y2-go/commit/84b49f9c6a294a72bd394fed5695d26276aaab35))

## 0.1.0 (2025-12-28)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/y2-intel/y2-go/compare/v0.0.2...v0.1.0)

### Features

* **api:** manual updates ([0257f69](https://github.com/y2-intel/y2-go/commit/0257f69db3ae027bd6cd990f859857c485eaf190))

## 0.0.2 (2025-12-24)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/y2-intel/y2-go/compare/v0.0.1...v0.0.2)

### Chores

* update SDK settings ([67676f2](https://github.com/y2-intel/y2-go/commit/67676f24c0fa0c6c37f3b6262479862c13626ed2))
