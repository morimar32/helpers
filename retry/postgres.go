package retry

import (
	pq "github.com/lib/pq"
)

func WithPostgresSupport() DbRetryOption {
	return func(o *DbRetry) {
		o.evalError = func(err error) bool {
			if postgreserr, ok := err.(*pq.Error); ok {
				switch postgreserr.Code.Name() {
				case "deadlock_detected":
				case "too_many_connections":
				case "query_canceled":
					return true
				default:
					return false
				}
			}
			return false
		}
	}
}
