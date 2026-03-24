# 우선순위 큐

**난이도:** 중
**출처:** [Java Priority Queue](https://www.hackerrank.com/challenges/java-priority-queue)

## 문제 설명

학교에서 학생들이 서비스를 받기 위해 대기하고 있다. 두 가지 이벤트가 발생할 수 있다:

1. **ENTER**: 학생이 우선순위 큐에 들어온다
2. **SERVED**: 가장 높은 우선순위의 학생이 서비스를 받고 큐에서 제거된다

각 학생은 고유한 ID, 이름, CGPA(학점)를 가진다. 우선순위는 다음 기준으로 결정된다:

1. CGPA가 가장 높은 학생이 먼저 서비스를 받는다
2. CGPA가 같으면 이름의 사전순(오름차순)으로 결정한다
3. CGPA와 이름이 모두 같으면 ID가 작은 학생이 먼저 서비스를 받는다

모든 이벤트를 처리한 후, 큐에 남아있는 학생들의 이름을 우선순위 순서대로 출력하시오. 큐가 비어있으면 "EMPTY"를 출력한다.

## 입력 형식

- 첫째 줄에 이벤트의 개수 N이 주어진다
- 다음 N개의 줄에 이벤트가 주어진다:
  - `ENTER name CGPA id`: 학생이 큐에 들어온다
  - `SERVED`: 가장 높은 우선순위의 학생이 서비스를 받는다

## 출력 형식

- 큐에 남아있는 학생들의 이름을 우선순위 순서대로 한 줄에 하나씩 출력한다
- 큐가 비어있으면 "EMPTY"를 출력한다

## 제약 조건

- 1 ≤ N ≤ 1000
- 0 ≤ CGPA ≤ 4.00
- 1 ≤ id ≤ 10^5
- 이름은 영문 알파벳으로만 구성된다

## 예제

### 예제 입력 1

```text
12
ENTER John 3.75 50
ENTER Mark 3.8 24
ENTER Shafaet 3.7 35
SERVED
SERVED
ENTER Samiha 3.85 36
SERVED
ENTER Ashley 3.9 42
ENTER Maria 3.6 46
ENTER Anik 3.95 49
ENTER Dan 3.95 50
SERVED
```

### 예제 출력 1

```text
Dan
Ashley
Shafaet
Maria
```
