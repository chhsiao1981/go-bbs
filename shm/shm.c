#include "shm.h"
#include <sys/shm.h>
#include <string.h>
#include <errno.h>
int isEExist() {
  return errno == EEXIST;
}
int isPtrLessThan0(void *p) {
  return p < 0;
}
void readwrapper(void *outptr, void *shmaddr, int offset, unsigned long n) {
  unsigned char *src = (unsigned char *)shmaddr + offset;
  memcpy(outptr, src, n);
}
void writewrapper(void *shmaddr, int offset, void *inptr, unsigned long n) {
  unsigned char *dst = (unsigned char *)shmaddr + offset;
  memcpy(dst, inptr, n);
}
void incuint32wrapper(void *shmaddr, int offset) {
  unsigned char *dst_b = (unsigned char *)shmaddr + offset;
  unsigned int *dst = (unsigned int *)dst_b;
  (*dst)++;
}
void set_or_uint32wrapper(void *shmaddr, int offset, void *inptr) {
    unsigned char *c_dst = (unsigned char *)shmaddr + offset;
    unsigned int *dst = (unsigned int *)c_dst;
    unsigned int *flag = (unsigned int *)inptr;
    *dst |= *flag;
}
void innerset_int32wrapper(void *shmaddr, int offsetSrc, int offsetDst) {
    unsigned char *c_src = (unsigned char *)shmaddr + offsetSrc;
    unsigned char *c_dst = (unsigned char *)shmaddr + offsetDst;
    int *src = (int *)c_src;
    int *dst = (int *)c_dst;
    *dst = *src;
}
int cmpwrapper(void *shmaddr, int offset, unsigned long n, void *cmpaddr) {
  unsigned char *cmp1 = (unsigned char *)shmaddr + offset;
  return memcmp(cmp1, cmpaddr, n);
}
void memsetwrapper(void *shmaddr, int offset, unsigned char c, unsigned long n) {
  unsigned char *dst = (unsigned char *)shmaddr + offset;
  memset(dst, c, n);
}
