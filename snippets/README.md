# Snippets

## Mesasure duration

To measure how long it takes for a function f to execute:

```golang
func f() {
	defer func(start time.Time) {
		statsd.Distribution("duration_s", time.Since(start).Seconds())
		...
	}(time.Now())
    ...
}
```
