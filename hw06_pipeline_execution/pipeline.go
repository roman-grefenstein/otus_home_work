package hw06_pipeline_execution //nolint:golint,stylecheck

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil {
		out := make(Bi)
		defer close(out)
		return out
	}

	for _, stage := range stages {
		out := make(Bi)
		go func(in In, out Bi) {
			defer close(out)
			for {
				select {
				case a, ok := <-in:
					if !ok {
						return
					}
					out <- a
				case <-done:
					return
				}
			}
		}(in, out)
		in = stage(out)
	}

	return in
}
