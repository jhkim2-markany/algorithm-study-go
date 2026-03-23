package validate

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// 67개 알고리즘 폴더 목록
// **Validates: Requirements 8.2, 9.4**
var algorithmFolders = []string{
	"01-implementation-and-simulation",
	"02-bruteforce",
	"03-sorting",
	"04-stack-and-queue",
	"05-hash",
	"06-prefix-sum",
	"07-math-and-number-theory",
	"08-binary-search",
	"09-parametric-search",
	"10-two-pointer-and-sliding-window",
	"11-greedy",
	"12-heap-and-priority-queue",
	"13-tree",
	"14-binary-tree",
	"15-graph-dfs",
	"16-graph-bfs",
	"17-backtracking",
	"18-divide-and-conquer",
	"19-dynamic-programming",
	"20-union-find",
	"21-shortest-path",
	"22-minimum-spanning-tree",
	"23-topological-sort",
	"24-graph-advanced",
	"25-segment-tree",
	"26-string-algorithm",
	"27-geometry",
	"28-combinatorics",
	"29-bitmask",
	"30-game-theory",
	"31-probability",
	"32-bitmask-dp",
	"33-maximum-flow",
	"34-primality-test",
	"35-offline-queries",
	"36-exponentiation-by-squaring",
	"37-knapsack",
	"38-dag",
	"39-coordinate-compression",
	"40-recursion",
	"41-euclidean-algorithm",
	"42-convex-hull",
	"43-bipartite-matching",
	"44-sieve-of-eratosthenes",
	"45-inclusion-exclusion",
	"46-lca",
	"47-sparse-table",
	"48-hashing",
	"49-modular-inverse",
	"50-floyd-warshall",
	"51-trie",
	"52-deque",
	"53-prime-factorization",
	"54-tree-dp",
	"55-lis",
	"56-sqrt-decomposition",
	"57-meet-in-the-middle",
	"58-zero-one-bfs",
	"59-flood-fill",
	"60-fft",
	"61-ternary-search",
	"62-euler-tour",
	"63-mcmf",
	"64-convex-hull-trick",
	"65-gaussian-elimination",
	"66-hld",
	"67-centroid-decomposition",
}

// 확장 알고리즘(27~53번)의 선수 학습 참조 매핑
var prerequisiteMap = map[string][]string{
	"28-combinatorics":              {"07-math-and-number-theory"},
	"30-game-theory":                {"19-dynamic-programming"},
	"31-probability":                {"07-math-and-number-theory", "19-dynamic-programming"},
	"32-bitmask-dp":                 {"19-dynamic-programming", "29-bitmask"},
	"33-maximum-flow":               {"15-graph-dfs", "16-graph-bfs", "24-graph-advanced"},
	"34-primality-test":             {"07-math-and-number-theory"},
	"35-offline-queries":            {"03-sorting", "25-segment-tree"},
	"36-exponentiation-by-squaring": {"18-divide-and-conquer", "07-math-and-number-theory"},
	"37-knapsack":                   {"19-dynamic-programming"},
	"38-dag":                        {"15-graph-dfs", "23-topological-sort", "19-dynamic-programming"},
	"39-coordinate-compression":     {"03-sorting"},
	"40-recursion":                  {"18-divide-and-conquer"},
	"41-euclidean-algorithm":        {"07-math-and-number-theory"},
	"42-convex-hull":                {"27-geometry", "03-sorting"},
	"43-bipartite-matching":         {"15-graph-dfs", "33-maximum-flow"},
	"44-sieve-of-eratosthenes":      {"07-math-and-number-theory"},
	"45-inclusion-exclusion":        {"28-combinatorics", "07-math-and-number-theory"},
	"46-lca":                        {"13-tree", "15-graph-dfs"},
	"47-sparse-table":               {"06-prefix-sum", "25-segment-tree"},
	"48-hashing":                    {"05-hash", "26-string-algorithm"},
	"49-modular-inverse":            {"07-math-and-number-theory", "36-exponentiation-by-squaring"},
	"50-floyd-warshall":             {"21-shortest-path"},
	"51-trie":                       {"13-tree", "26-string-algorithm"},
	"52-deque":                      {"04-stack-and-queue"},
	"53-prime-factorization":        {"07-math-and-number-theory", "34-primality-test"},
	"54-tree-dp":                    {"13-tree", "19-dynamic-programming"},
	"55-lis":                        {"08-binary-search", "19-dynamic-programming"},
	"56-sqrt-decomposition":         {"35-offline-queries", "25-segment-tree"},
	"57-meet-in-the-middle":         {"02-bruteforce", "03-sorting"},
	"58-zero-one-bfs":               {"16-graph-bfs", "52-deque", "21-shortest-path"},
	"59-flood-fill":                 {"15-graph-dfs", "16-graph-bfs"},
	"60-fft":                        {"07-math-and-number-theory", "18-divide-and-conquer"},
	"61-ternary-search":             {"08-binary-search"},
	"62-euler-tour":                 {"13-tree", "25-segment-tree", "15-graph-dfs"},
	"63-mcmf":                       {"33-maximum-flow", "21-shortest-path"},
	"64-convex-hull-trick":          {"19-dynamic-programming", "42-convex-hull"},
	"65-gaussian-elimination":       {"07-math-and-number-theory"},
	"66-hld":                        {"13-tree", "25-segment-tree", "15-graph-dfs"},
	"67-centroid-decomposition":     {"13-tree", "18-divide-and-conquer"},
}

// 허용된 표준 라이브러리 목록
var allowedImports = map[string]bool{
	"fmt": true, "bufio": true, "os": true, "math": true, "math/big": true,
	"sort": true, "strings": true, "strconv": true, "container/heap": true,
	"container/list": true, "io": true, "unicode": true, "bytes": true,
	"encoding/binary": true, "math/bits": true, "math/cmplx": true,
}

// rootDir은 프로젝트 루트 경로 (validate/ 의 상위 디렉토리)
var rootDir = filepath.Join("..")

// folderPath는 알고리즘 폴더의 전체 경로를 반환한다
func folderPath(folder string) string {
	return filepath.Join(rootDir, folder)
}

// fileExists는 파일 존재 여부를 확인한다
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// dirExists는 디렉토리 존재 여부를 확인한다
func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// readFileContent는 파일 내용을 읽어 반환한다
func readFileContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// listGoFiles는 디렉토리 내 .go 파일 목록을 반환한다
func listGoFiles(dir string) ([]string, error) {
	var goFiles []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") {
			goFiles = append(goFiles, filepath.Join(dir, entry.Name()))
		}
	}
	return goFiles, nil
}

// listSubDirs는 디렉토리 내 하위 디렉토리 목록을 반환한다
func listSubDirs(dir string) ([]string, error) {
	var dirs []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}
	return dirs, nil
}

// collectAllGoFiles는 알고리즘 폴더 내 모든 .go 파일을 재귀적으로 수집한다
func collectAllGoFiles(dir string) ([]string, error) {
	var goFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})
	return goFiles, err
}

// TestFolderStructureCompleteness verifies each of the 53 folders has the required structure.
// **Validates: Requirements 1.5, 2.1, 3.1, 4.1, 4.6, 5.1**
// Feature: algorithm-study-guide, Property 1: 폴더 구조 완전성
func TestFolderStructureCompleteness(t *testing.T) {
	for _, folder := range algorithmFolders {
		t.Run(folder, func(t *testing.T) {
			base := folderPath(folder)

			// README.md 존재 확인
			if !fileExists(filepath.Join(base, "README.md")) {
				t.Errorf("%s: README.md가 존재하지 않습니다", folder)
			}

			// theory.md 존재 확인
			if !fileExists(filepath.Join(base, "theory.md")) {
				t.Errorf("%s: theory.md가 존재하지 않습니다", folder)
			}

			// examples/ 디렉토리 존재 및 .go 파일 최소 1개 확인
			examplesDir := filepath.Join(base, "examples")
			if !dirExists(examplesDir) {
				t.Errorf("%s: examples/ 디렉토리가 존재하지 않습니다", folder)
			} else {
				goFiles, err := listGoFiles(examplesDir)
				if err != nil {
					t.Errorf("%s: examples/ 디렉토리 읽기 실패: %v", folder, err)
				} else if len(goFiles) < 1 {
					t.Errorf("%s: examples/ 디렉토리에 .go 파일이 없습니다", folder)
				}
			}

			// problems/ 디렉토리 존재 및 최소 3개 문제 폴더 확인
			problemsDir := filepath.Join(base, "problems")
			if !dirExists(problemsDir) {
				t.Errorf("%s: problems/ 디렉토리가 존재하지 않습니다", folder)
			} else {
				subDirs, err := listSubDirs(problemsDir)
				if err != nil {
					t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
				} else if len(subDirs) < 3 {
					t.Errorf("%s: problems/ 디렉토리에 문제 폴더가 %d개뿐입니다 (최소 3개 필요)", folder, len(subDirs))
				} else {
					// 각 문제 폴더에 problem.md, solution.go, explanation.md 확인
					for _, sub := range subDirs {
						subPath := filepath.Join(problemsDir, sub)
						if !fileExists(filepath.Join(subPath, "problem.md")) {
							t.Errorf("%s/problems/%s: problem.md가 존재하지 않습니다", folder, sub)
						}
						if !fileExists(filepath.Join(subPath, "solution.go")) {
							t.Errorf("%s/problems/%s: solution.go가 존재하지 않습니다", folder, sub)
						}
						if !fileExists(filepath.Join(subPath, "answer.go")) {
							t.Errorf("%s/problems/%s: answer.go가 존재하지 않습니다", folder, sub)
						}
						if !fileExists(filepath.Join(subPath, "explanation.md")) {
							t.Errorf("%s/problems/%s: explanation.md가 존재하지 않습니다", folder, sub)
						}
					}
				}
			}
		})
	}
}

// TestFolderNamingConvention verifies all 53 folder names match the naming pattern.
// **Validates: Requirements 1.3**
// Feature: algorithm-study-guide, Property 2: 폴더 명명 규칙
func TestFolderNamingConvention(t *testing.T) {
	pattern := regexp.MustCompile(`^[0-9]{2}-[a-z]+(-[a-z]+)*$`)
	for _, folder := range algorithmFolders {
		t.Run(folder, func(t *testing.T) {
			if !pattern.MatchString(folder) {
				t.Errorf("폴더명 '%s'이(가) 명명 규칙(두 자리 번호 + 케밥 케이스)을 만족하지 않습니다", folder)
			}
		})
	}
}

// TestTheoryDocStructure verifies all 53 theory.md files contain required sections.
// **Validates: Requirements 2.2, 2.3, 2.4, 2.5**
// Feature: algorithm-study-guide, Property 3: 이론 문서 구조 완전성
func TestTheoryDocStructure(t *testing.T) {
	requiredSections := []string{"개념", "동작 원리", "복잡도", "적합한 문제 유형"}
	for _, folder := range algorithmFolders {
		t.Run(folder, func(t *testing.T) {
			theoryPath := filepath.Join(folderPath(folder), "theory.md")
			content, err := readFileContent(theoryPath)
			if err != nil {
				t.Fatalf("%s: theory.md 읽기 실패: %v", folder, err)
			}
			for _, section := range requiredSections {
				if !strings.Contains(content, section) {
					t.Errorf("%s: theory.md에 '%s' 섹션이 없습니다", folder, section)
				}
			}
		})
	}
}

// TestProblemDocStructure verifies all problem.md files contain required sections.
// **Validates: Requirements 4.2, 4.3, 4.4, 4.5, 4.8**
// Feature: algorithm-study-guide, Property 4: 문제 파일 구조 완전성
func TestProblemDocStructure(t *testing.T) {
	requiredSections := []string{"문제 설명", "입력 형식", "출력 형식", "예제", "제약 조건"}
	difficultyMarkers := []string{"하", "중", "상"}

	for _, folder := range algorithmFolders {
		problemsDir := filepath.Join(folderPath(folder), "problems")
		subDirs, err := listSubDirs(problemsDir)
		if err != nil {
			t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			continue
		}
		for _, sub := range subDirs {
			t.Run(folder+"/"+sub, func(t *testing.T) {
				problemPath := filepath.Join(problemsDir, sub, "problem.md")
				content, err := readFileContent(problemPath)
				if err != nil {
					t.Fatalf("%s/problems/%s: problem.md 읽기 실패: %v", folder, sub, err)
				}
				for _, section := range requiredSections {
					if !strings.Contains(content, section) {
						t.Errorf("%s/problems/%s: problem.md에 '%s' 섹션이 없습니다", folder, sub, section)
					}
				}
				// 난이도 표기 확인 (하/중/상 중 하나)
				hasDifficulty := false
				for _, marker := range difficultyMarkers {
					if strings.Contains(content, marker) {
						hasDifficulty = true
						break
					}
				}
				if !hasDifficulty {
					t.Errorf("%s/problems/%s: problem.md에 난이도 표기(하/중/상)가 없습니다", folder, sub)
				}
			})
		}
	}
}

// TestExplanationDocStructure verifies all explanation.md files contain required sections.
// **Validates: Requirements 5.2, 5.3, 5.4**
// Feature: algorithm-study-guide, Property 5: 해설 문서 구조 완전성
func TestExplanationDocStructure(t *testing.T) {
	requiredSections := []string{"접근 방식", "핵심 아이디어", "복잡도 분석"}

	for _, folder := range algorithmFolders {
		problemsDir := filepath.Join(folderPath(folder), "problems")
		subDirs, err := listSubDirs(problemsDir)
		if err != nil {
			t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			continue
		}
		for _, sub := range subDirs {
			t.Run(folder+"/"+sub, func(t *testing.T) {
				explPath := filepath.Join(problemsDir, sub, "explanation.md")
				content, err := readFileContent(explPath)
				if err != nil {
					t.Fatalf("%s/problems/%s: explanation.md 읽기 실패: %v", folder, sub, err)
				}
				for _, section := range requiredSections {
					if !strings.Contains(content, section) {
						t.Errorf("%s/problems/%s: explanation.md에 '%s' 섹션이 없습니다", folder, sub, section)
					}
				}
			})
		}
	}
}

// TestDifficultyDistribution verifies each problems/ dir has at least one easy, medium, and hard folder.
// **Validates: Requirements 4.9**
// Feature: algorithm-study-guide, Property 6: 난이도 분포 균형
func TestDifficultyDistribution(t *testing.T) {
	for _, folder := range algorithmFolders {
		t.Run(folder, func(t *testing.T) {
			problemsDir := filepath.Join(folderPath(folder), "problems")
			subDirs, err := listSubDirs(problemsDir)
			if err != nil {
				t.Fatalf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			}

			hasEasy := false
			hasMedium := false
			hasHard := false
			for _, sub := range subDirs {
				if strings.Contains(sub, "easy") {
					hasEasy = true
				}
				if strings.Contains(sub, "medium") {
					hasMedium = true
				}
				if strings.Contains(sub, "hard") {
					hasHard = true
				}
			}
			if !hasEasy {
				t.Errorf("%s: problems/에 easy 난이도 폴더가 없습니다", folder)
			}
			if !hasMedium {
				t.Errorf("%s: problems/에 medium 난이도 폴더가 없습니다", folder)
			}
			if !hasHard {
				t.Errorf("%s: problems/에 hard 난이도 폴더가 없습니다", folder)
			}
		})
	}
}

// TestGoCodeConventions verifies all .go files follow code conventions.
// **Validates: Requirements 3.2, 3.4, 3.5, 7.2, 7.3**
// Feature: algorithm-study-guide, Property 7: Go 코드 규칙 준수
func TestGoCodeConventions(t *testing.T) {
	// 한국어 유니코드 범위 정규식 (가-힣)
	koreanPattern := regexp.MustCompile(`[\x{AC00}-\x{D7AF}]`)
	importPattern := regexp.MustCompile(`(?m)^import\s+\(\s*\n([\s\S]*?)\n\s*\)`)
	singleImportPattern := regexp.MustCompile(`(?m)^import\s+"([^"]+)"`)

	for _, folder := range algorithmFolders {
		base := folderPath(folder)
		goFiles, err := collectAllGoFiles(base)
		if err != nil {
			t.Errorf("%s: .go 파일 수집 실패: %v", folder, err)
			continue
		}
		for _, goFile := range goFiles {
			relPath := strings.TrimPrefix(goFile, rootDir+string(os.PathSeparator))
			t.Run(relPath, func(t *testing.T) {
				content, err := readFileContent(goFile)
				if err != nil {
					t.Fatalf("파일 읽기 실패: %v", err)
				}

				// package main 확인
				if !strings.Contains(content, "package main") {
					t.Errorf("'package main' 선언이 없습니다")
				}

				// func main() 확인
				if !strings.Contains(content, "func main()") {
					t.Errorf("'func main()' 함수가 없습니다")
				}

				// 한국어 주석 확인 (// 또는 /* */ 내에 한국어 포함)
				lines := strings.Split(content, "\n")
				hasKoreanComment := false
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if strings.HasPrefix(trimmed, "//") {
						if koreanPattern.MatchString(trimmed) {
							hasKoreanComment = true
							break
						}
					}
				}
				if !hasKoreanComment {
					t.Errorf("한국어 주석이 없습니다")
				}

				// 표준 라이브러리만 사용 확인
				matches := importPattern.FindStringSubmatch(content)
				if len(matches) > 1 {
					importBlock := matches[1]
					importLines := strings.Split(importBlock, "\n")
					for _, imp := range importLines {
						imp = strings.TrimSpace(imp)
						imp = strings.Trim(imp, "\"")
						if imp == "" {
							continue
						}
						if !allowedImports[imp] {
							t.Errorf("허용되지 않은 import: %s", imp)
						}
					}
				}
				singleMatches := singleImportPattern.FindAllStringSubmatch(content, -1)
				for _, m := range singleMatches {
					if len(m) > 1 && !allowedImports[m[1]] {
						t.Errorf("허용되지 않은 import: %s", m[1])
					}
				}
			})
		}
	}
}

// TestPrerequisiteReferences verifies extended algorithms (27-53) reference prerequisite folders in theory.md.
// **Validates: Requirements 8.6, 9.3**
// Feature: algorithm-study-guide, Property: 확장 알고리즘 선수 학습 참조
func TestPrerequisiteReferences(t *testing.T) {
	for folder, prereqs := range prerequisiteMap {
		t.Run(folder, func(t *testing.T) {
			theoryPath := filepath.Join(folderPath(folder), "theory.md")
			content, err := readFileContent(theoryPath)
			if err != nil {
				t.Fatalf("%s: theory.md 읽기 실패: %v", folder, err)
			}
			for _, prereq := range prereqs {
				if !strings.Contains(content, prereq) {
					t.Errorf("%s: theory.md에 선수 학습 참조 '%s'가 없습니다", folder, prereq)
				}
			}
		})
	}
}

// TestTheoryDocSupplementCompleteness verifies all theory.md files have detailed supplement sections.
// **Validates: Requirements 12.2, 12.4, 12.5, 12.9**
// Feature: algorithm-study-guide, Property 8: 이론 문서 상세 보충 완전성
func TestTheoryDocSupplementCompleteness(t *testing.T) {
	supplementSections := []struct {
		name     string
		patterns []string
	}{
		{"단계별 추적/Trace", []string{"단계별 추적", "Trace"}},
		{"실전 팁", []string{"실전 팁"}},
		{"관련 알고리즘 비교", []string{"관련 알고리즘 비교"}},
	}

	for _, folder := range algorithmFolders {
		t.Run(folder, func(t *testing.T) {
			theoryPath := filepath.Join(folderPath(folder), "theory.md")
			content, err := readFileContent(theoryPath)
			if err != nil {
				t.Fatalf("%s: theory.md 읽기 실패: %v", folder, err)
			}
			for _, section := range supplementSections {
				found := false
				for _, pattern := range section.patterns {
					if strings.Contains(content, pattern) {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("%s: theory.md에 '%s' 섹션이 없습니다", folder, section.name)
				}
			}
		})
	}
}

// TestProblemFolderFileCompleteness verifies all 201 problem folders contain the required 4 files.
// **Validates: Requirements 1.7, 5.1**
// Feature: hackerrank-style-refactor, Property 1: 문제 폴더 파일 완전성
func TestProblemFolderFileCompleteness(t *testing.T) {
	requiredFiles := []string{"problem.md", "solution.go", "answer.go", "explanation.md"}
	for _, folder := range algorithmFolders {
		problemsDir := filepath.Join(folderPath(folder), "problems")
		subDirs, err := listSubDirs(problemsDir)
		if err != nil {
			t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			continue
		}
		for _, sub := range subDirs {
			t.Run(folder+"/"+sub, func(t *testing.T) {
				subPath := filepath.Join(problemsDir, sub)
				for _, reqFile := range requiredFiles {
					if !fileExists(filepath.Join(subPath, reqFile)) {
						t.Errorf("%s/problems/%s: %s가 존재하지 않습니다", folder, sub, reqFile)
					}
				}
			})
		}
	}
}

// TestAnswerGoCodeConventions verifies all answer.go files follow Go code conventions.
// **Validates: Requirements 1.1, 3.5, 3.6, 3.7, 3.8**
// Feature: hackerrank-style-refactor, Property 2: answer.go Go 코드 규칙 준수
func TestAnswerGoCodeConventions(t *testing.T) {
	koreanPattern := regexp.MustCompile(`[\x{AC00}-\x{D7AF}]`)
	importPattern := regexp.MustCompile(`(?m)^import\s+\(\s*\n([\s\S]*?)\n\s*\)`)
	singleImportPattern := regexp.MustCompile(`(?m)^import\s+"([^"]+)"`)

	for _, folder := range algorithmFolders {
		problemsDir := filepath.Join(folderPath(folder), "problems")
		subDirs, err := listSubDirs(problemsDir)
		if err != nil {
			t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			continue
		}
		for _, sub := range subDirs {
			t.Run(folder+"/"+sub, func(t *testing.T) {
				answerPath := filepath.Join(problemsDir, sub, "answer.go")
				content, err := readFileContent(answerPath)
				if err != nil {
					t.Fatalf("answer.go 읽기 실패: %v", err)
				}
				if !strings.Contains(content, "package main") {
					t.Errorf("'package main' 선언이 없습니다")
				}
				if !strings.Contains(content, "func main()") {
					t.Errorf("'func main()' 함수가 없습니다")
				}
				lines := strings.Split(content, "\n")
				hasKoreanComment := false
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if strings.HasPrefix(trimmed, "//") && koreanPattern.MatchString(trimmed) {
						hasKoreanComment = true
						break
					}
				}
				if !hasKoreanComment {
					t.Errorf("한국어 주석이 없습니다")
				}
				matches := importPattern.FindStringSubmatch(content)
				if len(matches) > 1 {
					for _, imp := range strings.Split(matches[1], "\n") {
						imp = strings.TrimSpace(strings.Trim(strings.TrimSpace(imp), "\""))
						if imp != "" && !allowedImports[imp] {
							t.Errorf("허용되지 않은 import: %s", imp)
						}
					}
				}
				for _, m := range singleImportPattern.FindAllStringSubmatch(content, -1) {
					if len(m) > 1 && !allowedImports[m[1]] {
						t.Errorf("허용되지 않은 import: %s", m[1])
					}
				}
			})
		}
	}
}

// TestSolutionGoCodeConventions verifies all solution.go files in problems/ follow Go code conventions.
// **Validates: Requirements 1.3, 3.1, 3.2, 3.3, 3.4**
// Feature: hackerrank-style-refactor, Property 3: solution.go Go 코드 규칙 준수
func TestSolutionGoCodeConventions(t *testing.T) {
	koreanPattern := regexp.MustCompile(`[\x{AC00}-\x{D7AF}]`)
	importPattern := regexp.MustCompile(`(?m)^import\s+\(\s*\n([\s\S]*?)\n\s*\)`)
	singleImportPattern := regexp.MustCompile(`(?m)^import\s+"([^"]+)"`)

	for _, folder := range algorithmFolders {
		problemsDir := filepath.Join(folderPath(folder), "problems")
		subDirs, err := listSubDirs(problemsDir)
		if err != nil {
			t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			continue
		}
		for _, sub := range subDirs {
			t.Run(folder+"/"+sub, func(t *testing.T) {
				solutionPath := filepath.Join(problemsDir, sub, "solution.go")
				content, err := readFileContent(solutionPath)
				if err != nil {
					t.Fatalf("solution.go 읽기 실패: %v", err)
				}
				if !strings.Contains(content, "package main") {
					t.Errorf("'package main' 선언이 없습니다")
				}
				if !strings.Contains(content, "func main()") {
					t.Errorf("'func main()' 함수가 없습니다")
				}
				lines := strings.Split(content, "\n")
				hasKoreanComment := false
				for _, line := range lines {
					trimmed := strings.TrimSpace(line)
					if strings.HasPrefix(trimmed, "//") && koreanPattern.MatchString(trimmed) {
						hasKoreanComment = true
						break
					}
				}
				if !hasKoreanComment {
					t.Errorf("한국어 주석이 없습니다")
				}
				matches := importPattern.FindStringSubmatch(content)
				if len(matches) > 1 {
					for _, imp := range strings.Split(matches[1], "\n") {
						imp = strings.TrimSpace(strings.Trim(strings.TrimSpace(imp), "\""))
						if imp != "" && !allowedImports[imp] {
							t.Errorf("허용되지 않은 import: %s", imp)
						}
					}
				}
				for _, m := range singleImportPattern.FindAllStringSubmatch(content, -1) {
					if len(m) > 1 && !allowedImports[m[1]] {
						t.Errorf("허용되지 않은 import: %s", m[1])
					}
				}
			})
		}
	}
}

// TestSolutionGoEmptyFunctionPattern verifies all solution.go files have empty core function pattern.
// **Validates: Requirements 1.2, 1.4, 2.2**
// Feature: hackerrank-style-refactor, Property 4: solution.go 빈 핵심 함수 패턴
func TestSolutionGoEmptyFunctionPattern(t *testing.T) {
	funcPattern := regexp.MustCompile(`func\s+(\w+)\s*\(`)

	for _, folder := range algorithmFolders {
		problemsDir := filepath.Join(folderPath(folder), "problems")
		subDirs, err := listSubDirs(problemsDir)
		if err != nil {
			t.Errorf("%s: problems/ 디렉토리 읽기 실패: %v", folder, err)
			continue
		}
		for _, sub := range subDirs {
			t.Run(folder+"/"+sub, func(t *testing.T) {
				solutionPath := filepath.Join(problemsDir, sub, "solution.go")
				content, err := readFileContent(solutionPath)
				if err != nil {
					t.Fatalf("solution.go 읽기 실패: %v", err)
				}
				matches := funcPattern.FindAllStringSubmatch(content, -1)
				nonMainFuncs := 0
				for _, m := range matches {
					if len(m) > 1 && m[1] != "main" {
						nonMainFuncs++
					}
				}
				if nonMainFuncs < 1 {
					t.Errorf("main 외 함수가 없습니다 (최소 1개 필요)")
				}
				if !strings.Contains(content, "// 여기에 코드를 작성하세요") {
					t.Errorf("'// 여기에 코드를 작성하세요' 안내 주석이 없습니다")
				}
			})
		}
	}
}
