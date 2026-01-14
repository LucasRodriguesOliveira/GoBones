package pipeline

import (
	"fmt"

	"github.com/LucasRodriguesOliveira/GoBones/internal/http"
)

type HOOK_REGISTER_POSITION uint

const (
	PIPELINE_REGISTER_BEFORE HOOK_REGISTER_POSITION = 0
	PIPELINE_REGISTER_AFTER  HOOK_REGISTER_POSITION = 1
	PIPELINE_REGISTER_BOTH   HOOK_REGISTER_POSITION = 2
)

type Pipeline[T func(req *http.Request, res *http.Response) error] struct {
	before []T
	after  []T
}

func (p *Pipeline[T]) Register(
	fn func(req *http.Request, res *http.Response) error,
	pos HOOK_REGISTER_POSITION,
) {
	switch pos {
	case PIPELINE_REGISTER_BEFORE:
		p.before = append(p.before, fn)
	case PIPELINE_REGISTER_AFTER:
		p.after = append(p.after, fn)
	case PIPELINE_REGISTER_BOTH:
		p.before = append(p.before, fn)
		p.after = append(p.after, fn)
	}
}

func (p *Pipeline[T]) run(
	pipeline []T,
	req *http.Request,
	res *http.Response,
) error {
	for _, h := range pipeline {
		if err := h(req, res); err != nil {
			return err
		}
	}

	return nil
}

func (p *Pipeline[T]) Before(req *http.Request, res *http.Response) error {
	fmt.Println("Running Before Hooks")
	return p.run(p.before, req, res)
}

func (p *Pipeline[T]) After(req *http.Request, res *http.Response) error {
	fmt.Println("Running After Hooks")
	return p.run(p.after, req, res)
}
