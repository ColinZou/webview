package webview

/*
#include "native/webview_models.h"
#include "native/webview_methods.h"
#include "native/webview_go_glue.h"

#include <stdlib.h>
#include <stdint.h>

extern void goOnChildWindowCreatedCallback(int, webview_t);
extern void goOnChildWindowClosedCallback(int);
extern char* _goWebviewNativeCallback(int, char *);

extern void _webviewDispatchGoCallback(void*);
extern void _webviewBindingGoCallback(webview_t, char *, char *,uintptr_t);

void c_goOnChildWindowCreatedCallback(int window_id, webview_t window){
	goOnChildWindowCreatedCallback(window_id, window);
}
void c_webviewDispatchGoCallback(void* w) {
	_webviewDispatchGoCallback(w);
}
void c_webviewBindingGoCallback(webview_t w, char * c1, char * c2, uintptr_t p1){
	_webviewBindingGoCallback(w, c1, c2, p1);
}
void c_goOnChildWindowClosedCallback(int windowId) {
	goOnChildWindowClosedCallback(windowId);
}
char* c_goWebviewNativeCallback(int window_id, char *args) {
	return _goWebviewNativeCallback(window_id, args);
}
*/
import "C"
