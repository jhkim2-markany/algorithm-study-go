# 요구사항 문서

## 소개

67개 알고리즘 폴더(01~67번)의 모든 solution.go 파일을 HackerRank 스타일로 리팩토링한다. 현재 각 문제 폴더에는 완전한 풀이가 담긴 solution.go 하나만 존재한다. 이를 학습자가 핵심 알고리즘 함수만 구현하면 되는 solution.go(보일러플레이트 + 빈 함수 시그니처)와 완전한 정답 코드인 answer.go로 분리한다. 학습자는 solution.go를 열어 보일러플레이트가 준비된 상태에서 핵심 함수만 작성하고, 막히면 answer.go를 참고할 수 있다.

## 용어 정의

- **문제_폴더**: 각 알고리즘 폴더의 problems/ 하위에 위치하는 개별 문제 디렉토리 (예: 01-easy-matrix-rotation)
- **기존_solution**: 리팩토링 전 문제_폴더에 존재하는 완전한 풀이 코드 파일 (solution.go)
- **신규_solution**: 리팩토링 후 보일러플레이트와 빈 핵심 함수 시그니처를 포함하는 학습용 파일 (solution.go)
- **정답_파일**: 리팩토링 후 신규_solution과 동일한 보일러플레이트에 핵심_함수의 구현이 채워진 완성본 파일 (answer.go). 학습자가 solution.go를 완성했을 때의 최종 결과물과 동일한 형태이다
- **보일러플레이트**: package main 선언, import 문, main 함수(입출력 처리), 헬퍼 함수(readLine, checkError 등)를 포함하는 코드 골격
- **핵심_함수**: 학습자가 직접 구현해야 하는 알고리즘 로직을 담는 함수. main 함수에서 호출되며, 빈 함수 시그니처와 안내 주석만 포함한다
- **검증_스크립트**: 프로젝트 구조와 코드 규칙을 자동으로 검증하는 validate.sh 및 validate/validate_test.go

## 요구사항

### 요구사항 1: solution.go 파일 분리 (HackerRank 스타일 변환)

**사용자 스토리:** 학습자로서, 문제를 열었을 때 입출력 처리가 이미 준비되어 있고 핵심 알고리즘 함수만 구현하면 되는 환경을 원한다. 그래야 알고리즘 로직에만 집중할 수 있다.

#### 인수 조건

1. WHEN 리팩토링을 수행하면, THE 각 문제_폴더 SHALL 기존_solution의 완전한 코드를 정답_파일(answer.go)로 보존한다
2. WHEN 리팩토링을 수행하면, THE 각 문제_폴더 SHALL 신규_solution(solution.go)을 생성하여 보일러플레이트와 빈 핵심_함수 시그니처를 포함한다
3. THE 신규_solution SHALL package main 선언, 필요한 import 문, main 함수(입력 읽기 및 출력 처리 로직)를 포함한다
4. THE 신규_solution의 핵심_함수 SHALL 올바른 매개변수와 반환 타입을 가진 함수 시그니처를 포함하되, 함수 본문은 안내 주석("// 여기에 코드를 작성하세요")과 제로값 반환문만 포함한다
5. THE 신규_solution의 main 함수 SHALL 핵심_함수를 호출하여 결과를 출력하는 코드를 포함한다
6. THE 정답_파일(answer.go) SHALL 신규_solution과 동일한 보일러플레이트 구조를 가지되, 핵심_함수의 본문에 완전한 알고리즘 구현 코드를 포함한다. 즉 answer.go는 solution.go의 빈 함수를 채운 완성본이다
7. THE 리팩토링 SHALL 67개 알고리즘 폴더의 모든 문제_폴더(총 201개)에 적용한다

### 요구사항 2: 핵심 함수 설계 규칙

**사용자 스토리:** 학습자로서, 핵심 함수의 시그니처와 한국어 설명을 보고 무엇을 구현해야 하는지 명확히 이해하고 싶다. 그래야 함수의 역할과 입출력을 파악한 후 구현에 집중할 수 있다.

#### 인수 조건

1. THE 핵심_함수 SHALL 해당 문제의 알고리즘 로직을 캡슐화하는 의미 있는 함수명을 사용한다 (예: rotateMatrix, findSumPair, countInversions)
2. THE 핵심_함수 위에 SHALL 함수의 역할을 설명하는 한국어 주석을 포함한다 (예: // rotateMatrix는 N×N 행렬을 시계 방향으로 90도 회전한 결과를 반환한다)
3. THE 핵심_함수의 매개변수 SHALL 기존_solution의 main 함수에서 읽어들이는 입력 데이터를 기반으로 설계한다
4. THE 핵심_함수의 반환 타입 SHALL 기존_solution의 main 함수에서 출력하는 결과 데이터를 기반으로 설계한다
5. WHEN 기존_solution에 main 함수 외부에 정의된 헬퍼 함수가 존재하면, THE 신규_solution SHALL 해당 헬퍼 함수를 보일러플레이트에 포함하되 핵심 알고리즘 로직이 포함된 함수는 빈 시그니처로 제공한다
6. WHEN 기존_solution에 전역 변수가 존재하면, THE 신규_solution SHALL 해당 전역 변수를 보일러플레이트에 포함한다

### 요구사항 3: 코드 품질 유지

**사용자 스토리:** 학습자로서, 리팩토링 후에도 모든 코드가 기존과 동일한 품질 기준을 만족하길 원한다. 그래야 answer.go를 참고하여 학습하거나 solution.go를 완성한 후 바로 실행할 수 있다.

#### 인수 조건

1. THE 신규_solution SHALL `package main` 선언을 포함한다
2. THE 신규_solution SHALL `func main()` 함수를 포함한다
3. THE 신규_solution SHALL 한국어 주석을 최소 1개 이상 포함한다
4. THE 신규_solution SHALL 표준 라이브러리만 import한다
5. THE 정답_파일 SHALL `package main` 선언을 포함한다
6. THE 정답_파일 SHALL `func main()` 함수를 포함한다
7. THE 정답_파일 SHALL 한국어 주석을 최소 1개 이상 포함한다
8. THE 정답_파일 SHALL 표준 라이브러리만 import한다
9. THE 정답_파일 SHALL `go run answer.go` 명령으로 실행 가능한 완전한 코드를 포함한다
10. THE 신규_solution의 핵심_함수 위에 SHALL 함수의 목적, 매개변수 설명, 반환값 설명을 포함하는 한국어 주석 블록을 포함한다 (알고리즘 힌트는 포함하지 않는다)
11. THE 정답_파일의 핵심_함수 위에 SHALL 함수의 목적, 매개변수 설명, 반환값 설명, 알고리즘 힌트를 포함하는 상세한 한국어 주석 블록을 포함한다

### 요구사항 4: 검증 스크립트 업데이트

**사용자 스토리:** 학습자로서, 검증 스크립트가 answer.go 파일도 함께 검증하길 원한다. 그래야 프로젝트 전체의 구조적 정합성을 자동으로 확인할 수 있다.

#### 인수 조건

1. THE 검증_스크립트(validate.sh) SHALL 각 문제_폴더에 answer.go 파일이 존재하는지 확인한다
2. THE 검증_스크립트(validate.sh) SHALL 각 문제_폴더에 solution.go 파일이 존재하는지 계속 확인한다
3. THE Go 테스트(validate/validate_test.go) SHALL answer.go 파일에 대해 기존 solution.go와 동일한 코드 규칙(package main, func main(), 한국어 주석, 표준 라이브러리만 사용)을 검증한다
4. THE Go 테스트(validate/validate_test.go) SHALL solution.go 파일에 대해 기존 코드 규칙(package main, func main(), 한국어 주석, 표준 라이브러리만 사용)을 계속 검증한다
5. THE Go 테스트(validate/validate_test.go) SHALL 각 문제_폴더의 폴더 구조 검증 시 answer.go 존재 여부를 추가로 확인한다

### 요구사항 5: 리팩토링 후 파일 구조

**사용자 스토리:** 학습자로서, 리팩토링 후 각 문제 폴더의 파일 구조가 일관되길 원한다. 그래야 어떤 문제를 열어도 동일한 학습 경험을 할 수 있다.

#### 인수 조건

1. THE 리팩토링 후 각 문제_폴더 SHALL 다음 4개 파일을 포함한다: problem.md, solution.go, answer.go, explanation.md
2. THE 리팩토링 후 문제_폴더 SHALL 기존의 problem.md와 explanation.md를 변경 없이 유지한다
3. THE 리팩토링 SHALL 기존 파일 구조(알고리즘 폴더, problems/ 디렉토리, 난이도별 하위 폴더)를 변경하지 않는다
