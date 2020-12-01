#ifndef __GO_BBS_SHM_H__
#define __GO_BBS_SHM_H__

#include <sys/shm.h>
#include <string.h>
#include <errno.h>

int isEExist();
int isPtrLessThan0(void *p);
void readwrapper(void *outptr, void *shmaddr, int offset, unsigned long n);
void writewrapper(void *shmaddr, int offset, void *inptr, unsigned long n);
void incuint32wrapper(void *shmaddr, int offset);
void set_or_uint32wrapper(void *shmaddr, int offset, void *inptr);
void innerset_int32wrapper(void *shmaddr, int offsetSrc, int offsetDst);
int cmpwrapper(void *shmaddr, int offset, unsigned long n, void *cmpaddr);
void memsetwrapper(void *shmaddr, int offset, unsigned char c, unsigned long n);

#endif //__GO_BBS_SHM_H__