package main

import (
	"jpath/input"
	"jpath/parser"
	"testing"
)

// Benchmark_Jpath benchmarking function - generates a flame graph of usage
// To use, put a large file in /tmp/output.json (or whatever filename)
// and then update a complex query in the test below
// Then run this command
// go test -bench=Jpath -cpuprofile=cpuprof_jpath.out
// To render the flame graph, upload cpuprof_jpath.out on this page https://www.speedscope.app/
// refer: https://sathishvj.medium.com/flamegraphs-for-code-optimization-with-golang-and-speedscope-80c20725fdd2
func Benchmark_Jpath(b *testing.B) {
	json, _ := input.ParseInputJson("/tmp/output.json")
	_, _ = parser.ProcessExpression(".event_context.container[location_id=3857].changed_values", json)
}
