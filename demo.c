#include <stdio.h>
#include <resolv.h>
#include <arpa/nameser.h>

int main() {
  printf("sizeof(struct __res_state)=%zd\n", sizeof(struct __res_state));

  const size_t answer_buffer_len = 1024;
  char buffer[answer_buffer_len];
  int result = res_query("google.com", C_IN, T_PTR, buffer, answer_buffer_len);
  printf("res_query result=%d\n", result);
}

