package nson_test

import (
	"testing"
	"time"

	nson "github.com/danclive/nson-go"
)

// TestMarshalTimestamp 测试 time.Time 到 Timestamp 的序列化
func TestMarshalTimestamp(t *testing.T) {
	type TimeData struct {
		CreatedAt time.Time `nson:"created_at"`
		UpdatedAt time.Time `nson:"updated_at"`
	}

	// 创建测试时间（使用毫秒精度）
	now := time.UnixMilli(1732694400000)   // 2024-11-27 12:00:00 UTC
	later := time.UnixMilli(1732698000000) // 2024-11-27 13:00:00 UTC

	data := TimeData{
		CreatedAt: now,
		UpdatedAt: later,
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证类型和值
	createdAt, ok := m["created_at"].(nson.Timestamp)
	if !ok {
		t.Errorf("created_at should be Timestamp type, got %T", m["created_at"])
	}
	if createdAt != nson.Timestamp(1732694400000) {
		t.Errorf("expected Timestamp(1732694400000), got %v", createdAt)
	}

	updatedAt, ok := m["updated_at"].(nson.Timestamp)
	if !ok {
		t.Errorf("updated_at should be Timestamp type, got %T", m["updated_at"])
	}
	if updatedAt != nson.Timestamp(1732698000000) {
		t.Errorf("expected Timestamp(1732698000000), got %v", updatedAt)
	}
}

// TestUnmarshalTimestamp 测试 Timestamp 到 time.Time 的反序列化
func TestUnmarshalTimestamp(t *testing.T) {
	type TimeData struct {
		CreatedAt time.Time `nson:"created_at"`
		UpdatedAt time.Time `nson:"updated_at"`
	}

	m := nson.Map{
		"created_at": nson.Timestamp(1732694400000),
		"updated_at": nson.Timestamp(1732698000000),
	}

	var data TimeData
	err := nson.Unmarshal(m, &data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	expectedCreated := time.UnixMilli(1732694400000)
	expectedUpdated := time.UnixMilli(1732698000000)

	if data.CreatedAt.UnixMilli() != expectedCreated.UnixMilli() {
		t.Errorf("expected CreatedAt=%v, got %v", expectedCreated.UnixMilli(), data.CreatedAt.UnixMilli())
	}
	if data.UpdatedAt.UnixMilli() != expectedUpdated.UnixMilli() {
		t.Errorf("expected UpdatedAt=%v, got %v", expectedUpdated.UnixMilli(), data.UpdatedAt.UnixMilli())
	}
}

// TestTimestampRoundTrip 测试 time.Time 的往返转换
func TestTimestampRoundTrip(t *testing.T) {
	type TimeData struct {
		Timestamp time.Time `nson:"ts"`
	}

	original := TimeData{
		Timestamp: time.UnixMilli(1732694400000),
	}

	// Marshal
	m, err := nson.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result TimeData
	err = nson.Unmarshal(m, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// 比较毫秒时间戳（忽略纳秒部分）
	if original.Timestamp.UnixMilli() != result.Timestamp.UnixMilli() {
		t.Errorf("expected %v, got %v", original.Timestamp.UnixMilli(), result.Timestamp.UnixMilli())
	}
}

// TestMarshalId 测试 nson.Id 的序列化
func TestMarshalId(t *testing.T) {
	type UserData struct {
		ID   nson.Id `nson:"id"`
		Name string  `nson:"name"`
	}

	data := UserData{
		ID:   nson.Id{0x01, 0x02, 0x03, 0x04, 0x05},
		Name: "Alice",
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证类型和值
	id, ok := m["id"].(nson.Id)
	if !ok {
		t.Errorf("id should be nson.Id type, got %T", m["id"])
	}
	expectedID := nson.Id{0x01, 0x02, 0x03, 0x04, 0x05}
	if len(id) != len(expectedID) {
		t.Errorf("id length mismatch: expected %d, got %d", len(expectedID), len(id))
	}
	for i := range expectedID {
		if id[i] != expectedID[i] {
			t.Errorf("id[%d]: expected 0x%02X, got 0x%02X", i, expectedID[i], id[i])
		}
	}

	name, ok := m["name"].(nson.String)
	if !ok {
		t.Errorf("name should be String type, got %T", m["name"])
	}
	if name != nson.String("Alice") {
		t.Errorf("expected name=Alice, got %v", name)
	}
}

// TestUnmarshalId 测试 nson.Id 的反序列化
func TestUnmarshalId(t *testing.T) {
	type UserData struct {
		ID   nson.Id `nson:"id"`
		Name string  `nson:"name"`
	}

	m := nson.Map{
		"id":   nson.Id{0xAA, 0xBB, 0xCC, 0xDD, 0xEE},
		"name": nson.String("Bob"),
	}

	var data UserData
	err := nson.Unmarshal(m, &data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	expectedID := nson.Id{0xAA, 0xBB, 0xCC, 0xDD, 0xEE}
	if len(data.ID) != len(expectedID) {
		t.Errorf("ID length mismatch: expected %d, got %d", len(expectedID), len(data.ID))
	}
	for i := range expectedID {
		if data.ID[i] != expectedID[i] {
			t.Errorf("ID[%d]: expected 0x%02X, got 0x%02X", i, expectedID[i], data.ID[i])
		}
	}
	if data.Name != "Bob" {
		t.Errorf("expected Name=Bob, got %v", data.Name)
	}
}

// TestIdRoundTrip 测试 nson.Id 的往返转换
func TestIdRoundTrip(t *testing.T) {
	type Document struct {
		DocID nson.Id `nson:"doc_id"`
		Title string  `nson:"title"`
	}

	original := Document{
		DocID: nson.Id{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0},
		Title: "Test Document",
	}

	// Marshal
	m, err := nson.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result Document
	err = nson.Unmarshal(m, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(original.DocID) != len(result.DocID) {
		t.Errorf("DocID length mismatch: expected %d, got %d", len(original.DocID), len(result.DocID))
	}
	for i := range original.DocID {
		if original.DocID[i] != result.DocID[i] {
			t.Errorf("DocID[%d]: expected 0x%02X, got 0x%02X", i, original.DocID[i], result.DocID[i])
		}
	}
	if original.Title != result.Title {
		t.Errorf("expected Title=%v, got %v", original.Title, result.Title)
	}
}

// TestIdVsBinary 测试 nson.Id 和 []byte (Binary) 的区别
func TestIdVsBinary(t *testing.T) {
	type MixedData struct {
		ID   nson.Id `nson:"id"`
		Data []byte  `nson:"data"`
	}

	data := MixedData{
		ID:   nson.Id{0x01, 0x02, 0x03},
		Data: []byte{0x04, 0x05, 0x06},
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// ID 应该是 nson.Id 类型
	id, ok := m["id"].(nson.Id)
	if !ok {
		t.Errorf("id should be nson.Id type, got %T", m["id"])
	}
	if id.DataType() != nson.DataTypeID {
		t.Errorf("expected DataTypeID, got %v", id.DataType())
	}

	// Data 应该是 Binary 类型
	binary, ok := m["data"].(nson.Binary)
	if !ok {
		t.Errorf("data should be Binary type, got %T", m["data"])
	}
	if binary.DataType() != nson.DataTypeBINARY {
		t.Errorf("expected DataTypeBINARY, got %v", binary.DataType())
	}
}

// TestTimestampWithPointer 测试指针类型的 time.Time
func TestTimestampWithPointer(t *testing.T) {
	type OptionalTimeData struct {
		MaybeTime *time.Time `nson:"maybe_time"`
	}

	// 测试非 nil 指针
	t.Run("NonNil", func(t *testing.T) {
		now := time.UnixMilli(1732694400000)
		data := OptionalTimeData{
			MaybeTime: &now,
		}

		m, err := nson.Marshal(data)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		ts, ok := m["maybe_time"].(nson.Timestamp)
		if !ok {
			t.Errorf("maybe_time should be Timestamp type, got %T", m["maybe_time"])
		}
		if ts != nson.Timestamp(1732694400000) {
			t.Errorf("expected Timestamp(1732694400000), got %v", ts)
		}

		// Unmarshal
		var result OptionalTimeData
		err = nson.Unmarshal(m, &result)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if result.MaybeTime == nil {
			t.Fatal("MaybeTime should not be nil")
		}
		if now.UnixMilli() != result.MaybeTime.UnixMilli() {
			t.Errorf("expected %v, got %v", now.UnixMilli(), result.MaybeTime.UnixMilli())
		}
	})

	// 测试 nil 指针
	t.Run("Nil", func(t *testing.T) {
		data := OptionalTimeData{
			MaybeTime: nil,
		}

		m, err := nson.Marshal(data)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		_, ok := m["maybe_time"].(nson.Null)
		if !ok {
			t.Errorf("nil time.Time pointer should be Null, got %T", m["maybe_time"])
		}
	})
}

// TestComplexStructWithTimestampAndId 测试复杂结构
func TestComplexStructWithTimestampAndId(t *testing.T) {
	type User struct {
		ID        nson.Id   `nson:"id"`
		Name      string    `nson:"name"`
		CreatedAt time.Time `nson:"created_at"`
		UpdatedAt time.Time `nson:"updated_at"`
		Avatar    []byte    `nson:"avatar"`
	}

	original := User{
		ID:        nson.Id{0x01, 0x02, 0x03, 0x04},
		Name:      "Charlie",
		CreatedAt: time.UnixMilli(1732694400000),
		UpdatedAt: time.UnixMilli(1732698000000),
		Avatar:    []byte{0xFF, 0xFE, 0xFD},
	}

	// Marshal
	m, err := nson.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证类型
	_, ok := m["id"].(nson.Id)
	if !ok {
		t.Errorf("id should be nson.Id type, got %T", m["id"])
	}
	_, ok = m["created_at"].(nson.Timestamp)
	if !ok {
		t.Errorf("created_at should be Timestamp type, got %T", m["created_at"])
	}
	_, ok = m["updated_at"].(nson.Timestamp)
	if !ok {
		t.Errorf("updated_at should be Timestamp type, got %T", m["updated_at"])
	}
	_, ok = m["avatar"].(nson.Binary)
	if !ok {
		t.Errorf("avatar should be Binary type, got %T", m["avatar"])
	}

	// Unmarshal
	var result User
	err = nson.Unmarshal(m, &result)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(original.ID) != len(result.ID) {
		t.Errorf("ID length mismatch: expected %d, got %d", len(original.ID), len(result.ID))
	}
	for i := range original.ID {
		if original.ID[i] != result.ID[i] {
			t.Errorf("ID[%d]: expected 0x%02X, got 0x%02X", i, original.ID[i], result.ID[i])
		}
	}
	if original.Name != result.Name {
		t.Errorf("expected Name=%v, got %v", original.Name, result.Name)
	}
	if original.CreatedAt.UnixMilli() != result.CreatedAt.UnixMilli() {
		t.Errorf("expected CreatedAt=%v, got %v", original.CreatedAt.UnixMilli(), result.CreatedAt.UnixMilli())
	}
	if original.UpdatedAt.UnixMilli() != result.UpdatedAt.UnixMilli() {
		t.Errorf("expected UpdatedAt=%v, got %v", original.UpdatedAt.UnixMilli(), result.UpdatedAt.UnixMilli())
	}
	if len(original.Avatar) != len(result.Avatar) {
		t.Errorf("Avatar length mismatch: expected %d, got %d", len(original.Avatar), len(result.Avatar))
	}
	for i := range original.Avatar {
		if original.Avatar[i] != result.Avatar[i] {
			t.Errorf("Avatar[%d]: expected 0x%02X, got 0x%02X", i, original.Avatar[i], result.Avatar[i])
		}
	}
}
