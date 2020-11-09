#ifndef _BBSCRYPT_H_
#define _BBSCRYPT_H_

#ifdef PERL5
char *des_crypt(char *buf, char *salt);
#else
char *fcrypt(char *buf, char *salt);
#endif

#endif
