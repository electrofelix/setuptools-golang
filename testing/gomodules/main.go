package main

// #include <Python.h>
//
// void reversemsg();
//
// static struct PyMethodDef methods[] = {
//     {"reversemsg", (PyCFunction)reversemsg, NULL},
//     {NULL, NULL}
// };
//
// #if PY_MAJOR_VERSION >= 3
// static struct PyModuleDef module = {
//     PyModuleDef_HEAD_INIT,
//     "gomodules",
//     NULL,
//     -1,
//     methods
// };
//
// PyMODINIT_FUNC PyInit_gomodules(void) {
//     return PyModule_Create(&module);
// }
// #else
// PyMODINIT_FUNC initgomodules(void) {
//     Py_InitModule3("gomodules", methods, NULL);
// }
// #endif
import "C"

import (
	"fmt"

	"github.com/golang/example/stringutil"
)

//export reversemsg
func reversemsg() {
    fmt.Println(stringutil.Reverse("elpmaxe tset"))
}

func main() {}
