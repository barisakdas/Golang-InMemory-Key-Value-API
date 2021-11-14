package Helpers

var (
	errorNotFound        = "Cache data not found: "
	errorEmptyKey        = "Key can not be null"
	errorAlreadyExistKey = "Key is already exist."
)

// Cache interface
type Cache interface {
	// Set cache'e data kaydeder
	Set(key string, data string) string
	// Get cache'ten data alır
	Get(key string) (string, string)
	// DeleteAll Tüm verileri silecek
	DeleteAll()
	// GetAll : Cache içindeki tüm datayı dönecek
	GetAll() map[string]string
}

// MemoryCache : In memory'de data saklama sınıfı
type MemoryCache struct {
	data map[string]string
}

// NewMemoryCache : Yeni bir cache oluşturmak için gereken fonksiyon
func NewMemoryCache() Cache {
	return &MemoryCache{data: make(map[string]string)}
}

// GetAll : Tüm datayı getirecek olan fonksiyon
func (c *MemoryCache) GetAll() map[string]string {
	return c.data
}

// Get : İn Memory'de tutulan datayı alır.
func (c *MemoryCache) Get(key string) (string, string) {
	if value, oke := c.data[key]; oke == true {
		return value, ""
	}
	return "", errorNotFound
}

// Set : İn Memorye datayı kaydeder
func (c *MemoryCache) Set(key string, data string) string {
	if key == "" {
		return errorEmptyKey
	}
	x, _ := c.Get(key)
	if x != "" {
		return errorAlreadyExistKey
	}
	c.data[key] = data
	return ""
}

// DeleteAll : Cache üzerindeki tüm verileri silmek için gereken fonksiyon
func (c *MemoryCache) DeleteAll() {
	for k := range c.data {
		delete(c.data, k)
	}
}
