# Build Instructions

```
go build .\
go build -buildmode=plugin -o eng\eng.so .\eng
```

Running
```
.\plugin_test.exe
```

Debugging
```
gdb plugin_test.exe
b go.link.addmoduledata
run
// At breakpoint you can start running si until segfault.
```

Known Issues
* Right now the data structure that should sit in local.moduledata the address of which is loaded in go.link.addmoduledata is corrupt
