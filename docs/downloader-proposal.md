# Checksum manifest downloader proposal

This preserves the downloader usage proposal from [pull request #1](https://github.com/gosuda/unipath/pull/1) alongside the current Unipath protocol toolkit. The manifest reader, checksum verification, and command-line options below are not implemented by the current downloader command.

## Proposed manifest

A file named `sample.png.hashget` would contain the expected digest and candidate sources:

```yaml
hash:
  sha256: 262084257C2103702EF8A25705E3F8DBC1FA3823103AD7B954D54BDB77E6D89D
where:
  - https://www.google.co.kr/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png
  # Additional supported sources could include IPFS or torrent URLs.
```

The URL and digest above are the original proposal's example, not a live download fixture.

## Proposed behavior

Running `downloader *.hashget` in the same directory would try the manifest's sources and save the verified content as `sample.png`.

The downloader must compare the complete downloaded content with the expected SHA-256 digest before publishing the destination file. A failed download or a digest mismatch must leave no unverified destination file.

An optional `--file-log-split` flag would write progress for each manifest to a separate file such as `sample.png.hashget.log`.
