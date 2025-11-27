package main

import (
	"fmt"
	"log"
	"time"

	"github.com/danclive/nson-go"
)

// 用户信息
type User struct {
	ID        nson.Id  `nson:"id"`
	Name      string   `nson:"name"`
	Age       int32    `nson:"age"`
	Email     string   `nson:"email,omitempty"`
	CreatedAt uint64   `nson:"created_at"`
	UpdatedAt uint64   `nson:"updated_at,omitempty"`
	Profile   *Profile `nson:"profile,omitempty"`
	Tags      []string `nson:"tags"`
	Metadata  nson.Map `nson:"metadata,omitempty"`
}

// 用户资料
type Profile struct {
	Avatar   string            `nson:"avatar"`
	Bio      string            `nson:"bio,omitempty"`
	Location string            `nson:"location,omitempty"`
	Website  string            `nson:"website,omitempty"`
	Social   map[string]string `nson:"social,omitempty"`
	Stats    Stats             `nson:"stats"`
}

// 统计信息
type Stats struct {
	Followers int32 `nson:"followers"`
	Following int32 `nson:"following"`
	Posts     int32 `nson:"posts"`
}

func main() {
	fmt.Println("=== NSON Marshal/Unmarshal Demo ===")

	// 1. 创建用户对象
	user := User{
		ID:        nson.NewId(),
		Name:      "Alice Chen",
		Age:       28,
		Email:     "alice@example.com",
		CreatedAt: uint64(time.Now().Unix()),
		Tags:      []string{"developer", "golang", "open-source"},
		Profile: &Profile{
			Avatar:   "https://example.com/avatar.jpg",
			Bio:      "Full-stack developer, Go enthusiast",
			Location: "San Francisco, CA",
			Website:  "https://alice.dev",
			Social: map[string]string{
				"github":  "alice",
				"twitter": "@alice",
			},
			Stats: Stats{
				Followers: 1250,
				Following: 380,
				Posts:     156,
			},
		},
		Metadata: nson.Map{
			"verified":    nson.Bool(true),
			"rating":      nson.F64(4.8),
			"last_active": nson.Timestamp(uint64(time.Now().Unix())),
		},
	}

	fmt.Printf("Original User:\n")
	fmt.Printf("  ID: %s\n", user.ID.Hex())
	fmt.Printf("  Name: %s\n", user.Name)
	fmt.Printf("  Email: %s\n", user.Email)
	fmt.Printf("  Tags: %v\n", user.Tags)
	if user.Profile != nil {
		fmt.Printf("  Followers: %d\n", user.Profile.Stats.Followers)
	}
	fmt.Println()

	// 2. 序列化为 nson.Map
	m, err := nson.Marshal(user)
	if err != nil {
		log.Fatalf("Marshal failed: %v", err)
	}

	fmt.Println("Marshaled to nson.Map:")
	fmt.Printf("  Map contains %d fields\n", m.Len())

	// 使用强类型方法访问
	name, _ := m.GetString("name")
	age, _ := m.GetI32("age")
	tags, _ := m.GetArray("tags")
	profile, _ := m.GetMap("profile")

	fmt.Printf("  Name: %s (type: String)\n", name)
	fmt.Printf("  Age: %d (type: I32)\n", age)
	fmt.Printf("  Tags: %d items (type: Array)\n", len(tags))
	fmt.Printf("  Profile: %d fields (type: Map)\n", profile.Len())

	if stats, err := profile.GetMap("stats"); err == nil {
		followers, _ := stats.GetI32("followers")
		fmt.Printf("  Followers: %d (type: I32)\n", followers)
	}

	if metadata, err := m.GetMap("metadata"); err == nil {
		verified, _ := metadata.GetBool("verified")
		rating, _ := metadata.GetF64("rating")
		fmt.Printf("  Verified: %v (type: Bool)\n", verified)
		fmt.Printf("  Rating: %.1f (type: F64)\n", rating)
	}
	fmt.Println()

	// 3. 反序列化回结构体
	var user2 User
	if err := nson.Unmarshal(m, &user2); err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}

	fmt.Println("Unmarshaled back to struct:")
	fmt.Printf("  ID: %s\n", user2.ID.Hex())
	fmt.Printf("  Name: %s\n", user2.Name)
	fmt.Printf("  Email: %s\n", user2.Email)
	fmt.Printf("  Tags: %v\n", user2.Tags)
	if user2.Profile != nil {
		fmt.Printf("  Avatar: %s\n", user2.Profile.Avatar)
		fmt.Printf("  Location: %s\n", user2.Profile.Location)
		fmt.Printf("  Followers: %d\n", user2.Profile.Stats.Followers)
		fmt.Printf("  Following: %d\n", user2.Profile.Stats.Following)
		if len(user2.Profile.Social) > 0 {
			fmt.Printf("  Social: %v\n", user2.Profile.Social)
		}
	}
	fmt.Println()

	// 4. 演示 omitempty
	fmt.Println("=== Demonstrating omitempty ===")
	userMinimal := User{
		ID:        nson.NewId(),
		Name:      "Bob",
		Age:       25,
		CreatedAt: uint64(time.Now().Unix()),
		Tags:      []string{},
		// Email, UpdatedAt, Profile, Metadata 为空值
	}

	mMinimal, _ := nson.Marshal(userMinimal)
	fmt.Printf("Minimal user has %d fields (empty fields omitted)\n", mMinimal.Len())
	fmt.Printf("  Has 'email'? %v\n", mMinimal.Contains("email"))
	fmt.Printf("  Has 'profile'? %v\n", mMinimal.Contains("profile"))
	fmt.Printf("  Has 'metadata'? %v\n", mMinimal.Contains("metadata"))
	fmt.Println()

	// 5. 演示嵌入结构体
	fmt.Println("=== Demonstrating Embedded Structs ===")

	type BaseEntity struct {
		ID        nson.Id `nson:"id"`
		CreatedAt uint64  `nson:"created_at"`
		UpdatedAt uint64  `nson:"updated_at,omitempty"`
	}

	type Article struct {
		BaseEntity          // 匿名嵌入，字段会展开到顶层
		Title      string   `nson:"title"`
		Content    string   `nson:"content"`
		Author     string   `nson:"author"`
		Tags       []string `nson:"tags,omitempty"`
	}

	article := Article{
		BaseEntity: BaseEntity{
			ID:        nson.NewId(),
			CreatedAt: uint64(time.Now().Unix()),
		},
		Title:   "Getting Started with NSON",
		Content: "NSON is a binary serialization format...",
		Author:  "Alice",
		Tags:    []string{"tutorial", "nson", "golang"},
	}

	mArticle, _ := nson.Marshal(article)
	fmt.Printf("Article map has %d fields\n", mArticle.Len())

	// 嵌入的字段在顶层
	id, _ := mArticle.GetMapId("id")
	title, _ := mArticle.GetString("title")
	createdAt, _ := mArticle.GetU64("created_at")

	fmt.Printf("  ID: %s (from embedded BaseEntity)\n", id.Hex())
	fmt.Printf("  Title: %s\n", title)
	fmt.Printf("  CreatedAt: %d (from embedded BaseEntity)\n", createdAt)
	fmt.Println()

	// 6. 性能对比示例
	fmt.Println("=== Performance Example ===")

	// 多次序列化以展示缓存的优势
	iterations := 1000
	start := time.Now()

	for i := 0; i < iterations; i++ {
		nson.Marshal(user)
	}

	elapsed := time.Since(start)
	fmt.Printf("Marshaled %d times in %v\n", iterations, elapsed)
	fmt.Printf("Average: %v per operation\n", elapsed/time.Duration(iterations))
	fmt.Printf("Note: First call builds cache, subsequent calls are faster\n")
	fmt.Println()

	fmt.Println("=== Demo Complete ===")
}
