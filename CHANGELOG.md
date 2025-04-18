## v0.10.0 (2025-04-15)


- Merge pull request #44 from Drafteame/feat/AH/add-xray-traces
- feat: add xray traces (GAM-236)
- feat: add xray traces

## v0.9.0 (2024-12-19)


- feat: refactor tests to not use gnomock (#36)
- * feat: refactor test to not use gnomock
- * chore: set parallel to all tests
- * chore: add mongo serice to tests
- ci: bump goreleaser/goreleaser-action from 5 to 6 (#31)
- ci: bump github.com/fatih/color from 1.16.0 to 1.17.0 (#29)
- Bumps [github.com/fatih/color](https://github.com/fatih/color) from 1.16.0 to 1.17.0.
- [Release notes](https://github.com/fatih/color/releases)
- [Commits](https://github.com/fatih/color/compare/v1.16.0...v1.17.0)
- ---
updated-dependencies:
- dependency-name: github.com/fatih/color
  dependency-type: direct:production
  update-type: version-update:semver-minor
...
- Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
- ci: bump actions/setup-go from 4 to 5 (#26)
- Bumps [actions/setup-go](https://github.com/actions/setup-go) from 4 to 5.
- [Release notes](https://github.com/actions/setup-go/releases)
- [Commits](https://github.com/actions/setup-go/compare/v4...v5)
- ---
updated-dependencies:
- dependency-name: actions/setup-go
  dependency-type: direct:production
  update-type: version-update:semver-major
...
- Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
- ci: bump actions/setup-python from 4 to 5 (#25)
- Bumps [actions/setup-python](https://github.com/actions/setup-python) from 4 to 5.
- [Release notes](https://github.com/actions/setup-python/releases)
- [Commits](https://github.com/actions/setup-python/compare/v4...v5)
- ---
updated-dependencies:
- dependency-name: actions/setup-python
  dependency-type: direct:production
  update-type: version-update:semver-major
...
- Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

## v0.8.0 (2024-07-25)


- Merge pull request #33 from Drafteame/feat/add-with-search-orders
- feat: add withSearchOrders
- feat: add withSearchOrders

## v0.7.3 (2024-03-18)


- fix: use go 1.21 on pipelines (#24)

## v0.7.2 (2024-03-18)


- fix: release files (#23)

## v0.7.1 (2024-03-18)


- deps: upgrade (#22)
- * deps: upgrade to fix security issues

* ci: upgrade precommit hooks

* ci: change test formatter

## v0.7.0 (2024-03-18)


- build(deps): bump goreleaser/goreleaser-action from 4 to 5 (#16)
- Bumps [goreleaser/goreleaser-action](https://github.com/goreleaser/goreleaser-action) from 4 to 5.
- [Release notes](https://github.com/goreleaser/goreleaser-action/releases)
- [Commits](https://github.com/goreleaser/goreleaser-action/compare/v4...v5)

---
updated-dependencies:
- dependency-name: goreleaser/goreleaser-action
  dependency-type: direct:production
  update-type: version-update:semver-major
...

Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>
- build(deps): bump actions/checkout from 3 to 4 (#15)
- Bumps [actions/checkout](https://github.com/actions/checkout) from 3 to 4.
- [Release notes](https://github.com/actions/checkout/releases)
- [Changelog](https://github.com/actions/checkout/blob/main/CHANGELOG.md)
- [Commits](https://github.com/actions/checkout/compare/v3...v4)

---
updated-dependencies:
- dependency-name: actions/checkout
  dependency-type: direct:production
  update-type: version-update:semver-major
...

Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

## v0.6.0 (2024-03-18)


- build(deps): bump go.mongodb.org/mongo-driver from 1.12.0 to 1.12.1 (#13)
- Bumps [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver) from 1.12.0 to 1.12.1.
- [Release notes](https://github.com/mongodb/mongo-go-driver/releases)
- [Commits](https://github.com/mongodb/mongo-go-driver/compare/v1.12.0...v1.12.1)

---
updated-dependencies:
- dependency-name: go.mongodb.org/mongo-driver
  dependency-type: direct:production
  update-type: version-update:semver-patch
...

Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

## v0.5.0 (2023-10-11)


- Merge pull request #18 from Drafteame/feat/force-search-limit
- feat: force search limit
- fix: wrong variable
- fix: add max search limit variable
- feat: force search limit

## v0.4.4 (2023-09-12)


- Merge pull request #17 from Drafteame/feat/remove-default-sort-by-id
- fix: remove unnecesary default sort
- fix: remove unnecesary default sort

## v0.4.3 (2023-06-29)


- Merge pull request #11 from Drafteame/fix/remove_search_limit_in_constructor
- fix: remove search limit from search options constructor
- fix: remove search limit from search options constructor

## v0.4.2 (2023-06-26)


- Merge pull request #10 from Drafteame/reafactor/add-so-as-generic-type
- refactor: decouple search signature and generify search options type
- refactor: decouple search signature and generify search options type

## v0.4.1 (2023-06-26)


- Merge pull request #9 from Drafteame/reafactor/add-so-as-generic-type
- refactor: add search orders to search options as generic type
- refactor: add search orders to search options as generic type

## v0.4.0 (2023-06-26)


- feat: add search interfaces to be able to replace defined structs (#7)
- * feat: add search interfaces to be able to replace defined structs

* chore: add interface assertion to builtin structs

## v0.3.0 (2023-06-26)


- build(deps): bump go.mongodb.org/mongo-driver from 1.11.6 to 1.12.0 (#8)
- Bumps [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver) from 1.11.6 to 1.12.0.
- [Release notes](https://github.com/mongodb/mongo-go-driver/releases)
- [Commits](https://github.com/mongodb/mongo-go-driver/compare/v1.11.6...v1.12.0)

---
updated-dependencies:
- dependency-name: go.mongodb.org/mongo-driver
  dependency-type: direct:production
  update-type: version-update:semver-minor
...

Signed-off-by: dependabot[bot] <support@github.com>
Co-authored-by: dependabot[bot] <49699333+dependabot[bot]@users.noreply.github.com>

## v0.2.0 (2023-06-01)


- feat: add mongo driver options constructor (#5)
- * feat: add mongo driver options constructor

* fix: error on empty client

## v0.1.1 (2023-05-31)


- fix: no timestamps conf (#4)
- * feat: add fill featured generic repo

* feat: add HardDelete method

* feat: add HardDeleteMany function

* refactor: generify SortOrders to be able to build precedence

* feat: add order constants

* refactor: examples

* refactor: raname internal log methods

* fix: logger tests

* refactor: use primitive.ObjectID instead of pointer

* fix: add testing and implementation for no timestamps

* fix: add updated add field on soft deletes

* deps: upgrade

* fix: assert dates

* ci: remove unneded lines

## v0.1.0 (2023-05-30)


- feat: full featured (#1)
- * feat: add fill featured generic repo

* ci: add ci configuration

* feat: add HardDelete method

* feat: add HardDeleteMany function

* refactor: generify SortOrders to be able to build precedence

* feat: add order constants

* refactor: examples

* docs: change readme

* deps: upgrade precommit

* fix: simplify validations delete many

* refactor: raname internal log methods

* fix: logger tests

* refactor: generify SearchOptions to a static struct instead of implementation

* refactor: use Now helper instead clock.Now property

* feat: add projection to search options
- Create README.md
- Initial commit
