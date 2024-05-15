# Changelog

## [0.0.14](https://github.com/Excoriate/tftest/compare/v0.0.13...v0.0.14) (2024-04-28)


### Features

* **cloudprovider:** add support for new AWS services ([#25](https://github.com/Excoriate/tftest/issues/25)) ([f70dae1](https://github.com/Excoriate/tftest/commit/f70dae1b2ad10d56db8daecb8f394dfd2e424424))

## [0.0.13](https://github.com/Excoriate/tftest/compare/v0.0.12...v0.0.13) (2024-04-17)


### Refactoring

* enhance git repo validator function ([b791280](https://github.com/Excoriate/tftest/commit/b7912804beed53ff48c05cde9695b1b839f6a7da))
* expose APIs for finding git repo by different methods ([17e7960](https://github.com/Excoriate/tftest/commit/17e7960964a2e9c815674fde606cdbe2c560e2d9))

## [0.0.12](https://github.com/Excoriate/tftest/compare/v0.0.11...v0.0.12) (2024-04-16)


### Other

* adding S3 client aws cloud provider ([#22](https://github.com/Excoriate/tftest/issues/22)) ([36950d4](https://github.com/Excoriate/tftest/commit/36950d4f594bac5d80c51f312ad8db42b297cec1))

## [0.0.11](https://github.com/Excoriate/tftest/compare/v0.0.10...v0.0.11) (2024-04-16)


### Features

* add support for JSONPath testcases ([#18](https://github.com/Excoriate/tftest/issues/18)) ([cb46401](https://github.com/Excoriate/tftest/commit/cb46401affebe2dbfe2fe72d19e7c0650f9e7ca9))

## [0.0.10](https://github.com/Excoriate/tftest/compare/v0.0.9...v0.0.10) (2024-04-02)


### Features

* Add test for ProcessTemplFile function ([#15](https://github.com/Excoriate/tftest/issues/15)) ([064c951](https://github.com/Excoriate/tftest/commit/064c951ab491b8a5c845c231cb678aedcc6df394))

## [0.0.9](https://github.com/Excoriate/tftest/compare/v0.0.8...v0.0.9) (2024-04-01)


### Features

* Add function to check if a directory is a git repository ([#13](https://github.com/Excoriate/tftest/issues/13)) ([6ad6ec0](https://github.com/Excoriate/tftest/commit/6ad6ec0d8c40d7f86bc603dd2df3becd1babe8cb))

## [0.0.8](https://github.com/Excoriate/tftest/compare/v0.0.7...v0.0.8) (2024-03-28)


### Features

* Improve Terraform variables and paths handling ([#12](https://github.com/Excoriate/tftest/issues/12)) ([48b0537](https://github.com/Excoriate/tftest/commit/48b05375d5de08ca7382740fcf02366663657afb))
* **tfvars:** Add function to get .tfvars files from a directory ([#10](https://github.com/Excoriate/tftest/issues/10)) ([a01a7db](https://github.com/Excoriate/tftest/commit/a01a7db5e60a2e26a671ad2658bb04aad09eb323))


### Refactoring

* Refactor GetRelativePathFromGitRepo function signature and variable names ([#9](https://github.com/Excoriate/tftest/issues/9)) ([10bbfa5](https://github.com/Excoriate/tftest/commit/10bbfa54e9730870be9808fd512438d6b5a6d943))

## [0.0.7](https://github.com/Excoriate/tftest/compare/v0.0.6...v0.0.7) (2024-03-28)


### Features

* add support for absolute path in the setup terraform for parallelism ([#6](https://github.com/Excoriate/tftest/issues/6)) ([3cc6a3a](https://github.com/Excoriate/tftest/commit/3cc6a3ab2815c43593a3ba9edd97a0e377cfa2f5))


### Refactoring

* Update Terraform directory setup for parallelism ([#8](https://github.com/Excoriate/tftest/issues/8)) ([2d06eea](https://github.com/Excoriate/tftest/commit/2d06eea752b120d19063278ee6b3c623bbf3cbe3))

## [0.0.6](https://github.com/Excoriate/tftest/compare/v0.0.5...v0.0.6) (2024-03-18)


### Features

* add paralellism as a built-in function into the client creation ([7e83867](https://github.com/Excoriate/tftest/commit/7e83867bbe03ff8b6732c41c17f840fccec6c09e))

## [0.0.5](https://github.com/Excoriate/tftest/compare/v0.0.4...v0.0.5) (2024-03-18)


### Refactoring

* remove options creation on advanced options ([5cc9f6d](https://github.com/Excoriate/tftest/commit/5cc9f6da9b29417cf022e7df5c0a791bdfcd9bad))

## [0.0.4](https://github.com/Excoriate/tftest/compare/v0.0.3...v0.0.4) (2024-03-18)


### Other

* make golangci happy ([ce5d0bf](https://github.com/Excoriate/tftest/commit/ce5d0bfc88a6e04a007b7ba8d417efaf1cd4ae7b))
* remove structure saving during initialization ([43a79b9](https://github.com/Excoriate/tftest/commit/43a79b9365b0ccfc378c4589621e601ba1ddfea4))

## [0.0.3](https://github.com/Excoriate/tftest/compare/v0.0.2...v0.0.3) (2024-03-18)


### Other

* amend 404 link on readme.md ([ef42f17](https://github.com/Excoriate/tftest/commit/ef42f17e66e6207e46679e5be8cb3fc78e47fd4b))


### Refactoring

* avoid naming collision on packages ([e930ccf](https://github.com/Excoriate/tftest/commit/e930ccf0d0ecf642e9a578f5bca2649d9ff904cf))

## [0.0.2](https://github.com/Excoriate/tftest/compare/v0.0.1...v0.0.2) (2024-03-18)


### Features

* add stages ([10995d4](https://github.com/Excoriate/tftest/commit/10995d43bd8e06fc38cc28f0d0905736c1aeec85))
* add stages apis ([a105513](https://github.com/Excoriate/tftest/commit/a10551313cff7532c8f4959dab083c7bef8ffe6b))
* first structural commit ([1f2de4a](https://github.com/Excoriate/tftest/commit/1f2de4ad92ba4ff86d363b2f3234d481dddb7c42))
* first structural commit ([e1f2dc7](https://github.com/Excoriate/tftest/commit/e1f2dc7673c81ca9884c3743872975729bc773ed))


### Bug Fixes

* remove unused test ([54b9551](https://github.com/Excoriate/tftest/commit/54b95516f9b065c8acab100108fac3a65f4d2a64))


### Docs

* add docs ([23ecb66](https://github.com/Excoriate/tftest/commit/23ecb66955007179a6bf37c37dc6c2f9df73cec8))
