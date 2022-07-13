package builder

import (
	"github.com/evertotomalok/text-game/internal/application/shared"
	"github.com/evertotomalok/text-game/internal/core/domain"
)

func CreateChannels(numTask int) (chan uint, chan domain.Question) {
	tasks := make(chan uint, numTask)
	results := make(chan domain.Question, numTask)

	return tasks, results
}

func Worker(tasks <-chan uint, results chan<- domain.Question) {
	for i := range tasks {
		results <- shared.MainQuestions[i]
	}

}

func CreateWorkers(tasks <-chan uint, results chan<- domain.Question, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go Worker(tasks, results)
	}
}

func PopulateTasks(tasks chan uint, nums []uint) {
	for _, num := range nums {
		tasks <- num
	}
	close(tasks)
}

func QuestionsWorkerPoolInitiate(numTasks int, randomNumbers []uint) <-chan domain.Question {
	tasks, results := CreateChannels(numTasks)

	workersNum := 8
	if workersNum > numTasks {
		workersNum = numTasks
	}

	CreateWorkers(tasks, results, workersNum)
	PopulateTasks(tasks, randomNumbers)
	return results
}
