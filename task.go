/*
 * This file is part of the dupman/celery project.
 *
 * (c) 2022. dupman <info@dupman.cloud>
 *
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 *
 * Written by Temuri Takalandze <me@abgeo.dev>
 */

package celery

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocelery/gocelery"
)

type TaskInterface interface {
	GetNamespace() string
	Register()
}

type Task struct {
	Client    *gocelery.CeleryClient
	Namespace string
	Functions map[string]interface{}
}

func (t *Task) GetNamespace() string {
	return strings.ToLower(t.Namespace)
}

func (t *Task) Register() {
	for name, function := range t.Functions {
		taskName := fmt.Sprintf("dupman.%s.%s", t.GetNamespace(), name)
		t.Client.Register(taskName, function)
		log.Printf("Task %s registered", taskName)
	}
}

func RegisterTasks(celery *gocelery.CeleryClient, constructors ...func(celery *gocelery.CeleryClient) TaskInterface) {
	for _, constructor := range constructors {
		constructor(celery).Register()
	}
}
