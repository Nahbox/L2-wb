package strategypattern

import (
	"fmt"
)

// Интерфейс стратегии
type CompressionStrategy interface {
	CompressFile(fileName string)
}

// Конкретная стратегия - сжатие файлов с использованием gzip
type GzipCompressionStrategy struct{}

func (g *GzipCompressionStrategy) CompressFile(fileName string) {
	fmt.Printf("File %s compressed using gzip\n", fileName)
}

// Конкретная стратегия - сжатие файлов с использованием zip
type ZipCompressionStrategy struct{}

func (z *ZipCompressionStrategy) CompressFile(fileName string) {
	fmt.Printf("File %s compressed using zip\n", fileName)
}

// Контекст, использующий стратегию
type CompressionContext struct {
	strategy CompressionStrategy
}

func (c *CompressionContext) SetCompressionStrategy(strategy CompressionStrategy) {
	c.strategy = strategy
}

func (c *CompressionContext) Compress(fileName string) {
	c.strategy.CompressFile(fileName)
}

// StrategyPattern точка входа
func StrategyPattern() {
	// Создаем контекст с начальной стратегией (gzip)
	context := &CompressionContext{strategy: &GzipCompressionStrategy{}}

	// Сжимаем файл с текущей стратегией
	context.Compress("example.txt")

	// Переключаем стратегию на zip и сжимаем другой файл
	context.SetCompressionStrategy(&ZipCompressionStrategy{})
	context.Compress("example2.txt")
}
