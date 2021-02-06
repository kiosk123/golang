# golang 연습

- go를 설치한 후 GOPATH가 가르키는 디렉터리에 src 폴더가 없으면 만든다.
- src 디렉토리의 하위 디렉터리는 다음과 같은 규칙으로 생성
  - src/도메인명(ex: github.com)/유저아이디(ex: kiosk123)/프로젝트폴더
- golang [API](https://golang.org/pkg/) 

## 설치

GOPATH가 가르키는 경로로 이동후 명령 실행  
**go query** [프로젝트](https://github.com/PuerkitoBio/goquery)

```bash
cd $GOPATH
go get github.com/PuerkitoBio/goquery
```

**echo** [프로젝트](https://github.com/labstack/echo)
```bash
cd $GOPATH
go get github.com/labstack/echo
```

go fullstack 프레임워크 [buffalo](https://gobuffalo.io/en/)

