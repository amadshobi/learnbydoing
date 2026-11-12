package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"learn-cli/content"
	"learn-cli/internal/compiler"
	"learn-cli/internal/model"
	"learn-cli/internal/progress"
	"learn-cli/internal/runner"
	"learn-cli/internal/scaffold"
	"learn-cli/internal/validator"
)

func PrintHelp(w io.Writer) {
	helpText := `
█░░ █▀▀ ▄▀█ █▀█ █▄░█                  ▄█▄
█▄▄ ██▄ █▀█ █▀▄ █░▀█                ▄█████▄     │
                                  ▄█████████▄   │
──────── STAGED LEARNING ──────── ▀█████████▀  (o)
                                    ███████

PENGGUNAAN:
  learn [perintah] [opsi]

PERINTAH YANG TERSEDIA:
  stage list         Melihat daftar semua stage tantangan yang tersedia
                     Contoh: learn stage list
                             learn stage list --easy

  init <stage>       Membuat folder dan materi untuk stage tantangan
                     Contoh: learn init 1
                             learn init 1 --easy
                             learn init 1 --force

  test               Menguji kode program kamu dengan sample test case
                     (Jalankan di dalam folder stage, misal: easy-stage-1)

  submit             Menguji seluruh test case (termasuk kasus rahasia).
                     Jika lolos, progres disimpan dan membuka stage berikutnya.

  status             Melihat progres tantangan yang sudah kamu selesaikan

  help, -h, --help   Menampilkan panduan ini

TIPS PENGGUNAAN:
  1. Mulai dengan melihat daftar stage: learn stage list
  2. Buka stage pertama               : learn init 1
  3. Masuk ke folder stage            : cd learn-c/easy-stage-1
  4. Baca README.md dan kerjakan kodenya di file solusi
  5. Uji kodemu kapan saja            : learn test
  6. Selesai? Kumpulkan hasil         : learn submit

CATATAN:
  Flag '--force' pada 'learn init <stage> --force' berfungsi sebagai tombol RESET aman
  untuk mengembalikan kode ke template awal bersih jika kodemu error parah dan ingin
  mengulang dari nol (bukan aksi berbahaya/merusak sistem).
`
	fmt.Fprint(w, helpText)
}

func ExecuteArgs(args []string, out io.Writer) error {
	if len(args) == 0 {
		PrintHelp(out)
		return nil
	}

	command := args[0]
	switch command {
	case "-h", "--help", "help":
		PrintHelp(out)
		return nil
	case "stage":
		if len(args) > 1 && args[1] == "list" {
			return runStageList(args[2:], out)
		}
		return runStageList(args[1:], out)
	case "list":
		return runStageList(args[1:], out)
	case "init":
		return runInit(args[1:], out)
	case "test":
		return runTest(false, out)
	case "submit":
		return runTest(true, out)
	case "status":
		return runStatus(out)
	default:
		fmt.Fprintf(out, "[✗] Perintah %q tidak dikenali. Ketik 'learn --help' untuk melihat bantuan.\n", command)
		return nil
	}
}

func runStageList(args []string, out io.Writer) error {
	track := "easy"
	for _, a := range args {
		if a == "--easy" {
			track = "easy"
		}
	}

	stages, err := content.ListStages(track)
	if err != nil {
		fmt.Fprintf(out, "[✗] Gagal mengambil daftar stage untuk track %q: %v\n", track, err)
		return nil
	}

	fmt.Fprintf(out, "\nDAFTAR STAGE TANTANGAN [%s]:\n", strings.ToUpper(track))
	for _, s := range stages {
		fmt.Fprintf(out, "  %d. %-23s  jalankan : learn init %d --%s\n", s.Stage, s.Title, s.Stage, s.Track)
	}
	fmt.Fprintln(out)
	return nil
}

func runInit(args []string, out io.Writer) error {
	var forceFlag bool
	var easyFlag bool = true
	var stageArg string

	for _, a := range args {
		switch a {
		case "--force", "-f", "-force":
			forceFlag = true
		case "--easy":
			easyFlag = true
		default:
			if !strings.HasPrefix(a, "-") && stageArg == "" {
				stageArg = a
			}
		}
	}

	if stageArg == "" {
		fmt.Fprintln(out, "[✗] Harap tentukan nomor stage yang ingin dibuka. Contoh: learn init 1")
		return nil
	}

	stageNum, err := strconv.Atoi(stageArg)
	if err != nil {
		fmt.Fprintf(out, "[✗] Nomor stage %q tidak valid. Harus berupa angka.\n", stageArg)
		return nil
	}

	track := "easy"
	_ = easyFlag

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	res, err := scaffold.InitStage(cwd, track, stageNum, forceFlag)
	if err != nil {
		fmt.Fprintf(out, "[✗] %s\n", err.Error())
		return nil
	}

	fmt.Fprintf(out, "\n🎉 [✓] Stage %d (%s) berhasil dibuat!\n", stageNum, res.Title)
	fmt.Fprintf(out, "📁 Lokasi: %s\n\n", res.StageDir)
	fmt.Fprintln(out, "Langkah berikutnya:")
	fmt.Fprintf(out, "  1. Masuk ke folder stage : cd %s\n", filepath.Join("learn-c", res.FolderName))
	fmt.Fprintln(out, "  2. Buka dan baca panduan  : README.md")
	fmt.Fprintln(out, "  3. Tulis kodemu di        : main.c")
	fmt.Fprintln(out, "  4. Uji tokomu kapan saja : learn test")
	fmt.Fprintln(out, "  5. Submit hasil akhir    : learn submit")
	fmt.Fprintln(out)

	return nil
}

func findStageMetadata(startDir string) (*model.StageMeta, string, error) {
	curr := startDir
	for i := 0; i < 4; i++ {
		metaFile := filepath.Join(curr, scaffold.StageMetaFileName)
		if data, err := os.ReadFile(metaFile); err == nil {
			var meta model.StageMeta
			if err := json.Unmarshal(data, &meta); err == nil {
				return &meta, curr, nil
			}
		}
		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}
	return nil, "", fmt.Errorf("tidak berada di dalam folder stage challenge")
}

func runTest(isSubmit bool, out io.Writer) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	meta, stageDir, err := findStageMetadata(cwd)
	if err != nil {
		fmt.Fprintln(out, "[✗] Kamu saat ini tidak berada di dalam folder stage tantangan.")
		fmt.Fprintln(out, "    Silakan 'cd' masuk ke dalam folder stage terlebih dahulu.")
		fmt.Fprintln(out, "    Contoh: cd learn-c/easy-stage-1")
		return nil
	}

	// Find source file (main.c or main.cpp)
	sourceFile := filepath.Join(stageDir, "main.c")
	if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
		sourceFile = filepath.Join(stageDir, "main.cpp")
		if _, err := os.Stat(sourceFile); os.IsNotExist(err) {
			fmt.Fprintln(out, "[✗] File 'main.c' atau 'main.cpp' tidak ditemukan di folder ini!")
			return nil
		}
	}

	actionTitle := "UJI COBA (learn test)"
	if isSubmit {
		actionTitle = "FINAL VALIDATION & SUBMISSION (learn submit)"
	}

	fmt.Fprintf(out, "\n🚀 Menjalankan %s untuk Stage %d: %s\n", actionTitle, meta.Stage, meta.Title)
	fmt.Fprintf(out, "⚙️  Mengompilasi %s...\n", filepath.Base(sourceFile))

	tempDir, err := os.MkdirTemp("", "learn-build-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	binPath := filepath.Join(tempDir, "solution.exe")
	compileWarns, err := compiler.Compile(sourceFile, binPath)
	if err != nil {
		fmt.Fprintln(out, "\n[✗] GAGAL KOMPILASI!")
		fmt.Fprintln(out, "Pesan dari compiler:")
		fmt.Fprintln(out, "-------------------------------------------------------------")
		fmt.Fprintln(out, compileWarns)
		fmt.Fprintln(out, "-------------------------------------------------------------")
		fmt.Fprintln(out, "💡 Tips: Periksa kembali titik koma (;), kurung tutup kurawal (}), atau typo penulisan variabel.")
		return nil
	}

	stageContent, err := content.GetStageContent(meta.Track, meta.Stage)
	if err != nil {
		fmt.Fprintf(out, "[✗] Gagal memuat test case: %v\n", err)
		return nil
	}

	testsToRun := stageContent.Tests
	if !isSubmit {
		var sampleTests []model.TestCase
		for _, t := range stageContent.Tests {
			if !t.Hidden {
				sampleTests = append(sampleTests, t)
			}
		}
		testsToRun = sampleTests
	}

	fmt.Fprintf(out, "🧪 Menjalankan %d test case...\n\n", len(testsToRun))

	allPassed := true
	for i, tc := range testsToRun {
		res, err := runner.RunBinary(binPath, tc.Stdin, 3*time.Second)
		if err != nil {
			fmt.Fprintf(out, "[✗] Error sistem saat menjalankan binary: %v\n", err)
			allPassed = false
			break
		}

		eval := validator.Evaluate(tc, res)
		if eval.Passed {
			fmt.Fprintf(out, "  [✓] Test %d: %s -> LULUS\n", i+1, tc.Name)
		} else {
			allPassed = false
			fmt.Fprintf(out, "  [✗] Test %d: %s -> GAGAL\n", i+1, tc.Name)
			if eval.ErrorMsg != "" {
				fmt.Fprintf(out, "      %s\n", eval.ErrorMsg)
			}
			if eval.Diff != "" {
				fmt.Fprintln(out, "-------------------------------------------------------------")
				fmt.Fprintln(out, eval.Diff)
				fmt.Fprintln(out, "-------------------------------------------------------------")
			}
			if !isSubmit {
				break
			}
		}
	}

	fmt.Fprintln(out, "")
	if allPassed {
		if !isSubmit {
			fmt.Fprintln(out, "✨ Keren! Semua sample test berhasil dilewati!")
			fmt.Fprintln(out, "👉 Jalankan 'learn submit' untuk menguji test case rahasia dan mencatat progresmu!")
			fmt.Fprintln(out)
		} else {
			rootDir := filepath.Dir(stageDir)
			_ = progress.SaveProgress(rootDir, meta.Track, meta.Stage)

			nextStage := meta.Stage + 1
			fmt.Fprintln(out, "=============================================================")
			fmt.Fprintf(out, "🎉 SELAMAT! Kamu berhasil menyelesaikan Stage %d: %s!\n", meta.Stage, meta.Title)
			fmt.Fprintln(out, "=============================================================")
			fmt.Fprintln(out, "Progres kamu sudah disimpan.")
			fmt.Fprintf(out, "\nLanjut ke tantangan berikutnya:\n")
			fmt.Fprintf(out, "  learn init %d --easy\n\n", nextStage)
		}
	} else {
		fmt.Fprintln(out, "❌ Masih ada pengujian yang belum cocok. Cek kembali logika kodemu lalu coba lagi!")
		fmt.Fprintln(out)
	}

	return nil
}

func runStatus(out io.Writer) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	p, err := progress.LoadProgress(cwd)
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "\n📊 STATUS PROGRES BELAJAR:")
	fmt.Fprintln(out, "-------------------------------------------------------------")
	if len(p.Completed) == 0 {
		fmt.Fprintln(out, "Belum ada stage yang diselesaikan. Ayo mulai dengan 'learn init 1'!")
	} else {
		for key := range p.Completed {
			fmt.Fprintf(out, "  [✓] %s (Selesai)\n", key)
		}
	}
	fmt.Fprintln(out, "-------------------------------------------------------------")
	fmt.Fprintln(out)
	return nil
}
