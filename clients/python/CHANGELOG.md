# Changelog

## [3.16.0](https://github.com/phrase/strings-openapi/compare/python-v3.15.1...python-v3.16.0) (2026-02-11)


### Features

* **API:** Accept branch param in RepoSyncs#import ([#1035](https://github.com/phrase/strings-openapi/issues/1035)) ([fdfd542](https://github.com/phrase/strings-openapi/commit/fdfd54203510f94d546d771e1eeb6f4cc40f9b06))
* **API:** add pr_branch param to repo-sync export #SCD-594 ([#978](https://github.com/phrase/strings-openapi/issues/978)) ([7c3c19f](https://github.com/phrase/strings-openapi/commit/7c3c19fa6b7254315ce48a1ad420b97d529dbd9a))
* **API:** creating upload batches #SCD-687 ([#1022](https://github.com/phrase/strings-openapi/issues/1022)) ([83f5a41](https://github.com/phrase/strings-openapi/commit/83f5a411f1ebc63d26cdce0729093aa517d79b62))
* **python:** Relax strict urllib3 version pin #SCD-642 ([#1048](https://github.com/phrase/strings-openapi/issues/1048)) ([33028bb](https://github.com/phrase/strings-openapi/commit/33028bb4504bb5b3ec7af3d0fb7ce14bda0659de))


### Bug Fixes

* remove project from account locales ([#1025](https://github.com/phrase/strings-openapi/issues/1025)) ([06b780f](https://github.com/phrase/strings-openapi/commit/06b780f91d2fe22a22ec821aa18968ed3b50510a))

## [3.15.1](https://github.com/phrase/strings-openapi/compare/python-v3.15.0...python-v3.15.1) (2026-01-16)


### Bug Fixes

* **API:** Add linked_translation to translation type ([#1005](https://github.com/phrase/strings-openapi/issues/1005)) ([b18e64c](https://github.com/phrase/strings-openapi/commit/b18e64c948d40a9e5f996e8fdf6dee07d1812b02))

## [3.15.0](https://github.com/phrase/openapi/compare/python-v3.14.0...python-v3.15.0) (2026-01-09)


### Features

* **API:** Add missing project flags to project details schema #SCD-141 ([#993](https://github.com/phrase/openapi/issues/993)) ([93fbdd7](https://github.com/phrase/openapi/commit/93fbdd76150ed94374ef888ac92c0c9b626de7c8))


### Bug Fixes

* **API:** drop invalid account locale params ([#992](https://github.com/phrase/openapi/issues/992)) ([87af83c](https://github.com/phrase/openapi/commit/87af83c94fd7e47340b5393847e93a31127abbe8))

## [3.14.0](https://github.com/phrase/openapi/compare/python-v3.13.0...python-v3.14.0) (2025-12-11)


### Features

* **API:** Add create comparison endpoint for branches [SCD-549] ([#966](https://github.com/phrase/openapi/issues/966)) ([4099e32](https://github.com/phrase/openapi/commit/4099e32d42744f9d73346077ef30b0f4bd5c849e))

## [3.13.0](https://github.com/phrase/openapi/compare/python-v3.12.1...python-v3.13.0) (2025-12-09)


### Features

* **API:** add manual triggering of automations #SCM-953 ([#968](https://github.com/phrase/openapi/issues/968)) ([ebe1c68](https://github.com/phrase/openapi/commit/ebe1c68e472d87ea6c6cd2f76bb421406c30a274))
* **API:** add use_locale_fallback option to download #SCD-620 ([#976](https://github.com/phrase/openapi/issues/976)) ([9266c68](https://github.com/phrase/openapi/commit/9266c680d8717bb736aed83625858f0b3419da54))


### Bug Fixes

* **python:** pin urllib3 version to &lt;2.6 #SCD-642 ([#979](https://github.com/phrase/openapi/issues/979)) ([04177c5](https://github.com/phrase/openapi/commit/04177c5dbb6a918209ecd51d63655c95f5ed0445))

## [3.12.1](https://github.com/phrase/openapi/compare/python-v3.12.0...python-v3.12.1) (2025-11-10)


### Bug Fixes

* **API:** Add omit_translation_keys to Jobs#show endpoint ([#957](https://github.com/phrase/openapi/issues/957)) ([23ece66](https://github.com/phrase/openapi/commit/23ece66b05fc748a4598ce876cb2954ed1a385cc))

## [3.12.0](https://github.com/phrase/openapi/compare/python-v3.11.0...python-v3.12.0) (2025-10-27)


### Features

* **API:** Support omit_translation_keys param in Jobs#show ([#944](https://github.com/phrase/openapi/issues/944)) ([690a2ce](https://github.com/phrase/openapi/commit/690a2ce992359e7920367699145d2d4ff3a206fa))

## [3.11.0](https://github.com/phrase/openapi/compare/python-v3.10.0...python-v3.11.0) (2025-10-23)


### Features

* **API:** Support source_locale_id param in uploads ([#935](https://github.com/phrase/openapi/issues/935)) ([304a406](https://github.com/phrase/openapi/commit/304a4061aaac8001a3a9e10c5cf985b04c31dab2))

## [3.10.0](https://github.com/phrase/openapi/compare/python-v3.9.0...python-v3.10.0) (2025-10-16)


### Features

* List child branches ([#926](https://github.com/phrase/openapi/issues/926)) ([31d3b57](https://github.com/phrase/openapi/commit/31d3b57e0d1381149409b77c3e236b226cde22a1))

## [3.9.0](https://github.com/phrase/openapi/compare/python-v3.8.0...python-v3.9.0) (2025-10-01)


### Features

* add branch sync endpoint ([#912](https://github.com/phrase/openapi/issues/912)) ([3293917](https://github.com/phrase/openapi/commit/329391757b3f81574448b1b87506b85bfdeb6761))

## [3.8.0](https://github.com/phrase/openapi/compare/python-v3.7.0...python-v3.8.0) (2025-09-30)


### Features

* **API:** Add "reviewed" param for tanslations#update ([#891](https://github.com/phrase/openapi/issues/891)) ([3f77286](https://github.com/phrase/openapi/commit/3f7728669fa68d5911dfca39a9e39c24cd20c0a4))
* **SCM-732:** expose Automation API ([#910](https://github.com/phrase/openapi/issues/910)) ([6eea53d](https://github.com/phrase/openapi/commit/6eea53d06aca8f43623ba90e2e0a8cd6761c9a13))

## [3.7.0](https://github.com/phrase/openapi/compare/python-v3.6.0...python-v3.7.0) (2025-08-01)


### Features

* **API:** Add support for `update_custom_metadata` option on upload ([#882](https://github.com/phrase/openapi/issues/882)) ([5b486ea](https://github.com/phrase/openapi/commit/5b486eabf30c84402a8e31911a5ba6fe3b343a89))

## [3.6.0](https://github.com/phrase/openapi/compare/python-v3.5.0...python-v3.6.0) (2025-07-24)


### Features

* **API:** Job and job locale annotations #SCD-145 ([#867](https://github.com/phrase/openapi/issues/867)) ([95b6c2a](https://github.com/phrase/openapi/commit/95b6c2a2ab798d01f9bc0d53a22dbf817a5eb0ee))

## [3.5.0](https://github.com/phrase/openapi/compare/python-v3.4.0...python-v3.5.0) (2025-07-18)


### Features

* **API:** add updated_since param for downloads ([#858](https://github.com/phrase/openapi/issues/858)) ([db4196b](https://github.com/phrase/openapi/commit/db4196bfb1a45f628577bebb99aad7da1179c241))
* **API:** Add use_ordinal_rules attribute to request body for create/update Keys endpoints ([#868](https://github.com/phrase/openapi/issues/868)) ([29410b5](https://github.com/phrase/openapi/commit/29410b5d4edee9645f449d2d7de53fb953c1f0c2))
* **API:** Add use_ordinal_rules boolean [STRINGS-2273] ([#846](https://github.com/phrase/openapi/issues/846)) ([450ce0c](https://github.com/phrase/openapi/commit/450ce0cce6bb2064758c44b130b39dd5539c2681))

## [3.4.0](https://github.com/phrase/openapi/compare/python-v3.3.0...python-v3.4.0) (2025-05-23)


### Features

* **API:** Add omit_statistics param to tags#show ([#834](https://github.com/phrase/openapi/issues/834)) ([2058b18](https://github.com/phrase/openapi/commit/2058b18297133075885ac99770aee2e171811cd6))

## [3.3.0](https://github.com/phrase/openapi/compare/python-v3.2.0...python-v3.3.0) (2025-05-16)


### Features

* **API:** Add "verify_mentioned_translations" parameter to uploads ([#830](https://github.com/phrase/openapi/issues/830)) ([f2fdf60](https://github.com/phrase/openapi/commit/f2fdf60dd97c3908293cd457a765dafde603bd9b))
* **API:** Add processed & upload total translations on upload summary ([#824](https://github.com/phrase/openapi/issues/824)) ([081db68](https://github.com/phrase/openapi/commit/081db68d2ffcf5d66a81e07eec0a9572a1f9d633))
* **API:** add updated_since filter to job list #STRINGS-1555 ([#799](https://github.com/phrase/openapi/issues/799)) ([dc9b6ed](https://github.com/phrase/openapi/commit/dc9b6ed12e013231d397820449086c87fea2f8ba))
* **API:** Include roles in account response ([#811](https://github.com/phrase/openapi/issues/811)) ([dc27ee5](https://github.com/phrase/openapi/commit/dc27ee5117762222b6e1e6abb639f8e00c6a9101))
* **API:** Translations unreview & batch unreview ([#831](https://github.com/phrase/openapi/issues/831)) ([6a696db](https://github.com/phrase/openapi/commit/6a696db00d80d8acaf5887a08a7a97997566eb8f))


### Bug Fixes

* **API:** Remove request body from GET comments ([#817](https://github.com/phrase/openapi/issues/817)) ([2646001](https://github.com/phrase/openapi/commit/264600132e80ac03983e0ae86e99db3d6fb9080d))

## [3.2.0](https://github.com/phrase/openapi/compare/python-v3.1.0...python-v3.2.0) (2025-02-17)


### Features

* **API:** document tags attribute of an upload #STRINGS-1221 ([#790](https://github.com/phrase/openapi/issues/790)) ([fff505b](https://github.com/phrase/openapi/commit/fff505bdff35a0033fee06e505c42fe794c88562))

## [3.1.0](https://github.com/phrase/openapi/compare/python-v3.0.3...python-v3.1.0) (2025-02-17)


### Features

* **API:** Add locale_ids param to synchronous download endpoint [STRINGS-1492] ([#780](https://github.com/phrase/openapi/issues/780)) ([47186a4](https://github.com/phrase/openapi/commit/47186a44fc8c0b8e466636acf3d49413b1f29f30))
* **API:** Add source last updated at information on job details ([#777](https://github.com/phrase/openapi/issues/777)) ([c9b8423](https://github.com/phrase/openapi/commit/c9b8423766b4138980d0553502b3e18ca524f34e))

## [3.0.3](https://github.com/phrase/openapi/compare/python-v3.0.2...python-v3.0.3) (2025-01-29)


### Bug Fixes

* **API:** pass translation_key_ids when removing keys from job ([#771](https://github.com/phrase/openapi/issues/771)) ([f670e27](https://github.com/phrase/openapi/commit/f670e2763b1112fefd1812109b3c09def42b7bd2))

## [3.0.2](https://github.com/phrase/openapi/compare/python-v3.0.1...python-v3.0.2) (2025-01-06)


### Bug Fixes

* **CLI:** Adjust operationId for quality_performance_score #STRINGS-1104 ([#721](https://github.com/phrase/openapi/issues/721)) ([7aa3b9b](https://github.com/phrase/openapi/commit/7aa3b9b508d1d24a4af7f4977b1a2fead8bfda78))

## [3.0.1](https://github.com/phrase/openapi/compare/python-v3.0.0...python-v3.0.1) (2024-12-20)


### Bug Fixes

* **API:** Repo Sync Event errors field type #STRINGS-1074 ([#756](https://github.com/phrase/openapi/issues/756)) ([c7670e0](https://github.com/phrase/openapi/commit/c7670e04810f95359d72ba6346b5f626bfb77b6f))
* **API:** Repo Sync schema fixes #STRINGS-1074 ([#748](https://github.com/phrase/openapi/issues/748)) ([033be10](https://github.com/phrase/openapi/commit/033be1003fe01b5115de1f8ba2336d32b4862bfd))

## [3.0.0](https://github.com/phrase/openapi/compare/python-v2.6.1...python-v3.0.0) (2024-12-18)


### ⚠ BREAKING CHANGES

* Remove old Git sync endpoints. Replaced with new repo sync ([#735](https://github.com/phrase/openapi/issues/735))

### Features

* **API:** Add 'default_encoding' documentation ([#733](https://github.com/phrase/openapi/issues/733)) ([0139c51](https://github.com/phrase/openapi/commit/0139c51da747fbe7bc9929bcf3534aad7f22f39a))
* Remove old Git sync endpoints. Replaced with new repo sync ([#735](https://github.com/phrase/openapi/issues/735)) ([c3bd8ec](https://github.com/phrase/openapi/commit/c3bd8eccaabcfa1b1066ea4438971ac59833af46))


### Bug Fixes

* **API:** Add missing branch parameter to job comment endpoints #STRINGS-988 ([#724](https://github.com/phrase/openapi/issues/724)) ([64d399c](https://github.com/phrase/openapi/commit/64d399ced0980ac2a48366f91110047287a0c590))

## [2.6.1](https://github.com/phrase/openapi/compare/python-v2.6.0...python-v2.6.1) (2024-11-27)


### Bug Fixes

* **API:** Comment creation schema fix #STRINGS-866 ([#718](https://github.com/phrase/openapi/issues/718)) ([e201d13](https://github.com/phrase/openapi/commit/e201d1360c89698dd8d3642cc28f89dd0e50a1fb))

## [2.6.0](https://github.com/phrase/openapi/compare/python-v2.5.0...python-v2.6.0) (2024-11-20)


### Features

* **API:** Add Pagination header to POST search endpoints [[#457](https://github.com/phrase/openapi/issues/457)] ([#706](https://github.com/phrase/openapi/issues/706)) ([9a79fa3](https://github.com/phrase/openapi/commit/9a79fa31bb3b9d58272fa2f4e82d72d0d44a93a0))
* **API:** autotranslate param in key creation [STRINGS-786] ([#713](https://github.com/phrase/openapi/issues/713)) ([581d0ff](https://github.com/phrase/openapi/commit/581d0ff5f1d06757e5ddd9603b78fc8d435d68ee))


### Bug Fixes

* **Python:** Correctly encode nested upload parameters #STRINGS-864 ([#717](https://github.com/phrase/openapi/issues/717)) ([446fd60](https://github.com/phrase/openapi/commit/446fd60ba007a8761b994ed6bd07ccd625a263fd))

## [2.5.0](https://github.com/phrase/openapi/compare/python-v2.4.0...python-v2.5.0) (2024-10-02)


### Features

* Add translation key prefixes for upload and download ([#687](https://github.com/phrase/openapi/issues/687)) ([9c9c959](https://github.com/phrase/openapi/commit/9c9c959830631bcac8beaf1de30ab31755ac1ee5))


### Bug Fixes

* **API:** Format list is not paginated and authenticated #STRINGS-458 ([#690](https://github.com/phrase/openapi/issues/690)) ([25e90f4](https://github.com/phrase/openapi/commit/25e90f46513e70cf328be80c36ae785cead05851))

## [2.4.0](https://github.com/phrase/openapi/compare/python-v2.3.0...python-v2.4.0) (2024-09-03)


### Features

* Add update_translations_on_source_match ([#670](https://github.com/phrase/openapi/issues/670)) ([11003ac](https://github.com/phrase/openapi/commit/11003ace7353bf99893482ca4aa32214abf3e581))

## [2.3.0](https://github.com/phrase/openapi/compare/python-v2.2.0...python-v2.3.0) (2024-07-15)


### Features

* add repo sync events show endpoint ([#641](https://github.com/phrase/openapi/issues/641)) ([e1d9cfb](https://github.com/phrase/openapi/commit/e1d9cfb23e079fea2d9e5475dff9a4137f1f0154))
* **API:** Async downloads [TSI-2515] ([#642](https://github.com/phrase/openapi/issues/642)) ([6fcab5d](https://github.com/phrase/openapi/commit/6fcab5d4719f64e8e5dd49c327dc9348b384de4c))


### Bug Fixes

* add app_min_version and app_max_version param to releases ([#633](https://github.com/phrase/openapi/issues/633)) ([b384301](https://github.com/phrase/openapi/commit/b3843012460ace4c1d34c4373e5158595466adcb))

## [2.2.0](https://github.com/phrase/openapi/compare/python-v2.1.1...python-v2.2.0) (2024-06-12)


### Features

* **API:** Add OTA Release Triggers API [TSI-2485] ([#622](https://github.com/phrase/openapi/issues/622)) ([8cb91dc](https://github.com/phrase/openapi/commit/8cb91dcce2c19ca700cf9d0713fa74f28ad59434))

## [2.1.1](https://github.com/phrase/openapi/compare/python-v2.1.0...python-v2.1.1) (2024-05-31)


### Bug Fixes

* **CLI:** use id instead of repo_sync_id as param ([#618](https://github.com/phrase/openapi/issues/618)) ([7a1a0d9](https://github.com/phrase/openapi/commit/7a1a0d9115262dc7fa3065ee7cc76bec7a381ddb))
* **PHP:** Fix deserializing custom metadata values in key response ([#607](https://github.com/phrase/openapi/issues/607)) ([b6eeeba](https://github.com/phrase/openapi/commit/b6eeeba223e3eabec268a8f3d372afcb6abd09dd))

## [2.1.0](https://github.com/phrase/openapi/compare/python-v2.0.0...python-v2.1.0) (2024-04-29)


### Features

* **API:** Add 'update_translation_keys' for Uploads [TSI-2292] ([#578](https://github.com/phrase/openapi/issues/578)) ([4492ec0](https://github.com/phrase/openapi/commit/4492ec0a7c62f9de9ab1c1125071615bddcc26ce))

## [2.0.0](https://github.com/phrase/openapi/compare/python-v1.22.0...python-v2.0.0) (2024-04-23)


### ⚠ BREAKING CHANGES

* add missing required params ([#571](https://github.com/phrase/openapi/issues/571))

### Code Refactoring

* add missing required params ([#571](https://github.com/phrase/openapi/issues/571)) ([d810e9e](https://github.com/phrase/openapi/commit/d810e9ebc767e14ba9e56106de8c5774d9d6d178))

## [1.22.0](https://github.com/phrase/openapi/compare/python-v1.21.0...python-v1.22.0) (2024-04-22)


### Features

* Add linked-parent to translation details ([#570](https://github.com/phrase/openapi/issues/570)) ([2c6f432](https://github.com/phrase/openapi/commit/2c6f43253e24b670b71ac810c85dce0759c29403))

## [1.21.0](https://github.com/phrase/openapi/compare/python-v1.20.0...python-v1.21.0) (2024-04-17)


### Features

* **API:** Add Repo Sync [TSI-1923] ([#569](https://github.com/phrase/openapi/issues/569)) ([0bd1756](https://github.com/phrase/openapi/commit/0bd17562018cb045ff41cc1ff5008b9419a0ed12))

## [1.20.0](https://github.com/phrase/openapi/compare/python-v1.19.0...python-v1.20.0) (2024-04-10)


### Features

* **API:** add Linked Keys endpoints ([#555](https://github.com/phrase/openapi/issues/555)) ([4935dac](https://github.com/phrase/openapi/commit/4935dac58c787eaade2f1f65ce649f466b5e3a60))


### Bug Fixes

* (API) Add mandatory params to linked keys endpoints ([#564](https://github.com/phrase/openapi/issues/564)) ([08d9846](https://github.com/phrase/openapi/commit/08d9846bc224d349e2ade9abf28d733afb1e8be3))

## [1.19.0](https://github.com/phrase/openapi/compare/python-v1.18.2...python-v1.19.0) (2024-02-08)


### Features

* add query param for properties ([#542](https://github.com/phrase/openapi/issues/542)) ([b4e12d0](https://github.com/phrase/openapi/commit/b4e12d04fd2916351f9201e1e6de504143ecc9aa))

## [1.18.2](https://github.com/phrase/openapi/compare/python-v1.18.1...python-v1.18.2) (2024-02-05)


### Bug Fixes

* **API:** allow nullable value for job's due_date ([#534](https://github.com/phrase/openapi/issues/534)) ([38b51b5](https://github.com/phrase/openapi/commit/38b51b51095394f8ce769873140038abba628514))

## [1.18.1](https://github.com/phrase/openapi/compare/python-v1.18.0...python-v1.18.1) (2024-02-01)


### Bug Fixes

* **API:** Adjust documentation of QPS endpoint ([#525](https://github.com/phrase/openapi/issues/525)) ([4b4f1ac](https://github.com/phrase/openapi/commit/4b4f1acf28fbd13b3d16c37162cdccfa05c38ffa))

## [1.18.0](https://github.com/phrase/openapi/compare/python-v1.17.0...python-v1.18.0) (2024-01-17)


### Features

* Add reports locales endpoint to API [TSS-2439] ([#465](https://github.com/phrase/openapi/issues/465)) ([e03aa9f](https://github.com/phrase/openapi/commit/e03aa9f49f031517b36db715fe70e8e0b65a438b))
* **API:** add Custom Metadata endpoints ([#474](https://github.com/phrase/openapi/issues/474)) ([d407d8b](https://github.com/phrase/openapi/commit/d407d8be5ccddec1afde14a12804a7a616f77d7a))
* **API:** Add custom_metadata_filters param to locale download endpoint [TSI-2174] ([#478](https://github.com/phrase/openapi/issues/478)) ([3623478](https://github.com/phrase/openapi/commit/3623478fc1518b457ab018b5630a693081637d6e))
* **api:** Add QPS endpoint and documentation ([#521](https://github.com/phrase/openapi/issues/521)) ([d22c558](https://github.com/phrase/openapi/commit/d22c558adfbb7fcd13759e388c038744914e42fa))
* **API:** Add url field to uploads ([#481](https://github.com/phrase/openapi/issues/481)) ([7332a84](https://github.com/phrase/openapi/commit/7332a84f9958346f2fb28dee4b0353519ef466d5))


### Bug Fixes

* **API:** Create Custom Metadata endpoint fix [TSI-2222] ([#499](https://github.com/phrase/openapi/issues/499)) ([ce2ed94](https://github.com/phrase/openapi/commit/ce2ed9488e111fb5d9bc3810a78c47d23553c8b7))

## [1.17.0](https://github.com/phrase/openapi/compare/python-v1.16.0...python-v1.17.0) (2023-11-03)


### Features

* [TSI-2083] enable format_options argument for java-client   ([#426](https://github.com/phrase/openapi/issues/426)) ([faa8cb3](https://github.com/phrase/openapi/commit/faa8cb353ba9f1030b9f7cfd46b894b4d6d26e70))

## [1.16.0](https://github.com/phrase/openapi/compare/python-v1.15.0...python-v1.16.0) (2023-10-30)


### Features

* Update openapi-generator to v7 ([#418](https://github.com/phrase/openapi/issues/418)) ([524626f](https://github.com/phrase/openapi/commit/524626f5e914bfef6025d0e1c2cbc7a728d08f56))

## [1.15.0](https://github.com/phrase/openapi/compare/python-v1.14.1...python-v1.15.0) (2023-10-23)


### Features

* **API:** Add order param to comment list endpoints ([#441](https://github.com/phrase/openapi/issues/441)) ([441c9c4](https://github.com/phrase/openapi/commit/441c9c46169f8c5ac4e71ade09a95dab136314ef))

## [1.14.1](https://github.com/phrase/openapi/compare/python-v1.14.0...python-v1.14.1) (2023-10-23)


### Bug Fixes

* **python:** pass format_options correctly ([#436](https://github.com/phrase/openapi/issues/436)) ([3621400](https://github.com/phrase/openapi/commit/362140060ab90463185b3b1f729b79c08226bf55))

## [1.14.0](https://github.com/phrase/openapi/compare/python-v1.13.0...python-v1.14.0) (2023-10-13)


### Features

* **API:** Implement figma attachments endpoints ([#415](https://github.com/phrase/openapi/issues/415)) ([970e612](https://github.com/phrase/openapi/commit/970e612fda620ca882a221ef541036b8d200b675))

## [1.13.0](https://github.com/phrase/openapi/compare/python-v1.12.0...python-v1.13.0) (2023-09-12)


### Features

* Optionally tag only affected keys on upload [TSI-2041] ([#412](https://github.com/phrase/openapi/issues/412)) ([e8f958e](https://github.com/phrase/openapi/commit/e8f958e91469c2542f44ab68469c933688958383))
* **TSI-1946:** Add reviewed_at to translations ([#396](https://github.com/phrase/openapi/issues/396)) ([3e663d9](https://github.com/phrase/openapi/commit/3e663d971a99a816f0165dd6653a9a1e8a87c95e))

## [1.12.0](https://github.com/phrase/openapi/compare/python-v1.11.0...python-v1.12.0) (2023-08-28)


### Features

* **API:** Document new query parameters ([#393](https://github.com/phrase/openapi/issues/393)) ([770515a](https://github.com/phrase/openapi/commit/770515a9628122955bb3919405babf9392684eb9))

## [1.11.0](https://github.com/phrase/openapi/compare/python-v1.10.0...python-v1.11.0) (2023-08-24)


### Features

* **API:** Introduce comment replies endpoints ([#383](https://github.com/phrase/openapi/issues/383)) ([71351ac](https://github.com/phrase/openapi/commit/71351ac285f4f49976092e176c77b09f3485eb65))

## [1.10.0](https://github.com/phrase/openapi/compare/python-v1.9.1...python-v1.10.0) (2023-08-22)


### Features

* **TSE-950:** Document comment_reactions endpoints ([#380](https://github.com/phrase/openapi/issues/380)) ([f230244](https://github.com/phrase/openapi/commit/f230244e6e9c069b18edc4c35dd5e290fd14793b))


### Bug Fixes

* **schemas:** Fix gitlab_sync type mismatch ([#373](https://github.com/phrase/openapi/issues/373)) ([1cb1f65](https://github.com/phrase/openapi/commit/1cb1f650598c68afee6e2cd7c3c4ede1484aba35))

## [1.9.1](https://github.com/phrase/openapi/compare/python-v1.9.0...python-v1.9.1) (2023-07-28)


### Bug Fixes

* Fix gitlab_sync history status type mismatch ([#363](https://github.com/phrase/openapi/issues/363)) ([ebcaa4e](https://github.com/phrase/openapi/commit/ebcaa4e5dfcb2f73559a56c78b0f2512ca798375))
