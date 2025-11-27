package nson_test

import (
	"testing"

	nson "github.com/danclive/nson-go"
)

type Person struct {
	Name     string   `nson:"name"`
	Age      int32    `nson:"age"`
	Email    string   `nson:"email,omitempty"`
	Height   float32  `nson:"height"`
	IsActive bool     `nson:"is_active"`
	Tags     []string `nson:"tags"`
	Metadata nson.Map `nson:"metadata,omitempty"`
}

type Address struct {
	Street  string `nson:"street"`
	City    string `nson:"city"`
	ZipCode string `nson:"zip_code"`
}

type Employee struct {
	Person
	EmployeeID uint64   `nson:"employee_id"`
	Department string   `nson:"department"`
	Salary     float64  `nson:"salary"`
	Address    *Address `nson:"address,omitempty"`
}

func TestMarshalBasicTypes(t *testing.T) {
	person := Person{
		Name:     "Alice",
		Age:      30,
		Email:    "alice@example.com",
		Height:   1.65,
		IsActive: true,
		Tags:     []string{"developer", "golang"},
	}

	m, err := nson.Marshal(person)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证字段
	if name, err := m.GetString("name"); err != nil || name != "Alice" {
		t.Errorf("Expected name=Alice, got %v, err=%v", name, err)
	}

	if age, err := m.GetI32("age"); err != nil || age != 30 {
		t.Errorf("Expected age=30, got %v, err=%v", age, err)
	}

	if email, err := m.GetString("email"); err != nil || email != "alice@example.com" {
		t.Errorf("Expected email=alice@example.com, got %v, err=%v", email, err)
	}

	if isActive, err := m.GetBool("is_active"); err != nil || !isActive {
		t.Errorf("Expected is_active=true, got %v, err=%v", isActive, err)
	}

	tags, err := m.GetArray("tags")
	if err != nil {
		t.Fatalf("GetArray failed: %v", err)
	}
	if len(tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(tags))
	}
}

func TestUnmarshalBasicTypes(t *testing.T) {
	m := nson.Map{
		"name":      nson.String("Bob"),
		"age":       nson.I32(25),
		"email":     nson.String("bob@example.com"),
		"height":    nson.F32(1.75),
		"is_active": nson.Bool(true),
		"tags":      nson.Array{nson.String("designer"), nson.String("ui")},
	}

	var person Person
	if err := nson.Unmarshal(m, &person); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if person.Name != "Bob" {
		t.Errorf("Expected name=Bob, got %s", person.Name)
	}
	if person.Age != 25 {
		t.Errorf("Expected age=25, got %d", person.Age)
	}
	if person.Email != "bob@example.com" {
		t.Errorf("Expected email=bob@example.com, got %s", person.Email)
	}
	if person.Height != 1.75 {
		t.Errorf("Expected height=1.75, got %f", person.Height)
	}
	if !person.IsActive {
		t.Errorf("Expected is_active=true, got %v", person.IsActive)
	}
	if len(person.Tags) != 2 {
		t.Errorf("Expected 2 tags, got %d", len(person.Tags))
	}
}

func TestMarshalNested(t *testing.T) {
	employee := Employee{
		Person: Person{
			Name:     "Charlie",
			Age:      35,
			Height:   1.80,
			IsActive: true,
			Tags:     []string{"manager"},
		},
		EmployeeID: 12345,
		Department: "Engineering",
		Salary:     95000.50,
		Address: &Address{
			Street:  "123 Main St",
			City:    "San Francisco",
			ZipCode: "94105",
		},
	}

	m, err := nson.Marshal(employee)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// 验证嵌套字段
	if name, err := m.GetString("name"); err != nil || name != "Charlie" {
		t.Errorf("Expected name=Charlie, got %v, err=%v", name, err)
	}

	if empID, err := m.GetU64("employee_id"); err != nil || empID != 12345 {
		t.Errorf("Expected employee_id=12345, got %v, err=%v", empID, err)
	}

	addr, err := m.GetMap("address")
	if err != nil {
		t.Fatalf("GetMap failed: %v", err)
	}

	if city, err := addr.GetString("city"); err != nil || city != "San Francisco" {
		t.Errorf("Expected city=San Francisco, got %v, err=%v", city, err)
	}
}

func TestUnmarshalNested(t *testing.T) {
	m := nson.Map{
		"name":        nson.String("Diana"),
		"age":         nson.I32(28),
		"height":      nson.F32(1.68),
		"is_active":   nson.Bool(true),
		"tags":        nson.Array{nson.String("developer")},
		"employee_id": nson.U64(67890),
		"department":  nson.String("Product"),
		"salary":      nson.F64(85000.00),
		"address": nson.Map{
			"street":   nson.String("456 Oak Ave"),
			"city":     nson.String("New York"),
			"zip_code": nson.String("10001"),
		},
	}

	var employee Employee
	if err := nson.Unmarshal(m, &employee); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if employee.Name != "Diana" {
		t.Errorf("Expected name=Diana, got %s", employee.Name)
	}
	if employee.EmployeeID != 67890 {
		t.Errorf("Expected employee_id=67890, got %d", employee.EmployeeID)
	}
	if employee.Address == nil {
		t.Fatal("Expected address to be non-nil")
	}
	if employee.Address.City != "New York" {
		t.Errorf("Expected city=New York, got %s", employee.Address.City)
	}
}

func TestMarshalOmitEmpty(t *testing.T) {
	person := Person{
		Name:     "Eve",
		Age:      40,
		Height:   1.70,
		IsActive: true,
		Tags:     []string{},
		// Email is empty, should be omitted
	}

	m, err := nson.Marshal(person)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Email 应该被省略
	if m.Contains("email") {
		t.Error("Expected email to be omitted")
	}

	// Metadata 应该被省略
	if m.Contains("metadata") {
		t.Error("Expected metadata to be omitted")
	}

	// 但 name 应该存在
	if !m.Contains("name") {
		t.Error("Expected name to be present")
	}
}

func TestRoundTrip(t *testing.T) {
	original := Employee{
		Person: Person{
			Name:     "Frank",
			Age:      45,
			Email:    "frank@example.com",
			Height:   1.85,
			IsActive: true,
			Tags:     []string{"architect", "lead"},
			Metadata: nson.Map{
				"projects": nson.I32(15),
				"rating":   nson.F64(4.5),
			},
		},
		EmployeeID: 99999,
		Department: "Architecture",
		Salary:     120000.00,
		Address: &Address{
			Street:  "789 Elm St",
			City:    "Seattle",
			ZipCode: "98101",
		},
	}

	// Marshal
	m, err := nson.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result Employee
	if err := nson.Unmarshal(m, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// 验证
	if result.Name != original.Name {
		t.Errorf("Name mismatch: expected %s, got %s", original.Name, result.Name)
	}
	if result.Age != original.Age {
		t.Errorf("Age mismatch: expected %d, got %d", original.Age, result.Age)
	}
	if result.EmployeeID != original.EmployeeID {
		t.Errorf("EmployeeID mismatch: expected %d, got %d", original.EmployeeID, result.EmployeeID)
	}
	if result.Address.City != original.Address.City {
		t.Errorf("City mismatch: expected %s, got %s", original.Address.City, result.Address.City)
	}
}

func TestMarshalSlice(t *testing.T) {
	type Data struct {
		Numbers []int32   `nson:"numbers"`
		Floats  []float64 `nson:"floats"`
		Bytes   []byte    `nson:"bytes"`
	}

	data := Data{
		Numbers: []int32{1, 2, 3, 4, 5},
		Floats:  []float64{1.1, 2.2, 3.3},
		Bytes:   []byte{0x01, 0x02, 0x03},
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	numbers, err := m.GetArray("numbers")
	if err != nil || len(numbers) != 5 {
		t.Errorf("Expected 5 numbers, got %d, err=%v", len(numbers), err)
	}

	bytes, err := m.GetBinary("bytes")
	if err != nil || len(bytes) != 3 {
		t.Errorf("Expected 3 bytes, got %d, err=%v", len(bytes), err)
	}
}

func TestUnmarshalSlice(t *testing.T) {
	type Data struct {
		Numbers []int32 `nson:"numbers"`
		Bytes   []byte  `nson:"bytes"`
	}

	m := nson.Map{
		"numbers": nson.Array{nson.I32(10), nson.I32(20), nson.I32(30)},
		"bytes":   nson.Binary([]byte{0xAA, 0xBB, 0xCC}),
	}

	var data Data
	if err := nson.Unmarshal(m, &data); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(data.Numbers) != 3 || data.Numbers[0] != 10 {
		t.Errorf("Numbers mismatch: %v", data.Numbers)
	}

	if len(data.Bytes) != 3 || data.Bytes[0] != 0xAA {
		t.Errorf("Bytes mismatch: %v", data.Bytes)
	}
}

func TestMarshalMap(t *testing.T) {
	type Data struct {
		Attrs map[string]int32 `nson:"attrs"`
	}

	data := Data{
		Attrs: map[string]int32{
			"foo": 100,
			"bar": 200,
		},
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	attrs, err := m.GetMap("attrs")
	if err != nil {
		t.Fatalf("GetMap failed: %v", err)
	}

	if val, err := attrs.GetI32("foo"); err != nil || val != 100 {
		t.Errorf("Expected foo=100, got %v, err=%v", val, err)
	}
}

func TestNullPointer(t *testing.T) {
	type Data struct {
		Name    string   `nson:"name"`
		Address *Address `nson:"address"`
	}

	// 测试 nil 指针
	data := Data{
		Name:    "Test",
		Address: nil,
	}

	m, err := nson.Marshal(data)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// address 应该是 Null
	if !m.IsNull("address") {
		t.Error("Expected address to be Null")
	}

	// 反序列化
	var result Data
	if err := nson.Unmarshal(m, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Address != nil {
		t.Error("Expected address to be nil")
	}
}

// 性能测试
func BenchmarkMarshal(b *testing.B) {
	person := Person{
		Name:     "Benchmark",
		Age:      30,
		Email:    "bench@example.com",
		Height:   1.75,
		IsActive: true,
		Tags:     []string{"tag1", "tag2", "tag3"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := nson.Marshal(person)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	m := nson.Map{
		"name":      nson.String("Benchmark"),
		"age":       nson.I32(30),
		"email":     nson.String("bench@example.com"),
		"height":    nson.F32(1.75),
		"is_active": nson.Bool(true),
		"tags":      nson.Array{nson.String("tag1"), nson.String("tag2"), nson.String("tag3")},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var person Person
		err := nson.Unmarshal(m, &person)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRoundTrip(b *testing.B) {
	original := Person{
		Name:     "Benchmark",
		Age:      30,
		Email:    "bench@example.com",
		Height:   1.75,
		IsActive: true,
		Tags:     []string{"tag1", "tag2", "tag3"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, err := nson.Marshal(original)
		if err != nil {
			b.Fatal(err)
		}

		var result Person
		err = nson.Unmarshal(m, &result)
		if err != nil {
			b.Fatal(err)
		}
	}
}
