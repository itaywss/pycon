package __main__
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "benchmarks/dict.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßBenchmarkDictCreate := πg.InternStr("BenchmarkDictCreate")
		ßBenchmarkDictCreateFunc := πg.InternStr("BenchmarkDictCreateFunc")
		ßBenchmarkDictGetItem := πg.InternStr("BenchmarkDictGetItem")
		ßBenchmarkDictSetItem := πg.InternStr("BenchmarkDictSetItem")
		ßBenchmarkDictStringOnlyGetItem := πg.InternStr("BenchmarkDictStringOnlyGetItem")
		ßBenchmarkDictStringOnlySetItem := πg.InternStr("BenchmarkDictStringOnlySetItem")
		ßBenchmarkHashStrCached := πg.InternStr("BenchmarkHashStrCached")
		ßN := πg.InternStr("N")
		ßRunBenchmarks := πg.InternStr("RunBenchmarks")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßdict := πg.InternStr("dict")
		ßfoo := πg.InternStr("foo")
		ßfoobarfoobarfoobar := πg.InternStr("foobarfoobarfoobar")
		ßhash := πg.InternStr("hash")
		ßone := πg.InternStr("one")
		ßthree := πg.InternStr("three")
		ßtwo := πg.InternStr("two")
		ßweetest := πg.InternStr("weetest")
		ßxrange := πg.InternStr("xrange")
		var πTemp001 *πg.Object
		_ = πTemp001
		var πTemp002 []*πg.Object
		_ = πTemp002
		var πTemp003 []πg.Param
		_ = πTemp003
		var πTemp004 *πg.Object
		_ = πTemp004
		var πTemp005 *πg.Object
		_ = πTemp005
		var πTemp006 *πg.Object
		_ = πTemp006
		var πTemp007 *πg.Object
		_ = πTemp007
		var πTemp008 *πg.Object
		_ = πTemp008
		var πTemp009 *πg.Object
		_ = πTemp009
		var πTemp010 *πg.Object
		_ = πTemp010
		var πTemp011 *πg.Object
		_ = πTemp011
		var πTemp012 bool
		_ = πTemp012
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 15: """Benchmarks for dictionary operations."""
			πF.SetLineno(15)
			// line 19: import weetest
			πF.SetLineno(19)
			if πTemp002, πE = πg.ImportModule(πF, "weetest"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßweetest.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 22: def BenchmarkDictCreate(b):
			πF.SetLineno(22)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("BenchmarkDictCreate", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 *πg.Dict
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp005 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µ_ = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 24: d = {'one': 1, 'two': 2, 'three': 3}
					πF.SetLineno(24)
					πTemp007 = πg.NewDict()
					if πE = πTemp007.SetItem(πF, ßone.ToObject(), πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßtwo.ToObject(), πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					if πE = πTemp007.SetItem(πF, ßthree.ToObject(), πg.NewInt(3).ToObject()); πE != nil {
						continue
					}
					πTemp003 = πTemp007.ToObject()
					µd = πTemp003
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkDictCreate.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 27: def BenchmarkDictCreateFunc(b):
			πF.SetLineno(27)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp004 = πg.NewFunction(πg.NewCode("BenchmarkDictCreateFunc", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 πg.KWArgs
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp005 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µ_ = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 29: d = dict(one=1, two=2, three=3)
					πF.SetLineno(29)
					πTemp007 = πg.KWArgs{
						{"one", πg.NewInt(1).ToObject()},
						{"two", πg.NewInt(2).ToObject()},
						{"three", πg.NewInt(3).ToObject()},
					}
					if πTemp003, πE = πg.ResolveGlobal(πF, ßdict); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, nil, πTemp007); πE != nil {
						continue
					}
					µd = πTemp004
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkDictCreateFunc.ToObject(), πTemp004); πE != nil {
				continue
			}
			// line 32: def BenchmarkDictGetItem(b):
			πF.SetLineno(32)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp005 = πg.NewFunction(πg.NewCode("BenchmarkDictGetItem", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 33: d = {42: 123}
					πF.SetLineno(33)
					πTemp001 = πg.NewDict()
					if πE = πTemp001.SetItem(πF, πg.NewInt(42).ToObject(), πg.NewInt(123).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					µd = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µ_ = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 35: d[42]
					πF.SetLineno(35)
					πTemp004 = πg.NewInt(42).ToObject()
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µd, πTemp004); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkDictGetItem.ToObject(), πTemp005); πE != nil {
				continue
			}
			// line 38: def BenchmarkDictStringOnlyGetItem(b):
			πF.SetLineno(38)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp006 = πg.NewFunction(πg.NewCode("BenchmarkDictStringOnlyGetItem", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 39: d = {'foo': 123}
					πF.SetLineno(39)
					πTemp001 = πg.NewDict()
					if πE = πTemp001.SetItem(πF, ßfoo.ToObject(), πg.NewInt(123).ToObject()); πE != nil {
						continue
					}
					πTemp002 = πTemp001.ToObject()
					µd = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µ_ = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 41: d['foo']
					πF.SetLineno(41)
					πTemp004 = ßfoo.ToObject()
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					if πTemp005, πE = πg.GetItem(πF, µd, πTemp004); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkDictStringOnlyGetItem.ToObject(), πTemp006); πE != nil {
				continue
			}
			// line 44: def BenchmarkDictSetItem(b):
			πF.SetLineno(44)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp007 = πg.NewFunction(πg.NewCode("BenchmarkDictSetItem", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 45: d = {}
					πF.SetLineno(45)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µd = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µ_ = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 47: d[42] = 123
					πF.SetLineno(47)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewInt(123).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp005 = πg.NewInt(42).ToObject()
					if πE = πg.SetItem(πF, µd, πTemp005, πTemp004); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkDictSetItem.ToObject(), πTemp007); πE != nil {
				continue
			}
			// line 50: def BenchmarkDictStringOnlySetItem(b):
			πF.SetLineno(50)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp008 = πg.NewFunction(πg.NewCode("BenchmarkDictStringOnlySetItem", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µd *πg.Object = πg.UnboundLocal; _ = µd
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Dict
				_ = πTemp001
				var πTemp002 *πg.Object
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πTemp007 bool
				_ = πTemp007
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 51: d = {}
					πF.SetLineno(51)
					πTemp001 = πg.NewDict()
					πTemp002 = πTemp001.ToObject()
					µd = πTemp002
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp002, πE = πg.Iter(πF, πTemp005); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp006 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp006 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp004, πE = πg.Next(πF, πTemp002); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp007 = !isStop
					} else {
						πTemp007 = true
						µ_ = πTemp004
					}
					if πE != nil || !πTemp007 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 53: d['foo'] = 123
					πF.SetLineno(53)
					if πE = πg.Tie(πF, πg.TieTarget{Target: &πTemp004}, πg.NewInt(123).ToObject()); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µd, "d"); πE != nil {
						continue
					}
					πTemp005 = ßfoo.ToObject()
					if πE = πg.SetItem(πF, µd, πTemp005, πTemp004); πE != nil {
						continue
					}
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkDictStringOnlySetItem.ToObject(), πTemp008); πE != nil {
				continue
			}
			// line 56: def BenchmarkHashStrCached(b):
			πF.SetLineno(56)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "b", Def: nil}
			πTemp009 = πg.NewFunction(πg.NewCode("BenchmarkHashStrCached", "benchmarks/dict.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µb *πg.Object = πArgs[0]; _ = µb
				var µh *πg.Object = πg.UnboundLocal; _ = µh
				var µ_ *πg.Object = πg.UnboundLocal; _ = µ_
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 []*πg.Object
				_ = πTemp002
				var πTemp003 *πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 bool
				_ = πTemp005
				var πTemp006 bool
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					// line 57: """Hashes the same value repeatedly to exercise any hash caching logic."""
					πF.SetLineno(57)
					// line 58: h = hash  # Prevent builtins lookup each iteration.
					πF.SetLineno(58)
					if πTemp001, πE = πg.ResolveGlobal(πF, ßhash); πE != nil {
						continue
					}
					µh = πTemp001
					πTemp002 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µb, "b"); πE != nil {
						continue
					}
					if πTemp003, πE = πg.GetAttr(πF, µb, ßN, nil); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πTemp003, πE = πg.ResolveGlobal(πF, ßxrange); πE != nil {
						continue
					}
					if πTemp004, πE = πTemp003.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					if πTemp001, πE = πg.Iter(πF, πTemp004); πE != nil {
						continue
					}
					πF.PushCheckpoint(2)
					πTemp005 = false
				Label1:
					if πE != nil || πR != nil {
						continue
					}
					if πTemp005 {
						πF.PopCheckpoint()
						goto Label3
					}
					if πTemp003, πE = πg.Next(πF, πTemp001); πE != nil {
						isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
						if exc != nil {
							πE = exc
						} else if isStop {
							πE = nil
							πF.RestoreExc(nil, nil)
						}
						πTemp006 = !isStop
					} else {
						πTemp006 = true
						µ_ = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 60: h('foobarfoobarfoobar')
					πF.SetLineno(60)
					πTemp002 = πF.MakeArgs(1)
					πTemp002[0] = ßfoobarfoobarfoobar.ToObject()
					if πE = πg.CheckLocal(πF, µh, "h"); πE != nil {
						continue
					}
					if πTemp003, πE = µh.Call(πF, πTemp002, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp002)
					continue
				Label2:
					if πE != nil || πR != nil {
						continue
					}
				Label3:
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßBenchmarkHashStrCached.ToObject(), πTemp009); πE != nil {
				continue
			}
			if πTemp011, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp010, πE = πg.Eq(πF, πTemp011, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp012, πE = πg.IsTrue(πF, πTemp010); πE != nil {
				continue
			}
			if πTemp012 {
				goto Label1
			}
			goto Label2
			// line 63: if __name__ == '__main__':
			πF.SetLineno(63)
		Label1:
			// line 64: weetest.RunBenchmarks()
			πF.SetLineno(64)
			if πTemp010, πE = πg.ResolveGlobal(πF, ßweetest); πE != nil {
				continue
			}
			if πTemp011, πE = πg.GetAttr(πF, πTemp010, ßRunBenchmarks, nil); πE != nil {
				continue
			}
			if πTemp010, πE = πTemp011.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
