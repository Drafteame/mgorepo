package env

import (
	"os"
	"strconv"
)

var defaults map[string]any

func init() {
	defaults = make(map[string]any)
}

func GetString(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return getDefault[string](key)
	}

	return val
}

func GetInt(key string) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		return getDefault[int](key)
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		// TODO: debug error
		return getDefault[int](key)
	}

	return intVal
}

func GetBool(key string) bool {
	val, exists := os.LookupEnv(key)
	if !exists {
		return getDefault[bool](key)
	}

	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return getDefault[bool](key)
	}

	return boolVal
}

func SetDefaults(values map[string]any) {
	for key, val := range values {
		defaults[key] = val
	}
}

func getDefault[T any](key string) T {
	val, ok := defaults[key]
	if !ok {
		return *new(T)
	}

	cast, ok := val.(T)
	if !ok {
		return *new(T)
	}

	return cast
}
