package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(id int, ctx context.Context, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case info, ok := <-ch:
			if !ok {
				fmt.Printf("Worker %d: channel closed, exiting...\n", id)
				break
			}
			fmt.Printf("Worker %d: received %d\n", id, info)
		case <-ctx.Done():
			fmt.Printf("Worker %d: context done, exiting...\n", id)
			break
		}
	}
}

func main() {
	var n int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&n)

	ch := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Запуск воркеров
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i, ctx, ch, &wg)
	}

	// Обработка сигнала завершения
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		fmt.Println("Received shutdown signal, shutting down...")
		cancel()  // Отмена контекста
		close(ch) // Закрытие канала
	}()

	// Постоянная запись данных в канал
	for {
		x := rand.Intn(100)                // Генерация случайного числа
		ch <- x                            // Отправка числа в канал
		time.Sleep(200 * time.Millisecond) // Задержка
	}

	// Ожидание завершения всех воркеров
	wg.Wait()
	fmt.Println("All workers have finished.")
}
