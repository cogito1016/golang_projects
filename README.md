# golang_projects

누가 고랭 좋다고 고랭?

## 1.세팅

### 1.1.설치

- https://go.dev/learn/ 접속
- 최신버전 Go 컴파일러 설치

```bash
> go version
go version go1.20.6 darwin/arm64
```

### 1.2.VsCode Extension

- 익스텐션 : go

#### 1.2.1.GO익스텐션 종속성의 문제

- 이 확장 프로그램은 go, gopls, dlv 및 기타 선택적 도구에 따라 달라집니다. 종속성 중 하나라도 누락된 경우 ⚠️ 분석 도구 누락 경고가 표시됩니다. 경고를 클릭하여 종속성을 다운로드하세요.
- 확장프로그램 다운로드시 왼쪽하단에 GO아이콘과 버전이 표기되는데, 이걸 클릭하면 경고창을 확인 가능, 따라서 종속성다운로드가 가능하다.

```bash
Tools environment: GOPATH=~
Installing 1 tool at ~/bin in module mode.
  gopls

Installing golang.org/x/tools/gopls@latest (~/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go. :)
Installing golang.org/x/tools/gopls@latest (~/bin/gopls) SUCCEEDED

All tools successfully installed. You are ready to Go. :)
```

### 1.3.gopls가 작업영역에서 모듈을 찾지못하는 문제

```bash
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go list
```

- 오류는 위와같은데 저 링크들어가면 404뜸. 쫌마니당혹쓰

#### 1.3.1.원인

- 여러 디렉토리에 go module 존재하면 생기는 문제

#### 1.3.2.조치

- Settings > gopls > Edit in settings.json > gopls json 영역에 아래설정 추가

```json
"experimentalWorkspaceModule": true
```

---

## 2.문법

### 2.1.package 사용

- package는 import하여 가져온다.
- 대문자로 시작하는것들은 Export된다.
- 따라서 적재한 패키지의 함수를 사용하려면 대문자의것들을 참조해야한다.
