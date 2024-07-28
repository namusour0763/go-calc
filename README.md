# Go Calc

Go のテーブル駆動テスト (TDD) を体験するリポジトリ。

## 参考

[Goのテーブル駆動テストをわかりやすく書きたい](https://zenn.dev/kimuson13/articles/go_table_driven_test)

## 実行方法

```bash
takuya@DESKTOP-K4JMPSM:~/go-calc/calc$ go test -v
=== RUN   TestFee
=== RUN   TestFee/early_morning_5:00
=== PAUSE TestFee/early_morning_5:00
=== RUN   TestFee/err_2:00
=== PAUSE TestFee/err_2:00
=== RUN   TestFee/daytime_10:00
=== PAUSE TestFee/daytime_10:00
=== RUN   TestFee/midnight_20:00
=== PAUSE TestFee/midnight_20:00
=== CONT  TestFee/early_morning_5:00
=== CONT  TestFee/daytime_10:00
=== CONT  TestFee/err_2:00
=== CONT  TestFee/midnight_20:00
--- PASS: TestFee (0.00s)
    --- PASS: TestFee/early_morning_5:00 (0.00s)
    --- PASS: TestFee/daytime_10:00 (0.00s)
    --- PASS: TestFee/err_2:00 (0.00s)
    --- PASS: TestFee/midnight_20:00 (0.00s)
PASS
ok      calc/calc       0.002s
```
