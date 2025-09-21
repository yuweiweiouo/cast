package cast

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// cast 套件提供了各種類型轉換的功能
// 以下示範如何使用 cast 套件進行不同類型間的轉換

// 測試案例使用的基本結構
type A struct {
	Name string
}

type B struct {
	Title string
}

// 轉換函數: 從類型 A 轉換到類型 B
func AToB(ca Carrier[A]) (cb Carrier[B]) {
	a := ca.Get()
	return Carry(&B{
		Title: a.Name,
	})
}

// ===== 單一對象轉換測試 =====

// TestFromValueToValue 測試從值到值的轉換
// 情境: 輸入 A 結構體，輸出 B 結構體
func TestFromValueToValue(t *testing.T) {
	// 準備測試資料
	a := A{Name: "example"}

	// 執行轉換
	result := Use(AToB).FromValue(a).ToValue()

	// 驗證結果
	expected := B{Title: a.Name}
	assert.Equal(t, expected, result, "應該正確將 A.Name 轉換為 B.Title")
}

// TestFromPointerToValue 測試從指針到值的轉換
// 情境: 輸入 A 結構體指針，輸出 B 結構體
func TestFromPointerToValue(t *testing.T) {
	// 準備測試資料
	a := A{Name: "example"}

	// 執行轉換
	result := Use(AToB).FromPointer(&a).ToValue()

	// 驗證結果
	expected := B{Title: a.Name}
	assert.Equal(t, expected, result, "應該正確從 A 指針轉換到 B 值")
}

// TestFromValueToPointer 測試從值到指針的轉換
// 情境: 輸入 A 結構體，輸出 B 結構體指針
func TestFromValueToPointer(t *testing.T) {
	// 準備測試資料
	a := A{Name: "example"}

	// 執行轉換
	result := Use(AToB).FromValue(a).ToPointer()

	// 驗證結果
	expected := &B{Title: a.Name}
	assert.Equal(t, expected, result, "應該正確從 A 值轉換到 B 指針")
}

// TestFromPointerToPointer 測試從指針到指針的轉換
// 情境: 輸入 A 結構體指針，輸出 B 結構體指針
func TestFromPointerToPointer(t *testing.T) {
	// 準備測試資料
	a := A{Name: "example"}

	// 執行轉換
	result := Use(AToB).FromPointer(&a).ToPointer()

	// 驗證結果
	expected := &B{Title: a.Name}
	assert.Equal(t, expected, result, "應該正確從 A 指針轉換到 B 指針")
}

// ===== 切片對象轉換測試 =====

// TestFromValuesSliceToValuesSlice 測試從值切片到值切片的轉換
// 情境: 輸入 A 結構體切片，輸出 B 結構體切片
func TestFromValuesSliceToValuesSlice(t *testing.T) {
	// 準備測試資料
	as := []A{
		{Name: "example1"},
		{Name: "example2"},
		{Name: "example3"},
	}

	// 執行轉換
	result := Use(AToB).FromValues(as).ToValues()

	// 驗證結果
	expected := []B{
		{Title: as[0].Name},
		{Title: as[1].Name},
		{Title: as[2].Name},
	}
	assert.Equal(t, expected, result, "應該正確將 A 值切片轉換為 B 值切片")
}

// TestFromPointersSliceToValuesSlice 測試從指針切片到值切片的轉換
// 情境: 輸入 A 結構體指針切片，輸出 B 結構體切片
func TestFromPointersSliceToValuesSlice(t *testing.T) {
	// 準備測試資料
	as := []*A{
		{Name: "example1"},
		{Name: "example2"},
		{Name: "example3"},
	}

	// 執行轉換
	result := Use(AToB).FromPointers(as).ToValues()

	// 驗證結果
	expected := []B{
		{Title: as[0].Name},
		{Title: as[1].Name},
		{Title: as[2].Name},
	}
	assert.Equal(t, expected, result, "應該正確將 A 指針切片轉換為 B 值切片")
}

// TestFromValuesSliceToPointersSlice 測試從值切片到指針切片的轉換
// 情境: 輸入 A 結構體切片，輸出 B 結構體指針切片
func TestFromValuesSliceToPointersSlice(t *testing.T) {
	// 準備測試資料
	as := []A{
		{Name: "example1"},
		{Name: "example2"},
		{Name: "example3"},
	}

	// 執行轉換
	result := Use(AToB).FromValues(as).ToPointers()

	// 驗證結果
	expected := []*B{
		{Title: as[0].Name},
		{Title: as[1].Name},
		{Title: as[2].Name},
	}
	assert.Equal(t, expected, result, "應該正確將 A 值切片轉換為 B 指針切片")
}

// TestFromPointersSliceToPointersSlice 測試從指針切片到指針切片的轉換
// 情境: 輸入 A 結構體指針切片，輸出 B 結構體指針切片
func TestFromPointersSliceToPointersSlice(t *testing.T) {
	// 準備測試資料
	as := []*A{
		{Name: "example1"},
		{Name: "example2"},
		{Name: "example3"},
	}

	// 執行轉換
	result := Use(AToB).FromPointers(as).ToPointers()

	// 驗證結果
	expected := []*B{
		{Title: as[0].Name},
		{Title: as[1].Name},
		{Title: as[2].Name},
	}
	assert.Equal(t, expected, result, "應該正確將 A 指針切片轉換為 B 指針切片")
}

type AToBCaster struct {
	logger *log.Logger
}

func (c AToBCaster) Cast(ca Carrier[A]) (cb Carrier[B]) {
	c.logger.Println("Casting A to B By AToBCaster")
	return AToB(ca)
}

func TestUsingCasterInsteadOfCastFunc(t *testing.T) {
	// 準備測試資料
	caster := AToBCaster{
		logger: log.Default(),
	}
	a := A{Name: "example"}

	// 執行轉換
	result := UseCaster(caster).FromPointer(&a).ToValue()

	// 驗證結果
	expected := B{Title: a.Name}
	assert.Equal(t, expected, result, "應該正確從 A 指針轉換到 B 值")
}
