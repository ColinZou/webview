#include "webview.h"
#include <stddef.h>

#ifdef _WIN32
#include <Windows.h>
#endif
#define UNUSED(x) (void)x
webview_t w = NULL;
void increment(const char *seq, const char *req, void *arg) {
  UNUSED(req);
  char *url = (char *)arg;
  webview_navigate_popup(w, url, 1);
  webview_return(w, seq, 0, "");
}
#ifdef _WIN32
//int WINAPI WinMain(HINSTANCE hInt, HINSTANCE hPrevInst, LPSTR lpCmdLine,
//                   int nCmdShow) {
int main() {
#else
int main() {
#endif
  char url_arg[] = "http://127.0.0.1:3030/child.html";
  w = webview_create(1, NULL);
  webview_set_title(w, "Basic Example");
  webview_set_size(w, 480, 320, WEBVIEW_HINT_NONE);
  webview_bind(w, "openPopup", increment, &url_arg);
  webview_navigate(w, "http://127.0.0.1:3030/webview.html");
  webview_run(w);
  webview_destroy(w);
  return 0;
}
