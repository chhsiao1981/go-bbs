package shm

// #include <sys/shm.h>
// #include <string.h>
// #include <errno.h>
// int isEExist() {
//   return errno == EEXIST;
// }
// int isPtrLessThan0(void *p) {
//   return p < 0;
// }
// void readwrapper(void *outptr, void *shmaddr, int offset, unsigned long n) {
//   unsigned char *src = (unsigned char *)shmaddr + offset;
//   memcpy(outptr, src, n);
// }
// void writewrapper(void *shmaddr, int offset, void *inptr, unsigned long n) {
//   unsigned char *dst = (unsigned char *)shmaddr + offset;
//   memcpy(dst, inptr, n);
// }
// void incuint32wrapper(void *shmaddr, int offset) {
//   unsigned char *dst_b = (unsigned char *)shmaddr + offset;
//   unsigned int *dst = (unsigned int *)dst_b;
//   (*dst)++;
// }
// int cmpwrapper(void *shmaddr, int offset, unsigned long n, void *cmpaddr) {
//   unsigned char *cmp1 = (unsigned char *)shmaddr + offset;
//   return memcmp(cmp1, cmpaddr, n);
// }
import "C"
import (
	"unsafe"

	"github.com/PichuChen/go-bbs/types"
	log "github.com/sirupsen/logrus"
)

func CreateShm(key types.Key_t, size types.Size_t, isUseHugeTlb bool) (shmid int, shmaddr unsafe.Pointer, isNew bool, err error) {
	flags := 0600 | IPC_CREAT | IPC_EXCL
	if isUseHugeTlb {
		flags |= SHM_HUGETLB
	}
	shmid = shmget(key, size, flags)

	isEExist := int(C.isEExist()) != 0
	if isEExist {
		flags = 0600 | IPC_CREAT
		if isUseHugeTlb {
			flags |= SHM_HUGETLB
		}
		shmid = shmget(key, size, flags)
	}
	if shmid < 0 {
		log.Errorf("shm.CreateShm: unable to create shm: key: %v size: %v", key, size)
		return shmid, nil, false, ErrInvalidShm
	}

	shmaddr, err = shmat(shmid, nil, 0)
	if err != nil {
		return -1, nil, false, err
	}

	return shmid, shmaddr, !isEExist, nil
}

func OpenShm(key types.Key_t, size types.Size_t, is_usehugetlb bool) (shmid int, shmaddr unsafe.Pointer, err error) {
	flags := 0
	if is_usehugetlb {
		flags |= SHM_HUGETLB
	}
	shmid = shmget(key, size, flags)

	if shmid < 0 {
		log.Errorf("shm.OpenShm: unable to create shm: key: %v size: %v", key, size)
		return shmid, nil, ErrInvalidShm
	}

	shmaddr, err = shmat(shmid, nil, 0)
	if err != nil {
		return -1, nil, err
	}

	return shmid, shmaddr, nil
}

func CloseShm(shmid int) (err error) {
	cerrno := C.shmctl(C.int(shmid), C.IPC_RMID, nil)

	log.Infof("After close shm: errno: %v", cerrno)

	if int(cerrno) < 0 {
		return ErrUnableToCloseShm
	}

	return nil
}

func ReadAt(shmaddr unsafe.Pointer, offset int, size types.Size_t, outptr unsafe.Pointer) {

	C.readwrapper(outptr, shmaddr, C.int(offset), C.ulong(size))

	return
}

func WriteAt(shmaddr unsafe.Pointer, offset int, size types.Size_t, inptr unsafe.Pointer) {

	C.writewrapper(shmaddr, C.int(offset), inptr, C.ulong(size))

	return
}

func IncUint32(shmaddr unsafe.Pointer, offset int) {
	C.incuint32wrapper(shmaddr, C.int(offset))

	return
}

//Cmp
//
//memcmp(shmaddr+offset, cmpaddr, size)
//
//Return:
//	int: 0: shm == gomem, <0: shm < gomem, >0: shm > gomem
func Cmp(shmaddr unsafe.Pointer, offset int, size types.Size_t, cmpaddr unsafe.Pointer) int {
	cret := C.cmpwrapper(shmaddr, C.int(offset), C.ulong(size), cmpaddr)

	return int(cret)
}

func shmget(key types.Key_t, size types.Size_t, shmflg int) int {
	cshmid := C.shmget(C.int(key), C.ulong(size), C.int(shmflg))
	return int(cshmid)
}

func shmat(shmid int, shmaddr unsafe.Pointer, shmflg int) (unsafe.Pointer, error) {
	shmaddr = C.shmat(C.int(shmid), shmaddr, C.int(shmflg))
	if int(C.isPtrLessThan0(shmaddr)) != 0 {
		return nil, ErrUnableToAttachShm
	}

	return shmaddr, nil
}
