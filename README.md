# Downloader

## usage

아래와 같이 `sample.png.hashget` 파일을 작성합니다
```yaml
hash:
    sha128: 262084257C2103702EF8A25705E3F8DBC1FA3823103AD7B954D54BDB77E6D89D
where:
    - https://www.google.co.kr/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png
    # ipfs link
    # torrent link
    #...
//dsadsa
```

같은경로에서 도구를 실행합니다.
```bash
downloader *.hashget
```

`sample.png` 가 로컬에 저장 되어야 합니다.
파일 확보에 실패하거나, 확보된 파일의 해시가 맞지않을경우 저장되지 않습니다.

`--file-log-split` 파라미터를 추가하여 `sample.png.hashget.log` 와 같이 각 파일별 진행로그를 남길 수 있습니다.
