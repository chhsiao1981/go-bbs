#ifndef __GO_BBS_SHM_H__
#define __GO_BBS_SHM_H__

#include <sys/shm.h>
#include <string.h>
#include <errno.h>
#include <stdlib.h>

#define OFFSET_BOARD_HEADER_BRDNAME 0 //brd-name: IDLEN+1 (12+1)
#define SIZE_BOARD_HEADER_BRDNAME 13

#define OFFSET_BOARD_HEADER_TITLE 13 //title: BTLEN+1 (48+1)
#define SIZE_BOARD_HEADER_TITLE 49

#define SIZE_BOARD_HEADER 256

int isEExist();
int isPtrLessThan0(void *p);
void readwrapper(void *outptr, void *shmaddr, int offset, unsigned long n);
void writewrapper(void *shmaddr, int offset, void *inptr, unsigned long n);
void incuint32wrapper(void *shmaddr, int offset);
void set_or_uint32wrapper(void *shmaddr, int offset, void *inptr);
void innerset_int32wrapper(void *shmaddr, int offsetSrc, int offsetDst);
int cmpwrapper(void *shmaddr, int offset, unsigned long n, void *cmpaddr);
void memsetwrapper(void *shmaddr, int offset, unsigned char c, unsigned long n);

void set_bcacheptr(void *shmaddr, int offset);
void qsort_cmpboardname_wrapper(void *shmaddr, int offset, unsigned long n, unsigned long sz);
void qsort_cmpboardclass_wrapper(void *shmaddr, int offset, unsigned long n, unsigned long sz);
int cmpboardname(const void * i, const void * j);
int cmpboardclass(const void * i, const void * j);

#endif //__GO_BBS_SHM_H__
