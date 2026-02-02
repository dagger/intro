# Intro Toolchain

This [toolchain](https://docs.dagger.io/core-concepts/toolchains) helps you get started with Dagger in your project.

It provides [checks](https://docs.dagger.io/core-concepts/checks) to walk you through onboarding Dagger in your project.

## Get started

1. [Install Dagger](https://docs.dagger.io/getting-started/installation)

2. Install this [toolchain](https://docs.dagger.io/core-concepts/toolchains) in your project

```shell
# From the root of your projects repository
dagger init

dagger toolchain install github.com/dagger/intro
```

3. List the [checks](https://docs.dagger.io/core-concepts/checks) this toolchain provides:

```
dagger check -l

Name                      Description
install-intro             Install your first Toolchain
login-to-cloud            Login to Dagger Cloud with dagger login
install-more-toolchains   Install your next toolchain
```

4. Execute the [checks](https://docs.dagger.io/core-concepts/checks)

```
dagger check
```

The failed checks will provide instructions to assist you through your dagger onboarding!
