# Gilded Rose Refactoring Kata — Extended (Go)

## Требования

[Требования к чекам](GildedRose_receipt.md)

## Запуск

Требуется Go 1.18+.

```bash
cd go/04_gilded-rose-extended
go test ./...
go run texttest_fixture.go
```

Начните с `Receipt`, затем `UpdateReceipts`, затем `GenerateReceiptReport`. В стартовом коде время жёстко читается через `time.Now()`; создание seam для часов является частью упражнения. `Test_Foo` намеренно красный.
