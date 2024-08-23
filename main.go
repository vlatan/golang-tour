package main

import (
	"fmt"
	"golang-tour/basics"
	"golang-tour/concurrency"
	"golang-tour/flow"
	"golang-tour/generics"
	"golang-tour/interfaces"
	"golang-tour/methods"
	"golang-tour/types"
)

func RunBasics() {
	basics.SquareRoot(7)
	basics.RandomNumber(10)
	basics.PrintPi()

	fmt.Println(basics.Add(42, 13))
	fmt.Println(basics.Addition(42, 13))
	a, b := basics.Swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(basics.Split(17))

	basics.Vars()
	basics.VarsInit()
	basics.ShortVars()
	basics.BasicTypes()
	basics.ZeroValues()
	basics.TypeConversions()
	basics.TypeInference()
	basics.Constants()
	basics.NumericConstants()
}

func RunFlow() {
	flow.ForLoop()
	flow.WhileLoop()
	flow.ForeverLoop(1)
	fmt.Println(flow.IfStatement(2), flow.IfStatement((-4)))
	fmt.Println(flow.IfShort(3, 2, 10), flow.IfShort(3, 3, 20))
	fmt.Println(flow.IfElse(3, 2, 10), flow.IfElse(3, 3, 20))
	flow.ExcerciseLoops(6)
	flow.Switch()
	flow.SwitchEvaluationOrder()
	flow.SwitchNoCondition()
	flow.Defer()
	flow.DeferStacking()
}

func RunTypes() {
	types.Pointers()
	types.Structs()
	types.StructFields()
	types.StructPointers()
	types.StructLiterals()
	types.Array()
	types.Slice()
	types.SliceReference()
	types.SliceLiterals()
	types.SliceDefaults()
	types.SliceLengthCapacity()
	types.NilSlice()
	types.MakeSlice()
	types.SlicesOfSlices()
	types.AppendToSlice()
	types.Range()
	types.ExerciseSlices(20, 20)
	types.Map()
	types.MutatingMaps()
	types.ExerciseMaps("Fox Dog Fox")
	types.FunctionValues()
	types.FunctionClosures()
	types.ExcerciseClosure()
}

func RunMethods() {
	methods.Method()
	methods.PointerRecievers()
	methods.PointersFunctions()
	methods.PointerIndirection()
}

func RunInterfaces() {
	interfaces.Interfaces()
	interfaces.InterfaceIsImplicit()
	interfaces.InterfaceValues()
	interfaces.InterfaceNilUnderlyingValue()
	interfaces.NilIntrefaceValue()
	interfaces.EmptyInterface()
}

func RunTypeAssertions() {
	interfaces.TypeAssertions()
	interfaces.TypeSwitches(21)
	interfaces.TypeSwitches("hello")
	interfaces.TypeSwitches(true)
}

func RunStringers() {
	interfaces.Stringer()
	interfaces.ExerciseStringer()
}

func RunErrors() {
	interfaces.Error()
	interfaces.ExerciseErrors()
}

func RunReaders() {
	interfaces.Reader()
	interfaces.ExerciseReaders()
	interfaces.ExerciseRot1Reader()
}

func RunImages() {
	interfaces.Img()
	interfaces.ExerciseImages()
}

func RunGenerics() {
	generics.TypeParameters()
	generics.GenericTypes()
}

func RunConcurrency() {
	concurrency.Goroutine()
	concurrency.Channel()
	concurrency.BufferedChannel()
	concurrency.RangeClose()
	concurrency.Select()
	concurrency.DefaultSelect()
	concurrency.ExcerciseEquivalentBinaryTrees()
	concurrency.Mutex()
	concurrency.ExerciseWebCrawler()

}

func main() {
	RunBasics()
	RunFlow()
	RunTypes()
	RunMethods()
	RunInterfaces()
	RunTypeAssertions()
	RunStringers()
	RunErrors()
	RunReaders()
	RunImages()
	RunGenerics()
	RunConcurrency()
}
