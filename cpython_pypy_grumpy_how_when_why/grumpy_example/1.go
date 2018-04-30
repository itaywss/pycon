make: Nothing to be done for `all'.
package __main__
import πg "grumpy"
var Code *πg.Code
func init() {
	Code = πg.NewCode("<module>", "../grumpy_example/1.py", nil, 0, func(πF *πg.Frame, _ []*πg.Object) (*πg.Object, *πg.BaseException) {
		var πR *πg.Object; _ = πR
		var πE *πg.BaseException; _ = πE
		ßThread := πg.InternStr("Thread")
		ß__main__ := πg.InternStr("__main__")
		ß__name__ := πg.InternStr("__name__")
		ßfib := πg.InternStr("fib")
		ßjoin := πg.InternStr("join")
		ßmain := πg.InternStr("main")
		ßnum_of_threads := πg.InternStr("num_of_threads")
		ßrange := πg.InternStr("range")
		ßstart := πg.InternStr("start")
		ßstr := πg.InternStr("str")
		ßthreading := πg.InternStr("threading")
		ßtime := πg.InternStr("time")
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
		var πTemp007 bool
		_ = πTemp007
		for ; πF.State() >= 0; πF.PopCheckpoint() {
			switch πF.State() {
			case 0:
			default: panic("unexpected function state")
			}
			// line 2: import time
			πF.SetLineno(2)
			if πTemp002, πE = πg.ImportModule(πF, "time"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßtime.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 3: import threading
			πF.SetLineno(3)
			if πTemp002, πE = πg.ImportModule(πF, "threading"); πE != nil {
				continue
			}
			πTemp001 = πTemp002[0]
			if πE = πF.Globals().SetItem(πF, ßthreading.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 6: def fib(n):
			πF.SetLineno(6)
			πTemp003 = make([]πg.Param, 1)
			πTemp003[0] = πg.Param{Name: "n", Def: nil}
			πTemp001 = πg.NewFunction(πg.NewCode("fib", "../grumpy_example/1.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µn *πg.Object = πArgs[0]; _ = µn
				var πTemp001 *πg.Object
				_ = πTemp001
				var πTemp002 bool
				_ = πTemp002
				var πTemp003 []*πg.Object
				_ = πTemp003
				var πTemp004 *πg.Object
				_ = πTemp004
				var πTemp005 *πg.Object
				_ = πTemp005
				var πTemp006 *πg.Object
				_ = πTemp006
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					default: panic("unexpected function state")
					}
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp001, πE = πg.LE(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					if πTemp002, πE = πg.IsTrue(πF, πTemp001); πE != nil {
						continue
					}
					if πTemp002 {
						goto Label1
					}
					goto Label2
					// line 7: if n <= 1:
					πF.SetLineno(7)
				Label1:
					// line 8: return n
					πF.SetLineno(8)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					πR = µn
					continue
					goto Label2
				Label2:
					// line 9: return fib(n - 1) + fib(n - 2)
					πF.SetLineno(9)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µn, πg.NewInt(1).ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfib); πE != nil {
						continue
					}
					if πTemp005, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					πTemp003 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µn, "n"); πE != nil {
						continue
					}
					if πTemp004, πE = πg.Sub(πF, µn, πg.NewInt(2).ToObject()); πE != nil {
						continue
					}
					πTemp003[0] = πTemp004
					if πTemp004, πE = πg.ResolveGlobal(πF, ßfib); πE != nil {
						continue
					}
					if πTemp006, πE = πTemp004.Call(πF, πTemp003, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp003)
					if πTemp001, πE = πg.Add(πF, πTemp005, πTemp006); πE != nil {
						continue
					}
					πR = πTemp001
					continue
				}
				if πE != nil {
					πR = nil
				} else if πR == nil {
					πR = πg.None
				}
				return πR, πE
			}), πF.Globals()).ToObject()
			if πE = πF.Globals().SetItem(πF, ßfib.ToObject(), πTemp001); πE != nil {
				continue
			}
			// line 12: num_of_threads = 600
			πF.SetLineno(12)
			if πE = πF.Globals().SetItem(πF, ßnum_of_threads.ToObject(), πg.NewInt(600).ToObject()); πE != nil {
				continue
			}
			// line 17: def main():
			πF.SetLineno(17)
			πTemp003 = make([]πg.Param, 0)
			πTemp004 = πg.NewFunction(πg.NewCode("main", "../grumpy_example/1.py", πTemp003, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
				var µfib_num *πg.Object = πg.UnboundLocal; _ = µfib_num
				var µthreads *πg.Object = πg.UnboundLocal; _ = µthreads
				var µstart_time *πg.Object = πg.UnboundLocal; _ = µstart_time
				var µend_time *πg.Object = πg.UnboundLocal; _ = µend_time
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
				var πTemp007 []πg.Param
				_ = πTemp007
				var πTemp008 *πg.Object
				_ = πTemp008
				var πTemp009 *πg.Object
				_ = πTemp009
				var πTemp010 *πg.Object
				_ = πTemp010
				var πTemp011 *πg.Object
				_ = πTemp011
				var πTemp012 *πg.Object
				_ = πTemp012
				var πTemp013 *πg.Object
				_ = πTemp013
				var πTemp014 []*πg.Object
				_ = πTemp014
				var πTemp015 *πg.Object
				_ = πTemp015
				var πTemp016 *πg.Object
				_ = πTemp016
				var πR *πg.Object; _ = πR
				var πE *πg.BaseException; _ = πE
				for ; πF.State() >= 0; πF.PopCheckpoint() {
					switch πF.State() {
					case 0:
					case 1: goto Label1
					case 2: goto Label2
					default: panic("unexpected function state")
					}
					πTemp002 = πF.MakeArgs(2)
					πTemp002[0] = πg.NewInt(1).ToObject()
					πTemp002[1] = πg.NewInt(20).ToObject()
					if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
						µfib_num = πTemp003
					}
					if πE != nil || !πTemp006 {
						continue
					}
					πF.PushCheckpoint(1)            
					// line 20: threads = [threading.Thread(target=fib, args=(fib_num,)) for _ in range(num_of_threads)]
					πF.SetLineno(20)
					πTemp007 = make([]πg.Param, 0)
					πTemp004 = πg.NewFunction(πg.NewCode("<generator>", "../grumpy_example/1.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
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
						var πTemp007 πg.KWArgs
						_ = πTemp007
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								πTemp002 = πF.MakeArgs(1)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßnum_of_threads); πE != nil {
									continue
								}
								πTemp002[0] = πTemp003
								if πTemp003, πE = πg.ResolveGlobal(πF, ßrange); πE != nil {
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
								// line 20: threads = [threading.Thread(target=fib, args=(fib_num,)) for _ in range(num_of_threads)]
								πF.SetLineno(20)
								if πTemp003, πE = πg.ResolveGlobal(πF, ßfib); πE != nil {
									continue
								}
								if πE = πg.CheckLocal(πF, µfib_num, "fib_num"); πE != nil {
									continue
								}
								πTemp004 = πg.NewTuple1(µfib_num).ToObject()
								πTemp007 = πg.KWArgs{
									{"target", πTemp003},
									{"args", πTemp004},
								}
								if πTemp003, πE = πg.ResolveGlobal(πF, ßthreading); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, πTemp003, ßThread, nil); πE != nil {
									continue
								}
								if πTemp003, πE = πTemp004.Call(πF, nil, πTemp007); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp003, nil
							Label4:
								πTemp004 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp008, πE = πTemp004.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp008}, nil); πE != nil {
						continue
					}
					µthreads = πTemp003
					// line 23: start_time = time.time()
					πF.SetLineno(23)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp008, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					µstart_time = πTemp003
					// line 26: [thread.start() for thread in threads]
					πF.SetLineno(26)
					πTemp007 = make([]πg.Param, 0)
					πTemp008 = πg.NewFunction(πg.NewCode("<generator>", "../grumpy_example/1.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µthread *πg.Object = πg.UnboundLocal; _ = µthread
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µthreads); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µthread = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 26: [thread.start() for thread in threads]
								πF.SetLineno(26)
								if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µthread, ßstart, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp005, nil
							Label4:
								πTemp004 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp009, πE = πTemp008.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp009}, nil); πE != nil {
						continue
					}
					// line 29: [thread.join() for thread in threads]
					πF.SetLineno(29)
					πTemp007 = make([]πg.Param, 0)
					πTemp009 = πg.NewFunction(πg.NewCode("<generator>", "../grumpy_example/1.py", πTemp007, 0, func(πF *πg.Frame, πArgs []*πg.Object) (*πg.Object, *πg.BaseException) {
						var µthread *πg.Object = πg.UnboundLocal; _ = µthread
						var πTemp001 *πg.Object
						_ = πTemp001
						var πTemp002 bool
						_ = πTemp002
						var πTemp003 bool
						_ = πTemp003
						var πTemp004 *πg.Object
						_ = πTemp004
						var πTemp005 *πg.Object
						_ = πTemp005
						var πR *πg.Object; _ = πR
						var πE *πg.BaseException; _ = πE
						return πg.NewGenerator(πF, func(πSent *πg.Object) (*πg.Object, *πg.BaseException) {
							for ; πF.State() >= 0; πF.PopCheckpoint() {
								switch πF.State() {
								case 0:
								case 1: goto Label1
								case 2: goto Label2
								case 4: goto Label4
								default: panic("unexpected function state")
								}
								if πE = πg.CheckLocal(πF, µthreads, "threads"); πE != nil {
									continue
								}
								if πTemp001, πE = πg.Iter(πF, µthreads); πE != nil {
									continue
								}
								πF.PushCheckpoint(2)
								πTemp002 = false
							Label1:
								if πE != nil || πR != nil {
									continue
								}
								if πTemp002 {
									πF.PopCheckpoint()
									goto Label3
								}
								if πTemp004, πE = πg.Next(πF, πTemp001); πE != nil {
									isStop, exc := πg.IsInstance(πF, πE.ToObject(), πg.StopIterationType.ToObject())
									if exc != nil {
										πE = exc
									} else if isStop {
										πE = nil
										πF.RestoreExc(nil, nil)
									}
									πTemp003 = !isStop
								} else {
									πTemp003 = true
									µthread = πTemp004
								}
								if πE != nil || !πTemp003 {
									continue
								}
								πF.PushCheckpoint(1)            
								// line 29: [thread.join() for thread in threads]
								πF.SetLineno(29)
								if πE = πg.CheckLocal(πF, µthread, "thread"); πE != nil {
									continue
								}
								if πTemp004, πE = πg.GetAttr(πF, µthread, ßjoin, nil); πE != nil {
									continue
								}
								if πTemp005, πE = πTemp004.Call(πF, nil, nil); πE != nil {
									continue
								}
								πF.PushCheckpoint(4)
								return πTemp005, nil
							Label4:
								πTemp004 = πSent
								continue
							Label2:
								if πE != nil || πR != nil {
									continue
								}
							Label3:
							}
							return nil, πE
						}).ToObject(), nil
					}), πF.Globals()).ToObject()
					if πTemp010, πE = πTemp009.Call(πF, nil, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πg.ListType.Call(πF, πg.Args{πTemp010}, nil); πE != nil {
						continue
					}
					// line 31: end_time = time.time()
					πF.SetLineno(31)
					if πTemp003, πE = πg.ResolveGlobal(πF, ßtime); πE != nil {
						continue
					}
					if πTemp010, πE = πg.GetAttr(πF, πTemp003, ßtime, nil); πE != nil {
						continue
					}
					if πTemp003, πE = πTemp010.Call(πF, nil, nil); πE != nil {
						continue
					}
					µend_time = πTemp003
					// line 33: print('Number Of Threads: ' + str(num_of_threads) + ' Fib Number: ' + str(fib_num) + ' Time: ' + str(end_time - start_time))
					πF.SetLineno(33)
					πTemp002 = make([]*πg.Object, 1)
					πTemp014 = πF.MakeArgs(1)
					if πTemp015, πE = πg.ResolveGlobal(πF, ßnum_of_threads); πE != nil {
						continue
					}
					πTemp014[0] = πTemp015
					if πTemp015, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp016, πE = πTemp015.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πTemp013, πE = πg.Add(πF, πg.NewStr("Number Of Threads: ").ToObject(), πTemp016); πE != nil {
						continue
					}
					if πTemp012, πE = πg.Add(πF, πTemp013, πg.NewStr(" Fib Number: ").ToObject()); πE != nil {
						continue
					}
					πTemp014 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µfib_num, "fib_num"); πE != nil {
						continue
					}
					πTemp014[0] = µfib_num
					if πTemp013, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp015, πE = πTemp013.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πTemp011, πE = πg.Add(πF, πTemp012, πTemp015); πE != nil {
						continue
					}
					if πTemp010, πE = πg.Add(πF, πTemp011, πg.NewStr(" Time: ").ToObject()); πE != nil {
						continue
					}
					πTemp014 = πF.MakeArgs(1)
					if πE = πg.CheckLocal(πF, µend_time, "end_time"); πE != nil {
						continue
					}
					if πE = πg.CheckLocal(πF, µstart_time, "start_time"); πE != nil {
						continue
					}
					if πTemp011, πE = πg.Sub(πF, µend_time, µstart_time); πE != nil {
						continue
					}
					πTemp014[0] = πTemp011
					if πTemp011, πE = πg.ResolveGlobal(πF, ßstr); πE != nil {
						continue
					}
					if πTemp012, πE = πTemp011.Call(πF, πTemp014, nil); πE != nil {
						continue
					}
					πF.FreeArgs(πTemp014)
					if πTemp003, πE = πg.Add(πF, πTemp010, πTemp012); πE != nil {
						continue
					}
					πTemp002[0] = πTemp003
					if πE = πg.Print(πF, πTemp002, true); πE != nil {
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
			if πE = πF.Globals().SetItem(πF, ßmain.ToObject(), πTemp004); πE != nil {
				continue
			}
			if πTemp006, πE = πg.ResolveGlobal(πF, ß__name__); πE != nil {
				continue
			}
			if πTemp005, πE = πg.Eq(πF, πTemp006, ß__main__.ToObject()); πE != nil {
				continue
			}
			if πTemp007, πE = πg.IsTrue(πF, πTemp005); πE != nil {
				continue
			}
			if πTemp007 {
				goto Label1
			}
			goto Label2
			// line 36: if __name__ == '__main__':
			πF.SetLineno(36)
		Label1:
			// line 37: main()
			πF.SetLineno(37)
			if πTemp005, πE = πg.ResolveGlobal(πF, ßmain); πE != nil {
				continue
			}
			if πTemp006, πE = πTemp005.Call(πF, nil, nil); πE != nil {
				continue
			}
			goto Label2
		Label2:
		}
		return nil, πE
	})
	πg.RegisterModule("__main__", Code)
}
