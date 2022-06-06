# SBOM FUN!

This repo produces a multi-platform image when built with `ko`, where each platform-specific image has a different SBOM.

```
$ crane manifest ghcr.io/imjasonh/sbom-fun/github.com/imjasonh/sbom-fun | jq '.manifests[].platform'
{
  "architecture": "amd64",
  "os": "linux"
}
{
  "architecture": "arm",
  "os": "linux"
}
{
  "architecture": "arm64",
  "os": "linux"
}
{
  "architecture": "ppc64le",
  "os": "linux"
}
{
  "architecture": "s390x",
  "os": "linux"
}
```

Select a platform and get that image's digest:

```
$ crane digest ghcr.io/imjasonh/sbom-fun/github.com/imjasonh/sbom-fun --platform=linux/arm64
sha256:18f3efcc2edee3500445ee2b3a968834d1396cbcdd6300e4bd8b01b980c50b67
```

Then get that image's SBOM:

```
$ cosign download sbom ghcr.io/imjasonh/sbom-fun/github.com/imjasonh/sbom-fun@sha256:18f3efcc2edee3500445ee2b3a968834d1396cbcdd6300e4bd8b01b980c50b67 | jq -r '.components[].name'
github.com/BurntSushi/toml
github.com/google/go-github/v45
github.com/google/go-querystring
golang.org/x/crypto
```

Select a different platform, get a different SBOM:

```
$ crane digest ghcr.io/imjasonh/sbom-fun/github.com/imjasonh/sbom-fun --platform=linux/ppc64le
sha256:b4c3ee6f0f5093c27f3710a18209e6922f56b3a9e7efef53271f5b5d50b91745
```

Then get that image's SBOM:

```
$ cosign download sbom ghcr.io/imjasonh/sbom-fun/github.com/imjasonh/sbom-fun@sha256:b4c3ee6f0f5093c27f3710a18209e6922f56b3a9e7efef53271f5b5d50b91745 | jq -r '.components[].name'
github.com/BurntSushi/toml
github.com/google/go-cmp
```

