package streams

import "context"

func StreamWithMapping[T any, V any](currentOps *CompiledQueryable[T], mapper func(items T) V, ctx context.Context) <-chan V {

	channel := make(chan V)

	go func() {
		defer close(channel)

		for _, v := range *currentOps.Items {

			keep := true

			select {
			case <-ctx.Done():
				keep = false
				return
			default:
			}
			for _, op := range currentOps.Operators {

				select {
				case <-ctx.Done():
					keep = false
					return
				default:
				}

				switch op.OperatorType {

				case FilterQueryable:

					if !op.MetaData.Function(v) {
						keep = false
						break
					}

				case ThrottleQueryable:

					op.MetaData.Function(v)
				}

				if !keep {
					break
				}
			}

			if keep {

				select {
				case <-ctx.Done():
					keep = false
					break

				case channel <- mapper(v):
				}
			}
		}
	}()

	return channel
}

func (currentOps *CompiledQueryable[T]) Stream(ctx context.Context) <-chan T {

	channel := make(chan T)

	go func() {
		defer close(channel)

		for _, v := range *currentOps.Items {

			keep := true

			select {
			case <-ctx.Done():
				keep = false
				return
			default:
			}
			for _, op := range currentOps.Operators {

				select {
				case <-ctx.Done():
					keep = false
					return
				default:
				}

				switch op.OperatorType {

				case FilterQueryable:

					if !op.MetaData.Function(v) {
						keep = false
						break
					}

				case ThrottleQueryable:

					op.MetaData.Function(v)
				}

				if !keep {
					break
				}
			}

			if keep {

				select {
				case <-ctx.Done():
					keep = false
					break

				case channel <- v:
				}
			}
		}
	}()

	return channel
}
