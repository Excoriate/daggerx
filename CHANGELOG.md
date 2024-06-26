# Changelog

## [0.0.16](https://github.com/Excoriate/daggerx/compare/v0.0.15...v0.0.16) (2024-06-02)


### Features

* add plain string to command conversion ([0dcd3a4](https://github.com/Excoriate/daggerx/commit/0dcd3a48795c871bb269c17750aa92b9a2a7ea70))

## [0.0.15](https://github.com/Excoriate/daggerx/compare/v0.0.14...v0.0.15) (2024-06-02)


### Features

* add extensions for github, add new functionality for commands ([763194c](https://github.com/Excoriate/daggerx/commit/763194c88a42e1974a0872f5578fc56942d0bece))
* add extensions for github, add new functionality for commands ([5aa33d2](https://github.com/Excoriate/daggerx/commit/5aa33d22636fa2a1ce783ddcc05c7ada63097c37))

## [0.0.14](https://github.com/Excoriate/daggerx/compare/v0.0.13...v0.0.14) (2024-06-02)


### Features

* add cloudaws package ([5754835](https://github.com/Excoriate/daggerx/commit/57548350cc6b1a2b413ca64de0e1ae94fda4d276))

## [0.0.13](https://github.com/Excoriate/daggerx/compare/v0.0.12...v0.0.13) (2024-05-30)


### Features

* Add DaggerEngine implementation ([#37](https://github.com/Excoriate/daggerx/issues/37)) ([20f6b3a](https://github.com/Excoriate/daggerx/commit/20f6b3a32eb4f9a472c820b1e91f48a045306e33))
* Add tests for GetLogLevelFromEnv and NewLogger ([#36](https://github.com/Excoriate/daggerx/issues/36)) ([b36f920](https://github.com/Excoriate/daggerx/commit/b36f9202b540dbcaf90b859b1212e3543d702a25))

## [0.0.12](https://github.com/Excoriate/daggerx/compare/v0.0.11...v0.0.12) (2024-05-27)


### Features

* Add tests and implementation for converting DaggerCMD to string and generating sh command ([#30](https://github.com/Excoriate/daggerx/issues/30)) ([87a2b4d](https://github.com/Excoriate/daggerx/commit/87a2b4daad428275e10824c6df1609f156d50f2e))
* **cleaner:** Add RemoveCommas function and test cases ([#32](https://github.com/Excoriate/daggerx/issues/32)) ([14c94c6](https://github.com/Excoriate/daggerx/commit/14c94c62b5733c28a9ba350b407f028a514d44b6))


### Refactoring

* Move ToAnyType function from conv package to parser package for better structuring. ([#33](https://github.com/Excoriate/daggerx/issues/33)) ([4da61da](https://github.com/Excoriate/daggerx/commit/4da61dacd90db179648dc1b4542b1b079e77da6d))

## [0.0.11](https://github.com/Excoriate/daggerx/compare/v0.0.10...v0.0.11) (2024-05-27)


### Refactoring

* Refactored BuildArgs function to improve argument processing. ([#28](https://github.com/Excoriate/daggerx/issues/28)) ([b0e0061](https://github.com/Excoriate/daggerx/commit/b0e0061f88194e933e52913f946ab361af67ec06))

## [0.0.10](https://github.com/Excoriate/daggerx/compare/v0.0.9...v0.0.10) (2024-05-26)


### Features

* Add default image name logic ([#26](https://github.com/Excoriate/daggerx/issues/26)) ([fd2d4f2](https://github.com/Excoriate/daggerx/commit/fd2d4f2acedc4723b311b7e47c018c5fcfd9a3fd))

## [0.0.9](https://github.com/Excoriate/daggerx/compare/v0.0.8...v0.0.9) (2024-05-26)


### Refactoring

* Organize import statements in execmd/args.go ([#24](https://github.com/Excoriate/daggerx/issues/24)) ([a2b0dab](https://github.com/Excoriate/daggerx/commit/a2b0dabe03370a32397caa89c570953ff9ffb994))

## [0.0.8](https://github.com/Excoriate/daggerx/compare/v0.0.7...v0.0.8) (2024-05-17)


### Bug Fixes

* convert directly from string to dagger env vars ([b054699](https://github.com/Excoriate/daggerx/commit/b0546992d1902d5cc181a07ae21147d1c7b83d87))

## [0.0.7](https://github.com/Excoriate/daggerx/compare/v0.0.6...v0.0.7) (2024-05-17)


### Other

* fix shell-check issue ([e925f58](https://github.com/Excoriate/daggerx/commit/e925f58450be0c6493b85b3c5f9f2caf21dca31c))


### Refactoring

* remove non-required functions ([f925194](https://github.com/Excoriate/daggerx/commit/f92519468f5a4db1cac68b1263e37e4e608df8a8))

## [0.0.6](https://github.com/Excoriate/daggerx/compare/v0.0.5...v0.0.6) (2024-05-17)


### Refactoring

* adjust conversion functions ([1e662aa](https://github.com/Excoriate/daggerx/commit/1e662aa2521e881978289b57e01997ff95e9a218))

## [0.0.5](https://github.com/Excoriate/daggerx/compare/v0.0.4...v0.0.5) (2024-05-17)


### Refactoring

* adjust api structure, add type converter ([9b43cad](https://github.com/Excoriate/daggerx/commit/9b43cade168e058fae05698d0be8eb10d6cd382a))

## [0.0.4](https://github.com/Excoriate/daggerx/compare/v0.0.3...v0.0.4) (2024-05-17)


### Features

* add built-in base containers per stack ([b6e9124](https://github.com/Excoriate/daggerx/commit/b6e91247b53b1f6c126b24dfcfbb9b72078c960e))

## [0.0.3](https://github.com/Excoriate/daggerx/compare/v0.0.2...v0.0.3) (2024-05-16)


### Features

* Add test for BuildArgs function in execmd package ([#13](https://github.com/Excoriate/daggerx/issues/13)) ([336a12e](https://github.com/Excoriate/daggerx/commit/336a12ef2f46337e90089413b33a9b142cb5b6bd))

## [0.0.2](https://github.com/Excoriate/daggerx/compare/v0.0.1...v0.0.2) (2024-05-16)


### Features

* Add DaggerEnvVars struct and converter functions ([b237e91](https://github.com/Excoriate/daggerx/commit/b237e912e6714206f71a597a5c9d227b2665a15e))
* Add GitHub Actions for label management ([da292aa](https://github.com/Excoriate/daggerx/commit/da292aa2405daf41690a18b9af5652af59ce8753))
* **pkg/execmd:** Add command generation and container setup functions ([#10](https://github.com/Excoriate/daggerx/issues/10)) ([1794e13](https://github.com/Excoriate/daggerx/commit/1794e138b18cad2c75707461005bfedd5a782e85))


### Bug Fixes

* Correct spelling errors in issue templates ([cadddfe](https://github.com/Excoriate/daggerx/commit/cadddfe8b3bf7058f8576e9be0a10a71a977cb6c))
* Correct spelling errors in issue templates ([bffdde0](https://github.com/Excoriate/daggerx/commit/bffdde09fd7805ed72930369ffc7d37eadc87625))


### Other

* first commit ([b20b5c5](https://github.com/Excoriate/daggerx/commit/b20b5c570e49761d491cd1d1f3501af59c194900))
* Update configuration files for golangci-lint, markdownlint, and versions in README ([3e671ec](https://github.com/Excoriate/daggerx/commit/3e671ec2f9a93af513e56d909d5890d9875e39d4))
* Update installation instructions to use daggerx package ([339d878](https://github.com/Excoriate/daggerx/commit/339d87837334718f190a1f942aaab5e80addfcd2))


### Docs

* Added pull request template and issue templates ([3fb38a9](https://github.com/Excoriate/daggerx/commit/3fb38a96a69b140a4c323007edfb1f48a6b0e623))
* Added pull request template and issue templates ([7062e35](https://github.com/Excoriate/daggerx/commit/7062e35150d0fae361bcf866ad763e9639f6d1d4))


### Refactoring

* Delete changelog ([d8d7dba](https://github.com/Excoriate/daggerx/commit/d8d7dbac03a379c407467c244a335b88f0d321ba))
* Delete changelog ([f49663d](https://github.com/Excoriate/daggerx/commit/f49663d9cc57cbbc66463fa25472f3d630009735))
* **envvars:** Rename functions and add handling for empty inputs ([df2df6f](https://github.com/Excoriate/daggerx/commit/df2df6f35763daf62827ad2b86013b8862a51329))
